package routes

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"linkconverter-api/helpers/test_helpers"
	"linkconverter-api/models/responses"
	"net/http"
	"testing"
)

func TestReturnStatusInfo(t *testing.T) {

	var recorder, context = test_helpers.SetupEcho("")

	NewStatusRouter().Status(context)

	result := responses.StatusResponseModel{}
	json.Unmarshal(recorder.Body.Bytes(), &result)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, true, result.Status)
	assert.Equal(t, "Always true", result.Description)
}
