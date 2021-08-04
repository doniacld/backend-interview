package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

const ctURL = "http://localhost:8080/transaction"

func TestCreateTransaction(t *testing.T) {

	tt := []struct {
		description    string
		body           map[string]interface{}
		expectedOutput int
	}{
		{
			"nominal case",
			map[string]interface{}{
				"AccountID": "testaid1",
				"Amount":    100,
			},
			http.StatusCreated,
		},
		{
			"invalid account ID",
			map[string]interface{}{
				"AccountID": "testaid1_invalid",
			},
			http.StatusBadRequest},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			body, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatal(err)
			}

			response, err := http.Post(ctURL, "application/json", bytes.NewBuffer(body))
			assert.Nil(t, err)
			assert.NotNil(t, response)

			bodyErr, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			assert.EqualValues(t, tc.expectedOutput, response.StatusCode)
			assert.EqualValues(t, "invalid payload\n", bodyErr)
		})
	}
}
