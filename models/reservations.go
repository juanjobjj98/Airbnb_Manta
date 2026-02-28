package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	id         string
	guestName  string
	email      string
	guests     int
	pets       bool
	startDate  time.Time
	endDate    time.Time
	totalPrice float64
	netProfit  float64
	channel    SalesChannel
}

// Constructor con Reglas de Negocio
func NewReservation(guestName, email string, guests int, pets bool, start, end time.Time, totalPrice float64, channel SalesChannel) (*Reservation, error) {
	if start.After(end) || start.Equal(end) {
		return nil, errors.New("la fecha de salida debe ser posterior a la de ingreso")
	}
	if guests <= 0 {
		return nil, errors.New("debe haber al menos 1 huésped")
	}
	if totalPrice < 0 {
		return nil, errors.New("el precio no puede ser negativo")
	}

	// CÁLCULO DE INTELIGENCIA DE NEGOCIOS: Ganancia Neta
	comision := totalPrice * channel.GetCommissionRate()
	gananciaNeta := totalPrice - comision

	return &Reservation{
		id:         uuid.New().String(),
		guestName:  guestName,
		email:      email,
		guests:     guests,
		pets:       pets,
		startDate:  start,
		endDate:    end,
		totalPrice: totalPrice,
		netProfit:  gananciaNeta,
		channel:    channel,
	}, nil
}

// Getters para persistencia
func (r *Reservation) GetID() string            { return r.id }
func (r *Reservation) GetGuestName() string     { return r.guestName }
func (r *Reservation) GetEmail() string         { return r.email }
func (r *Reservation) GetGuests() int           { return r.guests }
func (r *Reservation) GetPets() bool            { return r.pets }
func (r *Reservation) GetStartDate() time.Time  { return r.startDate }
func (r *Reservation) GetEndDate() time.Time    { return r.endDate }
func (r *Reservation) GetTotalPrice() float64   { return r.totalPrice }
func (r *Reservation) GetNetProfit() float64    { return r.netProfit }
func (r *Reservation) GetChannel() SalesChannel { return r.channel }
