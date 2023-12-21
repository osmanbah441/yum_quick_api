package models

import (
	"database/sql"
	"time"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Ongoing   OrderStatus = "ongoing"
	Completed OrderStatus = "completed"
	Cancelled OrderStatus = "cancelled"
)

type Order struct {
	ID            int         `json:"id"`
	UserID        string      `json:"userId"`
	DeliveryDate  time.Time   `json:"deliveryDate"`
	TotalPrice    float64     `json:"totalPrice"`
	Quantity      int         `json:"quantity"`
	Status        OrderStatus `json:"status"`
	FormattedDate string      `json:"formattedDate"`
}

type orderModel struct {
	db *sql.DB
}

func (m *orderModel) GetOrderByID(id int) (*Order, error) {
	row := m.db.QueryRow("SELECT * FROM orders WHERE id = ?", id)
	order := &Order{}
	err := row.Scan(&order.ID, &order.UserID, &order.DeliveryDate, &order.TotalPrice, &order.Quantity, &order.Status, &order.FormattedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No order found
		}
		return nil, err // Handle other database errors
	}
	return order, nil
}

func (m *orderModel) GetAllOrders() ([]Order, error) {
	rows, err := m.db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.DeliveryDate, &order.TotalPrice, &order.Quantity, &order.Status, &order.FormattedDate)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (m *orderModel) CreateOrder(order *Order) error {
	stmt, err := m.db.Prepare("INSERT INTO orders (id, userId, deliveryDate, totalPrice, quantity, status, formattedDate) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID, order.UserID, order.DeliveryDate, order.TotalPrice, order.Quantity, order.Status, order.FormattedDate)
	return err
}

func (m *orderModel) UpdateOrder(order *Order) error {
	stmt, err := m.db.Prepare("UPDATE orders SET userId = ?, deliveryDate = ?, totalPrice = ?, quantity = ?, status = ?, formattedDate = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.UserID, order.DeliveryDate, order.TotalPrice, order.Quantity, order.Status, order.FormattedDate, order.ID)
	return err
}

func (m *orderModel) DeleteOrder(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM orders WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
