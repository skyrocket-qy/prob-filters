# Counting Bloom Filter

### Explanation

A Counting Bloom filter is an extension of the standard Bloom filter that supports the deletion of elements. Instead of using a single bit for each slot in the filter, it uses a small counter (e.g., 4 bits). When an item is added, the counters at its corresponding hash locations are incremented. When an item is deleted, the counters are decremented. A query checks if all corresponding counters are non-zero.

### Scenario: Managing a list of malicious URLs

Consider a web browser's security feature that blocks access to malicious URLs. A standard Bloom filter could be used to store the list of malicious sites, but it would be difficult to remove URLs that are no longer considered a threat.

A Counting Bloom filter is a better fit here.
1.  When a URL is identified as malicious, it is added to the filter by incrementing the relevant counters.
2.  If a URL is later deemed safe, it can be removed from the filter by decrementing the counters.
3.  This allows the filter to stay up-to-date with the latest threat intelligence without needing to be rebuilt from scratch.

### Comparison

*   **Pros**:
    *   Supports the deletion of elements, which is a major advantage over standard Bloom filters.
    *   Still relatively space-efficient, though it requires more space than a standard Bloom filter.
*   **Cons**:
    *   Requires more space (e.g., 4x or more) than a standard Bloom filter.
    *   The counters can overflow if an item is added too many times, which can be an issue if the same item is added repeatedly.
    *   The size of the counters limits how many times an element can be added before overflow occurs.

### Mathematical Foundations

A Counting Bloom Filter extends the standard Bloom Filter by replacing single bits with counters. When an element is added, the corresponding counters are incremented. When an element is removed, the counters are decremented. An element is considered present if all its corresponding counters are greater than zero.

The false positive rate calculation is similar to a standard Bloom filter, but the counter size introduces a new consideration: counter overflow. If a counter reaches its maximum value (e.g., 255 for an 8-bit counter) and an item is added again, it can lead to issues if that item is later removed, as the counter cannot be decremented further. The probability of counter overflow depends on the counter size and the number of items added.

### Implementation Considerations

*   **Counter Size**: The choice of counter size (e.g., 4-bit, 8-bit) is a trade-off between memory usage and the maximum number of times an element can be added before a counter overflows. Larger counters reduce the risk of overflow but increase memory consumption.
*   **Deletion Semantics**: Deletion in a Counting Bloom Filter is "soft." It decrements counters, but if a counter was incremented by a false positive, decrementing it might incorrectly remove a truly present item if its counter reaches zero. This means that while deletions are supported, they can introduce new types of errors if not handled carefully.
*   **Hash Functions**: Similar to standard Bloom filters, the quality and independence of hash functions are crucial for distributing elements evenly and minimizing collisions.
*   **Memory Overhead**: Counting Bloom filters require more memory than standard Bloom filters because each bit is replaced by a counter. For example, an 8-bit counter uses 8 times more memory than a single bit.

## Code Example

A basic Go implementation of the Counting Bloom Filter can be found [here](code/counting_bloom_filter.go).