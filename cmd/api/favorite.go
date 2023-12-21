package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/osmanbah441/yum_quick_api/internal/models"
)

// Handler function to get a favorite by ID
func (app *application) getFavorite(w http.ResponseWriter, r *http.Request) {

	favoriteID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	favorite, err := app.models.FavoriteModel.GetFavoriteByID(favoriteID)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorResponse(w, r, http.StatusNotFound, "favorite not found")
			return
		}
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"favorite": favorite}, http.StatusOK, nil)
}

// Handler function to create a favorite
func (app *application) createFavorite(w http.ResponseWriter, r *http.Request) {
	var favorite models.Favorite
	if err := json.NewDecoder(r.Body).Decode(&favorite); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON")
		return
	}

	err := app.models.FavoriteModel.CreateFavorite(&favorite)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"favorite": favorite}, http.StatusCreated, nil)
}

// Handler function to delete a favorite
func (app *application) deleteFavorite(w http.ResponseWriter, r *http.Request) {
	favoriteID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.FavoriteModel.DeleteFavorite(favoriteID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"message": "favorite deleted successfully"}, http.StatusOK, nil)
}

// Handler function to get all favorites for a user
func (app *application) getAllFavorites(w http.ResponseWriter, r *http.Request) {
	userID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	favorites, err := app.models.FavoriteModel.GetAllFavoritesForUser(userID)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	app.writeJSON(w, envelop{"favorites": favorites}, http.StatusOK, nil)
}
