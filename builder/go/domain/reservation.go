package domain

type (
	Reservation interface {
		MakeReservation(reservation ReservationBuilder)
	}
)

type RsvnType string

const (
	Agoda   RsvnType = "Agoda"
	Booking RsvnType = "Booking"
	Expedia RsvnType = "Expedia"
)

// ReservationBuilder TODO Pointer를 사용하는 부분이 있어서 아래의 코드에서 설정해야할 부분을 검토 필요, 메모리 상의 효율을 검토 필요!
type ReservationBuilder struct {
	RsvnType        RsvnType
	ReservationName string
}

func NewReservationBuilder() ReservationBuilder {
	return ReservationBuilder{}
}

func (r *ReservationBuilder) Type(rsvnType RsvnType) *ReservationBuilder {
	r.RsvnType = rsvnType

	return r
}

func (r *ReservationBuilder) RsvnName(reservationName string) *ReservationBuilder {
	r.ReservationName = reservationName

	return r
}

func (r *ReservationBuilder) Make() *ReservationBuilder {
	return r
}
