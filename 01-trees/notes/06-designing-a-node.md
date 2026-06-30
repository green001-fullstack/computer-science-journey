# Why each field exists in the Node struct.

### The value is to store the data of the node. The left and right field stores the reference to their children to show their relationship

# Why Left and Right are pointers.

### To copy the reference to their children, thus representing their relationship

# Why Value is not a pointer.

### It is the data specific to a node and not a reference

# What happens if a node has no children.

### The left and right will point to nil
