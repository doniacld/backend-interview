package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

const guURL = "http://localhost:8080/user"


func TestGetUser(t *testing.T) {

	tt := []struct {
		description    string
		body           map[string]interface{}
		expectedOutput int
	}{
		{
			"nominal case",
			map[string]interface{}{
				"ID": "testuid",
			},
			http.StatusOK,
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
