## gofk

A brainfuck interpreter in Go

### Reference

1. File of ASCII text ending with `.b`
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
| .    | Outputs value, if greater than a byte, modulo 256 first           |
| ,    | Requests one byte of input, and sets cell to value                |
