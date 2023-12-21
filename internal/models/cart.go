package models

import (
	"database/sql"
)

type Cart struct {
	ID           int     `json:"id"`
	UserID       string  `json:"userId"`
	DeliveryCost float64 `json:"deliveryCost"`
	Quantity     int     `json:"quantity"`
	Subtotal     float64 `json:"subtotal"`
	Total        float64 `json:"total"`
}

type cartModel struct {
	db *sql.DB
}

func (m *cartModel) GetCartByID(id int) (*Cart, error) {
	row := m.db.QueryRow("SELECT * FROM carts WHERE id = ?", id)
	cart := &Cart{}
	err := row.Scan(&cart.ID, &cart.UserID, &cart.DeliveryCost, &cart.Quantity, &cart.Subtotal, &cart.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No cart found
		}
		return nil, err // Handle other database errors
	}
	return cart, nil
}

func (m *cartModel) GetAllCartsForUser(userID int) ([]Cart, error) {
	rows, err := m.db.Query("SELECT * FROM carts WHERE userId = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	carts := []Cart{}
	for rows.Next() {
		cart := Cart{}
		err := rows.Scan(&cart.ID, &cart.UserID, &cart.DeliveryCost, &cart.Quantity, &cart.Subtotal, &cart.Total)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return carts, nil
}

func (m *cartModel) CreateCart(cart *Cart) error {
	stmt, err := m.db.Prepare("INSERT INTO carts (id, userId, deliveryCost, quantity, subtotal, total) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cart.ID, cart.UserID, cart.DeliveryCost, cart.Quantity, cart.Subtotal, cart.Total)
	return err
}

func (m *cartModel) UpdateCart(cart *Cart) error {
	stmt, err := m.db.Prepare("UPDATE carts SET userId = ?, deliveryCost = ?, quantity = ?, subtotal = ?, total = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cart.UserID, cart.DeliveryCost, cart.Quantity, cart.Subtotal, cart.Total, cart.ID)
	return err
}

func (m *cartModel) DeleteCart(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM carts WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
