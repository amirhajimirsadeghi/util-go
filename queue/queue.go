package queue

type Queue[E any] struct {
	Items []E
}

func (q *Queue[E]) Push(element E) {
	q.Items = append(q.Items, element)
}

func (q *Queue[E]) Pop() E {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	element := q.Items[0]
	q.Items = q.Items[1:]

	return element
}

func (q *Queue[E]) IsEmpty() bool {
	return len(q.Items) == 0
}