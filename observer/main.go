package main

func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@email.com"}
	observerSecond := &Customer{id: "xyz@email.com"}
	observerThird := &Customer{id: "efg@email.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()

	shirtItem.register(observerThird)

	shirtItem.updateAvailability()

	shirtItem.deregister(observerFirst)

	shirtItem.updateAvailability()
}
