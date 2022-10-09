package queue

import "testing"

func TestNew(t *testing.T) {
	capacity := 3
	q := New(capacity)

	if q.capacity != 3 {
		t.Errorf("Incorrect capacity: %v, want %v", q.capacity, capacity)
	}
}

func TestQueue_Append(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Failed to add an item to the queue")
		}
	}

	if q.Append(4) {
		t.Errorf("Unexpectedly was able to add an item to the queue")
	}
}

func TestQueue_Next(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Expected another item from the queue, missing item %v", i)
		}

		if item != i {
			t.Errorf("Got item in the wrong order: %v, want %v", item, i)
		}
	}

	_, ok := q.Next()
	if ok {
		t.Errorf("Expected queue to be empty.")
	}
}
