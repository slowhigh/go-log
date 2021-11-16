##This document summarizes what I consider important by referring to "https://golang.org/doc/effective_go".


1. Commentary

    1.1. Every package should have a package comment, a block comment preceding the package clause. For multi-file packages, the package comment only needs to be present in one file, and any one will do. The package comment should introduce the package and provide information relevant to the package as a whole. It will appear first on the godoc page and should set up the detailed documentation that follows.

    ```
    /*
    Package regexp implements a simple library for regularexpressions.
    The syntax of the regular expressions accepted is:
        regexp:
            concatenation { '|' concatenation }
        concatenation:
            { closure }
        closure:
            term [ '*' | '+' | '?' ]
        term:
            '^'
            '$'
            '.'
            character
            '[' [ '^' ] character-ranges ']'
            '(' regexp ')'
    */
    package regexp
    ```

    1.2. If every doc comment begins with the name of the item it describes, you can use the doc subcommand of the go tool and run the output through grep. Imagine you couldn't remember the name "Compile" but were looking for the parsing function for regular expressions, so you ran the command,

    ```
    $ go doc -all regexp | grep -i parse
    ```

    If all the doc comments in the package began, "This function...", grep wouldn't help you remember the name. But because the package starts each doccomment with the name, you'd see something like this, which recalls the word you're looking for.

    ```
    $ go doc -all regexp | grep -i parse
        Compile parses a regular expression and returns, if successful, a Regexp
        MustCompile is like Compile but panics if the expression cannot be parsed.
        parsed. It simplifies safe initialization of global variables holding
    $
    ```

    ```
    /*
    Package regexp implements a simple library for regularexpressions.

    The syntax of the regular expressions accepted is:

        regexp:
            concatenation { '|' concatenation }
        concatenation:
            { closure }
        closure:
            term [ '*' | '+' | '?' ]
        term:
            '^'
            '$'
            '.'
            character
            '[' [ '^' ] character-ranges ']'
            '(' regexp ')'
    */
    package regexp
    ```

    1.2. If every doc comment begins with the name of the item itdescribes, you can use the doc subcommand of the go tool and runthe output through grep. Imagine you couldn't remember the name"Compile" but were looking for the parsing function for regularexpressions, so you ran the command,
    
    ```
    $ go doc -all regexp | grep -i parse
    ```

    If all the doc comments in the package began, "This function...", grep wouldn't help you remember the name. But because the package starts each doc comment with the name, you'd see something like this, which recalls the word you're looking for.
    
    ```
    $ go doc -all regexp | grep -i parse
        Compile parses a regular expression and returns, if successful, a Regexp
        MustCompile is like Compile but panics if the expression cannot be parsed.
        parsed. It simplifies safe initialization of global variables holding
    $
    ```


2. Names

    2.1. Package names
    For instance, the buffered reader type in the bufio package is called Reader, not BufReader, because users see it as bufio.Reader, which is a clear, concise name. Moreover, because imported entities are always addressed with their package name, bufio.Reader does not conflict with io.Reader.

    the function to make new instances of ring.Ring—which is the definition of a constructor in Go—would normally be called NewRing, but since Ring is the only type exported by the package, and since the package is called ring, it's called just New, which clients of the package see as ring.New.

    2.2. Getters
    If you have a field called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner.

    The use of upper-case names for export provides the hook to discriminate the field from the method. A setter function, if needed, will likely be called SetOwner.
    
    ```
    owner := obj.Owner()
    if owner != user {
        obj.SetOwner(user)
    }
    ```
    
    2.3. Interface names
    By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

    2.4. MixedCaps
    Finally, the convention in Go is to use MixedCaps or mixedCaps rather than underscores to write multiword names.


3. Semicolons

    If the last token before a newline is an identifier (which includes words like int and float64), a basic literal such as a number or string constant, or one of the tokens

    break continue fallthrough return ++ -- ) } 

    One consequence of the semicolon insertion rules is that you cannot put the opening brace of a control structure (if, for, switch, or select) on the next line.

    Write them like this

    ```
    if i < f() {
        g()
    }
    ```

    not like this

    ```
    if i < f()  // wrong!
    {           // wrong!
        g()
    }
    ```


