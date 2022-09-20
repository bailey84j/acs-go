package tools

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	client "github.com/bailey84j/acs-go/acs/client"
)

// ============= C ==============

// postResource - perform a POST request
func PostResource(client *client.Client, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {
	return resourceMutability(client, "POST", endpoint, args...)
}

// ============= R ==============

// getResource - perform a GET request
func GetResource(client *client.Client, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {
	//validate that param is url query format
	uriEndpoint := client.BaseURL
	uriEndpoint.Path = path.Join(uriEndpoint.Path, endpoint)

	//param = strings.ReplaceAll(param, " ", "%%20")

	//uriEndpoint.RawQuery = param

	//fmt.Printf("%s\n", uriEndpoint.String())

	//uriEndpoint := client.BaseURL.String() + endpoint + param
	logPrint(PrintLog{"Using this URL: " + uriEndpoint.String(), "INFO", client.LogLevel})

	req, err := http.NewRequest("GET", uriEndpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %s", err)
	}

	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	err = client.SetAuth(req)
	if err != nil {
		return nil, fmt.Errorf("error set auth: %s", err)
	}

	fmt.Printf("na: ")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed with error: %s", err)
	}

	fmt.Printf("\ncode: %d\n", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		return nil, ResponseError(
			StatusCodeFailed(res.Request.Method, uriEndpoint.String(), res.StatusCode, http.StatusOK),
			nil,
			res.Body,
		)
	}

	body, err := BodyToMap(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse body with error %s", err)
	}

	return body, nil
}

// ============= U ==============

// patchResource - perform a PATCH request
func PatchResource(client *client.Client, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {

	return resourceMutability(client, "PATCH", endpoint, args...)
}

// ============= U (alt) ==============

// putResource - perform a PUT request
func PutResource(client *client.Client, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {

	return resourceMutability(client, "PUT", endpoint, args...)
}

// ============= D ==============

// deleteResource - perform a DELETE request
func DeleteResource(client *client.Client, endpoint string, args ...map[string]interface{}) error {

	//validate that param is url query format
	uriEndpoint := client.BaseURL
	uriEndpoint.Path = path.Join(uriEndpoint.Path, endpoint)

	uriEndpoint.RawQuery = ""

	//fmt.Printf("%s\n", uriEndpoint.String())

	req, err := http.NewRequest("DELETE", uriEndpoint.String(), nil)

	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}

	client.SetAuth(req)

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error while making tower request: %s", err)
	}

	if !(res.StatusCode == http.StatusAccepted || res.StatusCode == http.StatusNoContent) {
		return ResponseError(
			StatusCodeFailed(res.Request.Method, uriEndpoint.String(), res.StatusCode, http.StatusAccepted, http.StatusNoContent),
			nil,
			res.Body,
		)
	}

	return nil
}

// ============= Support ==============

// resourceMutability handles the mutable interfaces such as POST, PUT, and PATCH
func resourceMutability(client *client.Client, verb, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {
	// Needs to be fixed
	//validate that param is url query format
	uriEndpoint := client.BaseURL
	uriEndpoint.Path = path.Join(uriEndpoint.Path, endpoint)

	uriEndpoint.RawQuery = ""

	//fmt.Printf("%s\n", uriEndpoint.String())

	//uriEndpoint := client.BaseURL.String() + endpoint
	logPrint(PrintLog{"Using this URL: " + uriEndpoint.String(), "DEBUG", client.LogLevel})

	prepRequestBody := make(map[string]interface{})

	if len(args) > 0 { // Handle splatted arguments
		for _, arg := range args {
			for key, value := range arg {
				switch v := value.(type) {
				case int:
					prepRequestBody[key] = v
				case float64:
					prepRequestBody[key] = v
				case bool:
					prepRequestBody[key] = v
				case string:
					if v == "" {
						continue
					}
					prepRequestBody[key] = v
				case []interface{}:
					prepRequestBody[key] = v
				case map[string]interface{}:
					prepRequestBody[key] = v
				default:

					continue
				}
			}
		}
	}

	requestBody, err := json.Marshal(prepRequestBody)

	fmt.Printf("\n Request Body: %v\n\n", string(requestBody))

	//fmt.Printf(string(requestBody))

	if err != nil {
		return nil, fmt.Errorf("unable to marshal request body with error: %s", err)
	}

	req, err := http.NewRequest(verb, uriEndpoint.String(), bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	client.SetAuth(req)

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error while making tower request: %s", err)
	}

	if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated) {
		return nil, ResponseError(
			StatusCodeFailed(res.Request.Method, uriEndpoint.String(), res.StatusCode, http.StatusOK, http.StatusCreated),
			requestBody,
			res.Body,
		)
	}

	// Decode the response body
	body, err := BodyToMap(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse body with error %s", err)
	}

	return body, nil
}
