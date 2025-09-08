# Count-Min Sketch

### Explanation

A Count-Min Sketch is a probabilistic data structure that serves as a frequency table of events in a stream of data. It can be used to estimate the frequency of an item. It is similar to a counting Bloom filter, but it's designed to provide frequency estimates rather than just membership. A key characteristic is that it may overestimate frequencies, but it never underestimates them.

### Scenario: Identifying trending topics

A social media platform wants to identify trending topics or hashtags in real-time. Tracking the exact count for every hashtag would require a massive amount of memory. Instead, the platform can use a Count-Min Sketch.

1.  As new posts with hashtags are created, the hashtags are added to the sketch, which increments their estimated frequency.
2.  To find trending topics, the platform can query the sketch for the estimated frequency of various hashtags.
3.  This allows the platform to identify which hashtags are being used most frequently and are likely trending, without storing exact counts for every single hashtag.

### Comparison

*   **Pros**:
    *   Very space-efficient for estimating frequencies in a large stream of data.
    *   Fast, constant-time updates and queries.
    *   Can be used to solve a range of frequency-related problems, such as finding the most frequent items ("heavy hitters").
*   **Cons**:
    *   It always provides an overestimation of the true frequency, never an underestimation. The amount of error can be reduced by increasing the size of the sketch.
    *   It cannot delete items (though variations that support this exist).
    *   It does not store the items themselves, only their estimated frequencies.