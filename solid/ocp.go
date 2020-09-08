// The Open-Closed Principle

// OCP states that types should be open for extension, but closed for modificaiton.
// This is going to be demonstrated with the help of the enterprise pattern called:
// Specification

// Imagine here that we're operating a store, online store.
// We want to allow users to filter our items. And we were given a
// specification, product description, for all criterias that we need to filter.

package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Now imagine we've implement this filtering by color.
// This goes to production, and managers come back, and the say:
// "Doough, we also need filtering by size, sry yeah."
// So this means we have to jump to our Filter type, and change it.
// We are changing the Filter type by adding anothet method to it [FilterBySize].

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

// And then some time passes, and managers activate their thinking parts.
// They conlude that since we have filtering by size and color, why not have both.
// At the same time !!!

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Everything seems fine and countinues to work, but this shows a violation.
// The violation of the OCP.
// Because, what we're doing as we're going back and adding additional methods,
// we are interfering with something that's already been written and tested.

// The OCP is all about beeing open for extension.
// We want to be able to extend a scenarion by adding types, or free standing functions.
// But without modyfing something thats already written and tested.
// Leave Filter type alone! Stop adding methods to it.

// What we need is extendible setup.
// And that is what we can get if we use Specification pattern.

type Specification interface {
	IsSatisfied(p *Product) bool
}

// So the idea behind this interface is that we're testing wheter or not
// a product satisfies some criteria. So if we want to check for color:

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

// Also, check for size:
type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// Now we'll create a somewhat better filter.
// It's better, because we're unlikely to ever modify.
// Fact: if we don't have any specific settings for our filter, we don't need type.
// We could have a free standing function, but we're using types just to illustrate.
type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

// As a consequence, things do become a bit more involved.
// We need to build those specifications.
// But, we're a lot more flexible.

// So if want to filter by some other new criteria,
// all we need to do now is just create a new Specification.

// Recap:
// -> The interface type is open for extension
// -> But it's closed for modification
// -> We're unlike to ever modify Specificaiton interface or BetterFilter

// But wait!
// We had that briliant idea to filter by Size and Color.
// To implement that we need a composite Specification

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	booger := Product{"Booger", green, small}
	lbge := Product{"Large Big Green Egg", green, large}
	whale := Product{"Whale", blue, large}

	products := []Product{booger, lbge, whale}
	fmt.Printf("Green products (old):\n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large green products:\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
