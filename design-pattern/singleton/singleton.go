package sigleton

// interface는 해당 모듈의 모든 타입이 기본적으로 구현되야 하는 메소드를 정의 한다.
type Singleton interface {
	PlusOne() int
	MinusOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) PlusOne() int {
	s.count++
	return s.count
}

func (s *singleton) MinusOne() int {
	s.count--
	return s.count
}
