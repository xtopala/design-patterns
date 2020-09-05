// The Single Responsibility Principle

// SRP states that type should have one primary responsibility.
// And, as a result, it should have one result it should have one reason to change.
// That reason being somehow related to it's primary responsibility.

package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)

	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// Separation of Concerns (another term in addition to SRP):
// This basically means that different concers, or different problems,
// that system solves have to reside in different constructs.
// So whether they're attached to different structures or residing in different packages.
// That's on us to decide, but they have to be split up.

// God Object - that's an Anti-pattern to SoC

// example to this would be:
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	//...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	//...
}

// <- There are bad! These are breaking SRP.
// Beacause the responsibility of the Journal is to deal with the entries.
// It's not to deal with persistance.

// This seems conviniet to do, put all logic into one place.
// But imagine other types in a system which also need to be written to a persistance of some sort.
// And there are some common settings to the way we do it, for both Journal and other types.
// So that is one of the reasons why want to take persistance information and just
// put it in a seperate component (package) or seperate type.

// Let's say we go with Package approach ->

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// <- So this is not any more a method on Journal, this exists by itself
// We could call this module 'persistance' with different settings for Line Seperator

// In addition to this we could turn persistance to a struct ->

type Persistance struct {
	lineSeparator string
}

func (p *Persistance) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

// Just to recap:
// -> Avoid cross-cutting concerns <-
// -> This way we have common settings controled in just one place <-

func main() {
	j := Journal{}
	j.AddEntry("I ate a bug")
	j.AddEntry("I cried today")

	fmt.Println(j.String())

	// so we could save it with package
	SaveToFile(&j, "journal.txt")

	// or with separate Struct
	p := Persistance{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
