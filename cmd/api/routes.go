package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// User routes
	router.HandlerFunc(http.MethodGet, "/users/:id", app.getUser)
	router.HandlerFunc(http.MethodPost, "/users", app.createUser)
	router.HandlerFunc(http.MethodPut, "/users/:id", app.updateUser)
	router.HandlerFunc(http.MethodDelete, "/users/:id", app.deleteUser)
	router.HandlerFunc(http.MethodGet, "/users", app.getAllUsers)

	// Products routes
	router.HandlerFunc(http.MethodGet, "/products/:id", app.getProduct)
	router.HandlerFunc(http.MethodPost, "/products", app.createProduct)
	router.HandlerFunc(http.MethodPut, "/products/:id", app.updateProduct)
	router.HandlerFunc(http.MethodDelete, "/products/:id", app.deleteProduct)
	router.HandlerFunc(http.MethodGet, "/products", app.getAllProducts)

	// Order routes
	router.HandlerFunc(http.MethodGet, "/orders/:id", app.getOrder)
	router.HandlerFunc(http.MethodPost, "/orders", app.createOrder)
	router.HandlerFunc(http.MethodPut, "/orders/:id", app.updateOrder)
	router.HandlerFunc(http.MethodDelete, "/orders/:id", app.deleteOrder)
	router.HandlerFunc(http.MethodGet, "/orders", app.getAllOrders)

	// OrderItem routes
	router.HandlerFunc(http.MethodGet, "/orderitems/:id", app.getOrderItem)
	router.HandlerFunc(http.MethodPost, "/orderitems", app.createOrderItem)
	router.HandlerFunc(http.MethodPut, "/orderitems/:id", app.updateOrderItem)
	router.HandlerFunc(http.MethodDelete, "/orderitems/:id", app.deleteOrderItem)

	// Favorite routes
	router.HandlerFunc(http.MethodGet, "/favorites/:id", app.getFavorite)
	router.HandlerFunc(http.MethodPost, "/favorites", app.createFavorite)
	router.HandlerFunc(http.MethodDelete, "/favorites/:id", app.deleteFavorite)
	// router.HandlerFunc(http.MethodGet, "/users/:id/favorites", app.getAllFavorites)

	// Cart routes
	router.HandlerFunc(http.MethodGet, "/carts/:id", app.getCart)
	router.HandlerFunc(http.MethodPost, "/carts", app.createCart)
	router.HandlerFunc(http.MethodPut, "/carts/:id", app.updateCart)
	router.HandlerFunc(http.MethodDelete, "/carts/:id", app.deleteCart)
	// router.HandlerFunc(http.MethodGet, "/users/:userId/carts", app.getAllCarts)

	// CartItem routes
	router.HandlerFunc(http.MethodGet, "/cartitems/:id", app.getCartItem)
	router.HandlerFunc(http.MethodPost, "/cartitems", app.createCartItem)
	router.HandlerFunc(http.MethodPut, "/cartitems/:id", app.updateCartItem)
	router.HandlerFunc(http.MethodDelete, "/cartitems/:id", app.deleteCartItem)
	// router.HandlerFunc(http.MethodGet, "/carts/:cartId/cartitems", app.getAllCartItems)

	return router
}
