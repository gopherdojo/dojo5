package typinggame

import "math/rand"

func defaultSentencePickFnc(sentences []string) string {
	randomIdx := rand.Intn(len(sentences))
	return sentences[randomIdx]
}
