package fixtures

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/masseelch/go-api-skeleton/ent"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

const userCount = 5

func (r refs) user() *ent.User {
	m := r[userKey].([]*ent.User)
	return m[rand.Intn(len(m))]
}

func users(refs refs, c *ent.Client) error {
	p, err := bcrypt.GenerateFromPassword([]byte("passw0rd!"), 0)
	if err != nil {
		return err
	}

	b := make([]*ent.UserCreate, userCount+1)

	b[0] = c.User.Create().
		SetEmail("user@api.com").
		SetPassword(string(p)).
		SetEnabled(true).
		SetGroup(refs.group())

	for i := 1; i <= userCount; i++ {
		b[i] = c.User.Create().
			SetEmail(randomdata.Email()).
			SetPassword(string(p)).
			SetEnabled(true).
			SetGroup(refs.group())
	}

	refs[userKey], err = c.User.CreateBulk(b...).Save(context.Background())

	return err
}
