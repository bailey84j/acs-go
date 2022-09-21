package identity_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bailey84j/acs-go/acs"
)

type identityTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

