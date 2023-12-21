package models

import (
	"database/sql"
)

type OrderItem struct {
	ID         int     `json:"id"`
	OrderID    string  `json:"orderId"`
	ProductID  string  `json:"productId"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}

type orderItemModel struct {
	db *sql.DB
}

func (m *orderItemModel) GetOrderItemByID(id int) (*OrderItem, error) {
	row := m.db.QueryRow("SELECT * FROM order_items WHERE id = ?", id)
	orderItem := &OrderItem{}
	err := row.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.TotalPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No order item found
		}
		return nil, err // Handle other database errors
	}
	return orderItem, nil
}

func (m *orderItemModel) GetOrderItemsByOrderID(orderID int) ([]OrderItem, error) {
	rows, err := m.db.Query("SELECT * FROM order_items WHERE orderId = ?", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderItems := []OrderItem{}
	for rows.Next() {
		orderItem := OrderItem{}
		err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.TotalPrice)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (m *orderItemModel) CreateOrderItem(orderItem *OrderItem) error {
	stmt, err := m.db.Prepare("INSERT INTO order_items (id, orderId, productId, quantity, totalPrice) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderItem.ID, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.TotalPrice)
	return err
}

func (m *orderItemModel) UpdateOrderItem(orderItem *OrderItem) error {
	stmt, err := m.db.Prepare("UPDATE order_items SET orderId = ?, productId = ?, quantity = ?, totalPrice = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.TotalPrice, orderItem.ID)
	return err
}

func (m *orderItemModel) DeleteOrderItem(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM order_items WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
