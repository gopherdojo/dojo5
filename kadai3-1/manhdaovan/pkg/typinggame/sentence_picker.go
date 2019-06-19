package typinggame

import "math/rand"

func defaultSentencePickFnc(sentences []string) string {
	if len(sentences) == 0 {
		return ""
	}

	randomIdx := rand.Intn(len(sentences))
	return sentences[randomIdx]
}
