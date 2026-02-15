package main // Paquete principal requerido para construir el ejecutable.

import ( // Inicia bloque de importaciones.
	"bufio"   // Permite leer secuencias de texto complejas (con espacios) desde la consola.
	"fmt"     // Paquete para imprimir texto interactivo en pantalla.
	"log"     // Paquete para imprimir errores críticos y detener el programa si es necesario.
	"os"      // Proporciona acceso a las funciones del sistema operativo, como el teclado (Stdin).
	"strings" // Funciones para limpiar y procesar las cadenas de texto ingresadas.
	"time"    // Paquete para procesar y validar el texto convirtiéndolo a fechas reales.

	"github.com/JJ/Airbnb_Manta/models"       // Importa el paquete de negocio local.
	"github.com/JJ/Airbnb_Manta/repositories" // Importa el paquete de persistencia de datos local.
) // Cierra importaciones.

func main() { // Función principal donde inicia la ejecución del programa.
	fmt.Println("=========================================") // Imprime separador visual superior.
	fmt.Println("  SISTEMA DE GESTIÓN AIRBNB (Consola)    ") // Imprime el título del software.
	fmt.Println("=========================================") // Imprime separador visual inferior.

	// 1. Invoca la función que crea la base de datos automáticamente si no existe.
	err := repositories.InicializarBaseDeDatos() // Ejecuta la inicialización y atrapa un posible error.
	if err != nil {                              // Si ocurrió un error grave de conexión...
		log.Fatalf("❌ Error Crítico: %v\n", err) // ...imprime el error y aborta el programa inmediatamente.
	} // Cierra bloque if.
	fmt.Println("-----------------------------------------") // Separador visual.

	reader := bufio.NewReader(os.Stdin) // Crea un lector que captura todo lo escrito en el teclado hasta presionar 'Enter'.

	// 2. Bloque de ingreso interactivo: Nombre del huésped.
	fmt.Print("Ingrese el nombre del huésped: ") // Imprime la instrucción en la misma línea.
	huesped, _ := reader.ReadString('\n')        // Lee el texto ingresado incluyendo espacios hasta detectar el salto de línea.
	huesped = strings.TrimSpace(huesped)         // Limpia el texto eliminando espacios extra y el propio salto de línea invisible.

	formato := "2006-01-02" // Define el estándar estricto de formato de fecha en el lenguaje Go.

	// 3. Bloque de ingreso interactivo: Fecha de llegada.
	fmt.Print("Fecha de Ingreso (YYYY-MM-DD): ")                            // Solicita la fecha de entrada.
	strIngreso, _ := reader.ReadString('\n')                                // Captura lo tipeado en teclado.
	fechaIngreso, err := time.Parse(formato, strings.TrimSpace(strIngreso)) // Intenta transformar el texto ingresado en un objeto Time real.
	if err != nil {                                                         // Si el usuario escribe mal el formato (ej. 15-03-2026)...
		log.Fatal("❌ Formato inválido. Use YYYY-MM-DD (Ej: 2026-03-15)") // ...muestra el error y detiene el flujo.
	} // Cierra bloque if.

	// 4. Bloque de ingreso interactivo: Fecha de salida.
	fmt.Print("Fecha de Salida (YYYY-MM-DD): ")                           // Solicita la fecha de salida.
	strSalida, _ := reader.ReadString('\n')                               // Captura lo tipeado en teclado.
	fechaSalida, err := time.Parse(formato, strings.TrimSpace(strSalida)) // Convierte el texto en un objeto Time.
	if err != nil {                                                       // Si el formato es erróneo...
		log.Fatal("❌ Formato inválido. Use YYYY-MM-DD") // ...aborta la ejecución indicando el error.
	} // Cierra bloque if.

	// 5. Instanciación usando los principios de Programación Orientada a Objetos.
	canalAirbnb := models.NewAirbnbChannel()                                                    // Crea una instancia concreta del canal de venta (Cumple con la interfaz).
	nuevaReserva, err := models.NewReservation(huesped, fechaIngreso, fechaSalida, canalAirbnb) // Pasa los datos por las reglas de negocio en el modelo.
	if err != nil {                                                                             // Si el modelo rechaza la creación (ej. salida antes de entrada)...
		log.Fatalf("❌ Error de Negocio: %v\n", err) // ...detiene el programa imprimiendo la regla violada.
	} // Cierra bloque if.

	// 6. Persistencia de los datos validados hacia el servidor SQL.
	fmt.Println("\nGuardando en SQL Server...")        // Informa al usuario que inició el proceso de red.
	err = repositories.GuardarReservaSQL(nuevaReserva) // Envía el objeto validado a la capa de persistencia.
	if err != nil {                                    // Si hay falla de red o de SQL...
		log.Fatalf("❌ Fallo al guardar: %v\n", err) // ...muestra el mensaje de fallo y aborta.
	} // Cierra bloque if.

	// 7. Finalización exitosa.
	fmt.Printf("✅ ¡Éxito! Reserva de %s guardada correctamente.\n", nuevaReserva.GetGuestName()) // Extrae el nombre con el getter y confirma el fin del programa.
} // Cierra la función principal.
