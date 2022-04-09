package linkedlist

func (list *LinkedList) SearchEmail(email string) int {
	if list.head == nil {
		return 0
	}
	current := list.head
	for current != nil {
		if current.Email == email {
			return 1
		}
		current = current.next
	}
	return 0
}
