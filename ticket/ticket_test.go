package ticket

import "testing"

func TestTicketPrice(t *testing.T) {
	t.Run("should return 0 when age is 0", func(t *testing.T) {
		want := 0.0
		var age uint = 0

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})

	t.Run("Free ticket when age under 3", func(t *testing.T) {
		want := 0.0
		var age uint = 3

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})

	t.Run("Ticket $15 when age at 4 year old", func(t *testing.T) {
		want := 15.0
		var age uint = 4

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})

	t.Run("Ticket $15 when age is 15", func(t *testing.T) {
		want := 15.0
		var age uint = 15

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})

	t.Run("Ticket $30 when age over 15", func(t *testing.T) {
		want := 30.0
		var age uint = 16

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})

	t.Run("Ticket $30 when age is 50", func(t *testing.T) {
		want := 30.0
		var age uint = 50

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})

	t.Run("Ticket $5 when age over 50", func(t *testing.T) {
		want := 5.0
		var age uint = 51

		got := Price(age)

		if got != want {
			t.Errorf("Price(%d) = %f; want %f", age, got, want)
		}
	})
}
