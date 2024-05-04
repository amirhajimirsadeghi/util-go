package id

import (
	"math/rand"
	"sync"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

var mutex sync.Mutex

func New() string {
	mutex.Lock()
	defer mutex.Unlock()

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func NewWithTime(t time.Time) string {
	mutex.Lock()
	defer mutex.Unlock()

	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func ExtractTime(id string) (time.Time, error) {
	t, err := ulid.Parse(id)
	if err != nil {
		return time.Time{}, err
	}
	return ulid.Time(t.Time()), nil
}
