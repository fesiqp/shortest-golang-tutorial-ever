// Go is a language built by Google

// Every Go file is a package
// Every Go file should declare which package it belongs
package main

// And this is how we import one of them (or many of them).
// If you are importing just one package, the parentheses are optional.
import (
	"fmt"
	"log"
)

// Declaring package level var
var y int

// Initializing package level var
var x = 10

// This is how we initialize constants in Go
const (
	myName = "John"
	myAge  = 10
)

// With structs, we are able to create our own types (similar to a class in OOP).
// Lower case firstName -> private to the package.
// First letter uppercase FirstName -> public outside the package.
type person struct {
	firstName string
	lastName  string
}

//  Functions have three sets of parentheses:
//  The first is receiver -> attach a type to the function (similar to a method in OOP)
//  The second are the parameters
// The third are the returns (Golang supports multiple returns)

func (p person) speak() (int, string) {
	fmt.Println("Speak, " + p.firstName)
	return 10, "nice"
}

// NewPerson is constructor, this pattern is used a lot to generate structs (similar to an object in OOP)
// nil is "equivalent" to the NULL of other languages, and it means "unitialized"
func NewPerson(firstName, lastName string) (*person, error) {
	if firstName == "" || lastName == "" {
		return nil, fmt.Errorf("Empty parameters")
	}

	return &person{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

// Inheritance in Golang is quite straightforward:
// You just need to add the person struct as an anonymous field
// and the student struct will automatically inherit every method declared in person struct
// (field is similar to attributes in OOP)
type student struct {
	person
	course string
}

// NewStudent is an example of constructor for promoted structs
func NewStudent(firstName, lastName, course string) (*student, error) {
	p, err := NewPerson(firstName, lastName)
	if err != nil {
		return nil, fmt.Errorf("Empty parameters")
	}

	if course == "" {
		return nil, fmt.Errorf("Empty course")
	}

	return &student{
		person: *p, // When dealing with anonymous fields, you should pass it as a pointer to a struct, in this case, person
		course: course,
	}, nil
}

// This is the main function of every package
func main() {
	// Go is a strongly typed language.
	// We can declare variables like we did with y above.
	// Or we can use a short declaration, which implies the variable type.

	x := "My String" // x is a string
	z := 25.6        // z is a float64

	// Printing the variables type:
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", z)

	// Printing the variables value:
	fmt.Println(x + "\n")
	fmt.Println(z)

	// As you may have noted,
	// ending semicolons are completely optional.

	// Composite literals
	// Literals in source code allow to specify fixed values
	// like numbers, strings or booleans.

	// Go, like JS or Py, is able to declare compound types:
	//  arrays, maps, slices and structs.

	// "mySlice" is our slice example.
	mySlice := []int{2, 3, 4, 5}
	fmt.Println(mySlice)

	// myMakedSlice is initialized with 0 values, and it will start with lenght = 10
	// make can be used to initialize arrays, slices and maps
	myMakedSlice := make([]int, 10)
	fmt.Println(myMakedSlice)

	// "myArray" is our array example.
	myArray := [2]int{1, 2}
	fmt.Println(myArray)

	// "myMap" is our map example
	myMap := map[string]int{
		"Name1": 1,
		"Name2": 2,
	}
	fmt.Println(myMap)

	//  With this, we are able to create a new "person", and we call it a struct literal
	p1 := person{
		"John",
		"Marston",
	}
	fmt.Println(p1)

	// And with this, we are constructing a new person, and handling possible error, if it occurs
	p2, err := NewPerson("Dovahkiin", "the Dragonborn")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p2)

	// Calling a function binded to the type
	p1.speak()

	// Flow control
	// Go has only one loop: the for, but it's really flexible.
	// To know more, check the README.md
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for range works like a forEach in other languages
	for index, value := range mySlice {
		fmt.Println("Index:", index, "Value:", value)
	}

	// You can use for range to iterate over a map as well
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
