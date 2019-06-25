package typinggame

import (
	"math/rand"
	"time"
)

func defaultSentencePickFnc(sentences []string) string {
	if len(sentences) == 0 {
		return ""
	}

	rand.Seed(time.Now().Unix())
	randomIdx := rand.Intn(len(sentences))

	return sentences[randomIdx]
}
