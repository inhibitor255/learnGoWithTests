package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	t.Run("For Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.00

		if got != want {
			t.Errorf("want %g but got %g", want, got)
		}
	})
	t.Run("For Circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("want %g but got %g", want, got)
		}
	})
}
