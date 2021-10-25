package main

import "errors"

type Item struct {
	Key      string
	Value    string
	HashCode uint
	Next     *Item
}

func hashFunc(val string) (hashCode uint) {
	hashCode = 0

	for _, ch := range val {
		hashCode += uint(ch)
	}

	return
}

type HashTable struct {
	items []*Item
	size  uint
}

func New() *HashTable {
	return &HashTable{
		items: make([]*Item, 10),
		size:  10,
	}
}

func (ht *HashTable) Add(k, v string) {
	i := ht.getIndex(k)
	hc := hashFunc(k)

	head := ht.items[i]

	for head != nil {
		if head.Key == k && head.HashCode == hc {
			head.Value = v
			return
		}
		head = head.Next
	}

	ht.size++

	head = ht.items[i]
	ht.items[i] = &Item{
		Key:      k,
		Value:    v,
		HashCode: hc,
		Next:     head,
	}

	if float64(ht.size) >= 0.7 {
		temp := ht.items

		ht.items = make([]*Item, 2*ht.size)

		ht.size = 2 * ht.size

		for _, x := range temp {
			for x != nil {
				ht.Add(x.Key, x.Value)
				x = x.Next
			}
		}
	}
}

func (ht *HashTable) Remove(k string) (string, error) {
	i := ht.getIndex(k)
	hc := hashFunc(k)

	head := ht.items[i]

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

	ht.size--

	if prev != nil {
		prev.Next = head.Next
	} else {
		ht.items[i] = head.Next
	}
	return head.Value, nil
}

func (ht *HashTable) Get(k string) (string, error) {
	i := ht.getIndex(k)
	hc := hashFunc(k)

	head := ht.items[i]
	for head != nil {
		if head.Key == k && head.HashCode == hc {
			return head.Value, nil
		}
		head = head.Next
	}
	return "", errors.New("not found")
}

func (ht *HashTable) getIndex(k string) (i uint) {
	hc := hashFunc(k)

	i = hc % ht.size

	return

}
