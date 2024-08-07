package animals

import "github.com/macondiano100/GoLang/animals/internal/animals"

type Animal = int

const (
	DOG Animal = iota
	CAT
	PARROT
)

func DoAnimal(animal Animal) {
	switch animal {
	case DOG:
		animals.Bark()
	case CAT:
		animals.Meow()
	case PARROT:
		println("Not supported")
	}
}
