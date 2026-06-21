package shoppingcartservice

func init() {
	im := CreateInventory()
	cm := CreateCartMananger()

	// Add products to inventory
	product1 := Product{Id: "1", Name: "Laptop", Price: 999.99}
	product2 := Product{Id: "2", Name: "Mouse", Price: 29.99}

	im.addProduct(product1, 10)
	im.addProduct(product2, 20)

	// Create a cart and add items
	cm.CreateCart("user1")
	cart1 := cm.GetCart("user1")

	cm.AddToCart(cartID, product1, 2)

}
