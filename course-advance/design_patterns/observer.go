package main

import (
	"fmt"
	"time"
)

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item  -> No available
// Notify when item is available

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

type EmailClient struct {
	id string
}

func (e *EmailClient) getId() string {
	return e.id
}

func (e *EmailClient) updateValue(name string) {
	fmt.Printf("Email to %s, item %s is available\n", e.id, name)
}

func main() {
	nVidiaItem := NewItem("RTX 3080")

	fistObserver := &EmailClient{id: "12ab"}
	secondObserver := &EmailClient{id: "34cd"}

	nVidiaItem.register(fistObserver)
	nVidiaItem.register(secondObserver)
	fmt.Printf("Item %s is not available\n", nVidiaItem.name)
	time.Sleep(3 * time.Second)
	nVidiaItem.UpdateAvailable()

}
