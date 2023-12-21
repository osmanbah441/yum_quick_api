package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/osmanbah441/yum_quick_api/internal/models"
)

// Handler function to get a cart item by ID
func (app *application) getCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	cartItem, err := app.models.CartItemModel.GetCartItemByID(cartItemID)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorResponse(w, r, http.StatusNotFound, "cart item not found")
			return
		}
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"cartItem": cartItem}, http.StatusOK, nil)
}

// Handler function to create a cart item
func (app *application) createCartItem(w http.ResponseWriter, r *http.Request) {
	var cartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}

	err := app.models.CartItemModel.CreateCartItem(&cartItem)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"cartItem": cartItem}, http.StatusCreated, nil)
}

// Handler function to update a cart item
func (app *application) updateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var cartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}
	cartItem.ID = cartItemID

	err = app.models.CartItemModel.UpdateCartItem(&cartItem)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"cartItem": cartItem}, http.StatusOK, nil)
}

// Handler function to delete a cart item
func (app *application) deleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.CartItemModel.DeleteCartItem(cartItemID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"message": "cart item deleted successfully"}, http.StatusOK, nil)
}

// Handler function to get all cart items for a cart
func (app *application) getAllCartItems(w http.ResponseWriter, r *http.Request) {
	cartID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	cartItems, err := app.models.CartItemModel.GetAllCartItemsForCart(cartID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"cartItems": cartItems}, http.StatusOK, nil)
}
