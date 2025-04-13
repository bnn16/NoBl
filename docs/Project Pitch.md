# NoBl Project Plan - Project Pitch

* NoBl is a new, lightweight, and efficient interpreted programming language designed with a focus on simplicity and performance. It will follow a functional paradigm, drawing inspiration from JavaScript's flexibility but avoiding the complexities of object-oriented class structures.
* The project involves developing a complete language implementation, starting with a lexer, parser, and Abstract Syntax Tree (AST) to create an interpreter. Subsequently, a compiler will be built to generate bytecode, which will then be executed by a custom stack-based virtual machine.
* The core philosophy of NoBl is to provide developers with a tool that minimizes unnecessary overhead, allowing for rapid development and efficient execution.

## Challenges and Enterprise Software Development Relevance:

* **Complexity of Language Implementation:**
    * Designing and implementing a programming language from scratch is inherently complex. It requires a deep understanding of compilers, interpreters, and virtual machines. This project provides a very deep understanding and exercise in core computer science principles, which are highly relevant to enterprise software development.
    * In enterprise settings, developers often need to understand how software systems function at a low level to optimize performance and troubleshoot issues. This project will develop those skills.
* **Performance and Efficiency:**
    * The goal of creating a "No Bloat" language directly addresses the enterprise need for efficient and performant software. In enterprise environments, resource optimization is critical, especially in large-scale applications.
    * Creating a compiler and virtual machine allows for deep optimization of the language, which is a valuable skill in enterprise software development.
* **Understanding of Language Design:**
    * Understanding how languages are designed, and why certain design decisions are made, this is very important for enterprise software development. Many enterprise systems are built using Domain specific languages, and the understanding of how to make a language will help with the creation, and maintenance of those systems.

## Non-Functional Requirements:

* **Security:**
    * The language will be designed with security as a primary concern, aiming to minimize vulnerabilities and prevent common security flaws.
* **Memory Safety:**
    * A key goal is to prevent memory leaks and ensure efficient memory management, contributing to the overall stability and reliability of applications built with NoBl.
* **Performance:**
    * The language will be designed to be very performant. The compiled version of the language, utilizing the virtual machine, will be highly optimized.

## Initial Ideas (Architecture and Technology Selection):

* **Lexer, Parser, and AST:**
    * These components will be implemented from scratch via hand-written code. NO AI!!! 
* **Compiler and Virtual Machine:**
    * The compiler will translate the AST into bytecode, and the virtual machine will execute this bytecode using a stack-based architecture.
* **Technology Stack:**
    * C or Go are being considered for the implementation. C offers low-level control and performance, while Go provides concurrency and ease of development. The final choice will depend on further evaluation and experimentation.
* **Functional paradigm:**
    * The language will focus on functions as first class citizens.

## Why this is a good project.

* **Significant Learning Opportunity:** This project provides an in-depth exploration of core computer science concepts, including compilers, interpreters, virtual machines, and language design.
* **Development of Practical Skills:** I will gain hands-on experience in software development, low-level programming, and optimization techniques.
* **Demonstrates Advanced Technical Ability:** Successfully completing this project will showcase my ability to tackle complex technical challenges.
* **Relevance to Enterprise Software Development:** The skills and knowledge gained are highly applicable to real-world enterprise software development scenarios.
* **Unique and Challenging:** This project is ambitious and unique, setting it apart from typical software development projects like a boring webapp 🤓
* **Personal Growth:** I will learn how to breakdown large problems, and learn new technologies by myself.
