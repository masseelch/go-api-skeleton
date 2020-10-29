package auth

import (
	go_token "github.com/masseelch/go-token"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	invalidToken          = "ThisTokenDoesNotExist"
	lifeExpiredToken      = "lifeExpiredToken"
	idleExpiredToken      = "idleExpiredToken"
	validTokenBlockedUser = "validTokenBlockedUser"
	validToken            = "validToken"
)

func TestMiddleware(t *testing.T) {
	c, _, l, h := setup(t)
	defer c.Close()

	tests := []struct {
		name              string
		token             *string
		inHandlerCheck    func(*testing.T, test)
		afterHandlerCheck func(*testing.T, test)
	}{
		{
			"missing token",
			nil,
			assertHandlerNotCalled,
			equalCheck(
				http.StatusUnauthorized,
				jsonErrMsg(
					http.StatusUnauthorized,
					errMissingToken,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logTokenMissing,
				},
			),
		},
		{
			"token does not exist",
			&invalidToken,
			assertHandlerNotCalled,
			equalCheck(
				http.StatusUnauthorized,
				jsonErrMsg(
					http.StatusUnauthorized,
					errInvalidToken,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logTokenNotFound,
				},
			),
		},
		{
			"token lifetime expired",
			&lifeExpiredToken,
			assertHandlerNotCalled,
			equalCheck(
				http.StatusUnauthorized,
				jsonErrMsg(
					http.StatusUnauthorized,
					errExpiredToken,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logTokenExpired,
				},
			),
		},
		{
			"token idletime expired",
			&idleExpiredToken,
			assertHandlerNotCalled,
			equalCheck(
				http.StatusUnauthorized,
				jsonErrMsg(
					http.StatusUnauthorized,
					errExpiredToken,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logTokenExpired,
				},
			),
		},
		{
			"valid",
			&validToken,
			func(t *testing.T, tt test) {
				s := SessionFromContext(tt.req.Context())
				assert.NotNil(t, s)
				assert.Equal(t, go_token.Token(validToken), s.ID)

				assert.NotNil(t, s.Edges.User)
				assert.Equal(t, 2, s.Edges.User.ID)
			},
			func(t *testing.T, tt test) {
				assert.Equal(t, http.StatusOK, tt.rec.Code)

				assert.Equal(t, logUserAuthenticated, tt.h.LastEntry().Message)
				assert.Equal(t, logrus.InfoLevel, tt.h.LastEntry().Level)
			},
		},
		{
			"valid but blocked user",
			&validTokenBlockedUser,
			assertHandlerNotCalled,
			equalCheck(
				http.StatusUnauthorized,
				jsonErrMsg(
					http.StatusUnauthorized,
					errUserBlocked,
				),
				&logrus.Entry{
					Level:   logrus.InfoLevel,
					Message: logUserBlocked,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", nil)

			if tt.token != nil {
				req.Header.Set("Authorization", *tt.token)
			}

			Middleware(c, l)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				tt.inHandlerCheck(t, test{c, nil, l, h, rec, r})
			})).ServeHTTP(rec, req)

			tt.afterHandlerCheck(t, test{c, nil, l, h, rec, req})
		})
	}
}

func assertHandlerNotCalled(t *testing.T, tt test) {
	assert.Fail(t, "handler should not have been called")
}
