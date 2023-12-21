package models

import "database/sql"

type Models struct {
	UserModel      userModel
	ProductModel   productModel
	OrderModel     orderModel
	OrderItemModel orderItemModel
	FavoriteModel  favoriteModel
	CartItemModel  cartItemModel
	CartModel      cartModel
}

func New(db *sql.DB) *Models {
	return &Models{
		UserModel:      userModel{db: db},
		ProductModel:   productModel{db: db},
		OrderModel:     orderModel{db: db},
		OrderItemModel: orderItemModel{db: db},
		FavoriteModel:  favoriteModel{db: db},
		CartItemModel:  cartItemModel{db: db},
		CartModel:      cartModel{db: db},
	}
}
