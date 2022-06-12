package cart

import (
	"testing"
)

func TestCartEmpty(t *testing.T) {
	c := Cart{}
	if c.TotalPrice() != 0 {
		t.Errorf("got %v, want %v", c.TotalPrice(), 0)
	}
}

// func TestCart_TotalPrice(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		c    Cart
// 		want int
// 	}{
// 		{"TestTotalPrice",
// 			Cart{},
// 			12},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.c.TotalPrice(); got != tt.want {
// 				t.Errorf("Cart.TotalPrice() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
