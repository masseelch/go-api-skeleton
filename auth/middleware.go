package auth

import (
	"context"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/go-api-skeleton/ent/session"
	"github.com/masseelch/go-token"
	"github.com/masseelch/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	sessionCtx = "go_api_skeleton_session_ctx"
)

var (
	SessionIdleTime = 15 * time.Minute
	SessionLifeTime = 24 * time.Hour

	errExpiredToken = "Token Expired"
	errInvalidToken = "Invalid Token"
	errMissingToken = "Missing Token"

	logTokenExpired      = "token expired"
	logTokenMissing      = "token missing"
	logTokenNotFound     = "token not found"
	logTokenUpdateError  = "error updating token"
	logUserAuthenticated = "user authenticated"
)

func Middleware(c *ent.Client, log *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := getTokenFromRequest(r)
			if t == "" {
				log.Info(logTokenMissing)
				render.Unauthorized(w, r, errMissingToken)
				return
			}

			// Load the session identified by the given token.
			s, err := c.Session.Query().WithUser().Where(session.IDEQ(t)).Only(r.Context())
			if err != nil {
				log.Info(logTokenNotFound)
				render.Unauthorized(w, r, errInvalidToken)
				return
			}

			// Check if session has not expired yet and has not been idle longer than allowed.
			n := time.Now()
			if s.LifeTimeExpiredAt.Before(n) || s.IdleTimeExpiredAt.Before(n) {
				log.Info(logTokenExpired)
				render.Unauthorized(w, r, errExpiredToken)
				return
			}

			// Check if the user is enabled. If so abort.
			if !s.Edges.User.Enabled {
				log.Info(logUserBlocked)
				render.Unauthorized(w, r, errUserBlocked)
				return
			}

			// Update idle time expiration.
			s.IdleTimeExpiredAt = n.Add(SessionIdleTime)
			if _, err = c.Session.UpdateOne(s).Save(r.Context()); err != nil {
				log.Info(logTokenUpdateError)
				render.Unauthorized(w, r, errMissingToken)
				return
			}

			// Save session on request context.
			log.WithField("email", s.Edges.User.Email).Info(logUserAuthenticated)
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), sessionCtx, s)))
		})
	}
}

func SessionFromContext(ctx context.Context) *ent.Session {
	return ctx.Value(sessionCtx).(*ent.Session)
}

func getTokenFromRequest(r *http.Request) token.Token {
	// Try to get token from header.
	t := r.Header.Get("Authorization")

	// If there was no token in the header look for the "auth" query parameter
	if t == "" {
		t = r.URL.Query().Get("auth")
	}

	return token.Token(t)
}
