package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/osmanbah441/yum_quick_api/internal/models"
)

// Handler function to get an order by ID
func (app *application) getOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	order, err := app.models.OrderModel.GetOrderByID(orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorResponse(w, r, http.StatusNotFound, "order not found")
			return
		}
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"order": order}, http.StatusOK, nil)
}

// Handler function to create an order
func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}

	err := app.models.OrderModel.CreateOrder(&order)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"order": order}, http.StatusCreated, nil)
}

// Handler function to update an order
func (app *application) updateOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}
	order.ID = orderID

	err = app.models.OrderModel.UpdateOrder(&order)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"order": order}, http.StatusOK, nil)
}

// Handler function to delete an order
func (app *application) deleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.OrderModel.DeleteOrder(orderID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"message": "order deleted successfully"}, http.StatusOK, nil)
}

// Handler function to get all orders
func (app *application) getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := app.models.OrderModel.GetAllOrders()
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"orders": orders}, http.StatusOK, nil)
}
