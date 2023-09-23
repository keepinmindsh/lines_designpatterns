package domain

type RsvnType string

const (
	Agoda   RsvnType = "Agoda"
	Booking RsvnType = "Booking"
	Expedia RsvnType = "Expedia"
)

type ReservationBuilder struct {
	RsvnType RsvnType
}

func NewReservationBuilder() *ReservationBuilder {
	return &ReservationBuilder{}
}

func (r *ReservationBuilder) Type(rsvnType RsvnType) {
	r.RsvnType = rsvnType
}

func (r *ReservationBuilder) make() *ReservationBuilder {
	return r
}
