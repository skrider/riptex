# Roadmap

## Day 1

Learn basic Go using learn go by example. Scaffold project, break ground on writing the REPL
- Figure out a good framework that is extensible. Some way of transforming async UI input into affecting application state 
- Figure out how to move the cursor within the buffer. Probably will be more difficult than just reading from stdin. How do CLI apps like vim and nano do it?
- How much state will be stored? SHould the app track the user's actions, or be completely stateless and only dependent on the current state of the equation?
- Should I make this object-oriented, create a Step class with a Intermediate subclass, or do everything functional? Leaning more towards functional but this depends a lot on Go's capabilitites. OO could be quite nice for extensibility.
- Undo will be a tough feature to implement. Key question: how much of bash / terminal emulator's text editing functionality am I going to use? Not sure if reading from stdin affords you the ability to programmatically modify the input, or move the cursor. 

## Day 2

Cursor movement within REPL, implement wolfram intermediate if time. This is the "mvp" featureset.
- Fetch library within go, or curl?
- Probably need some kind of config parser or environment var to hold the worlfram API Key
- Will also need to collect input and transform it to output. Probably just write it to stdout so that you can pipe it.

## Day 3

Break ground on diff/propagate. Probably want to parse latex and form some kind of AST, identify which AST nodes are the same in the next step, and change the ones that are not.
- Should diff/prop be done atomically, from one step to the next, with no additional context? Or should it act on the entire equation, with as much context as possible?
- Should I hard-code LaTeX constructs into the parser, or write one that sort-of-works based on a few general rules?

## Day 4

Continue working on diff/propagate


