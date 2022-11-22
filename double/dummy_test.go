package double

import "testing"

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstname, lastname string) *Person {
	return &Person{}
}

func TestFindItShouldReturnsErrorWhenFirstnameOrLastnameEmpty(t *testing.T) {
	phonebook := &Phonebook{}
	want := ErrMissingArgs

	dd := DummySearcher{}
	_, got := phonebook.Find(dd, "", "")

	if got != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}
