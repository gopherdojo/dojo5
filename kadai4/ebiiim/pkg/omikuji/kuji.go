package omikuji

import "math/rand"

type lottery interface {
	Do() string
}

type SimpleOmikuji struct {
	random   *rand.Rand
	kujiList []string
}

func NewSimpleOmikuji(random *rand.Rand) *SimpleOmikuji {
	k := &SimpleOmikuji{random: random}
	k.kujiList = []string{
		"大吉",
		"吉",
		"中吉",
		"小吉",
		"末吉",
		"凶",
	}
	return k
}

func (k *SimpleOmikuji) Do() string {
	return k.kujiList[rand.Intn(len(k.kujiList))]
}
