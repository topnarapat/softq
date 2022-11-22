package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCall   bool
	whatIsFirstName string
}

func (ss *SpySearcher) Search(people []*Person, firstname, lastname string) *Person {
	ss.searchWasCall = true
	ss.whatIsFirstName = "Jane"
	return &Person{
		Firstname: firstname,
		Lastname:  lastname,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWasCall {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't")
	}

	if spy.whatIsFirstName != "Jane" {
		t.Errorf("Expected to call 'Search' with 'Jane' as firstname, but it wasn't")
	}

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}
