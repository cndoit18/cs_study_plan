package facade

import (
	"fmt"
	"io"
)

type OrderSystem struct {
	out io.Writer
}

func (o *OrderSystem) Settlement() float64 {
	fmt.Fprintf(o.out, "Settlement completed\n")
	return 20.0
}

type PaymentSystem struct {
	out io.Writer
}

func (p *PaymentSystem) Payment(amount float64) {
	fmt.Fprintf(p.out, "Payment of $%.2f\n", amount)

}

type ShopSystem struct {
	order   *OrderSystem
	payment *PaymentSystem
}

func (s *ShopSystem) Buy() {
	s.payment.Payment(s.order.Settlement())
}

func NewShopSystem(out io.Writer) *ShopSystem {
	return &ShopSystem{
		order:   &OrderSystem{out: out},
		payment: &PaymentSystem{out: out},
	}
}
