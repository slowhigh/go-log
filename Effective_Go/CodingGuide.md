##It was written for personal study purposes referring to "https://golang.org/doc/effective_go".


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

    ```
    func Append(slice, data []byte) []byte {
        fmt.Printf("1- %v - %p[%p]\n", slice, &slice, slice)
        l := len(slice)
        if l + len(data) > cap(slice) {  // reallocate
            // Allocate double what's needed, for future growth.
            newSlice := make([]byte, (l+len(data))*2)
            fmt.Printf("2- %v - %p[%p]\n", newSlice, &newSlice, newSlice)
            // The copy function is predeclared and works for any slice type.
            copy(newSlice, slice)
            fmt.Printf("3- %v - %p[%p]\n", newSlice, &newSlice, newSlice)
            slice = newSlice
        } 
        fmt.Printf("4- %v - %p[%p]\n", slice, &slice, slice)
        slice = slice[0:l+len(data)]
        fmt.Printf("5- %v - %p[%p]\n", slice, &slice, slice)
        copy(slice[l:], data)
        fmt.Printf("6- %v - %p[%p]\n", slice, &slice, slice)
        return slice
    }

    func main() {
        slice := make([]byte, 5, 10)
        slice2 := []byte { 2,2,2,2,2,2 }
        Append(slice, slice2)
    }

    // - print -
    // 1- [0 0 0 0 0] - 0xc000004078[0xc0000120b0]
    // 2- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc0000040c0[0xc000014198]
    // 3- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc0000040c0[0xc000014198]
    // 4- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc000004078[0xc000014198]
    // 5- [0 0 0 0 0 0 0 0 0 0 0] - 0xc000004078[0xc000014198]
    // 6- [0 0 0 0 0 2 2 2 2 2 2] - 0xc000004078[0xc000014198]
    ```

    We must return the slice afterwards because, although Append can modify the elements of slice, the slice itself (the run-time data structure holding the pointer, length, and capacity) is passed by value.


    6.6. Two-dimensional slices

    Go's arrays and slices are one-dimensional. To create the equivalent of a 2D array or slice, it is necessary to define an array-of-arrays or slice-of-slices, like this:

    ```
    type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
    type LinesOfText [][]byte     // A slice of byte slices.
    ```

    Because slices are variable-length, it is possible to have each inner slice be a different length. That can be a common situation, as in our LinesOfText example: each line has an independent length.

    ```
    text := LinesOfText{
        []byte("Now is the time"),
        []byte("for all good gophers"),
        []byte("to bring some fun to the party."),
    }
    ```

    Sometimes it's necessary to allocate a 2D slice, a situation that can arise when processing scan lines of pixels, for instance. There are two ways to achieve this. One is to allocate each slice independently; the other is to allocate a single array and point the individual slices into it. Which to use depends on your application. If the slices might grow or shrink, they should be allocated independently to avoid overwriting the next line; if not, it can be more efficient to construct the object with a single allocation. For reference, here are sketches of the two methods. First, a line at a time:

    ```
    // Allocate the top-level slice.
    picture := make([][]uint8, 5) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range picture {
		picture[i] = make([]uint8, 3)
		fmt.Printf("- %v - %p[%p]\n", picture, &picture, picture)
		fmt.Printf("- %v - %p[%p]\n", picture[i], &picture[i], picture[i])
	}

	fmt.Printf("- %v - %p\n", picture[0][0], &picture[0][0])
	fmt.Printf("- %v - %p\n", picture[0][1], &picture[0][1])
	fmt.Printf("- %v - %p\n", picture[0][2], &picture[0][2])

    // - [[0 0 0] [] [] [] []] - 0xc000098060[0xc0000e0000]
    // - [0 0 0] - 0xc0000e0000[0xc0000ac058]
    // - [[0 0 0] [0 0 0] [] [] []] - 0xc000098060[0xc0000e0000]
    // - [0 0 0] - 0xc0000e0018[0xc0000ac07b]
    // - [[0 0 0] [0 0 0] [0 0 0] [] []] - 0xc000098060[0xc0000e0000]
    // - [0 0 0] - 0xc0000e0030[0xc0000ac0a4]
    // - [[0 0 0] [0 0 0] [0 0 0] [0 0 0] []] - 0xc000098060[0xc0000e0000]
    // - [0 0 0] - 0xc0000e0048[0xc0000ac0b0]
    // - [[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]] - 0xc000098060[0xc0000e0000]
    // - [0 0 0] - 0xc0000e0060[0xc0000ac0c0]
    // - 0 - 0xc0000ac058
    // - 0 - 0xc0000ac059
    // - 0 - 0xc0000ac05a
    ```

    And now as one allocation, sliced into lines:

    ```
    // Allocate the top-level slice, the same as before.
	picture := make([][]uint8, 5) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]uint8, 3*5) // Has type []uint8 even though picture is [][]uint8.
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture {
		picture[i], pixels = pixels[:3], pixels[3:]
        fmt.Printf("- %v - %p[%p]\n", picture, &picture, picture)
        fmt.Printf("- %v - %p[%p]\n", picture[i], &picture[i], picture[i])
		fmt.Printf("- %v - %p[%p]\n", pixels, &pixels, pixels)
	}

    // - [[0 0 0] [] [] [] []] - 0xc000004078[0xc000112000]
    // - [0 0 0] - 0xc000112000[0xc0000120b0]
    // - [0 0 0 0 0 0 0 0 0 0 0 0] - 0xc000004090[0xc0000120b3]
    // - [[0 0 0] [0 0 0] [] [] []] - 0xc000004078[0xc000112000]
    // - [0 0 0] - 0xc000112018[0xc0000120b3]
    // - [0 0 0 0 0 0 0 0 0] - 0xc000004090[0xc0000120b6]
    // - [[0 0 0] [0 0 0] [0 0 0] [] []] - 0xc000004078[0xc000112000]
    // - [0 0 0] - 0xc000112030[0xc0000120b6]
    // - [0 0 0 0 0 0] - 0xc000004090[0xc0000120b9]
    // - [[0 0 0] [0 0 0] [0 0 0] [0 0 0] []] - 0xc000004078[0xc000112000]
    // - [0 0 0] - 0xc000112048[0xc0000120b9]
    // - [0 0 0] - 0xc000004090[0xc0000120bc]
    // - [[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]] - 0xc000004078[0xc000112000]
    // - [0 0 0] - 0xc000112060[0xc0000120bc]
    // - [] - 0xc000004090[0xc0000120bc]
    ```


    6.7. Maps

    Maps are a convenient and powerful built-in data structure that associate values of one type (the key) with values of another type (the element or value). The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. Slices cannot be used as map keys, because equality is not defined on them. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

    Maps can be constructed using the usual composite literal syntax with colon-separated key-value pairs, so it's easy to build them during initialization.

    ```
    var timeZone = map[string]int{
        "UTC":  0*60*60,
        "EST": -5*60*60,
        "CST": -6*60*60,
        "MST": -7*60*60,
        "PST": -8*60*60,
    }
    ```

    Assigning and fetching map values looks syntactically just like doing the same for arrays and slices except that the index doesn't need to be an integer.

    ```
    offset := timeZone["EST"]
    ```

    An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map.

    ```
    attended := map[string]bool{
        "Ann": true,
        "Joe": true,
        ...
    }

    if attended[person] { // will be false if person is not in the map
        fmt.Println(person, "was at the meeting")
    }
    ```

    Sometimes you need to distinguish a missing entry from a zero value. Is there an entry for "UTC" or is that 0 because it's not in the map at all? You can discriminate with a form of multiple assignment.

    ```
    var seconds int
    var ok bool
    seconds, ok = timeZone[tz]
    ```

    For obvious reasons this is called the “comma ok” idiom. In this example, if tz is present, seconds will be set appropriately and ok will be true; if not, seconds will be set to zero and ok will be false. Here's a function that puts it together with a nice error report:

    ```
    func offset(tz string) int {
        if seconds, ok := timeZone[tz]; ok {
            return seconds
        }
        log.Println("unknown time zone:", tz)
        return 0
    }
    ```

    To test for presence in the map without worrying about the actual value, you can use the blank identifier (_) in place of the usual variable for the value.

    ```
    _, present := timeZone[tz]
    ```

    To delete a map entry, use the delete built-in function, whose arguments are the map and the key to be deleted. It's safe to do this even if the key is already absent from the map.

    ```
    delete(timeZone, "PDT")  // Now on Standard Time
    ```

    

    6.8. Printing

    The functions live in the fmt package and have capitalized names: fmt.Printf, fmt.Fprintf, fmt.Sprintf and so on. The string functions (Sprintf etc.) return a string rather than filling in a provided buffer.

    You don't need to provide a format string. For each of Printf, Fprintf and Sprintf there is another pair of functions, for instance Print and Println. These functions do not take a format string but instead generate a default format for each argument. The Println versions also insert a blank between arguments and append a newline to the output while the Print versions add blanks only if the operand on neither side is a string. In this example each line produces the same output.

    ```
    fmt.Printf("Hello %d\n", 23)
    fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
    fmt.Println("Hello", 23)
    fmt.Println(fmt.Sprint("Hello ", 23))
    ```

    The formatted print functions fmt.Fprint and friends take as a first argument any object that implements the io.Writer interface; the variables os.Stdout and os.Stderr are familiar instances.

    First, the numeric formats such as %d do not take flags for signedness or size; instead, the printing routines use the type of the argument to decide these properties.

    ```
    var x uint64 = 1<<64 - 1
    fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

    // print
    // 18446744073709551615 ffffffffffffffff; -1 -1
    ```

    If you just want the default conversion, such as decimal for integers, you can use the catchall format %v (for “value”); the result is exactly what Print and Println would produce. Moreover, that format can print any value, even arrays, slices, structs, and maps. Here is a print statement for the time zone map defined in the previous section.

    ```
    fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)
    ```

    which gives output:

    ```
    map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
    ```

    For maps, Printf and friends sort the output lexicographically by key.

    When printing a struct, the modified format %+v annotates the fields of the structure with their names, and for any value the alternate format %#v prints the value in full Go syntax.

    ```
    type T struct {
        a int
        b float64
        c string
    }
    t := &T{ 7, -2.35, "abc\tdef" }
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)

	fmt.Printf("%q\n", t)
    fmt.Printf("%+q\n", t)
    fmt.Printf("%#q\n", t)
    fmt.Printf("%#q\n", timeZone)

	fmt.Printf("%x\n", t)
    fmt.Printf("%+x\n", t)
    fmt.Printf("%#x\n", t)
    fmt.Printf("%#x\n", timeZone)


    // &{7 -2.35 abc   def}
    // &{a:7 b:-2.35 c:abc     def}
    // &main.T{a:7, b:-2.35, c:"abc\tdef"}
    // map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}

    // &{'\a' %!q(float64=-2.35) "abc\tdef"}
    // &{'\a' %!q(float64=-2.35) "abc\tdef"}
    // &{'\a' %!q(float64=-2.35000) `abc       def`}
    // map[`CST`:'�' `EST`:'�' `MST`:'�' `PST`:'�' `UTC`:'\x00']

    // &{7 -0x1.2cccccccccccdp+01 61626309646566}
    // &{+7 -0x1.2cccccccccccdp+01 61626309646566}
    // &{0x7 -0x1.2cccccccccccdp+01 0x61626309646566}
    // map[0x435354:-0x5460 0x455354:-0x4650 0x4d5354:-0x6270 0x505354:-0x7080 0x555443:0x0]
    ```

    (Note the ampersands.) That quoted string format is also available through %q when applied to a value of type string or []byte. The alternate format %#q will use backquotes instead if possible. (The %q format also applies to integers and runes, producing a single-quoted rune constant.) Also, %x works on strings, byte arrays and byte slices as well as on integers, generating a long hexadecimal string, and with a space in the format (% x) it puts spaces between the bytes.

    Another handy format is %T, which prints the type of a value.

    ```
    fmt.Printf("%T\n", timeZone)

    // map[string]int
    ```

    If you want to control the default format for a custom type, all that's required is to define a method with the signature String() string on the type. For our simple type T, that might look like this.

    ```
    func (t *T) String() string {
        return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
    }
    fmt.Printf("%v\n", t)
    
    // to print in the format
    // 7/-2.35/"abc\tdef"
    ```
    
    (If you need to print values of type T as well as pointers to T, the receiver for String must be of value type; this example used a pointer because that's more efficient and idiomatic for struct types.)

    Our String method is able to call Sprintf because the print routines are fully reentrant and can be wrapped this way. There is one important detail to understand about this approach, however: don't construct a String method by calling Sprintf in a way that will recur into your String method indefinitely. This can happen if the Sprintf call attempts to print the receiver directly as a string, which in turn will invoke the method again. It's a common and easy mistake to make, as this example shows.

    ```
    type MyString string

    func (m MyString) String() string {
        return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
    }
    ```

    It's also easy to fix: convert the argument to the basic string type, which does not have the method.

    ```
    type MyString string

    func (m MyString) String() string {
        return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
    }
    ```

    Another printing technique is to pass a print routine's arguments directly to another such routine. The signature of Printf uses the type ...interface{} for its final argument to specify that an arbitrary number of parameters (of arbitrary type) can appear after the format.

    ```
    func Printf(format string, v ...interface{}) (n int, err error) {
    ```

    Within the function Printf, v acts like a variable of type []interface{} but if it is passed to another variadic function, it acts like a regular list of arguments. Here is the implementation of the function log.Println we used above. It passes its arguments directly to fmt.Sprintln for the actual formatting.

    ```
    import (
        "fmt"
        "log"
    )

    type MyString string

    func (m MyString) String() string {
        return fmt.Sprintf("MyString=%s", string(m))
    }

    func Println(v ...interface{}) {
        log.Output(2, fmt.Sprintln(v...)) // Output takes parameters (int, string)
    }

    func main() {
        var str MyString = "text"

        Println(str.String())

        // print
        // 2021/12/21 19:19:13 MyString=text
    }
    ```

    We write ... after v in the nested call to Sprintln to tell the compiler to treat v as a list of arguments; otherwise it would just pass v as a single slice argument.

    There's even more to printing than we've covered here. See the godoc documentation for package fmt for the details.

    By the way, a ... parameter can be of a specific type, for instance ...int for a min function that chooses the least of a list of integers:

    ```
    func Min(a ...int) int {
        min := int(^uint(0) >> 1)  // largest int
        for _, i := range a {
            if i < min {
                min = i
            }
        }
        return min
    }

    // '^' => reverse bit
    ```



    6.9. Append

    Now we have the missing piece we needed to explain the design of the append built-in function. The signature of append is different from our custom Append function above. Schematically, it's like this:

    ```
    func append(slice []T, elements ...T) []T
    ```

    where T is a placeholder for any given type. You can't actually write a function in Go where the type T is determined by the caller. That's why append is built in: it needs support from the compiler.


    ```
    x := []int{1,2,3}
    x = append(x, 4, 5, 6)
    fmt.Println(x)
    ```

    prints [1 2 3 4 5 6]. So append works a little like Printf, collecting an arbitrary number of arguments.

    But what if we wanted to do what our Append does and append a slice to a slice? Easy: use ... at the call site, just as we did in the call to Output above. This snippet produces identical output to the one above.

    ```
    x := []int{1,2,3}
    y := []int{4,5,6}
    x = append(x, y...)
    fmt.Println(x)
    ```

    Without that ..., it wouldn't compile because the types would be wrong; y is not of type int.


7. Initialization

    Complex structures can be built during initialization and the ordering issues among initialized objects, even among different packages, are handled correctly.


    7.1. Constants

    Constants in Go are just that—constant. They are created at compile time, even when defined as locals in functions, and can only be numbers, characters (runes), strings or booleans. Because of the compile-time restriction, the expressions that define them must be constant expressions, evaluatable by the compiler. For instance, 1<<3 is a constant expression, while math.Sin(math.Pi/4) is not because the function call to math.Sin needs to happen at run time.

    In Go, enumerated constants are created using the iota enumerator. Since iota can be part of an expression and expressions can be implicitly repeated, it is easy to build intricate sets of values.

    ```
    type ByteSize float64

    const (
        _           = iota // iota = 0, ignore first value by assigning to blank identifier
        KB ByteSize = 1 << (10 * iota) // iota = 1
        MB // = 1 << (10 * iota) , iota = 2
        GB
        TB
        PB
        EB
        ZB
        YB
    )
    ```

    The ability to attach a method such as String to any user-defined type makes it possible for arbitrary values to format themselves automatically for printing. Although you'll see it most often applied to structs, this technique is also useful for scalar types such as floating-point types like ByteSize.

    ```
    func (b ByteSize) String() string {
        switch {
        case b >= YB:
            return fmt.Sprintf("%.2fYB", b/YB)
        case b >= ZB:
            return fmt.Sprintf("%.2fZB", b/ZB)
        case b >= EB:
            return fmt.Sprintf("%.2fEB", b/EB)
        case b >= PB:
            return fmt.Sprintf("%.2fPB", b/PB)
        case b >= TB:
            return fmt.Sprintf("%.2fTB", b/TB)
        case b >= GB:
            return fmt.Sprintf("%.2fGB", b/GB)
        case b >= MB:
            return fmt.Sprintf("%.2fMB", b/MB)
        case b >= KB:
            return fmt.Sprintf("%.2fKB", b/KB)
        }
        return fmt.Sprintf("%.2fB", b)
    }
    ```
    

    The expression YB prints as 1.00YB, while ByteSize(1e13) prints as 9.09TB.

    The use here of Sprintf to implement ByteSize's String method is safe (avoids recurring indefinitely) not because of a conversion but because it calls Sprintf with %f, which is not a string format: Sprintf will only call the String method when it wants a string, and %f wants a floating-point value.



    7.2. Variables

    Variables can be initialized just like constants but the initializer can be a general expression computed at run time.

    ```
    var (
        home   = os.Getenv("HOME")
        user   = os.Getenv("USER")
        gopath = os.Getenv("GOPATH")
    )
    ```

    7.3. The init function

    Finally, each source file can define its own niladic init function to set up whatever state is required. (Actually each file can have multiple init functions.) And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.

    Besides initializations that cannot be expressed as declarations, a common use of init functions is to verify or repair correctness of the program state before real execution begins.

    ```
    func init() {
        if user == "" {
            log.Fatal("$USER not set")
        }
        if home == "" {
            home = "/home/" + user
        }
        if gopath == "" {
            gopath = home + "/go"
        }
        // gopath may be overridden by --gopath flag on command line.
        flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
    }
    ```


8. Methods

    8.1. Pointers vs. Values

    As we saw with ByteSize, methods can be defined for any named type (except a pointer or an interface); the receiver does not have to be a struct.

    In the discussion of slices above, we wrote an Append function. We can define it as a method on slices instead. To do this, we first declare a named type to which we can bind the method, and then make the receiver for the method a value of that type.

    ```
    type ByteSlice []byte

    func (slice ByteSlice) Append(data []byte) []byte {
        l := len(slice)
        if l + len(data) > cap(slice) {  // reallocate
            // Allocate double what's needed, for future growth.
            newSlice := make([]byte, (l+len(data))*2)
            // The copy function is predeclared and works for any slice type.
            copy(newSlice, slice)
            slice = newSlice
        }
        slice = slice[0:l+len(data)]
        copy(slice[l:], data)
        return slice
    }
    ```

    This still requires the method to return the updated slice. We can eliminate that clumsiness by redefining the method to take a pointer to a ByteSlice as its receiver, so the method can overwrite the caller's slice.

    ```
    func (p *ByteSlice) Append(data []byte) {
        slice := *p

        l := len(slice)
        if l+len(data) > cap(slice) { // reallocate
            // Allocate double what's needed, for future growth.
            newSlice := make([]byte, (l+len(data))*2)
            // The copy function is predeclared and works for any slice type.
            copy(newSlice, slice)
            slice = newSlice
        }
        slice = slice[0 : l+len(data)]
        copy(slice[l:], data)

        *p = slice
    }
    ```

    In fact, we can do even better. If we modify our function so it looks like a standard Write method, like this,

    ```
    func (p *ByteSlice) Append3(data []byte) (n int, err error) {
        slice := *p

        l := len(slice)
        if l+len(data) > cap(slice) { // reallocate
            // Allocate double what's needed, for future growth.
            newSlice := make([]byte, (l+len(data))*2)
            // The copy function is predeclared and works for any slice type.
            copy(newSlice, slice)
            slice = newSlice
        }
        slice = slice[0 : l+len(data)]
        copy(slice[l:], data)

        *p = slice
        return len(data), nil
    }
    ```

    then the type *ByteSlice satisfies the standard interface io.Writer, which is handy. For instance, we can print into one.

    ```
    var b ByteSlice
    fmt.Fprintf(&b, "This hour has %d days\n", 7)
    ```

    We pass the address of a ByteSlice because only *ByteSlice satisfies io.Writer. The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.

    This rule arises because pointer methods can modify the receiver; invoking them on a value would cause the method to receive a copy of the value, so any modifications would be discarded. The language therefore disallows this mistake. There is a handy exception, though. When the value is addressable, the language takes care of the common case of invoking a pointer method on a value by inserting the address operator automatically. In our example, the variable b is addressable, so we can call its Write method with just b.Write. The compiler will rewrite that to (&b).Write for us.

    By the way, the idea of using Write on a slice of bytes is central to the implementation of bytes.Buffer.


9. Interface and other types

    9.1. Interfaces

    Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here. We've seen a couple of simple examples already; custom printers can be implemented by a String method while Fprintf can generate output to anything with a Write method. Interfaces with only one or two methods are common in Go code, and are usually given a name derived from the method, such as io.Writer for something that implements Write.

    A type can implement multiple interfaces. For instance, a collection can be sorted by the routines in package sort if it implements sort.Interface, which contains Len(), Less(i, j int) bool, and Swap(i, j int), and it could also have a custom formatter. In this contrived example Sequence satisfies both.

    ```
    import (
        "fmt"
        "sort"
    )

    type Sequence []int

    func (s Sequence) Len() int {
        return len(s)
    }

    func (s Sequence) Less(i, j int) bool {
        return s[i] < s[j]
    }

    func (s Sequence) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
    }

    func(s Sequence) Copy() Sequence {
        copy := make(Sequence, 0, len(s))
        return append(copy, s...)
    }

    func (s Sequence) String() string {
        s = s.Copy()
        sort.Sort(s)
        str := "["
        for i , elem := range s {
            if i > 0 {
                str += " "
            }
            str += fmt.Sprint(elem)
        }
        return str + "]"
    }

    func main() {
        var s1 Sequence = []int{4,1,2,3}
        fmt.Println(s1.String())
    }
    ```


    9.2. Conversions

    The String method of Sequence is recreating the work that Sprint already does for slices. (It also has complexity O(N²), which is poor.) We can share the effort (and also speed it up) if we convert the Sequence to a plain []int before calling Sprint.

    ```
    func (s Sequence) String() string {
        s = s.Copy()
        sort.Sort(s)
        return fmt.Sprint([]int(s))
    }
    ```

    This method is another example of the conversion technique for calling Sprintf safely from a String method. Because the two types (Sequence and []int) are the same if we ignore the type name, it's legal to convert between them. The conversion doesn't create a new value, it just temporarily acts as though the existing value has a new type. (There are other legal conversions, such as from integer to floating point, that do create a new value.)

    It's an idiom in Go programs to convert the type of an expression to access a different set of methods. As an example, we could use the existing type sort.IntSlice to reduce the entire example to this:

    ```
    type Sequence []int

    // Method for printing - sorts the elements before printing
    func (s Sequence) String() string {
        s = s.Copy()
        sort.IntSlice(s).Sort()
        return fmt.Sprint([]int(s))
    }
    ```

    Now, instead of having Sequence implement multiple interfaces (sorting and printing), we're using the ability of a data item to be converted to multiple types (Sequence, sort.IntSlice and []int), each of which does some part of the job. That's more unusual in practice but can be effective.


    9.3. Interface conversions and type assertions

    Type switches are a form of conversion: they take an interface and, for each case in the switch, in a sense convert it to the type of that case. Here's a simplified version of how the code under fmt.Printf turns a value into a string using a type switch. If it's already a string, we want the actual string value held by the interface, while if it has a String method we want the result of calling the method.

    ```
    type Stringer interface {
        String() string
    }

    var value interface{} // Value provided by caller.
    switch str := value.(type) {
    case string:
        return str
    case Stringer:
        return str.String()
    }
    ```

    The first case finds a concrete value; the second converts the interface into another interface. It's perfectly fine to mix types this way.

    What if there's only one type we care about? If we know the value holds a string and we just want to extract it? A one-case type switch would do, but so would a type assertion. A type assertion takes an interface value and extracts from it a value of the specified explicit type. The syntax borrows from the clause opening a type switch, but with an explicit type rather than the type keyword:

    ```
    value.(typeName)
    ```

    and the result is a new value with the static type typeName. That type must either be the concrete type held by the interface, or a second interface type that the value can be converted to. To extract the string we know is in the value, we could write:

    ```
    str := value.(string)
    ```

    But if it turns out that the value does not contain a string, the program will crash with a run-time error. To guard against that, use the "comma, ok" idiom to test, safely, whether the value is a string:

    ```
    str, ok := value.(string)
    if ok {
        fmt.Printf("string value is: %q\n", str)
    } else {
        fmt.Printf("value is not a string\n")
    }
    ```

    If the type assertion fails, str will still exist and be of type string, but it will have the zero value, an empty string.

    As an illustration of the capability, here's an if-else statement that's equivalent to the type switch that opened this section.

    ```
    if str, ok := value.(string); ok {
        return str
    } else if str, ok := value.(Stringer); ok {
        return str.String()
    }
    ```

    9.4. Generality

    If a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the type itself. Exporting just the interface makes it clear the value has no interesting behavior beyond what is described in the interface. It also avoids the need to repeat the documentation on every instance of a common method.

    In such cases, the constructor should return an interface value rather than the implementing type. As an example, in the hash libraries both crc32.NewIEEE and adler32.New return the interface type hash.Hash32. Substituting the CRC-32 algorithm for Adler-32 in a Go program requires only changing the constructor call; the rest of the code is unaffected by the change of algorithm.

    A similar approach allows the streaming cipher algorithms in the various crypto packages to be separated from the block ciphers they chain together. The Block interface in the crypto/cipher package specifies the behavior of a block cipher, which provides encryption of a single block of data. Then, by analogy with the bufio package, cipher packages that implement this interface can be used to construct streaming ciphers, represented by the Stream interface, without knowing the details of the block encryption.

    The crypto/cipher interfaces look like this:

    ```
    type Block interface {
        BlockSize() int
        Encrypt(dst, src []byte)
        Decrypt(dst, src []byte)
    }
    ```

    type Stream interface {
        XORKeyStream(dst, src []byte)
    }
    Here's the definition of the counter mode (CTR) stream, which turns a block cipher into a streaming cipher; notice that the block cipher's details are abstracted away:

    ```
    // NewCTR returns a Stream that encrypts/decrypts using the given Block in
    // counter mode. The length of iv must be the same as the Block's block size.
    func NewCTR(block Block, iv []byte) Stream
    ```

    NewCTR applies not just to one specific encryption algorithm and data source but to any implementation of the Block interface and any Stream. Because they return interface values, replacing CTR encryption with other encryption modes is a localized change. The constructor calls must be edited, but because the surrounding code must treat the result only as a Stream, it won't notice the difference.


    9.5. Interfaces and methods

    