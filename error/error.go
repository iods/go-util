package error

// conventionally an error is something that is fatal to a program - go does not provide this
// in Go, an error is just a value that a function can return if something unexpected happens

/*
 Scenario: Something happens unexpected in a function.
 Example:  Function is invoked with the wrong arguments.
           Execution within the function did not go as planned.
 Response: Error is returned as a value if either happen.

*/

/*
 Errors in Go are not exceptions - coming from PHP
 Errors in Go are values - not like try and catch

pcs = Program Counters - up to 10? what would the Go limit or threshold be?

*/

// error handling is addressing something that happened unexpected, and performing a task in response to it

// goal one - how do we raise errors in Go
// goal two - how do we handle those errors in Go

// this is usually done with try and catch construct, my previous PHP experience

// simple and straight forward but the simplicity makes it kind of difficult to understand

// error handling is really about the decoration of debugging and managing output?
// that means it is better to understand the stacktrace first

/*

 Errors as Values

 - capturing errors as variables using Gos built-in `error` type
   - essentially creating error variables, like other types (like int and string)


 - when no errors exist, `error` holds the value of `nil`
    - if there is anything else in `error` then there is an error - "that is a value"

	- the try catch, turns into an if conditional

 - in other languages, errors are expected. they are called exceptions.
 - uncertain if it throws an exception, so you build around that, handle the uncertainty
 -

*/

// Go runtime package provides the information for providing stack traces

// returning an error signals to the caller that there was an error.
