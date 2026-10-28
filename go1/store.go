package reservations

import (
	"errors"
	"fmt"
)

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrOutOfStock          = errors.New("out of stock")
	ErrReservationNotFound = errors.New("reservation not found")
)

// Store holds products and reservations in memory.
type Store struct {
	products     map[string]*Product
	reservations map[string]*Reservation
	seq          int
}

func NewStore() *Store {
	return &Store{
		products:     make(map[string]*Product),
		reservations: make(map[string]*Reservation),
	}
}

// AddProduct seeds a product. Given, don't touch it.
func (s *Store) AddProduct(p *Product) {
	s.products[p.ID] = p
}

// nextID returns a fresh reservation ID: r1, r2, r3...
func (s *Store) nextID() string {
	s.seq++
	return fmt.Sprintf("r%d", s.seq)
}

// Confirm looks up the reservation and delegates the transition to it.
// WORKED EXAMPLE: notice the shape: find, guard nil, delegate.
// Your Reserve and Cancel follow the same pattern.
func (s *Store) Confirm(reservationID string) error {
	r := s.reservations[reservationID]
	if r == nil {
		return ErrReservationNotFound
	}
	return r.confirm()
}

// Reserve checks stock, decrements it, and records a Held reservation.
// INVARIANT: never let stock go below zero.
// ErrProductNotFound if the product doesn't exist; ErrOutOfStock if not enough stock.
func (s *Store) Reserve(productID string, qty int) (*Reservation, error) {
	r := s.products[productID]
	if r == nil {
		return nil, ErrProductNotFound
	}
	if r.Stock >= qty {
		r.Stock -= qty
		res := &Reservation{
			ID:        s.nextID(),
			ProductID: productID,
			Quantity:  qty,
			Status:    Held,
		}
		s.reservations[res.ID] = res
		return res, nil
	}
	return nil, ErrOutOfStock
}

// Cancel releases a Held reservation and RETURNS its quantity to the product's stock.
// Only return stock if the cancel transition was legal.
// ErrReservationNotFound if the reservation doesn't exist.
func (s *Store) Cancel(reservationID string) error {
	// TODO: implement
	return errors.New("TODO: Cancel")
}

// Reservations returns every reservation of a product, sorted by ID.
// A map has no order: range over it, collect into a slice, then sort.
// Unknown product or no reservations: an empty slice (not nil), never an error.
func (s *Store) Reservations(productID string) []*Reservation {
	// TODO: implement
	return nil
}
