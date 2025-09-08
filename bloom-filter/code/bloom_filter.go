package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter represents a probabilistic data structure for set membership testing.
type BloomFilter struct {
	bits       []bool
	hashFuncs  []hash.Hash64
	numHashes  int
	capacity   int
	falsePosRate float64
}

// New creates a new BloomFilter with a given capacity and desired false positive rate.
func New(capacity int, falsePosRate float64) *BloomFilter {
	// Optimal number of bits (m) and hash functions (k)
	// m = -(n * ln(p)) / (ln(2)^2)
	// k = (m / n) * ln(2)
	m := int(-float64(capacity) * (math.Log(falsePosRate) / (math.Log(2) * math.Log(2))))
	k := int(float64(m) / float64(capacity) * math.Log(2))

	if m < 1 {
		m = 1
	}
	if k < 1 {
		k = 1
	}

	hashFuncs := make([]hash.Hash64, k)
	for i := 0; i < k; i++ {
		hashFuncs[i] = fnv.New64a() // Using FNV-1a for simplicity
	}

	return &BloomFilter{
		bits:       make([]bool, m),
		hashFuncs:  hashFuncs,
		numHashes:  k,
		capacity:   capacity,
		falsePosRate: falsePosRate,
	}
}

// Add inserts an element into the Bloom filter.
func (bf *BloomFilter) Add(data []byte) {
	for _, hf := range bf.hashFuncs {
		hf.Reset()
		hf.Write(data)
		index := hf.Sum64() % uint64(len(bf.bits))
		bf.bits[index] = true
	}
}

// Contains checks if an element might be in the Bloom filter.
// Returns true if the element might be present, false if definitely not.
func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hf := range bf.hashFuncs {
		hf.Reset()
		hf.Write(data)
		index := hf.Sum64() % uint64(len(bf.bits))
		if !bf.bits[index] {
			return false // Definitely not in the set
		}
	}
	return true // Might be in the set (possible false positive)
}

// Example usage (main function for demonstration)
/*
import (
	"fmt"
	"math"
)

func main() {
	filter := New(1000, 0.01) // Capacity for 1000 items, 1% false positive rate

	itemsToAdd := []string{"apple", "banana", "cherry", "date", "elderberry"}
	for _, item := range itemsToAdd {
		filter.Add([]byte(item))
		fmt.Printf("Added: %s\n", item)
	}

	fmt.Println("\nChecking items:")
	checkItems := []string{"apple", "grape", "cherry", "kiwi", "date"}
	for _, item := range checkItems {
		if filter.Contains([]byte(item)) {
			fmt.Printf("'%s' might be in the set.\n", item)
		} else {
			fmt.Printf("'%s' is definitely NOT in the set.\n", item)
		}
	}

	// Demonstrate a false positive (if it occurs)
	fmt.Println("\nDemonstrating potential false positive:")
	falsePositiveCandidate := "zucchini" // A word not added
	if filter.Contains([]byte(falsePositiveCandidate)) {
		fmt.Printf("'%s' might be in the set (false positive).\n", falsePositiveCandidate)
	} else {
		fmt.Printf("'%s' is definitely NOT in the set.\n", falsePositiveCandidate)
	}
}
*/
