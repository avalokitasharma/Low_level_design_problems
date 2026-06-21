package shoppingcartservice

type InventoryItem struct {
	itemId  string
	qty     int
	product Product
	price   float64
}
