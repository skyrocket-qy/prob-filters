# Cuckoo Filter

### Explanation

A Cuckoo filter is another probabilistic data structure for set membership testing, similar to a Bloom filter. It is based on Cuckoo Hashing and, most notably, supports dynamic addition and deletion of items. This makes it a powerful alternative to Bloom filters in scenarios where the set of items changes frequently.

### Scenario: Filtering for recently accessed articles

Consider a content delivery network (CDN) that caches recently accessed articles. To avoid a slow lookup in the main storage (e.g., a database or object storage), the CDN can use a Cuckoo filter to keep track of which articles are currently in its cache.

1.  When an article is added to the cache, its ID is added to the Cuckoo filter.
2.  When a request for an article comes in, the CDN first checks the filter. If the filter says the article is not in the cache, the request is forwarded to the main storage.
3.  If the filter indicates the article *might* be in the cache, the CDN attempts to retrieve it from the cache.
4.  Crucially, when an article is removed from the cache (e.g., due to an eviction policy), its ID is also removed from the Cuckoo filter. This keeps the filter accurate over time.

### Comparison

*   **Pros**:
    *   Supports dynamic addition and **deletion** of items.
    *   Often has higher space efficiency (fewer bits per item) than Bloom filters for low false positive rates (e.g., < 3%).
*   **Cons**:
    *   Insertions can fail if the filter is too full, requiring the filter to be rebuilt with more space.
    *   Implementation is slightly more complex than a standard Bloom filter.