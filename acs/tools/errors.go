package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// ResponseError formats an error return message with the provided info
func ResponseError(errorMessage string, RequestBody []byte, ResponseBody io.ReadCloser) error {
	var requestBody map[string]interface{}
	_ = json.Unmarshal(RequestBody, &requestBody)

	responseBody, _ := BodyToMap(ResponseBody)

	return fmt.Errorf("%s\n RequestBody: %v\n ResponseBody: %v", errorMessage, requestBody, responseBody)
}

// StatusCodeFailed returns a standardized failed status code error message string
func StatusCodeFailed(method, url string, gotStatus int, wantStatus ...int) string {
	if len(wantStatus) < 1 {
		panic("No wantStatus provided for StatusCodeFailed, this is a bug in the provider.")
	}

	var expectedStatus string
	for _, status := range wantStatus {
		expectedStatus = expectedStatus + fmt.Sprint(status) + "/"
	}
	expectedStatus = strings.TrimSuffix(expectedStatus, "/")

	return fmt.Sprintf("status code failed, got %d, expected %s\n Endpoint: %s %s", gotStatus, expectedStatus, method, url)
}
