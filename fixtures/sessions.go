package fixtures

import (
	"context"
	"github.com/masseelch/go-api-skeleton/ent"
	"time"
)

func sessions(_ refs, c *ent.Client) error {
	_, err := c.Session.Create().
		SetID("valid_token").
		SetIdleTimeExpiredAt(time.Now().Add(24 * time.Hour)).
		SetLifeTimeExpiredAt(time.Now().Add(24 * time.Hour)).
		SetUserID(1).
		Save(context.Background())

	return err
}
