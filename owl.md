# 🦉

# `2023-10-04`

`11:02a` **`git rm <path-to-submodule>`**

`10:51a` Organizing this repo to run against `peregrine`. Working on UI in `peregrine.`

`7:15a` [[crisis_in_science]]

- https://news.ycombinator.com/item?id=37756570
- https://www.wsj.com/articles/how-preapproved-narratives-corrupt-science-false-studies-covid-climate-change-5bee0844
## `2023-10-03`


#### `2:07p`

Putting timed entries here will make this more like a feed. 

> What about converting this markdown into JSON into reddit or something.
#### `12:26p`

- `notes` cd into notes and vim
- `peregrine` take me to peregrine app folder
#### `10:12a`

- https://zod.dev/?id=basic-usage
- https://create-react-app.dev/docs/adding-typescript/
- https://headlessui.com 
- https://github.com/typescript-cheatsheets/react#reacttypescript-cheatsheets
- https://react-typescript-cheatsheet.netlify.app/docs/basic/setup
- https://create-react-app.dev/docs/adding-typescript/

`9:36a`

```
✗ npx tailwindcss init


Created Tailwind CSS config file: tailwind.config.js
```


## `2023-10-02`





#### react & go

`~12p` 

- https://learngolangonline.com/posts/golang-roadmap#resources-to-learn-gogolang

![[golang_concept_map.png]]

#### tikz in jupyter

`9:07a` *Getting `tikz `into Jupyter*
- https://dvisvgm.de/Downloads/

`9:49a` Kind of fun going down the rabbit hole with TeX here in Norlan Stacks.

https://www.ctan.org/pkg/kpathsea

```
(base) ➜  dvisvgm-3.1.1 which dvisvgm
/Library/TeX/texbin/dvisvgm
```

I didn't need to build again from source, I vaguely recall these feels when I was installing MacTex, which I remember being huge.
	- https://www.latex-project.org/get/ 
	- https://en.wikipedia.org/wiki/TeX_Live

So, this tool was already installed. 

`10:32a` I looked at the source code for a *texmagic* and saw it was using `convert` which was an alias for *ImageMagick*... 
So of all things, ended up installing that on the 2015 here.

- https://imagemagick.org/script/download.php#macosx

```
export MAGICK_HOME="$HOME/ImageMagick-7.0.10"
export PATH="$MAGICK_HOME/bin:$PATH"
export DYLD_LIBRARY_PATH="$MAGICK_HOME/lib/"
```

That involved allowing the libraries to run in System Preferences a bunch. 

`11:05` Went back to https://github.com/mkrphys/ipython-tikzmagic and http://cityinthesky.co.uk/opensource/pdf2svg/

I had deleted pdf2svg and replaced with an alias to dvisgm.

- https://tobywf.com/2018/01/beautiful-svg-diagrams-with-tikz-dvisvgm/
- https://laurentperrinet.github.io/sciblog/posts/2015-01-20-using-tikzmagic.html
#latex #tikz #jupyter


##  `2023-10-01`

`12:34p`

*Dogs.*

`12:24p`

`xxd | most a.out`
## `12p`

```sh
$ r2 /bin/ls   # open file in read-only
> aaa          # analyse the program (r2 -A)
> afl          # list all functions (try aflt, aflm)
> px 32        # print 32 byte hexdump current block
> s sym.main   # seek to main (using flag name)
> f~foo        # filter flags matching 'foo' (internal |grep)
> iS;is        # list sections and symbols (rabin2 -Ss)
> pdf; agf     # disassembly and ascii-art function graph
> oo+;w hello  # reopen in read-write and write a string
> ?*~...       # interactive filter in all command help
> q            # quit
```

```
[0x100003e90]> pd
```

*Working from iPad: radare2, cmake, vim*

`09:23` http://concurrencyfreaks.blogspot.com/2023/09/50-years-later-is-two-phase-locking.html
# `2023-09-30`


`3:43p` lot's done today: 

- DSP: wavelet transform
- C++: pointers vs references, heap vs stack
- ML:  Hubel, D. H.; Wiesel, T. N. (1959). "Receptive fields of single neurones in the cat's striate cortex." _Journal of Physiology_. 148: 574–591. 
	- Recounting the impact of Hubel and Wiesel https://www.ncbi.nlm.nih.gov/pmc/articles/PMC2718241/
	-  *Cognitron.* Fukushima https://en.wikipedia.org/wiki/Neocognitron
	- https://ai.meta.com/blog/ai-speech-brain-activity/ Jean Remi King
	- https://www.thelancet.com/journals/laneur/article/PIIS1474-4422(21)00427-0/fulltext
- CS1 citation style 

#### DSP

`~9a` I had forgotten wavelets video this morning. It was a review when I heard "convolution" in this video: 

https://www.youtube.com/watch?v=jnxqHcObNK4&t=213s - Wavelets: a mathematical microscope


### machine learning 

`~8a` https://venturebeat.com/business/todays-ai-is-not-science-its-alchemy-what-that-means-and-why-that-matters-the-ai-beat/

```
epistemology

ĭ-pĭs″tə-mŏl′ə-jē

### noun

1. The branch of philosophy that examines the nature of knowledge, its presuppositions and foundations, and its extent and validity.
2. The theory of cognition; that branch of logic which undertakes to explain how knowledge is possible. Probably first used by Ferrier.
3. The theory or science of the method or grounds of knowledge.
```

https://www.youtube.com/watch?v=gG5NCkMerHU - The Epistemology of Deep Learning - Yann LeCun

`9:01a` ConvNets (CNNs?) were inspired by the visual cortex. Hubel & Wiesel, 1959*

https://www.ncbi.nlm.nih.gov/pmc/articles/PMC2718241/ - Recounting the impact of Hubel and Wiesel. 

`3:17p` https://ai.meta.com/blog/ai-speech-brain-activity/  - Stéphane d'Ascoli

### C++
#### pointers

`5:23a` Daves Garage - mastering pointers

https://www.youtube.com/watch?v=DTxHyVn0ODg The Cherno - pointers

uninitialized a null are not the same kind of pointer
void pointer? 
function pointer?

```c++
int x = 10;
int* ptr = &x; // Pointer to int
*ptr = 20; // Modify the value of x indirectly through the pointer

```

