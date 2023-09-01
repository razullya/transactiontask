package delivery

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type CreateAccountRequest struct {
		Username       string  `json:"username,omitempty"`
		CurrentBalance float64 `json:"current_balance,omitempty"`
	}

	var resp CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Account.CreateAccount(resp.Username, resp.CurrentBalance); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, statusOK)
}

func (h *Handler) getAllAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	accs, err := h.service.Account.GetAllAccounts()
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, accs)
}

func (h *Handler) getTransactionByAccountID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	id := r.URL.Query().Get("id")

	transactions, err := h.service.Account.GetTransactionByAccountID(id)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, transactions)
}

func (h *Handler) getAccountByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	id := r.URL.Query().Get("id")

	acc, err := h.service.Account.GetAccountByID(id)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, acc)
}
