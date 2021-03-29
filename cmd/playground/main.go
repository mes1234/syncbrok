package main

type SimpleQueue struct {
	channels []bool
}

//Add item to end of queue
func (q *SimpleQueue) AddMsg() {
	ch := true // make(chan bool, 1)
	q.channels = append(q.channels, ch)
}

func main() {
	q := SimpleQueue{
		channels: make([]bool, 0),
	}

	for {
		q.AddMsg()
	}
}
