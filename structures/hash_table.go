package structures

const TABLE_SIZE = 7

type HashTable struct {
	array [TABLE_SIZE]*LinkedList
}

func hash(key string) int {
	sum := 0

	for _, char := range key {
		sum += int(char)
	}

	return sum % TABLE_SIZE
}

func InitHashTable() *HashTable {
	ht := &HashTable{}

	for idx := range ht.array {
		ht.array[idx] = &LinkedList{}
	}

	return ht
}

func (ht *HashTable) Insert(key string) {
	bucket := ht.array[hash(key)]
	bucket.Insert(key)
}

func (ht *HashTable) Search(key string) bool {
	bucket := ht.array[hash(key)]
	return bucket.Search(key)
}

func (ht *HashTable) Delete(key string) {
	bucket := ht.array[hash(key)]
	bucket.DeleteWithValue(key)
}
