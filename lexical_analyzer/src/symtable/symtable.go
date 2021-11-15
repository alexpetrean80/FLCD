package symtable

import (
	"fmt"
	"github.com/alexpetrean80/FLCD/utils"
	"log"
	"os"
)

type SymbolTable interface {
	Add(tkn string)
	utils.Outputter
	fmt.Stringer
}

// symTable defines a hash table data structure
type symTable struct {
	// items is a list of references to items in the hashtable
	items []*item
	// size represents the total amount of items in the hashtable
}

// New is a function for constructing the symTable struct, abstracting away unneeded details about it.
func New() *symTable {
	return &symTable{
		items: []*item{},
	}
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

// Add inserts a new value v at the val k, solving the collisions by chaining.
func (st *symTable) Add(val string) {
	if st.checkIfExists(val){
		return
	}
	i := st.getIndex(val)
	hc := hashFunc(val)

	var head *item
	if len(st.items) != 0 {
		head = st.items[i]
	}

	sym := &item{
		val:      val,
		hashCode: hc,
		next:     head,
	}
	st.items = append(st.items, sym)
	head = st.items[i]
}

func (st symTable) checkIfExists(val string) bool {
	for _, sym := range st.items {
		if sym.val == val {
			return true
		}
		head := sym
		for head != nil {
			if head.val == val {
				return true
			}
			head = head.next
		}
	}
	return false
}

// getIndex retrieves the index at which the item with val k is found in the hashtable.
func (st *symTable) getIndex(k string) (i uint) {
	hc := hashFunc(k)

	if len(st.items) == 0 {
		return 0
	}
	i = hc % uint(len(st.items))

	return
}

func (st symTable) String() string {
	s := ""
	for _, item := range st.items {
		if item == nil {
			s += "nil\n"
			continue
		}
		s += item.String()
		s += "\n"
	}
	return s
}

func (st symTable) Output(outFile string) {
	if err := os.WriteFile(outFile, []byte(st.String()), 0666); err != nil {
		log.Fatalf("cannot write pif to file: %s", err.Error())
	}
}

// item represents an entity stored in the symTable.
type item struct {
	val      string
	hashCode uint
	next     *item
}

func (i item) String() string {
	return fmt.Sprintf("{token: %s}", i.val)
}
