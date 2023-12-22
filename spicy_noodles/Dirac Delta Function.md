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







```tikz
\begin{document}
\begin{tikzpicture}
% Axes
	\draw[->] (0,0) -- (6,0) node[right] {$t$};
	\draw[->] (0,-1.5) -- (0,1.5) node[above] {$x(t)$};
% Grid lines
\foreach \x in {1,2,3,4,5}
	\draw (\x,-0.1) -- (\x,0.1);
\foreach \y in {-1,-0.5,0.5,1}
 	\draw (-0.1,\y) -- (0.1,\y);
% Signal
	\draw[pink,thick,domain=0:5,samples=200] plot (\x,{sin(2*pi*\x r)});
\end{tikzpicture}
\end{document}
```

Links

- https://en.wikipedia.org/wiki/Dirac_delta_function

---


[[Heavyside Step Function]]