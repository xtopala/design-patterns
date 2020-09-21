// The Problems with Singleton

// Now we'll see why singleton is not always the best idea,
// and explore some problems with it.

// And in most cases, the problem is that singleton quite
// often breaks the Dependency Inversion Principle.

// Let us see how this might affect us.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// We'll have everything same as in previous example.

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var instance *singletonDatabase
var once sync.Once

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetSingletonDB() *singletonDatabase {
	once.Do(func() {
		db := singletonDatabase{}
		caps, err := readData(".\\capitals.txt")
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

// So imagine we want to perform some sort of research,
// we want to get the total population of several cities.

func GetTotalPopulation(cities []string) int {
	res := 0
	for _, city := range cities {
		res += GetSingletonDB().GetPopulation(city)
	}
	return res
}

// So it might seem OK, like everything is fine here, and we can
// try testing this.
// And the test obviously passes, but there is a huge problem with
// the test that we are writing right now.

// That test is dependent upon data from a real life database and
// in real life we almost never test against a live database, because
// well these values they are essentially magic values.
// The database can change it at any time, so somebody can go into our
// capitals file and they can change something.
// And then, all of a sudden our test would break.

// Also, there's a performance consideration because even though we just
// want to do a unit test GetTotalPopulation(), but instead our test is turning
// from a unit test into an integration test, because we are actually going into
// the database.
// So we're testing not just that the sum of the cities and their population is
// correct, bu we're also testing that the database reads correctly and stuff.
// Which is totally not what we want!

// We want to be able to supply some sort of fake database with predictable data,
// which doesn't wimply actually reading from a disk.

// So this is the kind of issue that we have right here in that with a singleton.
// This is difficult!

// And the reason in because essentially if we look into GetTotalPopulation function,
// we're depending directly upon the singleton database.
// So there we're depending upon a concrete instance of the singleton database.

// And here is a good place to be reminded of -> the Dependency Inversion Principle

// -> Instead of depending on concrete implementations we want to depend upon
//    abstractions which typically implies depending upon an interface

// And certainly if we decide to depend upon an interface, this problem of reading
// a live database for the purposes of unit testing kind of goes away

func main() {
	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)

	ok := tp == (17500000 + 17400000) // <- has a problem!
	fmt.Println(ok)
}
