package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gustvision/backend-interview/pkg/account"
	"github.com/gustvision/backend-interview/pkg/user/dto"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func (h *handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("method", "create_transaction").Logger()

	var req dto.CreateTransactionReq

	// we assume the amount cannot be equal to 0
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.AccountID == "" || req.Amount == 0 {
		logger.Error().Err(err).Msg("invalid payload")
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// verify the presence of the user
	a, err := h.account.Fetch(ctx, account.Filter{ID: req.AccountID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch account")
		http.Error(w, "failed to fetch account", http.StatusInternalServerError)
		return
	}

	// insert transition
	t := account.Transaction{
		ID:        uuid.NewString(), // generate a unique ID
		Amount:    req.Amount,
		AccountID: req.AccountID,
		CreatedAt: int64(time.Now().Second()), // generated time of the transaction to now
	}

	// insert the transaction in db
	err = h.account.InsertTransaction(ctx, t)
	if err != nil {
		logger.Error().Err(err).Msg(fmt.Sprintf("failed to insert transaction %#v", req))
		http.Error(w, fmt.Sprintf("failed to insert transaction %#v", req), http.StatusInternalServerError)
		return
	}

	// update the total amount of the given account by removing the amount from the total
	err = h.account.UpdateTotal(ctx, account.Filter{ID: a.ID, UserID: a.UserID, Total: a.Total - req.Amount})
	if err != nil {
		logger.Error().Err(err).Msg(fmt.Sprintf("failed to update total account %#v", req))
		http.Error(w, fmt.Sprintf("failed to total account %#v", req), http.StatusInternalServerError)
		return
	}

	// write response, successful insertion of data
	w.WriteHeader(http.StatusCreated)
}
