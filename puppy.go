package puppy

import (
	"fmt"

	dog "github.com/kimhu11/dog1"
)

func Version() {
	fmt.Println("Version v1.1.0")
}

func Bark() string {
	return "woof woof!"
}

func DogBark() string {
	return dog.WhenGrownUp(Bark())
}
