## Notes

### Parsing

Everyone who has ever programmed has probably heard about parsers, mostly by encountering a “parser error”. Or maybe heard or even said something like “we need to parse this”, “after it’s parsed”, “the parser blows up with this input”. The word “parser” is as common as “compiler”, “interpreter” and “programming language”.

But what is a parser exactly? What is its job and how does it do it?

A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure – giving a structural representation of the input, checking for correct syntax in the process. […] The parser is often preceded by a separate lexical analyser, which creates tokens from the sequence of input characters;

A parser turns its input into a data structure that represents the input. That sounds pretty abstract, so let me illustrate this with an example. Here is a little bit of JavaScript:

```js
const input = '{"name": "Bogdan", "age": 21}';
const obj = JSON.parse(input);

console.log(obj.name); // "Bogdan"
console.log(obj.age); // 21
```

Our input is just some text, a string. We then pass it to a parser hidden behind the JSON.parse function and receive an output value. This output is the data structure that represents the input: a JavaScript object with two fields named name and age, their values also corresponding to the input. We can now easily work with this data structure as demonstrated by accessing the name and age fields.

In most interpreters and compilers the data structure used for the internal representation of the source code is called a “syntax tree” or an “abstract syntax tree” (AST for short). The “abstract” is based on the fact that certain details visible in the source code are omitted in the AST. Semicolons, newlines, whitespace, comments, braces, bracket and parentheses – depending on the language and the parser these details are not represented in the AST, but merely guide the parser when constructing it.

A small example should make things clearer. Let’s say that we have the following source code:

```js
if (3 * 5 > 10) {
  return "hello";
} else {
  return  "goodbye";
}
```
And let’s say we are using JavaScript, have a MagicLexer, a MagicParser and the AST is built out of JavaScript objects :

```js
const input = 'if (3 * 5 > 10) { return "hello"; } else { return  "goodbye"; }';
const tokens = MagicLexer.parse(input);

console.log(tokens); 
// {
//   type: "if-statement",
//   condition: {
//     type: "operator-expression",
//     operator: ">",
//     left: {
//       type: "operator-expression",
//       operator: "*",
//       left: { type: "integer-literal", value: 3 },
//       right: { type: "integer-literal", value: 5 }
//     },
//     right: { type: "integer-literal", value: 10 }
//   },
//   consequence: {
//     type: "return-statement",
//     returnValue: { type: "string-literal", value: "hello" }
//   },
//   alternative: {
//     type: "return-statement",
//     returnValue: { type: "string-literal", value: "goodbye" }
//   }
// }
```

As you can see, the output of the parser, the AST, is pretty abstract: there are no parentheses, no semicolons and no braces. But it does represent the source code pretty accurately.


### Why not using a parser generator?
A parser generator is a tool that takes a formal grammar as input and produces source code for a parser that recognizes that grammar. Parser generators are often used to create parsers for programming languages, data formats, or other structured text.

The main reason is that I want to learn how to write a parser. I want to understand how it works, what the different parts are and how they fit together. I want to be able to write my own parser for my own language, and I want to be able to do it without relying on a tool that does it for me.

Full description of the EcmaScript syntax, in BNF https://tomcopeland.blogs.com/EcmaScript.html. A parser generator would take something like this and turn it into compilable C++ code.

Most people that recommend using a parser generator, when others want to get started with interpreters and compilers, only do so because they’ve written a parser themselves before. They’ve seen the problems and solutions available and decided it’s better to use an existing tool for the job. And they’re correct - when you want to get something done and are in a production environment, where correctness and robustness are priorities.

### Writing my own parser for NoBl

There are two main strategies when parsing a programming language: top-down parsing or bottom-up parsing. A lot of slightly different forms of each strategy exist. For example, “recursive descent parsing”, “Early parsing” or “predictive parsing” are all variations of top down parsing.

The parser that I am going to be writing today is a recursive descent parser. And in particular, it’s a “top down operator precedence” parser, sometimes called “Pratt parser”, after its inventor Vaughan Pratt. https://en.wikipedia.org/wiki/Operator-precedence_parser#Pratt_parsing https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html

A recursive descent parser is a top-down parser that processes input based on a set of recursive functions, where each function corresponds to a grammar rule. It parses the input from left to right, constructing a parse tree by matching the grammar’s production rules. This parser is simple to implement and is suitable for LL(1) grammars, where decisions can be made based on a single lookahead token. 

My parser won’t be the fastest of all time, I won’t have formal proof of its correctness and its error-recovery process and detection of erroneous syntax won’t be bullet proof. The last one is especially hard to get right without extensive study of the theory surrounding parsing. But what I'm going to have is a fully working parser for the NoBl programming language that’s open for extensions and improvements, easy to understand and a great start to further dive into the topic of parsing.



