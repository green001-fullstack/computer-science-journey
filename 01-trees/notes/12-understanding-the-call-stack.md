# What a call stack is.

### Call stack is a phenomenon in computer programming that describes where certain operations such as function calls are paused and stored for the purpose of resuming them later in an ordered manner. The last operation to enter the stack comes out first. 

# What a stack frame is.

### Represents each layer of paused function in a call stack

# Why recursion works.

### Recursion works because the program does not forget the paused functions and enables them to be completed by saving the operations in a call stack.

# Why recursion can overflow.

### If the height of the tree is too much

# Why the space complexity depends on tree height.

### At every level, the left child returns before the right child enters, so the right child replaces that space

# My Biggest Insight.

### This was the discovery of the reason why the space complexity of recursive function depends on the height of the tree. 