package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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
				"Amount":    100,
			},
			http.StatusCreated,
			true,
		},
		{
			"invalid account ID",
			map[string]interface{}{
				"AccountID": "testaid1_invalid",
			},
			http.StatusBadRequest,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			body, err := json.Marshal(tc.body)
			// TODO assert expected error
			if err != nil {
				panic(err)
			}

			response, err := http.Post(ctURL, "application/json", bytes.NewBuffer(body))
			assert.Nil(t, err)
			assert.NotNil(t, response)

			assert.EqualValues(t, tc.expectedOutput, response.StatusCode)
		})
	}
}
