package abstract_factory

import "testing"

func Test_MotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())
	
	sportBike, ok := motorbikeVehicle.(Motorbike) // 타입 형변환
	if !ok {
		t.Fatal("struct assertion has failed")
	}

	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())
}

func Test_CarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())

	luxuryCar, ok := carVehicle.(Car) // 타입 형변환
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
}