package fixtures

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/go-api-skeleton/ent/schema"
	"math/rand"
	"time"
)

const jobCount = 5

func (r refs) job() *ent.Job {
	m := r[jobKey].([]*ent.Job)
	return m[rand.Intn(len(m))]
}

func jobs(refs refs, c *ent.Client) error {
	states := []string{schema.JobStateOpen, schema.JobStateClosed, schema.JobStateBilled}

	b := make([]*ent.JobCreate, jobCount)
	for i := 0; i < jobCount; i++ {
		t, err := time.Parse(randomdata.DateOutputLayout, randomdata.FullDate())
		if err != nil {
			return err
		}

		b[i] = c.Job.Create().
			SetDate(t).
			SetTask(randomdata.Paragraph()).
			SetState(states[rand.Intn(len(states))]).
			SetReport("").
			SetRest("").
			SetNote("").
			SetCustomerName(randomdata.SillyName()).
			SetRiskAssessmentRequired(randomdata.Boolean()).
			SetMaintenanceRequired(randomdata.Boolean()).
			AddUsers(refs.user())
	}

	var err error
	refs[jobKey], err = c.Job.CreateBulk(b...).Save(context.Background())

	return err
}
