package shoppingcartservice

type InventoryManager struct {
	inventory map[string]InventoryItem
}

func CreateInventory() *InventoryManager {
	return &InventoryManager{
		inventory: make(map[string]InventoryItem),
	}
}

func (im *InventoryManager) addProduct(p Product, qty int)     {}
func (im *InventoryManager) removeProduct(p Product, qty int)  {}
func (im *InventoryManager) incrQty(p Product, qty int)        {}
func (im *InventoryManager) decrQty(p Product, qty int)        {}
func (im *InventoryManager) restockProduct(p Product, qty int) {}

// todo: Add locks
