package urand

import (
	"math/rand"
	"time"
)

func RandPercent() int32 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Int31n(99)
}
