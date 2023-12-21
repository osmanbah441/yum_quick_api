package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/osmanbah441/yum_quick_api/internal/models"
)

// Handler function to get a product by ID
func (app *application) getProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	product, err := app.models.ProductModel.GetProductByID(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorResponse(w, r, http.StatusNotFound, "product not found")
			return
		}
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"product": product}, http.StatusOK, nil)
}

// Handler function to create a product
func (app *application) createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}

	err := app.models.ProductModel.CreateProduct(&product)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"product": product}, http.StatusCreated, nil)
}

// Handler function to update a product
func (app *application) updateProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}
	product.ID = productID

	err = app.models.ProductModel.UpdateProduct(&product)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"product": product}, http.StatusOK, nil)
}

// Handler function to delete a product
func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.ProductModel.DeleteProduct(productID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"message": "product deleted successfully"}, http.StatusOK, nil)
}

// Handler function to get all products
func (app *application) getAllProducts(w http.ResponseWriter, r *http.Request) {
	// todo use user id
	products, err := app.models.ProductModel.GetAllProducts(0)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"products": products}, http.StatusOK, nil)
}
