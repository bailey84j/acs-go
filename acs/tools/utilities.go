package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

// BodyToMap converts an io.Closer to a map and returns the datastructure, if there is an error the body will be returned with the error
func BodyToMap(Body io.ReadCloser) (map[string]interface{}, error) {
	var body map[string]interface{}

	if Body == nil {
		err := errors.New("unable to parse nil io.ReadCloser")

		body = make(map[string]interface{})
		body["parse_error"] = fmt.Sprintf("%s", err)

		return body, err
	}
	err := json.NewDecoder(Body).Decode(&body)

	if err != nil {
		body = make(map[string]interface{})
		body["parse_error"] = fmt.Sprintf("%s", err)

		if strings.Contains(err.Error(), "EOF") {
			delete(body, "parse_error")
		}

		return nil, err
	}

	return body, nil
}

func jsontourlparm(jsonstring string) string {
	r := strings.NewReplacer(": ", "=",
		",", "&",
		"{", "",
		"}", "",
		" ", "%20",
		"\"", "",
	)
	result := r.Replace(jsonstring)
	return result
}

type PrintLog struct {
	Msg      string
	Level    string
	LogLevel string
}

func logPrint(data PrintLog) {
	// variables declaration

	switch data.LogLevel {
	case "info":
		if data.Level == "INFO" {
			log.Printf(data.Level + ":\tsnow-go\t" + data.Msg)
		}

	case "warning":
		if data.Level == "WARNING" {
			log.Printf(data.Level + ":\tsnow-go\t" + data.Msg)
		}
	case "debug":
		log.Printf(data.Level + ":\tsnow-go\t" + data.Msg)
	default:
		if data.Level == "ERROR" {
			log.Printf(data.Level + ":\tsnow-go\t" + data.Msg)
		}
	}
}

func Find(slice []string, val string) (int, bool) {
	if len(slice) < 1 {
		return -1, false
	}

	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
