package auth

import (
	"bytes"
	"context"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	c, v, l, h := setup(t)
	defer c.Close()

	tests := []struct {
		name     string
		postData map[string]string
		check func(*testing.T, test)
	}{
		{
			"no body",
			nil,
			equalCheck(
				http.StatusBadRequest,
				jsonErrMsg(
					http.StatusBadRequest,
					errMalformedRequest,
				),
				&logrus.Entry{
					Level:   logrus.ErrorLevel,
					Message: logParseError,
				},
			),
		},
		{
			"empty body",
			map[string]string{},
			equalCheck(
				http.StatusBadRequest, jsonErrMsg(
					http.StatusBadRequest,
					map[string]string{"email": errRequired, "password": errRequired},
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logMalformedCredentials,
				},
			),
		},
		{
			"email empty",
			map[string]string{"email": ""},
			equalCheck(
				http.StatusBadRequest, jsonErrMsg(
					http.StatusBadRequest,
					map[string]string{"email": errRequired, "password": errRequired},
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logMalformedCredentials,
				},
			),
		},
		{
			"malformed email",
			map[string]string{"email": "invalid_mail"},
			equalCheck(
				http.StatusBadRequest, jsonErrMsg(
					http.StatusBadRequest,
					map[string]string{"email": "This is not a valid email.", "password": errRequired},
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logMalformedCredentials,
				},
			),
		},
		{
			"password missing",
			map[string]string{"email": "info@mail.com"},
			equalCheck(
				http.StatusBadRequest, jsonErrMsg(
					http.StatusBadRequest,
					map[string]string{"password": errRequired},
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logMalformedCredentials,
				},
			),
		},
		{
			"user does not exist",
			map[string]string{"email": "info@mail.com", "password": "passw0rd"},
			equalCheck(
				http.StatusUnauthorized, jsonErrMsg(
					http.StatusUnauthorized,
					errBadCredentials,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logUserNotFound,
				},
			),
		},
		{
			"wrong password",
			map[string]string{"email": disabledMail, "password": "passw0rd"},
			equalCheck(
				http.StatusUnauthorized, jsonErrMsg(
					http.StatusUnauthorized,
					errBadCredentials,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logPasswordMismatch,
				},
			),
		},
		{
			"user blocked",
			map[string]string{"email": disabledMail, "password": pw},
			equalCheck(
				http.StatusForbidden, jsonErrMsg(
					http.StatusForbidden,
					errUserBlocked,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logUserBlocked,
				},
			),
		},
		{
			"login",
			map[string]string{"email": enabledMail, "password": pw},
			func(t *testing.T, tt test) {
				assert.Equal(t, http.StatusOK, tt.rec.Code)

				d := make(map[string]interface{})
				assert.NoError(t, json.Unmarshal(tt.rec.Body.Bytes(), &d))

				// Log entry exists.
				assert.Equal(t, logUserLoggedIn, h.LastEntry().Message)
				assert.Equal(t, logrus.InfoLevel, h.LastEntry().Level)

				// A token has to be returned. Exactly 64 chars long.
				assert.Contains(t, d, "token")
				assert.IsType(t, "", d["token"])
				assert.Len(t, d["token"], 64)

				// There must be exactly one session now.
				n, err := c.Session.Query().Count(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, 5, n)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d []byte
			var err error

			if tt.postData != nil {
				d, err = json.Marshal(tt.postData)
				assert.NoError(t, err)
			}

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(d))
			req.Header.Set("Content-Type", "application/json")

			LoginHandler(c, v, l)(rec, req)
			tt.check(t, test{c, v, l, h, rec, req})
		})
	}
}
