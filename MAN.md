# NAME

riptex - manipulate latex equations and statements quickly

# SYNOPSIS

riptex

riptex --diff [steps...]

riptex --file problem1.txt >> homework.md

# DESCRIPTION

rIpTeX is a time-independent REPL-like CLI application for LaTeX. In a rIpTeX session, the user will type in an equation and proceed to transform it until the desired result has been reached. Each "REP" cycle constitutes one step in this process. The user can "time travel" within their session and make upstream changes, such as a forgotten negative sign, or switching a variable name. rIpTeX will then use a heuristic to attempt to propagate the change to downstream steps, saving the user from having to manually make this change. When the user is done, rIpTeX will export the equation to plaintext format.

# OPTIONS

--file [file] : initializes the REPL with lines read from a file 

--diff [steps...] : does not open a REPL, only outputs the result of attempting to propogate diff(eq1, eq2) to eq3, eq4, etc

--nodiff : does not attempt to diff and propagate during the session

# DEFINITIONS 

Equation - the subject of the rIpTeX session

Step - one step of transforming the equation from point A to point B. A session will probably have several steps.

Intermediate - rather than pushing a new step, a user can run a special command

# COMMANDS

TAB - move cursor to next "point of significance" inside the step. Inspired by emacs package org-cdlatex-mode.

S-TAB - move cursor to previous "point of significance" inside the step

RET - finish editing current step and move to next step

C-RET - finish editing current step, move to next step, and process side effects, such as diffing and propagating or executing an intermediate

C-up - move to previous step

C-down - move to next step

C-backspace - delete current step

C-z - undo

C-S-z - redo

# INTERMEDIATES

It's important to note that intermediates are steps in and of themselves, and so are still taken into account while diffing.

"wolfram" - sends content of previous step to Wolfram Alpha API. Inserts response into next step.

"A == 3.14159 + 1" - evaluates a Go arithmatic expression and stores it temporarily in A

"substitute" - substitutes all previously set numeric values into the equation, outputs this as the next step. One powerful thing you can do is chain a "subsititute" intermediate into a "wolfram" intermediate, which will usually evaluate the equation.

# SYNTAX / RESERVED WORDS

!! - halts the propagate sequence at this step
