package main

import (
	"encoding/json"
	"net/http"

	"github.com/gustvision/backend-interview/pkg/user"
	"github.com/gustvision/backend-interview/pkg/user/dto"
	"github.com/rs/zerolog/log"
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

	u, err := h.user.Fetch(ctx, user.Filter{ID: req.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch user")
		http.Error(w, "failed to fetch user", http.StatusInternalServerError)
		return
	}

	// # Compute user total
	var total float64
	/*

			totalAcc := func(a account.Account) error {
			total += a.Total
			// TODO not polite to not handle error here
			return nil
		}
			err, errc := h.account.FetchMany(ctx, account.Filter{UserID: req.ID}, totalAcc)
			if err != nil || errc != nil {
				logger.Error().Err(err).Msg(fmt.Sprintf("failed to fetch accounts for user %s", req.ID))
				http.Error(w, fmt.Sprintf("failed to fetch accounts for user %s", req.ID), http.StatusInternalServerError)
	*/

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
