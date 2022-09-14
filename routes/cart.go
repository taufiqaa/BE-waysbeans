package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", middleware.Auth(h.FindCartId)).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", h.UpdateCartTransaction).Methods("PATCH")
	r.HandleFunc("/cart-qty/{id}", h.UpdateQuantity).Methods("PATCH")
	r.HandleFunc("/cart/{id}", h.DeleteCart).Methods("DELETE")
}
