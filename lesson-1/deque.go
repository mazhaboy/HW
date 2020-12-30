package main

import (
	"fmt"
)
//implement DEQUEUE
type Deque struct {
	front *Node
	back *Node
}
type Node struct {
	value interface{}
	next *Node
}


func main() {
	deque:=NewDeque()
	PushFront(deque,3)
	PushFront(deque,2)
	PushFront(deque,1)
	PushBack(deque,4)
	PushBack(deque,5)
	PushBack(deque,6)
	Print(deque)
	fmt.Println(PopFront(deque))
	fmt.Println(PopBack(deque))
	Print(deque)
	PushFront(deque,"start")
	PushBack(deque, "finish")
	Print(deque)
	Peek(deque.front)
	Peek(deque.back)


}
func Print(list *Deque){
	l:=list.front
	for l!=nil{
		fmt.Println("value: ",l.value)
		l=l.next
	}
}
func NewDeque()*Deque{
	return new(Deque)
}
func PushFront(deque *Deque, value interface{}){
	new_node:=&Node{
		value: value,
		next:nil,
	}
	if deque.front==nil && deque.back==nil{
		deque.front=new_node
		deque.back=new_node
	} else {
		x:=deque.front
		deque.front=new_node
		deque.front.next=x
	}
}
func PushBack(deque *Deque, value interface{}){
	new_node:=&Node{
		value: value,
		next:nil,
	}
	if deque.front==nil && deque.back==nil{
		deque.front=new_node
		deque.back=new_node
	} else {
		deque.back.next=new_node
		deque.back=new_node
	}
}
func PopFront(deque *Deque)interface{}{

	if deque.front==nil && deque.back==nil{
		return nil
	} else {
		x:=deque.front.next
		deleted:=deque.front.value
		deque.front=nil
		deque.front=x
		return deleted
	}
}
func PopBack(deque *Deque)interface{}{
	if deque.front==nil && deque.back==nil{
		return nil
	} else {
		y:=deque.front
		x:=deque.back
		for y.next!=x{
			y=y.next
		}
		deque.back=nil
		deque.back=y
		deque.back.next=nil
		return x.value
	}
}
func Peek(node *Node){
	if node!=nil {
		fmt.Println("value: ", node.value)
	} else {
		fmt.Println("value: ", "null")
	}
}
//value:  1
//value:  2
//value:  3
//value:  4
//value:  5
//value:  6
//1
//6
//value:  2
//value:  3
//value:  4
//value:  5
//value:  start
//value:  2
//value:  3
//value:  4
//value:  5
//value:  finish
//value:  start
//value:  finish

