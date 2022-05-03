package builder

type vehicle struct {
	structure string
	wheels    int
	seats     int
}

type Vehicle interface {
	GetStructure() string
	GetWheels() int
	GetSeats() int
}

func (v *vehicle) GetStructure() string {
	return v.structure
}

func (v *vehicle) GetWheels() int {
	return v.wheels
}

func (v *vehicle) GetSeats() int {
	return v.seats
}

type vehicleBuilder struct {
	structure string
	wheels    int
	seats     int
}

type VehicleBuilder interface {
	SetStructure(structure string) VehicleBuilder
	SetWheels(wheels int) VehicleBuilder
	SetSeats(seats int) VehicleBuilder
	Build() Vehicle
}

func (vb *vehicleBuilder) SetStructure(structure string) VehicleBuilder {
	vb.structure = structure
	return vb
}

func (vb *vehicleBuilder) SetWheels(wheels int) VehicleBuilder {
	vb.wheels = wheels
	return vb
}

func (vb *vehicleBuilder) SetSeats(seats int) VehicleBuilder {
	vb.seats = seats
	return vb
}

func (vb *vehicleBuilder) Build() Vehicle {
	return &vehicle{
		structure: vb.structure,
		wheels:    vb.wheels,
		seats:     vb.seats,
	}
}

func New() VehicleBuilder {
	return &vehicleBuilder{}
}