- Pointers are variables that store the memory address of another variable.
- They can be reassigned to point to different variables or memory locations.
- Pointers can be null (not pointing to anything) or uninitialized.
- They require dereferencing with the `*` operator to access the value they point to.
###### uninitialized pointer

```c++

int* uninitializedPtr; // Declare an uninitialized pointer 
// uninitializedPtr is not pointing to a valid memory location 
// Accessing uninitializedPtr before assigning a value would result in undefined behavior


```

###### null pointer

```c++
int* nullPtr = nullptr; // Declare a null pointer 
// nullPtr is explicitly set to point to nothing, represented by nullptr 
// It's safe to check and use nullPtr without accessing memory
```

###### void pointer

```c++
void* voidPtr; // Declare a void pointer
int x = 10;
voidPtr = &x; // Assign the address of an int to the void pointer
// voidPtr can point to objects of any type, but you need to cast it to access the value
int* intPtr = static_cast<int*>(voidPtr);
std::cout << *intPtr << std::endl; // Accessing the int value through casting

```

###### function pointer

```c++

// Define a function that takes two integers and returns their sum
int Add(int a, int b) {
    return a + b;
}

// Declare a function pointer and point it to the Add function
int (*functionPtr)(int, int) = Add;

// Call the function through the pointer
int result = functionPtr(5, 3);
std::cout << "Result: " << result << std::endl;

```


*Dereferencing?* 

`memset(pointer_name, 0)`

`delete[]`

*Endianness?*

#### references

```c++
int x = 10;
int& ref = x; // Reference to int
ref = 20; // Modify the value of x directly through the reference
```

- References are essentially aliases to existing variables.
- They must be initialized when declared, and once bound to a variable, they cannot be rebound to another variable.
- References are automatically dereferenced; you access the value directly.
- They cannot be null or uninitialized.

https://www.youtube.com/watch?v=IzoFn3dfsPA
*Pretty much the same thing for the computer, usually. Syntactic sugar.*

- **References are automatically dereferenced; you access the value directly.**

References are easier to read. References can't be null or uninitialized. 

#### reference vs pointer

- **Mutability:** Pointers can be reassigned to point to different variables, making them more flexible in terms of what they can reference. References are fixed once initialized, providing safety but less flexibility.
    
- **Initialization:** Pointers can be null or uninitialized, but references must be initialized when declared.
    
- **Dereferencing:** Pointers require explicit dereferencing with `*` to access the value, while references directly access the value.
    
- **Syntax:** Pointers use `*` to declare and dereference, while references use `&` when declared.
    
- **Nullability:** Pointers can be set to null (nullptr) to indicate they are not pointing to any valid memory location, while references always refer to an existing variable.

#### stack vs heap

`6:04a` Stack vs Heap? https://www.youtube.com/watch?v=wJ1L2nSIV1s

###### stack allocation:
###### heap allocation: 

---
REFERENCES & CITATIONS 
1. Hubel, D. H.; Wiesel, T. N. (1959). "Receptive fields of single neurones in the cat's striate cortex." _Journal of Physiology_. 148: 574–591.
2. Last, F. MI.; Others, A. (2023). "This would be the title of the reference." *Research Publishing Today.* 0: 123. 
3. Chicken, C. C.; Chickens. "Chicken, chick chick chicken chicks." *Chicken: Chickens Chicken*. 67 


# `2023-09-29`


`12:07a` https://www.educative.io/blog/cpp-memory-management

#### C++

`8:04a` 

```
**Key Differences:**

1. **C++ supports Object-Oriented Programming (OOP):** C++ extends C by adding features like classes, objects, inheritance, and polymorphism, making it suitable for OOP.
    
2. **Standard Template Library (STL):** C++ includes the STL, which provides a collection of template classes and functions for common data structures and algorithms.
    
3. **Namespace:** C++ introduces the concept of namespaces to prevent naming conflicts, allowing you to organize code better.
    
4. **Function Overloading:** C++ allows you to define multiple functions with the same name but different parameter lists, which is not possible in C.
    
5. **Exception Handling:** C++ provides try-catch blocks for handling exceptions, while C relies on error codes and manual error checking.
    
6. **New and Delete Operators:** C++ introduces `new` and `delete` operators for dynamic memory allocation and deallocation, whereas C uses `malloc` and `free`.
    
7. **I/O Streams:** C++ replaces C's `printf` and `scanf` with stream-based I/O, making it more versatile and type-safe.
    
8. **Type Safety:** C++ is generally more type-safe than C, thanks to features like strong typing, function prototypes, and better type checking.
    
9. **Header Files:** In C, header files typically have a ".h" extension, while C++ often omits the extension and uses ".hpp" or ".h" for templates.
```

- https://cmake.org/cmake/help/latest/guide/tutorial/index.html
- https://www.freecodecamp.org/news/learn-modern-cpp-by-building-an-audio-plugin/
- https://www.educative.io/blog/cpp-memory-management
- https://www.educative.io/blog/intermediate-cpp-tutorial
###### peregrine 

https://antonz.org/sql-upsert/
https://www.etymonline.com/columns/post/who-lusts-for-certainty-lusts-for-lies

`9a` Eric S. Raymond's catb.org (Cathedral and The Bazar)

- http://www.catb.org/esr/structure-packing/
- http://www.catb.org/esr/time-programming/
- 

`9:53a` https://www.youtube.com/watch?v=habHK6wLkic how was the Natural Logarithm discovered?
## `2023-09-28`



$$F(n) = F(n-1) + F(n-2)$$


`9:39p`

```c


#include <stdio.h>

int fibonacci(int n) {
    if (n <= 1) {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}




int pointers() {
    int num = 10; // Declare an integer variable
    int *ptr;    // Declare a pointer variable

    ptr = &num;  // Assign the address of 'num' to the pointer

    printf("Value of num: %d\n", num);
    printf("Address of num: %p\n", &num);
    printf("Value of ptr: %p\n", ptr);
    printf("Value pointed by ptr: %d\n", *ptr);

    *ptr = 20;   // Change the value of 'num' indirectly through the pointer

    printf("New value of num: %d\n", num);

    return 0;
}


int main() {
    int n = 10; // Change this value to calculate a different number of Fibonacci sequence terms

    printf("Fibonacci Sequence up to %d terms:\n", n);
    for (int i = 0; i < n; i++) {
        printf("%d ", fibonacci(i));
    }
pointers();
    
    return 0;
}


/*  


*/




```

