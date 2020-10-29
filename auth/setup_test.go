package auth

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/go-api-skeleton/ent/enttest"
	"github.com/masseelch/go-api-skeleton/util"
	go_token "github.com/masseelch/go-token"
	"github.com/sirupsen/logrus"
	logrusTest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	errRequired  = "This value is required."
	disabledMail = "disabled@example.com"
	enabledMail  = "enabled@example.com"
	pw           = "test_test_1234"
)

type test struct {
	c   *ent.Client
	v   *validator.Validate
	l   *logrus.Logger
	h   *logrusTest.Hook
	rec *httptest.ResponseRecorder
	req *http.Request
}

func setup(t *testing.T) (c *ent.Client, v *validator.Validate, l *logrus.Logger, h *logrusTest.Hook) {
	ctx := context.Background()
	c = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), 0)

	if _, err := c.User.CreateBulk(
		// Disabled user.
		c.User.Create().SetEmail(disabledMail).SetEnabled(false).SetPassword(string(hash)),
		// Enabled user.
		c.User.Create().SetEmail(enabledMail).SetEnabled(true).SetPassword(string(hash)),
	).Save(ctx); err != nil {
		panic(err)
	}

	if _, err := c.Session.CreateBulk(
		// Life has expired.
		c.Session.Create().SetID(go_token.Token(lifeExpiredToken)).SetIdleTimeExpiredAt(time.Now().Add(10)).SetLifeTimeExpiredAt(time.Now().Add(-10)),
		// Idle has expired.
		c.Session.Create().SetID(go_token.Token(idleExpiredToken)).SetIdleTimeExpiredAt(time.Now().Add(10)).SetLifeTimeExpiredAt(time.Now().Add(-10)),
		// Valid but blocked user.
		c.Session.Create().SetID(go_token.Token(validTokenBlockedUser)).SetIdleTimeExpiredAt(time.Now().Add(time.Hour)).SetLifeTimeExpiredAt(time.Now().Add(time.Hour)).SetUserID(1),
		// Valid.
		c.Session.Create().SetID(go_token.Token(validToken)).SetIdleTimeExpiredAt(time.Now().Add(time.Hour)).SetLifeTimeExpiredAt(time.Now().Add(time.Hour)).SetUserID(2),
	).Save(ctx); err != nil {
		panic(err)
	}

	v = util.Validator()

	l, h = logrusTest.NewNullLogger()

	return
}

func equalCheck(code int, body string, l *logrus.Entry) func(*testing.T, test) {
	return func(t *testing.T, tt test) {
		assert.Equal(t, code, tt.rec.Code)
		assert.JSONEq(t, body, tt.rec.Body.String())

		if l != nil {
			assert.Equal(t, l.Message, tt.h.LastEntry().Message)
			assert.Equal(t, l.Level, tt.h.LastEntry().Level)
		}
	}
}

func jsonErrMsg(code int, data interface{}) string {
	m, _ := json.Marshal(map[string]interface{}{
		"code":   code,
		"status": http.StatusText(code),
		"errors": data,
	})

	return string(m)
}
