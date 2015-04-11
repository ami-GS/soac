package main

import (
	"fmt"
	"github.com/ami-GS/soac"
)

func main() {
	c := soac.NewChanger()
	fmt.Printf("int out = %s\n", c.Underline().White().Apply(123456))
	c.Reset()
	fmt.Printf("%s is %s\n", c.Bblue().Red().Bold().Apply("This"), c.Apply("continuing!!"))
	c2 := soac.NewChanger()
	c2.Bred().White().Underline()
	fmt.Printf("%s is %s\n", c2.Apply("This"), c.Apply("combination!!"))
}
