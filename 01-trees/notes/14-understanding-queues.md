# What a queue is and the problem it solves

### It is a data structure that operate on the principle of first in first out. It process data that enter the queue first before the later ones. A queue solves the problem of processing items in the exact order they were discovered or added.

# FIFO

### FIFO (First In, First Out) means the first element added to the queue is the first element removed.

# Enqueue.

### Adding to the back of a queue

# Dequeue

### Remove from the beginning of the queue(the first member of the queue)

# Why are we storing *Node in the queue instead of Node

### The queue stores pointers because a tree is built from linked nodes. By storing pointers, the queue processes the actual nodes in the tree rather than making copies of them.

# Why BFS naturally uses a queue

### Because BFS wants to process every node at a level before going deeper into other levels. Queue enables it to achieve this

# Why the queue stores pointers to nodes.

### Queues need to modify the data in memory and not a duplicate

# My Biggest Insight

### I realized that algorithms naturally pair with data structures. BFS uses a queue because a queue processes nodes in the exact order they are discovered.