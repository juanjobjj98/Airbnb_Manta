package models

type SalesChannel interface {
	GetCommissionRate() float64
	GetChannelName() string
}

type AirbnbChannel struct{ commissionRate float64 }

func (a *AirbnbChannel) GetCommissionRate() float64 { return a.commissionRate }
func (a *AirbnbChannel) GetChannelName() string     { return "Airbnb" }

type BookingChannel struct{ commissionRate float64 }

func (b *BookingChannel) GetCommissionRate() float64 { return b.commissionRate }
func (b *BookingChannel) GetChannelName() string     { return "Booking" }

type DirectChannel struct{ commissionRate float64 }

func (d *DirectChannel) GetCommissionRate() float64 { return d.commissionRate }
func (d *DirectChannel) GetChannelName() string     { return "Directo" }

func CreateChannel(channelType string) SalesChannel {
	switch channelType {
	case "Booking":
		return &BookingChannel{0.15} // 15% Booking
	case "Directo":
		return &DirectChannel{0.00} // 0% Directo
	default:
		return &AirbnbChannel{0.15} // 15% Airbnb
	}
}
