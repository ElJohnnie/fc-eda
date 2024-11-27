package web

import (
	"encoding/json"
	"net/http"

	"fc-eda/internal/usecase/get_account"

	"github.com/go-chi/chi"
)

type WebAccountHandler struct {
	GetAccountUseCase get_account.GetAccountUseCase
}

func NewWebAccountHandler(getAccountUseCase get_account.GetAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		GetAccountUseCase: getAccountUseCase,
	}
}

func (h *WebAccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var dto get_account.GetAccountInputDTO
	dto.ID = id
	output, err := h.GetAccountUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
