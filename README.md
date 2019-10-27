## bff

A simple brainfuck interpreter in Go

### Usage

1. Clone repo and build (optionally using `-tags TRACE`)
2. `$ ./bff <file-ending-in-.b-or-.bf>`

### Tracing

This interpreter comes with a tracing mechanism to print out a count of all the
instructions that were executed. Build the binary using the tag `TRACE`:
```
$ go build -tags TRACE
$ ./bff bf/hello.bf

Hello World!

*** Tracing activated, printing instruction count: ***
-  --  66
]  --  80
.  --  13
+  --  368
[  --  17
>  --  184
<  --  178
-----
TOTAL: 906
```

**Warning:** Building with tracing will cause the program to run _much_
slower.

### Reference

1. File of ASCII text ending with `.bf`
2. A movable pointer to manipulate an array, of at least 30,000 cells
3. Each cell holds one byte, initial value of zero
4. Instructions operate on cells, they are:

| Name | Value                                                             |
| :--: | -----                                                             |
| +    | Increment value                                                   |
| -    | Decrement value                                                   |
| >    | Move pointer to the right                                         |
| <    | Move pointer to the left                                          |
| [    | Checks value, if `0` control passes to following matching `]`     |
| ]    | Checks value, if nonzero control passes to following matching `[` |
| .    | Outputs value, if greater than a byte                             |
| ,    | Requests one byte of input, and sets cell to value                |

