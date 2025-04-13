# NoBl Design Documentation
_Last updated: April 13, 2025_

## Language Philosophy
NoBl (No Bloat) is designed as a lightweight, efficient programming language focusing on simplicity and performance. The guiding principles are:

1. **Minimalism**: Include only essential features, avoiding redundant syntax
2. **Functional Paradigm**: Emphasize functional programming concepts
3. **Readability**: Maintain clear, unambiguous syntax
4. **Performance**: Design for efficient execution

## Syntax Design

### Core Syntax Elements

#### Variable Bindings
```
const identifier = expression;
```
NoBl uses `const` for variable bindings to encourage immutability and functional programming patterns.

#### Functions
```
function(param1, param2) { 
    expression; 
    return expression;
}

const sumNumbers = function (a,b) {
    return a + b;
}

const sum = sumNumbers(a, b)
```
Functions are first-class citizens, designed to be concise and expressive.

#### Control Flow
```
if (condition) {
    expression;
} else {
    expression;
}
```
Traditional if-else structure for readability and familiarity.

### Operators

| Category | Operators | Description |
|----------|-----------|-------------|
| Arithmetic | +, -, *, / | Standard arithmetic operations |
| Comparison | ==, !=, <, > | Equality and ordering comparisons |
| Logical | ! | Logical negation |

### Type System
The initial implementation supports:
- Integers
- Booleans
- Functions

Future plans include:
- Strings
- Arrays
- Hash maps

## Language Feature Rationale

### Why 'const' instead of 'let' or 'var'?
- Encourages immutable data patterns
- Helps prevent side effects and makes code easier to reason about
- Aligns with functional programming principles

### Why C-style syntax?
- Familiar to most programmers
- Clear block structure with braces
- Semicolons provide explicit statement termination

### Excluded Features
- Classes and object-oriented constructs
- Complex type systems
- Exception handling mechanisms
- Special operator precedence rules beyond mathematical convention

## Implementation Considerations

### Memory Management
- Planning for a simple garbage collection approach
- No manual memory management needed by the user

### Performance Goals
- Fast startup time
- Efficient function calls
- Minimal memory overhead

## Future Design Directions

### Planned Features
1. Pattern matching
2. Simple module system
3. Built-in functions for common operations
4. Immutable data structures

### Potential Extensions
1. Concurrent programming primitives
2. REPL with enhanced debugging features
3. Foreign function interface for C integration

## Design Influences
NoBl's design is influenced by:
- JavaScript's flexible functions
- Go and C's minimalism
