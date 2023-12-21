package models

import (
	"database/sql"
)

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	ImageURL      string  `json:"imageUrl"`
	Category      string  `json:"category"`
	AverageRating float64 `json:"averageRating"`
	Inventory     int     `json:"inventory"`
	IsFavorite    bool    `json:"isFavorite"` // Added for favorite information
}

type productModel struct {
	db *sql.DB
}

func (m *productModel) GetProductByID(id int) (*Product, error) {
	row := m.db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	product := &Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.ImageURL, &product.Category, &product.AverageRating, &product.Inventory)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No product found
		}
		return nil, err // Handle other database errors
	}
	return product, nil
}

func (m *productModel) GetAllProducts(userID int) ([]Product, error) {
	rows, err := m.db.Query(`
        SELECT p.*, IF(f.productId IS NOT NULL, TRUE, FALSE) AS isFavorite
        FROM products p
        LEFT JOIN favorites f ON p.id = f.productId AND f.userId = ?
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.ImageURL, &product.Category, &product.AverageRating, &product.Inventory, &product.IsFavorite)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (m *productModel) CreateProduct(product *Product) error {
	stmt, err := m.db.Prepare("INSERT INTO products (id, name, price, description, imageUrl, category, averageRating, inventory) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price, product.Description, product.ImageURL, product.Category, product.AverageRating, product.Inventory)
	return err
}

func (m *productModel) UpdateProduct(product *Product) error {
	stmt, err := m.db.Prepare("UPDATE products SET name = ?, price = ?, description = ?, imageUrl = ?, category = ?, averageRating = ?, inventory = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Description, product.ImageURL, product.Category, product.AverageRating, product.Inventory, product.ID)
	return err
}

func (m *productModel) DeleteProduct(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

// Additional methods as needed, for example:
// func (m *productModel)
