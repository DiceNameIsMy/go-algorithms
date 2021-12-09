package structures

type Hasher interface {
	Hash(key string) int
}

type SimpleHasher struct {
}

func (h *SimpleHasher) Hash(key string) int {
	return len(key) % 20
}

// Bucket to be stored in the hash table
type Bucket struct {
	data []string
}

func (b *Bucket) add(item string) {
	_, exists := b.exists(item)
	if !exists {
		b.data = append(b.data, item)
	}
}

func (b *Bucket) exists(item string) (int, bool) {
	for i := 0; i < len(b.data); i++ {
		if b.data[i] == item {
			return i, true
		}
	}
	return -1, false
}

func (b *Bucket) remove(item string) bool {
	idx, exists := b.exists(item)
	if !exists {
		return false
	}
	b.data[idx] = b.data[len(b.data)-1]
	b.data = b.data[:len(b.data)-1]
	return true
}

type HashTable struct {
	Hasher  Hasher
	Buckets [20]*Bucket
}

// add value to the hash table
func (h *HashTable) Add(value string) {
	idx := h.Hasher.Hash(value)
	if h.Buckets[idx] != nil {
		h.Buckets[idx].add(value)
	} else {
		h.Buckets[idx] = &Bucket{[]string{value}}
	}
}

// check if item exists in the hash table
func (h *HashTable) Exists(value string) bool {
	idx := h.Hasher.Hash(value)
	if h.Buckets[idx] == nil {
		return false
	} else {
		_, exists := h.Buckets[idx].exists(value)
		return exists
	}
}

// remove item from the hash table
// returns true if item was removed, false if value does not exist
func (h *HashTable) Remove(value string) bool {
	idx := h.Hasher.Hash(value)
	if h.Buckets[idx] == nil {
		return false
	} else {
		return h.Buckets[idx].remove(value)
	}
}
