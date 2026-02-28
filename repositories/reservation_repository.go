package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JJ/Airbnb_Manta/models"
	_ "github.com/microsoft/go-mssqldb"
)

type ReservaWeb struct {
	ID           string
	Huesped      string
	Email        string
	Personas     int
	Mascotas     bool
	FechaIngreso time.Time
	FechaSalida  time.Time
	PrecioTotal  float64
	GananciaNeta float64
	CanalVenta   string
}

func GuardarReservaSQL(reserva *models.Reservation) error {
	conn := "server=localhost;port=1433;database=AirbnbManta;trusted_connection=yes;"
	db, _ := sql.Open("sqlserver", conn)
	defer db.Close()

	query := `INSERT INTO Reservas (ID, Huesped, Email, Personas, Mascotas, FechaIngreso, FechaSalida, PrecioTotal, GananciaNeta, CanalVenta) 
			  VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10)`

	_, err := db.Exec(query,
		reserva.GetID(), reserva.GetGuestName(), reserva.GetEmail(),
		reserva.GetGuests(), reserva.GetPets(), reserva.GetStartDate(),
		reserva.GetEndDate(), reserva.GetTotalPrice(), reserva.GetNetProfit(),
		reserva.GetChannel().GetChannelName())
	return err
}

func ObtenerTodasLasReservas() ([]ReservaWeb, error) {
	conn := "server=localhost;port=1433;database=AirbnbManta;trusted_connection=yes;"
	db, _ := sql.Open("sqlserver", conn)
	defer db.Close()

	query := "SELECT ID, Huesped, Email, Personas, Mascotas, FechaIngreso, FechaSalida, PrecioTotal, GananciaNeta, CanalVenta FROM Reservas ORDER BY FechaIngreso DESC"
	filas, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer filas.Close()

	var lista []ReservaWeb
	for filas.Next() {
		var r ReservaWeb
		filas.Scan(&r.ID, &r.Huesped, &r.Email, &r.Personas, &r.Mascotas, &r.FechaIngreso, &r.FechaSalida, &r.PrecioTotal, &r.GananciaNeta, &r.CanalVenta)
		lista = append(lista, r)
	}
	return lista, nil
}

// EliminarReservaSQL borra un registro por su ID
func EliminarReservaSQL(id string) error {
	conn := "server=localhost;port=1433;database=AirbnbManta;trusted_connection=yes;"
	db, err := sql.Open("sqlserver", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Reservas WHERE ID = @p1", id)
	return err
}

// VerificarDisponibilidad comprueba si ya existe una reserva en ese rango de fechas
func VerificarDisponibilidad(ingreso, salida time.Time) error {
	conn := "server=localhost;port=1433;database=AirbnbManta;trusted_connection=yes;"
	db, err := sql.Open("sqlserver", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Lógica SQL: Hay choque si una reserva existente entra ANTES de que la nueva salga,
	// Y sale DESPUÉS de que la nueva entre.
	query := `SELECT COUNT(*) FROM Reservas WHERE FechaIngreso < @p1 AND FechaSalida > @p2`

	var cantidad int
	// @p1 = salida nueva, @p2 = ingreso nuevo
	err = db.QueryRow(query, salida, ingreso).Scan(&cantidad)
	if err != nil {
		return err
	}

	if cantidad > 0 {
		return fmt.Errorf("las fechas seleccionadas ya están ocupadas por otra reserva")
	}

	return nil
}
