// Client code
package observer

import (
	"strings"
	"testing"
)

func Test_Observer(t *testing.T) {
	shirtItem := newItem("Nike Shirt")
	jeansItem := newItem("Adidas Jeans")

	observerFirst := &customer{id: "first@gmail.com"}
	observerSecond := &customer{id: "second@gmail.com"}
	observerThird := &customer{id: "third@gmail.com"}

	shirtItem.register(observerFirst)

	shirtItem.register(observerSecond)
	jeansItem.register(observerSecond)

	jeansItem.register(observerThird)

	shirtItem.updateAvailability()
	jeansItem.updateAvailability()

	isRecivedNikeShirtMail := false
	isRecivedAdidasJeansMail := false
	for _, mail := range observerFirst.inbox {
		if strings.Contains(mail, shirtItem.name) {
			isRecivedNikeShirtMail = true
		} else if strings.Contains(mail, jeansItem.name) {
			isRecivedAdidasJeansMail = true
		}
	}

	if !isRecivedNikeShirtMail || isRecivedAdidasJeansMail {
		t.Error("first customer's inbox is wrong")
	}

	isRecivedNikeShirtMail = false
	isRecivedAdidasJeansMail = false
	for _, mail := range observerSecond.inbox {
		if strings.Contains(mail, shirtItem.name) {
			isRecivedNikeShirtMail = true
		} else if strings.Contains(mail, jeansItem.name) {
			isRecivedAdidasJeansMail = true
		}
	}

	if !isRecivedNikeShirtMail || !isRecivedAdidasJeansMail {
		t.Error("second customer's inbox is wrong")
	}

	isRecivedNikeShirtMail = false
	isRecivedAdidasJeansMail = false
	for _, mail := range observerThird.inbox {
		if strings.Contains(mail, shirtItem.name) {
			isRecivedNikeShirtMail = true
		} else if strings.Contains(mail, jeansItem.name) {
			isRecivedAdidasJeansMail = true
		}
	}

	if isRecivedNikeShirtMail || !isRecivedAdidasJeansMail {
		t.Error("third customer's inbox is wrong")
	}
}
