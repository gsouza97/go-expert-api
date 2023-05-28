package entity

import (
	"errors"
	"time"

	"github.com/gsouza97/go-expert-api/pkg/entity"
)

var (
	ErrRequiredID    = errors.New("id is required")
	ErrInvalidID     = errors.New("id is invalid")
	ErrRequiredName  = errors.New("name is required")
	ErrRequiredPrice = errors.New("price is required")
	ErrInvalidPrice  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrRequiredID
	}

	if _, err := entity.ParseId(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrRequiredName
	}

	if p.Price == 0 {
		return ErrRequiredPrice
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
