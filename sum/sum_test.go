package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should return 3 when 1 and 2", func(t *testing.T) {
		// Arrange
		want := 3

		// Act
		got := sum(1, 2)

		// Assert
		if got != want {
			t.Errorf("sum(1, 2) = %d; want 3", got)
		}
	})

	t.Run("should return 1 when 1 and 0", func(t *testing.T) {
		want := 1

		got := sum(1, 0)

		if got != want {
			t.Errorf("sum(1, 0) = %d; want 1", got)
		}
	})

	t.Run("should return -2 when -1 and -1", func(t *testing.T) {
		want := -2

		got := sum(-1, -1)

		if got != want {
			t.Errorf("sum(-1, -1) = %d; want -2", got)
		}
	})

	t.Run("should multi parameters", func(t *testing.T) {
		want := 0

		got := sum([]int{}...)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("including sign integer", func(t *testing.T) {
		want := 7
		xs := []int{2, 3, 3, -1}

		got := sum(xs...) // got := sum([]int{2, 3, 3, -1}...) // ... = (2, 3, 3, -1)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
