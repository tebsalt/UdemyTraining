package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/* closure helps us limit the scope of va
riables used by multiple functions.
without closure, for 2 or more funcs to have access to the same variable, that variable would need to be package scope.

Anonymous function
a function without a name ... func(), no name ahead of it, like func main () for example
to put a function inside a function, it has to be done this way.
It's called a func expression,
when you assign a function inside a variable
*/
