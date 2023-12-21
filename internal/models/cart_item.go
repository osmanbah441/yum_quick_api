package models

import (
	"database/sql"
)

type CartItem struct {
	ID         int     `json:"id"`
	CartID     string  `json:"cartId"`
	ProductID  string  `json:"productId"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}

type cartItemModel struct {
	db *sql.DB
}

func (m *cartItemModel) GetCartItemByID(id int) (*CartItem, error) {
	row := m.db.QueryRow("SELECT * FROM cart_items WHERE id = ?", id)
	cartItem := &CartItem{}
	err := row.Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.TotalPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No cart item found
		}
		return nil, err // Handle other database errors
	}
	return cartItem, nil
}

func (m *cartItemModel) GetAllCartItemsForCart(cartID int) ([]CartItem, error) {
	rows, err := m.db.Query("SELECT * FROM cart_items WHERE cartId = ?", cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cartItems := []CartItem{}
	for rows.Next() {
		cartItem := CartItem{}
		err := rows.Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.TotalPrice)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (m *cartItemModel) CreateCartItem(cartItem *CartItem) error {
	stmt, err := m.db.Prepare("INSERT INTO cart_items (id, cartId, productId, quantity, totalPrice) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cartItem.ID, cartItem.CartID, cartItem.ProductID, cartItem.Quantity, cartItem.TotalPrice)
	return err
}

func (m *cartItemModel) UpdateCartItem(cartItem *CartItem) error {
	stmt, err := m.db.Prepare("UPDATE cart_items SET cartId = ?, productId = ?, quantity = ?, totalPrice = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cartItem.CartID, cartItem.ProductID, cartItem.Quantity, cartItem.TotalPrice, cartItem.ID)
	return err
}

func (m *cartItemModel) DeleteCartItem(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM cart_items WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
