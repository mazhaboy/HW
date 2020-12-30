package main

import "fmt"

// implement QUEUE

var size = 0
var queue = new(Node)


func main() {
	Push(queue, 1)
	Push(queue, 2)
	Print(queue)
	Push(queue, 3)
	Print(queue)
	Pop(queue)
	fmt.Println(queue.value)

}

type Node struct {
	value int
	next *Node
}

func Print(queue *Node) {
	q := queue
	for q != nil {
		fmt.Println("value: ", q.value)
		q = q.next
	}

}

// 1 2 3
// 1 4

func Push(queue *Node, value int) bool {
	n:= &Node{
		value: value,
		next:  nil,
	}

	if size == 0 {
		queue.value = value  	// queue.next = n >> queue.value = value
		size++
	} else {					// add else
		for queue.next!=nil{
			queue=queue.next
		}
		queue.next=n  			// add queue.next=n
		size++
	}

	return true
}
// При такой реализации QUEUE функция Pop невозможна
func Pop(queue *Node) (int, bool) { //Мы не можем на прямую с помощью функций менять адрес ссылки
	if size == 0 {
		return 0, false
	}
	t:=queue.next
	queue=nil //Например queue=nil
	queue=t
	return queue.value, true
}

