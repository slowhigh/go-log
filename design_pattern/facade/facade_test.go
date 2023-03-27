// Client code
package facade

import "testing"

func Test_Facade(t *testing.T) {
	walletFacade := newWalletFacade("abc", 1234)

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}

	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}
}