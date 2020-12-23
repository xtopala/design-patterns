// Chain of Responsibility - Method Chain

// Here we'll take a look at something that might
// get called a Method chain because, well, there's gonna
// be a linked list of method invocations.

// Imagine that we're working on some sort of a game.
// The game has a bunch of creatures and these creatures are
// participating in it doing all sort of messed up things we can imagine,
// but mostly engage in combat with other creatures.

package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

// So with this setup what we want to be able to do
// in our game is we want to be able to applt modifiers to a creature.
// For example, a creature might be roaming the grounds and it
// picks up a magic sword, and starts going Rambo on everybody left and right.
// Or, it tries to eat a mushroom and gets poisoned.

// We want to modify the aspects of the creature and we want to
// build a stack or a list of these modifications, so that they can be
// applied to a creature one after another.

// In other words, we want to be able to have some sort of modifier.
// And idea is that we can apply more than one modifier to a creature and
// whenever we apply additional modifiers we can stick them on the end of an
// already existing modifier.

type Modifier interface {
	Add(m Modifier)
	Handle()
}

// With the interface sorted out, we also need some sort of concrete type.

type CreatureModifier struct {
	creature *Creature
	next     Modifier // a linked list of modifiers [singly linked list]
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle() // <- this doesn't do much, just calls every single element
		//	  				  one after another
	}
}

// Why are we doing this?
// Well, we're going to aggregate this type and we're going to actually
// make use of all of this stuff.

// So, now that we have a Creature modifier lets actually make a constructor for it.

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

// ↑↑↑ We don't specify the next modifier, because that needs to be added later.

// Now, we want something concrete.
// A modifier that actually does something to a creature.
// For an example, let's suppose there's an item that the creature
// can pick up, which doubles the creatures attack.

type DoubleAttackModifier struct {
	CreatureModifier
}

// -> What we do here is we're simply aggregating creature modifier here

// But now we can have a constructor:

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{
		CreatureModifier{
			creature: c,
		},
	}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle() // <- we propagate the application of every single
	//								  one of these modifiers in the hierarchy
}

// Now, let's suppose we want to have some sort of increased defense modifier.

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(c *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{
		CreatureModifier{
			creature: c,
		},
	}
}

func (i *IncreaseDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("Increasing", i.creature.Name, "\b's defense")
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}

// What if at some point goblin gets hit with a spell
// and that spell disables any of the other modifiers.
// How can we actually disable this entire list of modifiers?

type NoBufsModifier struct {
	CreatureModifier
}

func NewNoBufsModifier(c *Creature) *NoBufsModifier {
	return &NoBufsModifier{
		CreatureModifier{
			creature: c,
		},
	}
}

func (d *NoBufsModifier) Handle() {
	// empty for a reason
}

// ↑↑↑ This means that every single modifier which comes after this one
// 	   is actualy not going to be applied, which we wanted in the 1st place

// This modifier has prevented the traversal of the entire linked list,
// thereby preventing the application of any other modifier.

// Recap:
// -> This implementation of Chain of Responsibility is called Method Chain
//	  because we're following a linked list of these modifiers and we're calling
//	  that magical Handler method on it
// -> We can also prevent the invocation of this Handler method if we decide
//	  that we don't want to invoke it from whatever implementer we have constructed

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	root.Add(NewNoBufsModifier(goblin))

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreaseDefenseModifier(goblin))
	root.Handle()

	fmt.Println(goblin.String())
}
