package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gsouza97/go-expert-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToTestDBAndMigrate() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.Product{})
	return db, nil
}

func TestCreateProduct(t *testing.T) {
	db, err := ConnectToTestDBAndMigrate()
	if err != nil {
		t.Error(err)
	}

	product, _ := entity.NewProduct("Product 1", 10.0)
	productDB := NewProductDB(db)

	err = productDB.CreateProduct(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := ConnectToTestDBAndMigrate()
	if err != nil {
		t.Error(err)
	}
	productDB := NewProductDB(db)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = db.Create(product).Error
		assert.NoError(t, err)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := ConnectToTestDBAndMigrate()
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	err = db.Create(product).Error
	assert.NoError(t, err)

	productDB := NewProductDB(db)
	p, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := ConnectToTestDBAndMigrate()
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	err = db.Create(product).Error
	assert.NoError(t, err)

	productDB := NewProductDB(db)
	product.Name = "Product 2"
	product.Price = 20.0
	err = productDB.Update(product)
	assert.NoError(t, err)

	p, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", p.Name)
	assert.Equal(t, 20.0, p.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := ConnectToTestDBAndMigrate()
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	err = db.Create(product).Error
	assert.NoError(t, err)

	productDB := NewProductDB(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
