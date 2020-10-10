// Proxy - Protection Proxy

// This proxy performs access control.
// It basically tries to check whether or not
// the object it's trying to proxy is actually allowed
// to be accessed.

package main

import "fmt"

// So let's suppose that we're simulating the process
// of cars and other wehicles being driven.

type Driven interface {
	Drive()
}

// Now what we can do is we can construct a car.

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

// So this is our starting point.
// Now imagine that we want the car to only be driven
// if we have a driver and if that driver is actually old enough.

// So what we do then is we build a protection proxy on top of the car.
// Once again, we would be reusing the car somehow, but we would also be
// specifying the driver.

type Driver struct {
	Age int
}

// Now we can build a car proxy.

type CarProxy struct {
	car    Car
	driver *Driver
}

// The idea is that whenever somebody wants to make a new car
// what they in fact get is they get this car proxy.

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

// So this car proxy, just like the ordinary car, it has to
// implement that driven interface.

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}

// So this example shows a very common kind of pattern
// where first of all we're starting out with an object or
// some sort of structure which can be use as is without any
// verification but also subsequently we want to have additional
// verification, some additional checks being made whenever somebody
// actually uses this struct.

// We can have it as a new factory function
// and we can also introduce dependencies, so we need to remember
// the original car doesn't have any parameters we need to feed it.
// It doesn't have anything at all, whereas here we have to explicitly
// specify the driver, without which we wouldn't be able to even construct
// this structure in the first place.

func main() {
	// car := NewCarProxy(&Driver{12})
	car := NewCarProxy(&Driver{22})
	car.Drive()
}
