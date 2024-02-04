package shelf

import "bookshelf/domain"

type Shelf struct {
	index  int
	shelfs []string
}

func (s *Shelf) HasNext() bool {
	if s.index < len(s.shelfs) {
		return true
	}
	return false
}

func (s *Shelf) Next() interface{} {
	if s.HasNext() {
		user := s.shelfs[s.index]
		s.index++
		return user
	}
	return nil
}

func NewBookShelf(shelfs []string) domain.Iterator {
	return &Shelf{
		shelfs: shelfs,
	}
}
