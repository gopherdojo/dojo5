package game

import (
	"github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/questions"
)

var ExportGetQuestion = (*Game).getQuestion

func (g *Game) ExportSetQs(qs *questions.Questions) {
	g.qs = qs
}
