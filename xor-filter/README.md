# XOR Filter

### Explanation

An XOR filter is a type of probabilistic data structure used for set membership testing. They are faster and more memory-efficient than Bloom and Cuckoo filters. XOR filters are a type of perfect hash function, meaning they have no false positives for items in the set they were built with. However, they are static; once constructed, items cannot be added or removed.

### Scenario: Serving static assets from a CDN

A Content Delivery Network (CDN) needs to determine quickly if an asset (e.g., an image or a JavaScript file) is stored at an edge location. The set of assets is large but changes infrequently.

An XOR filter can be built containing the fingerprints of all assets. When a request for an asset comes in, the CDN can use the XOR filter to very quickly check if the asset is supposed to be in the cache. Because XOR filters are extremely fast and small, this check is highly efficient. When the set of assets is updated, a new filter must be constructed and distributed to the edge locations.

### Comparison

*   **Pros**:
    *   Uses less space than Bloom or Cuckoo filters (e.g., ~1.23 bits per item).
    *   Faster lookups than Bloom or Cuckoo filters.
    *   No false positives for items that are in the set.
*   **Cons**:
    *   It is a static data structure; it cannot be modified after it is built. Adding or removing items requires a complete rebuild of the filter.
    *   The construction process is more complex than for a Bloom filter.