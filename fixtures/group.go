package fixtures

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/masseelch/go-api-skeleton/ent"
	"math/rand"
)

const groupCount = 2

func (r refs) group() *ent.Group {
	m := r[groupKey].([]*ent.Group)
	return m[rand.Intn(len(m))]
}

func groups(refs refs, c *ent.Client) error {
	b := make([]*ent.GroupCreate, groupCount)

	for i := 0; i < groupCount; i++ {
		b[i] = c.Group.Create().SetTitle(randomdata.SillyName())
	}

	var err error
	refs[groupKey], err = c.Group.CreateBulk(b...).Save(context.Background())

	return err
}
