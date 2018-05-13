package main

import "fmt"

const (
	size = 10
)

//Que -- sdfdsf
type Que []string

var currPtr int

//MakeQ - make it
func MakeQ() Que {
	q := make(Que, size)
	return q
}

//Enqueue - Enqueue
func (q *Que) Enqueue(s string) {
	fmt.Println("enqueue ", s)
	if currPtr+1 < size {
		currPtr++
		(*q)[currPtr] = s
	}

}

//Dequeue - dequeue
func (q *Que) Dequeue() string {
	fmt.Println("dequeue ")
	var s string
	if currPtr > -1 {
		s = (*q)[currPtr]
		currPtr--
	}
	return s
}

//CurrSize - cs
func (q *Que) CurrSize() int {
	fmt.Println("curr size")
	return currPtr + 1
}

func main() {

	q1 := MakeQ()
	var input string
	fmt.Println("Enter values - end with EOF")
	for input != "EOF" {
		fmt.Scanln(&input)
		q1.Enqueue(input)
	}
	for q1.CurrSize() > 0 {
		fmt.Println(q1.Dequeue())
	}
}
