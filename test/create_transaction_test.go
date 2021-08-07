package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

const ctURL = "http://localhost:8080/transaction"

func TestCreateTransaction(t *testing.T) {

	tt := []struct {
		description    string
		body           map[string]interface{}
		expectedOutput int
		expectedErr    bool
	}{
		{
			"nominal case",
			map[string]interface{}{
				"AccountID": "testaid1",
				"Amount":    100.0,
			},
			http.StatusCreated,
			true,
		},
		{
			"invalid account ID",
			map[string]interface{}{
				"AccountID": "testaid1_invalid",
				"Amount":    100.0,
			},
			http.StatusInternalServerError,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			body, err := json.Marshal(tc.body)
			if err != nil {
				log.Err(err)
			}
			response, err := http.Post(ctURL, "application/json", bytes.NewBuffer(body))
			assert.NotNil(t, response)
			assert.EqualValues(t, tc.expectedOutput, response.StatusCode)
		})
	}
}
