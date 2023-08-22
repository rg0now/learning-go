# Calculator

Implement the below functions. Use [the `math.Round` function](https://pkg.go.dev/math#Round) to round floats to the nearest integer, [`math.Sqrt`](https://pkg.go.dev/math#Sqrt) to compute square roots, and the `int()` function to convert a float to an integer.

1. Write a function with the below signature that receives two `float64` arguments `x` and `y` and returns the arithmetic mean of `x` and `y` rounded to the nearest integer (so the result for `x=1.2` and `y=2.4` would be 2):

   ```go
   func amean(x,y float64) int
   ```

2. Write a function with the below signature that receives `x` and `y` as *strings* and returns the same output as before or an arbitrary error if any of the input strings could not be parsed to a float. Use the [`strconv.ParseFloat` function](https://pkg.go.dev/strconv#ParseFloat) to parse a `string` to a `float64`. 
   
   ```go
   func ameanString(x,y string) (int, error)
   ```

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.
