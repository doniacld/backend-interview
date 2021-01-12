package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gustvision/backend-interview/pkg/account"
	"github.com/gustvision/backend-interview/pkg/user"
	"github.com/gustvision/backend-interview/pkg/user/dto"
)

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("method", "get_user").Logger()

	var req dto.GetUserReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ID == "" {
		logger.Error().Err(err).Msg("invalid payload")
		http.Error(w, "invalid payload", http.StatusBadRequest)

		return
	}

	u, err := h.user.Fetch(ctx, user.Filter{req.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch user")
		http.Error(w, "failed to fetch user", http.StatusInternalServerError)

		return
	}

	var total float64

	if err := h.account.FetchMany(ctx, account.Filter{UserID: u.ID}, func(a account.Account) error {
		if err := h.account.FetchManyTransaction(ctx, account.FilterTransaction{
			AccountID: a.ID,
		}, func(t account.Transaction) error {
			total += t.Amount

			return nil
		}); err != nil {
			logger.Error().Err(err).Msg("failed to fetch transaction")
			http.Error(w, "failed to fetch transaction", http.StatusInternalServerError)

			return err
		}

		return nil
	}); err != nil {
		logger.Error().Err(err).Msg("failed to fetch account")
		http.Error(w, "failed to fetch account", http.StatusInternalServerError)

		return
	}

	// #Marshal results.
	raw, err := json.Marshal(dto.GetUserResp{
		User:  u,
		Total: total,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal PCs")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		return
	}

	logger.Info().Msg("success")
}
