package fixtures

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/masseelch/go-api-skeleton/ent"
	"math/rand"
)

const accountCount = 5

func (r refs) account() *ent.Account {
	m := r[accountKey].([]*ent.Account)
	return m[rand.Intn(len(m))]
}

func accounts(refs refs, c *ent.Client) error {
	b := make([]*ent.AccountCreate, accountCount)
	for i := 0; i < accountCount; i++ {
		// t, err := time.Parse(randomdata.DateOutputLayout, randomdata.FullDate())
		// if err != nil {
		// 	return err
		// }

		b[i] = c.Account.Create().
			SetTitle(randomdata.SillyName()).
			AddUsers(refs.user())
	}

	var err error
	refs[accountKey], err = c.Account.CreateBulk(b...).Save(context.Background())

	return err
}
