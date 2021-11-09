package web

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetFixedValue(t *testing.T) {
	value, err := mockbyhand.GetFixedValue("http://localhost:8080")
	if err != nil {
		t.Error(err)
	}

	if value != "actualvalue" {
		t.Errorf("expected 'actualvalue', got %s", value)
	}

}