package cli

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func Test_formatResponse_For_Error(t *testing.T) {
	response := formatResponse(nil, errors.New("an error"))
	assert.Equal(t, "error: an error", response, "error was not formatted as expected")
}

func Test_formatResponse_For_Success(t *testing.T) {
	response := formatResponse([]byte(`"x"`), nil)
	assert.Equal(t, `"x"`, response, "string was not formatted as expected")
}

func Test_formatResponse_For_Success_Produces_Valid_Json(t *testing.T) {
	bytes, _ := json.Marshal(&map[string]string{"name": "abhijat"})
	response := formatResponse(bytes, nil)
	assert.NotContains(t, response, "error:")

	n := map[string]string{}
	err := json.Unmarshal([]byte(response), &n)
	
	assert.Nil(t, err)
	assert.Equal(t, n["name"], "abhijat")
}

func Test_formatResponseForSuccessFailsOnInvalidJson(t *testing.T) {
	response := formatResponse([]byte("x"), nil)
	assert.Contains(t, response, "error: invalid")
}