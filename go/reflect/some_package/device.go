package some_package

import (
	"fmt"
	"math/rand"
)

type Device struct {
	id       int
	Name     string
	Location [2]float64
	Address  string
}

func NewDevice(name, address string, location [2]float64) Device {
	return Device{
		id:       rand.Int(),
		Name:     name,
		Location: location,
		Address:  address,
	}
}

func (d *Device) change(p [2]float64, name string) (string, error) {
	d.Location = p
	d.Name = name
	fmt.Println("device location is", d.Location)
	return d.Name + d.Address, nil
}


func (d *Device) Change(p [2]float64, name string) (string, error) {
	d.Location = p
	d.Name = name
	fmt.Println("device location is", d.Location)
	return d.Name + d.Address, nil
}