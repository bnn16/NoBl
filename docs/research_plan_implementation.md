# Individual Research Plan: Designing & Implementing NoBl Programming Language
_Last updated: April 13, 2025_

## Research Process & Methodology

This research follows the DOT framework (Library, Lab, and Showroom) to systematically investigate programming language implementation approaches. The methodology involves:

1. **Literature Study Phase**: Analysis of academic papers, books, and online resources
2. **Comparative Analysis Phase**: Examination of existing language implementations
3. **Prototyping Phase**: Creation of experimental implementations to validate concepts
4. **Evaluation Phase**: Assessment of results against predefined criteria

For each research question, I followed this process to ensure thorough investigation and evidence-based decision-making.

## Research Progress Update - Lexical Analysis Phase

### Completed Research Activities

#### Literature Review on Lexer Design
I conducted a review of lexer design principles, focusing on:
- Academic resources on compiler construction and language implementation
- Open-source language implementations
- Performance considerations in lexical analysis

The literature review revealed multiple approaches to lexer implementation, including:
- Table-driven lexers using state machines
- Hand-written recursive descent lexers
- Regular expression-based tokenizers
- Character-by-character scanners with lookahead

After comparing these approaches against my criteria of simplicity, performance, and extensibility, I selected the character-by-character approach with minimal lookahead for NoBl.

#### Analysis of Go-Based Lexer Implementations
I examined several Go-based lexers to understand idiomatic patterns and performance characteristics:
- The Go language's own lexer implementation
- Lexers in multiple open-source Go projects
- Rob Pike's lexical scanning approach

This analysis revealed that the most performant Go lexers typically use straightforward character-by-character scanning rather than complex state machines, which aligned with my design goals for NoBl.

#### Token Design Pattern Evaluation
I systematically evaluated different token representation patterns, considering:
- Memory efficiency
- Type safety
- Debugging capabilities
- Performance characteristics

While numeric constants would be more memory-efficient, I opted for string-based token types to improve debugging and readability. This decision was informed by findings from the literature review that suggested development speed and maintainability were often more valuable than marginal performance gains in token representation.

### Research Questions Addressed

#### 1. ✅ What lexer implementation approach offers the best balance of simplicity and performance?

**Decision: Character-by-character scanner with minimal lookahead**

**Reasoning:**
- **Simplicity**: The character-by-character approach is conceptually straightforward and maps well to Go's strengths.
- **Performance**: Benchmarking of different approaches showed that for small to medium inputs, the performance difference between optimized state machines and direct character scanning was negligible in Go.
- **Extensibility**: The selected approach makes it easier to add new token types without restructuring the entire lexer. This was demonstrated when I added support for multi-character operators (==, !=).
- **Memory Usage**: The approach requires minimal buffer management compared to regex-based approaches, which often build up intermediate representations.

**Implementation Details:**
The implementation uses a `Lexer` struct that tracks:
- Current position in the input
- Next position (for lookahead)
- Current character under examination

This design allowed for efficient implementation of the `NextToken()` method that forms the core of the lexer.

#### 2. ✅ What core language features should be included in the lexical structure?

**Decision: Focus on essential operators and language constructs common in functional languages**

**Rationale:**
- **Minimalism**: Research into functional programming languages showed that a relatively small set of operators and keywords can provide significant expressive power.
- **Familiarity**: Analysis of popular languages indicated that certain syntax patterns (e.g., C-style function calls and conditionals) are widely recognized and reduce learning curve.
- **Expressiveness**: Studies of language design suggest that orthogonality (small number of concepts that can be combined in predictable ways) leads to more powerful languages.

**Token Selection Process:**
I followed a systematic process to determine which tokens to include:
1. Created a baseline set from common operators across multiple languages
2. Added functional programming-specific tokens
3. Removed redundant or rarely used operators
4. Added tokens specific to NoBl's design philosophy

This process resulted in a token set that balances minimalism with expressiveness, supporting all core features without unnecessary complexity.

**Implementation Impact:**
The selected token set directly influenced the lexer design:
- Switch statement for single-character tokens
- Specialized handling for identifiers and numbers
- Lookahead for multi-character operators
- Keyword map for efficient identification

### Ongoing Research Questions

#### 3. What parser implementation approaches offer the best balance of simplicity and performance?

**Current Research Status:**
After reviewing multiple parsing techniques including LL, LR, and recursive descent approaches, I've narrowed the focus to recursive descent parsing with Pratt parsing for expressions.

**Preliminary Findings:**
- **Recursive Descent**: Offers good balance between implementation simplicity and parsing power
- **Pratt Parsing**: Provides elegant solution for expression parsing with operator precedence
- **AST Design**: Tree structure significantly impacts performance of later interpretation phases

**Evaluation Criteria:**
I'm evaluating parser approaches based on:
1. Implementation complexity
2. Performance characteristics with complex input
3. Error reporting capabilities
4. Extensibility for future language features

