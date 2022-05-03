package builder

import "testing"

type testCase struct {
	structure string
	wheels    int
	seats     int
}

func TestBuilderPattern(t *testing.T) {
	testCaseArr := []testCase{
		{
			structure: "Car",
			wheels:    4,
			seats:     5,
		},
		{
			structure: "Bike",
			wheels:    2,
			seats:     1,
		},
	}

	for i, tc := range testCaseArr {
		vehicle := New()
		buildedVehicle := vehicle.SetStructure(tc.structure).SetWheels(tc.wheels).SetSeats(tc.seats).Build()

		if buildedVehicle.GetStructure() != tc.structure {
			t.Errorf("Wrong structure. - %d", i)
		}

		if buildedVehicle.GetWheels() != tc.wheels {
			t.Errorf("Wrong wheels. - %d", i)
		}

		if buildedVehicle.GetSeats() != tc.seats {
			t.Errorf("Wrong seats. - %d", i)
		}
	}
}
