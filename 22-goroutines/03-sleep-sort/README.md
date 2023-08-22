# Sleep sort

Sleep sort works by starting a separate goroutine for each item to be sorted. Each task sleeps for
an interval corresponding to the item's sort key, and then emits the item. Items are then collected
sequentially by the main thread.

Write a `sleepSort` function that performs a forward sleep-sort on the slice
of unsigned integers (`uint`) received as input. 

You can assume the input is in the range [1,50] and every item `x` should induce a wait of `x*10`
milliseconds if you sort in the forward direction, and `500 - x*10` milliseconds if you sort in the
reverse direction. This way, the function is guaranteed to terminate under half a second. 

Implement your code in a function `sleepSort` with the below signature:

``` go
// sleepSort returns the input uint-slice sorted in the forward order.
func sleepSort(input []uint) []uint {
    ...
}
```

Insert the code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.

HINTS:
- use a separate goroutine per item and a channel to return the result from each goroutine,
- `for` loop variables and closures may interact strangely in Go, remember the GOTCHA from the
  seminar on loop constructs!
