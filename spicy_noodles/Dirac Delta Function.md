```tikz
\begin{document}
\begin{tikzpicture} 
% Axis 
\draw[->] (-2,0) -- (2,0) node[right] {}; \draw[->] (0,-0.5) -- (0,1.5) node[above] {}; 
% Dirac Delta arrow 
\draw[pink, thick, -stealth] (0,0) -- (0,1) node[above, blue] {}; 
% Label for t0 
\draw[dashed] (0,0) node[below] {}; \end{tikzpicture}
\end{document}
```

*Impulse Function*

$$
\delta(t - t_0)
$$

$$
\int_{-\infty}^{\infty} \delta(t - t_0) \, dt = 1
$$






Links

- https://en.wikipedia.org/wiki/Dirac_delta_function

---


[[Heavyside Step Function]]