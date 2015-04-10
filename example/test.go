package main

import (
	"fmt"
	"github.com/ami-GS/soac"
)

func main() {
	c := soac.NewChanger()
	fmt.Printf("%s\n", c.Set(1).Underline().White().Out())
	c.Reset()
	fmt.Printf("%s\n", c.Set(1).Out())
	fmt.Printf("%s\n", c.Set("aiueoKAKIKU!!").Yellow().Underline().Bblue().Out())
	c.Reset()
	fmt.Printf("%s\n", c.Set("aiueoKAKIKU!!").Underline().Yellow().Out())
}