4. Control structures

    4.1. If

    In Go a simple if looks like this:

    ```
    if x > 0 {
        return y
    }
    ```


    4.2. For

    The Go for loop is similar to—but not the same as—C's. It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.

    ```
    // Like a C for
    for init; condition; post { }

    // Like a C while
    for condition { }

    // Like a C for(;;)
    for { }
    ```

    ```
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    ```

    ```
    for key, value := range oldMap {
        newMap[key] = value
    }
    ```

    ```
    for key := range m {
        if key.expired() {
            delete(m, key)
        }
    }
    ```

    ```
    sum := 0
    for _, value := range array {
        sum += value
    }
    ```

    ```
    for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
        fmt.Printf("character %#U starts at byte position %d\n", char, pos)
    }

    // - prints -
    // character U+65E5 '日' starts at byte position 0
    // character U+672C '本' starts at byte position 3
    //character U+FFFD '�' starts at byte position 6
    // character U+8A9E '語' starts at byte position 7
    ```

    Finally, Go has no comma operator and ++ and -- are statements not expressions. Thus if you want to run multiple variables in a for you should use parallel assignment (although that precludes ++ and --).

    ```
    a := [9]int{1,2,3,4,5,6,7,8,9}
    
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    
    fmt.Println(a)
    ```

    4.3. Switch

    Go's switch is more general than C's. The expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found, and if the switch has no expression it switches on true. It's therefore possible—and idiomatic—to write an if-else-if-else chain as a switch.

    ```
    func unhex(c byte) byte {
        switch {
        case '0' <= c && c <= '9':
            return c - '0'
        case 'a' <= c && c <= 'f':
            return c - 'a' + 10
        case 'A' <= c && c <= 'F':
            return c - 'A' + 10
        }
        return 0
    }
    ```

    There is no automatic fall through, but cases can be presented in comma-separated lists.

    ```
    a:= 3;
    
    switch {
        case a < 9:
            fmt.Print("9")
        case a < 8:
            fmt.Print("8")
        case a < 7:
            fmt.Print("7")
        case a < 6:
            fmt.Print("6")
    }

    // prints
    // 9876 (X) => 9 (O)
    ```

    ```
    func shouldEscape(c byte) bool {
        switch c {
        case ' ', '?', '&', '=', '#', '+', '%':
            return true
        }
        return false
    }
    ```

    Although they are not nearly as common in Go as some other C-like languages, break statements can be used to terminate a switch early. Sometimes, though, it's necessary to break out of a surrounding loop, not the switch, and in Go that can be accomplished by putting a label on the loop and "breaking" to that label. This example shows both uses.

    ```
    Loop:     
	
    for a := 0; a < 10; a++ {
        switch {        
            case a < 5:
                fmt.Print("1")
            	break Loop

            case a < 6:
                fmt.Print("2")
            	break Loop

            case a < 7:
                fmt.Print("3")
                break Loop

            case a < 8:
                fmt.Print("4")
            	break Loop
        }
    }
    
    fmt.Println("end")

    // prints
    // 1end
    ```

    4.3. Type switch

    A switch can also be used to discover the dynamic type of an interface variable. Such a type switch uses the syntax of a type assertion with the keyword type inside the parentheses. If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. It's also idiomatic to reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case.

    ```
    var t interface{}
    t = functionOfSomeType()
    switch t := t.(type) {
    default:
        fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
    case bool:
        fmt.Printf("boolean %t\n", t)             // t has type bool
    case int:
        fmt.Printf("integer %d\n", t)             // t has type int
    case *bool:
        fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
    case *int:
        fmt.Printf("pointer to integer %d\n", *t) // t has type *int
    }
    ```

