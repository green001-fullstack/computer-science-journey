# What inorder traversal is

### Traverses the left subtrees first, then the parent node and then the right subtrees

# The phrase Left → Me → Right.

### Left - traverse the left subtree first
### Me - Process the parent node
### Right - Traverse the right subtree

# Why only one line changed from preorder.

### The recursive structure of the algorithm remains the same. The only difference is the position of the line that processes the current node. Moving this line changes the traversal order.

# Why the algorithm still works for every subtree.

### Because every subtree is a tree on its own
