package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/osmanbah441/yum_quick_api/internal/models"
)

// Handler function to get an order item by ID
func (app *application) getOrderItem(w http.ResponseWriter, r *http.Request) {

	orderItemID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	orderItem, err := app.models.OrderItemModel.GetOrderItemByID(orderItemID)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorResponse(w, r, http.StatusNotFound, "order item not found")
			return
		}
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"orderItem": orderItem}, http.StatusOK, nil)
}

// Handler function to create an order item
func (app *application) createOrderItem(w http.ResponseWriter, r *http.Request) {
	var orderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}

	err := app.models.OrderItemModel.CreateOrderItem(&orderItem)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"orderItem": orderItem}, http.StatusCreated, nil)
}

// Handler function to update an order item
func (app *application) updateOrderItem(w http.ResponseWriter, r *http.Request) {
	orderItemID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var orderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}
	orderItem.ID = orderItemID

	err = app.models.OrderItemModel.UpdateOrderItem(&orderItem)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"orderItem": orderItem}, http.StatusOK, nil)
}

// Handler function to delete an order item
func (app *application) deleteOrderItem(w http.ResponseWriter, r *http.Request) {
	orderItemID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.OrderItemModel.DeleteOrderItem(orderItemID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"message": "order item deleted successfully"}, http.StatusOK, nil)
}
