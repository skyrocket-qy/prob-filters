# t-digest

### Explanation

A t-digest is a probabilistic data structure for estimating the rank of an element in a sorted sequence of numbers, and for estimating percentiles (e.g., the 99th percentile). It is particularly useful for summarizing the distribution of a large dataset or a stream of data in a space-efficient way.

### Scenario: Monitoring API response times

An e-commerce platform wants to monitor the response times of its API endpoints. It needs to calculate the 95th and 99th percentile response times to ensure that most users are having a good experience. Storing all response times to calculate these percentiles would be very memory-intensive.

Instead, the platform can use a t-digest. It adds each response time to the t-digest, which maintains a compact summary of the distribution. The platform can then query the t-digest to get accurate estimates of the 95th and 99th percentile response times.

### Comparison

*   **Pros**:
    *   Very space-efficient for estimating percentiles.
    *   Can be merged, allowing for distributed calculations.
    *   More accurate at the tails of the distribution (e.g., p99, p99.9) than some other methods.
*   **Cons**:
    *   It is an approximation, not an exact calculation.
    *   The accuracy depends on the size of the digest.