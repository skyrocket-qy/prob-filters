# Probabilistic Models for Pre-Database Filtering

This repository provides an overview and practical scenarios for various probabilistic data structures, useful for optimizing database access, filtering, and cardinality estimation in software engineering.

Each model is detailed in its own dedicated folder, providing explanations, use-case scenarios, and comparisons.

## Choosing the Right Probabilistic Model

To help you select the most suitable probabilistic model for your needs, consider the following categories and their primary applications:

### 1. Membership Testing (Is an item in a set?)

These models are ideal when you need to quickly check if an element is part of a collection, often to avoid more expensive lookups (e.g., database queries).

*   **[Bloom Filter](bloom-filter/README.md)**: For basic membership testing where false positives are acceptable and deletions are not required. Highly space-efficient.
*   **[Counting Bloom Filter](counting-bloom-filter/README.md)**: Similar to a Bloom Filter, but supports the deletion of items. Uses more space than a standard Bloom Filter.
*   **[Cuckoo Filter](cuckoo-filter/README.md)**: Supports dynamic additions and deletions, and can be more space-efficient than Bloom Filters for low false positive rates.
*   **[Quotient Filter](quotient-filter/README.md)**: Often more space-efficient than Bloom and Cuckoo filters, and supports merging and resizing.
*   **[XOR Filter](xor-filter/README.md)**: Extremely fast and space-efficient for static sets. No false positives for items in the set, but cannot be modified after construction.

### 2. Cardinality Estimation (How many unique items are there?)

These models are used to estimate the number of unique elements in a large dataset or stream, using minimal memory.

*   **[HyperLogLog](hyperloglog/README.md)**: Estimates the number of unique items (cardinality) in very large datasets with high accuracy and extremely low memory usage.

### 3. Frequency Estimation & Top Items (How often does an item appear? What are the most frequent items?)

These models help in estimating the frequency of items in a stream or identifying the most frequent items.

*   **[Count-Min Sketch](count-min-sketch/README.md)**: Estimates the frequency of items in a data stream. May overestimate frequencies but never underestimates.
*   **[Top-K](top-k/README.md)**: Identifies the most frequent items (top K) in a data stream.

### 4. Quantile & Percentile Estimation (What is the value at a certain percentile?)

These models are used to estimate quantiles and percentiles from a dataset or stream, providing insights into data distribution.

*   **[t-digest](t-digest/README.md)**: Estimates percentiles (e.g., 95th, 99th) from large datasets or streams, particularly accurate at the tails of the distribution.

### 5. Sorted Data Structures (Maintaining sorted order with probabilistic guarantees)

While not strictly a "filter" in the same sense as others, this data structure uses probabilistic methods to achieve efficient sorted operations.

*   **[Skip List](skip-list/README.md)**: A probabilistic data structure for maintaining sorted data, offering performance similar to balanced trees with simpler implementation, especially good for concurrent access.

---

Navigate to each model's directory for detailed information, scenarios, and comparisons.