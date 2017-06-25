Macro Preprocessor
==================
[Ken Leidal](ken@poshdevelopment.com)

A simple language agnostic string manipulation precompiler for defining macros for code generation.
The macros are similar in syntax to C macros.  For example, the macro SUM can
be defined as follows (see examples/sum3/macros.go):

```
//#define SUM (TYPE)
func SumTYPE(a TYPE, b TYPE) TYPE {
	return a + b
}

//#end
```

And executed on the source file (examples/sum3/source.go):

```
package main

//#import ./macros.go

//#macro SUM (int)

//#macro SUM (float64)
```

To produce the generated code (examples/sum3/target/go):

```
// .----------------------------------------.
// |      WARNING: AUTO-GENERATED CODE      |
// .----------------------------------------.
// This code was automatically generated using
// https://github.com/PoshDev/macro-preprocessor .
// The original source file was source.go.
// XXX: DO NOT CHANGE THIS FILE!!!
package main


func Sumint(a int, b int) int {
	return a + b
}


func Sumfloat64(a float64, b float64) float64 {
	return a + b
}

```

This makes eliminating boiler-plate code in a low-level strongly typed language like Go easy.

Note that since this is language agnostic, no AST's are used,
so matched macro argument names will be replaced, even in literals.

This is a simple script we are using for a few of our projects, but if you like it
and/or have ideas, feel free to file issues/PRs
