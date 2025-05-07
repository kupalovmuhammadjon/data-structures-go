package main

import (
	"log"

	datastructuresgo "github.com/kupalovmuhammadjon/data-structures-go"
)

func main() {
	doublyLinkedList := datastructuresgo.NewDoublyLinkedList()

	doublyLinkedList.Append(1)
	doublyLinkedList.Append(2)
	doublyLinkedList.Append(3)
	doublyLinkedList.Prepend(4)

	log.Println(doublyLinkedList.Find(4))
	doublyLinkedList.Delete(4)
	log.Println(doublyLinkedList.Find(4))
	doublyLinkedList.InsertAfter(5, (doublyLinkedList.Find(3)))
	doublyLinkedList.InsertAfter(6, (doublyLinkedList.Find(3)))
	doublyLinkedList.InsertAfter(7, (doublyLinkedList.Find(3)))
	doublyLinkedList.Remove((doublyLinkedList.Find(3)))
	doublyLinkedList.Print()
	doublyLinkedList.MoveToFront((doublyLinkedList.Find(7)))
	doublyLinkedList.Print()
	doublyLinkedList.MoveToBack((doublyLinkedList.Find(1)))
	doublyLinkedList.Print()
	doublyLinkedList.MoveAfter((doublyLinkedList.Find(5)), (doublyLinkedList.Find(2)))
	doublyLinkedList.Print()
	doublyLinkedList.MoveBefore((doublyLinkedList.Find(5)), (doublyLinkedList.Find(2)))
	doublyLinkedList.Print()

}


