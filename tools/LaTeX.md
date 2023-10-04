
$$
\int_{0}^{2\pi} \sin(x) \, dx = 0 =  e^{i\pi} + 1
$$


$$\frac{\Delta y}{\Delta x}$$

```tikz
\begin{document}
  \begin{tikzpicture}[domain=0:4]
    \draw[very thin,color=gray] (-0.1,-1.1) grid (3.9,3.9);
    \draw[->] (-0.2,0) -- (4.2,0) node[right] {$x$};
    \draw[->] (0,-1.2) -- (0,4.2) node[above] {$f(x)$};
    \draw[color=red]    plot (\x,\x)             node[right] {$f(x) =x$};
    \draw[color=blue]   plot (\x,{sin(\x r)})    node[right] {$f(x) = \sin x$};
    \draw[color=orange] plot (\x,{0.05*exp(\x)}) node[right] {$f(x) = \frac{1}{20} \mathrm e^x$};
  \end{tikzpicture}
\end{document}
```





$$ y = mx + b $$








#### propositional logic symbols 

If you are new to logic, you should know that logic involves a number of “special symbols”. To input these symbols using Markdown, you will need to insert a specific command between two dollar signs, e.g. `$\someCommand$`. Let’s look at one example. In propositional logic, there is what is called the caret, which looks like this: ∧

To get the caret to display, you will need to type

```
$\wedge$
```

Here are the rest of the symbols you will need:

|                         |                     |
| ----------------------- | ------------------- |
| **Propositional Logic** |                     |
| ¬                       | `$\neg$`            |
| →                       | `$\rightarrow$`     |
| ↔                       | `$\leftrightarrow$` |
| ∨                       | `$\vee$`            |
| ∧                       | `$\wedge$`          |
| ⊢                       | `$\vdash$`          |
| ⊣                       | `$\dashv$`          |
| **Predicate Logic**     |                     |
| ∀                       | `$\forall$`         |
| ∃                       | `$\exists$`         |
| ∈                       | `$\in$`             |
| ⊨                       | `$\models$`         |
| **Modal Logic**         |                     |
| □                       | `$\Box$`            |
| ◊                       | `$\Diamond$`        |

from http://www.davidagler.com/teaching/logic/handouts/supplemental_material/MarkdownForSymbolicLogic.html





###### links



- https://tex.stackexchange.com/questions/11/what-are-good-learning-resources-for-a-latex-beginner
- https://tex.stackexchange.com/questions/253077/how-do-you-create-a-set-in-latex
- https://github.com/artisticat1/obsidian-tikzjax/blob/main/README.md


---
[[2023-08-23]]


