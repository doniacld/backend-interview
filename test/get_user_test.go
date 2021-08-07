package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

const guURL = "http://localhost:8080/user"

func TestGetUser(t *testing.T) {

	tt := []struct {
		description    string
		body           map[string]interface{}
		expectedOutput int
		expectedErr    bool
	}{
		{
			"nominal case",
			map[string]interface{}{
				"ID": "testuid",
			},
			http.StatusOK,
			false,
		},
		{
			"unknown ID",
			map[string]interface{}{
				"ID": "testuid_invalid",
			},
			http.StatusInternalServerError, // for the moment, should be a 404
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			body, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatal(err)
			}

			request, err := http.NewRequest("GET", guURL, bytes.NewBuffer(body))
			client := &http.Client{}
			response, err := client.Do(request)

			assert.NotNil(t, response)

			_, err = ioutil.ReadAll(response.Body)
			if err != nil {
				log.Err(err)
			}

			assert.EqualValues(t, tc.expectedOutput, response.StatusCode)
		})
	}
}
