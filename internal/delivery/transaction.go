package delivery

import (
	"encoding/json"
	"net/http"
	"task/models"
)

func (h *Handler) createTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	var resp models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Transaction.CreateTransaction(resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, statusOK)
}

func (h *Handler) getTransactionByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	id := r.URL.Query().Get("id")

	acc, err := h.service.Transaction.GetTransactionByID(id)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, acc)
}

func (h *Handler) getAllTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	accs, err := h.service.Transaction.GetAllTransactions()
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, accs)
}
