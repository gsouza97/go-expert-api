package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gsouza97/go-expert-api/internal/dto"
	"github.com/gsouza97/go-expert-api/internal/entity"
	"github.com/gsouza97/go-expert-api/internal/infra/database"
	entityPkg "github.com/gsouza97/go-expert-api/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductDBInterface
}

func NewProductHandler(db database.ProductDBInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create product godoc
// @Summary      Create product
// @Description  Create product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request    body     dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products [post]
// @Security	 ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.CreateProduct(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get single product godoc
// @Summary      Get product
// @Description  Get product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 id    	path     string    true  "product id"  Format(uuid)
// @Success      200  {object}  entity.Product
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
// @Security	 ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// retorna o json do produto encontrado
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// List products godoc
// @Summary      List products
// @Description  List all product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 page    	query     string  	false  "page number"
// @Param		 limit   	query     string  	false  "page limit"
// @Success      200  {array}   entity.Product
// @Failure      500  {object}  Error
// @Router       /products [get]
// @Security	 ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := chi.URLParam(r, "sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// retorna o json dos produtos encontrados
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Update product godoc
// @Summary      Update product
// @Description  Update product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 id    		path     string    				 true  "product id"  Format(uuid)
// @Param        request    body     dto.CreateProductInput  true  "product request"
// @Success      200
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [put]
// @Security	 ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete product godoc
// @Summary      Delete product
// @Description  Delete product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 id    path    string    true  "product id"  Format(uuid)
// @Success      200
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [delete]
// @Security	 ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
