package repositories // Define que este archivo pertenece al paquete 'repositories'.

import ( // Inicia el bloque de importaciones.
	"database/sql" // Paquete estándar para interactuar con bases de datos SQL.
	"fmt"          // Paquete para impresión en consola.

	_ "github.com/microsoft/go-mssqldb" // Importa el driver de SQL Server de forma anónima (necesario para init).
) // Cierra importaciones.

// InicializarBaseDeDatos verifica que la BD y la tabla existan; si no, las crea automáticamente.
func InicializarBaseDeDatos() error { // Retorna un error si algo falla en el proceso.
	fmt.Println("Verificando infraestructura de Base de Datos en SQL Server...") // Imprime mensaje de estado.

	// 1. Paso uno: Conectarse a la base de datos del sistema 'master' de SQL Server.
	connMaster := "server=localhost;port=1433;database=master;trusted_connection=yes;" // Cadena de conexión usando Autenticación de Windows.
	dbMaster, err := sql.Open("sqlserver", connMaster)                                 // Abre el pool de conexiones hacia 'master'.
	if err != nil {                                                                    // Si hay un error de sintaxis en la conexión...
		return fmt.Errorf("error conectando a master: %v", err) // ...retorna el error y aborta.
	} // Cierra bloque if.
	defer dbMaster.Close() // Asegura que la conexión a 'master' se cierre al terminar la función.

	// Consulta SQL para crear la base de datos 'AirbnbManta' solo si no existe previamente.
	queryCrearDB := `
		IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'AirbnbManta')
		BEGIN
			CREATE DATABASE AirbnbManta;
		END
	` // Cierra la cadena multilínea.
	_, err = dbMaster.Exec(queryCrearDB) // Ejecuta la consulta en el servidor SQL.
	if err != nil {                      // Si SQL Server rechaza la creación...
		return fmt.Errorf("error al crear base de datos: %v", err) // ...retorna el error de ejecución.
	} // Cierra bloque if.

	// 2. Paso dos: Conectarse a la base de datos específica 'AirbnbManta' que acabamos de asegurar que existe.
	connApp := "server=localhost;port=1433;database=AirbnbManta;trusted_connection=yes;" // Cadena de conexión apuntando a AirbnbManta.
	dbApp, err := sql.Open("sqlserver", connApp)                                         // Abre el nuevo pool de conexiones.
	if err != nil {                                                                      // Verifica errores de inicialización del driver.
		return fmt.Errorf("error conectando a AirbnbManta: %v", err) // Retorna el error descriptivo.
	} // Cierra bloque if.
	defer dbApp.Close() // Asegura el cierre de la conexión a la BD de la aplicación.

	// Consulta SQL para crear la tabla 'Reservas' con todas sus columnas si esta no existe.
	queryCrearTabla := `
		IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Reservas' AND xtype='U')
		BEGIN
			CREATE TABLE Reservas (
				ID NVARCHAR(50) PRIMARY KEY,
				Huesped NVARCHAR(100) NOT NULL,
				FechaIngreso DATE NOT NULL,
				FechaSalida DATE NOT NULL,
				CanalVenta NVARCHAR(50) NOT NULL
			);
		END
	` // Cierra la cadena multilínea.
	_, err = dbApp.Exec(queryCrearTabla) // Ejecuta el script de creación de tabla.
	if err != nil {                      // Si hay errores (ej. permisos insuficientes)...
		return fmt.Errorf("error al crear tabla Reservas: %v", err) // ...retorna el error.
	} // Cierra bloque if.

	fmt.Println("✓ Infraestructura lista (DB: AirbnbManta, Tabla: Reservas).") // Mensaje de éxito al usuario.
	return nil                                                                 // Retorna nil indicando que la inicialización fue completamente exitosa.
} // Cierra la función.
