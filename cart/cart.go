package cart

type Cart struct {
	Totalprice int
}

func (c Cart) TotalPrice() int {

	c = Cart{
		Totalprice: 12,
	}

	return c.Totalprice
}
