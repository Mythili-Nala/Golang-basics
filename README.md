# Golang-basics

A neat explaination of Golang Basics 

GOLANG
------

	- Go is a procedural programming language, launched by Google as an open source programming language.
	- It also supports enviroment adopting patterns alike other dynamic languages.
	
  Golang has been designed in such a way that instead of using threading it uses Goroutine, which is similar to threading but consumes very less memory. Like threading consumes 1MB whereas Goroutine consumes 2KB of memory, hence at the same time, we can have millions of goroutine triggered. So the above-discussed point makes golang a strong language that handles concurrency like C++ and Java.


	Advantages:

	Flexible- It is concise, simple and easy to read.
	Concurrency- It allows multiple process running simultaneously and effectively.
	Quick Outcome- Its compilation time is very fast.
	Library- It provide a rich standard library.
	Garbage collection- It is a key feature of go. Go excels in giving a lot of control over memory allocation and has dramatically reduced latency in the most recent versions of the garbage collector.

	Disadvantages:

	It has no support for generics.
	The packages distributed with this programming language is quite useful but Go is not so object-oriented in the conventional sense.
	There is absence of some libraries especially a UI tool kit.

	Features of golang:- 
	
	- Language design
	- Package Management
	- Powerful library
	- Testing support
	- platform Independent 
  
Identifiers 
============
	- There is no limit for identifier length
	- 25 keywords are present
	- Complex numbers can also be defined with the help of complex_ datatype
	- Strings are immutable in go
	- rune represents int32
	- In Golang, we can use any symbol or character in the writing system in the world because it uses the UTF-8 encoding

Operators in Golang
=====================
	- -, +, !, &, *, <-, and ^ are also known as unary operators and the precedence of unary operators is higher. ++ and — operators are from statements they are not expressions,
		 so they are out from the operator hierarchy.

	- Golang has 6 bitwise operators, in which AND NOT (&^) is new, which is a bit clear operator

	- Misc operator 
		*, &, <-(This is a receive operator, used to receive value from channel)


==================================

 Golang doesn't support implicit type conversion, we need to do it explicitly.
 There is no concept of type casting, but exists type conversion


Short declaration operator (:-)
=============================

	- It is used to declare and initialize the variables only inside the functions.
	- There is no need to put type. If you, it will give error.


