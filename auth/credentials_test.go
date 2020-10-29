package auth

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCredentialsFromRequest(t *testing.T) {
	const jsonHeader = "application/json"
	const xmlHeader = "application/xml"
	const formHeader = "application/x-www-form-urlencoded"

	const email = "info@example.com"
	const pw = "test124"

	e := credentials{email, pw}

	// Create an empty request.
	r := reqWithValuesAndHeader(nil, "")
	c, err := credentialsFromRequest(r)
	assert.NoError(t, err)
	assert.Empty(t, c.Email)
	assert.Empty(t, c.Password)

	// Create an non-empty request without header set.
	r = reqWithValuesAndHeader(url.Values{"email": []string{email}, "password": []string{pw}}, "")
	c, err = credentialsFromRequest(r)
	assert.NoError(t, err)
	assert.Empty(t, c.Email)
	assert.Empty(t, c.Password)

	// Create form-url-encoded request with email set but empty password.
	r = reqWithValuesAndHeader(url.Values{"email": []string{email}}, formHeader)
	c, err = credentialsFromRequest(r)
	assert.NoError(t, err)
	assert.Equal(t, email, c.Email)
	assert.Empty(t, c.Password)

	// Create form-url-encoded request.
	r = reqWithValuesAndHeader(url.Values{"email": []string{email}, "password": []string{pw}}, formHeader)
	c, err = credentialsFromRequest(r)
	assert.NoError(t, err)
	assert.Equal(t, email, c.Email)
	assert.Equal(t, pw, c.Password)

	// Create json request.
	d, _ := json.Marshal(e)
	r = reqWithValuesAndHeader(d, jsonHeader)
	c, err = credentialsFromRequest(r)
	assert.NoError(t, err)
	assert.Equal(t, email, c.Email)
	assert.Equal(t, pw, c.Password)

	// Create xml request.
	d, _ = xml.Marshal(e)
	r = reqWithValuesAndHeader(d, xmlHeader)
	c, err = credentialsFromRequest(r)
	assert.NoError(t, err)
	assert.Equal(t, email, c.Email)
	assert.Equal(t, pw, c.Password)
}

func reqWithValuesAndHeader(data interface{}, contentType string) *http.Request {
	var body io.Reader

	switch d := data.(type) {
	case url.Values:
		body = strings.NewReader(d.Encode())
	case []byte:
		body = bytes.NewReader(d)
	}

	r := httptest.NewRequest(http.MethodPost, "/", body)
	if contentType != "" {
		r.Header.Set("Content-Type", contentType + "; charset=utf-8")
	}

	return r
}