package main

import (
	"fmt"
	"testtdd/cart"
)

type Cart struct {
	Totalprice int
}

func main() {
	c := cart.Cart{}
	d := c.TotalPrice()
	fmt.Println(d)
}
