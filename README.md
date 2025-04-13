# NoBl Programming Language

## Introduction

NoBl (No Bloat) is a lightweight, efficient interpreted programming language focusing on simplicity and performance. Following a functional paradigm, the project involves developing a complete language implementation, including lexer, parser, AST, compiler, and virtual machine.

The core philosophy of NoBl is to provide developers with a tool that minimizes unnecessary overhead, allowing for rapid development and efficient execution.

## Current Status

**Current Phase:** Parser Implementation (Week 8/18)

The project has completed the lexical analysis phase and is now focused on implementing the parser component. The lexer can successfully tokenize NoBl source code, and development is underway on a recursive descent parser with Pratt parsing for expressions.

### Completed Components:
- ✅ Token system implementation
- ✅ Lexical analyzer (lexer)
- ✅ Basic REPL (Read-Eval-Print Loop) for testing
- ✅ Comprehensive lexer testing

### In Progress:
- 🔄 AST (Abstract Syntax Tree) design
- 🔄 Recursive descent parser implementation
- 🔄 Pratt parsing for expressions

## Language Syntax

NoBl uses a clean, C-style syntax that will be familiar to most programmers:

```
// Variable binding
const x = 5;

// Function definition
const add = function(a, b) {
    return a + b;
};

// Function call
const result = add(5, 10);

// Conditional
if (result > 10) {
    return true;
} else {
    return false;
}
```

### Key Language Features

- **Functional Paradigm**: Functions are first-class citizens
- **Immutability**: Encourages immutable data patterns through `const` declarations
- **No OOP Bloat**: Focus on simplicity without classes and inheritance hierarchies
- **Minimal Syntax**: Only essential language constructs

## Documentation

This project is documented through several key documents:

1. **[Learning Outcomes Coverage](docs/learning_outcomes.md)** - Describes how the project addresses each learning outcome
2. **[Research Plan](docs/research_plan.md)** - Details research methodology, questions, and findings
3. **[Research Plan Implementation](docs/research_plan_implementation.md)** - Implementation findings and answers to research questions.
4. **[Implementation Details](docs/implementation_details.md)** - Technical documentation of the language implementation
5. **[Design Documentation](docs/design_documentation.md)** - Explores language design decisions and rationale
6. **[Project Pitch](docs/Project%20Pitch.md)** - Initial Project Pitch

## Timeline

| Phase | Timeline | Status |
|-------|----------|--------|
| Research & Design | Weeks 1-4 | ✅ Complete |
| Lexical Analysis | Weeks 5-8 | ✅ Complete |
| Parsing | Weeks 9-12 | 🔄 In Progress |
| Evaluation & Standard Library | Weeks 13-14 | ⏱️ Pending |
| Optimization | Weeks 15-16 | ⏱️ Pending |
| Documentation & Distribution | Weeks 17-18 | ⏱️ Pending |

## Running NoBl

Currently, you can run the NoBl REPL to experiment with the lexer:

```bash
go run main.go
```

This will start an interactive session where you can type NoBl code and see the tokens produced by the lexer.

## Development

### Technology Stack
- **Implementation Language**: Go
- **Testing**: Go's built-in testing framework
- **Development Environment**: Standard Go toolchain

### Project Structure
- `token/` - Token definitions and utilities
- `lexer/` - Lexical analyzer implementation
- `repl/` - Interactive read-eval-print loop
- `parser/` - Parser implementation (in progress)
- `ast/` - Abstract Syntax Tree definitions (in progress)

## Future Work

After completing the parser, upcoming work includes:
- Interpreter/evaluator implementation
- Standard library development
- Virtual machine implementation
- Memory management and optimization
- Distribution and packaging

## Learning Outcomes

This project demonstrates advanced software engineering skills across multiple domains:
- Language design and implementation
- Recursive descent parsing
- Interpreter construction
- Performance optimization
- Software architecture

## Acknowledgements

This project is influenced by several resources:
- "Writing an Interpreter in Go" by Thorsten Ball
- "Crafting Interpreters" by Robert Nystrom
- "Top Down Operator Precedence" by Vaughan Pratt
