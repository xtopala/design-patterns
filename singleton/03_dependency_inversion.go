// Sinlgeton and Dependency Inversion

// We really want our tests to depend on some sort of abstraction,
// so that instead of depending on the concrete singleton database,
// we can substitute this database with something else.

// We need to make a dummy database.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// Well first of all, we need to introduce some sort of abstraction
// which has something in common between the real one the dummy one.

type Database interface {
	GetPopulation(name string) int
}

// And we'll use stuff, we previously wrote.
// Now we already have Singleton database implement our new interface.

// So now what we need is to slightly change the wat that we calculate
// the total population so that we now introduce a dependency.

// Our old GetTotalPopulation() is unfortunately hard coded to use
// the singleton, and that means there's no way for us to substitute
// something in there.

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

func GetTotalPopulation(cities []string) int {
	res := 0
	for _, city := range cities {
		res += GetSingletonDB().GetPopulation(city)
	}
	return res
}

// <- We need a better function than this.

func GetTotalPopulationEx(db Database, cities []string) int {
	res := 0
	for _, city := range cities {
		res += db.GetPopulation(city)
	}
	return res
}

// Right now, we provide a database as opposed to just expecting
// function will find one for us.
// Now we have additional flexibility.

// So if we do really want to test against the real life database, we could.
// But really what we want to be able to do is to get away from using real life
// databases and use some dummies.

// What we want is we want a database with predictable values.
// We just want a map of let's say three values [alpha, beta, gamma]

// So we'll make a new type here.

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

// So now we can write a proper unit test, not some flaky integration test.

// Takeaway:
// -> The Singleton isn't really that scary
// -> The scary part is depending directly on the Singleton as opposed to
//    depending on some interface that the Singleton implements

// Because if we depend on the interface we can substitute that interface.
// But of course, in order to be able to substitute something we have to provide
// an API aware that something can be plugged in.

// And that's exactly what we did by making our modification of the GetTotalPopulation()
// which now takes the database up on which it operates.

func main() {
	// cities := []string{"Seoul", "Mexico City"}
	// tp := GetTotalPopulationEx(GetSingletonDB(), cities)

	// ok := tp == (17500000 + 17400000)
	// fmt.Println(ok)

	// <- No good, we need new approach!

	names := []string{"alpha", "gamma"} // <- This will not change, no brittle test
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}
