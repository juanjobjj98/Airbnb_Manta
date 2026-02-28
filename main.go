package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JJ/Airbnb_Manta/models"
	"github.com/JJ/Airbnb_Manta/repositories"
)

type ReservaEntradaJSON struct {
	Huesped      string  `json:"huesped"`
	Email        string  `json:"email"`
	Personas     int     `json:"personas"`
	Mascotas     bool    `json:"mascotas"`
	FechaIngreso string  `json:"fecha_ingreso"`
	FechaSalida  string  `json:"fecha_salida"`
	Precio       float64 `json:"precio"`
	Canal        string  `json:"canal"`
}

func main() {
	repositories.InicializarBaseDeDatos()

	http.Handle("/", http.FileServer(http.Dir("./static")))

	// 1. POST /api/reservas (Crear Reserva con validación de Overbooking)
	http.HandleFunc("POST /api/reservas", func(w http.ResponseWriter, r *http.Request) {
		var input ReservaEntradaJSON

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, `{"error": "JSON inválido"}`, http.StatusBadRequest)
			return
		}

		ingreso, _ := time.Parse("2006-01-02", input.FechaIngreso)
		salida, _ := time.Parse("2006-01-02", input.FechaSalida)
		canal := models.CreateChannel(input.Canal)

		reserva, err := models.NewReservation(
			input.Huesped, input.Email, input.Personas, input.Mascotas,
			ingreso, salida, input.Precio, canal)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		err = repositories.VerificarDisponibilidad(ingreso, salida)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		repositories.GuardarReservaSQL(reserva)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Reserva creada exitosamente"})
	})

	// 2. GET /api/reservas (Listar Todas)
	http.HandleFunc("GET /api/reservas", func(w http.ResponseWriter, r *http.Request) {
		reservas, _ := repositories.ObtenerTodasLasReservas()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reservas)
	})

	// 3. DELETE /api/reservas/{id} (Eliminar)
	http.HandleFunc("DELETE /api/reservas/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := repositories.EliminarReservaSQL(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "No se pudo eliminar"})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Eliminada con éxito"})
	})

	// 4. GET /api/reservas/{id}
	http.HandleFunc("GET /api/reservas/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Detalle de la reserva " + id})
	})

	// 5. GET /api/canales
	http.HandleFunc("GET /api/canales", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{"Airbnb", "Booking", "Directo"})
	})

	// 6. GET /api/kpis
	http.HandleFunc("GET /api/kpis", func(w http.ResponseWriter, r *http.Request) {
		reservas, _ := repositories.ObtenerTodasLasReservas()
		var totalNeto float64
		for _, res := range reservas {
			totalNeto += res.GananciaNeta
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"total_reservas": len(reservas),
			"ganancia_neta":  totalNeto,
		})
	})

	// 7. GET /api/ping
	http.HandleFunc("GET /api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"estado": "Activo"})
	})

	// 8. POST /api/reportes/async (CONCURRENCIA Y REPORTE PROFESIONAL PARA EXCEL)
	http.HandleFunc("POST /api/reportes/async", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"mensaje": "Generando reporte profesional en segundo plano. Revisa tu carpeta en 3 segundos.",
		})

		go func() {
			fmt.Println("\n[CONCURRENCIA] Extrayendo datos financieros de SQL Server...")
			time.Sleep(3 * time.Second)

			reservas, err := repositories.ObtenerTodasLasReservas()
			if err != nil {
				fmt.Println("[ERROR] No se pudo leer la base de datos.")
				return
			}

			timestamp := time.Now().Format("2006-01-02_15-04-05")
			nombreArchivo := fmt.Sprintf("Reporte_Financiero_%s.csv", timestamp)

			archivo, err := os.Create(nombreArchivo)
			if err != nil {
				fmt.Println("[ERROR] No se pudo crear el archivo.")
				return
			}
			defer archivo.Close()

			// TRUCO PRO 1: Escribir el BOM UTF-8 para que Excel lea bien los acentos
			archivo.WriteString("\xEF\xBB\xBF")

			escritor := csv.NewWriter(archivo)
			// TRUCO PRO 2: Cambiar el separador a PUNTO Y COMA para el Excel en español
			escritor.Comma = ';'
			defer escritor.Flush()

			// Escribimos las cabeceras (ahora con acentos sin miedo)
			escritor.Write([]string{"Huésped", "Email", "Fecha Ingreso", "Fecha Salida", "Canal de Venta", "Ganancia Neta ($)"})

			var sumaTotal float64
			for _, res := range reservas {
				escritor.Write([]string{
					res.Huesped,
					res.Email,
					res.FechaIngreso.Format("2006-01-02"),
					res.FechaSalida.Format("2006-01-02"),
					res.CanalVenta,
					// Formateamos el número para que no tenga problemas
					fmt.Sprintf("%.2f", res.GananciaNeta),
				})
				sumaTotal += res.GananciaNeta
			}

			// Fila final de totales (dejamos columnas en blanco para que cuadre al final)
			escritor.Write([]string{"", "", "", "", "TOTAL ACUMULADO:", fmt.Sprintf("%.2f", sumaTotal)})

			fmt.Printf("[CONCURRENCIA] ✅ ÉXITO: Archivo profesional '%s' generado.\n\n", nombreArchivo)
		}()
	})

	fmt.Println("🚀 Servidor Web RESTful (8 Servicios) Activo -> http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
