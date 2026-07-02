# What Level-Order Traversal is

### Level-order traversal is a tree traversal algorithm that visits nodes one level at a time, from left to right (or according to the enqueue order), before moving to the next level.

# The BFS algorithm in your own words

### 1. Put the root into a queue.
### 2. While the queue is not empty:
###     - Remove the front node.
###     - Process it.
###     - Add its children to the queue if they exist.
### 3. Repeat until the queue is empty.

# Why a queue makes the algorithm possible

### Queue process the earliest data that enter it before processing the latest ones

# Time complexity

### O(n) because every node is visited once

# Space complexity

### O(w). It depends on how wide the tree is at a particular level.

# How BFS differs from DFS

### DFS process nodes along a path, it goes deeper before coming back to process the nodes along another path. BFS process all the nodes at a level before going deeper

# My Biggest Insight.

### I realized that BFS is actually a very simple algorithm once the queue is in place. The queue is what keeps track of the nodes waiting to be processed, allowing the algorithm to naturally move level by level.