Deadlock and Default Case in Select Statement
=======================
	- In Go language, the select statement is just like a switch statement, but in the select statement, the case statement refers to communication, i.e. sent or receive operation on the channel.
	 
     select{

     case SendOrReceive1: // Statement
     case SendOrReceive2: // Statement
     case SendOrReceive3: // Statement
     .......
     default: // Statement

Deadlock: When you trying to read or write data from the channel but the channel does not have value. So, it blocks the current execution of the goroutine and passes the control to other goroutines, but if there is no other goroutine available or other goroutines are sleeping due to this situation program will crash. This phenomenon is known as deadlock.


=============================

	- In golang main() doesnt take any argument and doesn't return anything (unlike other static typed languages)
	- main() is present in main package only
	- init() is same as main(), doesn't take any argument and doesn't return anything
	- each package has a init(), called when package is initialised
	- we can create multiple init() , they will be Executed in a alphabetical order, this function is declared implicitly, so you cannot reference it from anywhere 
	- init() will always be executed before main(), so doesnt depend on main()
	- The main purpose of the init() function is to initialize the global variables that cannot be initialized in the global context.

	***- in golang we need not to store the values returned by a function.

	- Blank Identifier (_), you can use multiple _ in your program, so we can say Golang program can have multiple variables using the same identifier name which is the blank identifier.
	- In golang methods are similar to the functions except one difference, i.e, the method contains a receiver argument in it. 
	With the help of the receiver argument, the method can access the properties of the receiver. Here, the receiver can be of struct type or non-struct type. 
	When you create a method in your code the receiver and receiver type must present in the same package. 
	And you are not allowed to create a method in which the receiver type is already defined in another package including inbuilt type like int, string, etc. 
	If you try to do so, then the compiler will give an error.


Array and slices
=======
	- In golang arrays are of fixed length but slices arent.
	- In Go language, an array is of value type not of reference type. So when the array is assigned to a new variable, then the changes made in the new variable do not affect the original array
	- If we try to copy an array with the address operator (&), then any changes made in the new array will be showcased in the actual array.


	- In Go language slice is more powerful, flexible, convenient than an array, and is a lightweight data structure.
	- It is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array
	- Internally, slice and an array are connected with each other, a slice is a reference to an underlying array. 
  
	A slice contains three components:
		pointer, length, capacity. 

	- whenever we create a slice using string literal, it first creates an array of it and then creates a reference of it to a slice.
	- we can create slice using make function as well.
		
		var a = make([]int, len, cap)
	- we can also create a nill value slice. And it wont contain a reference to any array.
	- slices are reference type, so if we make any change in slice then it will be reflected in refereenced array as well.
	- we cant compare two slices using == operator. If we try to do so then compiler will throw an error.
	- we can only compare a slice with nil
	- we can sort slices using sort package, it contains three methods : int, float64, and string
	- we can copy a slice to another using copy() 
	-	it returns the number of items copied.
	- items will be copied as per the size of our slice.
	- since slice is a reference type, when it is passsed to a function. All changes made will be reflected in the actual one. But not in the case of array.

Comparing two slices:-
	
	In the Go slice, you are allowed to compare two slices of the byte type with each other using Compare() function. This function returns an integer value which represents that these slices are equal or not and the values are:

		If the result is 0, then slice_1 == slice_2.
		If the result is -1, then slice_1 < slice_2.
		If the result is +1, then slice_1 > slice_2.

	- We can also check the equality of the slices using Equal() method (only for byte slices)
	- Equal() returns either true or false.
        - We can also trim the byte slice using Trim()
        - We can also perform splitting functionality using Split().


Structure
==========
	- We can compare two structs, there are two ways:- 
		- with the help of == operator
		- with the help of reflect package (reflect.DeepEqual()) 
	- We can have nested structs.
	
String
===========
		
	- String in Golang are different from other langs, here every character is represented by one or more bytes using UTF_8 encoding.
	- String are immutable, string is a readonly slice of bytes.
	- individual element of a string can be accessed in different ways.
	- we can also print the individual bytes  of the string.
	- we can compare strings through two methods:-
		- using relational operators           //returns bool value
		- using strings.Compare()             // always returns 0 if two of them are equal
	- concatenation :- 
		fmt.Sprintf()
		strings.Join(string_name, "seq")
	- to trim a string	;-
		strings.Trim()
		strings.TrimLeft()
		strings.TrimRight()
		strings.TrimSpace()
		strings.TrimSuffix()
		strings.TrimPrefix()

	- we can split a string using strings.Split()
	- we can check, whether the string is containing the value or not using strings.Contains()
	- we can repeat strings using strings.Repeat(str, int) function.

	- we can find the index of a string in a string using following functions.
		strigns.Index()
		strigns.IndexAny()
		strigns.IndexByte()				
	- we can also count the occurance of a specific instance in a string using strings.Count().


Pointers 
=========
	- Pointers are the variables that store the address of another variable. The address is always in the form of hexadecimal form (0XFFAF)
	- A zero-value pointer will always have a <nil> value.
	- Golang doesn't provide any support of pointer arithmetic. unlike (c, c++)
	- By default array is value type. But with the help of the pointers we can pass the reference of the array.

	- we can make the slice out of the array then we can pass it to any function, the changes will be reflected in the actual array.
	- structure with pointers.  
	- we can compare the pointers using == or != operators.
	- we can also apply the cap() and the len() functions, on the pointer to an array.


Interfaces
=============
	- Interfaces don't describe data, they describe behaviours.
	- In Golang, the interfaces are implemented implicitly. And we don't have any specific keyword to specify for the implementaion of the interface
	- Interface Types: The interface is of two types one is static and another one is dynamic type. 
		The static type is the interface itself. 
		But interface does not have a static value so it always points to the dynamic values.
		A variable of the interface type containing the value of the Type which implements the interface 
		so the value of that Type is known as dynamic value and the type is the dynamic type. It is also known as concrete value and concrete type.
	- In an interface, a variable of an interface type can contain any value which implements the interface. This property helps interfaces to achieve polymorphism in the Go language.
   
    =========
	 Type Assertion:-
		- Is an operation applied to the values of the interface.
		- It simply means that extracting the values of the interface.

	
Concurrency
===========
	- Goroutines
	=============
		- Golang has the special feature named Goroutines (It is a lightweighted thread)
		- These are the functions or methods which executed independently or simultaneously in the connection with other goroutines present in the system.
		- The cost of creating goroutines are very less as compared of creating threads.
		- Every program contains atleast a single go routine, that is known as main Goroutine
		- All the goroutines work under the main one, if main one is terminated then all the corresponding ones will also be terminated.
		- You can create your own Goroutine simply by using go keyword as a prefixing to the function or method call.

		- When we call a function as a goroutine, goroutine call returns immedietlty.
			The program does not wait for the execution of the goroutine. They always move forward to the next line and ignore the result generated by the goroutine.
			
Advantages of Goroutines
	Goroutines are cheaper than threads.
	Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.
	Goroutines can communicate using the channel and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines.
	Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.

Multiple goroutines
==============
	- We can control the execution of the goroutines with the help of the time package.
	- Goroutines are not hardware dependent, unlike threads
	- they are managed by runtime.

Channels in golang
====================
	- Channels are the medium through which goroutines communicate with each other. With the help of channels they can send or receive data.
	- Channels are by default bidirectional.
	- In Go language, a channel is created using chan keyword and it can only transfer data of the same type, different types of data are not allowed to transport from the same channel.
	- int, float64, bool and string data are safe to be sent over the channel.
	But pointer and reference type like slice, map are not safe.
	When you use any of them, just make sure that they are accessible by only channel.

- A channel that can only receive data or a channel that can only send data is the unidirectional channel.
===============

Errors in go
=============
	- Go does not have an exception mechanism like try/catch in Java, we cannot throw an exception in Go.
		but go uses a different mechanism that is defer-panic-and-recover.
	- go handles simple errors for functoins and methods by returning an error object. This object maybe the only or the last return value. Its value is nil if there is no error.
	- we should always check errors at the calling statements:-
		- it returns two variables, one is value and other is an error code which is nil in case of success (no error), !=nil in case of an error.
	- the error is always checked after the function call.
	- the user defines the panic function. User raises the panic. When we raise a panic() command, the program will stop the execution.
	- panic() is just like the Throw keyword in java.
	- when we use defer statement and we have raised a panic, defer will run before the termination of the program. But this statement should preceed the panic statement.
	- using recover(), we cannot prevent the termination of the program but it will let us recover from the panic situation.
