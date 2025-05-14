package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasPerimeter: 40},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6, Width: 6}, hasPerimeter: 24},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := tC.shape.Perimeter()

			if got != tC.hasPerimeter {
				t.Errorf("%#v got %g want %g", tC.shape, got, tC.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasArea: 100.00},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6, Width: 6}, hasArea: 36},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := tC.shape.Area()

			if got != tC.hasArea {
				t.Errorf("%#v got %g want %g", tC.shape, got, tC.hasArea)
			}
		})
	}
}

func ExampleTriangle_Area() {
	triangle := Triangle{Width: 6, Base: 4, Height: 4}
	fmt.Println(triangle.Area())

	// Output: 8
}

func ExampleRectangle_Perimeter() {
	triangle := Rectangle{Width: 6, Height: 4}
	fmt.Println(triangle.Perimeter())

	// Output: 20
}
