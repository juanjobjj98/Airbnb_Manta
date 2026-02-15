package models // Define que este archivo pertenece al paquete lógico 'models'.

import ( // Inicia el bloque de importaciones.
	"errors" // Paquete estándar para crear y manejar mensajes de error.
	"fmt"    // Paquete estándar para formatear textos de salida.
	"time"   // Paquete estándar para manejar fechas y horas.

	"github.com/google/uuid" // Paquete externo para generar identificadores únicos.
) // Cierra el bloque de importaciones.

// Reservation es la estructura principal que encapsula los datos de una reserva.
type Reservation struct { // Inicia la declaración de la estructura.
	id        string       // Atributo privado: Identificador único de la reserva.
	guestName string       // Atributo privado: Nombre completo del huésped.
	startDate time.Time    // Atributo privado: Fecha exacta de llegada.
	endDate   time.Time    // Atributo privado: Fecha exacta de salida.
	channel   SalesChannel // Atributo privado: Interfaz de la plataforma de venta.
} // Cierra la estructura.

// NewReservation es el constructor seguro que valida las reglas de negocio antes de crear la reserva.
func NewReservation(guestName string, start, end time.Time, channel SalesChannel) (*Reservation, error) { // Recibe parámetros y retorna reserva o error.
	if start.After(end) || start.Equal(end) { // Verifica si la fecha de inicio es igual o posterior a la salida.
		return nil, errors.New("la fecha de inicio debe ser anterior a la fecha de fin") // Retorna nulo y un error descriptivo.
	} // Cierra la validación de fechas.
	if guestName == "" { // Verifica si el nombre del huésped está vacío.
		return nil, errors.New("el nombre del huésped es obligatorio") // Retorna nulo y un error descriptivo.
	} // Cierra la validación de nombre.
	if channel == nil { // Verifica si no se asignó un canal de venta.
		return nil, errors.New("se debe especificar un canal de venta") // Retorna nulo y un error descriptivo.
	} // Cierra la validación del canal.

	return &Reservation{ // Si todo es correcto, retorna la nueva instancia.
		id:        uuid.New().String(), // Genera un ID único y lo convierte a string.
		guestName: guestName,           // Asigna el nombre validado.
		startDate: start,               // Asigna la fecha de inicio validada.
		endDate:   end,                 // Asigna la fecha de fin validada.
		channel:   channel,             // Asigna el canal de venta.
	}, nil // Retorna nil en el espacio del error porque la creación fue exitosa.
} // Cierra el constructor.

// Overlaps es el motor lógico para prevenir overbooking comparando fechas.
func (r *Reservation) Overlaps(other *Reservation) bool { // Recibe otra reserva para comparar.
	return r.startDate.Before(other.endDate) && r.endDate.After(other.startDate) // Retorna true si las fechas chocan entre sí.
} // Cierra el método.

// Bloque de Getters: Permiten lectura segura de atributos privados (Encapsulamiento).
func (r *Reservation) GetID() string            { return r.id }        // Retorna el ID.
func (r *Reservation) GetGuestName() string     { return r.guestName } // Retorna el nombre del huésped.
func (r *Reservation) GetStartDate() time.Time  { return r.startDate } // Retorna la fecha de ingreso.
func (r *Reservation) GetEndDate() time.Time    { return r.endDate }   // Retorna la fecha de salida.
func (r *Reservation) GetChannel() SalesChannel { return r.channel }   // Retorna el objeto del canal de venta.

// String sobrescribe el método de impresión para mostrar la reserva en formato legible.
func (r *Reservation) String() string { // Retorna un string formateado.
	return fmt.Sprintf("Reserva: %s | Desde: %s Hasta: %s | Canal: %s", // Define la plantilla de texto.
		r.guestName, r.startDate.Format("2006-01-02"), r.endDate.Format("2006-01-02"), r.channel.GetChannelName()) // Inyecta las variables en la plantilla.
} // Cierra el método.
