package structs

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}

// Calculate the perimeter of a square/rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Calculate the area of a square/rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}


