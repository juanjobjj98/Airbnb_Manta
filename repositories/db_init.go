package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func InicializarBaseDeDatos() error {
	fmt.Println("Verificando infraestructura de BD...")

	connMaster := "server=localhost;port=1433;database=master;trusted_connection=yes;"
	dbMaster, err := sql.Open("sqlserver", connMaster)
	if err != nil {
		return err
	}
	defer dbMaster.Close()

	_, err = dbMaster.Exec("IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'AirbnbManta') CREATE DATABASE AirbnbManta")
	if err != nil {
		return err
	}

	connApp := "server=localhost;port=1433;database=AirbnbManta;trusted_connection=yes;"
	dbApp, err := sql.Open("sqlserver", connApp)
	if err != nil {
		return err
	}
	defer dbApp.Close()

	// Tabla expandida con los nuevos datos
	queryTabla := `
		IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Reservas' AND xtype='U')
		BEGIN
			CREATE TABLE Reservas (
				ID NVARCHAR(50) PRIMARY KEY,
				Huesped NVARCHAR(100) NOT NULL,
				Email NVARCHAR(100),
				Personas INT,
				Mascotas BIT,
				FechaIngreso DATE NOT NULL,
				FechaSalida DATE NOT NULL,
				PrecioTotal DECIMAL(10,2),
				GananciaNeta DECIMAL(10,2),
				CanalVenta NVARCHAR(50) NOT NULL
			);
		END
	`
	_, err = dbApp.Exec(queryTabla)
	return err
}
