package models

import "github.com/google/uuid"

// Property define una unidad de alojamiento en Manta.
// Usamos minúsculas para que el ID sea privado (Encapsulamiento).
type Property struct {
	id       string
	Name     string
	Address  string
	Capacity int
}

// NewProperty crea una nueva propiedad con un ID único de Google.
func NewProperty(name, address string, capacity int) *Property {
	return &Property{
		id:       uuid.New().String(),
		Name:     name,
		Address:  address,
		Capacity: capacity,
	}
}

// GetID permite obtener el ID de forma segura.
func (p *Property) GetID() string {
	return p.id
}
