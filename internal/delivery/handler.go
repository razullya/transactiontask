package delivery

import (
	"encoding/json"
	"net/http"
	"task/internal/service"
)

var statusOK = Status{
	Success: true,
}

type Status struct {
	Success bool `json:"success"`
}

type Handler struct {
	Mux     *http.ServeMux
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Mux:     http.NewServeMux(),
		service: services,
	}
}

func (h *Handler) InitRoutes() {

	h.Mux.HandleFunc("/account/create", h.createAccount) //

	h.Mux.HandleFunc("/account/all", h.getAllAccounts)
	h.Mux.HandleFunc("/account/transactions", h.getTransactionByAccountID)
	h.Mux.HandleFunc("/account", h.getAccountByID)

	h.Mux.HandleFunc("/transaction/create", h.createTransaction)
	h.Mux.HandleFunc("/transaction/all", h.getAllTransaction)
	h.Mux.HandleFunc("/transaction", h.getTransactionByID)
}

func (h *Handler) response(w http.ResponseWriter, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
