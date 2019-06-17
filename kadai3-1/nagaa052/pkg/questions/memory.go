package questions

import "errors"

type inMemQ struct{}

var words = []string{
	"Germany",
	"Italy",
	"Russia",
	"Ukraine",
	"Ireland",
	"Iceland",
	"Albania",
	"Belgium",
	"Bulgaria",
	"Denmark",
	"Estonia",
	"Finland",
	"France",
	"Greece",
	"Greenland",
	"Hungary",
	"Cyprus",
	"Croatia",
	"Macedonia",
	"Malta",
	"Monaco",
	"Norway",
	"Austria",
	"Netherlands",
	"Poland",
	"Portugal",
	"Romania",
	"Switzerland",
	"Spain",
	"Sweden",
	"Serbia",
	"Montenegro",
}

func newInMemQ() (*inMemQ, error) {
	return &inMemQ{}, nil
}

func (iq *inMemQ) GetSize() int {
	return len(words)
}

func (iq *inMemQ) GetOne(index int) (*Question, error) {
	if len(words) <= index {
		return nil, errors.New("Question was not found")
	}

	return &Question{
		Word: words[index],
	}, nil
}
