// The Composite Pattern - Illustration through Neural Networks

// We're going to try and implement very simple neurons.

package main

// Going to start this with a deffinition of a Neuron:

type Neuron struct {
	In, Out []*Neuron
}

// Now, we want to able to connect one neuron to another:

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

// Situation gets more complicated and it's not convinient to
// work with when we want to work with whole neuron layers, as
// those are basically a collection of neurons all stored together.

type NeuronLayer struct {
	Neurons []Neuron
}

// We can have a factory function for making a nuron layer
// of a particular size.

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

// What we want is if we could somehow have a single method
// for connecting neurons to neurons, neurons to layers, layers to neurons,
// and layers to layers. Pfeew !

// How can we get this to work considering that we need to somehow
// be able to iterate every single neuron inside either a neuron
// [Yeah, this sounds weird -> iterate every single neuron inside a
// scalar object] or inside a neuron layer.
// Indeed, within neuron layers it's a bit easier.

// So imagine we have some sort of interface [Neuron Interface]

type NeuronInterface interface {
	Iter() []*Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}

	return result
}

// <- This is how we would collect a pointer to every single Neuron
//    inside in our own layer. That's simple enough.

// But another problem is we need to have the same thing on a neuron.

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

// Since we've implemented this interface in both a scalar object,
// which has a single neuron, as well as a collection object, which
// is a neuron layer, we can write a single unifying Connect function.

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

// Now this workd because Iter() is defined on any type which
// implements NeuronInterface, so when we go and iterate a neuron layer
// we get every single pointer of that neuron layer.
// But if we iterate a scalar object, if we iterate a signle neruon
// what happens is we just get a pointer to ourselves. A lonely Neuron.

// In a way, we get a scalar object to masquerade as if it were a collection.

// Now those 4 calls in main() become valid, they become completely legal
// because it no longer matters what we're passing in so long as every single
// type that we're passing in there implements that neuron interface.

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)
	// ↑↑↑ We want to be able to do this with 1 function, not 4
}
