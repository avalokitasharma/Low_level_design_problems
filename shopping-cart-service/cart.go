package shoppingcartservice

import "time"

type Cart struct {
	id        string
	userId    string
	items     map[string]CartItem
	createdAt time.Time
}

func (c *Cart) addItem(cartItem CartItem) {
	c.items[cartItem.itemId] = cartItem
}
func (c *Cart) removeItem(id string) {
	delete(c.items, id)
}
func (c *Cart) updateItem(id string, qty int) {
	item := c.items[id]
	item.SetQty(qty)
	c.items[id] = item
}
func (c *Cart) getTotalPrice() float64 {
	return 0
}
func (c *Cart) isExpired() bool {
	return false
}
