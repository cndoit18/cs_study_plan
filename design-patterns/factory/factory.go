package factory

import "fmt"

// Transport 描述一个运输单位的抽象
type Transport interface {
	Deliver()
}

var _ Transport = &Truck{}

type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("truck deliver")
}

var _ Transport = &Ship{}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("ship deliver")
}

// 简单工厂
func CreateSimpleTransport(t string) Transport {
	if t == "road" {
		return &Truck{}
	}
	return &Ship{}
}

// 抽象工厂
type Factory interface {
	CreateTransport() Transport
}

var _ Factory = &RoadLogistics{}

type RoadLogistics struct {
}

func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

var _ Factory = &SeaLogistics{}

type SeaLogistics struct {
}

func (r *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

func CreateAbstractFactory(t string) Factory {
	if t == "road" {
		return &RoadLogistics{}
	}
	return &SeaLogistics{}
}
