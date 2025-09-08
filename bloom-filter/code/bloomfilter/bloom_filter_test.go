package bloomfilter

import (
	"fmt"
	"math"
	"testing"
)

func TestNew(t *testing.T) {
	capacity := 1000
	falsePosRate := 0.01
	bf := New(capacity, falsePosRate)

	if bf.capacity != capacity {
		t.Errorf("Expected capacity %d, but got %d", capacity, bf.capacity)
	}

	if bf.falsePosRate != falsePosRate {
		t.Errorf("Expected false positive rate %f, but got %f", falsePosRate, bf.falsePosRate)
	}

	// Check if m and k are calculated correctly using ceiling
	mFloat := -float64(capacity) * math.Log(falsePosRate) / (math.Ln2 * math.Ln2)
	kFloat := mFloat / float64(capacity) * math.Ln2
	expectedM := int(math.Ceil(mFloat))
	expectedK := int(math.Ceil(kFloat))

	if int(bf.size) != expectedM {
		t.Errorf("Expected bits length %d, but got %d", expectedM, bf.size)
	}

	if bf.numHashes != expectedK {
		t.Errorf("Expected numHashes %d, but got %d", expectedK, bf.numHashes)
	}
}

func TestAddAndContains(t *testing.T) {
	bf := New(100, 0.01)
	item := []byte("hello")
	bf.Add(item)

	if !bf.Contains(item) {
		t.Errorf("Expected to find item 'hello', but it was not found")
	}

	nonExistentItem := []byte("world")
	if bf.Contains(nonExistentItem) {
		t.Logf("False positive for 'world', which is acceptable but should be noted.")
	}
}

func TestFalsePositiveRate(t *testing.T) {
	capacity := 10000
	falsePosRate := 0.01
	bf := New(capacity, falsePosRate)

	// Add items to the filter
	for i := 0; i < capacity; i++ {
		item := []byte(fmt.Sprintf("item-%d", i))
		bf.Add(item)
	}

	// Check for false positives
	falsePositives := 0
	numChecks := 10000
	for i := 0; i < numChecks; i++ {
		// Check for items that were not added to the filter
		item := []byte(fmt.Sprintf("not-in-filter-%d", i))
		if bf.Contains(item) {
			falsePositives++
		}
	}

	observedFPR := float64(falsePositives) / float64(numChecks)
	t.Logf("Observed false positive rate: %f", observedFPR)

	// We expect the observed FPR to be close to the configured one.
	// Allow a small margin for error due to probabilistic nature.
	if observedFPR > falsePosRate*1.5 {
		t.Errorf("Observed false positive rate %f is higher than the acceptable margin of %f", observedFPR, falsePosRate*1.5)
	}
}

func TestZeroCapacity(t *testing.T) {
	bf := New(0, 0.1)
	if bf.capacity != 0 {
		t.Error("should handle zero capacity")
	}
	// m and k should be at least 1
	if bf.size == 0 || bf.numHashes == 0 {
		t.Error("m and k should be at least 1, even for zero capacity")
	}
}
