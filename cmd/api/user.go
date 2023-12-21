package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/osmanbah441/yum_quick_api/internal/models"
)

// Handler function to get a user by ID
func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	user, err := app.models.UserModel.GetUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorResponse(w, r, http.StatusNotFound, "user not found")
			return
		}
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"user": user}, http.StatusOK, nil)
}

// Handler function to create a user
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}

	err := app.models.UserModel.CreateUser(&user)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"user": user}, http.StatusCreated, nil)
}

// Handler function to update a user
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}
	user.ID = userID

	err = app.models.UserModel.UpdateUser(&user)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"user": user}, http.StatusOK, nil)
}

// Handler function to delete a user
func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.UserModel.DeleteUser(userID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"message": "user deleted successfully"}, http.StatusOK, nil)
}

// Handler function to get all users
func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.models.UserModel.GetAllUsers()
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"users": users}, http.StatusOK, nil)
}
