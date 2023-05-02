package main

func scope() func() int {
	outer_var := 2
	foo := func() int {
		return outer_var
	}
	return foo
}

/*
// Closures, lexically scoped: Functions can access values that were in scope when defining the function
func scope() func() int{
outer_var := 2
foo := func() int { return outer_var}
return foo
}
func another_scope() func() int{
// won't compile because outer_var and foo not defined in this scope
outer_var = 999
return foo
}
// Closures: don't mutate outer vars, instead redefine them!
func outer() (func() int, int) {
outer_var := 2
inner := func() int {
outer_var += 99 // attempt to mutate outer_var from outer scope
return outer_var // => 101 (but outer_var is a newly redefined variable visible only inside inner)
}
return inner, outer_var // => 101, 2 (outer_var is still 2, not mutated by foo!)
}

*/
