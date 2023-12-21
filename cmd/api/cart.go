package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/osmanbah441/yum_quick_api/internal/models"
)

// Handler function to get a cart by ID
func (app *application) getCart(w http.ResponseWriter, r *http.Request) {
	cartID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	cart, err := app.models.CartModel.GetCartByID(cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorResponse(w, r, http.StatusNotFound, "cart not found")
			return
		}
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"cart": cart}, http.StatusOK, nil)
}

// Handler function to create a cart
func (app *application) createCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}

	err := app.models.CartModel.CreateCart(&cart)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"cart": cart}, http.StatusCreated, nil)
}

// Handler function to update a cart
func (app *application) updateCart(w http.ResponseWriter, r *http.Request) {
	cartID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}
	cart.ID = cartID

	err = app.models.CartModel.UpdateCart(&cart)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"cart": cart}, http.StatusOK, nil)
}

// Handler function to delete a cart
func (app *application) deleteCart(w http.ResponseWriter, r *http.Request) {
	cartID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.CartModel.DeleteCart(cartID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"message": "cart deleted successfully"}, http.StatusOK, nil)
}

// Handler function to get all carts for a user
func (app *application) getAllCarts(w http.ResponseWriter, r *http.Request) {
	userID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	carts, err := app.models.CartModel.GetAllCartsForUser(userID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"carts": carts}, http.StatusOK, nil)
}
