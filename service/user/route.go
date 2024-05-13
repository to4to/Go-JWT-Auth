package user
import (
	"net/http"
	"github.com/gorilla/mux"
)

type Handler struct {



	 
}


func NewHandler() *Handler {
	return &Handler{}
}


func (h *Handler) RegisterRoutesf(router *mux.Router) {
	router.HandleFunc("/register", h.Register).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")

}


func (h *Handler)handleLogin(w http.ResponseWriter, r *http.Request){

}