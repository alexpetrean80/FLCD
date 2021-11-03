package symtable

import "errors"

// Item represents an entity stored in the SymbolTable.
type Item struct {
	Key      string
	Value    string
	HashCode uint
	Next     *Item
}

// hashFunc computes the hashcode of a string and returns it.
// The hashcode is the sum of all characters' ASCII codes.
func hashFunc(val string) (hashCode uint) {
	hashCode = 1

	for _, ch := range val {
		hashCode += uint(ch)
	}

	return
}

// SymbolTable defines a hash table data structure
type SymbolTable struct {
	// items is a list of references to items in the hashtable
	items []*Item
	// size represents the total amount of items in the hashtable
	size uint
}

// New is a function for constructing the SymbolTable struct, abstracting away unneeded details about it.
func New() *SymbolTable {
	return &SymbolTable{
		items: make([]*Item, 11),
		size:  11,
	}
}

// Add inserts a new value v at the key k, solving the colisions by chaining.
func (st *SymbolTable) Add(k, v string) {
	i := st.getIndex(k)
	hc := hashFunc(k)

	head := st.items[i]

	for head != nil {
		if head.Key == k && head.HashCode == hc {
			head.Value = v
			return
		}
		head = head.Next
	}

	st.size++

	head = st.items[i]
	st.items[i] = &Item{
		Key:      k,
		Value:    v,
		HashCode: hc,
		Next:     head,
	}

	if float64(st.size) >= 0.7 {
		temp := st.items

		st.items = make([]*Item, 3*st.size)

		st.size = 3 * st.size

		for _, x := range temp {
			for x != nil {
				st.Add(x.Key, x.Value)
				x = x.Next
			}
		}
	}
}

// Remove deletes the item at key k with the correct hashcode, returning its value.
// If the key is not found, then the method returns an empty string and a non-nil error.
func (st *SymbolTable) Remove(k string) (string, error) {
	i := st.getIndex(k)
	hc := hashFunc(k)

	head := st.items[i]

	var prev *Item

	for head != nil {
		if head.Key == k && hc == head.HashCode {
			break
		}
		prev = head
		head = head.Next
	}

	if head == nil {
		return "", errors.New("not found")
	}

	st.size--

	if prev != nil {
		prev.Next = head.Next
	} else {
		st.items[i] = head.Next
	}
	return head.Value, nil
}

// Get retrieves the value at key k with the correct hashcode.
// If the key does not exist, then the method returns an empty string and a non-nil error.
func (st *SymbolTable) Get(k string) (string, error) {
	i := st.getIndex(k)
	hc := hashFunc(k)

	head := st.items[i]
	for head != nil {
		if head.Key == k && head.HashCode == hc {
			return head.Value, nil
		}
		head = head.Next
	}
	return "", errors.New("not found")
}

// getIndex retrieves the index at which the item with key k is found in the hashtable.
func (st *SymbolTable) getIndex(k string) (i uint) {
	hc := hashFunc(k)

	i = hc % st.size

	return
}