**Next Steps in Research:**
- Complete prototype implementation of recursive descent parser
- Benchmark against alternative approaches
- Analyze error reporting quality

#### 4. What memory management strategies will provide safety without significant overhead?

**Research Plan:**
This question will be investigated after the parser implementation, focusing on:
- Garbage collection approaches suitable for interpreted languages
- Reference counting vs. tracing collectors
- Memory safety guarantees in functional languages

**Initial Research Activities:**
- Literature review of memory management in similar languages
- Analysis of Go's garbage collector as a potential model
- Investigation of memory safety vulnerabilities in interpreted languages

## Next Research Steps

### 1. Parser Design Investigation

**Research Questions:**
- How can Pratt parsing be optimized for the specific needs of NoBl's expression syntax?
- What AST design will best support both interpretation and future compilation?
- How can error recovery be implemented to provide useful feedback during parsing?

**Methodology:**
- Implement prototype parsers using different approaches
- Benchmark performance on complex expressions
- Evaluate error reporting quality

**Expected Outcomes:**
- Optimized Pratt parsing implementation
- AST design documentation
- Error handling strategy

### 2. AST Design Patterns

**Research Questions:**
- What AST node design provides the best balance of simplicity and extensibility?
- How should different expression types be represented in the AST?
- What visitor patterns will best support interpretation and future optimizations?

**Methodology:**
- Analyze AST designs in existing language implementations
- Prototype different AST structures
- Evaluate impact on interpreter implementation

**Expected Outcomes:**
- Comprehensive AST design documentation
- Node hierarchy specification
- Visitor pattern implementation

### 3. Benchmark Prototypes

**Research Questions:**
- What performance characteristics do different parsing approaches exhibit?
- How do AST designs impact memory usage during parsing?
- What optimizations yield the greatest performance improvements?

**Methodology:**
- Implement benchmark suite for parser operations
- Measure performance across varied input types
- Profile memory usage during parsing

**Expected Outcomes:**
- Quantitative performance data
- Memory usage analysis
- Optimization recommendations

## Resources and Literature Review

### Key Academic Sources
- **Aho, A.V., Lam, M.S., Sethi, R., & Ullman, J.D. (2006).** *Compilers: Principles, Techniques, and Tools.* - Provided foundational understanding of lexer and parser design principles
- **Pratt, V.R. (1973).** *Top Down Operator Precedence.* - Original paper on Pratt parsing technique, key influence for expression parsing strategy

### Practical Implementation Resources
- **Thorsten Ball's "Writing an Interpreter in Go"** - Provided practical implementation patterns for lexer design and token handling
- **Rob Pike's "Lexical Scanning in Go"** - Influenced the character-by-character scanning approach and error handling strategies
- **"Crafting Interpreters" by Robert Nystrom** - Offered insights into interpreter architecture and design considerations

### Code Repositories Analyzed
- **Go Programming Language Source Code** - Studied for idiomatic Go patterns in lexer implementation
- **Lua Implementation** - Analyzed for minimalist language design and implementation techniques

## Research Impact on Implementation

The research findings directly influenced several key aspects of the NoBl implementation:

1. **Lexer Structure**:
   - Adopted character-by-character scanning based on performance research
   - Implemented minimal lookahead mechanism informed by token complexity analysis
   - Designed error handling approach based on lessons from existing implementations

2. **Token Design**:
   - Selected string-based token representation for improved debugging
   - Created efficient keyword lookup mechanism based on map performance research
   - Implemented specialized character classification functions based on lexing patterns

3. **REPL Implementation**:
   - Designed interactive feedback mechanism based on usability research
   - Implemented token visualization approach for educational purposes
   - Created error reporting mechanism informed by user experience studies

## Reflection and Adaptation

Throughout the research process, I've continuously refined my approach based on new findings:

- **Initial Plan Adjustment**: Originally intended to use regular expressions for tokenization, but research revealed performance limitations
- **Methodology Refinement**: Expanded comparative analysis to include more implementation approaches after initial results were inconclusive
- **Scope Adaptation**: Added focused research on Pratt parsing after discovering its relevance to NoBl's expression handling needs

## Conclusion and Path Forward

The research completed thus far has provided a solid foundation for the lexical analysis phase of NoBl. The character-by-character scanner with minimal lookahead has proven to be an effective approach, balancing simplicity, performance, and extensibility.

As the project moves into the parsing phase, the research focus will shift to recursive descent parsing with Pratt parsing for expressions. The lessons learned from the lexical analysis phase, particularly around simplicity and extensibility, will inform the parser implementation.

The DOT framework has proven effective for guiding this research, and I will continue to apply it as the project progresses. By systematically investigating design options, prototyping implementations, and evaluating results, I aim to create a language implementation that truly embodies the "No Bloat" philosophy while providing powerful functional programming capabilities.
