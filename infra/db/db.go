package db

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func generateULID() string {
	t := time.Now()
	e := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), e).String()
}
