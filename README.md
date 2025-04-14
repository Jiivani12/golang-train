# golang-train
session1
in this session, we learn about basic language from GOLANG


    basic knowledge used on terminal :
    -cd.. (going back to previous folder)

    -go mod init <nama-project> (used for made Go Mofules or Initialising a project)
    -go run (used for execution a files)
    -go build (used for making a new file on the same folder)
    -go get (used for downloading new package)

    basic knowledge :
    -fmt.Print = (basic output)
    -fmt.Printf = (the output is depends from flag that we used and have the same function with Println but need extra on naming structure; ex : %T (to knew data types from variable), %s (mostly used on string), %d (used to formatting interger data type from 0-9)) (extra information = https://pkg.go.dev/fmt) (and add \n at the end to make new line(like Println))
    -fmt.Println = (having the same function with Printf but more neat)

    1. Variables (var)
        Naming : use camelCase (ex: favoriteColor ="blue")
        -Variable WITH data type : (tbh im not sure how to explain the variable with data type, but its more like more clearly because after u use var, you add the specific data types on it; ex : <var name string = "Golang">)

        -Variable WITHOUT data type : (we can make some variables without adding the data types, and its called Type Inference that the system from golang will choose on its own the data types based from the value that we add before on variable)
            ex : 
            var name = "Golang"
            var age = 25

            fmt.Printf("%T, %T", name, age)
            (the output will shown as <string, int>)

        -Variable WITHOUT data types (Short Declaration) (:=)
            we can make variables using Short Declaration (:=) without using var 
            ex :
            name := "Golang"
            age := 23

            fmt.Printf("%T, %T", name, age)
            (the output will shown as Golang 23)

        -Multiple Variables Declaration : (we can make more than one variables in one line)
        
        -Underscore Variable : (well basicly adding underscore (_)just to avoid error from useless variables)

    2. Data Types
        -Interger (int = number) : (interger is a non-decimal numeric data type in the GO Language)
        They are again arithmetic types and they represents a> interger types or b> floating point values throughout the program.
            -uint(8,16,32,64) : unsigned 8/16/32/64-bit intergers (0 to xxx)
            -int(8,16,32,64) : signed 8/16/32/64-bit intergers (-xxx to xxx)

        -Float  (the value of an n-bit interger is n bits and is represented using two's complement arithmetic operations.)
            -float32 : IEEE-754 32-bit floating-point numbers
            -float64 : IEEE-754 64-bit floating-point numbers
            -complex64 : complex numbers with float32 real and imaginary parts
            -complex128 : complex numbers with float64 real and imaginary parts

            -verb%f : used to format decimal numbers to 6 digit
            -verb%.nf : used to reduce the decimal numbers depends on 'n'
        (tbh i still dont get the float, maybe used for higher function?)

        -Bool (in Go, bool is the single type representing Boolean Values. it is widely used in logical and conditional operations. it can accept only two values wich are "true" or "false")

        -String (a string type represents the set of string values. its value is a sequence of bytes. strings are immutable types that is once created, it is not possible to change the conttents of a string. the predeclared string type is string.)

        -Nil & Zero Value
            -Zero value from string is ""
            -Zero value from bool is false
            -Zero value from non-decimal numeric is 0
            -Zero value from decimal numeric is 0.0
    
    3. Constants & Operators
        -Constant (conts) on Go language means the value is absolute. for example : Phi, gravity, speed of light, ect.

        -Operators (Arithmetic Operators) basicly using basic math inside the codes.
            - + (adds two operands)
            - - (substracts second operand from the first)
            - * (multiples both operands)
            - / (divides the numerator by the denominator)
            - % (modulus operator; gives the remainder after an interger division)
            - ++ (increment operator. it increases the interger value by one)
            - -- (decrement operator. it decreases the interger value by one)

        -Operators (Relational Operators) we can use relational operators / comparison operators to check the truth of a condition which will produce a value with boolean data type.




        

