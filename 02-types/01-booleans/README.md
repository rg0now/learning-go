# Logical operators

Solve the below exercises using boolean operands to represent logical values.

1. Write a function that returns the logical *inverse* of a logical value. Use the below signature:
   
   ```go
   func inverse(b bool) bool
   ```
   
2. Write a function that returns the logical "AND" of two logical values. Use the below signature:

   ```go
    func and(x, y bool) book
   ```

3. Write a function called `deMorgan` that takes two logical operands $A$ and $B$, returns $\overline{A+B}$. You cannot use the built-in logical operators of Go in this function, only the two functions implemented above. Use the below /De Morgan equation/ to obtain a form that you can express given the above two functions.

   $$\overline{A OR B} = \overline{A} AND \overline{B}$$

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.


