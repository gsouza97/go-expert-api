package database

import (
	"github.com/gsouza97/go-expert-api/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (db *ProductDB) CreateProduct(product entity.Product) error {
	return db.DB.Create(product).Error
}

func (db *ProductDB) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := db.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (db *ProductDB) FindAll(page int, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = db.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at" + sort).Find(&products).Error
	} else {
		err = db.DB.Order("created_at" + sort).Find(&products).Error
	}

	return products, err
}

func (db *ProductDB) Update(product *entity.Product) error {
	_, err := db.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return db.DB.Save(product).Error
}

func (db *ProductDB) Delete(id string) error {
	product, err := db.FindByID(id)
	if err != nil {
		return err
	}
	return db.DB.Delete(product).Error
}
