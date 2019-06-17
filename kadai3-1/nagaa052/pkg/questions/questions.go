package questions

// Questions is Get a questions, etc.
type Questions struct {
	Store
}

// Store is interface for acquiring questions data
type Store interface {
	GetSize() int
	GetOne(index int) (*Question, error)
}

var _ Store = &inMemQ{}

// New is Initialize Questions
func New() (*Questions, error) {
	imq, err := newInMemQ()
	if err != nil {
		return nil, err
	}
	return NewWithStore(imq)
}

// NewWithStore is Questions initialization by specifying Store
func NewWithStore(s Store) (*Questions, error) {
	return &Questions{
		Store: s,
	}, nil
}

// Question is have a Word
type Question struct {
	Word string
}

// IsCorrect is Determine if the questions and answers are correct
func (q *Question) IsCorrect(answer string) bool {
	return q.Word == answer
}
