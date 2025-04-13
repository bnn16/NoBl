# NoBl Implementation Details
_Last updated: April 12, 2025_

## Introduction

This document outlines the current state of the NoBl (No Bloat) programming language implementation, specifically focusing on the lexical analysis (lexer) component. NoBl is designed to be a lightweight, efficient interpreted programming language with a focus on simplicity and performance, following a functional paradigm.

## Implementation Progress

### Completed Components

1. **Token System**: Defined a flexible token structure that supports all necessary language elements including:
   - Identifiers and literals (IDENTIFIER, INT)
   - Operators (+, -, *, /, <, >, ==, !=, !)
   - Delimiters ((), {}, ,, ; )
   - Keywords (function, const, true, false, if, else, return)

2. **Lexical Analyzer (Lexer)**: Implemented a complete lexer that:
   - Processes input character by character
   - Converts character sequences into meaningful tokens
   - Handles whitespace appropriately
   - Identifies keywords, identifiers, literals, and operators
   - Supports single and multi-character tokens (e.g., = vs ==)

3. **REPL (Read-Eval-Print Loop)**: Created a basic interactive environment that:
   - Reads user input
   - Tokenizes the input using the lexer
   - Displays the resulting tokens for inspection

### Current Architecture

The lexer implementation follows a straightforward design:

- **Character-by-character processing**: The lexer reads the input string one character at a time, tracking current position and next read position.
- **Token generation**: Based on the current character and context, appropriate tokens are generated.
- **Lookahead mechanism**: For multi-character tokens like "==", the lexer uses a peek function to determine what token to produce.
- **Identifier and number handling**: Special functions to read identifier and number literals as complete tokens.

```
Input String -> Lexer -> Stream of Tokens -> (Future Parser)
```

## Design Decisions

### Language Syntax Design

NoBl's syntax is intentionally familiar to programmers who know C-like languages while remaining minimal. Some key syntax decisions include:

1. **Keywords**: Using `function` for function declarations and `const` for variable bindings
2. **Function Declarations**: Following the format `function(param1, param2) { body }`
3. **Operators**: Supporting standard arithmetic, comparison, and logical operators

### Implementation Decisions

1. **Token Representation**: Using string constants as token types for readability and easier debugging, despite potential performance tradeoffs.
2. **Character Set Limitation**: Currently limited to ASCII encoding for simplicity. A future improvement would be to support full Unicode.
3. **Error Handling**: Basic error tokens (ILLEGAL) are generated for unrecognized input.
4. **Future Extensibility**: The lexer is designed to be easily extended with new token types as needed.

## Next Steps: Parser Implementation

The next phase of development will focus on implementing the parser. The parser will:

1. **Consume tokens** produced by the lexer
2. **Build an Abstract Syntax Tree (AST)** representing the program structure
3. **Handle operator precedence** correctly in expressions
4. **Support all language constructs** including functions, conditionals, and variable declarations

The parser implementation will follow a recursive descent approach with Pratt parsing for expressions, which is particularly well-suited for handling operator precedence in a simple and elegant way.

## Challenges and Solutions

### Challenges Encountered

1. **Multi-character token handling**: Implementing tokens like "==" and "!=" required lookahead capabilities.
   - Solution: Implemented a `peekChar()` method that looks at the next character without consuming it.

2. **Keyword identification**: Distinguishing keywords from identifiers.
   - Solution: Created a lookup map in the token package to efficiently identify keywords.

### Future Challenges

1. **Parser complexity**: Handling expression parsing with correct precedence will require careful implementation.
2. **Error reporting**: Providing meaningful error messages during parsing.
3. **AST design**: Creating a flexible yet efficient Abstract Syntax Tree structure.

## Testing Approach

The current implementation includes comprehensive tests for the lexer component:

1. **Simple token tests**: Verify basic token recognition
2. **Complex token tests**: Ensure correct tokenization of complex program structures
3. **Edge case handling**: Test boundary conditions like EOF

## Conclusion

The lexer implementation provides a solid foundation for the NoBl language interpreter. With the lexical analysis component complete, the project is now ready to move on to the parsing phase, which will translate the token stream into a structured representation of the program that can later be evaluated.