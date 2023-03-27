// Client code
package decorator

import "testing"

func Test_Decorator(t *testing.T) {
	pizza := &veggieMania{} //15

	pizzaWithCheese := &cheeseTopping{ // 15 + 10
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &tomatoTopping{ // 15 + 10 + 7
		pizza: pizzaWithCheese,
	}

	if pizzaWithCheeseAndTomato.getPrice() != (15 + 10 + 7) {
		t.Error("Wrong total price.")
	}
}