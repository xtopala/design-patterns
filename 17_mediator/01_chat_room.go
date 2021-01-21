// Mediator - Chat Room

// The best and most succint example of the
// -> Mediator design pattern is the simulation
// of a chatroom.
// And that's exactly what we're going to build.

package main

import "fmt"

// So, we'll begin by defining a participant of the chat room.

type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

// Let's construct that person.

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

// And then what we're going to have a bunch of methods on a Person.
// First of all, what we need to be able to do is we need
// to be able to receive a message from somewhere.
// We can receive it from another person or we can receive it
// from the Chatroom itself.

// You get kicked out, the chatroom can tell you:
// -> You received a ban hammer <-
// Or maybe the room got disconnected or deleted or something.

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s\n", sender, message)
	fmt.Printf("[%s's chat session] %s", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

// Then we'll also have a way of actually saying somethin in the room.

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

// As we can see what's happening here is that every
// single Person, every single participant in the chatroom
// has a pointer to the room they're in.

// So they can join the room and when they join it,
// they can receive messages and they can also send messages
// whether it's messages for an entire room or messages for
// specific individuals.

// Ok, now we need an actual room.

type ChatRoom struct {
	people []*Person
}

// What we want to do with this is that we can first of all
// support the idea of broadcasting.
// So we can specify a message as well as the source of
// that message and we can send it to every participant in the Chatroom.

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

// Also, we need to send a targeted message from
// one participant to another.

func (c *ChatRoom) Message(source, destination, message string) {
	for _, p := range c.people {
		if p.Name == destination {
			p.Receive(source, message)
		}
	}
}

// And the final thing we want to add is a method for
// actually getting a person to join the Chatroom.

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

// Now we have everything ready, and we can try out
// our little chatroom.

// Recap:
// ->	Here, the Mediator was the Chatroom
// ->	Bunch of people in the Chatroom, but they're not
//		directly aware of one another
// ->	Meaning, they don't have pointers to one another
// ->	Meadiator, on the other hand, is able to inform every
//		participant of something happening
// -> 	Mediator is just a central component that everyone knows
// 		and everyone can connect to
// -> 	People can communicate with one another without being afraid
//		that for example they will call something on a nil pointer
// -> 	If we had a private message where we had to have a pointer
//		to the person we're sending the private message to we could run
//		into a situation where that person no longer exists, because maybe
// 		they left the system altogether
// ->	So the mediator here provides an additional layer of safety as well

func main() {
	room := ChatRoom{}

	fatAbbot := NewPerson("Fat Abbot")
	rudy := NewPerson("Rudy")

	room.Join(fatAbbot)
	room.Join(rudy)

	fmt.Println(fatAbbot.Name)

	fatAbbot.Say("Hey hey hey. What's goin' on, Rudy?")
	rudy.Say("Man, Fat Abbot, you need to lose weight!")
	fatAbbot.Say("You know somethin', Rudy? You're like school in summertime.")
	rudy.Say("School in summertime?")
	fatAbbot.Say("Yeah, b***h, school in summertime!")

	yolanda := NewPerson("Yolanda")
	room.Join(yolanda)

	fatAbbot.Say("Hey Yolanda. Why is your eye all black and blue and s**t?")
	yolanda.Say("Maaan, Fat Abbot. My stepdad popped me in my eye.")
	fatAbbot.Say("Stepdad? You gotta off his ass!")
	yolanda.Say("Really?")
	fatAbbot.Say("Yeah, b***h! Snatch his ass in a bear trap! Leave that m***erfu***r swingin' from a tree so high nobody finds him for days! Glock-glock, you know what I'm sayin'? Dumbassed m***erfu***r pullin' s**t! Damn!")

	yolanda.PrivateMessage("Fat Abbot", "You're right, Fat Abbot. Thanks!")
}
