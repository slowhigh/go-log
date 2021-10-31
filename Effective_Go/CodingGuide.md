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
            fmt.Println("9")
        case a < 8:
            fmt.Println("8")
        case a < 7:
            fmt.Println("7")
        case a < 6:
            fmt.Println("6")
    }

    // prints
    ```