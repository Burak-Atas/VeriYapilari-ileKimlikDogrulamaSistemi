package linkedlist

type Node struct {
	ID       string
	Name     string
	Surname  string
	Email    string
	Password string
	next     *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertLast(ıd, name, surname, email, password string) {
	data := &Node{
		ID:       ıd,
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
	}
	if list.head == nil {
		list.head = data
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

func (list *LinkedList) RemoveByValue(i string) bool {
	if list.head == nil {
		return false
	}
	if list.head.ID == i {
		list.head = list.head.next
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.ID == i {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) GetFirst() (string, bool) {
	if list.head == nil {
		return "", false
	}
	return list.head.ID, true
}

func (list *LinkedList) GetLast() (string, bool) {
	if list.head == nil {
		return "", false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.ID, true
}

func (list *LinkedList) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}
	return count
}

func (list *LinkedList) GetItems() []string {
	var items []string
	current := list.head
	for current != nil {
		items = append(items, current.ID)
		current = current.next
	}
	return items
}

type Get interface {
	GetItems() []int
}
