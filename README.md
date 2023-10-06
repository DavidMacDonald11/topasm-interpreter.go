# README
Have you ever wanted to program in assembly, but didn't want to be tethered to one architecture, stuck reading through antiquated documentation? Introducing:

### Topasm - The obviously practical assembly (language)!
Unlike traditional assembly languages, Topasm isn't held down by the necessity of usefulness.
Topasm is a simple asm built for the angry hyper horse - AHH! - CPU architecture.
Whereas x86 or ARM CPUs are expensive and overheat, AHH! CPUs are free and never overheat.
In fact, you already own one!

### Writing Topasm Code
Newlines are significant and each line should only contain one instruction.
All characters after a semicolon are ignored as a comment.

A register is a 64-bit unsigned integer used to hold intermediate values.
There are 10 of them: `#0, #1, #2, ..., #9`.

A value is defined to be one of the following:
* An integer literal (e.g. `307`)
* A character literal (e.g. `'\n'`)
* A value held in a register (e.g. `#7`)
In the following instructions, replace %val with a value and %reg with a register.

Be sure to check out the [examples](examples) folder to see Topasm in action!
Here is an exhaustive list of instructions in Topasm:

#### Simple Value Instructions
- Move: `move %val into %reg`
Set the value of %reg to %val.

- Addition: `add %val into %reg`
Add %val into %reg.

- Subtraction: `sub %val from %reg`
Subtract %val from %reg.

- Increment: `inc %reg`
Increase the value of %reg by 1.

- Decrement: `dec %reg`
Decrease the value of %reg by 1.

#### Advanced Value Instructions
- Multiplication: `mul %reg`
Multiplies the value of the #0 register by %reg.
The product is stored in #0.

- Division: `div %reg`
Divides the value of the #0 register by %reg.
The quotient is stored in #0 and the remainder in #1.

- Compare: `comp %val with %val`
Compares the first %val (a) against the second %val (b).
The eq flag is set to the result of `a == b`.
The lt flag is set to the result of `a < b`.
These flags are only ever changed by another `comp` call.

#### Conditional Instructions
- Label: `%id:`
A label is not actually an instruction, but is important nonetheless.
It marks its location in code and can be jumped to.
The %id is any non-keyword identifier made up of letters and underscores.

- Non-Conditional Jump: `jump %id`
Changes command flow to the label %id.

- Jump If Not Equal: `jumpNE %id`
Jumps to %id if the eq flag is not set.

- Jump If Equal: `jumpEQ %id`
Jumps to %id if the eq flag is set.

- Jump If Less Than: `jumpLT %id`
Jumps to %id if the lt flag is set.

- Jump If Greater Than: `jumpGT %id`
Jumps to %id if both the lt and eq flags are not set.

- Jump If Less Than Or Equal: `jumpLTE %id`
Jumps to %id if either the lt or eq flags are set.

- Jump If Greater Than Or Equal: `jumpGTE %id`
Jumps to %id if the lt flag is not set.

#### Other Instructions
- Call: `call %id`
Calls a built-in function called %id.
Here is an exhaustive list of built-in functions:
    - Print Char: `printc`
    Prints out the value of #0 converted to an ASCII character.
    - Print Int: `printi`
    Prints out the value of #0.
