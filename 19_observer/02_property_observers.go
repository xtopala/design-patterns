// Observer - Property Observers

// One very common use of the Observer design pattern
// is to implement notifications related to property changes
// on some object.

// A struct might change somehow and we want to be
// informed that there is in fact a change.
// Let's checkout how this can work.

// We'll use some stuff from previous example, but
// we'll introduce some modifications.

package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

type Observer interface {
	Notify(data interface{})
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

// Instead of specifying the person's name,
// let's suppose that we want to be informed when
// a person's age changes.

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

// <- This is how we make a person ¯\_(ツ)_/¯

// Now, imagine that we want to get notifications about
// a person getting older or indeed getting younger, if they've
// found some magical fountain.

// Well, we have to send information about the person's property
// being changed, but what do we mean by property?
// Notice that we've made -> age <- lowercase.
// The idea of properties is that a property is a combination
// of a getter and a setter.

// So for the getter, we would have a method called Age()
// and for the setter we would have a method called SetAge()

// Properties are not very idiomatic in Go,
// unless of course we want to do the kind of thing that
// we're doing here.

// We just change notifications, because the idea is that
// when we change the age, when we call SetAge(), not only do
// we change the actual value, but we also want to use this as
// an Observable and get it firing some sort of event telling others
// that the age has changed.

// And this might only be one set of properties that we want
// to notify about.

// So when we have a change notification, when we
// want to tell somebody that something has changed, we can
// have a separate struct for encoding this information.

type PropertyChange struct {
	Name  string // "Age" ; "Height"
	Value interface{}
}

func (p *Person) Age() int { return p.age }

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

// Now let's suppose that we have some
// traffic menagement company or system of whatever.
// The Traffic Menagement wants to be informed about
// a person's age and if the person is too young to drive
// they would keep monitoring their age, but as soon as a person
// turn's certain age, they no longer care about person's age.
// The Traffic Menagement congratulates the person on being able to
// drive and then they unsubscribe from the Observable.

type TrafficMenagement struct {
	o Observable
}

func (t TrafficMenagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("Grats, you can drive now!")
			t.o.Unsubscribe(t) // <- this look hella weird, but it's ok
		}
	}
}

// Recap:
// -> We've been able to demonstrate, horefully, that there is a case
//	  for using properties, meaning a combination of a getter and a setter
//	  as opposed to just ordinary fields
// -> That case is when we want notifications of property changes
// -> However, this approach does have certain problems with dependencies,
//	  and that's what we're going to take a look at in the next example

func main() {
	p := NewPerson(15)
	t := TrafficMenagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting the age to:", i)
		p.SetAge(i)
	}
}
