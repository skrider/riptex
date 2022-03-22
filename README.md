# rIpTeX

I spend a lot of time writing equations for school, mainly for homework assignments. To get faster at this I switched to using LaTeX embedded inside markdown files, mainly to leverage the almighty power of snippets and ctrl-c ctrl-v. My workflow comes out to: type equation, copy and paste it onto next line, alter equation in such a way that is logically valid but gets me closer to my goal of solving it, repeat until solved. Example:

```latex
\begin{align*}
\frac{d}{dt} x(t) &= -2x(t) + 2u(t) \\
x(t) &= e^{at}x(0) + \int_0^t e^{a(t - \tau)}bu(\tau) d\tau \\
|x(t)| &= |e^{at}x(0) + \int_0^t e^{a(t - \tau)}bu(\tau) d\tau| \\
&\leq |e^{at}x(0)| + |\int_0^t e^{a(t - \tau)}bu(\tau) d\tau| \\
&\leq |e^{at}x(0)| + \int_0^t |e^{a(t - \tau)}bu(\tau)| d\tau \\
&= |e^{at}x(0)| + \int_0^t |be^{a(t - \tau)}||u(\tau)| d\tau \\
&\leq |\epsilon e^{at}| + |b \epsilon e^{at}|\int_0^t e^{-a\tau} d\tau \\
&= |\epsilon e^{at}| - |b \epsilon e^{at}|\frac{1}{a}(1 - e^{-at}) \\
&= \epsilon e^{-2t} + \epsilon (1 - e^{-2t}) \\
&= \epsilon \\
\end{align*}
```

To streamline this process even more, I want to write a tool that accelerates this process of copy-pasting and modifying lines, while providing the user with a helpful assortment of tricks that, while falling far short of a symbolic math manipulation engine, will no doubt speed things up.

rIpTeX is a time-independent REPL-like CLI application for LaTeX. In a rIpTeX session, the user will type in an equation and proceed to transform it until the desired result has been reached. Each "REP" cycle constitutes one step in this process. The user can "time travel" within their session and make upstream changes, such as a forgotten negative sign, or switching a variable name. rIpTeX will then use a heuristic to attempt to propagate the change to downstream steps, saving the user from having to manually make this change. When the user is done, rIpTeX will export the equation to plaintext format.

rIpTeX will have a modular architecture, enabling *future* features such as autocompletion, snippets, versatile exporting, potentially things like equation graphing and symbolic manipulation, and more to be added.
