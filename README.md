# golang-train
session1
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
        -Interger (int = number)
        

