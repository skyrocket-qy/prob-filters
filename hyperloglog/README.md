# HyperLogLog

### Explanation

HyperLogLog is a probabilistic algorithm used for the count-distinct problem, which is the process of finding the number of unique elements in a dataset (also known as the cardinality). It can estimate the cardinality of very large datasets with a high degree of accuracy while using a very small and fixed amount of memory. It does not store the items themselves.

### Scenario: Counting unique visitors

A popular news website wants to count the number of unique visitors it receives each day. Storing every visitor's IP address or user ID in a set to count the unique entries would be memory-intensive, especially for a high-traffic site.

With HyperLogLog, the website can process a stream of visitor IDs and add each one to a HyperLogLog structure. At any time, the website can get a highly accurate estimate of the number of unique visitors by querying the structure, which might only be a few kilobytes in size. This is far more efficient than storing millions of unique identifiers.

### Comparison

*   **Pros**:
    *   Extremely space-efficient. It can estimate the cardinality of a set of billions of items with a typical error rate of around 2% using only 1.5 kB of memory.
    *   Fast, constant-time insertions.
    *   The union of two HyperLogLog structures can be computed, which allows for distributed or parallel counting.
*   **Cons**:
    *   The result is an approximation, not an exact count.
    *   It cannot retrieve the actual items that were added, only the estimated count of unique items.
    *   It does not support the deletion of items.