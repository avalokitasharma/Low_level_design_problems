package shoppingcartservice

type CartItem struct {
	userId             string
	itemId             string
	qty                int
	pricePerUnit       float64
	discountInPerecent float64
}

func (i *CartItem) GetItemTotalPrice() float64 {
	return float64(i.qty) * (i.pricePerUnit) * i.discountInPerecent
}
func (i *CartItem) SetQty(q int) {
	i.qty = q
}
