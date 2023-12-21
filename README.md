# Yum Quick API

The Yum Quick API is a RESTful API that allows users to manage users, products, carts, orders, favorites, and cart items. It provides endpoints to perform CRUD operations for various resources.

## Getting Started

To get started with the Yum Quick API, follow these steps:

### Prerequisites

- Go programming language installed
- MySQL database installed and running

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/osmanbah441/yum-quick-api.git
   ```

2. Navigate to the project directory:
   ```bash
   cd yum-quick-api
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Set up the database:
   - Use the provided MySQL script to create the necessary database and tables. 
   `internal/models/yum_quick_db.sql`

### Usage

1. Run the application:
   ```bash
   go run ./cmd/api/main.go
   ```

2. Access the API endpoints using tools like `curl` or Postman:

   - **Users**
     - Retrieve all users: `GET /users`
     - Retrieve a user by ID: `GET /users/{id}`
     - Create a new user: `POST /users`
     - Update a user: `PUT /users/{id}`
     - Delete a user: `DELETE /users/{id}`

   - **Products**
     - Retrieve all products: `GET /products`
     - Retrieve a product by ID: `GET /products/{id}`
     - Create a new product: `POST /products`
     - Update a product: `PUT /products/{id}`
     - Delete a product: `DELETE /products/{id}`

   - **Carts**
     - Retrieve all carts: `GET /carts`
     - Retrieve a cart by ID: `GET /carts/{id}`
     - Create a new cart: `POST /carts`
     - Update a cart: `PUT /carts/{id}`
     - Delete a cart: `DELETE /carts/{id}`

   - **Orders**
     - Retrieve all orders: `GET /orders`
     - Retrieve an order by ID: `GET /orders/{id}`
     - Create a new order: `POST /orders`
     - Update an order: `PUT /orders/{id}`
     - Delete an order: `DELETE /orders/{id}`

   - **Favorites**
     - Retrieve all favorites: `GET /favorites`
     - Retrieve a favorite by ID: `GET /favorites/{id}`
     - Create a new favorite: `POST /favorites`
     - Delete a favorite: `DELETE /favorites/{id}`

   - **Cart Items**
     - Retrieve all cart items: `GET /cartitems`
     - Retrieve a cart item by ID: `GET /cartitems/{id}`
     - Create a new cart item: `POST /cartitems`
     - Update a cart item: `PUT /cartitems/{id}`
     - Delete a cart item: `DELETE /cartitems/{id}`

3. Ensure proper authentication and authorization mechanisms are implemented as per your application requirements.

## Contributing

Contributions are welcome! To contribute to the Yum Quick API, please follow these guidelines:
- Fork the repository
- Create a new branch for your changes: `git checkout -b feature/your-feature`
- Make changes and commit them: `git commit -m 'Add your changes'`
- Push to the branch: `git push origin feature/your-feature`
- Submit a pull request

