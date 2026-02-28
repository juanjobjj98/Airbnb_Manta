package models

import (
	"testing"
	"time"
)

// PRUEBA UNITARIA 1: Validación de Reglas de Negocio
func TestNewReservationFechasInvalidas(t *testing.T) {
	canal := NewAirbnbChannel()
	ingreso := time.Now()
	salida := ingreso.Add(-24 * time.Hour) // Salida antes del ingreso (Error)

	_, err := NewReservation("Juan Reyes", "juan@test.com", 2, false, ingreso, salida, 100.0, canal)

	if err == nil {
		t.Errorf("Se esperaba un error por fechas inválidas, pero la reserva se creó.")
	}
}

// PRUEBA UNITARIA 2: Cálculo de Ganancia Neta (Inteligencia de Negocios)
func TestCalculoGananciaNetaAirbnb(t *testing.T) {
	canal := NewAirbnbChannel() // 15% de comisión
	ingreso := time.Now()
	salida := ingreso.Add(24 * time.Hour)
	precioBruto := 100.0

	reserva, err := NewReservation("Prueba", "p@test.com", 1, false, ingreso, salida, precioBruto, canal)

	if err != nil {
		t.Fatalf("No se esperaba error, pero se obtuvo: %v", err)
	}

	// 100 - (100 * 0.15) = 85.0
	gananciaEsperada := 85.0
	if reserva.GetNetProfit() != gananciaEsperada {
		t.Errorf("Error matemático. Esperado: %.2f, Obtenido: %.2f", gananciaEsperada, reserva.GetNetProfit())
	}
}
