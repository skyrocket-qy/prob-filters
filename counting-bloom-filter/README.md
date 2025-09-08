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