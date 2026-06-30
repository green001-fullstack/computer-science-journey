# What recursion means in the context of trees.

### Recursion allows us to apply the same algorithm to every node in a tree instead of writing a separate function for each node. Since every subtree is itself a tree, the same function naturally works for every subtree.

# Why the base case (node == nil) is necessary.

### The base case prevents us from trying to visit a node that does not exist. It also tells the recursive algorithm when to stop exploring a path.

# How the phrase Me → Left → Right maps directly to code.

### Me
#### Visit the current node by processing or printing its value.

### Left
#### Apply the preorder algorithm to the left subtree.

### Right
#### Apply the preorder algorithm to the right subtree.

# Why the same function works for the whole tree and every subtree.

### Because every subtree is a tree on its own

# My Biggest Insight

### I realized that recursion is not a magical concept. It simply means applying the same rule to every subtree because every subtree is itself a tree.