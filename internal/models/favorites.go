package models

import (
	"database/sql"
)

type Favorite struct {
	ID        int    `json:"id"`
	UserID    string `json:"userId"`
	ProductID string `json:"productId"`
}

type favoriteModel struct {
	db *sql.DB
}

func (m *favoriteModel) GetFavoriteByID(id int) (*Favorite, error) {
	row := m.db.QueryRow("SELECT * FROM favorites WHERE id = ?", id)
	favorite := &Favorite{}
	err := row.Scan(&favorite.ID, &favorite.UserID, &favorite.ProductID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No favorite found
		}
		return nil, err // Handle other database errors
	}
	return favorite, nil
}

func (m *favoriteModel) GetAllFavoritesForUser(userID int) ([]Favorite, error) {
	rows, err := m.db.Query("SELECT * FROM favorites WHERE userId = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	favorites := []Favorite{}
	for rows.Next() {
		favorite := Favorite{}
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.ProductID)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return favorites, nil
}

func (m *favoriteModel) CreateFavorite(favorite *Favorite) error {
	stmt, err := m.db.Prepare("INSERT INTO favorites (id, userId, productId) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(favorite.ID, favorite.UserID, favorite.ProductID)
	return err
}

func (m *favoriteModel) DeleteFavorite(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM favorites WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
