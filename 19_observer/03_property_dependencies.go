// Observer - Property Dependencies

// In the previous example we've been told that there are
// certain problems with using change notifications on properties.
// And by property we mean a combination of a getter and the setter.
// So what are those problems?

// Let's imagine that for a given person we decide to make
// a read only property.
// A propery which is computed from some other value as opposed
// to just being used directly.

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

type PropertyChange struct {
	Name  string // "Age" ; "Height"
	Value interface{}
}

func (p *Person) Age() int { return p.age }

// So for example, let's suppose we want a boolean method
// telling us whether a person can vote.

func (p *Person) CanVote() bool {
	return p.age >= 18
}

// <- Well, the problem is that we want change notification for this as well.
//    But remember, change notifications can only happen in setters.
//    It can't really happen in getters right here.

// So the question is, where exectly would we send the notification
// for the change in the voting status?

// The answer of course is that we would do this in the setter
// for the age, because we're changing it there, and age affects
// the result of CanVote().

// It make sense that not only do we set the notification for age
// but right here somewhere we also set the notification for CanVote().

// Let's imagine, once again, that we have some sort of structure
// for an electoral roll.

type ElectoralRoll struct{}

// And we want to implement the Observer interface on it.
// And let's suppose that we want to notify a person or to notify
// the system when a person's voting status has changed to true, when
// they can finally vote.

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Grats, you can vote!")
		}
	}
}

// <- So this is our handler for the changes in a persons's CanVote()

// But, the question is where exectly would we generate an event?
// Where would we fire an event where the argument is CanVote?
// And the answer is we would have to do this in SetAge.
// There is no other place for us to do it, but in order to be able
// to do this we need to make sure that CanVote has actually changed.

// That the voting status changed from false to true, or other way around,
// because otherwise we don't want to fire any change notification.
// And this is where things get really difficult, because remember the changing
// of the age changes the result of CanVote.

// And if we want to make sure that can vote has in fact changed,
// we need to cache it's previous value and then compare it to the
// new value.
// Anoying, I know.

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

// <- So this is how we would implement the dependent property
// 	  because essentially what we have is we have CanVote which is
//	  a property which depends on the property or indeed the age field.

// Going back to our scenario, let's connect everything together.
// Ok, now this works, but what's the problem?
// There's always something.

// Recap:
// -> The problem here is the problem of dependencies
// -> The problem is that the voting status depends on age and
//	  age gets modified as part of the SetAge setter
// -> The SetAge setter becomes very large, it starts caching previous
// 	  values of all the properties it affects and then compares the previous
// 	  values of the affected properties to the current values and sends the notifications
// -> This does not scale
// -> If we have lots of properties dependent upon the age or if we have one property
// 	  such as CanVote that depends on multiple other different properties then we
//	  end up with a nightmare
// -> And we don't do things this way, we don't define dependency properties inside
// 	  the setters of the properties
// -> Instead we try to build some sort of higher level framework, some sort of map
//	  where all the dependencies between all the different properties are catalogued and
//    then subsequently we iterate through this map and we perform the notifications in a
//	  more regularized way
// -> This example is just a simple illustration of how we can get complexity
//	  virtually out of nowhere and suddenly there is no Go language mechanism that
// 	  would help us with this

func main() {
	p := NewPerson(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to:", i)
		p.SetAge(i)
	}
}
