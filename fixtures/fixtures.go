package fixtures

import (
	"github.com/masseelch/go-api-skeleton/ent"
	"math/rand"
	"time"
)

const (
	userKey = iota
	accountKey
)

type refs map[uint]interface{}

func Load(c *ent.Client) error {
	rand.Seed(time.Now().Unix())
	refs := make(refs)

	if err := users(refs, c); err != nil {
		return err
	}
	if err := sessions(refs, c); err != nil {
		return err
	}
	if err := accounts(refs, c); err != nil {
		return err
	}

	return nil
}
