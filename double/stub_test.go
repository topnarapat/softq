package double

import "testing"

type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstname, lastname string) *Person {
	return &Person{
		Firstname: firstname,
		Lastname:  lastname,
		Phone:     ss.phone,
	}
}

func TestFindReturnPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}

	ss := StubSearcher{
		phone: fakePhone,
	}
	phone, _ := phonebook.Find(ss, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}
