package main

/*
Personal notes.

Go is a general purpose, compiled, statically typed, object-oriented
(in its own way) language with built-in testing support.

package and file names should be good: short, concise, evocative.
By convention, packages are given lower case, single-word names;
there should be no need for underscores or mixedCaps.

the package in src/encoding/base64 is imported as "encoding/base64"
but has the name base64, not encoding_base64 and not encodingBase64.
*/

func main() {
	helloWorld()
	types()
	inference()
	normalSyntaxInGo()
	math()
}