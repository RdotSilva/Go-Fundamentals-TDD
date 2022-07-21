package structs

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Calculate the perimeter of a square/rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Calculate the area of a square/rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}


