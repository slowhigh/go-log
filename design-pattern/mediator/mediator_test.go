// Client code
package mediator

import "testing"

func Test_Mediator(t *testing.T) {
	stationManager := newStationManager()

	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}

	freightTrain := &freightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()

	if stationManager.isPlatformFree {
		t.Error("PassengerTrain has not arrived.")
	}

	if len(stationManager.trainQueue) != 0 {
		t.Errorf("Train Queue : %v", stationManager.trainQueue)
	}

	freightTrain.arrive()

	if len(stationManager.trainQueue) != 1 {
		t.Errorf("Train Queue : %v", stationManager.trainQueue)
	}

	passengerTrain.depart()

	if len(stationManager.trainQueue) != 0 {
		t.Errorf("Train Queue : %v", stationManager.trainQueue)
	}
	
	freightTrain.depart()
	
	if !stationManager.isPlatformFree {
		t.Error("FreightTrain has not departed.")
	}
}