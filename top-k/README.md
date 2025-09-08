# Top-K

### Explanation

Top-K is a probabilistic data structure that allows you to find the most frequent items in a data stream. It is a generalization of the "heavy hitters" problem. It keeps a small, fixed-size list of the most frequent items seen so far.

### Scenario: Real-time "Top 10" list

A video streaming service wants to show a real-time list of the "Top 10 most-watched videos right now". Tracking the exact watch count for millions of videos in real-time is challenging.

Instead, the service can use a Top-K algorithm. As users watch videos, the video IDs are fed into the Top-K structure. At any time, the service can query the structure to get a list of the top 10 most frequent video IDs, which represents the most-watched videos.

### Comparison

*   **Pros**:
    *   Very space-efficient, as it only stores a small number of items (K).
    *   Can provide real-time estimates of the most frequent items.
*   **Cons**:
    *   It is an approximation. It may not always find the true top K items, especially if the frequencies are close.
    *   The accuracy depends on the size of the structure and the algorithm used (e.g., "heavy hitters" with Count-Min Sketch).