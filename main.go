package main

import (
	"fmt"
	"testtdd/cart"
)

type Cart struct {
	Totalprice int
}

func main() {
	//ok
	c := cart.Cart{}
	d := c.TotalPrice()
	fmt.Println(d)
}
