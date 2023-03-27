package abstract_factory

type SportMotorbike struct{}

// s변수를 사용할거면 표시해도 되고 안할거면 안해도 된다.
func (s *SportMotorbike) NumWheels() int {
	return 2
}

func (*SportMotorbike) NumSeats() int {
	return 1
}

// 해당 메소드가 구현이 안되어 있어서 빌드 전에 오류가 나지 않는다.
func (*SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}