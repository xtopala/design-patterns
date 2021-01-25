// Observer - Observer and Observable

// The whole idea behind the Observer design pattern
// is that we have some components which generate certain
// events and other components which want to be notified about
// these events happening.

// So essentially some component does something or
// some information comes into the system and we want
// notifications in other components.

// Our big questin is how can we do this and furthermore
// how can we do this in the most general way possible.

// Let's take a look at how we can implement a general
// porpose kind of framework, most of the times it's just a
// bunch of structures and interfaces to support this idea of
// notification.

package main

import (
	"container/list"
	"fmt"
)

// There's gonna be two participants to this story.
// One is typically called -> Observable and the other
// is called the -> Observer.

type Observable struct {
	subs *list.List // list of Observers, that are connected to this
}

type Observer interface {
	Notify(data interface{})
}

// <- The idea here is that the observer has a method
// 	  which gets notified effectively.
// 	  This method gets called whenever there's something
//	  happening.

// When something happens the Observer has a method -> Notify

// So what do we want the Observable to be able to do?
// We want to be able to subscribe to the events that happen
// on the observable.

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

// Similarly to this, we can have a way of unsubscribing.

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

// <- We remove this Observer from the set, because
// sometimes we want to be notified about some event
// happening but only to a point after which we no longer
// really care and we don't want any notifications to happen.

// And the final method that we want ont the Observable
// is some method for actually firing the event.
// The method that will notify the Observer that something happens.

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

// We've set everything up, but now what we can do is
// we can actually use the Observer interface as well as
// the Observable struct to build some sort of a scenario.

// So let's suppose that we have a person that maybe cathes
// a cold and we just want to have some sort of doctor's service
// to be informed that a person has become ill, and maybe it's time
// to visit them to see what's going on.

type Person struct {
	Observable
	Name string
}

// We'll also make a factory, which churns out people persons.

func NewPerson(name string) *Person {
	return &Person{
		Name:       name,
		Observable: Observable{new(list.List)},
	}
}

// Then, we can simulate the person cathing some disease.

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

// <- This is the place where we want to perform the notification.
// 	  We use the information that we have inside the Observable.

// Now we want some sort of doctor service.

type DoctorService struct{}

// <- We use this so we can implement the Observer interface

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s", data.(string))
}

// <- This is the method that gets called whenever a person falls ill.

// So now that we have this whole scenarion,
// let's see how we can actually use all of it.

// Recap:
// -> We can imagine we can have more than one Observer for a given Observable
// -> We could have different handles for different objects that
//	  inherit the Observable interface

func main() {
	p := NewPerson("Leto Atreides")
	ds := &DoctorService{}
	p.Subscribe(ds)

	p.CatchACold()
}