5. Functions

    5.1. Multiple return values

    One of Go's unusual features is that functions and methods can return multiple values.

    In C, a write error is signaled by a negative count with the error code secreted away in a volatile location. In Go, Write can return a count and an error: “Yes, you wrote some bytes but not all of them because you filled the device”. The signature of the Write method on files from package os is:

    ```
    func (file *File) Write(b []byte) (n int, err error)
    ```

    5.2. Multiple return values

    The return or result "parameters" of a Go function can be given names and used as regular variables, just like the incoming parameters. When named, they are initialized to the zero values for their types when the function begins; if the function executes a return statement with no arguments, the current values of the result parameters are used as the returned values.

    ```
    func nextInt(b []byte, pos int) (value, nextPos int) {
    ```

    Because named results are initialized and tied to an unadorned return, they can simplify as well as clarify. Here's a version of io.ReadFull that uses them well:

    ```
    func ReadFull(r Reader, buf []byte) (n int, err error) {
        for len(buf) > 0 && err == nil {
            var nr int
            nr, err = r.Read(buf)
            n += nr
            buf = buf[nr:]
        }
        return
    }
    ```


    5.3. Defer

    Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file.

    ```
    // Contents returns the file's contents as a string.
    func Contents(filename string) (string, error) {
        f, err := os.Open(filename)
        if err != nil {
            return "", err
        }
        defer f.Close()  // f.Close will run when we're finished.

        var result []byte
        buf := make([]byte, 100)
        for {
            n, err := f.Read(buf[0:])
            result = append(result, buf[0:n]...) // append is discussed later.
            if err != nil {
                if err == io.EOF {
                    break
                }
                return "", err  // f will be closed if we return here.
            }
        }
        return string(result), nil // f will be closed if we return here.
    }
    ```


    The arguments to the deferred function (which include the receiver if the function is a method) are evaluated when the defer executes, not when the call executes. Besides avoiding worries about variables changing values as the function executes, this means that a single deferred call site can defer multiple function executions. Here's a silly example.

    ```
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
    ```

    Deferred functions are executed in LIFO order, so this code will cause 4 3 2 1 0 to be printed when the function returns. A more plausible example is a simple way to trace function execution through the program.

    We can do better by exploiting the fact that arguments to deferred functions are evaluated when the defer executes. The tracing routine can set up the argument to the untracing routine. This example:

    ```
    func trace(s string) string {
        fmt.Println("entering:", s)
        return s
    }

    func un(s string) {
        fmt.Println("leaving:", s)
    }

    func a() {
        defer un(trace("a"))
        fmt.Println("in a")
    }

    func b() {
        defer un(trace("b"))
        fmt.Println("in b")
        a()
    }

    func main() {
        b()
    }
    ```

    prints

    ```
    entering: b
    in b
    entering: a
    in a
    leaving: a
    leaving: b
    ```

    For programmers accustomed to block-level resource management from other languages, defer may seem peculiar, but its most interesting and powerful applications come precisely from the fact that it's not block-based but function-based.


