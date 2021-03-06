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
	difference := Subtract(5, 1)
	fmt.Println(difference)
	// Output: 4
}

func TestMultiply(t *testing.T) {
	difference := Multiply(6, 3)
	expected := 18

	if difference != expected {
		t.Errorf("expected '%d' but got '%d'", expected, difference)
	}
}

func ExampleMultiply() {
	product := Multiply(2, 5)
	fmt.Println(product)
	// Output: 10
}

func TestDivide(t *testing.T) {
	quotient := Divide(18, 3)
	expected := 6

	if quotient != expected {
		t.Errorf("expected '%d' but got '%d'", expected, quotient)
	}
}

func ExampleDivide() {
	quotient := Divide(20, 5)
	fmt.Println(quotient)
	// Output: 4
}