// The Interface Segregation Principle

// This is the simplest principle out of the SOLID design.
// It states that we shouldn't put too much in the interface.
// We shouldn't try to throw everything and sink into just one single interface.
// And then sometimes it makes sense to break up the interface into several smaller ones.

// Lets say we have a type Document that stores some information about the documents,
// and we want to make an interface that allows people to build different machines,
// different constructs for operating on the documents (printing, scanning, sending...)

// So one approach would be just make a single Machine interface

package main

type Document struct {
	// ...
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// Now this is fine if we want that hip multi-function printer, that everybody talks about.

type MultiFunctionPrinter struct {
	// ...
}

func (m MultiFunctionPrinter) Print(d Document) {
	// ...
}

func (m MultiFunctionPrinter) Fax(d Document) {
	// ...
}

func (m MultiFunctionPrinter) Scan(d Document) {
	// ...
}

// Now imagine that we have someone not buying into the hype, and is working
// with a good old fashioned printer. It doesn't have scanning or faxing capabilities.
// But it prints!

type OldFashionedPrinter struct {
	// ..
}

// But, because we want to implement that interface, because maybe
// some other APIs rely on Machine interface, we have to implement this anyway.
// We are being forced into implementation so we go ahead through all those
// similar and pesky notions to implement the Machine and we end up with the same stuff.

func (o OldFashionedPrinter) Print(d Document) {
	// this makes sense, the old one can print
}

// Deprecated...
func (o OldFashionedPrinter) Fax(d Document) {
	// how to fax ?!?
	panic("operation not supported, buy a new one")
}

// Deprecated...
func (o OldFashionedPrinter) Scan(d Document) {
	// how to scan ?!?
	panic("operation not supported, buy a new one")
}

// <- Problem!
// Well those methods are not really deprecated, we're lying to the user a little bit.
// But the hope of having 'deprecated' there is in that some IDE's will in the end
// cross out Scan and Fax functionalities and signal us that we should avoid them.
// No bueno!

// But really, the problem was created because we've put too much into an interface.
// We've put all of that in a single interface [Print, Fax, Scan] and then we now expect
// everyone to implement this even if they actually don't have this functionality.

// To deal with this, we adhere to the -> Interface segregation principle
// Which states that we need to try to break up an interface into separate parts.
// And be sure that those parts are what users will definetely need.

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// And this allows us to compose different types out of interfaces that we'll actually need.

type MyPrinter struct {
	// ...
}

func (m MyPrinter) Print(d Document) {
	//...
}

type Photocopier struct {
	//...
}

func (p Photocopier) Print(d Document) {
	//...
}

func (p Photocopier) Scan(d Document) {
	//...
}

// Don't forget, we can always combine interfaces.

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// What we could also do is, let's say we already have Printer and Scanner
// implemented as separate components, we can use -> the Decorator Design Pattern [more on that latter]

// Decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

// Recap:
// -> With ISP approach we have very granular kind of definitions
// -> We grab interfaces that we need and we don't have any extra members
// -> We always have an ability of combining the interfaces

func main() {

}
