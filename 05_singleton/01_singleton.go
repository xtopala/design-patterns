// The Singleton

// Before we start looking at problems that people
// encounter with a Singleton, we need to talk about
// how to actually construct a singleton and what is the
// basic motivation for doing so.

// So imagine we have some database with a bunch of cities
// as well as their respective populations, and we want to load
// it into memory.

// And oviously, if we're going to do it, we want to do it once.
// We need only one instance of whatever struct is actually keeping
// this memory.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

// <- Notice that this is private
// So that is a hint that this shouldn't be created directly.
// I should be done by using some sort of factory function for initializing.

// Now we need some utility method for actually getting to our data.

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// Now the problem is that we want people to only ever access one instance
// of this struct.
// We can do this by using something like this.

// Now there are actually different options here.

// But, we need a feature -> Thread Safety
// In terms of making the whole thing thread safe, because thread
// safey is important here. We don't want multiple threads to start
// initializing this object at the same time.
// We want control for it.

// One of those options is to use sync.Once, and the other option
// is using just the package level init function.

// sync.Once init() -- thread safety

var instance *singletonDatabase
var once sync.Once

// <- This construct ensures that something gets called only once
//	  and then it's done.

// Another feature that want is -> Lazines
// This basically means that we only construct the database,
// we only read it from a disk to memory, whenever somebody asks for it.

// So lazines is not going to be guaranteed in the init function, unfortunately,
// but it can be guaranteed using sync.Once inside our own function.

// Now we'll need a fucntion for actually getting or creating
// that database of ours.

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

// <- Here we keep returning the instance pointer if
// 	  there are several callers who are attempting to get data f
//    rom this db.

func main() {
	db := GetSingletonDB()
	pop := db.GetPopulation("Seoul")

	fmt.Println("Population of Seoul = ", pop)
}
