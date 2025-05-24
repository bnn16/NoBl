**Document Title: Building NoBl: A Journey into Interpreter Construction**

**Page Allocation (Approximate):**

*   **Page 1:** Title Page (Your Name, Course, Project Title, Date)
*   **Page 2:** Table of Contents
*   **Page 3-4:** Chapter 1: Introduction - The Genesis of NoBl
    *   1.1. Why Build a Programming Language?
    *   1.2. Introducing NoBl: Philosophy and Features
    *   1.3. The Roadmap: Lexing, Parsing, and Evaluation
    *   1.4. Choosing Go: A Practical Foundation
*   **Page 5-6:** Chapter 2: Lexical Analysis - Breaking Down the Code
    *   2.1. From Source to Tokens: The Role of the Lexer
    *   2.2. Defining NoBl's Vocabulary: Tokens
    *   2.3. Implementing the NoBl Lexer
    *   2.4. The First Interactive Step: A Read-Lex-Print-Loop (RLPL)
*   **Page 7-9:** Chapter 3: Parsing - Building Structure from Tokens
    *   3.1. The Parser's Task: Syntax Analysis and ASTs
    *   3.2. Designing NoBl's Abstract Syntax Tree (AST)
    *   3.3. Implementing the NoBl Parser: Top-Down Operator Precedence (Pratt Parsing)
    *   3.4. Parsing Statements and Expressions in NoBl
    *   3.5. Evolving Interaction: A Read-Parse-Print-Loop (RPPL)
*   **Page 10-12:** Chapter 4: Evaluation - Bringing NoBl to Life
    *   4.1. The Heart of the Interpreter: The Evaluator
    *   4.2. Representing Values: NoBl's Object System
    *   4.3. Evaluating Expressions and Statements
    *   4.4. Managing State: Environments and Bindings
    *   4.5. The Power of Functions: Closures and Application
    *   4.6. The Complete Cycle: The Read-Eval-Print-Loop (REPL)
*   **Page 13-14:** Chapter 5: Extending NoBl - Richer Data Types and Built-ins
    *   5.1. Adding Strings
    *   5.2. Implementing Arrays
    *   5.3. Introducing Hashes (Dictionaries)
    *   5.4. Built-in Functions: Expanding Capabilities (e.g., `len`, `puts`)
*   **Page 15:** Chapter 6: Conclusion - Reflections and Future Directions
    *   6.1. Summary of the Journey
    *   6.2. Key Learnings
    *   6.3. Potential Enhancements for NoBl

---

**Page 1: Title Page**

*   **Project Title:** Building NoBl: A Journey into Interpreter Construction
*   **Your Name:** Bogdan Nikolov
*   **Student ID:** 501077
*   **Course:** Advanced Software
*   **University:** Fontys University
*   **Date:** 24.05.2025

---

**Page 2: Table of Contents**

## Table of Contents

