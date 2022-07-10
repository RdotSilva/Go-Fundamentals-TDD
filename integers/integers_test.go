package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestSubtract(t *testing.T) {
	difference := Subtract(4, 2)
	expected := 2

	if difference != expected {
		t.Errorf("expected '%d' but got '%d'", expected, difference)
	}
}

func ExampleSubtract() {
	sum := Subtract(5, 1)
	fmt.Println(sum)
	// Output: 4
}