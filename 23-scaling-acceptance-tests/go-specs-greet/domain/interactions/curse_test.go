package interactions_test

import (
	"testing"

	"github.com/tiagocorrea/go-specs-greet/domain/interactions"
	"github.com/tiagocorrea/go-specs-greet/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.MeanGreetAdapter(interactions.Curse),
	)
}