`7:41p` Documenting break in routine [[../2023-09-28]]. At McDonalds.

This system had been working well with Math topics, LaTeX and Tikz. 


Cody can write in Markdown:

```c
// Cody, could you write a demo 'hello world' in C?
#include <stdio.h>

int main() {
  printf("Hello, world!\n");
  return 0;

  // Cody, how about a function which will print the Fibonacci sequence?

  int fib(int n) {
    if (n == 0) {
      return 0;
      } else if (n == 1) {
        return 1;
        } else {
          return fib(n - 1) + fib(n - 2);

// Cody, please write a function which demonstrates a linked list
int linked_list_demo( ) {

}

//Cody, please write a function which demonstrates a binary tree
int binary_tree_demo( ) {}

  // Cody, please write a function which demonstrates a binary search tree

  // Cody, please write a function which demonstrates a heap

  // Cody, please write a function which demonstrates a graph

// Cody, please write a function which demonstrates a queue

  void graph_demo( ) {}




```

## `2023-09-25`

- https://major.io/p/quadlets-replace-docker-compose/

## `2023-09-24`

`10:34a` https://cstheory.stackexchange.com/questions/53343/recent-advances-in-computer-science-since-2010

`4:22a` https://en.wikipedia.org/wiki/Faraday%27s_law_of_induction

https://en.wikipedia.org/wiki/Kirchhoff%27s_circuit_laws

https://en.wikipedia.org/wiki/Kerckhoffs%27s_principle

_Different things, both relevant, starting yesterday afternoon or last night._

https://www.petitcolas.net/kerckhoffs/crypto_militaire_2.pdf

`1:55p`

https://en.wikipedia.org/wiki/Gustav_Kirchhoff
-> https://en.wikipedia.org/wiki/Zeeman_effect
-> https://en.wikipedia.org/wiki/Lorentz_ether_theory

`2023-09-23`

Worked with GPT in French.

## `2023-09-21`

`12:15a`

###### convolution

I don't really understand it yet:

$$ a * b $$

