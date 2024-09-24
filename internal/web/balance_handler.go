package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com.br/cristian.scherer/eda-balance/internal/usecase/find_balance"
)

type WebBalanceHandler struct {
	usecase find_balance.FindBalanceUseCase
}

func NewWebAccountHandler(FindBalanceUseCase find_balance.FindBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		usecase: FindBalanceUseCase,
	}
}

func (h *WebBalanceHandler) FindBalance(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/balances/")

	dto := &find_balance.FindBalanceInputDTO{
		AccountID: id,
	}

	output, err := h.usecase.Executa(*dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
