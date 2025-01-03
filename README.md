# Unintuitive Behavior of Uninitialized Integer Variables in Go

This example demonstrates a subtle but important aspect of Go's handling of uninitialized integer variables.  While the zero value for integers is well-defined, the behavior of their addresses can be surprising to those unfamiliar with the language's memory management.

**The Bug:**
The code compares an uninitialized integer to 0, as well as comparing the address of the integer to `nil`. While `i == 0` evaluates to true as expected, `&i == nil` evaluates to false even though the variable has not explicitly been assigned a value, illustrating that Go allocates memory even for uninitialized variables.

**The Solution:**
The solution shows the correct way to handle uninitialized integers. For clarity, it demonstrates how to explicitly initialize to the zero value or how to check the specific value of the integer.