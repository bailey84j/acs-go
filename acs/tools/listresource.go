package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	client "github.com/bailey84j/acs-go/acs/client"
	"github.com/jmespath/go-jmespath"
)

// ============= C ==============

// postResource - perform a POST request
func postResource(client *client.Client, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {
	return resourceMutability(client, "POST", endpoint, args...)
}

// ============= R ==============

// getResource - perform a GET request
func get(client *client.Client, endpoint string, param string) (map[string]interface{}, error) {
	//validate that param is url query format
	uriEndpoint := client.BaseURL
	uriEndpoint.Path = path.Join(uriEndpoint.Path, endpoint)

	param = strings.ReplaceAll(param, " ", "%%20")

	uriEndpoint.RawQuery = param

	//fmt.Printf("%s\n", uriEndpoint.String())

	//uriEndpoint := client.BaseURL.String() + endpoint + param
	logPrint(PrintLog{"Using this URL: " + uriEndpoint.String(), "INFO", client.LogLevel})

	req, err := http.NewRequest("GET", uriEndpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %s", err)
	}

	client.SetAuth(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed with error: %s", err)
	}

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
func patchResource(client *client.Client, sysID string, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {
	if sysID != "" {
		endpoint = endpoint + "/" + sysID
	}

	return resourceMutability(client, "PATCH", endpoint, args...)
}

// ============= U (alt) ==============

// putResource - perform a PUT request
func putResource(client *client.Client, sysID string, endpoint string, args ...map[string]interface{}) (map[string]interface{}, error) {
	if sysID != "" {
		endpoint = endpoint + "/" + sysID
	}

	return resourceMutability(client, "PUT", endpoint, args...)
}

// ============= D ==============

// deleteResource - perform a DELETE request
func deleteResource(client *client.Client, sysID string, endpoint string) error {
	if sysID != "" {
		endpoint = endpoint + "/" + sysID + "/"
	}

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

func msitosnowtable(msi map[string]interface{}, LogLevel string) (tablefields, error) {
	var snowfields tablefields

	msiByte, err := json.Marshal(msi)
	if err != nil {
		return nil, fmt.Errorf("msitosnowtable: First Marshal failure: %s", err)
	}
	var data interface{}

	err = json.Unmarshal(msiByte, &data)
	if err != nil {
		return nil, fmt.Errorf("msitosnowtable: failed to Convert Table Structure into JSON: %s", err)
	} else {
		logPrint(PrintLog{"msitosnowtable: Converted Table structure from sys_dictionary, for " + "xxNAme" + " into JSON", "INFO", LogLevel})
	}

	jQueryResult, err := jmespath.Search("result[].{name: element, mandatory: mandatory, internal_type: internal_type.value, reference_value: reference.value}", data)
	if err != nil {
		return nil, fmt.Errorf("msitosnowtable: failed to get Table Structure: %s", err)
	} else {
		logPrint(PrintLog{"msitosnowtable: Run JSON Query against sys_dictionary to return, name, mandatory, internal_type(value) and reference ", "DEBUG", LogLevel})
	}
	jQueryResultJSON, _ := json.Marshal(jQueryResult)
	if err != nil {
		return nil, fmt.Errorf("msitosnowtable: Marshal failure: %s", err)
	}

	err = json.Unmarshal(jQueryResultJSON, &snowfields)
	if err != nil {
		return nil, fmt.Errorf("msitosnowtable: Unmarshal failure: %s", err)
	}

	return snowfields, nil
}

func newfileUploadRequestclient(client *client.Client, params map[string]string, filepath string) (map[string]interface{}, error) {
	//(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	//validate that param is url query format

	uriEndpoint := client.BaseURL
	uriEndpoint.Path = path.Join(uriEndpoint.Path, "file")
	//uriEndpoint.Path = path.Join(uriEndpoint.Path, endpoint)

	//uriEndpoint := client.BaseURL.String() + endpoint
	logPrint(PrintLog{"Using this URL: " + uriEndpoint.String(), "DEBUG", client.LogLevel})

	// accept bytes in map string only
	// loop fallthrough

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("uploadFile", fi.Name())
	if err != nil {
		return nil, err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}

	querystring := ""
	for key, val := range params {
		querystring = querystring + key + "=" + val + "&"
	}

	querystring = querystring + "file_name=" + fi.Name()

	fmt.Printf("Query Parms test: %s", querystring)

	uriEndpoint.RawQuery = querystring
	/*
		for key, val := range params {
			_ = writer.WriteField(key, val)
		}*/
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	fmt.Printf("body: %+v", body)

	req, err := http.NewRequest("POST", uriEndpoint.String(), body)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	client.SetAuth(req)

	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// Set Content-Length
	req.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))

	writer.Close()

	fmt.Printf("req: %+v", req.Header)

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error while making tower request: %s", err)
	}

	if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated) {
		return nil, ResponseError(
			StatusCodeFailed(res.Request.Method, uriEndpoint.String(), res.StatusCode, http.StatusOK, http.StatusCreated),
			body.Bytes(),
			res.Body,
		)
	}

	// Decode the response body
	returnbody, err := BodyToMap(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse body with error %s", err)
	}

	return returnbody, nil
}
