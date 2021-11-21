package firstpackageingo

import "fmt"

const Description = "This is my first package in go!"

var Lion = "Leo"

type Felidae struct {
	Name  string
	Color string
}

func (f *Felidae) Reset() {
	f.Color = ""
	f.Name = ""
}

func (f Felidae) String() string {
	return fmt.Sprintf("The lion %s and color %s is a large mammal of the Felidae family", f.Name, f.Color)
}

func Hello() {
	fmt.Println("Hello World!")
}
