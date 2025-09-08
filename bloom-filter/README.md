# Bloom Filter

### Explanation

A Bloom filter is a space-efficient probabilistic data structure used to test whether an element is a member of a set. It consists of a bit array and multiple hash functions. False positive matches are possible, but false negatives are not. This means the filter can tell you if an element *might* be in the set or is *definitely not* in the set.

### Scenario: Checking for existing usernames

Imagine a service like Gmail or Twitter, where usernames must be unique. Every time a new user signs up, the service has to check if the desired username is already taken. This would require a database query for every attempt.

Instead, we can use a Bloom filter that stores all existing usernames. When a user tries to register a new username:
1. The username is checked against the Bloom filter.
2. If the filter returns "definitely not in the set," we know the username is available without needing to query the database.
3. If the filter returns "might be in the set" (a positive match), we then perform a database query to confirm. Since there's a chance of a false positive, the database is the source of truth.

This approach significantly reduces the load on the database, as most invalid username attempts are filtered out before a query is ever made.

### Comparison

*   **Pros**:
    *   Very space-efficient compared to storing all items in a hash set.
    *   Fast, constant-time insertions and lookups (O(k), where k is the number of hash functions).
    *   No false negatives.
*   **Cons**:
    *   False positives are possible. The rate can be tuned by adjusting the size of the filter and the number of hash functions.
    *   Cannot delete elements from a standard Bloom filter (though variations like Counting Bloom Filters exist).
    *   The size of the filter must be decided in advance based on the expected number of items.