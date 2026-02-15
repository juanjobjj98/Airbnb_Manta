package repositories // Define el paquete al que pertenece este archivo.

import ( // Inicia bloque de importaciones.
	"database/sql" // Paquete de base de datos estándar de Go.
	"fmt"          // Paquete para formatear strings de error.

	"github.com/JJ/Airbnb_Manta/models" // Importa tus modelos para extraer la información.
	_ "github.com/microsoft/go-mssqldb" // Importa el driver de SQL Server en segundo plano.
) // Cierra importaciones.

// GuardarReservaSQL ejecuta un comando INSERT para guardar una reserva en la base de datos.
func GuardarReservaSQL(reserva *models.Reservation) error { // Recibe el puntero de la reserva a guardar.
	connString := "server=localhost;port=1433;database=AirbnbManta;trusted_connection=yes;" // Define la cadena de conexión.
	db, err := sql.Open("sqlserver", connString)                                            // Inicializa la conexión.
	if err != nil {                                                                         // Si la conexión inicial falla...
		return fmt.Errorf("error abriendo conexión: %v", err) // ...devuelve el error.
	} // Cierra bloque if.
	defer db.Close() // Asegura la liberación de recursos de red y memoria.

	// Prepara la consulta parametrizada para evitar inyección SQL (@p1, @p2, etc.).
	query := `
		INSERT INTO Reservas (ID, Huesped, FechaIngreso, FechaSalida, CanalVenta) 
		VALUES (@p1, @p2, @p3, @p4, @p5)` // Cierra el texto de la consulta.

	// Ejecuta la consulta mapeando los getters de la reserva a los parámetros @p de SQL.
	_, err = db.Exec(query, // Pasa la consulta base.
		reserva.GetID(),                       // @p1: Extrae el ID privado.
		reserva.GetGuestName(),                // @p2: Extrae el nombre del huésped.
		reserva.GetStartDate(),                // @p3: Extrae la fecha de entrada.
		reserva.GetEndDate(),                  // @p4: Extrae la fecha de salida.
		reserva.GetChannel().GetChannelName(), // @p5: Extrae el nombre del canal desde la interfaz de ventas.
	) // Finaliza los parámetros de ejecución.

	if err != nil { // Si la inserción a nivel de base de datos falla...
		return fmt.Errorf("error ejecutando INSERT: %v", err) // ...retorna el error de base de datos.
	} // Cierra bloque if.

	return nil // Retorna nil, confirmando que la reserva está guardada en el disco.
} // Cierra la función.
