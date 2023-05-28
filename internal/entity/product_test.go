package entity

import (
	"testing"

	"github.com/gsouza97/go-expert-api/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.NotEmpty(t, product.ID.String())
	assert.NotEmpty(t, product.CreatedAt.String())
}

func TestNewProduct_InvalidName(t *testing.T) {
	product, err := NewProduct("", 10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrRequiredName, err)
	assert.Nil(t, product)
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	product, err := NewProduct("Product 1", -10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
	assert.Nil(t, product)
}

func TestNewProduct_RequiredPrice(t *testing.T) {
	product, err := NewProduct("Product 1", 0.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrRequiredPrice, err)
	assert.Nil(t, product)
}

func TestProductValidate(t *testing.T) {
	product := &Product{
		ID:    entity.NewId(),
		Name:  "Product 1",
		Price: 10.0,
	}

	err := product.Validate()
	assert.Nil(t, err)
}
