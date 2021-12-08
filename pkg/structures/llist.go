package structures

type nodeStr struct {
	data string
	next *nodeStr
}

// LinkedList
type LListStr struct {
	head   *nodeStr
	length int
}

// Adds element to the head of the list
func (l *LListStr) Prepend(data string) {
	newNode := &nodeStr{data, l.head}
	l.head = newNode
	l.length++
}

// Removes head element and returns it's data
func (l *LListStr) Poll() (string, bool) {
	if l.IsEmpty() {
		return "", false
	}
	data := l.head.data
	l.head = l.head.next
	l.length--

	return data, true
}

func (l *LListStr) IsEmpty() bool {
	return l.length == 0
}
