package shoppingcartservice

type CartManager struct {
	carts map[string]Cart
}

func CreateCartMananger() *CartManager {
	return &CartManager{
		carts: make(map[string]Cart),
	}
}
func (cm *CartManager) CreateCart(userId string) {
	cm.carts[userId] = Cart{}
}
func (cm *CartManager) GetCart(userId string) Cart {
	return cm.carts[userId]
}
func (cm *CartManager) AddToCart(userId string, itemId string, qty int)  {}
func (cm *CartManager) UpdateCart(userId string, itemId string, qty int) {}
func (cm *CartManager) Checkout(userId string)                           {}
