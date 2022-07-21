package structs

// Calculate the perimeter of a square/rectangle
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Calculate the area of a square/rectangle
func Area(width float64, height float64) float64 {
	return width * height
}