package omikuji_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gopherdojo/dojo5/kadai4/ebiiim/pkg/omikuji"
	"github.com/seehuhn/mt19937"
)

func TestNewEmbeddedKuji(t *testing.T) {
	var (
		randGenTime = rand.New(rand.NewSource(time.Now().UnixNano()))
		randGenMT   = rand.New(mt19937.New())
	)
	cases := []struct {
		name   string
		random *rand.Rand
	}{
		{"normal", randGenTime},
		{"mersenne_twister", randGenMT},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			k := omikuji.NewSimpleOmikuji(c.random)
			k.Do()
		})
	}
}