- **Chapter 1: Introduction - The Genesis of NoBl** (Page 3)
  - [1.1. Why Build a Programming Language?](#11-why-build-a-programming-language) (Page 3)
  - [1.2. Introducing NoBl: Philosophy and Features](#12-introducing-nobl-philosophy-and-features) (Page 3)
  - [1.3. The Roadmap: Lexing, Parsing, and Evaluation](#13-the-roadmap-lexing-parsing-and-evaluation) (Page 4)
  - [1.4. Choosing Go: A Practical Foundation](#14-choosing-go-a-practical-foundation) (Page 4)

- **Chapter 2: Lexical Analysis - Breaking Down the Code** (Page 5)
  - [2.1. From Source to Tokens: The Role of the Lexer](#21-from-source-to-tokens-the-role-of-the-lexer) (Page 5)
  - [2.2. Defining NoBl's Vocabulary: Tokens](#22-defining-nobls-vocabulary-tokens) (Page 5)
  - [2.3. Implementing the NoBl Lexer](#23-implementing-the-nobl-lexer) (Page 6)
  - [2.4. The First Interactive Step: A Read-Lex-Print-Loop (RLPL)](#24-the-first-interactive-step-a-read-lex-print-loop-rlpl) (Page 6)

- **Chapter 3: Parsing - Building Structure from Tokens** (Page 7)
  - [3.1. The Parser's Task: Syntax Analysis and ASTs](#31-the-parsers-task-syntax-analysis-and-asts) (Page 7)
  - [3.2. Designing NoBl's Abstract Syntax Tree (AST)](#32-designing-nobls-abstract-syntax-tree-ast) (Page 7)
  - [3.3. Implementing the NoBl Parser: Top-Down Operator Precedence (Pratt Parsing)](#33-implementing-the-nobl-parser-top-down-operator-precedence-pratt-parsing) (Page 8)
  - [3.4. Parsing Statements and Expressions in NoBl](#34-parsing-statements-and-expressions-in-nobl) (Page 9)
  - [3.5. Evolving Interaction: A Read-Parse-Print-Loop (RPPL)](#35-evolving-interaction-a-read-parse-print-loop-rppl) (Page 9)

- **Chapter 4: Evaluation - Bringing NoBl to Life** (Page 10)
  - [4.1. The Heart of the Interpreter: The Evaluator](#41-the-heart-of-the-interpreter-the-evaluator) (Page 10)
  - [4.2. Representing Values: NoBl's Object System](#42-representing-values-nobls-object-system) (Page 10)
  - [4.3. Evaluating Expressions and Statements](#43-evaluating-expressions-and-statements) (Page 11)
  - [4.4. Managing State: Environments and Bindings](#44-managing-state-environments-and-bindings) (Page 12)
  - [4.5. The Power of Functions: Closures and Application](#45-the-power-of-functions-closures-and-application) (Page 12)
  - [4.6. The Complete Cycle: The Read-Eval-Print-Loop (REPL)](#46-the-complete-cycle-the-read-eval-print-loop-repl) (Page 12)

- **Chapter 5: Extending NoBl - Richer Data Types and Built-ins** (Page 13)
  - [5.1. Adding Strings](#51-adding-strings) (Page 13)
  - [5.2. Implementing Arrays](#52-implementing-arrays) (Page 13)
  - [5.3. Introducing Hashes (Dictionaries)](#53-introducing-hashes-dictionaries) (Page 14)
  - [5.4. Built-in Functions: Expanding Capabilities](#54-built-in-functions-expanding-capabilities) (Page 14)

- **Chapter 6: Conclusion - Reflections and Future Directions** (Page 15)
  - [6.1. Summary of the Journey](#61-summary-of-the-journey) (Page 15)
  - [6.2. Key Learnings](#62-key-learnings) (Page 15)
  - [6.3. Potential Enhancements for NoBl](#63-potential-enhancements-for-nobl) (Page 15)

- **Chapter 7: References** (Page 16)
---

**Page 3**

**Chapter 1: Introduction - The Genesis of NoBl**

**1.1. Why Build a Programming Language?**

The endeavor of creating a programming language from scratch is often perceived as a Herculean task, reserved for seasoned computer scientists. However, it is also one of the most enlightening journeys a developer can undertake. Building a language, even a simple one, demystifies the tools we use daily. It provides a profound understanding of how code, which starts as mere text, is transformed into actions performed by a computer. This project, the creation of the "NoBl" (No Bloat) programming language, was embarked upon with this spirit of learning and discovery, aiming to peel back the layers of abstraction and understand the fundamental mechanics of language implementation.

The primary motivation was not to create the next industry-standard language, but to gain a practical, hands-on understanding of compiler and interpreter theory. By constructing each component, from tokenization to evaluation, one can truly appreciate the intricacies and design decisions involved in language development.

**1.2. Introducing NoBl: Philosophy and Features**

NoBl, short for "No Bloat," is conceived as a simple, dynamic, and expressive programming language. Its design philosophy emphasizes clarity, minimalism, and ease of understanding, both for the user of the language and for its implementer. The goal was to build a language that is powerful enough to be interesting, yet simple enough to be implemented from the ground up within the scope of this project.

NoBl inherits several characteristics from languages like JavaScript, Ruby, and Lisp, but with a C-like syntax to offer a degree of familiarity. Its core features include:

*   **C-like Syntax:** Familiar curly braces, semicolons, and keyword-based statements.
*   **Variable Bindings:** Using `let` statements (e.g., `let x = 10;`).
*   **Primitive Data Types:** Integers and Booleans as foundational types.
*   **Arithmetic Expressions:** Standard infix operators (`+`, `-`, `*`, `/`) with operator precedence.
*   **Boolean Expressions:** Comparison operators (`<`, `>`, `==`, `!=`) and logical negation (`!`).
*   **Conditionals:** `if/else` expressions.
*   **Functions:** First-class and higher-order functions, supporting closures. This means functions are values that can be assigned to variables, passed as arguments, and returned from other functions, carrying their lexical scope with them.
*   **Data Structures:** Strings, Arrays, and Hashes (key-value maps).
*   **Built-in Functions:** A small set of essential functions provided by the interpreter itself.

The "No Bloat" aspect refers to the intention of keeping the language core small and focused, avoiding unnecessary complexities while still providing a rich learning experience in its construction.

---

**Page 4**

**1.3. The Roadmap: Lexing, Parsing, and Evaluation**

The construction of the NoBl interpreter follows a well-trodden path in language implementation, typically broken down into three main stages:

1.  **Lexical Analysis (Lexing):** This is the first phase where the raw source code, a sequence of characters, is broken down into a series of "tokens." Tokens are the basic building blocks of the language – like words in a sentence. For example, `let x = 5;` would be transformed into tokens representing `let`, the identifier `x`, the assignment operator `=`, the integer `5`, and the semicolon `;`. This process is handled by a component called a "lexer" or "tokenizer."

2.  **Parsing (Syntactic Analysis):** The stream of tokens produced by the lexer is then fed into a "parser." The parser's job is to analyze the syntactic structure of the tokens, ensuring they form valid statements and expressions according to the language's grammar. The output of this phase is typically an Abstract Syntax Tree (AST). An AST is a tree-like data structure that represents the code's structure in a way that's easier for the interpreter to work with, abstracting away details like whitespace or parentheses that are only syntactically important.

3.  **Evaluation (Interpretation):** With the AST constructed, the final stage is to "evaluate" or "interpret" it. This is where the code "comes to life." The interpreter traverses the AST, node by node, performing the operations and computations defined by the code. For example, an AST node representing `5 + 10` would be evaluated to produce the result `15`. This stage involves managing program state, handling function calls, and producing the final output or effects of the program.

Each of these stages builds upon the previous one, culminating in a working interpreter for NoBl.

**1.4. Choosing Go: A Practical Foundation**

The NoBl interpreter is implemented in the Go programming language. This choice was deliberate, guided by several of Go's strengths that align well with the educational goals of this project:

*   **Simplicity and Readability:** Go has a relatively small, C-like syntax that is easy to learn and read. This is crucial for a project where understanding the implementation details is paramount. The code written for the interpreter aims to be clear and directly translatable to the concepts being implemented.
*   **Strong Standard Library:** Go's standard library provides essential tools without requiring external dependencies for core functionalities like string manipulation, data structures (maps, slices), and I/O.
*   **Built-in Tooling:** Go comes with excellent built-in tools for formatting (`gofmt`), testing, and building, which streamline the development process and allow focus on the interpreter logic rather than on setting up a complex development environment.
*   **Performance:** While not the primary goal, Go offers good performance, which is a desirable characteristic for an interpreter. Its static typing and compiled nature provide a solid base.
*   **Concurrency (Not Used Extensively, but a Plus):** Although NoBl's initial implementation doesn't focus on concurrency, Go's built-in support for goroutines and channels makes it a language well-suited for potential future extensions in that direction.

For this project, Go's straightforwardness allows the concepts of interpreter design to shine through without being obscured by complex language features or boilerplate.

---

**Page 5**

**Chapter 2: Lexical Analysis - Breaking Down the Code**

**2.1. From Source to Tokens: The Role of the Lexer**

The journey of interpreting NoBl source code begins with lexical analysis, commonly known as "lexing" or "tokenizing." Raw source code, as a human writes it, is just a long string of characters. For a computer program to understand this code, it first needs to be broken down into meaningful, discrete units. These units are called "tokens."

Consider a simple NoBl statement: `let age = 21;`

The lexer scans this string character by character and groups them into a sequence of tokens. For the statement above, the lexer would ideally produce something like:
`[LET, IDENTIFIER("age"), ASSIGN, INTEGER(21), SEMICOLON, EOF]`

Each token has a type (e.g., `LET`, `IDENTIFIER`) and often a literal value (e.g., the string "age" for the identifier, the number 21 for the integer). The `EOF` (End Of File) token is special; it signals that there's no more input to process.

The lexer's role is crucial. It simplifies the subsequent parsing phase by providing a higher-level, structured view of the source code. The parser doesn't have to deal with individual characters, whitespace, or comments (which the lexer typically discards or handles appropriately). It works with a clean stream of tokens.

**2.2. Defining NoBl's Vocabulary: Tokens**

Before implementing the lexer, we must first define the "vocabulary" of NoBl – the set of all possible tokens our language will recognize. For NoBl, these tokens can be categorized as:

*   **Keywords:** Reserved words with special meaning in the language, such as `let`, `fn` (for function), `if`, `else`, `return`, `true`, `false`.
*   **Identifiers:** Names given to variables and functions by the programmer (e.g., `age`, `myFunction`, `x`).
*   **Literals:** Representations of fixed values:
    *   **Integers:** e.g., `10`, `0`, `123`.
    *   **Strings:** (To be added later) e.g., `"hello"`.
*   **Operators:** Symbols that perform operations, such as:
    *   Assignment: `=`
    *   Arithmetic: `+`, `-`, `*`, `/`
    *   Comparison: `<`, `>`, `==`, `!=`
    *   Logical: `!` (bang for negation)
*   **Delimiters:** Symbols used for grouping or separation, such as:
    *   Parentheses: `(`, `)`
    *   Braces: `{`, `}`
    *   Brackets: `[`, `]` (for arrays, to be added later)
    *   Comma: `,`
    *   Semicolon: `;`
    *   Colon: `:` (for hashes, to be added later)
*   **Special Tokens:**
    *   `ILLEGAL`: Represents a character or sequence of characters that the lexer doesn't recognize.
    *   `EOF`: Signifies the end of the input.

In our Go implementation, a token is typically represented by a struct containing its type (an enumerable or string constant) and its literal string value as it appeared in the source code.

```go
// token/token.go (Conceptual)
package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

// Example TokenTypes
const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    IDENT = "IDENT" // add, foobar, x, y, ...
    INT   = "INT"   // 1343456

    ASSIGN = "="
    PLUS   = "+"
    // ... other operators and delimiters

    LET      = "LET"
    FUNCTION = "FUNCTION"
    // ... other keywords
)
```

---

**Page 6**

**2.3. Implementing the NoBl Lexer**

The NoBl lexer is designed as a struct that holds the input string, the current position being examined, and the next reading position (to peek ahead). Its core functionality is encapsulated in a `NextToken()` method. Each call to `NextToken()` consumes characters from the input and returns the next recognized token.

The lexer's logic within `NextToken()` typically involves:

1.  **Skipping Whitespace:** Whitespace characters (spaces, tabs, newlines) are usually insignificant for NoBl's syntax (unlike Python, for example) and are consumed and ignored.
2.  **Character-based Dispatch:** The lexer looks at the current character to decide what kind of token it might be.
    *   If it's a letter, it reads a sequence of letters and digits to form an identifier or a keyword. A lookup in a keyword map then determines if the identifier is actually a keyword.
    *   If it's a digit, it reads a sequence of digits to form an integer literal.
    *   If it's an operator like `=`, `+`, `!`, it might check the next character to see if it's part of a two-character operator (e.g., `==`, `!=`).
    *   If it's a delimiter like `(`, `)`, `{`, `}`, `;`, it typically forms a token directly.
    *   If it's a quotation mark, it starts reading a string literal (covered in extensions).
3.  **Advancing Position:** After a token is formed, the lexer advances its internal pointers to the next character to be processed.
4.  **Handling EOF and Unknown Characters:** If the end of the input is reached, an `EOF` token is returned. If an unrecognized character is encountered, an `ILLEGAL` token is produced.

```go
// lexer/lexer.go (Conceptual)
package lexer

import "nobl/token"

type Lexer struct {
    input        string
    position     int  // current position in input (points to current char)
    readPosition int  // current reading position (after current char)
    ch           byte // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar() // Initialize l.ch, l.position, l.readPosition
    return l
}

// readChar advances pointers and updates l.ch
func (l *Lexer) readChar() { /* ... */ }

// NextToken identifies and returns the next token
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.ch {
    case '=':
        if l.peekChar() == '=' { // Check for '=='
            // ... create EQ token
        } else {
            tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
        }
    case '+':
        tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
    // ... cases for other characters, digits, letters
    case 0: // Null byte, signals EOF
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(l.ch) {
            // readIdentifier, then check if it's a keyword
            // ...
            return tok // Early return for identifiers/keywords
        } else if isDigit(l.ch) {
            // readNumber
            // ...
            return tok // Early return for numbers
        } else {
            tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
        }
    }
    l.readChar() // Advance for single-character tokens
    return tok
}

func (l *Lexer) peekChar() byte { /* ... looks at input[l.readPosition] ... */ }
func isLetter(ch byte) bool { /* ... checks if ch is a letter or '_' ... */ }
func isDigit(ch byte) bool { /* ... checks if ch is a digit ... */ }
func (l *Lexer) skipWhitespace() { /* ... consumes whitespace ... */ }
```
This iterative process of calling `NextToken()` allows the lexer to transform the entire source code string into a sequence of tokens for the parser.

**2.4. The First Interactive Step: A Read-Lex-Print-Loop (RLPL)**

To test and interact with the lexer incrementally, a simple Read-Lex-Print Loop (RLPL) was developed. This is a precursor to the full Read-Eval-Print Loop (REPL) that will be the ultimate interactive interface for NoBl.

The RLPL:
1.  **Reads** a line of input from the user.
2.  Passes this input to a new instance of the NoBl **Lexer**.
3.  Repeatedly calls `NextToken()` on the lexer instance until an `EOF` token is encountered.
4.  **Prints** each token (its type and literal value) to the console.
5.  **Loops** back to step 1.

This RLPL provided an immediate way to verify that the lexer was correctly identifying different NoBl constructs like keywords, identifiers, integers, and operators. It served as an invaluable debugging tool during the early stages of development.

---

**Page 7**

**Chapter 3: Parsing - Building Structure from Tokens**

**3.1. The Parser's Task: Syntax Analysis and ASTs**

Once the lexer has successfully transformed the NoBl source code into a stream of tokens, the next crucial stage is parsing, also known as syntactic analysis. The parser takes this sequence of tokens and attempts to build a more structured representation of the code, typically an Abstract Syntax Tree (AST).

The AST is a hierarchical tree structure that mirrors the syntactic structure of the source code. It's "abstract" because it omits details that are only important for the raw text, like semicolons (in many cases), specific types of whitespace, or the exact delimiters used for grouping if the grouping itself is represented by the tree structure. Each node in the AST represents a construct in the source code, such as a statement, an expression, an identifier, or a literal.

For example, the NoBl code `let x = 5 + 10;` might be represented by an AST like:

```
Program
  └── LetStatement (Token: LET)
        ├── Name: Identifier (Token: IDENT, Value: "x")
        └── Value: InfixExpression (Token: PLUS)
              ├── Left: IntegerLiteral (Token: INT, Value: 5)
              ├── Operator: "+"
              └── Right: IntegerLiteral (Token: INT, Value: 10)
```

The parser has two main responsibilities:
1.  **Syntax Checking:** It verifies that the sequence of tokens conforms to the grammatical rules of NoBl. If it encounters a sequence that violates these rules (e.g., `let = 5 x;`), it reports a syntax error.
2.  **AST Construction:** If the syntax is correct, the parser constructs the AST, which then becomes the input for the evaluation phase.

For NoBl, we chose to write a recursive descent parser, specifically one that employs the "Top-Down Operator Precedence" parsing technique, also known as Pratt parsing. This approach is elegant for handling expressions with varying operator precedences (e.g., `*` having higher precedence than `+`).

**3.2. Designing NoBl's Abstract Syntax Tree (AST)**

Before implementing the parser, we needed to define the structure of our AST nodes. In Go, this was achieved by defining interfaces and structs.

A core set of interfaces:
*   `Node`: The base interface for all AST nodes. It requires a `TokenLiteral()` method to return the literal value of the token associated with the node (useful for debugging).
*   `Statement`: An interface for nodes that represent statements (which typically don't produce values). It embeds `Node`.
*   `Expression`: An interface for nodes that represent expressions (which typically produce values). It embeds `Node`.

Concrete AST node types were then defined as structs implementing these interfaces:

*   `Program`: The root node of any AST, holding a slice of `Statement`s.
*   `LetStatement`: Represents `let x = <expression>;`. Fields: `Token` (the `LET` token), `Name` (*Identifier), `Value` (Expression).
*   `ReturnStatement`: Represents `return <expression>;`. Fields: `Token` (the `RETURN` token), `ReturnValue` (Expression).
*   `ExpressionStatement`: Represents an expression used as a statement (e.g., `x + 5;`). Fields: `Token` (first token of expression), `Expression` (the actual expression).
*   `Identifier`: Represents an identifier (e.g., `x`). Fields: `Token` (the `IDENT` token), `Value` (string name).
*   `IntegerLiteral`: Represents an integer (e.g., `5`). Fields: `Token` (the `INT` token), `Value` (int64).
*   `Boolean`: Represents `true` or `false`. Fields: `Token`, `Value` (bool).
*   `PrefixExpression`: Represents expressions like `!true` or `-5`. Fields: `Token` (the prefix operator, e.g., `!`), `Operator` (string), `Right` (Expression).
*   `InfixExpression`: Represents expressions like `5 + 5`. Fields: `Token` (the infix operator, e.g., `+`), `Left` (Expression), `Operator` (string), `Right` (Expression).
*   `IfExpression`: Represents `if (<condition>) { <consequence> } else { <alternative> }`. Fields: `Token` (the `IF` token), `Condition` (Expression), `Consequence` (*BlockStatement), `Alternative` (*BlockStatement, optional).
*   `BlockStatement`: A sequence of statements enclosed in `{}`. Fields: `Token` (the `{` token), `Statements` ([]Statement).
*   `FunctionLiteral`: Represents `fn(<parameters>) { <body> }`. Fields: `Token` (the `FN` token), `Parameters` ([]*Identifier), `Body` (*BlockStatement).
*   `CallExpression`: Represents `myFunction(arg1, arg2)`. Fields: `Token` (the `(` token), `Function` (Expression, typically an Identifier or FunctionLiteral), `Arguments` ([]Expression).

Each of these structs also implements methods like `statementNode()` or `expressionNode()` (dummy methods to satisfy interfaces) and `String()` for debugging and pretty-printing the AST.

---

**Page 8**

**3.3. Implementing the NoBl Parser: Top-Down Operator Precedence (Pratt Parsing)**

The NoBl parser maintains two tokens: `curToken` (the current token being processed) and `peekToken` (the next token, used for lookahead). It advances these tokens using a `nextToken()` helper. The main entry point is `ParseProgram()`, which iteratively calls `parseStatement()` until it encounters an `EOF` token, populating the `Program` AST node.

The `parseStatement()` method decides which specific parsing function to call based on `curToken.Type`:
*   If `token.LET`, call `parseLetStatement()`.
*   If `token.RETURN`, call `parseReturnStatement()`.
*   Otherwise, assume it's an expression and call `parseExpressionStatement()`.

The real elegance of Pratt parsing comes into play in `parseExpression(precedence)`. This is the heart of expression parsing. The Pratt parser associates parsing functions with token types, distinguishing between *prefix* and *infix* positions.

*   **Prefix parsing functions (`prefixParseFn`)**: These are called when a token appears at the beginning of an expression or sub-expression (e.g., the `-` in `-5`, an identifier `x`, an integer `10`, or an opening parenthesis `(`).
*   **Infix parsing functions (`infixParseFn`)**: These are called when a token appears after an expression, acting as an infix operator (e.g., the `+` in `5 + 10`, or the `[` in `myArray[0]`).

The parser stores these associations in maps:
`prefixParseFns map[token.TokenType]prefixParseFn`
`infixParseFns map[token.TokenType]infixParseFn`

The `parseExpression(precedence)` method works roughly as follows:
1.  Find the `prefixParseFn` for `curToken.Type`. If none, it's an error (or an unexpected token).
2.  Call this `prefixParseFn`. This parses the initial part of the expression (e.g., an integer literal, an identifier, or a prefix expression like `!foo`). Let this result be `leftExp`.
3.  Then, it enters a loop: as long as `peekToken` is not a semicolon (or other statement terminator) and `peekToken`'s precedence is higher than the `precedence` argument passed to `parseExpression`:
    a.  Find the `infixParseFn` for `peekToken.Type`. If none, the `leftExp` is complete for this precedence level, so break the loop.
    b.  Advance tokens (`nextToken()`, so `curToken` is now the infix operator).
    c.  Call the `infixParseFn`, passing `leftExp` as its left-hand side. This infix function will, in turn, call `parseExpression` (recursively) to parse its right-hand side, passing its own (the operator's) precedence.
    d.  The result of the `infixParseFn` becomes the new `leftExp`.
4.  Return `leftExp`.

Precedences are defined for operators:
```go
// parser/parser.go (Conceptual Precedences)
const (
    _ int = iota
    LOWEST
    EQUALS      // ==
    LESSGREATER // > or <
    SUM         // +
    PRODUCT     // *
    PREFIX      // -X or !X
    CALL        // myFunction(X)
    INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
    token.EQ:       EQUALS,
    token.NOT_EQ:   EQUALS,
    token.LT:       LESSGREATER,
    // ... and so on for other operators like PLUS, ASTERISK, LPAREN (for CALL), LBRACKET (for INDEX)
}
```
The `LPAREN` token is registered as an infix operator to handle function calls (`identifier(args)`), and `LBRACKET` is registered as an infix operator to handle array indexing (`array[index]`). This allows the parser to seamlessly handle expressions like `add(1, 2)[0] * 5`.

Error handling is also integrated. The parser collects errors in a slice. If `expectPeek(expectedType)` is called and `peekToken` is not of the `expectedType`, an error is recorded.

---

**Page 9**

**3.4. Parsing Statements and Expressions in NoBl**

Building on the Pratt parsing mechanism, specific parsing methods were developed for each NoBl construct:

*   **`parseLetStatement()`:**
    *   Expects `curToken` to be `LET`.
    *   Calls `expectPeek(token.IDENT)` to consume the identifier name.
    *   Calls `expectPeek(token.ASSIGN)` to consume the `=` sign.
    *   Calls `nextToken()` to move to the start of the expression.
    *   Calls `parseExpression(LOWEST)` to parse the value expression.
    *   Consumes an optional semicolon.
    *   Constructs and returns an `ast.LetStatement`.

*   **`parseReturnStatement()`:**
    *   Expects `curToken` to be `RETURN`.
    *   Calls `nextToken()` to move to the start of the expression.
    *   Calls `parseExpression(LOWEST)` to parse the return value expression.
    *   Consumes an optional semicolon.
    *   Constructs and returns an `ast.ReturnStatement`.

*   **`parseExpressionStatement()`:**
    *   Calls `parseExpression(LOWEST)` to parse the expression.
    *   Consumes an optional semicolon if `peekToken` is a semicolon.
    *   Constructs and returns an `ast.ExpressionStatement`.

*   **Prefix Functions (registered in `prefixParseFns`):**
    *   `parseIdentifier()`: Returns an `ast.Identifier` using `curToken`.
    *   `parseIntegerLiteral()`: Parses `curToken.Literal` into an int64 and returns an `ast.IntegerLiteral`.
    *   `parseBoolean()`: Returns an `ast.Boolean` based on `curToken` (`TRUE` or `FALSE`).
    *   `parsePrefixExpression()`: Creates an `ast.PrefixExpression`. `curToken` is the operator (e.g., `!`, `-`). Calls `nextToken()` then `parseExpression(PREFIX)` for the right-hand operand.
    *   `parseGroupedExpression()`: For `( expression )`. `curToken` is `LPAREN`. Calls `nextToken()`, then `parseExpression(LOWEST)`, then `expectPeek(token.RPAREN)`.
    *   `parseIfExpression()`: For `if (condition) consequence else alternative`. Handles `expectPeek` for `LPAREN`, `RPAREN`, `LBRACE`. Parses condition expression, consequence `BlockStatement`, and optional `else` alternative `BlockStatement`.
    *   `parseFunctionLiteral()`: For `fn (parameters) { body }`. Handles `expectPeek` for `LPAREN`. Calls a helper to parse comma-separated parameters until `RPAREN`. Expects `LBRACE` and parses the `BlockStatement` for the body.
    *   `parseArrayLiteral()`: For `[element1, element2, ...]`. `curToken` is `LBRACKET`. Calls a helper `parseExpressionList(token.RBRACKET)` to parse comma-separated expressions until `RBRACKET`.
    *   `parseHashLiteral()`: For `{key1: value1, ...}`. `curToken` is `LBRACE`. Loops, parsing key expression, expecting `COLON`, parsing value expression. Handles commas and expects `RBRACE`.

*   **Infix Functions (registered in `infixParseFns`):**
    *   `parseInfixExpression(left ast.Expression)`: `curToken` is the infix operator. Creates `ast.InfixExpression` with `left`. `precedence = curPrecedence()`. Calls `nextToken()`. Calls `parseExpression(precedence)` for the right-hand operand.
    *   `parseCallExpression(function ast.Expression)`: `curToken` is `LPAREN`. `function` is the `left` part (e.g., identifier of the function). Calls `parseExpressionList(token.RPAREN)` to get arguments. Returns `ast.CallExpression`.
    *   `parseIndexExpression(left ast.Expression)`: `curToken` is `LBRACKET`. `left` is the array/hash being indexed. Calls `nextToken()`, then `parseExpression(LOWEST)` for the index, then `expectPeek(token.RBRACKET)`. Returns `ast.IndexExpression`.

*   **Helper `parseBlockStatement()`:** `curToken` is `LBRACE`. Loops, calling `parseStatement()` until `RBRACE` or `EOF`.
*   **Helper `parseExpressionList(endToken token.TokenType)`:** Parses a list of comma-separated expressions until `endToken`. Used for function arguments and array elements.

This systematic approach, combining registered parsing functions with precedence climbing, allows the NoBl parser to handle a wide range of expressions and statements robustly.

**3.5. Evolving Interaction: A Read-Parse-Print-Loop (RPPL)**

With the parser implemented, the interactive loop was upgraded from an RLPL to an RPPL (Read-Parse-Print Loop).
1.  **Reads** input.
2.  Passes it to the **Lexer** to get tokens.
3.  Passes the tokens to the **Parser** to get an AST (`Program` node).
4.  If parser errors occurred, they are printed.
5.  Otherwise, the `String()` method of the `Program` AST node is called, which recursively calls `String()` on its children, effectively **printing** a string representation of the parsed code structure.
6.  **Loops**.

This RPPL was invaluable for verifying the parser's output, ensuring correct AST construction and operator precedence handling by visually inspecting the stringified AST.

---

**Page 10**

**Chapter 4: Evaluation - Bringing NoBl to Life**

**4.1. The Heart of the Interpreter: The Evaluator**

With a lexer tokenizing source code and a parser building an Abstract Syntax Tree (AST), the next and most exciting stage is evaluation. This is where the NoBl code, represented by the AST, is executed, and its meaning is realized. The component responsible for this is the evaluator.

For NoBl, we implemented a tree-walking evaluator. This means the evaluator recursively traverses the AST, node by node. At each node, it performs an action or computes a value based on the type of the node. The core of the evaluator is a central function, conventionally named `Eval`.

The `Eval` function takes an AST node and the current "environment" (which we'll discuss shortly) as input and returns an "object" representing the result of evaluating that node.

```go
// evaluator/evaluator.go (Conceptual Signature)
func Eval(node ast.Node, env *object.Environment) object.Object {
    // ... logic to dispatch based on node type ...
}
```
This recursive structure is powerful. For instance, to evaluate an infix expression like `2 + 3`, `Eval` would:
1.  Be called with the `InfixExpression` node for `2 + 3`.
2.  Recursively call `Eval` on the left child (the `IntegerLiteral` node for `2`), which would return an internal representation of the integer 2.
3.  Recursively call `Eval` on the right child (the `IntegerLiteral` node for `3`), returning an internal representation of 3.
4.  Finally, it would perform the addition on these internal representations and return a new internal representation of 5.

**4.2. Representing Values: NoBl's Object System**

During evaluation, we need a way to represent the values that NoBl expressions produce. These are not the AST nodes themselves (which represent code structure) but the actual data (like the number 5, the boolean `true`, or a function).

To achieve this, NoBl has an internal object system. All values manipulated by the NoBl interpreter are wrapped in Go structs that implement a common `object.Object` interface:

```go
// object/object.go (Conceptual)
package object

type ObjectType string

type Object interface {
    Type() ObjectType
    Inspect() string // Returns a string representation for display
}
```
Specific object types were then defined:

*   **`Integer`**: Wraps an `int64` value. `Inspect()` returns its string form.
    ```go
    type Integer struct { Value int64 }
    ```
*   **`Boolean`**: Wraps a `bool` value. `Inspect()` returns `"true"` or `"false"`. For efficiency, only two instances, `TRUE` and `FALSE`, are pre-allocated and reused.
    ```go
    type Boolean struct { Value bool }
    var TRUE = &Boolean{Value: true}
    var FALSE = &Boolean{Value: false}
    ```
*   **`Null`**: Represents the absence of a value. It has no underlying Go value. A single `NULL` instance is reused.
    ```go
    type Null struct{}
    var NULL = &Null{}
    ```
*   **`ReturnValue`**: A special wrapper object used internally to signal that a `return` statement has been encountered. It wraps the actual value being returned.
    ```go
    type ReturnValue struct { Value Object }
    ```
*   **`Error`**: Represents runtime errors. It wraps an error message string.
    ```go
    type Error struct { Message string }
    ```
*   **`Function`**: Represents NoBl functions. This is key for first-class functions. It stores:
    *   `Parameters` ([]*ast.Identifier): The function's formal parameters.
    *   `Body` (*ast.BlockStatement): The AST for the function's body.
    *   `Env` (*Environment): The environment in which the function was defined. This is crucial for closures.
    ```go
    type Function struct {
        Parameters []*ast.Identifier
        Body       *ast.BlockStatement
        Env        *Environment
    }
    ```
*   **`StringObject`**, **`ArrayObject`**, **`HashObject`**: (Covered in Chapter 5) Wrappers for string, slice, and map data respectively.
*   **`Builtin`**: Wraps a Go function that implements a NoBl built-in function.
    ```go
    type BuiltinFunction func(args ...Object) Object
    type Builtin struct { Fn BuiltinFunction }
    ```
This object system provides a uniform way for `Eval` to handle all values produced and manipulated by NoBl programs.

---

**Page 11**

**4.3. Evaluating Expressions and Statements**

The `Eval` function uses a large `switch` statement on the type of the input `ast.Node` to determine how to evaluate it:

*   **`ast.Program`**: Evaluates each statement in the program sequentially. If a `ReturnValue` or `Error` object is encountered, evaluation stops immediately, and that object is propagated up. Otherwise, the result of the last statement is returned.
*   **`ast.ExpressionStatement`**: Evaluates the underlying expression by calling `Eval` on `node.Expression`.
*   **`ast.IntegerLiteral`**: Creates and returns a new `object.Integer` wrapping the literal's value.
*   **`ast.Boolean`**: Returns the pre-allocated `object.TRUE` or `object.FALSE`.
*   **`ast.PrefixExpression`**:
    1.  Evaluates the right-hand operand: `right := Eval(node.Right, env)`.
    2.  If `right` is an error, return it.
    3.  Calls a helper `evalPrefixExpression(node.Operator, right)`.
        *   For `!`: `evalBangOperatorExpression(right)` returns `FALSE` if `right` is `TRUE`, `NULL`, or any truthy integer. Returns `TRUE` otherwise.
        *   For `-`: If `right` is not an `Integer`, returns an `Error`. Otherwise, returns a new `Integer` with the negated value.
*   **`ast.InfixExpression`**:
    1.  Evaluates both operands: `left := Eval(node.Left, env)`, `right := Eval(node.Right, env)`.
    2.  Handles errors from operand evaluation.
    3.  Calls a helper `evalInfixExpression(node.Operator, left, right)`.
        *   If both `left` and `right` are `Integer`s: performs arithmetic (`+`, `-`, `*`, `/`) or comparison (`<`, `>`, `==`, `!=`) and returns an `Integer` or `Boolean`.
        *   If `operator` is `==` or `!=`: performs pointer comparison for `Boolean`s or `Null` (since `TRUE`, `FALSE`, `NULL` are singletons).
        *   For other type mismatches (e.g., `integer + boolean`), returns an `Error`.
*   **`ast.BlockStatement`**: Evaluates statements sequentially. If a `ReturnValue` or `Error` is encountered, it's immediately returned without evaluating further statements in the block.
*   **`ast.IfExpression`**:
    1.  Evaluates the condition: `condition := Eval(node.Condition, env)`.
    2.  If `condition` is an error, return it.
    3.  Calls `isTruthy(condition)`: (non-`FALSE` and non-`NULL` objects are truthy).
    4.  If truthy, evaluates `node.Consequence`.
    5.  Else, if `node.Alternative` exists, evaluates it. Otherwise, returns `NULL`.
*   **`ast.ReturnStatement`**:
    1.  Evaluates the return value: `val := Eval(node.ReturnValue, env)`.
    2.  If `val` is an error, return it.
    3.  Wraps `val` in an `object.ReturnValue` and returns it. This signals to an enclosing `evalProgram` or `evalBlockStatement` to stop further execution and unwrap this value.
*   **`ast.LetStatement`**:
    1.  Evaluates the value expression: `val := Eval(node.Value, env)`.
    2.  If `val` is an error, return it.
    3.  Stores the binding in the current environment: `env.Set(node.Name.Value, val)`. Returns `NULL` (as `let` statements don't produce values).
*   **`ast.Identifier`**: Looks up the identifier's value in the current environment: `val, ok := env.Get(node.Value)`. If not found (`!ok`), returns an `Error`. Otherwise, returns `val`. Also checks for built-in functions if not found in `env`.
*   **`ast.FunctionLiteral`**: Creates an `object.Function` instance, capturing the parameters, body, and the *current* environment (`env`). This captured environment is crucial for closures.
*   **`ast.CallExpression`**: This is more involved:
    1.  Evaluate the function part (which could be an identifier or another function literal): `function := Eval(node.Function, env)`. If error, return.
    2.  Evaluate all argument expressions: `args := evalExpressions(node.Arguments, env)`. If any argument evaluation results in an error, return that error.
    3.  Call `applyFunction(function, args)`:
        *   If `function` is an `object.Function`:
            a.  Create a new, "extended" environment: `extendedEnv := NewEnclosedEnvironment(function.Env)`. The outer environment is the function's *captured* environment.
            b.  Bind the argument values (`args`) to the function's parameter names in `extendedEnv`.
            c.  Evaluate the function's body (`function.Body`) using this `extendedEnv`: `evaluated := Eval(function.Body, extendedEnv)`.
            d.  If `evaluated` is a `ReturnValue`, unwrap it and return the inner value. Otherwise, return `evaluated`.
        *   If `function` is an `object.Builtin`: Call the underlying Go function `builtin.Fn(args...)`.
        *   Otherwise, return an "not a function" error.

This comprehensive dispatch mechanism allows `Eval` to correctly interpret all valid NoBl AST structures.

---

**Page 12**

**4.4. Managing State: Environments and Bindings**

A crucial aspect of any programming language interpreter is managing state, primarily through variable bindings. In NoBl, this is handled by an "Environment" object.

An `object.Environment` essentially acts as a symbol table or a frame in a call stack. It stores key-value pairs where keys are string identifiers (variable names) and values are `object.Object` instances (the values these variables are bound to).

```go
// object/environment.go (Conceptual)
package object

type Environment struct {
    store map[string]Object
    outer *Environment // Pointer to an enclosing (outer) environment
}

func NewEnvironment() *Environment {
    s := make(map[string]Object)
    return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new environment that "extends" an outer one.
func NewEnclosedEnvironment(outer *Environment) *Environment {
    env := NewEnvironment()
    env.outer = outer
    return env
}

func (e *Environment) Get(name string) (Object, bool) {
    obj, ok := e.store[name]
    if !ok && e.outer != nil { // If not found, try the outer environment
        obj, ok = e.outer.Get(name)
    }
    return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
    e.store[name] = val
    return val
}
```

*   **`Get(name)`**: Retrieves the value associated with `name`. If the name is not found in the current environment's `store`, and an `outer` environment exists, it recursively calls `Get` on the `outer` environment. This mechanism allows for lexical scoping.
*   **`Set(name, val)`**: Associates `val` with `name` in the current environment's `store`.

When the interpreter starts, a top-level global environment is created.
*   When a `let` statement like `let x = 10;` is evaluated, `Eval` calls `env.Set("x", evaluatedInteger10)`.
*   When an `Identifier` `x` is evaluated, `Eval` calls `env.Get("x")`.

The `outer` field is fundamental for implementing lexical scoping and closures, as seen in function evaluation.

**4.5. The Power of Functions: Closures and Application**

Functions in NoBl are first-class citizens. This means they can be:
*   Assigned to variables: `let add = fn(a, b) { a + b; };`
*   Passed as arguments to other functions.
*   Returned from other functions.

When an `ast.FunctionLiteral` (e.g., `fn(x) { x + y; }`) is evaluated, an `object.Function` is created. Crucially, this object stores not only its parameters and body (as AST nodes) but also a reference to the environment active *at the time of its definition*. This is the "closure" part.

When this `object.Function` is later called (applied):
1.  A new environment (`extendedEnv`) is created.
2.  The `outer` environment of this `extendedEnv` is set to the *function's captured environment* (`function.Env`), not necessarily the environment of the call site.
3.  The arguments passed to the function are bound to the function's parameter names within this `extendedEnv`.
4.  The function's body is then evaluated within this `extendedEnv`.

This process ensures that when the function body is executed, it has access to:
*   Its own parameters (bound in `extendedEnv`).
*   Any variables that were in scope when the function was *defined* (accessible via `extendedEnv.outer`, which points to `function.Env`).

This enables powerful constructs like:
```nobl
let newAdder = fn(x) {
  fn(y) { x + y }; // Inner function closes over 'x' from newAdder's scope
};

let addFive = newAdder(5); // addFive is now fn(y) { 5 + y }
addFive(3); // Evaluates to 8
```
Here, `addFive` "remembers" the `x` from `newAdder`'s environment, even after `newAdder` has finished executing. This is the essence of closures.

**4.6. The Complete Cycle: The Read-Eval-Print-Loop (REPL)**

With the evaluator capable of handling expressions, statements, bindings, and functions, the interactive loop was finally upgraded to a full REPL:
1.  **Read**: A line of NoBl code is read from the user.
2.  Pass it to the **Lexer** to get tokens.
3.  Pass the tokens to the **Parser** to get an AST. If parser errors, print them and loop.
4.  **Eval**: The AST is passed to `Eval` along with the current top-level environment.
5.  **Print**: If `Eval` returns an `object.Object` (and not `nil`, which can happen for `let` statements or if the last statement was an expression producing `null`), its `Inspect()` method is called, and the result is printed.
6.  **Loop**: The process repeats, maintaining the state of the top-level environment across inputs.

The REPL is the primary interface for interacting with NoBl, allowing users to define variables, create functions, and see the results of their computations immediately. It showcases the entire pipeline: lexing, parsing, and evaluation working in concert.

---

**Page 13**

**Chapter 5: Extending NoBl - Richer Data Types and Built-ins**

With the core interpreter machinery for integers, booleans, and functions in place, the next logical step was to enrich NoBl with more complex data types and essential built-in functionalities. This involved revisiting each stage of the interpreter: adding new tokens, extending the parser to recognize new literals and expressions, defining new object types in the object system, and updating the evaluator to handle these new constructs.

**5.1. Adding Strings**

Strings are a fundamental data type in most programming languages, representing sequences of characters.

1.  **Lexing:**
    *   A new `token.STRING` type was defined.
    *   The lexer was updated to recognize double-quoted character sequences (e.g., `"hello"`). It reads characters between the opening and closing `"` to form the string literal.
2.  **Parsing:**
    *   An `ast.StringLiteral` node type was added, holding the `token.STRING` and its string `Value`.
    *   A prefix parsing function for `token.STRING` was registered with the parser, which simply creates an `ast.StringLiteral` node.
3.  **Object System & Evaluation:**
    *   An `object.StringObject` type was created, wrapping a Go `string`. It implements the `object.Object` interface and a `Hashable` interface (for use as hash keys later).
    *   The evaluator's `Eval` function was updated: when it encounters an `ast.StringLiteral`, it creates and returns an `object.StringObject`.
    *   String concatenation using the `+` operator was implemented. In `evalInfixExpression`, a case was added: if both operands are `StringObject`s and the operator is `+`, a new `StringObject` is returned, containing the concatenated value. Other operators for strings (like `-` or `*`) result in an error.

**5.2. Implementing Arrays**

Arrays in NoBl are ordered collections of heterogeneous elements.

1.  **Lexing:**
    *   New token types `token.LBRACKET` (`[`) and `token.RBRACKET` (`]`) were defined. The comma `token.COMMA` was already present.
    *   The lexer was updated to recognize `[` and `]` as distinct tokens.
2.  **Parsing:**
    *   An `ast.ArrayLiteral` node type was defined, holding the `token.LBRACKET` and a slice of `ast.Expression`s for its `Elements`.
    *   A prefix parsing function for `token.LBRACKET` was registered. It calls `parseExpressionList(token.RBRACKET)` to parse the comma-separated element expressions until a closing `]`.
    *   An `ast.IndexExpression` node type was defined for array access (e.g., `myArray[0]`). It has a `Left` expression (the array) and an `Index` expression.
    *   An infix parsing function for `token.LBRACKET` was registered. When `myArray[...` is parsed, `myArray` is the `left` operand to this function. It then parses the `Index` expression and expects a `token.RBRACKET`. It has a high precedence (INDEX).
3.  **Object System & Evaluation:**
    *   An `object.ArrayObject` type was created, wrapping a Go slice `[]object.Object`.
    *   `Eval` for `ast.ArrayLiteral`: Evaluates each element expression using `evalExpressions` and stores the resulting `object.Object`s in a new `object.ArrayObject`.
    *   `Eval` for `ast.IndexExpression`:
        *   Evaluates the `Left` operand (the array) and the `Index` operand.
        *   If `Left` is an `ArrayObject` and `Index` is an `Integer`:
            *   Performs bounds checking.
            *   If valid, returns the element at that index from the `ArrayObject.Elements`.
            *   If out of bounds, returns `NULL`.
        *   Otherwise, returns an error (e.g., "index operator not supported for type X").

---

**Page 14**

**5.3. Introducing Hashes (Dictionaries)**

Hashes (also known as maps or dictionaries) provide key-value storage.

1.  **Lexing:**
    *   A new token type `token.COLON` (`:`) was defined. `LBRACE` (`{`) and `RBRACE` (`}`) were already available (used for block statements).
    *   The lexer was updated to recognize `:` as a distinct token.
2.  **Parsing:**
    *   An `ast.HashLiteral` node type was defined. It holds the `token.LBRACE` and a `map[ast.Expression]ast.Expression` for its `Pairs`.
    *   A prefix parsing function for `token.LBRACE` was registered. It loops, parsing a key `ast.Expression`, expecting a `token.COLON`, then parsing a value `ast.Expression`. It continues as long as it finds commas, until a closing `RBRACE`.
3.  **Object System & Evaluation:**
    *   A `Hashable` interface was defined in the object system:
        ```go
        type Hashable interface {
            HashKey() HashKey
        }
        type HashKey struct { Type ObjectType; Value uint64 }
        ```
        `Integer`, `Boolean`, and `StringObject` types were made to implement `Hashable`. `StringObject` uses `fnv.New64a()` for hashing its string value. `Integer` uses its `int64` value. `Boolean` uses 0 or 1. The `HashKey` struct includes the `ObjectType` to prevent collisions between, say, the integer `1` and the boolean `true` (if `true` also hashed to `1`).
    *   An `object.HashObject` type was created. It stores `Pairs` as `map[object.HashKey]object.HashPair`, where `HashPair` is `{ Key, Value object.Object }`. Storing the original `Key` object is useful for inspection.
    *   `Eval` for `ast.HashLiteral`:
        *   Iterates over the AST pairs. For each, evaluates the key and value expressions.
        *   Checks if the evaluated key object implements `Hashable`. If not, an error.
        *   Computes the `HashKey` for the key.
        *   Stores the `object.HashPair` (containing the evaluated key and value objects) in the new `object.HashObject.Pairs` map, using the computed `HashKey`.
    *   `Eval` for `ast.IndexExpression` (extending the array logic):
        *   If `Left` is a `HashObject`:
            *   Evaluates the `Index` operand.
            *   Checks if the evaluated `Index` object is `Hashable`. If not, an error.
            *   Computes its `HashKey`.
            *   Looks up this `HashKey` in `HashObject.Pairs`. If found, returns the `Value` from the `HashPair`.
            *   If not found, returns `NULL`.

**5.4. Built-in Functions: Expanding Capabilities**

Built-in functions are Go functions made available within NoBl. They allow NoBl to perform operations that would be difficult or impossible to implement in NoBl itself (e.g., I/O, or low-level data manipulation).

1.  **Object System:**
    *   An `object.BuiltinFunction` type alias for `func(args ...object.Object) object.Object` was defined.
    *   An `object.Builtin` struct was created, wrapping an `object.BuiltinFunction`. It implements `object.Object`.
2.  **Registration & Lookup:**
    *   A global map `builtins map[string]*object.Builtin` was created in the evaluator.
    *   When `evalIdentifier` fails to find an identifier in the current environment, it now also checks this `builtins` map.
3.  **Application:**
    *   The `applyFunction` helper in the evaluator was extended: if the function to be applied is an `object.Builtin`, it simply calls the wrapped Go function `builtin.Fn(args...)`.
4.  **Example Built-ins:**
    *   `len(arg)`:
        *   Expects one argument.
        *   If `arg` is `StringObject`, returns its length as an `Integer`.
        *   If `arg` is `ArrayObject`, returns its number of elements as an `Integer`.
        *   Otherwise, returns an error.
    *   `puts(args...)`: A variadic function. Iterates through `args`, calls `Inspect()` on each, prints it to standard output, and returns `NULL`.
    *   Array functions: `first(arr)`, `last(arr)`, `rest(arr)`, `push(arr, elem)`. These operate on `ArrayObject`s, perform argument type/count checks, and return new `ArrayObject`s or elements (arrays are immutable in NoBl; `push` returns a new array).

These extensions significantly increased NoBl's expressiveness and utility, demonstrating how a simple core can be incrementally enhanced.

---

**Page 15**

**Chapter 6: Conclusion - Reflections and Future Directions**

**6.1. Summary of the Journey**

The development of the NoBl interpreter, from simple character streams to a functioning language with first-class functions, closures, and diverse data structures, has been an intensive and rewarding process. This journey mirrored the classic stages of interpreter construction:

1.  **Lexical Analysis:** We started by transforming raw NoBl source code into a structured sequence of tokens, defining the fundamental vocabulary of our language.
2.  **Parsing:** We then built a Pratt parser to convert this token stream into an Abstract Syntax Tree (AST), giving hierarchical grammatical structure to the code and enabling robust handling of operator precedence.
3.  **Evaluation:** The AST was then brought to life by a tree-walking evaluator. This involved designing an internal object system to represent NoBl's runtime values (integers, booleans, strings, arrays, hashes, functions) and an environment model to manage variable bindings and lexical scope, which was key to implementing closures.
4.  **Extension:** Finally, the core language was extended with richer data types and built-in functions, significantly enhancing its practical utility.

Throughout this process, an interactive Read-Eval-Print-Loop (REPL) served as a constant companion, evolving from a simple token printer (RLPL) to an AST visualizer (RPPL), and finally, to a fully interactive NoBl shell.

**6.2. Key Learnings**

This project provided invaluable insights into several core computer science concepts:

*   **Language Design Trade-offs:** Every feature, from syntax choices to evaluation semantics (e.g., truthiness, array immutability), involves design decisions with cascading effects. The "No Bloat" philosophy helped guide these choices towards simplicity.
*   **The Power of Abstraction:** The AST provides a crucial level of abstraction between the raw syntax and the evaluation logic. Similarly, the internal object system and environment model abstract the details of runtime data and state management.
*   **Recursive Structures and Algorithms:** Both the parser (recursive descent) and the evaluator (tree-walking) are fundamentally recursive, elegantly handling the nested nature of programming language constructs.
*   **Modularity of Interpreter Components:** The distinct phases of lexing, parsing, and evaluation, while interconnected, are largely modular. This modularity allows for focused development and easier extension. For example, adding a new data type primarily involves changes to the object system and evaluator, with minimal or localized changes to the lexer and parser.
*   **Test-Driven Development:** While not explicitly detailed page-by-page in this summary, the original development process heavily relied on writing tests before or alongside implementation. This was crucial for verifying correctness at each small step, especially for complex components like the Pratt parser or the closure mechanism.

**6.3. Potential Enhancements for NoBl**

While NoBl is functional, it remains a relatively simple language. There are numerous avenues for future development and enhancement:

*   **More Data Types:** Floating-point numbers, characters.
*   **Richer Control Flow:** Loops (`for`, `while`), `switch` statements.
*   **Error Handling Improvements:** More detailed error messages, line numbers, perhaps even a try-catch mechanism.
*   **Standard Library:** A more comprehensive set of built-in functions.
*   **Modules/Imports:** A system for organizing code into multiple files.
*   **Performance Optimization:** The current tree-walking evaluator is not the most performant. Future work could explore:
    *   Bytecode compilation and a Virtual Machine (VM).
    *   Just-In-Time (JIT) compilation.
*   **Garbage Collection:** NoBl currently relies on Go's garbage collector. Implementing a custom GC for NoBl objects (if Go's GC was not available or if NoBl were implemented in a language like C) would be a significant undertaking.
*   **Static Type System:** Adding optional or mandatory static typing for improved error detection and potential performance gains.

---

**Page 16**

**References**

Aho, A. V., Lam, M. S., Sethi, R., & Ullman, J. D. (2007). *Compilers: Principles, techniques, & tools* (2nd ed.). Pearson/Addison Wesley.

Ball, T. (2018). *Writing an interpreter in Go*. Self-published.

Crockford, D. (n.d.). *Top down operator precedence*. [http://javascript.crockford.com/tdop/tdop.html](http://javascript.crockford.com/tdop/tdop.html)

Donovan, A. A. A., & Kernighan, B. W. (2015). *The Go programming language*. Addison-Wesley Professional.

Nystrom, R. (2021). *Crafting interpreters*. Genever Benning.

Pratt, V. R. (1973). Top down operator precedence. In *Proceedings of the 1st annual ACM SIGACT-SIGPLAN symposium on principles of programming languages* (pp. 41–51). Association for Computing Machinery. [https://doi.org/10.1145/512927.512931](https://doi.org/10.1145/512927.512931)

The Go Authors. (n.d.). *The Go programming language specification*. [https://go.dev/ref/spec](https://go.dev/ref/spec)

---

The development of the NoBl programming language and interpreter, as detailed in this report, was primarily guided by the methodologies, concepts, and step-by-step implementation approach presented by Ball (2018). Many of the architectural decisions and implementation strategies for NoBl were directly inspired by or adapted from this work. 