[But what is a convolution? 3blue1brown](https://www.youtube.com/watch?v=IaSGqQa5O-M)

Some new vocabulary in here that I'm enjoying. He times the `np.convolve(arr1, arr2)` function from Numpy against FFTconvolve in Scipy.

> "Expanding a polynomial and collecting all like terms is the same thing as a convolution."

$$ O(N^2) $$

_Linear system of equations..._

FFT ...

Roots of unity...

https://www.youtube.com/watch?v=IaSGqQa5O-M
_Continuous convolution._

Probability: Normal distribution.

`1:07a` Central Limit Theorem https://mathoverflow.net/questions/182752/central-limit-theorem-via-maximal-entropy

`1:27a` Gaussian...

`2a` Gauss discovered an algo for FFT even before Fourier himself, but it was published in 19th century Latin in his collected works.

https://www.youtube.com/watch?v=h7apO7q16V0 # The Fast Fourier Transform (FFT): Most Ingenious Algorithm Ever?

What's interesting here is groking the multiplication of polynomials being related to Fast Fourier Transform being related to convolution...

Due to harmonics only the first parts of the signal are needed for FFT.

Interpolation is the inverse DFT matrix...

`7a` Probability density function. #statistics

The [standard normal distribution](https://en.wikipedia.org/wiki/Normal_distribution#Standard_normal_distribution "Normal distribution") has probability density:

$$ f(x) = \frac{1}{2\pi}e^{-x^2/2} $$

https://en.wikipedia.org/wiki/Eigenvalues_and_eigenvectors

`8:34a`

###### container networking

https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking
https://developers.redhat.com/blog/2019/05/17/an-introduction-to-linux-virtual-interfaces-tunnels
https://iximiuz.com/en/posts/container-networking-is-simple/ #network-engineering

http://peter.eisentraut.org/blog/2023/09/20/grouping-digits-in-sql #sql I feel like there were some other SQL notes saved recently.

`10:33a`

`11:01a` https://en.wikipedia.org/wiki/Discrete_cosine_transform

https://mathshistory.st-andrews.ac.uk/Biographies/Cauchy/

`11:38a` https://jakobs.dev/solving-mnist-with-gzip/

## `2023-09-18`

`9:18` 4423

Making some progress classifying Signals and Systems math. 

```
transfer function
unit step
Laplace
Fourier
Mutual Inductance
```

`2:43p`

###### Cardinal Sine (sinc)

$$ sinc(x) = \frac{sin(x)}{x} $$

`2:24p`

[Unit Impulse Signal (quick video)](https://www.youtube.com/watch?v=suFh1FTETRI)

`2:02p`

- Diraq Delta function is related to and nearly the same as the _impulse function_.
- Heavyside function is related to _unit step function_.

`1:57p`

```
What are some applications in DSP of these functions?

ChatGPT

Digital signal processing (DSP) relies on various mathematical functions, including step functions and impulse functions, for a wide range of applications. Here are some common applications of these functions in DSP:

**Impulse Response Analysis**: Impulse functions are used to analyze the impulse response of linear time-invariant (LTI) systems. By convolving the impulse response with input signals, you can predict the system's output for any input, making it essential for system modeling and filtering in DSP.

**Convolution**: Both step and impulse functions play a crucial role in convolution operations. Convolution is used for various DSP tasks like filtering, image processing, and speech recognition. Step functions are often used to model unit sample sequences.

**Filter Design**: Step functions help in designing and analyzing filters. For example, a step input can be used to determine the step response of a filter, providing insights into its characteristics, such as settling time and overshoot.

**System Analysis**: Step responses are used to analyze system behavior, including settling time, stability, and frequency response. Step inputs can reveal a system's response to sudden changes, which is useful in control systems and communications.

**Signal Generation**: Impulse functions are used to create signals with specific frequency content. Fourier analysis can be applied to impulse responses to understand signal frequency components, enabling the generation of signals with desired properties.

**Signal Reconstruction**: In applications like audio and image compression, step and impulse functions are used to understand the effects of quantization and reconstruction. This aids in maintaining signal quality during compression and decompression processes.

**Control Systems**: Step responses are used to analyze and design control systems. Understanding how a system responds to step inputs is critical for stability and performance analysis in control engineering.

**Sampling and Reconstruction**: In digital communication and audio processing, step functions are used to understand the effects of sampling and reconstruction. This knowledge helps ensure accurate data transmission and high-quality audio playback.

**Digital Filters**: Impulse functions are fundamental in designing digital filters, such as finite impulse response (FIR) and infinite impulse response (IIR) filters. Impulse responses define the filter's behavior and frequency response.

These applications highlight the significance of step and impulse functions in DSP for tasks ranging from system analysis and signal processing to control systems and communication. They provide a foundation for understanding, analyzing, and manipulating digital signals in various domains.
```

###### Heavyside function (H(x)) vs Step Function (u(t))

$$ H(x) \ \ u(t) $$

```
The Heaviside function and the step function are closely related mathematical functions used in different contexts, but they serve similar purposes. Here's the key difference between them:

1. **Step Function (u(t))**: The step function, often denoted as u(t), is a continuous function defined as follows:

    u(t) = 0 for t < 0 u(t) = 1 for t ≥ 0

    In other words, the step function is zero for negative values of 't' and one for non-negative values of 't'. It represents an instantaneous jump from 0 to 1 at t = 0 and is used to model sudden changes or switches in a system.

2. **Heaviside Function (H(t))**: The Heaviside function, denoted as H(t), is closely related to the step function. It is defined as follows:

    H(t) = 0 for t < 0 H(t) = 1/2 for t = 0 H(t) = 1 for t > 0

    Unlike the step function, the Heaviside function has a value of 1/2 at t = 0, creating a smoother transition at that point. It is often used in engineering and physics to model systems with more gradual responses.


In summary, the main difference is that the step function makes an abrupt transition from 0 to 1 at t = 0, while the Heaviside function has a continuous, but smooth, transition through 1/2 at t = 0. The choice between them depends on the specific modeling requirements of the problem at hand.
```

###### LTI, FIR, IIR

- linear time-invariant
- finite impulse response
- infinite impulse response

###### Dirac delta vs. impulse function

```

The terms "Dirac delta" and "impulse function" are often used interchangeably, but they can have slightly different interpretations depending on the context. Let's clarify the differences:

1. **Dirac Delta (δ(t))**: The Dirac delta, often denoted as δ(t), is a mathematical concept used in the context of distributions or generalized functions. It's not a function in the traditional sense, but rather a distribution. The Dirac delta is defined as follows:

    ∫[from -∞ to +∞] δ(t) dt = 1 δ(t) = 0 for t ≠ 0

    The Dirac delta is concentrated at t = 0 and has an area under its curve equal to 1. It's used to model an idealized impulse or a point source in mathematics and physics. It's crucial in the theory of signals and systems, especially when dealing with continuous-time signals.

2. **Impulse Function (or Impulse Signal)**: The term "impulse function" is more general and can refer to both continuous and discrete cases. In the continuous domain, it often refers to the Dirac delta function. However, in the discrete domain, it represents a sequence of discrete impulses. For example:

    Impulse function in the discrete domain: δ[n] = 1 for n = 0 δ[n] = 0 for n ≠ 0

    This discrete impulse function is used in discrete signal processing and serves a similar role as the Dirac delta but in the discrete-time context.


In summary, the Dirac delta is a specific mathematical distribution used to model impulses in continuous-time signals, while the term "impulse function" can refer to either the Dirac delta in the continuous case or discrete impulses in the discrete case, depending on the context.
```

###### the integral of the impulse function is the step function

```
From its definition it follows that the integral of the impulse function is just the step function:

δ(t)dt = u(t)

Thus, defining the derivative of the unit step function as the impulse function is justified.
```

`1:41p`

```
\delta(t) = \frac{d}{dt} H(t)
```

$$ \delta(t) = \frac{d}{dt} H(t) $$

`1:35p`

https://en.wikipedia.org/wiki/Heaviside_step_function

```tikz
\begin{document}
\begin{tikzpicture}
	% Axes
	\draw[->] (-2,0) -- (2,0) node[right] {$t$};
	\draw[->] (0,-0.5) -- (0,1.5) node[above] {$H(t)$};
	% Horizontal line
	\draw[blue,thick] (-2,0) -- (0,0);
	\draw[blue,thick] (0,1) -- (2,1);
	% Vertical line at t=0
	\draw[red,thick,dashed] (0,-0.2) -- (0,1.2);
	% Labels
	\node[red] at (0.2,-0.2) {$0$}; \node[blue] at (1.5,1.2) {$H(t)$};
\end{tikzpicture}
\end{document}
```

$$ H(t) := \int\_{-\infty}^{t} \delta(\tau) d\tau $$

`12:38p`

Hull function?


https://mu.krj.st/wave/

`8:06a`

##### "emergence in complex systems"

https://www.ncbi.nlm.nih.gov/pmc/articles/PMC6207336/pdf/pone.0206489.pdf

> Soy Yo

**Granger Causality** https://en.wikipedia.org/wiki/Granger_causality

`8:38a`

###### impulse function, LaTeX, signals, DSP and music... электротехника

$$\int_{∞}^{−∞} 𝛿(𝑡)d𝑡=1 $$

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
	\draw[blue,thick,domain=0:5,samples=200] plot (\x,{sin(2*pi*\x r)});
\end{tikzpicture}
\end{document}
```

$$ x(t) = \sin(2\pi t) $$

$$ sinc = sin(x)/x $$

`~9:45a` _convolution._ What does it mean in EE and AI?

https://mu.krj.st/wave/ this was a great little site Austin found.

In this article it gets into WAV format and 1-bit music or whatever it is.

`10:20a` Who are the authors What are the titles of these various books and who is the author?

_Designing Audio Effect Plugins in C++_ - Will C. Pirkle
_Circuit Theory_ - Wikibooks
_Signals and Systems_ - Wikibooks
_Signals and Systems_ - LibreTexts
_Book 10_ - Basic Books in Science (electromotive things)

- Faraday discovered that the changing flux of the magnetic field, passing through a wire loop, could produce an electric field...

Music specific titles:
_THE ART OF VA FILTER DESIGN_ - Vadim Zavalishin
_The Technique of Electronic Music_ - Thomas Wells
The Fourier series for the sawtooth-wave function is ...?
_Modern Recording Techniques_ - David Miles Huber

- Transmission Loss $ TL = 14.5logM + 23 $
  _Music: A Mathematical Offering_ - Dave Benson
- Fletcher–Munson curves, or isophons.

_Yamaha Sound Reinforcement Handbook_ -

- For example, the wavelength of a 100 Hz tone is: 1130 flJsec. + 100 cycles/sec. =11.3 ft..Icycle
- What is the **Absorption Coefficient?**

_Otto's DIY cookbook_

```
Lesson 1: Op-Amp Basics Prerequisites

- Understanding the concept of voltages and currents.

- Understanding resistors and the usage of Ohm’s Law.

- Practical understanding of Kirchhoff's circuit laws.

- Understanding the definition of input and output impedance.
```

`10:41a` Maybe a bit of a spiral back to fundamentals:

- unit circle trig
- engineering maths first aid
- _Calculus_ - LibreTexts - there was something about the 'textbooky' graphics here which helped me move past traumas

The `spicy_noodles` folder is really messy compared to the cohesion I'm experiencing this morning around DSP, which is going better than I expected. Some amount of ambiguity is happening re: process or personal reflection. Need to edit.

#### sinc(x) into 'mexican hat'

https://www.allaboutcircuits.com/technical-articles/what-is-the-sinc-function-and-why-is-it-important-in-electrical-engineering/

`2023-09-17`

- https://blog.lancedb.com/llms-rag-the-missing-storage-layer-for-ai-28ded35fa984

## `2023-09-15`

Where am I at with DSP prerequisites?

## `2023-09-12`

`8:39a`
`8:49a` Quadratic Equations: `x` at 0 on the graph is where

$$
ax+bx^2+c = 0
$$

comes from.

## `2023-09-11`

> Editing from phone.

`7:20a09-07`

Quadratic Equations.

$$
ax + bx^2 + c
$$

https://en.m.wikipedia.org/wiki/Completing_the_square

Want to maybe add this to the historical math notes. [[2023-09-11]] I commented quite a bit on process and history. I should probably do a bit to copy into LaTeX here.

`11a` [GCSE Maths - How to Factorise Quadratics by Completing the Square (Part 1 - Factorise When a=1) #52](https://www.youtube.com/watch?v=ov3r4hXBU98 "GCSE Maths - How to Factorise Quadratics by Completing the Square (Part 1 - Factorise When a=1) #52")

‘5:13’ looking at commutative property of multiplication led to discovery of ‘math in Victorian England’. There was hardly Math in that time until a dramatic increase in the later period.

`6:49p` https://www.altdevarts.com/p/onboarding-floating-point

## `2023-09-07`

`9:39a` https://en.wikipedia.org/wiki/Differentiation_rules starting to get the hang of these.

- https://tutorial.math.lamar.edu/Classes/CalcI/DiffTrigFcns.aspx

#### essence of multivariable calculus

https://youtu.be/Jk_k3q9RoMU Essence of Multivariable Calculus

`10a` Getting to the end of the above video, which is reading ahead or forward on a learning spiral. I've been listening passively.

> Navier-Stokes theorem, etc... they are all 'the same thing' https://en.wikipedia.org/wiki/Navier–Stokes_equations

`12:30p`

## `2023-09-06`

`7:30a` https://www.ezzeriesa.com/index/Writing%20efa7772a3ae64a8193c7ef981ac619fc/How%20to%20level%20up%20beyond%20ETLs%200987988358654751a8538a05bde18f96.html

`11:26a` Doing well with GoodNotes, Khan and GPT: Chain Rule, Product Rule.

`1:37p` [[2023-09-06]] this is a link to more reflective journaling, which now lives separately from the Owl's technical notes. Noted here because I've written a lot there today, but also made strides on this end.

#### chain rule, product rule.

- https://www.khanacademy.org/math/old-ap-calculus-ab/ab-derivative-rules/ab-basic-diff-rules/a/basic-differentiation-review
- https://youtu.be/WxTrxxW0qeM Product Rule example, Khan Academy
- https://www.khanacademy.org/math/ap-calculus-ab/ab-differentiation-2-new/ab-3-1a/v/chain-rule-introduction

##### reviewing `2023` notebook in GoodNotes

- https://en.wikipedia.org/wiki/Helmholtz_machine
- http://www.cs.utoronto.ca/~hinton/helmholtz.html

#### people in AI

- https://en.wikipedia.org/wiki/Geoffrey_Hinton
- https://en.wikipedia.org/wiki/Peter_Norvig
- https://en.wikipedia.org/wiki/Stuart_J._Russell
- https://en.wikipedia.org/wiki/Andrej_Karpathy
- https://martin.kleppmann.com

#### chip huyen

> runs a discord server:

https://huyenchip.com

https://en.wikipedia.org/wiki/Artificial_intelligence#knowledge-representation

https://docs.google.com/document/d/10K3pYTNvreVy5hl2EqWf_LX3mMW4CQw1TdMrHplMu00/edit#heading=h.ygqitxw2cpfv

https://github.com/chiphuyen/dmls-book/blob/main/basic-ml-review.md

`2:22p` https://www.recurse.com/about

https://www.ucl.ac.uk/gatsby/

> The Gatsby Computational Neuroscience Unit is a research centre at UCL supported by the Gatsby Charitable Foundation. Our work encompasses theoretical and computational neuroscience, computational statistics, machine learning and AI, threads drawn together by our focus on the mathematical foundations of adaptive intelligent behaviour.

`2:55p` convo with ChatGPT about how to think about Algorithms, which has been productive.

- https://en.wikipedia.org/wiki/Metric_space
- https://en.wikipedia.org/wiki/Cosine_similarity

https://en.wikipedia.org/wiki/Gensim

> **Gensim** is an [open-source](https://en.wikipedia.org/wiki/Open-source_software "Open-source software") library for unsupervised [topic modeling](https://en.wikipedia.org/wiki/Topic_model "Topic model"), [document indexing](https://en.wikipedia.org/wiki/Information_retrieval "Information retrieval"), retrieval by similarity, and other [natural language processing](https://en.wikipedia.org/wiki/Natural_language_processing "Natural language processing") functionalities, using modern statistical [machine learning](https://en.wikipedia.org/wiki/Machine_learning "Machine learning"). Gensim is implemented in [Python](<https://en.wikipedia.org/wiki/Python_(programming_language)> "Python (programming language)") and [Cython](https://en.wikipedia.org/wiki/Cython "Cython") for performance. Gensim is designed to handle large text collections using data streaming and incremental online algorithms, which differentiates it from most other machine learning software packages that target only in-memory processing.

`2:59p` Gensim demo in `peregrine` ...?

## `2023-09-05`

`1p`

- https://stackoverflow.com/questions/35534766/how-to-create-submodule-in-existing-repo

`12:51p`

- https://dev.to/chefgs/deploy-kubernetes-resources-in-minikube-cluster-using-terraform-1p8o

## `2023-09-04`

`4:14p` The determinant: result of transformation of a vector.

https://www.youtube.com/watch?v=Ip3X9LOh2dk&list=PL0-GT3co4r2y2YErbmuJw2L5tW4Ew2O5B&index=7

$$
\det(A) = ad - bc
$$

`4:28p`

- https://bvanderlei.github.io/jupyter-guide-to-linear-algebra/Applications_IP.html

## `2023-09-02`

`10a`

https://avestura.dev/blog/explaining-the-postgres-meme

https://akashhamirwasia.com/blog/internals-of-async-await-in-javascript/

## `2023-09-01`

https://www.cs.cmu.edu/Groups/AI/util/html/faqs/ai/neural/faq.html

`9:49a` https://automatetheboringstuff.com/chapter12/ `openpyxl`

`8:30a` https://github.com/karpathy/nanoGPT

`7a`

> Moon Enters Aries - Time to begin new projects. Exercise patience and selflessness where possible. Look out for new opportunities.

`12a`

## `2023-08-31`

### 🌚

_blue moon_

`11:30p` https://status.openai.com

`11p` Playing with MacOS grapher... for the first time that it felt relevant. Maybe the second. This time well enough to keep it in the dock as I am dealing with functions a lot lately.

I think N would enjoy the openAI status updates.

`10:30p`

###### squeeze theorem

$$

\lim_{{x \to 0}} \frac{\sin(x)}{x} = 1


$$

$$ \cos(x) \leq \sin(x)/x \leq 1 $$

- https://en.wikipedia.org/wiki/Squeeze_theorem
- https://web.math.ucsb.edu/~jcs/SqueezeTheorem.pdf

`10p`

##### instantaneous rate of change: the derivative

https://math.stackexchange.com/questions/2665752/what-does-approach-zero-really-mean

###### $\delta t$

\*means the change in time, or any other variable being considered.\*\*

Then $\delta t\to0$ simply means that you are forcing two points in time closer and closer together. This is an extremely important concept because it
forms the fundamentals of calculus.

The gradient of a line is given by 1$$\frac{\delta y}{\delta x}=\frac{y_2-y_1}{x_2-x_1}$$ and we can apply this to curves as well. We have $$\frac{f(x)-f(a)}{x-a}$$ at some point $x=a$ and this is known as the average rate of change. But to find the rate of change (or slope) exactly at the point $(a,f(a))$, we let $x=a+t$ and take the _limit_ as $t\to0$. $$\lim_{t\to0}\frac{f(x)-f(a)}{x-a}=\lim_{t\to0}\frac{f(a+t)-f(a)}{t}$$ This is known as the **instantaneous rate of change**, or more commonly, the _derivative_.

###### example

Take the function $f(x)=\sin x$ and we want to find its slope at $x=\pi$.

Then the definition of the derivative is $$\begin{align}f'(\pi)=\lim_{t\to0}\frac{f(\pi+t)-f(\pi)}{t}&amp;=\lim_{t\to0}\frac{\sin(\pi+t)-\sin\pi}{t}\\&amp;=\lim_{t\to0}\frac{\sin\pi\cos t+\sin t\cos\pi}{t}\\&amp;=-\lim_{t\to0}\frac{\sin t}{t}\end{align}$$ which is a very famous limit and it evaluates to $1$ (see the squeeze theorem)

Hence: $f'(\pi)=-1$.

##### SQL JOIN

`10p` Reviewing SQL joins:

[SQL Joins Explained](https://www.youtube.com/watch?v=9yeOJ0ZMUYw)

https://www.youtube.com/watch?v=Jh_pvk48jHA&t=200s examples

- INNER
- LEFT
- RIGHT
- FULL (Postgres only, not MySQL or SQLite)

```sql
SELECT *
FROM table1
INNER JOIN base
ON table1.column_name = table2.column_name


```

##### limits

`9p`

```tikz
\begin{document} \begin{tikzpicture} \draw[->] (-2,0) -- (2,0) node[right] {$x$}; \draw[->] (0,-1) -- (0,2) node[above] {$f(x)$}; \draw[dashed] (0.5,0) -- (0.5,1.5) node[above] {$0$}; \draw[domain=-1.8:1.8,smooth,variable=\x,blue] plot ({\x},{0.5*(\x)^2 + 0.5}); \draw[dotted] (0.5,0.5) -- (0,0.5) node[left] {$L$}; \draw[->] (1,-0.3) -- (1.8,-0.9) node[midway,above,sloped] {$f(x)$ approaching $L$ as $x$ gets closer to $0$}; \end{tikzpicture} \end{document}
```

```tikz
\begin{document} \begin{tikzpicture} \draw[->] (-1,0) -- (5,0) node[right] {$x$}; \draw[->] (0,-1) -- (0,3) node[above] {$f(x)$}; \draw[dashed] (2,0) -- (2,2) node[above] {$a$}; \draw[domain=0.5:4.5,smooth,variable=\x,blue] plot ({\x},{0.5*(\x-2)^2 + 1}); \draw[dotted] (2,1) -- (0,1) node[left] {$L$}; \draw[->] (2.5,0.8) -- (3.5,1.8) node[midway,above,sloped] {$f(x)$ approaching $L$}; \end{tikzpicture} \end{document}
```

`8:45p`

https://magazine.sebastianraschka.com/p/ahead-of-ai-10-state-of-computer

`8:30p`

```tikz
\begin{document}
\begin{tikzpicture}[domain=0:4]
    % Axis
    \draw[->] (-1,0) -- (6,0) node[right] {$x$};
    \draw[->] (0,-1) -- (0,5) node[above] {$y$};

    % Function
    \draw[domain=0.5:5.5,smooth,variable=\x,blue] plot ({\x},{0.2*\x^2}) node[right] {$f(x)$};

    % Point and tangent
    \coordinate (A) at (2,0.8);
    \coordinate (T) at (4,1.6);

    \draw[dashed] (A) -- (A |- 0,0) node[below] {$x$};
    \draw[dashed] (A) -- (0,0.8) node[left] {$f(x)$};
    \draw[dashed] (T) -- (T |- 0,0) node[below] {$x+dx$};
    \draw[dashed] (T) -- (0,1.6) node[left] {$f(x+dx)$};

    \draw[thick,red] (A) -- (T) node[midway, above left] {\small $dy$};

    % Tangent line
    \draw[thick,orange] (1.5,-0.5) -- (5.5,2.5) node[right] {Tangent Line};

    % Differential dx
    \draw[<->,purple] (2,-0.2) -- (4,-0.2) node[midway,below] {\small $dx$};

    % Derivative label
    \node[align=center] at (3.5,-1) {\small Derivative: \\ $\frac{dy}{dx} = f'(x)$};
\end{tikzpicture}
\end{document}

```

4:30p Venn diagram of strengths and interests

[ Understanding Llama 2 and the New Code Llama LLMs](https://magazine.sebastianraschka.com/p/ahead-of-ai-11-new-foundation-models) ([sebastianraschka.com](https://news.ycombinator.com/from?site=sebastianraschka.com))

- https://news.ycombinator.com/item?id=37321032

`8:45a` https://realpython.com/python-data-types/

## `2023-08-30`

`10:35p`

- https://stackoverflow.com/questions/426963/replace-tabs-with-spaces-in-vim

​ `:set tabstop=2 shiftwidth=2 expandtab | retab`

https://developer.mozilla.org/en-US/docs/Web/HTML/Quirks_Mode_and_Standards_Mode?utm_source=mozilla&utm_medium=firefox-console-errors&utm_campaign=default

`5p` stackoverflow.com/questions/16862627/json-stringify-output-to-div-in-pretty-print-way

`2p`

`uvicorn hello_falcon:app --host 0.0.0.0 --port 8000 --reload`

- [ ] Clean up docs to get collaborators up-and-running faster
- [ ] DB Browser (sqlite) for Nick
- [x] Run Redis on Podman engine
- [x] Run Falcon on Podman engine and connect to Redis
  - Note, this involved getting IP from inside (ssh) running Redis container

```
base) ➜  app git:(main) ✗ uvicorn hello_falcon:app
- Redis working with JSON.
 - Wow.
INFO:     Started server process [89980]
INFO:     Waiting for application startup.
INFO:     Application startup complete.
INFO:     Uvicorn running on http://127.0.0.1:8000 (Press CTRL+C to quit)
^C
INFO:     Shutting down
INFO:     Waiting for application shutdown.
INFO:     Application shutdown complete.
```

`12:46p`

`podman run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest`

`12:35p`

Success running Flask, then Falcon in Podman.

`podman build -t hello:falcon -f Falconfile .`

- https://redis.io/docs/data-types/json/
- https://falcon.readthedocs.io/en/stable/user/quickstart.html#learning-by-example

`10:28a`

- `%s/import//g` copy imports into `requirements.txt` and remove-all `import` keywords
- `sudo systemctl start/stop/status sshd.service`

`9:30a`

- https://math.stackexchange.com/questions/23902/what-is-the-practical-difference-between-a-differential-and-a-derivative#23914

## `2023-08-29`

`8:30p`

`8p`

https://developer.ibm.com/series/kubernetes-learning-path/

`5p`

[https://tailscale.com](https://tailscale.com/)

https://github.com/x0rz/EQGRP

https://news.ycombinator.com/item?id=37307681

https://www.modular.com/blog/how-mojo-gets-a-35-000x-speedup-over-python-part-2

https://github.com/ashvardanian/Stringzilla

https://news.ycombinator.com/item?id=37271162

https://superfastpython.com/numpy-share-array-processes/

## `2023-08-28`

Chain Rule, Power Rule are used in Differentiation. I knew this on one level probably for years but it clicked over the last couple days. 🚩personal observation. Zettlekasten: stick to meaningful connections.

Power rule is something about moving the exponent.

## `2023-08-27`

`1:11p` https://adamj.eu/tech/2021/05/11/python-type-hints-args-and-kwargs/

https://github.com/satwikkansal/wtfpython

`1:41p` https://sourcegraph.com/get-cody

## `2023-08-26`

`8:21a` ChatGPT thread on process improvement. How to spin `peregrine` out into something or get ahead of my personal narrative or brand.

https://en.wikipedia.org/wiki/A_Manual_for_Writers_of_Research_Papers,_Theses,_and_Dissertations

`1p` https://www.quora.com/How-can-I-do-AI-research-independently-without-college?share=1

## `2023-08-25`

https://en.wikipedia.org/wiki/Automatic_differentiation#Python

- https://vmartin.fr/comprendre-la-differenciation-automatique-en-30-lignes-de-python-fr.html _au français_
- https://photonlines.substack.com/p/grokking-hash-array-mapped-tries

- https://github.com/bsilverthorn/maccarone _delegated_ AI code assistance

From Peter Norvig:

- https://github.com/norvig/pytudes

`8:15a`

I have about 45 min until the library opens and then and hour until I'm supposed to meet with N.

- Automatic Differentiation
  - to find the partial derivative

SQLite changelog

- https://sqlite.org/releaselog/3_43_0.html
- https://blog.whimslab.io/how-modern-sql-databases-are-changing-web-development-4d1f15b0daea

> _I wish I could outline these changes and learn what I don't know._
> You can, with GPT.

- [x] plugin for `tikz` in Obsidian

https://github.com/artisticat1/obsidian-tikzjax/blob/main/README.md

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

`8:41a` This worked instantly. Cheers to that.

`11a`

$$
slope = \frac{\Delta y}{\Delta x}
$$

Libeinitz Notation:

$$
derrivative = \frac{d y}{d x}
$$

LaGrange Notation

$$
f'(x_1)
$$

`2:47p`

###### `lima`

_Ended up getting podman to work which simplified things._

`brew install lima`

- https://opensource.com/article/21/9/run-containers-mac-lima
- https://github.com/lima-vm/lima
-

`2:51p` Also downloading Rancher... we'll see.

`3:18p` Out of space, emptied trash, etc.

https://mark-riedl.medium.com/toward-agi-what-is-missing-c2f0d878471a

https://www.khanacademy.org/math/algebra/x2f8bb11595b61c86:forms-of-linear-equations/x2f8bb11595b61c86:intro-to-slope-intercept-form/v/slope-intercept-form

https://www.youtube.com/watch?v=S0_qX4VJhMQ&list=PL0-GT3co4r2wlh6UHTUeQsrf3mlS2lk6x&index=3 essence of calculus

https://www.youtube.com/watch?v=MQzG1hfhow4 neural network from scratch

https://github.com/VikParuchuri/zero_to_gpt/blob/master/explanations/linreg.ipynb
https://github.com/VikParuchuri/zero_to_gpt/blob/master/explanations/dense.ipynb

## `2023-08-20`



**Taylor's Theorem**

_And Taylor Series, relating back to logarithms and Euler's number._

- https://en.wikipedia.org/wiki/Taylor%27s_theorem
- https://en.wikipedia.org/wiki/Taylor_series

**Motivation**

I want to actually tie this back to the logarithm `ipynb` ... today this is an attempt at how to track ongoing studies day over day.

[[2023-08-20]] _This was one way I had been linking concept documents to journal entries._

LaTeX for relevant concepts:

- https://tex.stackexchange.com/questions/512201/how-to-write-taylors-series-of-sinx-in-pstricks
- https://www.alanshawn.com/latex3-tutorial/
- https://www.tug.org/TUGboat/tb24-3/wong.pdf (just for fun, I guess)
## `2023-08-19`

##### LaTeX for pretty graphs of Taylor Series

- https://tex.stackexchange.com/questions/512201/how-to-write-taylors-series-of-sinx-in-pstricks

`4:20p` Apple support. 

*Apple Pencil was under warranty. That helps a lot.*

# `2023-08-18`

```
 Taylor Series or Taylor's Theorem 
- Polynomials
- Differentiable Function 
- Differentiation 
  - Leibinitz notation
  - Lagrange notation
- natural logarithm 
- Euler's number
```
## `2023-08-08`

`8:37p` trying to get some other little CLI things going. `viteris`.

https://pastebin.com/gWggYb0a

```
brew isntall irssi
/connect irc.scuttled.net
```

https://irssi.org/documentation/help/connect/

https://quadpoint.org/articles/irssi/



## `2023-08-02`

```bash
 echo "---------";env;echo "-----------"
 alias
```
## `2023-07-08`

- [[2023-07-08#SchemDraw]]
- [[2023-07-08#txtai]]
- [[2023-07-08#En Español, por favor.]]
## `2023-07-01`

[[2023-07-01]] has a lot of links on ML...

## `2023-06-04`


`12a` Opened up notebook from `2006`. Mysticism, math, DSP and music theory all in one place. I signed it `6-22-06`. 

I have decided, in honor of my ancestor John Napier, to study the logarithm this evening, starting on paper and hopefully moving into a Jupyter notebook or something. 

### logarithms

*logos* *arithmetic*

`12:23a` Noting my memories of confusing `logarithm` with `alogrithm` probably 10 years ago, maybe around the time at Scrib in Boulder. I think at that time I was studying algorithms for the first time and knew they were different, but had an earlier memory of confusing both words, probably going back to middle school.  

- https://math.stackexchange.com/questions/61209/what-algorithm-is-used-by-computers-to-calculate-logarithms

$$\log_216 = 4$$
because

$$ 2^4 = 2 \times 2 \times 2 \times 2 = 16$$
`12:36a` Interesting observation of passage of time, noticed the other night that 10min felt much longer. I looked up 'time dilation' because I wasn't sure about the usage of the word. 

- https://en.wikipedia.org/wiki/Time_dilation

Also, thought maybe the mini was running hot, couldn't log in remote. But I had forgotten I removed the local keyboard's batteries. 
Either way, the mini is not powerful enough and I need to consolidate environments. There is still a bunch of stuff on there though. 

`12:51a` Coming up on `1a` and then it will be only a few hours until my walk. I made good first steps toward properly understand logarithms. 

##### `log` links
- https://en.wikipedia.org/wiki/Logarithm *review with historical context*
- https://en.wikipedia.org/wiki/Mirifici_Logarithmorum_Canonis_Descriptio
- https://en.wikipedia.org/wiki/Prosthaphaeresis

- https://artofproblemsolving.com/wiki/index.php/Logarithm
- https://www.overleaf.com/articles/logarithms/skwshvtxmpqt.pdf

##### python things
- https://stackoverflow.com/questions/23832864/python-logarithm-function
- https://docs.sympy.org/dev/modules/solvers/solvers.html


##### slide rule & analog computers

- https://mathematicalmysteries.org/slide-rule/
- https://blog.cambridgecoaching.com/slide-rules-logarithms-and-analog-computers

Found a touch point with the *chain rule*.

https://math.stackexchange.com/questions/112235/derivatives-of-logarithmic-functions-and-the-chain-rule
https://math.stackexchange.com/questions/4060611/chain-rule-of-logarithm


`1a`

- simple Khan Academy intro
- https://youtu.be/Z5myJ8dg_rM

`1:48a` closing out tabs

`1:57a` https://www.youtube.com/watch?v=PupNgv49_WY *more advanced logarithm things*

`2a`

https://www.indiatoday.in/education-today/gk-current-affairs/story/remembering-the-person-behind-all-computer-software-john-napier-1204345-2018-04-04

_A Plaine Discovery of the Whole Revelation of St. John_.

https://en.wikipedia.org/wiki/Systemantics

`2:30a` https://doc.lagout.org/readme.txt good **French** site




---
## about

This file was created to capture technical notes separate from personal reflection and journaling. I had been using a single daily entry to capture everything, which after a while made it difficult to refer to ongoing studies.

So, as a rule cathartic journaling and reflections on process belong in those daily entries, whereas hard information belongs here.

Additional mind-mapping, vaults organized around topic, etc are still being explored. Ex. a Math file in historical order could be split into a vault, I don’t know.

[[2023-06-15]] has some interesting chronology 
[[2023-06-04]] This file has a lot of reflection.
