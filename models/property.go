package models // Define que este archivo pertenece al paquete 'models'.

import ( // Inicia el bloque de importaciones.
	"errors" // Para el manejo de errores en el constructor.

	"github.com/google/uuid" // Para generar identificadores únicos de propiedad.
) // Cierra el bloque de importaciones.

// Property define una unidad de alojamiento físico (Ej. departamento en Manta).
type Property struct { // Inicia la estructura.
	id           string         // Atributo privado: Identificador único.
	name         string         // Atributo privado: Nombre de la propiedad.
	address      string         // Atributo privado: Ubicación física.
	capacity     int            // Atributo privado: Número máximo de huéspedes.
	reservations []*Reservation // Atributo privado: Lista de reservas asignadas.
} // Cierra la estructura.

// NewProperty es el constructor seguro para instanciar propiedades.
func NewProperty(name, address string, capacity int) (*Property, error) { // Recibe los datos y valida.
	if name == "" || address == "" { // Verifica que no existan campos de texto vacíos.
		return nil, errors.New("nombre y dirección obligatorios") // Retorna error si faltan datos.
	} // Cierra validación.
	return &Property{ // Retorna la instancia de la propiedad.
		id:           uuid.New().String(),     // Genera ID alfanumérico.
		name:         name,                    // Asigna el nombre.
		address:      address,                 // Asigna la dirección.
		capacity:     capacity,                // Asigna la capacidad.
		reservations: make([]*Reservation, 0), // Inicializa el arreglo dinámico (slice) vacío para evitar errores de memoria nula.
	}, nil // Indica que no hubo errores.
} // Cierra el constructor.

// GetName permite obtener el nombre de la propiedad encapsulada.
func (p *Property) GetName() string { return p.name } // Retorna el valor de name.
