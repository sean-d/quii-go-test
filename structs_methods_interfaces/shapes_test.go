package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{height: 10.0, width: 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// Iteration 3: table driven tests

	areaTests := []struct {
		testName string
		shape    Shape
		hasArea  float64
	}{
		{"rectangle area", &Rectangle{12, 6}, 72},
		{"circle area", &Circle{10.0}, 314.1592653589793},
		{"triangle area", &Triangle{12, 6}, 36.0},
	}

	for _, tableTest := range areaTests {
		got := tableTest.shape.Area()
		if got != tableTest.hasArea {
			t.Errorf("%#v: got %g want %g", tableTest.shape, got, tableTest.hasArea)
		}
	}

	// Iteration 2: helper function approach with interface use
	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}
	//
	//t.Run("rectangle area", func(t *testing.T) {
	//	rect := Rectangle{12.0, 6.0}
	//	checkArea(t, &rect, 72.0)
	//})
	//
	//t.Run("circle area", func(t *testing.T) {
	//	circle := Circle{10.0}
	//	checkArea(t, &circle, 314.1592653589793)
	//})

	// Iteration 1: traditional testing
	//t.Run("rectangle area", func(t *testing.T) {
	//	rect := Rectangle{height: 12.0, width: 6.0}
	//	got := rect.Area()
	//	want := 72.0
	//
	//	if got != want {
	//		t.Errorf("got %.2f want %.2f", got, want)
	//	}
	//})
	//
	//t.Run("circle area", func(t *testing.T) {
	//	circle := Circle{radius: 10}
	//	got := circle.Area()
	//	want := 314.1592653589793
	//
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//})

}
