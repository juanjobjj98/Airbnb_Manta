package models // Define que este archivo pertenece al paquete lógico 'models'.

// SalesChannel es una interfaz que define el contrato para las plataformas de venta.
// Esto permite aplicar polimorfismo en la aplicación.
type SalesChannel interface { // Inicia la declaración de la interfaz.
	GetCommissionRate() float64 // Método obligatorio que debe retornar un decimal (float64).
	GetChannelName() string     // Método obligatorio que debe retornar texto (string).
} // Cierra la interfaz.

// AirbnbChannel es una estructura (clase) que implementará la interfaz SalesChannel.
type AirbnbChannel struct { // Inicia la declaración de la estructura.
	commissionRate float64 // Atributo privado que guarda el porcentaje de comisión.
} // Cierra la estructura.

// NewAirbnbChannel es el constructor que inicializa un nuevo canal de Airbnb.
func NewAirbnbChannel() *AirbnbChannel { // Retorna un puntero a AirbnbChannel.
	return &AirbnbChannel{ // Retorna la dirección de memoria de la nueva instancia.
		commissionRate: 0.15, // Asigna un 15% de comisión por defecto.
	} // Cierra la inicialización.
} // Cierra la función constructora.

// GetCommissionRate es la implementación del método de la interfaz para AirbnbChannel.
func (a *AirbnbChannel) GetCommissionRate() float64 { // Recibe el puntero del canal.
	return a.commissionRate // Retorna el valor encapsulado de la comisión.
} // Cierra el método.

// GetChannelName es la implementación del método de la interfaz para AirbnbChannel.
func (a *AirbnbChannel) GetChannelName() string { // Recibe el puntero del canal.
	return "Airbnb" // Retorna explícitamente el nombre de la plataforma.
} // Cierra el método.
