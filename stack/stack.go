package stack

type Stack[E any] struct {
	Items []E
}

func New[E any]() *Stack[E] {
	return &Stack[E]{Items: []E{}}
}

func (q *Stack[E]) Push(element E) {
	q.Items = append(q.Items, element)
}

func (q *Stack[E]) Pop() E {
	if q.IsEmpty() {
		panic("Stack is empty")
	}
	element := q.Items[len(q.Items)-1]
	q.Items = q.Items[:len(q.Items)-1]

	return element
}

func (q *Stack[E]) IsEmpty() bool {
	return len(q.Items) == 0
}
