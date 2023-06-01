package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type (
	ITransport interface {
		setName(n string)
		setSpeed(s int)
	}
	transport struct {
		name  string
		speed int
	}
	scooter struct {
		transport
	}
	bike struct {
		transport
	}
)

func (t *transport) setName(n string) {
	t.name = n
}

func (t *transport) setSpeed(s int) {
	t.speed = s
}

func newScooter() *scooter {
	return &scooter{
		transport: transport{
			name:  "Scooter",
			speed: 4,
		},
	}
}

func newBike() *bike {
	return &bike{
		transport: transport{
			name:  "Bike",
			speed: 9,
		},
	}
}

func getTransport(tt string) (ITransport, error) {
	if tt == "scooter" {
		return newScooter(), nil
	}
	if tt == "bike" {
		return newBike(), nil
	}
	return nil, fmt.Errorf("Wrong type")
}

func main6() {
	scooter, _ := getTransport("scooter")
	bike, _ := getTransport("bike")

	fmt.Println(scooter)
	fmt.Println(bike)
}
