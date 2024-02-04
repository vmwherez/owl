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



[[Wave Equation]]