6. Data

    6.1. Allocation with new

    Go has two allocation primitives, the built-in functions new and make. They do different things and apply to different types, which can be confusing, but the rules are simple. Let's talk about new first. It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it. That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T. In Go terminology, it returns a pointer to a newly allocated zero value of type T.

    Since the memory returned by new is zeroed, it's helpful to arrange when designing your data structures that the zero value of each type can be used without further initialization. This means a user of the data structure can create one with new and get right to work. For example, the documentation for bytes.Buffer states that "the zero value for Buffer is an empty buffer ready to use." Similarly, sync.Mutex does not have an explicit constructor or Init method. Instead, the zero value for a sync.Mutex is defined to be an unlocked mutex.

    The zero-value-is-useful property works transitively. Consider this type declaration.

    ```
    type SyncedBuffer struct {
        lock    sync.Mutex
        buffer  bytes.Buffer
    }
    ```

    Values of type SyncedBuffer are also ready to use immediately upon allocation or just declaration. In the next snippet, both p and v will work correctly without further arrangement.

    ```
    p := new(SyncedBuffer)  // type *SyncedBuffer
    var v SyncedBuffer      // type  SyncedBuffer
    ```

    6.2. Constructors and composite literals

    Sometimes the zero value isn't good enough and an initializing constructor is necessary, as in this example derived from package os.

    ```
    func NewFile(fd int, name string) *File {
        if fd < 0 {
            return nil
        }
        f := new(File)
        f.fd = fd
        f.name = name
        f.dirinfo = nil
        f.nepipe = 0
        return f
    }
    ```

    There's a lot of boiler plate in there. We can simplify it using a composite literal, which is an expression that creates a new instance each time it is evaluated.
    
    ```
    func NewFile(fd int, name string) *File {
        if fd < 0 {
            return nil
        }
        f := File{fd, name, nil, 0}
        return &f
    }
    ```

    Note that, unlike in C, it's perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns. In fact, taking the address of a composite literal allocates a fresh instance each time it is evaluated, so we can combine these last two lines.

    ```
    return &File{fd, name, nil, 0}
    ```

    The fields of a composite literal are laid out in order and must all be present. However, by labeling the elements explicitly as field:value pairs, the initializers can appear in any order, with the missing ones left as their respective zero values. Thus we could say

    ```
    return &File{fd: fd, name: name}
    ```

    As a limiting case, if a composite literal contains no fields at all, it creates a zero value for the type. The expressions new(File) and &File{} are equivalent.

    Composite literals can also be created for arrays, slices, and maps, with the field labels being indices or map keys as appropriate. In these examples, the initializations work regardless of the values of Enone, Eio, and Einval, as long as they are distinct.

    ```
    a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    ```


    6.3. Allocation with make

    It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T). The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use.

    For instance,

    ```
    make([]int, 10, 100)
    ```

    allocates an array of 100 ints and then creates a slice structure with length 10 and a capacity of 100 pointing at the first 10 elements of the array. (When making a slice, the capacity can be omitted; see the section on slices for more information.) In contrast, new([]int) returns a pointer to a newly allocated, zeroed slice structure, that is, a pointer to a nil slice value.

    ```
    var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
    var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

    // Unnecessarily complex:
    var p *[]int = new([]int)
    *p = make([]int, 100, 100)

    // Idiomatic:
    v := make([]int, 100)
    ```

    Remember that make applies only to maps, slices and channels and does not return a pointer. To obtain an explicit pointer allocate with new or take the address of a variable explicitly.


    6.4. Arrays

    Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation, but primarily they are a building block for slices, the subject of the next section.

    There are major differences between the ways arrays work in Go and C. In Go,

    - Arrays are values. Assigning one array to another copies all the elements.
    - In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
    - The size of an array is part of its type. The types [10]int and [20]int are distinct.

    The value property can be useful but also expensive; if you want C-like behavior and efficiency, you can pass a pointer to the array.

    ```
    func Sum(a *[3]float64) (sum float64) {
        for _, v := range *a {
            sum += v
        }
        return
    }

    array := [...]float64{7.0, 8.5, 9.1}
    x := Sum(&array)  // Note the explicit address-of operator
    ```

    But even this style isn't idiomatic Go. Use slices instead.


    6.5. Slices

    Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

    Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array. A Read function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the Read method of the File type in package os:

    ```
    func (f *File) Read(buf []byte) (n int, err error)
    ```

    The method returns the number of bytes read and an error value, if any. To read into the first 32 bytes of a larger buffer buf, slice (here used as a verb) the buffer.

    ```
    n, err := f.Read(buf[0:32])
    ```

    Such slicing is common and efficient. In fact, leaving efficiency aside for the moment, the following snippet would also read the first 32 bytes of the buffer.

    ```
    var n int
    var err error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
    ```

    The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume. Here is a function to append data to a slice. If the data exceeds the capacity, the slice is reallocated. The resulting slice is returned. The function uses the fact that len and cap are legal when applied to the nil slice, and return 0.