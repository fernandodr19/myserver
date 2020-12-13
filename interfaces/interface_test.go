package interfaces

import "testing"

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, form Form, expected float64) {
		result := form.Area()
		if result != expected {
			t.Errorf("Result %.2f; Expected %.2f", result, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		verifyArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		verifyArea(t, circle, 314.1592653589793)
	})
}
