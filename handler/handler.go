package handler

import (
	"net/http"

	"bookstore/domain/book"
	"bookstore/domain/order"
	"bookstore/domain/user"
)

type Handler interface {
	// User
	CreateUser(w http.ResponseWriter, r *http.Request)
	// TODO: delete
	//Login(w http.ResponseWriter, r *http.Request)

	// Book
	ListBooks(w http.ResponseWriter, r *http.Request)

	// Order
	CreateOrder(w http.ResponseWriter, r *http.Request)
	GetOrderHistory(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	UserSvc  user.Service
	BookSvc  book.Service
	OrderSvc order.Service
}

func NewHandler(userSvc user.Service, bookSvc book.Service, orderSvc order.Service) Handler {
	return &handler{
		UserSvc:  userSvc,
		BookSvc:  bookSvc,
		OrderSvc: orderSvc,
	}
}
