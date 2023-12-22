#### language

[[Think-C.pdf]]
[[bgc_usl_c_1.pdf]]

[[pointers]]
[[keywords]]
[[hexadecimal]]

#### compiling
[[linking & compiling]]

#### POSIX
file system, interprocess communication (IPC), network programming
[[POSIX]]

Realtime Audio in the BSD kernel [[URN_NBN_fi_jyu-2005243.pdf]]

#### reverse engineering 
[[r2]]
##### assembly 




#### links

- https://floooh.github.io/2019/09/27/modern-c-for-cpp-peeps.html 
- http://staff.washington.edu/gmobus/Academics/TCES202/Moodle/OO-ProgrammingInC.html
- https://github.com/microsoft/checkedc
- https://github.com/google/sanitizers
- http://blog.pkh.me/p/37-gcc-undefined-behaviors-are-getting-wild.html
- https://tmewett.com/c-tips/
- https://best.openssf.org/Compiler-Hardening-Guides/Compiler-Options-Hardening-Guide-for-C-and-C++.html
- https://beej.us/guide/bgipc/html/
- https://ocw.mit.edu/courses/6-s096-introduction-to-c-and-c-january-iap-2013/pages/lectures-and-assignments/data-structures-debugging/
- https://github.com/PacktPublishing/Extreme-C/blob/master/ch04-process-memory-structure/ExtremeC_examples_chapter4_8.c
- https://llvm.org
- https://lldb.llvm.org
- https://www.deusinmachina.net/p/c-strings-and-my-slow-descent-to
- https://awesomekling.github.io/MutexProtected-A-C++-Pattern-for-Easier-Concurrency/
- https://progforperf.github.io/references.html

## Memory Management


Nasal Demons - [http://www.catb.org/jargon/html/N/nasal-demons.html](http://www.catb.org/jargon/html/N/nasal-demons.html)

- https://blog.trailofbits.com/2022/11/10/divergent-representations-variable-overflows-c-compiler/
- https://www.memorysafety.org
- https://deepu.tech/memory-management-in-programming/
- http://blog.llvm.org/2011/05/what-every-c-programmer-should-know_14.html 

https://github.com/pkivolowitz/asm_book/tree/main

https://www.latencytop.org/announcement.html latencytop

https://www.cppstories.com/p/start-here/ C++ news, changes explained)



#### misc 

https://en.wikipedia.org/wiki/Tak_%28function%29
https://github.com/aike/tarai_language/blob/master/player.js
http://aikelab.net/tarai/

[![](https://static.wixstatic.com/media/a20762_3f2de2596b1f4dd29dd481fdfeda06b3~mv2.png/v1/fill/w_79,h_80,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/a20762_3f2de2596b1f4dd29dd481fdfeda06b3~mv2.png)](https://www.vult-dsp.com/)
vult
https://vult-dsp.github.io/vult/tutorials/generate-c/
https://www.vult-dsp.com/vult-language

http://ccr.sigcomm.org/online/files/p83-keshavA.pdf


## Stack vs Heap

#### ChatGPT:
Here are some key differences between stack memory and heap memory in C:

1.  Allocation: Stack memory is allocated automatically by the system when a function is called, whereas heap memory is allocated explicitly using functions like `malloc()` and `calloc()`.
2.  Size: The amount of stack memory is typically limited and determined by the system, whereas heap memory is generally larger and can be dynamically allocated as needed.
3.  Speed: Accessing stack memory is generally faster than accessing heap memory because it involves simply incrementing or decrementing a pointer, whereas accessing heap memory requires additional overhead to locate the appropriate memory block.
4.  Scope: Variables allocated on the stack are typically only accessible within the function where they are declared, whereas variables allocated on the heap can be accessed from any part of the program, as long as a valid pointer to the memory block is available.
5.  Management: Stack memory is automatically managed by the system, which means there is no need to explicitly allocate or free it. Heap memory, on the other hand, must be explicitly allocated and freed to avoid memory leaks or dangling pointers.   
6.  Fragility: Accessing heap memory incorrectly can lead to a wide range of bugs and security vulnerabilities, such as buffer overflows and memory leaks. In contrast, accessing stack memory incorrectly usually leads to an immediate crash or other obvious error, which can make it easier to detect and fix problems.

In general, stack memory is faster and easier to manage, but has limited size and scope, whereas heap memory is more flexible, but requires explicit management and can be more fragile. Careful consideration of the trade-offs between stack and heap memory is necessary to write efficient and memory-safe programs.


https://shop.learncodethehardway.org/products/contents/3/?page=1&query=

Maybe move to its own file eventually... 

- ThinkC
- ThinkOS
- LearnCTheHardway





## Massacring C Pointers

I'm taking a break from debugging books to talk about a calamitous shitshow of textbook writing: *Mastering C Pointers: Tools for Programming Power*, by Robert J. Traister.

I learned of the book through a [talk](https://www.youtube.com/watch?v=8SUkrR7ZfTA) by Brian Kernighan where he refers to the book as probably “the worst C programming textbook ever written.”\[[1](https://wozniak.ca/blog/2018/06/25/1/index.html#fn.1)\] He doesn't name it but [with some help](https://twitter.com/johnregehr/status/1003411102791692288) I was able to track down his obliquely accurate reference.

This book has become my [white whale](https://www.urbandictionary.com/define.php?term=white%20whale). Since I started reading debugging books, and especially now that I'm digging through older ones, I find bits of advice that simply don't work today. While some of it could be construed as useless or idiotic, I've always found the authors come from a position of earnestness, attempting to draw the best conclusions based on decent principles and what they knew at the time they wrote it. In some cases they may not have known much, but they're honestly and humbly trying to impart some wisdom.

When Kernighan put up the following example, I saw what seemed to be the opposite of that.

```C
char *combine(s, t)
char \*s, \*t;
{

      int x, y;
      char r\[100\];

      strcpy(r, s);
      y = strlen(r);
      for (x = y; *t != '\\0'; ++x)
           r\[x\] = *t++;

      r\[x\] = '\\0';

      return(r);

}
```

This program (formatting preserved) is taken from the first edition of the book (p.146). It's horrific (Kernighan calls it “malpractice”). It does not exhibit the genuineness I have seen with, say, books from the late 1970s on how to debug BASIC programs. I didn't know what it was. Deceit? Laziness? Unremitting ignorance? What is the mindset of someone who writes this, presumably thinking it's a good idea? Can the whole book be that bad? Kernighan said it was. I had to know.

The book has two editions: the first was published in 1990, the second in 1993. The fact that there are two editions piqued my curiosity even more. It sold enough to make another version? Is that horrible example corrected? I obtained a copy of each and read them.

Reviewing the book is pointless. Kernighan was right: it's garbage. And the second edition only makes things worse. Teaching from either one would be a breach of ethics. If you follow Traister’s coding practices, even adjusted for today's standards, you are *guaranteed* to create defects and vicious, latent bugs. A subtly pernicious aspect of the book is the casual tone of the writing. It’s informal enough that if you don’t really know much about C what he says sort of makes sense, despite the sloppy terminology and mixed, inaccurate metaphors.

Any trained programmer will recognize the lessons as worthless. His terminology is all over the map and typically inaccurate, if not plainly wrong. Expressions “return a value.” Values are “handed” into locations. Constants are “written directly into the program.” A union is a “specialized pointer.” The terminology isn’t even consistent. Micro-optimizations are stressed at all times, and program efficiency is valued over comprehension. He can’t even define a pointer correctly: “A pointer is a special variable that returns the address of a memory location.” It does not take long to realize Traister has no idea what he’s talking about.

Although I’m not going to go over the content (that would take far too long since there’s nothing redeeming there), I did take extensive notes. You can [read them](https://wozniak.ca/blog/2018/06/25/1/notes.html), if you are so inclined.

I must, however, take a moment to single out the code. It is universally bad and much of it is simply wrong. (Imagine trying to learn programming principles from a book that contains a large number of programs that don’t even compile.) I’ve [transcribed](https://wozniak.ca/blog/2018/06/25/1/code.html) some of the programs and annotated them with comments so that you can get a taste of how inept Traister is as a C programmer. One thing to keep in mind (both for the programs and the notes) is that C89/C90 was new at the time and that the code was written on (and for) MS-DOS systems of the late 1980s/early 1990s. Things were a bit different then.

Enough about the material. I want to explore the question of how something so wrong even got written. It’s not that everything in the book is wrong, but it feels like when it’s right, it’s right by accident.

Traister has [written](https://www.amazon.com/Robert-J.-Traister/e/B001H6UPHY/ref=sr_ntt_srch_lnk_1?qid=1529540256&sr=1-1) other books, some about electronics and some about programming, one called *Going from BASIC to C*. In *Mastering C Pointers*, he talks about a product he created called CBREEZE that converts BASIC code to C. Throughout the book he makes passing, roundabout references to BASIC and uses terminology that suggests he’s written a lot of BASIC code. For example, there is a whole chapter on using pointers to access memory, where reading and writing memory is instead called “peeking” and “poking”, based on the PEEK and POKE instructions in BASIC. He also says that it took him a couple of tries to learn C coming from BASIC. In short, I’m convinced he’s knowledgeable about BASIC and has worked on writing software for small, electronic devices.

Why is this important? As I read the book (and if you read my notes, you know where this is going) I started to notice something in the wording and tone. The further I progressed the more I became convinced of it, and I think it explains how he managed to mangle the explanation of C pointers so badly.

I don’t think he understands the call stack.

My argument for this interpretation requires a little knowledge of BASIC and embedded devices.

With BASIC, the key thing to know about most implementations at the time is that there were no functions and no scope aside from the global scope.\[[2](https://wozniak.ca/blog/2018/06/25/1/index.html#fn.2)\] The closest thing to a function in BASIC is the GOSUB command. The GOSUB command jumps to a line and executes code until it gets to a RETURN statement, where control is transferred back to the line following the GOSUB command. Within a GOSUB you can jump somewhere else with another GOSUB. The control follows a stack principle, but no arguments are passed. GOSUB routines are a way to factor out common code, but that common code has to work on global variables. (And yes, it’s as terrible as it sounds.)

Now, consider the case of simple electronic devices. Even today some embedded devices, usually programmed using C, do not have a call stack that dynamically allocates space for automatic variables. There simply isn’t enough memory for it. Instead, the compiler lays out memory such that each function’s local variables have fixed memory addresses (a “compiled stack” model). The only stack you have is for return addresses and it is probably handled in hardware.

Suppose you’re used to writing BASIC for small memory electronic devices and you learn about C. You read about pointers and realize something: it’s possible to write a subroutine that can change variables *without knowing their names*. It’s manna from heaven! You don’t have to devote global variables to being the “parameters” of your subroutines anymore. Life is great.

This is the mindset I think Traister had and never got past. In the book there is one fleeting mention of the stack in reference to excessive (automatic) memory allocation. (On MS-DOS, if the space for local variables is too large in the program, it might not compile.) He consistently describes variables as having “exclusive” addresses in the program. His writing about pointers suggests that he thinks, for each function, space is set aside to hold the local variables for the duration of the program, but you can only access them when inside the function in which they were declared. So pointers are really powerful because you can provide this address to another function and it can change the value using only a parameter.

Further evidence for his lack of understanding is that he frequently cites ridiculous space micro-optimizations within functions, such as avoiding the use of integers for index variables, if possible. Another one, mentioned often, is local `char` arrays that have a fixed size. There are good reasons to not use them but his are not among them. His admonishment is that they waste space. Technically, that is true, but they don't exist until they're on the stack. And he never talks about global or file variables. He only refers to locals with “exclusive” addresses “set aside” for variables.

This interpretation runs into some problems once you start asking how functions with malloc will work, but it's worth pointing out that there is almost no discussion about memory management. In a book devoted to C pointers, that's a toxic mix of gross negligence and incompetence. There is literally one short paragraph devoted to talking about the `free` function—and it's characterized as a “side note.”

Another sticking point in this interpretation is Traister’s incomprehensible approach to writing functions that take a variable number of arguments. He does this by passing an arbitrary number of arguments to a function (the first being the number of arguments) and accessing them using offsets from the address of the first argument. This suggests he has some idea about parameters being passed in a dynamic fashion, but it is so spectacularly wrong\[[3](https://wozniak.ca/blog/2018/06/25/1/index.html#fn.3)\] you’re left wondering if he even tried his programs out before publishing them.

Honestly, this is the most generous interpretation of the text I could come up with, and it still paints a terrible picture. Occam's Razor suggests that Traister is just clueless. But like [analyzing a terrible movie](https://www.youtube.com/watch?v=01l1WIC9mBo) that somehow gets made, it's more fun to reason through the “behind the scenes” parts.

Given the ineptness of the book, you'd think it was self-published. You would be wrong. It was published through Academic Press, which was a division of Harcourt, Brace & World at the time, but is now an “imprint of Elsevier.”

In the preface of the second edition it says that the first edition was reviewed “by a professional C programmer hired by the publisher.” That programmer said it should not be published. That programmer was right, but the publisher went ahead and published it anyway.

Since there was a second edition, the assumption is that the book sold well. According to WorldCat, *Mastering C Pointers* is in at least [242 libraries](https://www.worldcat.org/title/mastering-c-pointers/oclc/898329209&referer=brief_results), most appearing to be the first edition, but I didn’t check them all.\[[4](https://wozniak.ca/blog/2018/06/25/1/index.html#fn.4)\] It claims to be one of the first books to tackle the subject of pointers in C, which is often a sticking point for novice programmers. The lack of material in this area at the time is probably why it sold. I could not track down a review of this book anywhere (and yes, I looked through scans of Byte Magazine et al.), but I did find reviews for other C books in the 1980s and what I found suggested that pointers were not covered well, if at all. In other words, like many books—and tech books in particular—it sold because of its title and good timing.\[[5](https://wozniak.ca/blog/2018/06/25/1/index.html#fn.5)\]

If you browse search results for other books by Traister you’ll find a lot of questionable sounding titles: *Making money with your microcomputer* (1982), *Leaping from BASIC to C++* (1994), *Learn C in Two Weeks with Run/C and CBreeze* (1987). The breadth of topics covered in his works seems exhausting: *Beginner’s Guide to Reading Schematics* (1991), *Astronomy and Telescopes: A Beginner’s Handbook* (1983), *Make your own professional home video recordings* (1982), *Cave Exploring* (1983), just to name a few. You start to wonder if this is actually the same person who writes all this. Then you start to wonder if maybe all these books just touch on the topics, and are churned out mainly to try and make a buck. I had a hard time finding a review for any of them (the best resource was [archive.org](https://archive.org/search.php?query=robert%20j%20traister&sin=TXT)). Practically all I could find about his books was ads for them—and even then it wasn’t that much, which is odd given the apparent volume of output from him. I did manage to find two reviews for his book *Programming in C for the Microcomputer User* (1984): one was favourable (80 Micro, Nov. 1984), the other was not (Practical Computing, Oct. 1985). The only other book I can find of his that seems to have some staying power is the schematics book, which went to a third edition.\[[6](https://wozniak.ca/blog/2018/06/25/1/index.html#fn.6)\]

Ultimately, the aspect of *Mastering C Pointers* that truly disturbs me is that there are probably a fair number of people who actually learned C pointers from it. There’s no way to know how much of an impact this book had on programmers in the 1990s, but given the number of copies in libraries it must have had some. It’s hard not to wonder how much of the terrible C code that has made its way into production can be attributed to the awful advice in Traister’s travesty of a text.

- [Annotated code samples from the book](https://wozniak.ca/blog/2018/06/25/1/code.html)
- [Notes I took](https://wozniak.ca/blog/2018/06/25/1/notes.html)

Thanks to John Regehr for helping me track down the book in the first place. The title was ~stolen~ inspired by one of his tweets.

June 25, 2018

\[[1](https://wozniak.ca/blog/2018/06/25/1/index.html#fnr.1)\]

His mention of the book comes at 27:44 in the linked talk. Also, if you watch the talk, be forewarned that an alarm goes off at around the 55 minute mark and it's quite jarring.

\[[2](https://wozniak.ca/blog/2018/06/25/1/index.html#fnr.2)\]

This is not strictly true. You could define a function, but it could only take one argument and had to be defined by an expression that could fit on one line.

\[[3](https://wozniak.ca/blog/2018/06/25/1/index.html#fnr.3)\]

Of course, it’s more nuanced than this because you can actually make his code work in some cases. See the discussion in the code samples for more information.

\[[4](https://wozniak.ca/blog/2018/06/25/1/index.html#fnr.4)\]

I tried to get some data on the number of times it had been loaned out by my alma mater’s library system, but was not able to get anything.

\[[5](https://wozniak.ca/blog/2018/06/25/1/index.html#fnr.5)\]

This is not meant to say anything particular about the tech publishing industry at the time. I think the same opportunism that was in effect then is alive and well today, and probably always has been. Arguably, it’s worse today.

\[[6](https://wozniak.ca/blog/2018/06/25/1/index.html#fnr.6)\]

I only know this because of an odd, rambling [promotional video](https://www.youtube.com/watch?v=7hX1BkWlUso) for the book where Traister’s name is forgotten by the new author.

## Undefined Behavior

(and what every C programmer should know)

```
By Chris Lattner
May 13, 2011
```
### 1/3

People occasionally ask why LLVM-compiled code sometimes generates SIGTRAP signals when the optimizer is turned on. After digging in, they find that Clang generated a "ud2" instruction (assuming X86 code) - the same as is generated by __builtin_trap(). There are several issues at work here, all centering around undefined behavior in C code and how LLVM handles it.

This blog post (the first in a series of three) tries to explain some of these issues so that you can better understand the tradeoffs and complexities involved, and perhaps learn a few more of the dark sides of C. It turns out that C is *not* a "high level assembler" like many experienced C programmers (particularly folks with a low-level focus) like to think, and that C++ and Objective-C have directly inherited plenty of issues from it.

## Introduction to Undefined Behavior

Both LLVM IR and the C programming language have the concept of "undefined behavior". Undefined behavior is a broad topic with a lot of nuances. The best introduction I've found to it is a post on [John Regehr's Blog](http://blog.regehr.org/archives/213). The short version of this excellent article is that many seemingly reasonable things in C actually have undefined behavior, and this is a common source of bugs in programs. Beyond that, any undefined behavior in C gives license to the implementation (the compiler and runtime) to produce code that formats your hard drive, does completely unexpected things, [or worse](http://www.catb.org/jargon/html/N/nasal-demons.html). Again, I would highly recommend reading [John's article](http://blog.regehr.org/archives/213).

Undefined behavior exists in C-based languages because the designers of C wanted it to be an extremely efficient low-level programming language. In contrast, languages like Java (and many other 'safe' languages) have eschewed undefined behavior because they want safe and reproducible behavior across implementations, and willing to sacrifice performance to get it. While neither is "the right goal to aim for," if you're a C programmer you really should understand what undefined behavior is.

Before getting into the details, it is worth briefly mentioning what it takes for a compiler to get good performance out a broad range of C apps, because: 

**There is no magic bullet**. 

At a very high level, compilers produce high performance apps by a) doing a good job at bread and butter algorithms like register allocation, scheduling, etc. b) knowing lots and lots of "tricks" (e.g. peephole optimizations, loop transformations, etc), and applying them whenever profitable. c) being good at eliminating unnecessary abstractions (e.g. redundancy due to macros in C, inlining functions, eliminating temporary objects in C++, etc) and d) not screwing anything up. 

> While any of the optimizations below may sound trivial, it turns out that saving just one cycle out of a critical loop can make some codec run 10% faster or take 10% less power.



## Advantages of Undefined Behavior in C, with Examples


Before getting into the dark side of undefined behavior and LLVM's policy and behavior when used as a C compiler, I thought it would be helpful to consider a few specific cases of undefined behavior, and talk about how each enables better performance than a safe language like Java. You can look at this either as "optimizations enabled" by the class of undefined behavior or as the "overhead avoided" that would be required to make each case defined. While the compiler optimizer could eliminate some of these overheads some of the time, to do so in general (for every case) would require solving the halting problem and many other "interesting challenges".

It is also worth pointing out that both Clang and GCC nail down a few behaviors that the C standard leaves undefined. The things I'll describe are both undefined according to the standard and treated as undefined behavior by both of these compilers in their default modes.

### **use of an uninitialized variable** 

This is commonly known as source of problems in C programs and there are many tools to catch these: from compiler warnings to static and dynamic analyzers. This improves performance by not requiring that all variables be zero initialized when they come into scope (as Java does). For most scalar variables, this would cause little overhead, but stack arrays and malloc'd memory would incur a memset of the storage, which could be quite costly, particularly since the storage is usually completely overwritten.

### signed integer overflow

If arithmetic on an 'int' type (for example) overflows, the result is undefined. One example is that "INT_MAX+1" is not guaranteed to be INT_MIN. This behavior enables certain classes of optimizations that are important for some code. 

For example, knowing that INT_MAX+1 is undefined allows optimizing "X+1 > X" to "true". Knowing the multiplication "cannot" overflow (because doing so would be undefined) allows optimizing "X*2/2" to "X". While these may seem trivial, these sorts of things are commonly exposed by inlining and macro expansion. A more important optimization that this allows is for "<=" loops like this:

```c
for (i = 0; i <= N; ++i) { ... }
```

In this loop, the compiler can assume that the loop will iterate exactly N+1 times if "i" is undefined on overflow, which allows a broad range of loop optimizations to kick in. On the other hand, if the variable is defined to wrap around on overflow, then the compiler must assume that the loop is possibly infinite (which happens if N is INT_MAX) - which then disables these important loop optimizations. This particularly affects 64-bit platforms since so much code uses "int" as induction variables.

It is worth noting that unsigned overflow is guaranteed to be defined as 2's complement (wrapping) overflow, so you can always use them. The cost to making signed integer overflow defined is that these sorts of optimizations are simply lost (for example, a common symptom is a ton of sign extensions inside of loops on 64-bit targets). Both Clang and GCC accept the "-fwrapv" flag which forces the compiler to treat signed integer overflow as defined (other than divide of INT_MIN by -1).

### oversized shift amounts

Shifting a uint32_t by 32 or more bits is undefined. My guess is that this originated because the underlying shift operations on various CPUs do different things with this: for example, X86 truncates 32-bit shift amount to 5 bits (so a shift by 32-bits is the same as a shift by 0-bits), but PowerPC truncates 32-bit shift amounts to 6 bits (so a shift by 32 produces zero). Because of these hardware differences, the behavior is completely undefined by C (thus shifting by 32-bits on PowerPC could format your hard drive, it is ***not\*** guaranteed to produce zero). The cost of eliminating this undefined behavior is that the compiler would have to emit an extra operation (like an 'and') for variable shifts, which would make them twice as expensive on common CPUs.

### dereferencing

#### Dereferences of Wild Pointers and Out of Bounds Array Accesses

Dereferencing random pointers (like NULL, pointers to free'd memory, etc) and the special case of accessing an array out of bounds is a common bug in C applications which hopefully needs no explanation. To eliminate this source of undefined behavior, array accesses would have to each be range checked, and the ABI would have to be changed to make sure that range information follows around any pointers that could be subject to pointer arithmetic. This would have an extremely high cost for many numerical and other applications, as well as breaking binary compatibility with every existing C library.

#### Dereferencing a NULL Pointer

Contrary to popular belief, dereferencing a null pointer in C is undefined. It is *not defined to trap*, and if you mmap a page at 0, it is *not defined to access that page*. This falls out of the rules that forbid dereferencing wild pointers and the use of NULL as a sentinel. NULL pointer dereferences being undefined enables a broad range of optimizations: in contrast, Java makes it invalid for the compiler to move a side-effecting operation across any object pointer dereference that cannot be proven by the optimizer to be non-null. This significantly punishes scheduling and other optimizations. In C-based languages, NULL being undefined enables a large number of simple scalar optimizations that are exposed as a result of macro expansion and inlining.

> If you're using an LLVM-based compiler, you can dereference a "volatile" null pointer to get a crash if that's what you're looking for, since volatile loads and stores are generally not touched by the optimizer. There is currently no flag that enables random NULL pointer loads to be treated as valid accesses or to make random loads know that their pointer is "allowed to be null".

### **violating type rules**

It is undefined behavior to cast an int* to a float* and dereference it (accessing the "int" as if it were a "float"). C requires that these sorts of type conversions happen through memcpy: using pointer casts is not correct and undefined behavior results. The rules for this are quite nuanced and I don't want to go into the details here (there is an exception for char*, vectors have special properties, unions change things, etc). This behavior enables an analysis known as "Type-Based Alias Analysis" (TBAA) which is used by a broad range of memory access optimizations in the compiler, and can significantly improve performance of the generated code. For example, this rule allows clang to optimize this function:

```c
float *P;
 void zero_array() {
   int i;
   for (i = 0; i < 10000; ++i)
     P[i] = 0.0f;
 }
```


into "`memset(P, 0, 40000)`". This optimization also allows many loads to be hoisted out of loops, common subexpressions to be eliminated, etc. This class of undefined behavior can be disabled by passing the -fno-strict-aliasing flag, which disallows this analysis. When this flag is passed, Clang is required to compile this loop into 10000 4-byte stores (which is several times slower), because it has to assume that it is possible for any of the stores to change the value of P, as in something like this:

```c
int main() {
  P = (float*)&P;  // cast causes TBAA violation in zero_array.
  zero_array();
}
```

This sort of type abuse is pretty uncommon, which is why the standard committee decided that the significant performance wins were worth the unexpected result for "reasonable" type casts. It is worth pointing out that Java gets the benefits of type-based optimizations without these drawbacks because it doesn't have unsafe pointer casting in the language at all.

Anyway, I hope that this gives you an idea of some of the classes of optimizations enabled by undefined behavior in C. There are many other kinds of course, including sequence point violations like "foo(i, ++i)", race conditions in multithreaded programs, violating 'restrict', divide by zero, etc.

In our [next post](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know_14.html), we'll discuss why undefined behavior in C is a pretty scary thing if performance is not your only goal. In our final post in the series, we'll talk about how LLVM and Clang handle it.

### 2/3 

#### why undefined behavior is often a scary and terrible thing for c programmers

In [Part 1](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know.html) of our series, we discussed what undefined behavior is, and how it allows C and C++ compilers to produce higher performance applications than "safe" languages. This post talks about how "unsafe" C really is, explaining some of the highly surprising effects that undefined behavior can cause. In [Part #3](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know_21.html), we talk about what friendly compilers can do to mitigate some of the surprise, even if they aren't required to. 

I like to call this "**Why undefined behavior is often a scary and terrible thing for C programmers**". :-)

## Interacting Compiler Optimizations Lead to Surprising Results


A modern compiler optimizer contains many optimizations that are run in specific orders, sometimes iterated, and change as the compiler evolves over time (e.g. new releases come out). Also, different compilers often have substantially different optimizers. Because optimizations run at different stages, emergent effects can occur due to previous optimizations changing the code.

Lets take a look at a silly example (simplified from an exploitable bug that was found in the Linux Kernel) to make this more concrete:

```c
void contains_null_check(int *P) {
  int dead = *P;
  if (P == 0)
    return;
  *P = 4;
}
```


In this example, the code "clearly" checks for the null pointer. If the compiler happens to run "Dead Code Elimination" before a "Redundant Null Check Elimination" pass, then we'd see the code evolve in these two steps:

```c
void contains_null_check_after_DCE(int *P) {
  //int dead = *P;     // deleted by the optimizer.
  if (P == 0)
    return;
  *P = 4;
}
```

and then:

```c
void contains_null_check_after_DCE_and_RNCE(int *P) {
  if (P == 0)   // Null check not redundant, and is kept.
    return;
  *P = 4;
}
```


However, if the optimizer happens to be structured differently, it could run RNCE before DCE. This would give us these two steps:

```c
void contains_null_check_after_RNCE(int *P) {
  int dead = *P;
  if (false)  // P was dereferenced by this point, so it can't be null 
    return;
  *P = 4;
}
```


and then dead code elimination runs:

```c
void contains_null_check_after_RNCE_and_DCE(int *P) {
  //int dead = *P;
  //if (false)
  //  return;
  *P = 4;
}
```


To many (reasonable!) programmers, deleting the null check from this function would be very surprising (and they'd probably file a bug against the compiler :). However, both "contains_null_check_after_DCE_and_RNCE" and "contains_null_check_after_RNCE_and_DCE" are perfectly valid optimized forms of "contains_null_check" according to the standard, and both of the optimizations involved are important for the performance of various applications.

> While this is intentionally a simple and contrived example, this sort of thing happens all the time with inlining: inlining a function often exposes a number of secondary optimization opportunities. This means that if the optimizer decides to inline a function, a variety of local optimizations can kick in, which change the behavior of the code. This is both perfectly valid according to the standard, and important for performance in practice.

## Undefined Behavior and Security Don't Mix Well


The C family of programming languages is used to write a wide range of security critical code, such as kernels, setuid daemons, web browsers, and much more. This code is exposed to hostile input and bugs can lead to all sorts of exploitable security problems. One of the widely cited advantages of C is that it is relatively easy to understand what is going on when you read the code.

However, undefined behavior takes this property away. After all, most programmers would think that "contains_null_check" would do a null check above. While this case isn't too scary (the code will probably crash in the store if passed a null check, which is relatively easy to debug) there are a wide range of *very reasonable* looking C fragments that are completely invalid. This problem has bit many projects (including the Linux Kernel, OpenSSL, glibc, etc) and even led to CERT issuing a [vulnerability note](http://www.kb.cert.org/vuls/id/162289) against GCC (though my personal belief is that all widely-used optimizing C compilers are vulnerable to this, not just GCC).

Lets look at an example. Consider this carefully written C code:

```c
void process_something(int size) {
  // Catch integer overflow.
  if (size > size+1)
    abort();
  ...
  // Error checking from this code elided.
  char *string = malloc(size+1);
  read(fd, string, size);
  string[size] = 0;
  do_something(string);
  free(string);
}
```

This code is checking to make sure that the malloc is big enough to hold the data read from the file (because a nul terminator byte needs to be added), bailing out if an integer overflow error occurs. However, this is [exactly the example we gave before](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know.html#signed_overflow) in which the compiler is allowed to (validly) optimize out the check. This means that it is perfectly possible for the compiler to turn this into:

```c
void process_something(int *data, int size) {
  char *string = malloc(size+1);
  read(fd, string, size);
  string[size] = 0;
  do_something(string);
  free(string);
}
```


When being built on a 64-bit platform, it is quite likely that this is an exploitable bug when "size" is INT_MAX (perhaps the size of a file on disk). Lets consider how terrible this is: a code auditor reading the code would very reasonably think that a proper overflow check is happening. Someone testing the code would find no problem unless they specifically tested that error path. The secure code seems to work, until someone goes ahead and exploits the vulnerability. All in all, this is a surprising and quite scary class of bugs. Fortunately, the fix is simple in this case: just use "size == INT_MAX" or similar.

As it turns out, integer overflow is a security problem for many reasons. Even if you are using fully defined integer arithmetic (either by using `-fwrapv` or by using unsigned integers), there is a [wholly different class](http://en.wikipedia.org/wiki/Integer_overflow#Security_ramifications) of integer overflow bug possible. Fortunately, this class is visible in the code and knowledgable security auditors are usually aware of the problem.

## Debugging Optimized Code May Not Make Any Sense

Some people (for example, low level embedded programmers who like to look at generated machine code) do all of their development with optimizations turned on. Because code **frequently** has bugs when it is being developed, these folks end up seeing a disproportionate number of surprising optimizations that can lead to difficult-to-debug behaviors at runtime. For example, accidentally leaving out the "i = 0" in the "zero_array" example from the first article allows the compiler to completely discard the loop (compiling zero_array into "return;") because it is a use of an uninitialized variable.

Another interesting case that bit someone recently happened when they had a (global) function pointer. A simplified example looks like this:

```c
static void (*FP)() = 0;
static void impl() {
  printf("hello\n");
}
void set() {
  FP = impl;
}
void call() {
  FP();
}
```

which clang optimizes into:

```c
void set() {}
void call() {
  printf("hello\n");
}
```

It is allowed to do this because calling a null pointer is undefined, which permits it to assume that set() must be called before call(). In this case, the developer forgot to call "set", did not crash with a null pointer dereference, and their code broke when someone else did a debug build.

The upshot is that it is a fixable issue: if you suspect something weird is going on like this, try building at -O0, where the compiler is much less likely to be doing any optimizations at all.

## "Working" code that uses undefined behavior can "break" as the compiler evolves or changes 

We've seen many cases where applications that "appear to be work" suddenly break when a newer LLVM is used to build it, or when the application was moved from GCC to LLVM. While LLVM does occasionally have a bug or two itself :-), this is most often because of latent bugs in the application that are now being exposed by the compiler. This can happen all sorts different ways, two examples are:

1. an uninitialized variable which was zero initialized by luck "before", and now it shares some other register that isn't zero. This is commonly exposed by register allocation changes.

2. an array overflow on the stack which starts clobbering a variable that actually matters, instead of something that was dead. This is exposed when the compiler rearranges how it packs things on the stack, or gets more aggressive about sharing stack space for values with non-overlapping lifetimes.

The important and scary thing to realize is that just about *any* optimization based on undefined behavior can start being triggered on buggy code at any time in the future. Inlining, loop unrolling, memory promotion and other optimizations will keep getting better, and a significant part of their reason for existing is to expose secondary optimizations like the ones above.

To me, this is deeply dissatisfying, partially because the compiler inevitably ends up getting blamed, but also because it means that huge bodies of C code are land mines just waiting to explode. This is even worse because...

## There is No Reliable Way to Determine if a Large Codebase Contains Undefined Behavior

Making the landmine a much much worse place to be is the fact that there is **no good way** to determine whether a large scale application is free of undefined behavior, and thus not susceptible to breaking in the future. There are many useful tools that can help find **some** of the bugs, but nothing that gives full confidence that your code won't break in the future. Lets look at some of these options, along with their strengths and weaknesses:

1. The [Valgrind](http://valgrind.org/) [memcheck tool](http://valgrind.org/info/tools.html#memcheck) is a fantastic way to find all sorts of uninitialized variables and other memory bugs. Valgrind is limited because it is quite slow, it can only find bugs that still exist in the generated machine code (so it [can't find things the optimizer removes](http://blog.regehr.org/archives/519)), and doesn't know that the source language is C (so it can't find shift-out-of-range or signed integer overflow bugs).

2. Clang has an experimental `-fcatch-undefined-behavior` mode that inserts runtime checks to find violations like shift amounts out of range, some simple array out of range errors, etc. This is limited because it slows down the application's runtime and it can't help you with random pointer dereferences (like Valgrind can), but it can find other important bugs. Clang also fully supports the `-ftrapv` flag (not to be confused with `-fwrapv`) which causes signed integer overflow bugs to trap at runtime (GCC also has this flag, but it is completely unreliable/buggy in my experience). Here is a quick demo of `-fcatch-undefined-behavior`:

```c
$ cat t.c
int foo(int i) {
  int x[2];
  x[i] = 12;
  return x[i];
}

int main() {
  return foo(2);
}

$ clang t.c 
$ ./a.out 
$ clang t.c -fcatch-undefined-behavior 
$ ./a.out 
Illegal instruction
```

3. Compiler warning messages are good for finding some classes of these bugs, like uninitialized variables and simple integer overflow bugs. It has two primary limitations: 1) it has no dynamic information about your code as it executes, and 2) it must run very quickly because any analysis it does slows down compile time.

4. [The Clang Static Analyzer](http://clang-analyzer.llvm.org/) performs a much deeper analysis to try to find bugs (including use of undefined behavior, like null pointer dereferences). You can think of it as generating souped up compiler warning messages, because it is not bound by the compile time constraints of normal warnings. The primary disadvantages of the static analyzer is that it 1) doesn't have dynamic information about your program as it runs, and 2) is not integrated into normal workflows for many developers (though its integration into [Xcode 3.2 and later](http://developer.apple.com/technologies/mac/snowleopard/static.html) is fantastic). 

5. The [LLVM "Klee" Subproject](http://klee.llvm.org/) uses symbolic analysis to "try every possible path" through a piece of code to find bugs in the code and it **produces a testcase**. It is a great little project that is mostly limited by not being practical to run on large-scale applications.

6. While I have never tried it, the [C-Semantics tool](http://code.google.com/p/c-semantics/) by Chucky Ellison and Grigore Rosu is a very interesting tool that can apparently find some classes of bugs (such as sequence point violations). It is still a research prototype, but may be useful for finding bugs in (small and self-contained) programs. I recommend reading [John Regehr's post about it](http://blog.regehr.org/archives/523) for more information.

> The end result of this is that we have lots of tools in the toolbox to find some bugs, but no good way to prove that an application is free of undefined behavior. Given that there are lots of bugs in real world applications and that C is used for a broad range of critical applications, this is pretty scary. In our [final article](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know_21.html), I look at various options that C compilers have when dealing with undefined behavior, with a specific focus on [Clang](http://clang.llvm.org/).

## Why can't you warn when optimizing based on undefined behavior?

People often ask why the compiler doesn't produce warnings when it is taking advantage of undefined behavior to do an optimization, since any such case might actually be a bug in the user code. The challenges with this approach are that it is 1) likely to generate far too many warnings to be useful - because these optimizations kick in all the time when there is no bug, 2) it is really tricky to generate these warnings only when people want them, and 3) we have no good way to express (to the user) how a series of optimizations combined to expose the opportunity being optimized. Lets take each of these in turn:

**It is "really hard" to make it actually useful**

Lets look at an example: even though invalid type casting bugs are frequently exposed by type based alias analysis, it would not be useful to produce a warning that "the optimizer is assuming that P and P[i] don't alias" when optimizing "zero_array" (from [Part #1 of our series](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know.html)).

```c
float *P;
 void zero_array() {
   int i;
   for (i = 0; i < 10000; ++i)
     P[i] = 0.0f;
}
```


Beyond this "false positive" problem, a logistical problem is that the optimizer doesn't have enough information to generate a reasonable warning. First of all, it is working on an already-abstract representation of the code ([LLVM IR](http://llvm.org/docs/LangRef.html)) which is quite different from C, and second, the compiler is highly layered to the point where the optimization trying to "hoist a load from P out of the loop" doesn't know that TBAA was the analysis that resolved the pointer alias query. Yes, this is "the compiler guy whining" part of the article :), but it really is a hard problem.

**It is hard to generate these warnings \*only\* when people want them.**

Clang implements numerous warnings for simple and obvious cases of undefined behavior, such as out of range shifts like "x << 421". You might think that this is a simple and obvious thing, but it turns out that this is hard, because [people don't want to get warnings about undefined behavior in dead code](http://llvm.org/bugs/show_bug.cgi?id=5544) (see also [the duplicates](http://llvm.org/bugs/show_bug.cgi?id=6933)).

This dead code can take several forms: a macro that expands out in a funny way when when passed a constant, we've even had complaints that we warn in cases that would require [control flow analysis](http://llvm.org/bugs/show_bug.cgi?id=9322) of switch statements to prove that cases are not reachable. This is not helped by the fact that switch statements in C are [not necessarily properly structured](http://en.wikipedia.org/wiki/Duff's_device).

The solution to this in Clang is a growing infrastructure for handling "runtime behavior" warnings, along with code to prune these out so that they are not reported if we later find out that the block is unexecutable. This is something of an arms race with programmers though, because there are always idioms that we don't anticipate, and doing this sort of thing in the frontend means that it doesn't catch every case people would want it to catch.

**Explaining a series of optimizations that exposed an opportunity.**

If the frontend has challenges producing good warnings, perhaps we can generate them *from the optimizer* instead! The biggest problem with producing a useful warning here is one of data tracking. A compiler optimizer includes dozens of optimization passes that each change the code as it comes through to canonicalize it or (hopefully) make it run faster. If the inliner decides to inline a function, this may expose other opportunities for [optimizing away an "X*2/2"](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know.html#signed_overflow), for example.

While I've given [relatively simple and self-contained examples](http://blog.llvm.org/2011/05/what-every-c-programmer-should-know.html#optimizations) to demonstrate these optimizations, most of the cases where they kick in are in code coming from macro instantiation, inlining, and other abstraction-elimination activities the compiler performs. The reality is that humans don't commonly write such silly things directly. For warnings, this means that in order to relay back the issue to the users code, the warning would have to reconstruct exactly how the compiler got the intermediate code it is working on. We'd need the ability to say something like:

> ```
> warning: after 3 levels of inlining (potentially across files with Link Time Optimization), some common subexpression elimination, after hoisting this thing out of a loop and proving that these 13 pointers don't alias, we found a case where you're doing something undefined. This could either be because there is a bug in your code, or because you have macros and inlining and the invalid code is dynamically unreachable but we can't prove that it is dead.
> ```


**Unfortunately, we simply don't have the internal tracking infrastructure to produce this, and even if we did, the compiler doesn't have a user interface good enough to express this to the programmer.**

Ultimately, undefined behavior is valuable to the optimizer because it is saying "this operation is invalid - you can assume it never happens". In a case like "*P" this gives the optimizer the ability to reason that P cannot be NULL. In a case like "*NULL" (say, after some constant propagation and inlining), this allows the optimizer to know that the code must not be reachable. The important wrinkle here is that, because it cannot solve the halting problem, the compiler cannot know whether code is actually dead (as the C standard says it must be) or whether it is a bug that was exposed after a (potentially long) series of optimizations. Because there isn't a generally good way to distinguish the two, almost all of the warnings produced would be false positives (noise).

## Clang's Approach to Handling Undefined Behavior


Given the sorry state that we're in when it comes to undefined behavior, you might be wondering what Clang and LLVM are doing to try to improve the situation. I mentioned a couple of them already: the [Clang Static Analyzer](http://clang-analyzer.llvm.org/), [Klee project](http://klee.llvm.org/), and the `-fcatch-undefined-behavior` flag are useful tools for tracking down some classes of these bugs. The problem is that these aren't as widely used as the compiler is, so anything we can do directly in the compiler offers even higher goodness than doing it in these other tools. Keep in mind though that the compiler is limited by not having dynamic information and by being limited to what it can without burning lots of compile time.

Clang's first step to improve the world's code is to turn on a whole lot more warnings by default than other compilers do. While some developers are disciplined and build with "`-Wall -Wextra`" (for example), many people don't know about these flags or don't bother to pass them. Turning more warnings on by default catches more bugs more of the time.

The second step is that Clang generates warnings for many classes of undefined behavior (including dereference of null, oversized shifts, etc) that are obvious in the code to catch some common mistakes. Some of the caveats are mentioned above, but these seem to work well in practice.

The third step is that the LLVM optimizer generally takes much less liberty with undefined behavior than it could. Though the standard says that any instance of undefined behavior has completely unbound effects on the program, this is not a particularly useful or developer friendly behavior to take advantage of. Instead, the LLVM optimizer handles these optimizations in a few different ways (the links describe rules of LLVM IR, not C, sorry!):

1. Some cases of undefined behavior are silently transformed into implicitly trapping operations if there is a good way to do that. For example, with Clang, this C++ function:

   ```c
   int *foo(long x) {
     return new int[x];
   }
   ```

   compiles into this X86-64 machine code:

   ```assembly
   __Z3fool:
           movl $4, %ecx
           movq %rdi, %rax
           mulq %rcx
           movq $-1, %rdi        # Set the size to -1 on overflow
           cmovnoq %rax, %rdi    # Which causes 'new' to throw std::bad_alloc
           jmp __Znam
   ```

   instead of the code GCC produces:

   ```assembly
   __Z3fool:
           salq $2, %rdi
           jmp __Znam             # Security bug on overflow!
   ```

   >  The difference here is that we've decided to invest a few cycles in preventing a potentially serious integer overflow bug that can lead to buffer overflows and exploits (operator new is typically fairly expensive, so the overhead is almost never noticable). The GCC folks have been aware of this since at least 2005 but haven't fixed this at the time of this writing.

   

2. Arithmetic that [operates on undefined values](http://llvm.org/docs/LangRef.html#undefvalues) is considered to produce a undefined value instead of producing undefined behavior. The distinction is that undefined values can't format your hard drive or produce other undesirable effects. A useful refinement happens in cases where the arithmetic would produce the same output bits given any possible instance of the undefined value. For example, the optimizer assumes that the result of "undef & 1" has zeros for its top bits, treating only the low bit as undefined. This means that ((undef & 1) >> 1) is defined to be 0 in LLVM, not undefined.

3. Arithmetic that dynamically executes an undefined operation (such as a signed integer overflow) generates a logical [trap value](http://llvm.org/docs/LangRef.html#trapvalues) which poisons any computation based on it, but that does not destroy your entire program. This means that logic downstream from the undefined operation may be affected, but that your entire program isn't destroyed. This is why the optimizer ends up deleting code that operates on uninitialized variables, for example.

4. Stores to null and calls through null pointers are turned into a __builtin_trap() call (which turns into a trapping instruction like "ud2" on x86). These happen all of the time in optimized code (as the result of other transformations like inlining and constant propagation) and we used to just delete the blocks that contained them because they were "obviously unreachable".

   While (from a pedantic language lawyer standpoint) this is strictly true, we quickly learned that people do occasionally dereference null pointers, and having the code execution just fall into the top of the next function makes it very difficult to understand the problem. From the performance angle, the most important aspect of exposing these is to squash downstream code. Because of this, clang turns these into a runtime trap: if one of these is actually dynamically reached, the program stops immediately and can be debugged. The drawback of doing this is that we slightly bloat code by having these operations and having the conditions that control their predicates.

5. The optimizer does go to some effort to "do the right thing" when it is obvious what the programmer meant (such as code that does "*(int*)P" when P is a pointer to float). This helps in many common cases, but you really don't want to rely on this, and there are lots of examples that you might think are "obvious" that aren't after a long series of transformations have been applied to your code.

6. Optimizations that don't fall into any of these categories, such the zero_array and set/call examples in Part #1 are optimized as described, silently, without any indication to the user. We do this because we don't have anything useful to say, and it is very uncommon for (buggy) real-world code to be broken by these optimizations.


One major area of improvement we can make is with respect to trap insertion. I think it would be interesting to add an (off-by-default) warning flag that would cause the optimizer to warn whenever it generates a trap instruction. This would be extremely noisy for some codebases, but could be useful for others. The first limiting factor here is the infrastructure work to make the optimizer produce warnings: it doesn't have useful source code location information unless debugging information is turned on (but this could be fixed).

The other, more significant, limiting factor is that the warning wouldn't have any of the "tracking" information to be able to explain that an operation is the result of unrolling a loop three times and inlining it through four levels of function calls. At best we'll be able to point out the file/line/column of the original operation, which will be useful in the most trivial cases, but is likely to be extremely confusing in other cases. In any event, this hasn't been a high priority for us to implement because a) it isn't likely to give a good experience b) we won't be able to turn it on by default, and c) is a lot of work to implement.

## Using a Safer Dialect of C (and other options)


A final option you have if you don't care about "ultimate performance", is to use various compiler flags to enable dialects of C that eliminate these undefined behaviors. For example, using the `-fwrapv` flag eliminates undefined behavior that results from signed integer overflow (however, note that it does **not** eliminate possible integer overflow security vulnerabilities). The `-fno-strict-aliasing` flag disables Type Based Alias Analysis, so you are free to ignore these type rules. If there was demand, we could add a flag to Clang that implicitly zeros all local variables, one that inserts an "and" operation before each shift with a variable shift count, etc. Unfortunately, there is no tractable way to **completely** eliminate undefined behavior from C without breaking the ABI and completely destroying its performance. The other problem with this is that you're not writing C anymore, you're writing a similar, but non-portable dialect of C.

If writing code in a non-portable dialect of C isn't your thing, then the `-ftrapv`and `-fcatch-undefined-behavior` flags (along with the other tools mentioned before) can be useful weapons in your arsenal to track down these sorts of bugs. Enabling them in your debug builds can be a great way to find related bugs early. These flags can also be useful in production code if you are building security critical applications. While they provide no guarantee that they will find all bugs, they do find a useful subset of bugs.

Ultimately, the real problem here is that C just isn't a "safe" language and that (despite its success and popularity) many people do not really understand how the language works. In its decades of evolution prior to standardization in 1989, C migrated from being a "low level systems programming language that was a tiny layer above PDP assembly" to being a "low level systems programming language, trying to provide decent performance by *breaking many people's expectations*". On the one hand, these C "cheats" almost always work and code is generally more performant because of it (and in some cases, *much* more performant). On the other hand, the places where C cheats are often some of the most surprising to people and typically strike at the worst possible time.

C is much more than a portable assembler, sometimes in very surprising ways. I hope this discussion helps explain some of the issues behind undefined behavior in C, at least from a compiler implementer's viewpoint.

-[Chris Lattner](http://nondot.org/sabre/)

- [ Regular Expression Commands](https://blog.llvm.org/2011/04/regular-expression-commands.html) (`lldb`)

`2022-12-25 | 12:01p`

## Ask HN: Learn C in 2023?

[Ask HN: Learn C in 2023?](https://news.ycombinator.com/item?id=34105770)

209 points by [0x008](https://news.ycombinator.com/user?id=0x008) [2 days ago](https://news.ycombinator.com/item?id=34105770) | [hide](https://news.ycombinator.com/hide?id=34105770&auth=47dff431d72ed9cd886255087c90c146236a0380&goto=item%3Fid%3D34105770) | [past](https://hn.algolia.com/?query=Ask%20HN%3A%20Learn%20C%20in%202023%3F&type=story&dateRange=all&sort=byDate&storyText=false&prefix&page=0) | [favorite](https://news.ycombinator.com/fave?id=34105770&auth=47dff431d72ed9cd886255087c90c146236a0380) | [208 comments](https://news.ycombinator.com/item?id=34105770)

Hello,

I always have the itch to learn C, juts because it was such an influental language and when you are digging around a lot in GNU/linux you will have contact with it one way or another.

What is the way you would suggest learning C in 2023? I am already familiar with other languages (rust..) but would like to have an introduction to the C basics and program some hands-on projects in C, so I can learn the pecularities of the language.

[https://www.learn-c.org/](https://www.learn-c.org/)

If you have lots of time: [https://hal.inria.fr/hal-02383654](https://hal.inria.fr/hal-02383654)

If you can't be bothered reading a whole book: [https://matt.sh/howto-c](https://matt.sh/howto-c)

Exercises: [https://www.codestepbystep.com/problem/list/c](https://www.codestepbystep.com/problem/list/c) and [https://exercism.org/tracks/c](https://exercism.org/tracks/c)

Once you have syntax and basic algorithms down well, watch this, the only 2 hour YouTube video I'll ever recommend: [https://m.youtube.com/watch?v=443UNeGrFoM](https://m.youtube.com/watch?v=443UNeGrFoM)

Both r/cprogramming and r/C_programming are active, also full of lazy students trying to get people to do their homework. If you come by, describe your problem well with code. Say you're learning for yourself, not for school.

Together C & C++ is a good Discord if you prefer live chat: [https://discord.gg/tccpp](https://discord.gg/tccpp)

 

From the learn-c link:

```
    /* draws a point at 10, 5 */
    struct point p;
    p.x = 10;
    p.y = 5;
    draw(p);
```

I wish that would be a bit more 'modern', for instance:

```
    struct point p = {
        .x = 10,
        .y = 5
    };
    draw(p);
```

...or even:

```
    draw((struct point){ .x = 10, .y = 5 });
```

...main reason being that this avoids any accidents with unitialized data if the struct grows (the compound literal initialization makes sure that any struct members that are not mentioned are set to zero, while with the 'old school' approach you might end up with random junk in the struct).

[reply](https://news.ycombinator.com/reply?id=34107443&goto=item%3Fid%3D34105770%2334107443)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110397&how=up&auth=b4fe76604da2dfe732ca5c909955e52f2b18e78b&goto=item%3Fid%3D34105770#34110397)

[suprjami](https://news.ycombinator.com/user?id=suprjami) [1 day ago](https://news.ycombinator.com/item?id=34110397) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107443) | [next](https://news.ycombinator.com/item?id=34105770#34107236) [–]

  

I also love designated initialisers. C99 ftw.

[reply](https://news.ycombinator.com/reply?id=34110397&goto=item%3Fid%3D34105770%2334110397)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107236&how=up&auth=568bedf48cc4d85f724ab3a84e1f906d2f8784d6&goto=item%3Fid%3D34105770#34107236)

[voidhorse](https://news.ycombinator.com/user?id=voidhorse) [2 days ago](https://news.ycombinator.com/item?id=34107236) | [parent](https://news.ycombinator.com/item?id=34105770#34106174) | [prev](https://news.ycombinator.com/item?id=34105770#34107443) | [next](https://news.ycombinator.com/item?id=34105770#34107434) [–]

  

+1 to [https://hal.inria.fr/hal-02383654](https://hal.inria.fr/hal-02383654). It covers the language in extreme depth, and you likely won't remember it all on a first read but you'll walk away with a pretty good general understanding of the language. Plus, the author is an active contributor to the language.

I love C. As dangerous and as tedious it can be for large projects, I think it's my top choice for solo endeavors that don't require too many string manipulations. I don't think any language has come close to achieving the same balance of elegance and simplicity--maybe scheme? (zig comes close but has a substantial number of features that move it away from simplicity (e.g. payload capture))

[reply](https://news.ycombinator.com/reply?id=34107236&goto=item%3Fid%3D34105770%2334107236)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107434&how=up&auth=a71a4f443afbf987a3d4bff64151b18d2a996810&goto=item%3Fid%3D34105770#34107434)

[MonkeyClub](https://news.ycombinator.com/user?id=MonkeyClub) [2 days ago](https://news.ycombinator.com/item?id=34107434) | [parent](https://news.ycombinator.com/item?id=34105770#34106174) | [prev](https://news.ycombinator.com/item?id=34105770#34107236) | [next](https://news.ycombinator.com/item?id=34105770#34106815) [–]

  

> Once you have syntax and basic algorithms down well, watch this, the only 2 hour YouTube video I'll ever recommend: > [https://m.youtube.com/watch?v=443UNeGrFoM](https://m.youtube.com/watch?v=443UNeGrFoM)

Hard +1 on this!

I think Eskil Steenberg work in C is really top notch - the video is heavily recommended, as is an overview of his work in general, he does some awesome stuff with C.

[reply](https://news.ycombinator.com/reply?id=34107434&goto=item%3Fid%3D34105770%2334107434)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106815&how=up&auth=b041474eb978fa321e90b6b523db53125d732f20&goto=item%3Fid%3D34105770#34106815)

[mattpallissard](https://news.ycombinator.com/user?id=mattpallissard) [2 days ago](https://news.ycombinator.com/item?id=34106815) | [parent](https://news.ycombinator.com/item?id=34105770#34106174) | [prev](https://news.ycombinator.com/item?id=34105770#34107434) | [next](https://news.ycombinator.com/item?id=34105770#34107001) [–]

  

#c on libera is a genuinely helpful place full of neckbeards.

[reply](https://news.ycombinator.com/reply?id=34106815&goto=item%3Fid%3D34105770%2334106815)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107455&how=up&auth=d4ef57bd9f84cb8340715b8b28eca0fd44d5d65e&goto=item%3Fid%3D34105770#34107455)

[bogomipz](https://news.ycombinator.com/user?id=bogomipz) [2 days ago](https://news.ycombinator.com/item?id=34107455) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34106815) | [next](https://news.ycombinator.com/item?id=34105770#34110493) [–]

  

Sorry if this is silly question but what is libera?

[reply](https://news.ycombinator.com/reply?id=34107455&goto=item%3Fid%3D34105770%2334107455)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107584&how=up&auth=6cb7f4d4e0ec0304975b3a1c7a53aae12007e1c2&goto=item%3Fid%3D34105770#34107584)

[flippant](https://news.ycombinator.com/user?id=flippant) [2 days ago](https://news.ycombinator.com/item?id=34107584) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107455) | [next](https://news.ycombinator.com/item?id=34105770#34116032) [–]

  

It's an IRC network. [https://libera.chat/](https://libera.chat/)

[reply](https://news.ycombinator.com/reply?id=34107584&goto=item%3Fid%3D34105770%2334107584)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34116032&how=up&auth=7584ec4ca9de1568c56e8f0b4e53f1fba9dfc68f&goto=item%3Fid%3D34105770#34116032)

[cinntaile](https://news.ycombinator.com/user?id=cinntaile) [1 day ago](https://news.ycombinator.com/item?id=34116032) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107455) | [prev](https://news.ycombinator.com/item?id=34105770#34107584) | [next](https://news.ycombinator.com/item?id=34105770#34110493) [–]

  

A sibling comment already mentioned what it is so I'll not repeat that. Just some additional context if you're somewhat familiar with IRC. It's the spiritual successor of Freenode after it was taken over by a member of that community.

[reply](https://news.ycombinator.com/reply?id=34116032&goto=item%3Fid%3D34105770%2334116032)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110493&how=up&auth=2853efce2d0b13ef07a2ecd0d9fc716cfeefa37b&goto=item%3Fid%3D34105770#34110493)

[queuebert](https://news.ycombinator.com/user?id=queuebert) [1 day ago](https://news.ycombinator.com/item?id=34110493) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34106815) | [prev](https://news.ycombinator.com/item?id=34105770#34107455) | [next](https://news.ycombinator.com/item?id=34105770#34107001) [–]

  

#c-jeopardy is great fun as well.

[reply](https://news.ycombinator.com/reply?id=34110493&goto=item%3Fid%3D34105770%2334110493)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107001&how=up&auth=f55cf43dbe4eaafbf56c1c6db2bc64d46a4455f6&goto=item%3Fid%3D34105770#34107001)

[galangalalgol](https://news.ycombinator.com/user?id=galangalalgol) [2 days ago](https://news.ycombinator.com/item?id=34107001) | [parent](https://news.ycombinator.com/item?id=34105770#34106174) | [prev](https://news.ycombinator.com/item?id=34105770#34106815) | [next](https://news.ycombinator.com/item?id=34105770#34107177) [–]

  

Any recommendations on static analyzers, unit tests, coverage for c? Can clang tidy give hints about modern c they way it does about modern c++?

[reply](https://news.ycombinator.com/reply?id=34107001&goto=item%3Fid%3D34105770%2334107001)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107696&how=up&auth=4b9b865177ef710bd2c6d806578f65b46343f23e&goto=item%3Fid%3D34105770#34107696)

[gregfjohnson](https://news.ycombinator.com/user?id=gregfjohnson) [2 days ago](https://news.ycombinator.com/item?id=34107696) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107001) | [next](https://news.ycombinator.com/item?id=34105770#34108449) [–]

  

Not exactly your question, but valgrind is an amazing tool. It does execution-time analysis of your code and provides an in-depth report on memory usage. It will find and report memory leaks, use of uninitialized variables, bad memory references, etc. I love programming in C, but there are a lot of ways to corrupt memory, get segfaults, etc. valgrind is a useful and powerful way to guard against many of these kinds of problems. C++ is no better. The way I see it, programming in C is like lying on a bed of nails. Doable if you are very careful. With C++, they put a sheet over the bed of nails.

[reply](https://news.ycombinator.com/reply?id=34107696&goto=item%3Fid%3D34105770%2334107696)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107755&how=up&auth=4837b7135963351aee5e28dcada2f8181c84b65c&goto=item%3Fid%3D34105770#34107755)

[galangalalgol](https://news.ycombinator.com/user?id=galangalalgol) [2 days ago](https://news.ycombinator.com/item?id=34107755) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107696) | [next](https://news.ycombinator.com/item?id=34105770#34108449) [–]

  

I haven't used valgrind in years, more used to asan, but I've been told it can do things asan can't?

Edit: to stretch the metaphor, modern c++ is like a thick rubber pad on the bed of nails, but often on the job you are hot bunking and you don't know which c++ your predecessor used to make the bed... Or you are in a hurry and you aren't sure which one you grabbed.

[reply](https://news.ycombinator.com/reply?id=34107755&goto=item%3Fid%3D34105770%2334107755)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109480&how=up&auth=2d853405993f776fd8ce6c90bc1b43cc7353ba1a&goto=item%3Fid%3D34105770#34109480)

[mtklein](https://news.ycombinator.com/user?id=mtklein) [1 day ago](https://news.ycombinator.com/item?id=34109480) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107755) | [next](https://news.ycombinator.com/item?id=34105770#34107889) [–]

  

If you're used to ASAN, think of Valgrind like ASAN, TSAN, MSAN, and UBSAN all at once at the same time, and then has other facilities like profiling.

In any situation where I have control of how the code is built, I personally prefer to run those various SANs instead of Valgrind, usually three parallel builds: one with -fsanitize=address,integer,undefined, one -fsanitize=thread if more than one thread is involved in the project, and one -fsanitize=memory too if Linux is available. That gets you almost all the benefits of Valgrind faster than Valgrind.

But Valgrind is still an exceptionally powerful tool to learn and use. Its chief advantage to me is to be able to work with an unmodified binary. Its chief downside is it can be slow, and occasionally you may run into issues with support for more obscure instructions like vcvtps2ph.

MSAN is really the star of the sanitizer world, and is worth spinning up a Linux VM for. It tracks the initialized-ness of all your data and catches any use (really, branch) on anything uninitialized. I like to think of ASAN as "did I get my pointers right?" and MSAN as "did I get everything else right?".

If you're not familiar with integer sanitizer, it's like UBSAN except it checks for things that are often mistakes rather than things that are actually undefined. The most common situation I've found is code that's hashing data will often pop up a warning about unsigned integer overflow, which I choose to suppress as finely grained as possible with something like __attribute__((no_sanitize("unsigned-integer-overflow", "unsigned-shift-base"))). That way you can still find unintentional unsigned integer problems (e.g. with size_t).

[reply](https://news.ycombinator.com/reply?id=34109480&goto=item%3Fid%3D34105770%2334109480)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107889&how=up&auth=7f5cef9a4fb44364d3d6da4a0126607dd0fc9d8f&goto=item%3Fid%3D34105770#34107889)

[swinglock](https://news.ycombinator.com/user?id=swinglock) [2 days ago](https://news.ycombinator.com/item?id=34107889) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107755) | [prev](https://news.ycombinator.com/item?id=34105770#34109480) | [next](https://news.ycombinator.com/item?id=34105770#34108449) [–]

  

Ideally you use both, as both can find some more obscure issues the other may not. At least in the past. Either in isolation also works well, though Valgrind is really a suite of many tools.

Don't bother with this stuff (or with "modern C") if you're just beginning to learn the language. You don't need it for small programs. No one is going to yell at you if your first effort isn't perfect. And no one is going to yell at you if you read the "wrong" book, or try things out in the "wrong" order. Pick any learning resource and get started now. You'll probably get distracted and move on to something else. Come back to it later or not.

[reply](https://news.ycombinator.com/reply?id=34108449&goto=item%3Fid%3D34105770%2334108449)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34120717&how=up&auth=14189f37604be7c6faf0547a1685947f3347a2de&goto=item%3Fid%3D34105770#34120717)

[seba_dos1](https://news.ycombinator.com/user?id=seba_dos1) [21 hours ago](https://news.ycombinator.com/item?id=34120717) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108449) | [next](https://news.ycombinator.com/item?id=34105770#34109128) [–]

  

Do bother with this stuff, it's super helpful and helps you understand your mistakes. Writing C without using sanitizers is like playing a game on hard mode, even for simple programs (especially when you're still learning).

You may leave static analyzers for later though, since they're more likely to bother you with unimportant stuff - but not much later.

[reply](https://news.ycombinator.com/reply?id=34120717&goto=item%3Fid%3D34105770%2334120717)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109128&how=up&auth=dab63739b24a5f0a143414585659958bde5c1d05&goto=item%3Fid%3D34105770#34109128)

[galangalalgol](https://news.ycombinator.com/user?id=galangalalgol) [1 day ago](https://news.ycombinator.com/item?id=34109128) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108449) | [prev](https://news.ycombinator.com/item?id=34105770#34120717) | [next](https://news.ycombinator.com/item?id=34105770#34114855) [–]

  

Well C was my second language. I learned it on an amiga 500. But I quickly transitioned to a 386 (yes i know, a downgrade) but my dad put turbo c++ on it and I never wrote much ansi c after that. A bit for odd embedded targets, but mostly c++ my whole career. So I can write old fashioned c, and I don't start new projects in c, but writing c abi shims for c++ libraries or reading modern c is something I do with some frequency. Learning it would be fun maybe as I've never really used c99 or newer for anything.

[reply](https://news.ycombinator.com/reply?id=34109128&goto=item%3Fid%3D34105770%2334109128)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109354&how=up&auth=04ddd70a9a19834458cc98c13e2e25b76cb47715&goto=item%3Fid%3D34105770#34109354)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [1 day ago](https://news.ycombinator.com/item?id=34109354) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34109128) | [next](https://news.ycombinator.com/item?id=34105770#34114855) [–]

  

My comment is for the OP, not for you. Go nuts learning C tooling and have fun.

[reply](https://news.ycombinator.com/reply?id=34109354&goto=item%3Fid%3D34105770%2334109354)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34114855&how=up&auth=b7efa311d5a5494a52fdc4221d5c94f0e78eea66&goto=item%3Fid%3D34105770#34114855)

[kaba0](https://news.ycombinator.com/user?id=kaba0) [1 day ago](https://news.ycombinator.com/item?id=34114855) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108449) | [prev](https://news.ycombinator.com/item?id=34105770#34109128) | [next](https://news.ycombinator.com/item?id=34105770#34107578) [–]

  

I’m not sure, it does help to even just know about whether your program is faulty, and C doesn’t like telling it you, or in unintuitive ways only.

[reply](https://news.ycombinator.com/reply?id=34114855&goto=item%3Fid%3D34105770%2334114855)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107578&how=up&auth=5b94ae8f1b841d85c84dcc1df913c1862510dbfa&goto=item%3Fid%3D34105770#34107578)

[flohofwoe](https://news.ycombinator.com/user?id=flohofwoe) [2 days ago](https://news.ycombinator.com/item?id=34107578) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107001) | [prev](https://news.ycombinator.com/item?id=34105770#34108449) | [next](https://news.ycombinator.com/item?id=34105770#34110431) [–]

  

Set warning level to max, use clang static analyzer or clang tidy (I'm not quite sure how those two relate to each other though), and especially ASAN (Address Sanitizer) and UBSAN (Undefined Behaviour Sanitizer) (and if you write threaded code: TSAN (ThreadSanitizer).

One thing that Apple did right in Xcode and where other IDEs lagged behind is easy access to those tools from the UI.

[reply](https://news.ycombinator.com/reply?id=34107578&goto=item%3Fid%3D34105770%2334107578)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107794&how=up&auth=1e009f67450c3601d7e52789fdbfa7e8faf5166d&goto=item%3Fid%3D34105770#34107794)

[galangalalgol](https://news.ycombinator.com/user?id=galangalalgol) [2 days ago](https://news.ycombinator.com/item?id=34107794) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107578) | [next](https://news.ycombinator.com/item?id=34105770#34110431) [–]

  

That sounds like all the things I do for c++. One thing I really like is that clang tidy can flag the "old way" to do stuff. Was hoping it had that for c.

[reply](https://news.ycombinator.com/reply?id=34107794&goto=item%3Fid%3D34105770%2334107794)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110431&how=up&auth=30634bb4331235f072ba94b5bcbaffef21233c5d&goto=item%3Fid%3D34105770#34110431)

[suprjami](https://news.ycombinator.com/user?id=suprjami) [1 day ago](https://news.ycombinator.com/item?id=34110431) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107001) | [prev](https://news.ycombinator.com/item?id=34105770#34107578) | [next](https://news.ycombinator.com/item?id=34105770#34107502) [–]

  

You can't rely on one tool, different ones catch different things. No one tool is best.

I usually use "-Wall -Wextra -Wpedantic", the GCC static analyser "-fanalyzer", cppcheck, and Clang scan-build.

That's good enough for my personal projects.

I have a couple of blog posts about this:

[https://superjamie.github.io/2022/05/24/c-static-analysis-to...](https://superjamie.github.io/2022/05/24/c-static-analysis-tools)

[https://superjamie.github.io/2022/11/04/c-static-analysis-to...](https://superjamie.github.io/2022/11/04/c-static-analysis-tools-part-2)

[reply](https://news.ycombinator.com/reply?id=34110431&goto=item%3Fid%3D34105770%2334110431)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107502&how=up&auth=88e55a5adb8ae5245a594b0fa0e5b993bd5c306c&goto=item%3Fid%3D34105770#34107502)

[cillian64](https://news.ycombinator.com/user?id=cillian64) [2 days ago](https://news.ycombinator.com/item?id=34107502) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107001) | [prev](https://news.ycombinator.com/item?id=34105770#34110431) | [next](https://news.ycombinator.com/item?id=34105770#34107177) [–]

  

One useful tool is running clang or gcc with —Werror -Wpedantic which will point out a lot of things which are useful to know about, even if occasionally a false positive.

[reply](https://news.ycombinator.com/reply?id=34107502&goto=item%3Fid%3D34105770%2334107502)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107177&how=up&auth=3f31f0de18cfbf9ae9a4851f87e7496ef4b0466a&goto=item%3Fid%3D34105770#34107177)

[mozball](https://news.ycombinator.com/user?id=mozball) [2 days ago](https://news.ycombinator.com/item?id=34107177) | [parent](https://news.ycombinator.com/item?id=34105770#34106174) | [prev](https://news.ycombinator.com/item?id=34105770#34107001) | [next](https://news.ycombinator.com/item?id=34105770#34107405) [–]

  

What is your opinion on K&R?

[reply](https://news.ycombinator.com/reply?id=34107177&goto=item%3Fid%3D34105770%2334107177)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107373&how=up&auth=050af7a6fe1996dd6d94171355fc042fdd495d0e&goto=item%3Fid%3D34105770#34107373)

[fn-mote](https://news.ycombinator.com/user?id=fn-mote) [2 days ago](https://news.ycombinator.com/item?id=34107373) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107177) | [next](https://news.ycombinator.com/item?id=34105770#34108298) [–]

  

It's great for what it is, but the language and style have come a long way since then. It would be better to learn "Modern C" (this is the title of a book but also at least learn C99 not K&R C).

E.g., people like the book by Jens Gustedt. I'm ambivalent about it, but it sure machine-guns the important points.

[1] [https://hal.inria.fr/hal-02383654/file/ModernC.pdf](https://hal.inria.fr/hal-02383654/file/ModernC.pdf)

[reply](https://news.ycombinator.com/reply?id=34107373&goto=item%3Fid%3D34105770%2334107373)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107614&how=up&auth=5123da438a42912dd5fdf5df542aa524159b6ad2&goto=item%3Fid%3D34105770#34107614)

[Koshkin](https://news.ycombinator.com/user?id=Koshkin) [2 days ago](https://news.ycombinator.com/item?id=34107614) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107373) | [next](https://news.ycombinator.com/item?id=34105770#34108298) [–]

  

> _not K&R C_

K&R the 2nd edition of book teaches ANSI C. (It is still the best intro into the basics of C.)

[reply](https://news.ycombinator.com/reply?id=34107614&goto=item%3Fid%3D34105770%2334107614)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108298&how=up&auth=f48461e9325c9f7548da72695848ba815468cc14&goto=item%3Fid%3D34105770#34108298)

[stusmall](https://news.ycombinator.com/user?id=stusmall) [2 days ago](https://news.ycombinator.com/item?id=34108298) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107177) | [prev](https://news.ycombinator.com/item?id=34105770#34107373) | [next](https://news.ycombinator.com/item?id=34105770#34107485) [–]

  

Not the person you are replying to, but here are my thoughts. It sounds like OP is looking to learn C for historical and contextual reasons, and for that K&R is perfect. The history of C is very K&R-centric.

For modern, practical C programing it is missing a lot though. It is a short, simple, from the original authors introduction to the language itself. There is nothing about how to test your code, how to set it up for long term maintainability, how to mitigate some of the security downfalls or even any of the features added since C89.

It's a good technical book. A great one for it's time and subject matter. But that caveat is important, it needs to be taken in context of the time. It shouldn't be the first, and absolutely not the only, book you read if you are writing modern production C in 2023.

[reply](https://news.ycombinator.com/reply?id=34108298&goto=item%3Fid%3D34105770%2334108298)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108384&how=up&auth=1d3ce05825c45f115be8328c76b87a0e1a550878&goto=item%3Fid%3D34105770#34108384)

[asveikau](https://news.ycombinator.com/user?id=asveikau) [2 days ago](https://news.ycombinator.com/item?id=34108384) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108298) | [next](https://news.ycombinator.com/item?id=34105770#34107485) [–]

  

> or how to mitigate some of the security downfalls.

I think I'd go further, there are places where it advocates downright unsafe practices. Eg. It's been a long time, but I thought I remember it introducing gets() unironically, without any caveat. For those that do not know, there is no safe use of that interface possible.

[reply](https://news.ycombinator.com/reply?id=34108384&goto=item%3Fid%3D34105770%2334108384)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108475&how=up&auth=657cd40f98b30c36e3c61cbb748fc2b8bd911f94&goto=item%3Fid%3D34105770#34108475)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [2 days ago](https://news.ycombinator.com/item?id=34108475) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108384) | [next](https://news.ycombinator.com/item?id=34105770#34107485) [–]

  

If you're just beginning with C, it doesn't matter. If anything, it's better to be exposed to these unsafe practices and make mistakes with them in the very safe environment of your own home, where absolutely no one cares what you're doing.

K&R C is an excellent learning resource primarily because it's extremely strong pedagogically. Read K&R and then move onto something else later when you're done with it. If you eventually get a job doing something with C, then you'll learn on the job what's important to the job you're doing.

[reply](https://news.ycombinator.com/reply?id=34108475&goto=item%3Fid%3D34105770%2334108475)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109028&how=up&auth=15c3243dda0913ac981f73fe4179ae28a7cb0027&goto=item%3Fid%3D34105770#34109028)

[asveikau](https://news.ycombinator.com/user?id=asveikau) [1 day ago](https://news.ycombinator.com/item?id=34109028) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108475) | [next](https://news.ycombinator.com/item?id=34105770#34109397) [–]

  

I've taught C at university level, and used to give crash courses to new college grads and interns in industry.

The one thing I impart the most from the beginning is to develop good instincts for secure interfaces. You should be able to look at function signatures and see alarm bells for unsafe practices. For someone skilled at C this is second nature, you can see bad practices from a distance, often at the function signature level.

I think it would be unwise to neglect this.

[reply](https://news.ycombinator.com/reply?id=34109028&goto=item%3Fid%3D34105770%2334109028)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109817&how=up&auth=56b473ab67124cf7f1c27ecd3d527024ec687fb5&goto=item%3Fid%3D34105770#34109817)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [1 day ago](https://news.ycombinator.com/item?id=34109817) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34109028) | [next](https://news.ycombinator.com/item?id=34105770#34109397) [–]

  

It's unwise to neglect it in the long run. It's perfectly reasonable to ignore it initially. There are different valid pedagogical approaches here. Your approach may be suitable for the type of student you were interacting with, but the profile you're familiar with doesn't represent every person who might be interested in learning C.

[reply](https://news.ycombinator.com/reply?id=34109817&goto=item%3Fid%3D34105770%2334109817)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110202&how=up&auth=79fb8adaa6102adf624e5298d120e74270b4aedf&goto=item%3Fid%3D34105770#34110202)

[asveikau](https://news.ycombinator.com/user?id=asveikau) [1 day ago](https://news.ycombinator.com/item?id=34110202) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34109817) | [next](https://news.ycombinator.com/item?id=34105770#34109397) [–]

  

I think a good approach is to have someone able to review code. So they can experiment and make all the common mistakes, then have feedback on it.

[reply](https://news.ycombinator.com/reply?id=34110202&goto=item%3Fid%3D34105770%2334110202)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109397&how=up&auth=f7b2d9bd7ecfbc6425ad25c2b64389d81ece10ee&goto=item%3Fid%3D34105770#34109397)

[stusmall](https://news.ycombinator.com/user?id=stusmall) [1 day ago](https://news.ycombinator.com/item?id=34109397) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108475) | [prev](https://news.ycombinator.com/item?id=34105770#34109028) | [next](https://news.ycombinator.com/item?id=34105770#34108549) [–]

  

C is one of those things that unless you have proper testing, tooling and linting you don't know if you code is incorrect. It may have subtle memory bugs, rely on undefined behavior or crumple in exploitable ways under bad input. Having a feedback loop is really important in learning and improving. Undefined behavior helps subvert it.

For the exact example of above, how would a developer learn _not_ to use gets? The best place would be a section on red flags in API design in the book. If you don't have that you need to find it through experience. Lints, fuzzing, sanitizers can help, but once again, none of this is covered in K&R. Some of it just didn't exist at the time or if it did, it was very primitive compared to what we have today.

Developers could pick up bad habits and be productive for years with them before they get bit by them. I know I did. I look back at some of the practices I picked up when first learning and cringe a bit. Especially since I think a fair amount of that code is still in production.

[reply](https://news.ycombinator.com/reply?id=34109397&goto=item%3Fid%3D34105770%2334109397)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109644&how=up&auth=ebb3dd1880a435b95d0a810720b27d6e382eba34&goto=item%3Fid%3D34105770#34109644)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [1 day ago](https://news.ycombinator.com/item?id=34109644) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34109397) | [next](https://news.ycombinator.com/item?id=34105770#34108549) [–]

  

All of this is true, but if you are just starting with C, none of it matters.

If you continue using C, then you should eventually learn this stuff, but to overburden someone with a list of 100 things they must do right is the surest way to frustrate them to the point where they give up.

If OP eventually writes C code that goes into production for a company, it is the company's responsibility to have procedures and safeguards which ensure that OP writes reasonable code. Ideally, one such safeguard will be a senior colleague to introduce them to the thornier parts of C.

To expect a beginner to learn all of the finer points of a topic from the start is bad pedagogy.

On the other hand, K&R provides a set of exercises which are fun, graded, interesting, and challenging. By the time you finish them, you have a working knowledge of a large and useful subset of C. You can then proceed to reading another book (or consulting some other reference). Invariably, you will find that they disagree with K&R to some extent. Learning to spot these differences and begin to understand how C has evolved is itself an important lesson.

[reply](https://news.ycombinator.com/reply?id=34109644&goto=item%3Fid%3D34105770%2334109644)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111316&how=up&auth=47dcb50c524de8c2eeb9b587fe28d54d026c7e25&goto=item%3Fid%3D34105770#34111316)

[stusmall](https://news.ycombinator.com/user?id=stusmall) [1 day ago](https://news.ycombinator.com/item?id=34111316) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34109644) | [next](https://news.ycombinator.com/item?id=34105770#34108549) [–]

  

> If OP eventually writes C code that goes into production for a company, it is the company's responsibility to have procedures and safeguards which ensure that OP writes reasonable code. Ideally, one such safeguard will be a senior colleague to introduce them to the thornier parts of C.

You've mentioned something like this in reply to a few comments in here. This is a strange argument. It isn't too much to ask of a book to lay down more of a foundation so you don't need a 1 on 1 tutor later to correct misconceptions.

I get people like K&R because it is short and easy to read. Some subjects aren't best served by short and easy to read. The subject was well served by it when C was new and novel. We live in a very different world.

[reply](https://news.ycombinator.com/reply?id=34111316&goto=item%3Fid%3D34105770%2334111316)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111794&how=up&auth=dbb32634245f8850db2359749e1d5b070188ca53&goto=item%3Fid%3D34105770#34111794)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [1 day ago](https://news.ycombinator.com/item?id=34111794) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34111316) | [next](https://news.ycombinator.com/item?id=34105770#34108549) [–]

  

You don't actually seem to be addressing the quoted text... but I can reply to your arguments.

Sure, it would be ideal if there were a new book on C which did a better job of explaining some of the developments since K&R 2ed was written. If that book exists, I'm not aware of it. All of the books I've seen recommended are either too long, or poorly written. K&R's strong advantage isn't that it's short, it's that it's very strong pedagogically.

A "how do I learn C?" thread comes up on HN about once a month. In each of these threads, people will steer people clear of K&R, quoting something or another about how K&R isn't appropriate for this day and age (note: you do not explain why you think it isn't appropriate!). The tacit assumption here is that you can only read one book, or that if you read a book which isn't "perfect", you are somehow tainting yourself with corrupted knowledge. I think this is nonsense. An important meta-skill for learning is the ability to discern between good and bad information. The only way to build that skill is by consulting multiple sources.

Note that in my actual reply to OP, I don't specifically advocate for reading K&R, but instead to take a multipronged approach where you primarily draw from whatever learning medium you're most comfortable with. Most people I know that are in their early 20s or younger prefer learning from videos (I'm a math teacher and have lots of experience with people this age).

[reply](https://news.ycombinator.com/reply?id=34111794&goto=item%3Fid%3D34105770%2334111794)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34112348&how=up&auth=dadf6c820b4d1e923b0d38e1cb4932c0413cdea3&goto=item%3Fid%3D34105770#34112348)

[stusmall](https://news.ycombinator.com/user?id=stusmall) [1 day ago](https://news.ycombinator.com/item?id=34112348) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34111794) | [next](https://news.ycombinator.com/item?id=34105770#34108549) [–]

  

I think you've got a bit of a strawman in your head. Look back up at some of my comments you are replying to, namely the top one in this chain. A friend describes this as "violent agreement." I agree that it can be best to have a multipronged approach. That's why I said "It's a good technical book. A great one for it's time and subject matter. But that caveat is important, it needs to be taken in context of the time. It shouldn't be the first, and absolutely not the only, book you read if you are writing modern production C in 2023." I also address its short comings in that post.

[reply](https://news.ycombinator.com/reply?id=34112348&goto=item%3Fid%3D34105770%2334112348)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34112408&how=up&auth=f246da85a35b53235091acae1878d6eab9b175bc&goto=item%3Fid%3D34105770#34112408)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [1 day ago](https://news.ycombinator.com/item?id=34112408) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34112348) | [next](https://news.ycombinator.com/item?id=34105770#34108549) [–]

  

I didn't read your other post. Glad we're in violent agreement!

[reply](https://news.ycombinator.com/reply?id=34112408&goto=item%3Fid%3D34105770%2334112408)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108549&how=up&auth=3ddad60d669b5e8c97f3639109d85c56a2bf6317&goto=item%3Fid%3D34105770#34108549)

[convolvatron](https://news.ycombinator.com/user?id=convolvatron) [2 days ago](https://news.ycombinator.com/item?id=34108549) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108475) | [prev](https://news.ycombinator.com/item?id=34105770#34109397) | [next](https://news.ycombinator.com/item?id=34105770#34107485) [–]

  

I kinda disagree. global hidden state under APIs is a really awful idea to put in someone's head.

I think it also presents the notion to neophytes that there is magic under there

[reply](https://news.ycombinator.com/reply?id=34108549&goto=item%3Fid%3D34105770%2334108549)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108732&how=up&auth=975616752fc7c5c40752fdd87a0e591a70ba3d28&goto=item%3Fid%3D34105770#34108732)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [2 days ago](https://news.ycombinator.com/item?id=34108732) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108549) | [next](https://news.ycombinator.com/item?id=34105770#34107485) [–]

  

Unfortunately, global hidden state is an unavoidable fact of programming. I don't see how one could be very good at programming if they were kept blissfully ignorant of it. Sometimes there _is_ magic under there, and you just have to do your best to empirically figure out how that magic works.

[reply](https://news.ycombinator.com/reply?id=34108732&goto=item%3Fid%3D34105770%2334108732)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110711&how=up&auth=161cf67aa16e4807c71256adc3a02f5c434f37ad&goto=item%3Fid%3D34105770#34110711)

[asveikau](https://news.ycombinator.com/user?id=asveikau) [1 day ago](https://news.ycombinator.com/item?id=34110711) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108732) | [next](https://news.ycombinator.com/item?id=34105770#34107485) [–]

  

libc has a few interfaces with implicit global storage. This leads to lack of thread safety and other issues. A good library will let the caller provide storage for these cases, or provide the caller a handle with which to pass it later/free it.

A trivial example would be strtok vs strtok_r. There are some, like, say, asctime(), that operate against a global string buffer.

Sometimes interfaces "look like" they're returning a global buffer, for example readdir(), but that structure can hang off the DIR* struct ... Nonetheless there is readdir_r()

[reply](https://news.ycombinator.com/reply?id=34110711&goto=item%3Fid%3D34105770%2334110711)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107485&how=up&auth=3ad48ec2bd84fe0f8835a49c64391a48a48629d4&goto=item%3Fid%3D34105770#34107485)

[flohofwoe](https://news.ycombinator.com/user?id=flohofwoe) [2 days ago](https://news.ycombinator.com/item?id=34107485) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107177) | [prev](https://news.ycombinator.com/item?id=34105770#34108298) | [next](https://news.ycombinator.com/item?id=34105770#34110714) [–]

  

It's much too awkward IMHO, C99 and onward almost feels like a whole new language (mainly because of struct literal and designated init support).

[reply](https://news.ycombinator.com/reply?id=34107485&goto=item%3Fid%3D34105770%2334107485)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108935&how=up&auth=ad640cc12a64eaf9495e86077088aa8c15836699&goto=item%3Fid%3D34105770#34108935)

[rightbyte](https://news.ycombinator.com/user?id=rightbyte) [1 day ago](https://news.ycombinator.com/item?id=34108935) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107485) | [next](https://news.ycombinator.com/item?id=34105770#34110714) [–]

  

Not being forced to declare symbols in the beginning of the scope is the real game changer I say. K&R or C89 without e.g. GNU extensions is really annoying to use.

[reply](https://news.ycombinator.com/reply?id=34108935&goto=item%3Fid%3D34105770%2334108935)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34118732&how=up&auth=ca1e6624c55c793ad5de1da1c3111e5f0c0b3e2a&goto=item%3Fid%3D34105770#34118732)

[bluecalm](https://news.ycombinator.com/user?id=bluecalm) [1 day ago](https://news.ycombinator.com/item?id=34118732) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34108935) | [next](https://news.ycombinator.com/item?id=34105770#34110714) [–]

  

Idk about that. I think it's code with everything declared at the beginning of the scope is easier to read and understand. Every time I see code written in this style (Git, Linux) it's easier to build a mental model of what is happening.

[reply](https://news.ycombinator.com/reply?id=34118732&goto=item%3Fid%3D34105770%2334118732)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110714&how=up&auth=03b6f1b0bd6eb3e26267263aacbe8a1cc30e97fc&goto=item%3Fid%3D34105770#34110714)

[sn9](https://news.ycombinator.com/user?id=sn9) [1 day ago](https://news.ycombinator.com/item?id=34110714) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107177) | [prev](https://news.ycombinator.com/item?id=34105770#34107485) | [next](https://news.ycombinator.com/item?id=34105770#34110474) [–]

  

It's one of the best examples of technical writing in the field. A good introduction to the basics of the language. Read it for culture.

Read something else for safe C writing in a modern environment. Books like _Modern C_ and _Effective C_ are good starts.

[reply](https://news.ycombinator.com/reply?id=34110714&goto=item%3Fid%3D34105770%2334110714)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110474&how=up&auth=573cb17deba50b2c9fc1f6e1dd2ac01ec6bc6d03&goto=item%3Fid%3D34105770#34110474)

[suprjami](https://news.ycombinator.com/user?id=suprjami) [1 day ago](https://news.ycombinator.com/item?id=34110474) | [root](https://news.ycombinator.com/item?id=34105770#34106174) | [parent](https://news.ycombinator.com/item?id=34105770#34107177) | [prev](https://news.ycombinator.com/item?id=34105770#34110714) | [next](https://news.ycombinator.com/item?id=34105770#34107405) [–]

  

It's wonderfully written. I'm a big fan of both Dennis Ritchie and Brian Kernighan personally.

However the book was last updated in 1988 and I prefer C99+ for almost every feature it adds. If I'm going to recommend just one book it will be something newer.

[reply](https://news.ycombinator.com/reply?id=34110474&goto=item%3Fid%3D34105770%2334110474)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107405&how=up&auth=ed2dc0bb17336785932a3e4a2863e82ddf83e054&goto=item%3Fid%3D34105770#34107405)

[andreygrehov](https://news.ycombinator.com/user?id=andreygrehov) [2 days ago](https://news.ycombinator.com/item?id=34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34106174) | [next](https://news.ycombinator.com/item?id=34105770#34107476) [–]

  

I'm not a C engineer, but I think I have an interesting recommendation to consider.

Since you are already familiar with other languages, you obviously don't need to know the basics of C. The basics are the same everywhere. Instead, you'd likely want to build that mindset on how to build good software with C.

I suggest you read other people's code and try to deep dive into the whys.

I heard Redis is a well-written software.

```
    git clone https://github.com/redis/redis.git
    git log –reverse
    git checkout ed9b544e...
```

We can now see the very first commit made by Salvatore Sanfilippo. From here, we can start exploring. By looking at the commit diffs and git log messages, you'll be able to understand the motivation behind the changes.

When you read a book, the most interesting parts are usually after all the theory is explained at the beginning. Somewhere in the middle of the book they start showing you code samples and that's when things become really interesting.

What I like about the reverse git-log approach is that it immediately throws you into the middle section of the book. This approach is both fun and eye opening.

Good luck!

[reply](https://news.ycombinator.com/reply?id=34107405&goto=item%3Fid%3D34105770%2334107405)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108518&how=up&auth=40679bbf08e44e6e459c7ac1e6f62a95136b1918&goto=item%3Fid%3D34105770#34108518)

[asveikau](https://news.ycombinator.com/user?id=asveikau) [2 days ago](https://news.ycombinator.com/item?id=34108518) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [next](https://news.ycombinator.com/item?id=34105770#34108350) [–]

  

For me, when I learned more than 20 years ago, reading source and manpages from Linux and other Unix like projects was a source of inspiration.

I'd recommend OpenBSD libc. [https://github.com/openbsd/src/tree/master/lib/libc](https://github.com/openbsd/src/tree/master/lib/libc)

String functions especially are a good way to get into the C way of thinking. Since C's approach to strings is unique.

I also learned a lot by reading manpages of libc functions or Unix utilities and thinking about how they were implemented, and writing my own little versions.

[reply](https://news.ycombinator.com/reply?id=34108518&goto=item%3Fid%3D34105770%2334108518)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108350&how=up&auth=57763d1bcb851efff9c03f2c7c699a4e2239928e&goto=item%3Fid%3D34105770#34108350)

[ozim](https://news.ycombinator.com/user?id=ozim) [2 days ago](https://news.ycombinator.com/item?id=34108350) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34108518) | [next](https://news.ycombinator.com/item?id=34105770#34113291) [–]

  

Well I tried it and ed9b544e... is god awful, documentation has so many typos. God file 'redis.c' with 3k lines having even commented out code and some comments that I would not pass my code review.

But I love it and it is great recommendation.

Going to move with commits to see how it unfolds :)

[reply](https://news.ycombinator.com/reply?id=34108350&goto=item%3Fid%3D34105770%2334108350)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34113291&how=up&auth=c2ea17bcc6303a74a988af6e6dbae36a937070ef&goto=item%3Fid%3D34105770#34113291)

[stcroixx](https://news.ycombinator.com/user?id=stcroixx) [1 day ago](https://news.ycombinator.com/item?id=34113291) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34108350) | [next](https://news.ycombinator.com/item?id=34105770#34108182) [–]

  

I think this is important just to get the idioms down that you otherwise only get on the job. Could pick almost any big active C project. I used to do this and I’d see some construct I didn’t understand and consult my K&R book to learn it. I’d do this in parallel with a book though for sure, K&R preferably. C is so basic that if you pursue it further it will be about learning the projects/companies way of using C.

[reply](https://news.ycombinator.com/reply?id=34113291&goto=item%3Fid%3D34105770%2334113291)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108182&how=up&auth=7af32b11a60d753cebeb1b3c1ecb3f6071ea63c4&goto=item%3Fid%3D34105770#34108182)

[brezelgoring](https://news.ycombinator.com/user?id=brezelgoring) [2 days ago](https://news.ycombinator.com/item?id=34108182) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34113291) | [next](https://news.ycombinator.com/item?id=34105770#34109537) [–]

  

I've got a single digit number of contributions to projects that I like, mostly because of how daunting it is to familiarize oneself with a codebase, and this 'git log -reverse' trick seems incredible as an entry point to a new project.

That being said, how do you feel about using this technique for repos with 2k+ commits? It seems unrealistic to read all of them, or am I being short-sighted here? I get demotivated just thinking about it.

[reply](https://news.ycombinator.com/reply?id=34108182&goto=item%3Fid%3D34105770%2334108182)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108398&how=up&auth=aa07d17011761e7812808914c4bdbb2f1ab92474&goto=item%3Fid%3D34105770#34108398)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [2 days ago](https://news.ycombinator.com/item?id=34108398) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34108182) | [next](https://news.ycombinator.com/item?id=34105770#34109293) [–]

  

Reading 2k+ commits is a big waste of time.

Forcing yourself to read every commit to a project you "like" in order contribute to it is putting the cart before the horse. Just because you "like" a project doesn't mean you will have the skills to contribute usefully to it. This isn't a bad thing. I "like" GCC, but I don't have the skills to contribute to it. I'm not losing any sleep over this, and I'm definitely not thinking about reading every commit to GCC in chronological order.

If you use a project, you might find a bug or some deficiency in it. Then it might happen that you have the skills to quickly fix this problem. Either that, or this bug causes you so many problems or slows you down so much that it becomes worth it for you to develop the skills to fix the problem.

[reply](https://news.ycombinator.com/reply?id=34108398&goto=item%3Fid%3D34105770%2334108398)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34129000&how=up&auth=d65ce0cc591c17ac17f677dc91545206afa5b9c3&goto=item%3Fid%3D34105770#34129000)

[brezelgoring](https://news.ycombinator.com/user?id=brezelgoring) [38 minutes ago](https://news.ycombinator.com/item?id=34129000) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34108398) | [next](https://news.ycombinator.com/item?id=34105770#34109293) [–]

  

I like this approach to FOSS contributions, thanks for the answer, sfpotter.

[reply](https://news.ycombinator.com/reply?id=34129000&goto=item%3Fid%3D34105770%2334129000)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109293&how=up&auth=33f9a3ce6a01885788d387a006ea6d9852beccdb&goto=item%3Fid%3D34105770#34109293)

[baby](https://news.ycombinator.com/user?id=baby) [1 day ago](https://news.ycombinator.com/item?id=34109293) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34108182) | [prev](https://news.ycombinator.com/item?id=34105770#34108398) | [next](https://news.ycombinator.com/item?id=34105770#34109537) [–]

  

I'm not sure why GP's post is upvoted so much, IMO it's a bad way to learn. I would instead find an entry point, a self-contained task that doesn't touch a lot of the codebase that you can use to learn something. Then, as you get more comfortable with some part of the codebase you can expand your scope by looking at tasks that touch other things. In general reading code is 50% of the work, the other 50% is to play with the codebase or try to make changes to it and see what happens.

[reply](https://news.ycombinator.com/reply?id=34109293&goto=item%3Fid%3D34105770%2334109293)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34128948&how=up&auth=6eefe30c9e341cad0d203cb74c55aaeb5d94f066&goto=item%3Fid%3D34105770#34128948)

[brezelgoring](https://news.ycombinator.com/user?id=brezelgoring) [42 minutes ago](https://news.ycombinator.com/item?id=34128948) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109293) | [next](https://news.ycombinator.com/item?id=34105770#34109537) [–]

  

That sounds reasonable, and I do agree that inching forward seems like a more viable approach to contributing in FOSS.

Thanks for the answer.

[reply](https://news.ycombinator.com/reply?id=34128948&goto=item%3Fid%3D34105770%2334128948)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109537&how=up&auth=6be8fd638bc9917406fd827cd0eb2d8c9f094a53&goto=item%3Fid%3D34105770#34109537)

[daniele_dll](https://news.ycombinator.com/user?id=daniele_dll) [1 day ago](https://news.ycombinator.com/item?id=34109537) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34108182) | [next](https://news.ycombinator.com/item?id=34105770#34109271) [–]

  

Fully agree with the principle but maybe let's not call Redis "well-written" :)

Redis is a great piece of software because of how much it impacted the "evolution" of Internet not because of the code quality, Redis's code resemble spaghetti code, a lot.

I will push myself and say that nowadays people would get fired for writing code like that!

And before someone start to throw stones (just for the sake of it), here a few examples: - Redis source code is FULL of GIANT files with thousands and thousands of lines of code, decoupling and separation of concerns are basic things! A few examples -- [https://github.com/redis/redis/blob/unstable/src/module.c](https://github.com/redis/redis/blob/unstable/src/module.c) -- [https://github.com/redis/redis/blob/unstable/src/cluster.c](https://github.com/redis/redis/blob/unstable/src/cluster.c) -- [https://github.com/redis/redis/blob/unstable/src/redis-cli.c](https://github.com/redis/redis/blob/unstable/src/redis-cli.c) -- [https://github.com/redis/redis/blob/unstable/src/networking....](https://github.com/redis/redis/blob/unstable/src/networking.c)

- talking about auto generated, the repo contains commited auto generated code, a valid reason though really slips my mind (yes, of course I can think of a few, but ... really? :)))

- [https://github.com/redis/redis/blob/9c7c6924a019b902996fc4b6...](https://github.com/redis/redis/blob/9c7c6924a019b902996fc4b608541f475daa649d/src/server.h#L1257) ... that struct feels like an "Italian minestrone", basically a kind of soup where you throw in, literally, all the vegetables you have around and it's great for a soup, a bit less for software engineering though :)

- the naming structure is a bit random at times, there is a mix of camel case, camel case mixed with underscores from time to time, pascal case

This kind of staff might be alright for a 0.3 or 0.4, maybe even for a 1.0, definitely not for a 7.x that has had 15 years to evolve :/ Even in cachegrand, that hasn't got to a v0.2 yet, I tried to avoid these kinda of dramas, in C is quick and easy to get there if you are not careful.

There are plenty of other small things but most of the big dramas should be there.

[reply](https://news.ycombinator.com/reply?id=34109537&goto=item%3Fid%3D34105770%2334109537)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109758&how=up&auth=0f350671309c47d351e3798092d9247bfa25d898&goto=item%3Fid%3D34105770#34109758)

[palmtree3000](https://news.ycombinator.com/user?id=palmtree3000) [1 day ago](https://news.ycombinator.com/item?id=34109758) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109537) | [next](https://news.ycombinator.com/item?id=34105770#34109715) [–]

  

Re: auto generated code: it saves ci time which is a pretty huge benefit. You still need to regenerate the code to make sure it's up to date, but you can do that in parallel with tests that rely on that code.

[reply](https://news.ycombinator.com/reply?id=34109758&goto=item%3Fid%3D34105770%2334109758)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109889&how=up&auth=b6c3dd117ca1c7369eefbe6aef0a002fd1bff01a&goto=item%3Fid%3D34105770#34109889)

[daniele_dll](https://news.ycombinator.com/user?id=daniele_dll) [1 day ago](https://news.ycombinator.com/item?id=34109889) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109758) | [next](https://news.ycombinator.com/item?id=34105770#34109715) [–]

  

This is a valid reason, but makes sense for code bases where the generation takes tens of minutes or more not seconds.

[reply](https://news.ycombinator.com/reply?id=34109889&goto=item%3Fid%3D34105770%2334109889)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109715&how=up&auth=ca3ab8c30ead4d56daea1a142ef68941fd15a0ea&goto=item%3Fid%3D34105770#34109715)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [1 day ago](https://news.ycombinator.com/item?id=34109715) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109537) | [prev](https://news.ycombinator.com/item?id=34105770#34109758) | [next](https://news.ycombinator.com/item?id=34105770#34109271) [–]

  

Here's one example of when to commit auto-generated code: if you use a tool like Cython and want to distribute sources to an end user, it is frequently recommended to distribute the C sources generated by Cython instead of the Cython code itself. The given reason is that the end user may not have access to Cython, but will likely have access to a C compiler. (This is frequently the case on scientific clusters where users don't have superuser privileges.)

[reply](https://news.ycombinator.com/reply?id=34109715&goto=item%3Fid%3D34105770%2334109715)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109870&how=up&auth=3c962931e9bc0532b83311cb38764f644061b6bc&goto=item%3Fid%3D34105770#34109870)

[daniele_dll](https://news.ycombinator.com/user?id=daniele_dll) [1 day ago](https://news.ycombinator.com/item?id=34109870) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109715) | [next](https://news.ycombinator.com/item?id=34105770#34109271) [–]

  

Importing a C file into a Python session is a practice that can ---easily--- cause ---a lot--- of headaches, the packages required to build the code as python library and make it importable will vary between systems and make it hard to debug in case of crashes. I personally would distribuite the pre-compiled binaries, CI + docker makes that much easier.

The fact that "can be done" doesn't necessarely mean that "it's the right way" to do it.

Going back to embedding auto-generated code, focusing on this specific case, are you telling me that a line in the requirements.txt, on a machine where the user is expected to have access to a compiler, is the reason to keep auto generated code in the repository? Sorry, personally it seems a flaky reason.

Also, the very fact that you mention that users don't have superuser privileges should be even more of a reason to use virtual environments, where you don't need super privileges, or if you can't or want to use virtual environments you can still have multiple paths in the PYTHONPATH env variable and you can install packages in a specific paths leveraging the --target of pip or the PYTHONUSERBASE env variable, solving the issue as well.

I would prefer to suggest my users to improve their runtime environment than enable them to improve the chaos.

[reply](https://news.ycombinator.com/reply?id=34109870&goto=item%3Fid%3D34105770%2334109870)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109937&how=up&auth=2dff8959eab8ba98cf4e1182cf7dab88187a5655&goto=item%3Fid%3D34105770#34109937)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [1 day ago](https://news.ycombinator.com/item?id=34109937) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109870) | [next](https://news.ycombinator.com/item?id=34105770#34109271) [–]

  

I don't agree that it's a good practice, I'm simply quoting one example to you since you mentioned you were not aware of any. IIRC, you can find many GitHub issue threads on the Cython repo where they go into more detail. It has been a while since I spent any time thinking about this so the details are no longer fresh.

Unfortunately, suggesting that users improve their runtime environment is frequently a nonstarter for different reasons. (One obvious reason: users either don't have time or technical experience to do so.)

[reply](https://news.ycombinator.com/reply?id=34109937&goto=item%3Fid%3D34105770%2334109937)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109271&how=up&auth=0558b24c6efc8018d8c9626106a0186525987978&goto=item%3Fid%3D34105770#34109271)

[baby](https://news.ycombinator.com/user?id=baby) [1 day ago](https://news.ycombinator.com/item?id=34109271) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34109537) | [next](https://news.ycombinator.com/item?id=34105770#34109552) [–]

  

It's funny because I would never think that this is a good way to learn a codebase. Code changes completely through time, and early phases where nothing is set can be extremely messy and filled with constant refactorings. Commit messages are rarely good. I would rather try to learn a codebase looking at the latest state and the docs.

[reply](https://news.ycombinator.com/reply?id=34109271&goto=item%3Fid%3D34105770%2334109271)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109788&how=up&auth=da2cd189ea7ccc062068304cbeda9f41123a9493&goto=item%3Fid%3D34105770#34109788)

[trenchgun](https://news.ycombinator.com/user?id=trenchgun) [1 day ago](https://news.ycombinator.com/item?id=34109788) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109271) | [next](https://news.ycombinator.com/item?id=34105770#34109552) [–]

  

> It's funny because I would never think that this is a good way to learn a codebase.

It is NOT for learning the codebase.

It is for learning how to program in C.

[reply](https://news.ycombinator.com/reply?id=34109788&goto=item%3Fid%3D34105770%2334109788)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111570&how=up&auth=e36a7d3c33127a79b7a7a6c2a3b5710be325dac7&goto=item%3Fid%3D34105770#34111570)

[baby](https://news.ycombinator.com/user?id=baby) [1 day ago](https://news.ycombinator.com/item?id=34111570) | [root](https://news.ycombinator.com/item?id=34105770#34107405) | [parent](https://news.ycombinator.com/item?id=34105770#34109788) | [next](https://news.ycombinator.com/item?id=34105770#34109552) [–]

  

Ah I see. I might try that with another language then! That makes more sense.

[reply](https://news.ycombinator.com/reply?id=34111570&goto=item%3Fid%3D34105770%2334111570)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109552&how=up&auth=f09e0f9e90762139f96e80a6d6a210b4a019dea7&goto=item%3Fid%3D34105770#34109552)

[pg314](https://news.ycombinator.com/user?id=pg314) [1 day ago](https://news.ycombinator.com/item?id=34109552) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34109271) | [next](https://news.ycombinator.com/item?id=34105770#34109146) [–]

  

I did that with Git itself. As a bonus you’ll learn more about a tool you’re probably using everyday. If I remember correctly, even the early commits were pretty high quality.

[reply](https://news.ycombinator.com/reply?id=34109552&goto=item%3Fid%3D34105770%2334109552)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109146&how=up&auth=cf5c03ed3be8005de7626da8f121486658c62cc4&goto=item%3Fid%3D34105770#34109146)

[shortlived](https://news.ycombinator.com/user?id=shortlived) [1 day ago](https://news.ycombinator.com/item?id=34109146) | [parent](https://news.ycombinator.com/item?id=34105770#34107405) | [prev](https://news.ycombinator.com/item?id=34105770#34109552) | [next](https://news.ycombinator.com/item?id=34105770#34107476) [–]

  

FYI, missing a second dash

```
  git log -–reverse
```

[reply](https://news.ycombinator.com/reply?id=34109146&goto=item%3Fid%3D34105770%2334109146)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107476&how=up&auth=c4f892ca358e89d1638ab56cdd5c97c07a1701b7&goto=item%3Fid%3D34105770#34107476)

[jerf](https://news.ycombinator.com/user?id=jerf) [2 days ago](https://news.ycombinator.com/item?id=34107476) | [prev](https://news.ycombinator.com/item?id=34105770#34107405) | [next](https://news.ycombinator.com/item?id=34105770#34106967) [–]

  

If you're fluent in Rust, you will have no problem learning C considered as a language. If C has any faults it is that it is too simple. Rust even sensitizes you nicely to the worst pokey bits of C, so you won't be wandering blind through the landscape of memory safety like most people learning C. I don't think you'd need books or anything, just grab some tutorials online and start banging away. It's not a hard language.

I say "C considered as a language", because in my opinion and experience the _real_ challenge in learning C is C considered as a _platform_, that is, learning how to build C projects, how to actually pull in dependencies, how to resolve them when they blow up, how to build cross-platform, understanding the several toolsets used by real projects to accomplish these tasks (raw make, autotools, cmake, etc.), dealing with library conflicts, etc. C is a nice little language, too simple for 2022 and something I wish people would really stop using, but it is at least simple. The C _platform_ is a clusterfuck nightmare disaster zone, if not multiple independent clusterfuck nightmare disaster zones.

I can say with a straight face I "know C", inasmuch as I know the language and have even done some non-trivial work in it. (Not much, but some.) I have never in my career learned C as a _platform_. This is a choice, and I still think it was even a good choice, but what I'm really highlighting in this post is that this is certainly something that is noticeable, that the fact that I "know C" has actually been of almost no utility whatsoever when I encounter problems and have to fix them because the problems are almost always in the _platform_. I have made do with a very cobbled-together understanding of the situation and a high degree of "google-fu" on this topic (beyond just "pop the error message in verbatim and hope for the best") and I've chosen not to pursue this any further.

So I would recommend to you that you consider your goals here; do you want to "know C" as a language or a platform? Both are fine options but they lead to different solutions.

[reply](https://news.ycombinator.com/reply?id=34107476&goto=item%3Fid%3D34105770%2334107476)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107806&how=up&auth=c24bc38a68bd2e501b031053b99541765d8c0218&goto=item%3Fid%3D34105770#34107806)

[aunyks](https://news.ycombinator.com/user?id=aunyks) [2 days ago](https://news.ycombinator.com/item?id=34107806) | [parent](https://news.ycombinator.com/item?id=34105770#34107476) | [next](https://news.ycombinator.com/item?id=34105770#34113327) [–]

  

> the real challenge in learning C is C considered as a platform, that is, learning how to build C projects, how to actually pull in dependencies, how to resolve them when they blow up, how to build cross-platform, understanding the several toolsets used by real projects to accomplish these tasks (raw make, autotools, cmake, etc.), dealing with library conflicts, etc.

Agreed. This is what's prevented me from learning C by contributing to large projects. I've always had trouble understanding the dependency management and build system of every project, as it seems everyone does it differently.

I kinda wish someone made a "C/C++ for Rust Developers", "CMake for the Rest of Us", or similar text that focuses enough on dep management and build systems.

[reply](https://news.ycombinator.com/reply?id=34107806&goto=item%3Fid%3D34105770%2334107806)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34113327&how=up&auth=23ff2894e42c5119ec03e98045728b14ecd63d3c&goto=item%3Fid%3D34105770#34113327)

[stcroixx](https://news.ycombinator.com/user?id=stcroixx) [1 day ago](https://news.ycombinator.com/item?id=34113327) | [parent](https://news.ycombinator.com/item?id=34105770#34107476) | [prev](https://news.ycombinator.com/item?id=34105770#34107806) | [next](https://news.ycombinator.com/item?id=34105770#34108377) [–]

  

Agreed, though often paying (not open source) C projects have much simpler build requirements because the company controls the hardware. Last time I got paid to write C on a non-trivial financial app, all we did was call make with exactly 1 external dependency.

[reply](https://news.ycombinator.com/reply?id=34113327&goto=item%3Fid%3D34105770%2334113327)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108377&how=up&auth=6a395f340beb5df8f62c987e0320e60f1942ecd8&goto=item%3Fid%3D34105770#34108377)

[KerrAvon](https://news.ycombinator.com/user?id=KerrAvon) [2 days ago](https://news.ycombinator.com/item?id=34108377) | [parent](https://news.ycombinator.com/item?id=34105770#34107476) | [prev](https://news.ycombinator.com/item?id=34105770#34113327) | [next](https://news.ycombinator.com/item?id=34105770#34106967) [–]

  

There wasn’t historically a C “platform” in the sense you describe it here due to C’s age. Rust built itself around a single “platform“ because Rust was built in the 21st century. (Python and Ruby didn’t have “platforms” either, but they had uncommonly useful standard libraries.)

[reply](https://news.ycombinator.com/reply?id=34108377&goto=item%3Fid%3D34105770%2334108377)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106967&how=up&auth=39f837b85d6c5155916a57501c40cd66bef161c8&goto=item%3Fid%3D34105770#34106967)

[liendolucas](https://news.ycombinator.com/user?id=liendolucas) [2 days ago](https://news.ycombinator.com/item?id=34106967) | [prev](https://news.ycombinator.com/item?id=34105770#34107476) | [next](https://news.ycombinator.com/item?id=34105770#34107198) [–]

  

Programming in C can turn you into a very disciplined programmer for the simple fact that you will need to track when you request memory and when you need to free it. Essentially you will become the garbage collector. I would recommend you to the read classic "The C Programming Language" (Kernighan & Ritchie) and "Expert C Programming Deep C Secrets" (Van Der Linden) which will get you far enough to write useful programs. I read half of Gustedt book and it has good advice, probably need to pick it up again. A good starting point is to implement your own data structures, like lists, double linked lists, dynamic arrays (had a lot of fun implementing this one). Also (this is personal preference) I would learn a bit of the autotools, they are painful to learn and feel kind of ugly, but once they are in place you basically will forget about a ton of things and your build will become a joy (there's a nice book by No Starch Press about it). I've managed to integrate Check ([https://libcheck.github.io/check/](https://libcheck.github.io/check/)) so now I have unit tests in C which is something nice.

[reply](https://news.ycombinator.com/reply?id=34106967&goto=item%3Fid%3D34105770%2334106967)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107232&how=up&auth=02f9d56c718a2a8eb090c8204284bb800c624ca7&goto=item%3Fid%3D34105770#34107232)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [2 days ago](https://news.ycombinator.com/item?id=34107232) | [parent](https://news.ycombinator.com/item?id=34105770#34106967) | [next](https://news.ycombinator.com/item?id=34105770#34107182) [–]

  

I don't recommend learning how to use Autotools unless you need to work on a project that uses it (and even then, try to avoid it).

Highly recommend learning CMake or Meson instead.

If you're just getting started, there's no need to tie yourself into knots learning "how to set up a project the right way". It's a waste of time. Your time would initially be better spent just calling the C compiler directly. This will help you learn a bit about the compilation and linking phase and will expose you to the C compiler's flags---useful skills for debugging automated builds later. Once you get to the point where you are working on a larger project, then you will be in a better position to decide what build tools are most useful for you.

[reply](https://news.ycombinator.com/reply?id=34107232&goto=item%3Fid%3D34105770%2334107232)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107391&how=up&auth=6f1ccbe444791773f1a433616f673ef51ecca5aa&goto=item%3Fid%3D34105770#34107391)

[liendolucas](https://news.ycombinator.com/user?id=liendolucas) [2 days ago](https://news.ycombinator.com/item?id=34107391) | [root](https://news.ycombinator.com/item?id=34105770#34106967) | [parent](https://news.ycombinator.com/item?id=34105770#34107232) | [next](https://news.ycombinator.com/item?id=34105770#34111124) [–]

  

> If you're just getting started, there's no need to tie ...

Yes, this is sane advice! Note that I'm not promoting autotools, I personally suffered it a lot before I got it for my builds, but is a tool that can be reliable across platforms. Yes I would leave the "how to set up a project" for the time when you really need it, no need to spend time on this if you're learning a language.

[reply](https://news.ycombinator.com/reply?id=34107391&goto=item%3Fid%3D34105770%2334107391)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111124&how=up&auth=ddfff465853661871b5a31008c6be575949939bc&goto=item%3Fid%3D34105770#34111124)

[mikewarot](https://news.ycombinator.com/user?id=mikewarot) [1 day ago](https://news.ycombinator.com/item?id=34111124) | [root](https://news.ycombinator.com/item?id=34105770#34106967) | [parent](https://news.ycombinator.com/item?id=34105770#34107232) | [prev](https://news.ycombinator.com/item?id=34105770#34107391) | [next](https://news.ycombinator.com/item?id=34105770#34108968) [–]

  

I have a C project I picked up (and got working thanks to help from HN users) that was heavily dependent on Autotools. I wanted to be able to compile it under Windows and couldn't make heads nor tails of it. After guessing and tweaking a config.h file to replace the one it didn't create, I was finally able to get gcc to compile everything using a batch file.

Eventually I'll be purging Autotools, make and all build systems, replacing it with a single batch file that wipes the slate then compiles. It only takes a few seconds to run, and removes a huge set of dependencies.

I hope to port the code to run under the Cosmopolitan library so I can be cross-platform and more.

[reply](https://news.ycombinator.com/reply?id=34111124&goto=item%3Fid%3D34105770%2334111124)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108968&how=up&auth=ef0326bc4f7c71b44ad285bee156d17f4a87cd16&goto=item%3Fid%3D34105770#34108968)

[tmtvl](https://news.ycombinator.com/user?id=tmtvl) [1 day ago](https://news.ycombinator.com/item?id=34108968) | [root](https://news.ycombinator.com/item?id=34105770#34106967) | [parent](https://news.ycombinator.com/item?id=34105770#34107232) | [prev](https://news.ycombinator.com/item?id=34105770#34111124) | [next](https://news.ycombinator.com/item?id=34105770#34108797) [–]

  

I disagree, I find autoconf and automake very pleasant in operation. Mainly due to the OOTB support for things like bindir, libdir, and so on, all properly documented via _/path/to/configure --help_.

[reply](https://news.ycombinator.com/reply?id=34108968&goto=item%3Fid%3D34105770%2334108968)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108797&how=up&auth=8e845dc7102e7a8f2a19807afdf5de1d10ae967d&goto=item%3Fid%3D34105770#34108797)

[varjag](https://news.ycombinator.com/user?id=varjag) [1 day ago](https://news.ycombinator.com/item?id=34108797) | [root](https://news.ycombinator.com/item?id=34105770#34106967) | [parent](https://news.ycombinator.com/item?id=34105770#34107232) | [prev](https://news.ycombinator.com/item?id=34105770#34108968) | [next](https://news.ycombinator.com/item?id=34105770#34107182) [–]

  

You probably don't need autotools as a beginner but if you intend to study those it makes no point starting with cmake. This is more of a C++ crowd thing with tiny mindshare among C programmers.

[reply](https://news.ycombinator.com/reply?id=34108797&goto=item%3Fid%3D34105770%2334108797)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107182&how=up&auth=de6c9f5d296d218e7d70bbf40161ed735fe0bb5f&goto=item%3Fid%3D34105770#34107182)

[gatane](https://news.ycombinator.com/user?id=gatane) [2 days ago](https://news.ycombinator.com/item?id=34107182) | [parent](https://news.ycombinator.com/item?id=34105770#34106967) | [prev](https://news.ycombinator.com/item?id=34105770#34107232) | [next](https://news.ycombinator.com/item?id=34105770#34107057) [–]

  

This. We called K&R as "The C bible" for a reason at Uni.

[reply](https://news.ycombinator.com/reply?id=34107182&goto=item%3Fid%3D34105770%2334107182)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107057&how=up&auth=2331e68c43d14000bafbbe702c41b9c0a74e5f2a&goto=item%3Fid%3D34105770#34107057)

[jasonjei](https://news.ycombinator.com/user?id=jasonjei) [2 days ago](https://news.ycombinator.com/item?id=34107057) | [parent](https://news.ycombinator.com/item?id=34105770#34106967) | [prev](https://news.ycombinator.com/item?id=34105770#34107182) | [next](https://news.ycombinator.com/item?id=34105770#34111532) [–]

  

I’m also a traditionalist and I think the classic K&R C Programming Language book is the best way to learn C—you also get an understanding why C and *nix-based systems seem to go so well.

[reply](https://news.ycombinator.com/reply?id=34107057&goto=item%3Fid%3D34105770%2334107057)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107155&how=up&auth=72a4858c0aeba791ba5614d9a9ba5a03cc4544af&goto=item%3Fid%3D34105770#34107155)

[liendolucas](https://news.ycombinator.com/user?id=liendolucas) [2 days ago](https://news.ycombinator.com/item?id=34107155) | [root](https://news.ycombinator.com/item?id=34105770#34106967) | [parent](https://news.ycombinator.com/item?id=34105770#34107057) | [next](https://news.ycombinator.com/item?id=34105770#34111532) [–]

  

I love books, even if I don't finish them, especially the technical ones. Today we're very used to recommend sites and youtube videos as primers for learning. In some situations that is ok but I wouldn't hesitate for even a second if I have to choose been something online or a book.

[reply](https://news.ycombinator.com/reply?id=34107155&goto=item%3Fid%3D34105770%2334107155)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111532&how=up&auth=dfefd06c0106578dd68138544d773711fbc4d889&goto=item%3Fid%3D34105770#34111532)

[wink](https://news.ycombinator.com/user?id=wink) [1 day ago](https://news.ycombinator.com/item?id=34111532) | [parent](https://news.ycombinator.com/item?id=34105770#34106967) | [prev](https://news.ycombinator.com/item?id=34105770#34107057) | [next](https://news.ycombinator.com/item?id=34105770#34108402) [–]

  

+1 for "Deep C Secrets" - excellent book. I don't use C enough to judge if it has stood up to the test of time, but I loved reading it.

[reply](https://news.ycombinator.com/reply?id=34111532&goto=item%3Fid%3D34105770%2334111532)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108402&how=up&auth=c267cb4e116b07515d26afbf957caeb43e0d4314&goto=item%3Fid%3D34105770#34108402)

[sicp-enjoyer](https://news.ycombinator.com/user?id=sicp-enjoyer) [2 days ago](https://news.ycombinator.com/item?id=34108402) | [parent](https://news.ycombinator.com/item?id=34105770#34106967) | [prev](https://news.ycombinator.com/item?id=34105770#34111532) | [next](https://news.ycombinator.com/item?id=34105770#34107198) [–]

  

Deep C Secrets is one of the best programming books ever written.

[reply](https://news.ycombinator.com/reply?id=34108402&goto=item%3Fid%3D34105770%2334108402)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107198&how=up&auth=531cd066f253aa78b322ddaa32cefe8a3efc0121&goto=item%3Fid%3D34105770#34107198)

[gwbas1c](https://news.ycombinator.com/user?id=gwbas1c) [2 days ago](https://news.ycombinator.com/item?id=34107198) | [prev](https://news.ycombinator.com/item?id=34105770#34106967) | [next](https://news.ycombinator.com/item?id=34105770#34106582) [–]

  

The thing about C is that its library call API is generally how different languages interoperate: If you write a library in [insert language here] and you want it to be callable in any other language, you need to make bindings for C.

Likewise, if you're using any other language and you want to use a general library, you'll need to know enough C so that you can successfully encapsulate concepts like pointers; and copy your program's data into data structures that C can consume.

It's all much, much easier if you have some C experience.

In my case, I haven't done much C in 20 years. What really helped me "keep up" was occasionally doing Mac programming in Objective C. Because it's really easy to drop into C, it helped refresh me on a lot of the concepts. Even if you're doing Swift, you can mix in an class written in Objective C that's mostly straight C in the methods.

[reply](https://news.ycombinator.com/reply?id=34107198&goto=item%3Fid%3D34105770%2334107198)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106582&how=up&auth=abfc4eebde3f13364f6b86241fbb783f80954044&goto=item%3Fid%3D34105770#34106582)

[thom](https://news.ycombinator.com/user?id=thom) [2 days ago](https://news.ycombinator.com/item?id=34106582) | [prev](https://news.ycombinator.com/item?id=34105770#34107198) | [next](https://news.ycombinator.com/item?id=34105770#34106463) [–]

  

Modern C by Jens Gustedt is quite good. This video gives a good overview of some modern patterns in C:

[https://www.youtube.com/watch?v=QpAhX-gsHMs](https://www.youtube.com/watch?v=QpAhX-gsHMs)

I think the best thing you can do is go find a unit testing library you like the look of, maybe some tools for address/leak sanitization, and spend a good amount of time learning CMake and creating a build/project workflow you're happy with. Coming from Rust you've got great tooling and good guardrails in place to keep you on the happy path, I think getting into that groove early in C where you can run experiments and break things in a sane environment is important.

[reply](https://news.ycombinator.com/reply?id=34106582&goto=item%3Fid%3D34105770%2334106582)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108306&how=up&auth=c4981debcca90430ded9281c13187038fe5200d9&goto=item%3Fid%3D34105770#34108306)

[fileeditview](https://news.ycombinator.com/user?id=fileeditview) [2 days ago](https://news.ycombinator.com/item?id=34108306) | [parent](https://news.ycombinator.com/item?id=34105770#34106582) | [next](https://news.ycombinator.com/item?id=34105770#34106463) [–]

  

Came also here to recommend this book. You can obtain it for free at the authors site: [https://gustedt.gitlabpages.inria.fr/modern-c/](https://gustedt.gitlabpages.inria.fr/modern-c/)

If you like it you could buy the print version.

There are some things he writes that are a bit controversial but generally I think it is very thorough. Sometimes it felt a bit dry to read to me.

[reply](https://news.ycombinator.com/reply?id=34108306&goto=item%3Fid%3D34105770%2334108306)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106463&how=up&auth=f05e989744a662a4d0b3531ee1ca24533d3c76d3&goto=item%3Fid%3D34105770#34106463)

[Waterluvian](https://news.ycombinator.com/user?id=Waterluvian) [2 days ago](https://news.ycombinator.com/item?id=34106463) | [prev](https://news.ycombinator.com/item?id=34105770#34106582) | [next](https://news.ycombinator.com/item?id=34105770#34106261) [–]

  

Write a CHIP-8 emulator in C!

In my opinion it’s a perfect project because:

- it is a good size. You can be done in 10-20 hours, or longer if that’s your pace (have fun!)

- a lot of examples and community support.

- you have a whole project without any overwhelming complexities: you’re just taking inputs, reading some data, running it through a switch statement, writing data, writing to a screen, then repeating.

- you have lots of “test code” aka games.

- it mostly lacks tedium: you’re not repeating ideas too often

- it’s fun!

[reply](https://news.ycombinator.com/reply?id=34106463&goto=item%3Fid%3D34105770%2334106463)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106809&how=up&auth=54bffa86a931ab7c67ef3ff709d4baf3aaa80ee5&goto=item%3Fid%3D34105770#34106809)

[mr_00ff00](https://news.ycombinator.com/user?id=mr_00ff00) [2 days ago](https://news.ycombinator.com/item?id=34106809) | [parent](https://news.ycombinator.com/item?id=34105770#34106463) | [next](https://news.ycombinator.com/item?id=34105770#34106261) [–]

  

“can be done in a weekend”

RIP me who has spent 3 weeks on this

Great project for sure though!

[reply](https://news.ycombinator.com/reply?id=34106809&goto=item%3Fid%3D34105770%2334106809)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106911&how=up&auth=00f40b137cd37cd8ca2abc223a1372366ef91ba7&goto=item%3Fid%3D34105770#34106911)

[Waterluvian](https://news.ycombinator.com/user?id=Waterluvian) [2 days ago](https://news.ycombinator.com/item?id=34106911) | [root](https://news.ycombinator.com/item?id=34105770#34106463) | [parent](https://news.ycombinator.com/item?id=34105770#34106809) | [next](https://news.ycombinator.com/item?id=34105770#34106261) [–]

  

haha I guess a “weekend” is a fuzzy unit of measure. I’ll amend.

[reply](https://news.ycombinator.com/reply?id=34106911&goto=item%3Fid%3D34105770%2334106911)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106261&how=up&auth=015343a273b940303fb5b9500ec5a6d658eb6532&goto=item%3Fid%3D34105770#34106261)

[jshrake](https://news.ycombinator.com/user?id=jshrake) [2 days ago](https://news.ycombinator.com/item?id=34106261) | [prev](https://news.ycombinator.com/item?id=34105770#34106463) | [next](https://news.ycombinator.com/item?id=34105770#34107727) [–]

  

I recommend _C Programming: A Modern Approach_ by K.N. King [http://knking.com/books/c2/index.html](http://knking.com/books/c2/index.html). My favorite part of this book is the exercises at the end of each chapter.

[reply](https://news.ycombinator.com/reply?id=34106261&goto=item%3Fid%3D34105770%2334106261)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106980&how=up&auth=9e58bdc6bf836022138c992439013c6fea33d51d&goto=item%3Fid%3D34105770#34106980)

[sarchertech](https://news.ycombinator.com/user?id=sarchertech) [2 days ago](https://news.ycombinator.com/item?id=34106980) | [parent](https://news.ycombinator.com/item?id=34105770#34106261) | [next](https://news.ycombinator.com/item?id=34105770#34107727) [–]

  

One of my favorite professors. His programming language concepts course was the most useful useful class I took in university.

I also recommend his book.

[reply](https://news.ycombinator.com/reply?id=34106980&goto=item%3Fid%3D34105770%2334106980)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107727&how=up&auth=62ef196872b17954617c7e46c663568deef64a5a&goto=item%3Fid%3D34105770#34107727)

[photochemsyn](https://news.ycombinator.com/user?id=photochemsyn) [2 days ago](https://news.ycombinator.com/item?id=34107727) | [prev](https://news.ycombinator.com/item?id=34105770#34106261) | [next](https://news.ycombinator.com/item?id=34105770#34107471) [–]

  

A main issue with learning C is getting the architecture right (header and source files), and figuring out the tooling for compiling and debugging and building (gcc and gdb and make/cmake for example). I like C in a Nutshell 2nd ed (Prinz/Crawford) 2016 as a reference for any questions about that:

[https://www.oreilly.com/library/view/c-in-a/9781491924174/](https://www.oreilly.com/library/view/c-in-a/9781491924174/)

As far as code examples, I find there's a lot of simple C graphics programs using the SDL2 library discussed on Youtube. Also simple ncurses programs are worth looking at. For example here's a video on making a tic-tac-toe game with pure C linked to github source code, just a few files:

[https://www.youtube.com/watch?v=gCVMkKgs3uQ](https://www.youtube.com/watch?v=gCVMkKgs3uQ)

[https://github.com/tsoding/profun/tree/master/procedural](https://github.com/tsoding/profun/tree/master/procedural)

For practice, CodeWars has a lot of nice simple C challenges to work on.

[reply](https://news.ycombinator.com/reply?id=34107727&goto=item%3Fid%3D34105770%2334107727)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109436&how=up&auth=54c5dd8199b4a73f6f9be63f78eb92eab2d07159&goto=item%3Fid%3D34105770#34109436)

[alar44](https://news.ycombinator.com/user?id=alar44) [1 day ago](https://news.ycombinator.com/item?id=34109436) | [parent](https://news.ycombinator.com/item?id=34105770#34107727) | [next](https://news.ycombinator.com/item?id=34105770#34107471) [–]

  

I've been using GPT this way. "Write the header file for a function that does X." Now write the implementation. Does a great job when prompted with a header file.

[reply](https://news.ycombinator.com/reply?id=34109436&goto=item%3Fid%3D34105770%2334109436)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107471&how=up&auth=7061664beee0b5095be2f7a3f23966e46ed8ae8c&goto=item%3Fid%3D34105770#34107471)

[dxbydt](https://news.ycombinator.com/user?id=dxbydt) [2 days ago](https://news.ycombinator.com/item?id=34107471) | [prev](https://news.ycombinator.com/item?id=34105770#34107727) | [next](https://news.ycombinator.com/item?id=34105770#34105819) [–]

  

One thing I always loved about C was how easily you could slip in and out of x86 assembler with the __asm__ tag. There is no ceremony & it seems like magic. Like if you have 3 C int's a,b & c, and you want to add a & b and store the result in c, you would do:

__asm__ ( "addl %%ebx, %%eax;" : "=a" (c) : "a" (a) , "b" (b) );

Then you could just continue in C as if nothing had happened, and a printf("%d",c) would give you the right result. When I worked as a quant, I saw such code snippets quite frequently. Those days (late 90s early 2k), individual programmers had a certain amount of freedom that is completely missing today. You would have a cohesive team with a C programmer, java programmer, a sysadmin, an R guy, a kdb/q quant - no issues. The front end was all java awt widgets, and you would jni to do the hard bits. jni would pull in the C code, which was peppered with little bits of __asm__ to speed things up. Everytime you needed to interact with the outer ecosystem, C would do an fwrite to the file system, then R would pick up the file, do its montecarlo simulations & write back to the filesystem, the algotrading bits would be in q, and even the sysadmin had his share of awk/sed to massage all the back & forth data. It was smooth & nice & pulled 100x its weight. Best of all, the individual programmers had respect for each other's skills & would learn from each other. These days, some asshole CIO would fire the whole bunch & replace them with 10 pythonistas & insist everything should be coded in one true language - and as one of those pythonistas, while I am happy I still have a job in this recession, it is kinda sad, those languages were so much more fun & it was exciting to watch the interplay of different languages bearing their own strengths, whereas python otoh is so terribly boring & everyday is same verbose torch.autograd.Function crap oh well back to the coal mines.

[reply](https://news.ycombinator.com/reply?id=34107471&goto=item%3Fid%3D34105770%2334107471)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108519&how=up&auth=5f52d5ab27f47179a924e0bb77d936fcea51828c&goto=item%3Fid%3D34105770#34108519)

[rramadass](https://news.ycombinator.com/user?id=rramadass) [2 days ago](https://news.ycombinator.com/item?id=34108519) | [parent](https://news.ycombinator.com/item?id=34105770#34107471) | [next](https://news.ycombinator.com/item?id=34105770#34105819) [–]

  

Man, that is one complex, interconnected and interesting System!

I would love to work on something like that even if it is just maintaining a legacy product. Seriously, let me know (contact in my profile).

On similar lines, I have been thinking of using Erlang/Elixir for distributed programming over the Web (even in the browser if possible) and dropping down to C/C++ and J programming languages as required on individual nodes.

[reply](https://news.ycombinator.com/reply?id=34108519&goto=item%3Fid%3D34105770%2334108519)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109747&how=up&auth=10d6e80a7089df409f796dba3f0a90ceb2daa100&goto=item%3Fid%3D34105770#34109747)

[dxbydt](https://news.ycombinator.com/user?id=dxbydt) [1 day ago](https://news.ycombinator.com/item?id=34109747) | [root](https://news.ycombinator.com/item?id=34105770#34107471) | [parent](https://news.ycombinator.com/item?id=34105770#34108519) | [next](https://news.ycombinator.com/item?id=34105770#34105819) [–]

  

If you like such systems, backend systems in investment banks is where its at. Most of my experience was at Goldman, but similar systems exist in Morgan, Deutsche, Citi, UBS, Bloomberg or old PE firms. Look for C programmer careers in trade publications like efinancialcareers.com, or work with Huxley, Selby Jennings etc. Combination of C/C++ and certification in Q/kdb - will keep you employed pretty much into your 60s. Those languages are intrinsically hard & imo a real pleasure to code in. Well, I liked it.

[reply](https://news.ycombinator.com/reply?id=34109747&goto=item%3Fid%3D34105770%2334109747)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34113534&how=up&auth=cc866cc32721dbb75f0c3175432115ff20341c2e&goto=item%3Fid%3D34105770#34113534)

[rramadass](https://news.ycombinator.com/user?id=rramadass) [1 day ago](https://news.ycombinator.com/item?id=34113534) | [root](https://news.ycombinator.com/item?id=34105770#34107471) | [parent](https://news.ycombinator.com/item?id=34105770#34109747) | [next](https://news.ycombinator.com/item?id=34105770#34105819) [–]

  

Thanks for the pointers, appreciate it very much.

[reply](https://news.ycombinator.com/reply?id=34113534&goto=item%3Fid%3D34105770%2334113534)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34105819&how=up&auth=e14ad5a7baf0cccc2e04db2b14d45a7542f64b61&goto=item%3Fid%3D34105770#34105819)

[AlexeyBrin](https://news.ycombinator.com/user?id=AlexeyBrin) [2 days ago](https://news.ycombinator.com/item?id=34105819) | [prev](https://news.ycombinator.com/item?id=34105770#34107471) | [next](https://news.ycombinator.com/item?id=34105770#34106610) [–]

  

Depends on your style of learning, but most people learn best by implementing something in a new language. Here are a few examples: a Brainfuck interpreter, a raytracer, even command line tools like ls, cat, grep are good exercises. If you are into electronics C is great for embedded work - I would buy a Raspberry Pi Pico and write C programs for it.

I would also suggest to enable all warnings your compiler allow for and sanitizers if possible. You will also want to learn to use a debugger that works on your operating system.

[reply](https://news.ycombinator.com/reply?id=34105819&goto=item%3Fid%3D34105770%2334105819)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106610&how=up&auth=69222867a5a33f364cd591303ceb3b357e96c26a&goto=item%3Fid%3D34105770#34106610)

[jviotti](https://news.ycombinator.com/user?id=jviotti) [2 days ago](https://news.ycombinator.com/item?id=34106610) | [prev](https://news.ycombinator.com/item?id=34105770#34105819) | [next](https://news.ycombinator.com/item?id=34105770#34106569) [–]

  

The language is only part of it. In my opinion, you would learn even more about C by consuming well-known libraries written in C and trying to do something useful with it and complementing that with some books, etc rather than just learning the syntax and semantics of C.

There are tons of useful/fun things you could do with C. Write some desktop apps with GTK, a web server with Mongoose ([https://mongoose.ws](https://mongoose.ws/)), etc.

[reply](https://news.ycombinator.com/reply?id=34106610&goto=item%3Fid%3D34105770%2334106610)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106569&how=up&auth=a3dc0979fa5b59a368ec9b4b2f6e752ef73d31c2&goto=item%3Fid%3D34105770#34106569)

[kovac](https://news.ycombinator.com/user?id=kovac) [2 days ago](https://news.ycombinator.com/item?id=34106569) | [prev](https://news.ycombinator.com/item?id=34105770#34106610) | [next](https://news.ycombinator.com/item?id=34105770#34107298) [–]

  

I'm a C# guy who learnt C over the past 2 years. It absolutely made me a better programmer. I learnt by going through K&R first and then Expert C programming by Peter van der Linden. I also went through OS books, books that talk about the machines like Understanding the machine. I program in many high level languages. The main difference I felt when learning C is, I felt like I needed to learn how the machine actually works (what happens when we run a program) to really appreciate C. I continue to pursue this learning, these days I'm looking into embedded stuff.

Also look at C programs. I study old, small unix programs like "more".

I must also say that I see people asking about learning C on sites like Stack Overflow and getting discouraged saying that it's outdated and all. These people are doing a great disservice.

[reply](https://news.ycombinator.com/reply?id=34106569&goto=item%3Fid%3D34105770%2334106569)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107073&how=up&auth=32c2555480595a18a63f2fc14da7986b50ef686e&goto=item%3Fid%3D34105770#34107073)

[tester756](https://news.ycombinator.com/user?id=tester756) [2 days ago](https://news.ycombinator.com/item?id=34107073) | [parent](https://news.ycombinator.com/item?id=34105770#34106569) | [next](https://news.ycombinator.com/item?id=34105770#34107298) [–]

  

As C# guy who had to work with C for a while - this language is minefield.

If you want to create your own language and want to experience what choices are subpar by modern standards - C is way to go.

If I had choice then I'd pick Rust anytime when messing with low level programming.

I don't understand fascination with C just because it's "tradition" "everywhere" or "gives you full control" - especially that, since it brought us countless security issues

[reply](https://news.ycombinator.com/reply?id=34107073&goto=item%3Fid%3D34105770%2334107073)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107282&how=up&auth=72b6cbea7f5bd0f7f80f8bf0bd09090f19fd281d&goto=item%3Fid%3D34105770#34107282)

[gwbas1c](https://news.ycombinator.com/user?id=gwbas1c) [2 days ago](https://news.ycombinator.com/item?id=34107282) | [root](https://news.ycombinator.com/item?id=34105770#34106569) | [parent](https://news.ycombinator.com/item?id=34105770#34107073) | [next](https://news.ycombinator.com/item?id=34105770#34107298) [–]

  

You'll wish you have decent C experience whenever you need to call an OS API directly; or whenever you need to write your own bindings to a library written in C.

PInvoke requires a working knowledge of C's concepts; you need to know the basics about pointers and how data is represented in order to do it correctly.

This statement comes from a lot of personal experience: I wrote a (currently closed source) adapter to call into SWMM from C#. ([https://www.openswmm.org/SWMM51007/swmm5-c](https://www.openswmm.org/SWMM51007/swmm5-c)). (Hopefully I can open-source it at some point.)

In [https://www.syncplicity.com/en](https://www.syncplicity.com/en), I wrote code in C# to interface with a closed-source driver to create a user-mode virtual disk drive. I also integrated with a different Virtual disk drive product that was written in C++ via PInvoke. I used PInvoke to create bindings between C# and Objective C for the Mac version as well.

[reply](https://news.ycombinator.com/reply?id=34107282&goto=item%3Fid%3D34105770%2334107282)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107340&how=up&auth=d2914e4d490ea5645843f1cfb023f7b7111ecf70&goto=item%3Fid%3D34105770#34107340)

[tester756](https://news.ycombinator.com/user?id=tester756) [2 days ago](https://news.ycombinator.com/item?id=34107340) | [root](https://news.ycombinator.com/item?id=34105770#34106569) | [parent](https://news.ycombinator.com/item?id=34105770#34107282) | [next](https://news.ycombinator.com/item?id=34105770#34107298) [–]

  

Maybe yes, but on the other hand I've never been paid for anything that used OS calls directly.

The only times I've used such APIs was creating simple keylogger for tests and very simple cheat for a video game (tested on non-official servers).

C# nowadays seems mostly like web backend development, tools, game scripting or automation, at least that's my perception.

Anyway even if WinAPI, then why not Rust?

[reply](https://news.ycombinator.com/reply?id=34107340&goto=item%3Fid%3D34105770%2334107340)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110379&how=up&auth=0b68a7bce9a1354c4bf51104a0ca9a4b29e04791&goto=item%3Fid%3D34105770#34110379)

[gwbas1c](https://news.ycombinator.com/user?id=gwbas1c) [1 day ago](https://news.ycombinator.com/item?id=34110379) | [root](https://news.ycombinator.com/item?id=34105770#34106569) | [parent](https://news.ycombinator.com/item?id=34105770#34107340) | [next](https://news.ycombinator.com/item?id=34105770#34107619) [–]

  

> Anyway even if WinAPI, then why not Rust?

Plenty of reasons:

First of all, most operating systems (and other lower-level APIs) you want to work with expose their APIs as C. These don't all have bindings in Rust. Even if you want to do them in Rust, you need to know enough C to be able to understand how to consume the API.

But then, there's plenty of legacy code in C. We aren't going to rewrite everything in Rust overnight.

More importantly, I think the "jury's still out" if Rust is better than C for embedded programming; or programming that's primarily direct memory manipulation. (And, if your embedded environment only gives you an API in C, you need to know enough C to write the API bindings for Rust before you can do embedded programming in Rust.)

Personally, I don't have any experience with unsafe Rust code. (I only do hobby projects in Rust.) In C# you can do C-style direct memory manipulation if you need to; but I'm not sure if Rust's "unsafe" is equivalent.

[reply](https://news.ycombinator.com/reply?id=34110379&goto=item%3Fid%3D34105770%2334110379)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107619&how=up&auth=8254894cc44c0c287a6a9f084926096e41715737&goto=item%3Fid%3D34105770#34107619)

[UncleEntity](https://news.ycombinator.com/user?id=UncleEntity) [2 days ago](https://news.ycombinator.com/item?id=34107619) | [root](https://news.ycombinator.com/item?id=34105770#34106569) | [parent](https://news.ycombinator.com/item?id=34105770#34107340) | [prev](https://news.ycombinator.com/item?id=34105770#34110379) | [next](https://news.ycombinator.com/item?id=34105770#34107298) [–]

  

> Anyway even if WinAPI, then why not Rust?

Because they said they already know rust and were asking about C…

[reply](https://news.ycombinator.com/reply?id=34107619&goto=item%3Fid%3D34105770%2334107619)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107824&how=up&auth=66cc2333211f405e428fdb44d736f09075b823ec&goto=item%3Fid%3D34105770#34107824)

[tester756](https://news.ycombinator.com/user?id=tester756) [2 days ago](https://news.ycombinator.com/item?id=34107824) | [root](https://news.ycombinator.com/item?id=34105770#34106569) | [parent](https://news.ycombinator.com/item?id=34105770#34107619) | [next](https://news.ycombinator.com/item?id=34105770#34107298) [–]

  

I meant in general, not just OP

[reply](https://news.ycombinator.com/reply?id=34107824&goto=item%3Fid%3D34105770%2334107824)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107298&how=up&auth=5c84a18b84f149e26cbe24ad2325b3329ad2e1bb&goto=item%3Fid%3D34105770#34107298)

[gptadmirer](https://news.ycombinator.com/user?id=gptadmirer) [2 days ago](https://news.ycombinator.com/item?id=34107298) | [prev](https://news.ycombinator.com/item?id=34105770#34106569) | [next](https://news.ycombinator.com/item?id=34105770#34108020) [–]

  

Are you me? I am learning C lately, been a few weeks. I am an SWE with 7 years of experience in JS/TS world.

Currently I am trying to learn how to test C correctly, seems there are lots of approaches.

It began when I decided to learn low level development a few months ago. I completed my Chip 8 emulator. Next I bought Arduino a month ago and I was playing around with it. Now I am reading AVR Programming book so I decided to learn C properly. On the side I already made a small keyboard with a microcontroller with QMK firmware. Soldering and stuffs.

Next what I’m trying to learn is some electronics knowledge just enough to make stuffs.

This is way more fun than my web development job!

Currently the resources I use: - Learn C The Hard Way - BeeJ’s guide to C - random video tutorials from the internet - Make: Arduino - Make: AVR programming

[reply](https://news.ycombinator.com/reply?id=34107298&goto=item%3Fid%3D34105770%2334107298)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108020&how=up&auth=3dca936fa7b53847612741a18f258f66b01e43c5&goto=item%3Fid%3D34105770#34108020)

[pkkm](https://news.ycombinator.com/user?id=pkkm) [2 days ago](https://news.ycombinator.com/item?id=34108020) | [prev](https://news.ycombinator.com/item?id=34105770#34107298) | [next](https://news.ycombinator.com/item?id=34105770#34106227) [–]

  

The K&R book is a great example of clear technical writing, but it's also quite terse. I'd recommend C Programming: A Modern Approach by K. N. King instead, as it uses a more gentle pedagogical approach.

[reply](https://news.ycombinator.com/reply?id=34108020&goto=item%3Fid%3D34105770%2334108020)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108110&how=up&auth=b8176b88958a41174e37a867ba7ef4bcfc810ff7&goto=item%3Fid%3D34105770#34108110)

[dragonwriter](https://news.ycombinator.com/user?id=dragonwriter) [2 days ago](https://news.ycombinator.com/item?id=34108110) | [parent](https://news.ycombinator.com/item?id=34105770#34108020) | [next](https://news.ycombinator.com/item?id=34105770#34106227) [–]

  

> The K&R book is a great example of clear technical writing, but it's also quite terse.

That's not really a “but”; eschewing superfluous verbosity is a hallmark of good technical writing.

IMO, a more significant strike against K&R, which is excellent for its time, is that it is hopelessly dated—unless there is a third edition I am not aware of—providing, in the second edition, a preview of the first edition of ANSI C, which was standardized just after it was released.

[reply](https://news.ycombinator.com/reply?id=34108110&goto=item%3Fid%3D34105770%2334108110)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106227&how=up&auth=b245a37695f759dfcb9f8d905f3e284464d6caa0&goto=item%3Fid%3D34105770#34106227)

[eggy](https://news.ycombinator.com/user?id=eggy) [2 days ago](https://news.ycombinator.com/item?id=34106227) | [prev](https://news.ycombinator.com/item?id=34105770#34108020) | [next](https://news.ycombinator.com/item?id=34105770#34107671) [–]

  

I suggest downloading and playing with Raylib[1]. It's written in C, not C++, and is a library for video games, but it can be used to produce apps too. It has a ton of examples for you to review and modify as you learn, and it is cross platform. There are even bindings to other languages if you decide to leave C and use your favorite language. It's a one-click download and setup. Couldn't be easier and more fun!

[1]  [https://www.raylib.com/](https://www.raylib.com/)

[reply](https://news.ycombinator.com/reply?id=34106227&goto=item%3Fid%3D34105770%2334106227)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106427&how=up&auth=acc0c0d7cf3d8044b8cde93f649f8dafb37670cb&goto=item%3Fid%3D34105770#34106427)

[stephc_int13](https://news.ycombinator.com/user?id=stephc_int13) [2 days ago](https://news.ycombinator.com/item?id=34106427) | [parent](https://news.ycombinator.com/item?id=34105770#34106227) | [next](https://news.ycombinator.com/item?id=34105770#34107671) [–]

  

SDL or Raylib are very good to quickly write interactive programs or simply play around with the samples.

[reply](https://news.ycombinator.com/reply?id=34106427&goto=item%3Fid%3D34105770%2334106427)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107671&how=up&auth=5b5b9c99d00c65464eda38639e77a8add931758a&goto=item%3Fid%3D34105770#34107671)

[jcadam](https://news.ycombinator.com/user?id=jcadam) [2 days ago](https://news.ycombinator.com/item?id=34107671) | [prev](https://news.ycombinator.com/item?id=34105770#34106227) | [next](https://news.ycombinator.com/item?id=34105770#34107590) [–]

  

C was the first language I learned after messing around with various flavors of BASIC way back in the day. I learned C using Turbo C for DOS and a big book, which is probably an 'obsolete' way of learning in 2023.

Haven't used it professionally in over a decade now, I'd be interested in learning Rust, but don't really have a need to do anything low(ish) level these days.

[reply](https://news.ycombinator.com/reply?id=34107671&goto=item%3Fid%3D34105770%2334107671)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108627&how=up&auth=855079f561c76a27b60c558707dd0198e9b7487f&goto=item%3Fid%3D34105770#34108627)

[agentultra](https://news.ycombinator.com/user?id=agentultra) [2 days ago](https://news.ycombinator.com/item?id=34108627) | [parent](https://news.ycombinator.com/item?id=34105770#34107671) | [next](https://news.ycombinator.com/item?id=34105770#34107590) [–]

  

Same here! I started with AmigaBASIC and 68000 assembly and learned C when my family switched over to Intel/DOS machines.

Addressing modes back then!

[reply](https://news.ycombinator.com/reply?id=34108627&goto=item%3Fid%3D34105770%2334108627)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108690&how=up&auth=86bea24caaad80c19fca8b59e9b01ef4697fb49e&goto=item%3Fid%3D34105770#34108690)

[jcadam](https://news.ycombinator.com/user?id=jcadam) [2 days ago](https://news.ycombinator.com/item?id=34108690) | [root](https://news.ycombinator.com/item?id=34105770#34107671) | [parent](https://news.ycombinator.com/item?id=34105770#34108627) | [next](https://news.ycombinator.com/item?id=34105770#34107590) [–]

  

Yep. Started with Apple BASIC on an Apple ][e, then Amos on the Amiga, then C on MSDOS :).

[reply](https://news.ycombinator.com/reply?id=34108690&goto=item%3Fid%3D34105770%2334108690)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107590&how=up&auth=5a187ef2e4d3726b4d45564ba185efe3c7d91b47&goto=item%3Fid%3D34105770#34107590)

[pjmlp](https://news.ycombinator.com/user?id=pjmlp) [2 days ago](https://news.ycombinator.com/item?id=34107590) | [prev](https://news.ycombinator.com/item?id=34105770#34107671) | [next](https://news.ycombinator.com/item?id=34105770#34107605) [–]

  

Learn C using the classical books, but also care about how to mitigate security issues while coding in C.

"Secure Coding in C and C++"

[https://www.oreilly.com/library/view/secure-coding-in/978013...](https://www.oreilly.com/library/view/secure-coding-in/9780132981989/)

"The CERT C Secure Coding Standard"

[https://www.oreilly.com/library/view/the-cert-c/978032160319...](https://www.oreilly.com/library/view/the-cert-c/9780321603197/)

MISRA-C

[https://www.misra.org.uk/](https://www.misra.org.uk/)

Frama-C

[https://frama-c.com/](https://frama-c.com/)

Also have a look at the static analysis and warnings available across all major C compilers.

[reply](https://news.ycombinator.com/reply?id=34107590&goto=item%3Fid%3D34105770%2334107590)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107605&how=up&auth=c7753f7a15931a36df601991835ffa8bbfdae316&goto=item%3Fid%3D34105770#34107605)

[adamrt](https://news.ycombinator.com/user?id=adamrt) [2 days ago](https://news.ycombinator.com/item?id=34107605) | [prev](https://news.ycombinator.com/item?id=34105770#34107590) | [next](https://news.ycombinator.com/item?id=34105770#34109037) [–]

  

I had the same itch for a long time. I eventually took the free Harvard CS50 course. Its the video lectures you can watch on demand. It gets into python at the end which I skipped as I already knew python.

I really enjoyed it and think the teacher is great. I don't think all of it is necessary if you are already a developer, but I didn't have a traditional education so it filled in a few gaps for me too.

[https://learning.edx.org/course/course-v1:HarvardX+CS50+X/bl...](https://learning.edx.org/course/course-v1:HarvardX+CS50+X/block-v1:HarvardX+CS50+X+type@sequential+block@6ea5a2ca5fdf422b8c6b4c993fe75083/block-v1:HarvardX+CS50+X+type@vertical+block@ca72c7c0d4d449e69a5b8d23f0f6cb0b)

[reply](https://news.ycombinator.com/reply?id=34107605&goto=item%3Fid%3D34105770%2334107605)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109037&how=up&auth=2de10d7818e090a6fcfdeb0b2e685ab67e687589&goto=item%3Fid%3D34105770#34109037)

[blaerk](https://news.ycombinator.com/user?id=blaerk) [1 day ago](https://news.ycombinator.com/item?id=34109037) | [prev](https://news.ycombinator.com/item?id=34105770#34107605) | [next](https://news.ycombinator.com/item?id=34105770#34107864) [–]

  

Excellent choice! I started off 2022 with learning c and it’s a lot of fun!

As some other commenters, I really recommend Modern C by Jens Gustedt.

If you’re on Linux you can use man pages for most or all of the c stdlib, it’s a really convenient way to check up on stuff directly from the terminal, less tempting to rabbit hole into something else while in a browser.

I started contributing to a really small open source project, it was less daunting than a large well written, well formatted project and you get to run into “accidents”, figure them out, fixing them, learning from the experience. I realise this is a matter of personal preference, for me it’s a nice learning experience.

Also having (or making) friends who knows c is really helpful.

Good luck!

[reply](https://news.ycombinator.com/reply?id=34109037&goto=item%3Fid%3D34105770%2334109037)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107864&how=up&auth=1c56ed53f11597675c6edce09b83a4aab2aa74fb&goto=item%3Fid%3D34105770#34107864)

[grose](https://news.ycombinator.com/user?id=grose) [2 days ago](https://news.ycombinator.com/item?id=34107864) | [prev](https://news.ycombinator.com/item?id=34105770#34109037) | [next](https://news.ycombinator.com/item?id=34105770#34108535) [–]

  

This is the best reference site I've found: [https://en.cppreference.com/w/c](https://en.cppreference.com/w/c)

There's also man pages for all the libc stuff. Try `man 2 listen` or `man 3 printf`.

The VSCode plugin for C is pretty decent as well, the Microsoft one called 'C/C++'.

I'd also recommend doing a readthrough of the GNU Make docs. They are succinct and full of useful stuff: [https://www.gnu.org/software/make/manual/make.html](https://www.gnu.org/software/make/manual/make.html)

[reply](https://news.ycombinator.com/reply?id=34107864&goto=item%3Fid%3D34105770%2334107864)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108535&how=up&auth=7cfe3b58848c161ba1fa4acc2366d20ecf5d3cea&goto=item%3Fid%3D34105770#34108535)

[sfpotter](https://news.ycombinator.com/user?id=sfpotter) [2 days ago](https://news.ycombinator.com/item?id=34108535) | [prev](https://news.ycombinator.com/item?id=34105770#34107864) | [next](https://news.ycombinator.com/item?id=34105770#34111035) [–]

  

The suggested way to learn C in 2023 is the same as the suggested way to learn all languages for the past decade or so.

Do you like videos? Find videos on YouTube and watch them.

Do you like books? Go on Amazon and search for "C programming".

Do you like reading code? Go on GitHub and start reading the source for any project.

Or just find any online resource and use that.

The most important thing is to write code to do something you're interested in. If you're interested in solving puzzles, solve some puzzles. If you want to write a game, write a game.

Above all, you shouldn't waste time concerning yourself initially with what the "right" thing to do is. Do NOT let the perfect be the enemy of the good. It's OK to ignore the advice of anyone who tells you you need to be concerned with how to set up your project the right way, or tells you to avoid certain "dangerous" functions in the standard library. Just avoid all dogma.

Since you're learning, thinking about these things now is a huge waste of time. It's much better to write code that actually does things. If you thoroughly engage with this process, you will figure this stuff out on your own naturally. Or you won't, and then someone will correct you later, at which point you will be in a better position to appreciate their advice.

_Biggest time-waster:_ reading this thread. These threads are mainly a venue for people with minimal experience to regurgitate things they have read elsewhere.

[reply](https://news.ycombinator.com/reply?id=34108535&goto=item%3Fid%3D34105770%2334108535)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111035&how=up&auth=fbcb311a9e2491be690cce79ecf76ccc7369f177&goto=item%3Fid%3D34105770#34111035)

[Arch-TK](https://news.ycombinator.com/user?id=Arch-TK) [1 day ago](https://news.ycombinator.com/item?id=34111035) | [prev](https://news.ycombinator.com/item?id=34105770#34108535) | [next](https://news.ycombinator.com/item?id=34105770#34110585) [–]

  

I think a good place to start is to mostly ignore most/all the advice in this thread. I haven't obviously read every response but I have struggled to find any response near the top which isn't recommending something which will give poor results or is outright wrong.

Good resources:

- #c on libera.chat *

- [https://www.iso-9899.info/](https://www.iso-9899.info/)

Good books:

- C Programming: A Modern Approach - K. N. King

- K&R2 with the errata at hand and the guidance of someone experienced when doing all the exercises

Good advice:

- Don't write a single line for which you are not sure that the behaviour is either well defined by the standard or by your particular implementation (and if the latter is the case, ensure you are well aware of the implications and document them appropriately).

- C is not a "try it and see" language, trying it will lead to undefined behavior and buggy code which you might not even become aware of until you have "learned" the language.

- Lots of C out there is littered with poor practices or heavily specialised to the task/target set of architectures/situation and is therefore not a good example for learning. Feel free to read lots of C once you have the fundamentals down, not before.

Things to avoid:

- Every online tutorial (no I haven't seen them all but none of the popular ones seem to instill the right discipline and are usually full of errors, one of the ones linked from this thread has a bug in the hello world program in the first exercise).

- Most books on the matter.

- Writing C.

* disclaimer: I am one of the founders of this channel.

[reply](https://news.ycombinator.com/reply?id=34111035&goto=item%3Fid%3D34105770%2334111035)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110585&how=up&auth=a2e19645a4210e38e30e4da266c96a10d8bec0ce&goto=item%3Fid%3D34105770#34110585)

[dariosalvi78](https://news.ycombinator.com/user?id=dariosalvi78) [1 day ago](https://news.ycombinator.com/item?id=34110585) | [prev](https://news.ycombinator.com/item?id=34105770#34111035) | [next](https://news.ycombinator.com/item?id=34105770#34107133) [–]

  

C is a very minimal language, you can learn it in a day if you are already familiar with C-like syntax. I would recommend using it on embedded systems, such as an Arduino (technically uses C++, but it's mostly hidden), so you can appreciate both the speed and minimal memory footprint and you would be abstracted from the complexity of operating systems.

[reply](https://news.ycombinator.com/reply?id=34110585&goto=item%3Fid%3D34105770%2334110585)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111562&how=up&auth=1825fced3e51152ba8eac7590d3373f3f255839e&goto=item%3Fid%3D34105770#34111562)

[wink](https://news.ycombinator.com/user?id=wink) [1 day ago](https://news.ycombinator.com/item?id=34111562) | [parent](https://news.ycombinator.com/item?id=34105770#34110585) | [next](https://news.ycombinator.com/item?id=34105770#34107133) [–]

  

If by learn you mean rote-memorize some syntax. I doubt most people who have never touched it will manage to use or even understand enough about pointers and arrays and stuff in a single day. Get some stuff to compile maybe, yeah.

[reply](https://news.ycombinator.com/reply?id=34111562&goto=item%3Fid%3D34105770%2334111562)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107133&how=up&auth=8072a81aeff07b29aa340d0c39161b87ab34625c&goto=item%3Fid%3D34105770#34107133)

[dthakur](https://news.ycombinator.com/user?id=dthakur) [2 days ago](https://news.ycombinator.com/item?id=34107133) | [prev](https://news.ycombinator.com/item?id=34105770#34110585) | [next](https://news.ycombinator.com/item?id=34105770#34108052) [–]

  

Arduino programming can be a pleasant intermediate step in learning C/C++. Specializing specially in C can be easier after that experience.

[reply](https://news.ycombinator.com/reply?id=34107133&goto=item%3Fid%3D34105770%2334107133)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108052&how=up&auth=da5e6173cdfdc934a753d26ccb9a4c4adf46eec2&goto=item%3Fid%3D34105770#34108052)

[jeffreyrogers](https://news.ycombinator.com/user?id=jeffreyrogers) [2 days ago](https://news.ycombinator.com/item?id=34108052) | [prev](https://news.ycombinator.com/item?id=34105770#34107133) | [next](https://news.ycombinator.com/item?id=34105770#34107660) [–]

  

The C Programming Language is a decent overview of the language. After that using it in a real project is the best way to learn. If you know Linux pretty well and know what syscalls are then you might be interested in reading the Linux source code and writing kernel modules (they're written in C). If none of that sounds familiar to you the overhead is probably too high for going that route to make sense.

In general, the reason to use C is because you are doing something that interfaces with the OS. But the problem with software at this level is that because the abstractions are so low level unless you understand something about the design/architecture of the software the code is really hard to understand.

So I would pick some project that is interesting to you and that you have some familiarity with already and try to implement a simple version of that in C. It is actually not too hard to write a WebSocket or HTTP server in C, so if you are familiar with web programming, writing one of those might be a fun project.

[reply](https://news.ycombinator.com/reply?id=34108052&goto=item%3Fid%3D34105770%2334108052)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107660&how=up&auth=a2ed75b596c207580833275a1f1b4edacc683a0d&goto=item%3Fid%3D34105770#34107660)

[satvikpendem](https://news.ycombinator.com/user?id=satvikpendem) [2 days ago](https://news.ycombinator.com/item?id=34107660) | [prev](https://news.ycombinator.com/item?id=34105770#34108052) | [next](https://news.ycombinator.com/item?id=34105770#34108236) [–]

  

Read K&R. Yes, it's explicitly not modern C, but it enables you to think about C from a foundational level from which you can then go on to learn modern C. Build a project in C to learn how and why it works, and you'll understand how memory safe languages like Rust came about. This is what we did in college, started with C then moved on to Rust, it worked well.

[reply](https://news.ycombinator.com/reply?id=34107660&goto=item%3Fid%3D34105770%2334107660)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108395&how=up&auth=1eb83f6f6bdf24ad7d7a4e69b55a0a29052f2dd6&goto=item%3Fid%3D34105770#34108395)

[squarefoot](https://news.ycombinator.com/user?id=squarefoot) [2 days ago](https://news.ycombinator.com/item?id=34108395) | [parent](https://news.ycombinator.com/item?id=34105770#34107660) | [next](https://news.ycombinator.com/item?id=34105770#34108236) [–]

  

What book would you suggest to someone who started with the K&R 2nd ed, then did other things, and now wants to catch up on the more modern C standards, at least the more widely supported ones?

[reply](https://news.ycombinator.com/reply?id=34108395&goto=item%3Fid%3D34105770%2334108395)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110306&how=up&auth=f2937675d0876283e8587d3408d16739f1dc9ae5&goto=item%3Fid%3D34105770#34110306)

[queuebert](https://news.ycombinator.com/user?id=queuebert) [1 day ago](https://news.ycombinator.com/item?id=34110306) | [root](https://news.ycombinator.com/item?id=34105770#34107660) | [parent](https://news.ycombinator.com/item?id=34105770#34108395) | [next](https://news.ycombinator.com/item?id=34105770#34108236) [–]

  

21st Century C: [https://www.oreilly.com/library/view/21st-century-c/97814493...](https://www.oreilly.com/library/view/21st-century-c/9781449344382/)

[reply](https://news.ycombinator.com/reply?id=34110306&goto=item%3Fid%3D34105770%2334110306)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108236&how=up&auth=f982a6749cda7e29b939ee87fe307362cb2f9357&goto=item%3Fid%3D34105770#34108236)

[stusmall](https://news.ycombinator.com/user?id=stusmall) [2 days ago](https://news.ycombinator.com/item?id=34108236) | [prev](https://news.ycombinator.com/item?id=34105770#34107660) | [next](https://news.ycombinator.com/item?id=34105770#34107790) [–]

  

If you are learning for historical context and curiosity reasons, K&R is a good place to start. It's short easy to ready but is missing a lot of important context. A lot has changed in the world since it was released.

If you are learning for practical reasons, like you have a project in mind, then it will be dependent on that project. C patterns and idioms vary a lot per project and time period. For a newer, more practical book I liked C Programming: A Modern Approach. Just don't go into it not expecting it to be _too_ modern. The subject matter is still still C. After that you'll still probably have a lot to learn for whatever project you have in mind. C has a lot of quirks, peculiarities and shortcomings. There is no standard way of dealing with them so each project has their own best practices or set of helper APIs/tools with some minimal overlap.

[reply](https://news.ycombinator.com/reply?id=34108236&goto=item%3Fid%3D34105770%2334108236)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107790&how=up&auth=00424fca5fe3ac0389a6a41492974b357762d88a&goto=item%3Fid%3D34105770#34107790)

[hombre_fatal](https://news.ycombinator.com/user?id=hombre_fatal) [2 days ago](https://news.ycombinator.com/item?id=34107790) | [prev](https://news.ycombinator.com/item?id=34105770#34108236) | [next](https://news.ycombinator.com/item?id=34105770#34107106) [–]

  

I learned a lot getting an Arduino and building stuff with it in C. With zero C and zero Arduino experience I wired together an e-ink display, SD card reader, and some buttons to build a basic e-reader.

I would have never been bothered to build something in C in an environment where I can reach for any higher level language.

[reply](https://news.ycombinator.com/reply?id=34107790&goto=item%3Fid%3D34105770%2334107790)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108057&how=up&auth=5cd880a11355bb6d8951fb26972dbb790821c612&goto=item%3Fid%3D34105770#34108057)

[kevin_thibedeau](https://news.ycombinator.com/user?id=kevin_thibedeau) [2 days ago](https://news.ycombinator.com/item?id=34108057) | [parent](https://news.ycombinator.com/item?id=34105770#34107790) | [next](https://news.ycombinator.com/item?id=34105770#34107106) [–]

  

The Arduino ecosystem is heavily dependent on C++.

[reply](https://news.ycombinator.com/reply?id=34108057&goto=item%3Fid%3D34105770%2334108057)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107106&how=up&auth=c44496efc3a6659936b5adc127f5faf8a760c5fc&goto=item%3Fid%3D34105770#34107106)

[Zanfa](https://news.ycombinator.com/user?id=Zanfa) [2 days ago](https://news.ycombinator.com/item?id=34107106) | [prev](https://news.ycombinator.com/item?id=34105770#34107790) | [next](https://news.ycombinator.com/item?id=34105770#34107939) [–]

  

I’d also recommend reading The Linux Programming Interface [1] as an alternative. While it describes the Linux API, it does so through tons of examples in C, so you can see how things are actually done in C-land.

[1] [https://man7.org/tlpi/](https://man7.org/tlpi/)

[reply](https://news.ycombinator.com/reply?id=34107106&goto=item%3Fid%3D34105770%2334107106)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107939&how=up&auth=374e6a8e9dbf40a6cf175ca0307e2e06217bebc8&goto=item%3Fid%3D34105770#34107939)

[Koshkin](https://news.ycombinator.com/user?id=Koshkin) [2 days ago](https://news.ycombinator.com/item?id=34107939) | [prev](https://news.ycombinator.com/item?id=34105770#34107106) | [next](https://news.ycombinator.com/item?id=34105770#34109203) [–]

  

As a C/C++ dev, I want to learn Java. With its huge and sophisticated ecosystem, platform independence, the safety of the language, good efficiency, it almost sounds like one must be insane not to use it these days... (Sorry for being off-topic here.)

[reply](https://news.ycombinator.com/reply?id=34107939&goto=item%3Fid%3D34105770%2334107939)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108035&how=up&auth=16d95eebea18e9ceda808ea2e609a79c42e83d0f&goto=item%3Fid%3D34105770#34108035)

[kevin_thibedeau](https://news.ycombinator.com/user?id=kevin_thibedeau) [2 days ago](https://news.ycombinator.com/item?id=34108035) | [parent](https://news.ycombinator.com/item?id=34105770#34107939) | [next](https://news.ycombinator.com/item?id=34105770#34109203) [–]

  

You'd be better off learning C# since it learned from Java's mistakes.

[reply](https://news.ycombinator.com/reply?id=34108035&goto=item%3Fid%3D34105770%2334108035)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109203&how=up&auth=882f745621d6455d35c089e054839c4232718ff1&goto=item%3Fid%3D34105770#34109203)

[phoehne](https://news.ycombinator.com/user?id=phoehne) [1 day ago](https://news.ycombinator.com/item?id=34109203) | [prev](https://news.ycombinator.com/item?id=34105770#34107939) | [next](https://news.ycombinator.com/item?id=34105770#34109326) [–]

  

I program in C professionally. It's a good language to learn how processors actually work. It's on the embedded side, where there are other options but either they're not mature (Rust) or were never really taken up (Ada). Embedded C it's just you and the hardware. It's a good way to learn about what's actually going on. Anything, like an Arduino or Arm cortex is great. Small warning about ESP32 boards - some are based on the Extensa cores, which means you need to use their compilers. Others are based on Risc V - do your research.

[reply](https://news.ycombinator.com/reply?id=34109203&goto=item%3Fid%3D34105770%2334109203)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109326&how=up&auth=29a5d44743ac88e2ccc070a8bb4bbc91ea613ecc&goto=item%3Fid%3D34105770#34109326)

[tmtvl](https://news.ycombinator.com/user?id=tmtvl) [1 day ago](https://news.ycombinator.com/item?id=34109326) | [prev](https://news.ycombinator.com/item?id=34105770#34109203) | [next](https://news.ycombinator.com/item?id=34105770#34106993) [–]

  

I am fond of _Notes on Data Structures and Programming Techniques_ by James Aspnes: [https://cs.yale.edu/homes/aspnes/classes/223/notes.html](https://cs.yale.edu/homes/aspnes/classes/223/notes.html) it has a short and sweet overview of C as well as a sections on problems with C. Another useful resource is Libc's info manual (in Emacs: _C-h i m Libc<return>_), it's a well-formatted, clear reference.

[reply](https://news.ycombinator.com/reply?id=34109326&goto=item%3Fid%3D34105770%2334109326)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106993&how=up&auth=f61dde6116ec0aba7d934c06594829ab7180b52e&goto=item%3Fid%3D34105770#34106993)

[nathias](https://news.ycombinator.com/user?id=nathias) [2 days ago](https://news.ycombinator.com/item?id=34106993) | [prev](https://news.ycombinator.com/item?id=34105770#34109326) | [next](https://news.ycombinator.com/item?id=34105770#34105988) [–]

  

[https://beej.us/guide/bgc/html/index-wide.html](https://beej.us/guide/bgc/html/index-wide.html)

[reply](https://news.ycombinator.com/reply?id=34106993&goto=item%3Fid%3D34105770%2334106993)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34105988&how=up&auth=2c9914dc5dcb03b8fa13172ff72c002ae6c54145&goto=item%3Fid%3D34105770#34105988)

[iExploder](https://news.ycombinator.com/user?id=iExploder) [2 days ago](https://news.ycombinator.com/item?id=34105988) | [prev](https://news.ycombinator.com/item?id=34105770#34106993) | [next](https://news.ycombinator.com/item?id=34105770#34106955) [–]

  

maybe try some embedded development, like raspberry, light up some leds. Embedded development is fun, educational and one of the last bastions of programming where C makes sense.

[reply](https://news.ycombinator.com/reply?id=34105988&goto=item%3Fid%3D34105770%2334105988)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106180&how=up&auth=f321c9ac2bf6ee2dfe443f4b6279b370cfe05bd4&goto=item%3Fid%3D34105770#34106180)

[5e92cb50239222b](https://news.ycombinator.com/user?id=5e92cb50239222b) [2 days ago](https://news.ycombinator.com/item?id=34106180) | [parent](https://news.ycombinator.com/item?id=34105770#34105988) | [next](https://news.ycombinator.com/item?id=34105770#34106955) [–]

  

ESP32 is quite nice. It is much lower level than Raspberry Pi or similar SBCs, and C really makes sense there.

It ships with pretty high-level libraries (connecting to a Wi-Fi network is ~200 lines of C code), and you'll be running on top of a small operating system (FreeRTOS). But you still have full access to hardware and all the fun of working directly with it. The dev boards are quite cheap at $3-4 per piece, and have lots of peripherals (including Wi-Fi and Bluetooth).

Have a look for ESP-WROOM-32D at aliexpress, or whatever the platform of choice in your country is. Some of the newer boards are built on RISC-V, if that's your thing.

Their SDK also has first-class Linux support (and the other two OSs, obviously).

[https://docs.espressif.com/projects/esp-idf/en/latest/esp32/](https://docs.espressif.com/projects/esp-idf/en/latest/esp32/)

[reply](https://news.ycombinator.com/reply?id=34106180&goto=item%3Fid%3D34105770%2334106180)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107239&how=up&auth=d372f6853b4dd7383b02634e63be076ba354eb46&goto=item%3Fid%3D34105770#34107239)

[II2II](https://news.ycombinator.com/user?id=II2II) [2 days ago](https://news.ycombinator.com/item?id=34107239) | [root](https://news.ycombinator.com/item?id=34105770#34105988) | [parent](https://news.ycombinator.com/item?id=34105770#34106180) | [next](https://news.ycombinator.com/item?id=34105770#34106955) [–]

  

I suspect they meant the Pico/RP2040.

Depending upon their goals, something like at ATmega may be a better choice. ESP32 and RP2040 development places a lot of emphasis on using libraries. That's fine, but it doesn't really highlight how C is different from most other languages. In contrast, there is a lot of information on low level programming for the ATmega. The RP2040 would be my second choice. I believe the libraries are entirely open source, which means you can read the datasheet and can reference the library source to see how something is implemented in C.

Granted, the ESP32 is a lot more fun if the intent is to build something.

[reply](https://news.ycombinator.com/reply?id=34107239&goto=item%3Fid%3D34105770%2334107239)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106955&how=up&auth=0eb9277d776a5786b9c9c58956fd0509fd91baed&goto=item%3Fid%3D34105770#34106955)

[linhns](https://news.ycombinator.com/user?id=linhns) [2 days ago](https://news.ycombinator.com/item?id=34106955) | [prev](https://news.ycombinator.com/item?id=34105770#34105988) | [next](https://news.ycombinator.com/item?id=34105770#34106082) [–]

  

K&R book. Nothing come close although it lacks on modern C.

[reply](https://news.ycombinator.com/reply?id=34106955&goto=item%3Fid%3D34105770%2334106955)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106082&how=up&auth=ec6f72016dc0622684bc2fb552e8a42a45f1611d&goto=item%3Fid%3D34105770#34106082)

[phendrenad2](https://news.ycombinator.com/user?id=phendrenad2) [2 days ago](https://news.ycombinator.com/item?id=34106082) | [prev](https://news.ycombinator.com/item?id=34105770#34106955) | [next](https://news.ycombinator.com/item?id=34105770#34114866) [–]

  

Well, GNU/linux is more "C plus GCC" than pure C. There are a lot of fancy things that GCC gives you that you'd otherwise have to do explicitly.

[reply](https://news.ycombinator.com/reply?id=34106082&goto=item%3Fid%3D34105770%2334106082)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106908&how=up&auth=e62d94400bf6dea96aef5cfefdcb9558e0488b09&goto=item%3Fid%3D34105770#34106908)

[torstenvl](https://news.ycombinator.com/user?id=torstenvl) [2 days ago](https://news.ycombinator.com/item?id=34106908) | [parent](https://news.ycombinator.com/item?id=34105770#34106082) | [next](https://news.ycombinator.com/item?id=34105770#34106626) [–]

  

C has the unique distinction of being criticized on the one hand for the complexities of vendor-specific extensions and portability concerns... and on the other for being insufficiently like Rust or Go, which are defined _entirely_ by their proprietary implementations.

I don't think I've ever seen this criticism leveled elsewhere. Do people criticize JavaScript programming because the ECMAScript standard doesn't give you everything you could want?

[reply](https://news.ycombinator.com/reply?id=34106908&goto=item%3Fid%3D34105770%2334106908)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107021&how=up&auth=403d777105d39256d3e419dbd5ed6f8529ab27d9&goto=item%3Fid%3D34105770#34107021)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [2 days ago](https://news.ycombinator.com/item?id=34107021) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34106908) | [next](https://news.ycombinator.com/item?id=34105770#34107030) [–]

  

There is no meaningful sense in which either Rust or Go’s reference compilers is “proprietary.” If you’re trying to say that there is no ISO standard for the two, that’s an entirely different problem (and one shared by the overwhelming majority of programming languages).

[reply](https://news.ycombinator.com/reply?id=34107021&goto=item%3Fid%3D34105770%2334107021)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108138&how=up&auth=de410004e3b5fbdd4a9c1ab52fed0a1481384c15&goto=item%3Fid%3D34105770#34108138)

[torstenvl](https://news.ycombinator.com/user?id=torstenvl) [2 days ago](https://news.ycombinator.com/item?id=34108138) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34107021) | [next](https://news.ycombinator.com/item?id=34105770#34107030) [–]

  

> _There is no meaningful sense in which either Rust or Go’s reference compilers is “proprietary.”_

Nobody is talking about the compilers. This, right here — the fact that you are conflating a language and one implementation — is _the whole point_ and is indicative of the tight control the respective organizations keep over their languages.

The direction of the languages is entirely up to their private owners, and that exclusivity is enforced via intellectual property law. They are exactly as proprietary as RHEL or Unix.

[reply](https://news.ycombinator.com/reply?id=34108138&goto=item%3Fid%3D34105770%2334108138)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108292&how=up&auth=f54088b0ade7e32a4f92f8f5a689641cd0ab8d13&goto=item%3Fid%3D34105770#34108292)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [2 days ago](https://news.ycombinator.com/item?id=34108292) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34108138) | [next](https://news.ycombinator.com/item?id=34105770#34108640) [–]

  

This is not a definition of "proprietary" used anywhere in the open source community. Private direction (which Rust notably does not have) does not somehow erase the fact that both are permissively licensed and have open codebases that you can use and contribute to to your heart's content.

I don't even know what you're referring to with IP law. The projects' names and logos?

[reply](https://news.ycombinator.com/reply?id=34108292&goto=item%3Fid%3D34105770%2334108292)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34119774&how=up&auth=82b73f7a96a5367ccaae5778048478bf74ab3a53&goto=item%3Fid%3D34105770#34119774)

[lelanthran](https://news.ycombinator.com/user?id=lelanthran) [23 hours ago](https://news.ycombinator.com/item?id=34119774) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34108292) | [next](https://news.ycombinator.com/item?id=34105770#34108920) [–]

  

> I don't even know what you're referring to with IP law. The projects' names and logos?

I think it's the fact that if you took the Rust codebase and forked[1] it to accept new keywords, the resulting language that it accepts is most definitely _not Rust_.

If you did the same to a C compiler[2], the resulting compiler would still be called a C compiler by everyone.

[1] Assuming that the license allows it.

[2] One who's licence allowed this.

[reply](https://news.ycombinator.com/reply?id=34119774&goto=item%3Fid%3D34105770%2334119774)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34119911&how=up&auth=b00ca41c6ba365638ab7bf293a55f27a0d682eef&goto=item%3Fid%3D34105770#34119911)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [23 hours ago](https://news.ycombinator.com/item?id=34119911) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34119774) | [next](https://news.ycombinator.com/item?id=34105770#34108920) [–]

  

Yes, the license allows it. Rust's compiler is _more permissively licensed_ than GCC, and is roughly as permissively licensed as LLVM.

This argument does not make sense: if you fork GCC or LLVM and change the language frontend for C, the thing you are parsing is _no longer standard C_. This is exactly the same as is true for Rust, except that the reference is the the living standard rather than an ISO standard.

By the original commenter's logic, the modified C compiler would effectively be proprietary: it no longer has a public, liberally licensed specification that it conforms to. But this isn't the way _anybody_ in the open source community uses the word "proprietary," because it precludes _any_ open-source work that doesn't have a corresponding formal specification.

[reply](https://news.ycombinator.com/reply?id=34119911&goto=item%3Fid%3D34105770%2334119911)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34120088&how=up&auth=07721631885b2c6d2a773dbffc3f1691bcdb74ef&goto=item%3Fid%3D34105770#34120088)

[lelanthran](https://news.ycombinator.com/user?id=lelanthran) [22 hours ago](https://news.ycombinator.com/item?id=34120088) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34119911) | [next](https://news.ycombinator.com/item?id=34105770#34121066) [–]

  

> This argument does not make sense: if you fork GCC or LLVM and change the language frontend for C, the thing you are parsing is no longer standard C.

It doesn't make sense to you, certainly.

To the type of person who've used a dozen different C compilers, all of them containing extensions to the language, and all of them still advertised as "C compilers", the argument makes perfect sense.

I mean, the input to the GCC C compiler is, _by default_, C with extensions, and yet even you still called it a C frontend, not a "SomeOtherLanguage" frontend.

You may not agree that the argument is valid or sound, but you can't with a straight face say that it is an unreasonable position to take.

[reply](https://news.ycombinator.com/reply?id=34120088&goto=item%3Fid%3D34105770%2334120088)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34120361&how=up&auth=faaa1a252ef5973b0c00e8d3dbffe5d4395d3f66&goto=item%3Fid%3D34105770#34120361)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [22 hours ago](https://news.ycombinator.com/item?id=34120361) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34120088) | [next](https://news.ycombinator.com/item?id=34105770#34121066) [–]

  

> You may not agree that the argument is valid or sound, but you can't with a straight face say that it is an unreasonable position to take.

That's not the argument I'm advancing: I'm saying it _does not make sense_ to call that modified C compiler "proprietary."

I don't care if you call it "C" or not, because everybody who writes C understands that compiler-specific extensions are commonplace. But the existence of non-formally-specified extensions doesn't somehow override the permissively licensed nature of the compiler.

Similarly, going back to the original argument: it does not make sense to call Rust "proprietary" because it only has a living specification, rather than a formal one. It's permissively licensed, and you can modify it to your heart's content. The only reason people wouldn't call your changes "Rust" is because the community doesn't share the same expectations re: vendor extensions that C's community does.

[reply](https://news.ycombinator.com/reply?id=34120361&goto=item%3Fid%3D34105770%2334120361)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34121438&how=up&auth=48b1f26cfc69c29fb144384c7b594d98a3bb4695&goto=item%3Fid%3D34105770#34121438)

[lelanthran](https://news.ycombinator.com/user?id=lelanthran) [20 hours ago](https://news.ycombinator.com/item?id=34121438) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34120361) | [next](https://news.ycombinator.com/item?id=34105770#34121066) [–]

  

> it does not make sense to call Rust "proprietary" because it only has a living specification,

It doesn't have a "living specification", it has a reference implementation. There's a difference. HTML is a living specification without a reference implementation. Rust (and many other languages) have a reference implementation without a living specification.

> The only reason people wouldn't call your changes "Rust" is because the community doesn't share the same expectations

That's not true; you cannot call it Rust because _Rust, the programming language_, is _trademarked_.

It is _protected_, you understand? There are legal barriers to calling your NewLanguage "Rust", as explained by the foundation themselves over here: [https://foundation.rust-lang.org/policies/logo-policy-and-me...](https://foundation.rust-lang.org/policies/logo-policy-and-media-guide/)

Writing your own language and calling it Rust is legally forbidden.[1] You are, of course, free to use the permissive _copyright_ but the name is trademarked.

Now, with C, OTOH, as an ISO standard it has no trademark so this argument does not even arise.

[1] As far as I can tell, it would be legally forbidden even if your new implementation of Rust behaves identical to the reference implementation. The best you can do is call it Rust compatible.

[reply](https://news.ycombinator.com/reply?id=34121438&goto=item%3Fid%3D34105770%2334121438)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34121066&how=up&auth=538dc4f49fb29073df016e3b33ce887fe46efb31&goto=item%3Fid%3D34105770#34121066)

[torstenvl](https://news.ycombinator.com/user?id=torstenvl) [20 hours ago](https://news.ycombinator.com/item?id=34121066) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34119911) | [prev](https://news.ycombinator.com/item?id=34105770#34120088) | [next](https://news.ycombinator.com/item?id=34105770#34108920) [–]

  

There is a difference between open _software_ and open _standards_. The two are orthogonal. Open software can use closed standards (though must usually reverse-engineer them); closed software can use open standards. See, e.g., LibreOffice support for .doc files; Microsoft's TCP/IP stack. Please make an attempt to engage with intellectual honesty, and stop intentionally conflating the two.

The issue at bar is whether Rust _the language_ is proprietary, especially as compared to C. Something is proprietary within the standard meaning of the term if it "is used, produced, or marketed under exclusive legal right of the inventor or maker." Merriam Webster, Online Edition, available at [https://www.merriam-webster.com/dictionary/proprietary](https://www.merriam-webster.com/dictionary/proprietary).

In this case, The Rust Foundation maintains and exercises exclusive legal rights over Rust that nobody exercises over C. If I make a C compiler with extensions, I can still call it C and I can still participate in C governance. C is an open standard. See [https://www.open-std.org/](https://www.open-std.org/). However, the same is not true for Rust, where governance happens entirely under the auspices of one vendor, who uses trademark law to ensure that no competitors can exist without its explicit approval. Specifically, any Rust variants other than those required for porting require explicit written permission. See [https://foundation.rust-lang.org/policies/logo-policy-and-me...](https://foundation.rust-lang.org/policies/logo-policy-and-media-guide/). Furthermore, even if I accepted my impotence in the governance process and even if I resolved to call my Rust variant something else (FeO?), competition in the market would _still_ be chilled by The Rust Foundation's refusal to commit to fair, reasonable, and non-discriminatory (FRAND) licensing (something that is required as part of the ISO standardization process). If a competitor ever did arise, patenting the borrow checker mechanism would be all it would take to completely crush them.

I am glad that rustc is open source. I am glad that The Rust Foundation is a non-profit. I accept on faith that they want to do the right thing. _However_, they have deliberately created a monoculture that they enforce through IP law. Rust _is_ proprietary, i.e., it is "something that is used, produced, or marketed under exclusive legal right of the inventor or maker." There are no open standards, there is no FRAND licensing obligation, there is no allowance for any Rust variants other than those explicitly permitted by the primary vendor of Rust compilers.

You brought up Python before, which is appropriate: Python explicitly _is_ a "benevolent dictator for life" language. The difference with Rust is that the dictator is a legal personhood instead of a single living, breathing person. But dictatorships, even benevolent and well-intentioned ones, are ill-suited to the requirements of stability and predictability that are so vital to systems programming.

Imagine, God forbid, that The Rust Foundation goes defunct at some point in the next 100 years, and its IP bought up by the next generation's SCO or Novell or whatever, and they decide to patent the borrow checker and whatever other innovations are in Rust. Do you think that would be a bad thing?

Because it can only be bad _if you agree that Rust is proprietary_. If Rust were an open standard (including the borrow checker as part of the standardized compilation process and subject to mandatory FRAND licensing), then that outcome would be mitigated against.

[reply](https://news.ycombinator.com/reply?id=34121066&goto=item%3Fid%3D34105770%2334121066)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34127866&how=up&auth=4d83ead27eba388bb541e2c799876fb1adfcd8aa&goto=item%3Fid%3D34105770#34127866)

[mustache_kimono](https://news.ycombinator.com/user?id=mustache_kimono) [2 hours ago](https://news.ycombinator.com/item?id=34127866) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34121066) | [next](https://news.ycombinator.com/item?id=34105770#34121352) [–]

  

> If a competitor ever did arise, patenting the borrow checker mechanism would be all it would take to completely crush them.

Apache 2 includes a patent grant? So I think most of this comment is clearly erroneous.

[reply](https://news.ycombinator.com/reply?id=34127866&goto=item%3Fid%3D34105770%2334127866)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34121352&how=up&auth=8bca41bb89150ee15313439cb597c7bf73290fe3&goto=item%3Fid%3D34105770#34121352)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [20 hours ago](https://news.ycombinator.com/item?id=34121352) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34121066) | [prev](https://news.ycombinator.com/item?id=34105770#34127866) | [next](https://news.ycombinator.com/item?id=34105770#34108920) [–]

  

This is a very large comment, filled with legitimate observations about how Rust (and most other programming language communities) choose to defend their name, logo, etc.

It has _absolutely_ nothing to do with how "proprietary" is used in the open source community. To make it as simple as possible: if your use of the word "proprietary" includes Python and Rust but excludes Borland C, you've failed to use to word in a way that is substantive and useful to _everyone else_ in this community.

The implied risk of Rust's trademark terms is entirely ahistorical: C did not win as a systems programming language because it satisfied FRAND or any other standard for openness. It won because Worse Is Better, _in spite_ of onerous licensing restrictions by companies with extremely expensive and proprietary C compilers.

[reply](https://news.ycombinator.com/reply?id=34121352&goto=item%3Fid%3D34105770%2334121352)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34121575&how=up&auth=4cb66bb695865da34adda8d9afa49654ca351b81&goto=item%3Fid%3D34105770#34121575)

[shrimp_emoji](https://news.ycombinator.com/user?id=shrimp_emoji) [19 hours ago](https://news.ycombinator.com/item?id=34121575) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34121352) | [next](https://news.ycombinator.com/item?id=34105770#34108920) [–]

  

>_It won because Worse Is Better_

I don't think this is why. I think it "won" because it was the best "high-level" language of the time and exploded to fill a huge niche, historically. In other words, historical luck.

I think that's also how JavaScript won.

Since it's always luck, and since the luckiest thing is rarely the "best" thing, it just looks like it's a selection for ironically worse quality from the outside.

[reply](https://news.ycombinator.com/reply?id=34121575&goto=item%3Fid%3D34105770%2334121575)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34122008&how=up&auth=89e3f7e7b4617ed68bb6fbdf3f3b294b64b25807&goto=item%3Fid%3D34105770#34122008)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [19 hours ago](https://news.ycombinator.com/item?id=34122008) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34121575) | [next](https://news.ycombinator.com/item?id=34105770#34108920) [–]

  

"Worse is Better" isn't meant to be derogatory: it's the term of art used by the Unix world to describe Unix and C's own adoption arc[1].

[1]: [https://cs.stanford.edu/people/eroberts/cs201/projects/2010-...](https://cs.stanford.edu/people/eroberts/cs201/projects/2010-11/WorseIsBetter/index43bb.html?title=Main_Page&oldid=86)

[reply](https://news.ycombinator.com/reply?id=34122008&goto=item%3Fid%3D34105770%2334122008)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108920&how=up&auth=d5f692d20090db05f76a23b7410a9d5cb97fb83d&goto=item%3Fid%3D34105770#34108920)

[torstenvl](https://news.ycombinator.com/user?id=torstenvl) [1 day ago](https://news.ycombinator.com/item?id=34108920) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34108292) | [prev](https://news.ycombinator.com/item?id=34105770#34119774) | [next](https://news.ycombinator.com/item?id=34105770#34108640) [–]

  

> _both are permissively licensed_

No, the _code_ of the _implementation_ is permissively licensed.

Do I really have to explain, on HN of all places, the importance of distinguishing between implementation and interface?

[reply](https://news.ycombinator.com/reply?id=34108920&goto=item%3Fid%3D34105770%2334108920)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109193&how=up&auth=16c252e1c690a74913276925a3168c040e494cb6&goto=item%3Fid%3D34105770#34109193)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [1 day ago](https://news.ycombinator.com/item?id=34109193) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34108920) | [next](https://news.ycombinator.com/item?id=34105770#34108640) [–]

  

This is incorrect on multiple levels: there is no organization that has "exclusive" control over Go, much less Rust. There's a company (Google) that _made_ Go, but they have no control over the source code itself beyond the terms of the non-proprietary license they provide it under[1]. Even the things that Google does w/r/t Go are not even remotely controversial in the context of the definition of proprietary software, like CLAs. Linux has a CLA!

Comparing this to Borland or Microsoft is bizarre: Microsoft and Borland actually _do_ have proprietary C compilers that they do not make publicly available in source code form. Those compilers in turn have all kinds of non-standard, proprietary extensions.

The _one_ thing that sets C apart from Rust and Go is the presence of an official ISO standard. But this isn't material to the definition of "proprietary" in the open source world. And, well, that ISO standard is itself copyrighted.

[1]: [https://github.com/golang/go/blob/master/LICENSE](https://github.com/golang/go/blob/master/LICENSE)

Edit: You shouldn't delete and re-write your comments. Here's the original comment I responded to:

> That's a dubious claim, and irrelevant besides. They are proprietary, in the standard meaning of the word, because a single private organization has exclusive control over each, but whatever word you choose to describe that is up to you.

> Google owns Go in a way that has never been true of Microsoft (or Borland or GNU or Apple...) when it comes to C.

Edit 2: For the edited comment: you're setting up a distinction that does not matter for programming languages that do not have formal specifications. There is no "interface" to distinguish against, and the claim that Rust (or Go, or Python, or any other language without an ISO spec) is tacitly reserving all rights to the design of the language itself is completely unsubstantiated.

[reply](https://news.ycombinator.com/reply?id=34109193&goto=item%3Fid%3D34105770%2334109193)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109303&how=up&auth=4d69a08d5f36b4643e48c3ddabf9e98a2be365ed&goto=item%3Fid%3D34105770#34109303)

[torstenvl](https://news.ycombinator.com/user?id=torstenvl) [1 day ago](https://news.ycombinator.com/item?id=34109303) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34109193) | [next](https://news.ycombinator.com/item?id=34105770#34109887) [–]

  

Your position that not having a standard is irrelevant for languages that don't have a standard is circular and bizarre.

The fact that you keep focusing on whether _compilers_ are proprietary instead of whether Rust and Go are proprietary is starting to seem deliberately evasive. You're free to extol the virtues of open source software, and I agree with you, but it's utterly beside the point.

.rs is analogous to .doc. The "right" way to handle it is defined solely with respect to the reference implementation. I view that as a problem. You may disagree, and that's fine, but don't pretend this is a conversation about something else.

[reply](https://news.ycombinator.com/reply?id=34109303&goto=item%3Fid%3D34105770%2334109303)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110184&how=up&auth=f8e6bf631333b06286db0283e716d06803fdc6f1&goto=item%3Fid%3D34105770#34110184)

[mustache_kimono](https://news.ycombinator.com/user?id=mustache_kimono) [1 day ago](https://news.ycombinator.com/item?id=34110184) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34109303) | [next](https://news.ycombinator.com/item?id=34105770#34109368) [–]

  

> The fact that you keep focusing on whether compilers are proprietary instead of whether Rust and Go are proprietary is starting to seem deliberately evasive.

"You keep using that word. It don't think you know what it means." -- Inigo Montoya

This is just bizarre. If MIT/Apache-licensed Rust is proprietary, then anything could be? AFAIK there is no POSIX definition for epoll/io_uring. I suppose epoll is proprietary too?

proprietary: "2: something that is used, produced, or marketed under exclusive legal right of the inventor or maker"

Can Rust be forked, modified, used/sold and distributed with your changes? If yes that hardly seems like an exclusive legal right. Your problem seems to be that people trust the Rust project, and don't trust you or a central committee (yet). But that doesn't make a project proprietary. It makes it implementation defined, which has its own problems, but those are not the same problems of as those of proprietary software.

[reply](https://news.ycombinator.com/reply?id=34110184&goto=item%3Fid%3D34105770%2334110184)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34121483&how=up&auth=216a58d4695d70a7836239b4de5eae67de9f65bb&goto=item%3Fid%3D34105770#34121483)

[torstenvl](https://news.ycombinator.com/user?id=torstenvl) [20 hours ago](https://news.ycombinator.com/item?id=34121483) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34110184) | [next](https://news.ycombinator.com/item?id=34105770#34109368) [–]

  

> _MIT/Apache-licensed Rust_

Rust is not MIT/Apache-licensed, rustc is.

Rust is licensed under terms that require explicit permission for use outside of a few very narrowly-defined circumstances.

[reply](https://news.ycombinator.com/reply?id=34121483&goto=item%3Fid%3D34105770%2334121483)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34123758&how=up&auth=127a0ebd0814dc2978ebe9263074cf94bcb34cc8&goto=item%3Fid%3D34105770#34123758)

[mustache_kimono](https://news.ycombinator.com/user?id=mustache_kimono) [14 hours ago](https://news.ycombinator.com/item?id=34123758) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34121483) | [next](https://news.ycombinator.com/item?id=34105770#34109368) [–]

  

> Rust is licensed under terms that require explicit permission for use outside of a few very narrowly-defined circumstances.

I think you mean the Rust trademark, and sorry, but you don't get to use IBM or Google's mark willy-nilly either? I'm not sure it's a realistic expectation on your part to be able to market your changes under the other guy's mark.

[reply](https://news.ycombinator.com/reply?id=34123758&goto=item%3Fid%3D34105770%2334123758)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109368&how=up&auth=a6da30bf491642ebcf3d3df61758e1891e754b89&goto=item%3Fid%3D34105770#34109368)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [1 day ago](https://news.ycombinator.com/item?id=34109368) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34109303) | [prev](https://news.ycombinator.com/item?id=34105770#34110184) | [next](https://news.ycombinator.com/item?id=34105770#34109887) [–]

  

Where's the circularity? This is the claim: you can't claim that a nonexistent thing is proprietary. Rust and Go do not (to my knowledge) have any sort of standard other than the _living standard_ of the reference implementation, which is itself not proprietary. _Anybody_ is welcome to copy or modify that living standard, per the terms of the permissive license.

This is the _overwhelmingly common_ state of affairs for programming languages: Python, PHP, Go, Rust, etc. do not have formal standards in the sense of C, but instead of a living standard in the form of their reference implementation. The reference implementation is permissively licensed, so the "standard," insofar as it can be said to exist, is also permissively licensed.

Edit: Again: this is just not what the definition of "proprietary" is, in the context of open source. There is a legitimate grievance that you're expressing about software standards and implementation diversity, but it has _absolutely no bearing_ on whether the reference implementation is proprietary or not.

By analogy: I, an open source developer, create a new file format $FOO and release my reference implementation of a $FOO parser under a permissive software license. Am I somehow _compelled_ to spend time writing a formal specification of $FOO to make it non-proprietary? It might be _nice_ for me to do, and nothing stops _anybody else_ from converting my living standard into a formal one, but I am absolutely not compelled to do it.

[reply](https://news.ycombinator.com/reply?id=34109368&goto=item%3Fid%3D34105770%2334109368)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109887&how=up&auth=faab6343ebbec5455252d6c733957384d78157f7&goto=item%3Fid%3D34105770#34109887)

[0xjnml](https://news.ycombinator.com/user?id=0xjnml) [1 day ago](https://news.ycombinator.com/item?id=34109887) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34109193) | [prev](https://news.ycombinator.com/item?id=34105770#34109303) | [next](https://news.ycombinator.com/item?id=34105770#34108640) [–]

  

All commits to the Go repository must be approved by two Google employees: [https://groups.google.com/g/golang-dev/c/K7oGURi0wTM/m/3mg1O...](https://groups.google.com/g/golang-dev/c/K7oGURi0wTM/m/3mg1OwvqAQAJ)

[reply](https://news.ycombinator.com/reply?id=34109887&goto=item%3Fid%3D34105770%2334109887)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109919&how=up&auth=65ff802a9d6774b3cdd6ee89deb4635bd1ff85c2&goto=item%3Fid%3D34105770#34109919)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [1 day ago](https://news.ycombinator.com/item?id=34109919) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34109887) | [next](https://news.ycombinator.com/item?id=34105770#34108640) [–]

  

Yes. Again, Google can do whatever it likes with its repositories. This is _entirely_ unrelated to whether or not the software is proprietary.

Compare Google's practices here to Qt's[1], CMake's[2], or any other large open source project that is primarily maintained by a single company.

To be abundantly clear: I don't think this is a _good_ way to run an open source project, but it is a _legitimate_ way to run one. So long as the source code and standards (insofar as they exist) are themselves permissive, then the community details are just that.

[1]: [https://wiki.qt.io/Qt_Contribution_Guidelines](https://wiki.qt.io/Qt_Contribution_Guidelines)

[2]: [https://github.com/Kitware/CMake/blob/master/CONTRIBUTING.rs...](https://github.com/Kitware/CMake/blob/master/CONTRIBUTING.rst)

[reply](https://news.ycombinator.com/reply?id=34109919&goto=item%3Fid%3D34105770%2334109919)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108640&how=up&auth=7e9a626414b43d6fb5712c090ce178bfc51e27ae&goto=item%3Fid%3D34105770#34108640)

[mustache_kimono](https://news.ycombinator.com/user?id=mustache_kimono) [2 days ago](https://news.ycombinator.com/item?id=34108640) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34108138) | [prev](https://news.ycombinator.com/item?id=34105770#34108292) | [next](https://news.ycombinator.com/item?id=34105770#34107030) [–]

  

> The direction of the languages is entirely up to their private owners, and that exclusivity is enforced via intellectual property law.

Can't you just fork the language, or, just re Rust, participate in the RFC process? Do we need _more_ bureaucracy?

[reply](https://news.ycombinator.com/reply?id=34108640&goto=item%3Fid%3D34105770%2334108640)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107030&how=up&auth=3502905bb82dc7ffc55ae574be02b06b835cd08f&goto=item%3Fid%3D34105770#34107030)

[alcuadrado](https://news.ycombinator.com/user?id=alcuadrado) [2 days ago](https://news.ycombinator.com/item?id=34107030) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34106908) | [prev](https://news.ycombinator.com/item?id=34105770#34107021) | [next](https://news.ycombinator.com/item?id=34105770#34106626) [–]

  

yes

[reply](https://news.ycombinator.com/reply?id=34107030&goto=item%3Fid%3D34105770%2334107030)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106626&how=up&auth=68b92083f3671342e5b18ed48a71b90159ff9f0b&goto=item%3Fid%3D34105770#34106626)

[jviotti](https://news.ycombinator.com/user?id=jviotti) [2 days ago](https://news.ycombinator.com/item?id=34106626) | [parent](https://news.ycombinator.com/item?id=34105770#34106082) | [prev](https://news.ycombinator.com/item?id=34105770#34106908) | [next](https://news.ycombinator.com/item?id=34105770#34106709) [–]

  

True. The author might enjoy studying something like OpenBSD ([https://github.com/openbsd/src](https://github.com/openbsd/src))

[reply](https://news.ycombinator.com/reply?id=34106626&goto=item%3Fid%3D34105770%2334106626)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106709&how=up&auth=814bfd193e0562afe769f714a31c9b0b9cd79efb&goto=item%3Fid%3D34105770#34106709)

[notwokeno](https://news.ycombinator.com/user?id=notwokeno) [2 days ago](https://news.ycombinator.com/item?id=34106709) | [parent](https://news.ycombinator.com/item?id=34105770#34106082) | [prev](https://news.ycombinator.com/item?id=34105770#34106626) | [next](https://news.ycombinator.com/item?id=34105770#34114866) [–]

  

"GNU C" has a number of features that make the language far more pleasant if you're willing to give up the portability and use them. A big one if you've been using modern languages is anonymous functions.

[reply](https://news.ycombinator.com/reply?id=34106709&goto=item%3Fid%3D34105770%2334106709)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107243&how=up&auth=20b790fbb179d7a8b62e735e5b6b301b614939fb&goto=item%3Fid%3D34105770#34107243)

[woodruffw](https://news.ycombinator.com/user?id=woodruffw) [2 days ago](https://news.ycombinator.com/item?id=34107243) | [root](https://news.ycombinator.com/item?id=34105770#34106082) | [parent](https://news.ycombinator.com/item?id=34105770#34106709) | [next](https://news.ycombinator.com/item?id=34105770#34114866) [–]

  

It’s worth noting that GNU C uses trampolines to implement their “nested functions” extension, which in turn requires an executable stack.

In other words, using that extension silently disables one of the most basic and effective security mitigations we have for C.

[reply](https://news.ycombinator.com/reply?id=34107243&goto=item%3Fid%3D34105770%2334107243)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34114866&how=up&auth=dba2864cfc08a4d0b82c74afa471237e247ecaad&goto=item%3Fid%3D34105770#34114866)

[sumosudo](https://news.ycombinator.com/user?id=sumosudo) [1 day ago](https://news.ycombinator.com/item?id=34114866) | [prev](https://news.ycombinator.com/item?id=34105770#34106082) | [next](https://news.ycombinator.com/item?id=34105770#34106385) [–]

  

No doubt the best book on the subject of modern C

[https://www.amazon.com/21st-Century-Tips-New-School/dp/14919...](https://www.amazon.com/21st-Century-Tips-New-School/dp/1491903899)

[reply](https://news.ycombinator.com/reply?id=34114866&goto=item%3Fid%3D34105770%2334114866)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106385&how=up&auth=70f9f1d08b87df51a51a6213c9d467a163380987&goto=item%3Fid%3D34105770#34106385)

[Tarucho](https://news.ycombinator.com/user?id=Tarucho) [2 days ago](https://news.ycombinator.com/item?id=34106385) | [prev](https://news.ycombinator.com/item?id=34105770#34114866) | [next](https://news.ycombinator.com/item?id=34105770#34107153) [–]

  

The C programming language book.

[reply](https://news.ycombinator.com/reply?id=34106385&goto=item%3Fid%3D34105770%2334106385)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106775&how=up&auth=0e0934941ef091295ecc0055a3b83241adef186e&goto=item%3Fid%3D34105770#34106775)

[rerdavies](https://news.ycombinator.com/user?id=rerdavies) [2 days ago](https://news.ycombinator.com/item?id=34106775) | [parent](https://news.ycombinator.com/item?id=34105770#34106385) | [next](https://news.ycombinator.com/item?id=34105770#34107153) [–]

  

Just like C, there is a stark beauty to the original K&R C Programming Language book. And there's an admirable clarity in the writing that has set the standard for all language manuals that have followed. It is, indeed, a terrific read. As long as you're not trying to learn C.

It's really out of date. It doesn't cover subsequent C standards; and it doesn't cover any of the best practices that have subsequently evolved to protect us from the things in C that are deeply awful (in retrospect) in the original. Or any of the evolved practical lore for writing good, portable, safe C code.

[reply](https://news.ycombinator.com/reply?id=34106775&goto=item%3Fid%3D34105770%2334106775)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107784&how=up&auth=b6a33a59921e3c3d87a8284381160fd5352c469d&goto=item%3Fid%3D34105770#34107784)

[Koshkin](https://news.ycombinator.com/user?id=Koshkin) [2 days ago](https://news.ycombinator.com/item?id=34107784) | [root](https://news.ycombinator.com/item?id=34105770#34106385) | [parent](https://news.ycombinator.com/item?id=34105770#34106775) | [next](https://news.ycombinator.com/item?id=34105770#34107153) [–]

  

> _As long as you're not trying to learn C._

That's insane. K&R (2nd ed.) can easily be the _only_ book you want to work through if you want to learn the basics of C.

[reply](https://news.ycombinator.com/reply?id=34107784&goto=item%3Fid%3D34105770%2334107784)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107153&how=up&auth=9d49d3aaf3fe468f8011396264e348307ddee234&goto=item%3Fid%3D34105770#34107153)

[hzhou321](https://news.ycombinator.com/user?id=hzhou321) [2 days ago](https://news.ycombinator.com/item?id=34107153) | [prev](https://news.ycombinator.com/item?id=34105770#34106385) | [next](https://news.ycombinator.com/item?id=34105770#34113023) [–]

  

Since you are already familiar with other languages, I suggest just pick a tool that you use daily that is open source in C, and start hacking it.

[reply](https://news.ycombinator.com/reply?id=34107153&goto=item%3Fid%3D34105770%2334107153)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34113023&how=up&auth=7feb581357716b0dfeddca334eac77201801605a&goto=item%3Fid%3D34105770#34113023)

[stcroixx](https://news.ycombinator.com/user?id=stcroixx) [1 day ago](https://news.ycombinator.com/item?id=34113023) | [prev](https://news.ycombinator.com/item?id=34105770#34107153) | [next](https://news.ycombinator.com/item?id=34105770#34107528) [–]

  

K&R probably taught more working C programmers than anything besides experience.

[reply](https://news.ycombinator.com/reply?id=34113023&goto=item%3Fid%3D34105770%2334113023)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107528&how=up&auth=1891792b682ae27004e94bb55505ab6065a8ad04&goto=item%3Fid%3D34105770#34107528)

[cptnapalm](https://news.ycombinator.com/user?id=cptnapalm) [2 days ago](https://news.ycombinator.com/item?id=34107528) | [prev](https://news.ycombinator.com/item?id=34105770#34113023) | [next](https://news.ycombinator.com/item?id=34105770#34108021) [–]

  

When I was learning C, I found most books only skimmed pointers, a topic I was having a hard time with. Reek's Pointers on C is a beginner C book in that it goes over all the novice stuff, but it has a concentrated focus on getting you to understand pointers. I recommend it.

[reply](https://news.ycombinator.com/reply?id=34107528&goto=item%3Fid%3D34105770%2334107528)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108021&how=up&auth=717b78e2c3578a2335b17dbcd6295c607e8bb827&goto=item%3Fid%3D34105770#34108021)

[aWidebrant](https://news.ycombinator.com/user?id=aWidebrant) [2 days ago](https://news.ycombinator.com/item?id=34108021) | [prev](https://news.ycombinator.com/item?id=34105770#34107528) | [next](https://news.ycombinator.com/item?id=34105770#34116921) [–]

  

It seems to me that the main reason to learn C would be to be able to dive deeper into the low-level implementation details of some hardware platform.

With that in mind, I would suggest picking a piece of hardware you want to learn more about, and using C programming to see what makes it tick.

[reply](https://news.ycombinator.com/reply?id=34108021&goto=item%3Fid%3D34105770%2334108021)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34116921&how=up&auth=16d093440dbaf29a30cbfb09998b0a17dfff2c03&goto=item%3Fid%3D34105770#34116921)

[lucidguppy](https://news.ycombinator.com/user?id=lucidguppy) [1 day ago](https://news.ycombinator.com/item?id=34116921) | [prev](https://news.ycombinator.com/item?id=34105770#34108021) | [next](https://news.ycombinator.com/item?id=34105770#34107678) [–]

  

I am finding that I do well with Udemy courses first and then find the books I need and docs that will fill out the holes.

[reply](https://news.ycombinator.com/reply?id=34116921&goto=item%3Fid%3D34105770%2334116921)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107678&how=up&auth=621a8fdda29e08d735c1bfaba9ab0aa046f878a7&goto=item%3Fid%3D34105770#34107678)

[selfhifive](https://news.ycombinator.com/user?id=selfhifive) [2 days ago](https://news.ycombinator.com/item?id=34107678) | [prev](https://news.ycombinator.com/item?id=34105770#34116921) | [next](https://news.ycombinator.com/item?id=34105770#34111147) [–]

  

There's was a great article on HN on how C is more than a programming language. It is a protocol for most of computing. I would suggest you to look into the actual influence of C on computing and protocols instead of just looking at software.

[reply](https://news.ycombinator.com/reply?id=34107678&goto=item%3Fid%3D34105770%2334107678)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111147&how=up&auth=0560a4228f197547f6b949aaf3d5f117e3050524&goto=item%3Fid%3D34105770#34111147)

[mikewarot](https://news.ycombinator.com/user?id=mikewarot) [1 day ago](https://news.ycombinator.com/item?id=34111147) | [prev](https://news.ycombinator.com/item?id=34105770#34107678) | [next](https://news.ycombinator.com/item?id=34105770#34107538) [–]

  

I'm not a C programmer, but I have a program written in C that I've adopted. So I to am interested in learning modern C. Is there a text that gives the current recommended techniques for C23?

[reply](https://news.ycombinator.com/reply?id=34111147&goto=item%3Fid%3D34105770%2334111147)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107538&how=up&auth=5f6e58b687647b6906d5f0fc5c5edeb5b98e1ab2&goto=item%3Fid%3D34105770#34107538)

[hasitseth](https://news.ycombinator.com/user?id=hasitseth) [2 days ago](https://news.ycombinator.com/item?id=34107538) | [prev](https://news.ycombinator.com/item?id=34105770#34111147) | [next](https://news.ycombinator.com/item?id=34105770#34114644) [–]

  

I suggest learning C for two non-usual reasons:

1. You study the wonderful Kerninghan & Richie book. It's a great work of art.

2. C's architecture is just raw computing. Try building a toy interpreter or compiler and you will know most of CS in bargain too!

Good luck my friend.

[reply](https://news.ycombinator.com/reply?id=34107538&goto=item%3Fid%3D34105770%2334107538)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34114644&how=up&auth=7a9eba2da62c4c1e1ae9c84e75c71484aa93ef8f&goto=item%3Fid%3D34105770#34114644)

[nivertech](https://news.ycombinator.com/user?id=nivertech) [1 day ago](https://news.ycombinator.com/item?id=34114644) | [prev](https://news.ycombinator.com/item?id=34105770#34107538) | [next](https://news.ycombinator.com/item?id=34105770#34109331) [–]

  

Learn the PDP-11 assembly first, and then you will understand many design decisions made in C (especially behind pointer arithmetics).

[reply](https://news.ycombinator.com/reply?id=34114644&goto=item%3Fid%3D34105770%2334114644)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109331&how=up&auth=ce6e19c5a54476e615d0f0cda36b6e9ba87955b8&goto=item%3Fid%3D34105770#34109331)

[pjdesno](https://news.ycombinator.com/user?id=pjdesno) [1 day ago](https://news.ycombinator.com/item?id=34109331) | [prev](https://news.ycombinator.com/item?id=34105770#34114644) | [next](https://news.ycombinator.com/item?id=34105770#34110092) [–]

  

Note that [IMO] one of the main values of learning C in 2023 is that it's small and simple enough to learn the precise semantics of almost all of its parts.

[reply](https://news.ycombinator.com/reply?id=34109331&goto=item%3Fid%3D34105770%2334109331)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109516&how=up&auth=59c13d3027f6f5dc1c111e6ac81ddd31b566d260&goto=item%3Fid%3D34105770#34109516)

[pg314](https://news.ycombinator.com/user?id=pg314) [1 day ago](https://news.ycombinator.com/item?id=34109516) | [parent](https://news.ycombinator.com/item?id=34105770#34109331) | [next](https://news.ycombinator.com/item?id=34105770#34110092) [–]

  

The 2018 ISO standard [1] is 520 pages. Not what I would consider small. Nor is it simple. The integer promotion rules and undefined behaviours will trip up even experienced C programmers.

Something like R5RS for Scheme (50 pages or so) is what I think of when I hear small and simple.

[1] [https://www.iso.org/standard/74528.html](https://www.iso.org/standard/74528.html)

[reply](https://news.ycombinator.com/reply?id=34109516&goto=item%3Fid%3D34105770%2334109516)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34110092&how=up&auth=e33dc6f01356746838861b804128f5903910be4f&goto=item%3Fid%3D34105770#34110092)

[p1esk](https://news.ycombinator.com/user?id=p1esk) [1 day ago](https://news.ycombinator.com/item?id=34110092) | [prev](https://news.ycombinator.com/item?id=34105770#34109331) | [next](https://news.ycombinator.com/item?id=34105770#34109367) [–]

  

I’d probably start with learning assembly. C maps to assembly very well, and you could understand a lot of what’s happening under the hood this way.

[reply](https://news.ycombinator.com/reply?id=34110092&goto=item%3Fid%3D34105770%2334110092)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109367&how=up&auth=45b53a40f4b3458ee557b7859e599242f504ab2b&goto=item%3Fid%3D34105770#34109367)

[TrackerFF](https://news.ycombinator.com/user?id=TrackerFF) [1 day ago](https://news.ycombinator.com/item?id=34109367) | [prev](https://news.ycombinator.com/item?id=34105770#34110092) | [next](https://news.ycombinator.com/item?id=34105770#34106873) [–]

  

Great language to learn data structures and algorithms. If you still like the language itself after going through a course like that, awesome.

[reply](https://news.ycombinator.com/reply?id=34109367&goto=item%3Fid%3D34105770%2334109367)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106873&how=up&auth=a3c93fb017b99279f84f26f9592e481169a27120&goto=item%3Fid%3D34105770#34106873)

[jay754](https://news.ycombinator.com/user?id=jay754) [2 days ago](https://news.ycombinator.com/item?id=34106873) | [prev](https://news.ycombinator.com/item?id=34105770#34109367) | [next](https://news.ycombinator.com/item?id=34105770#34107960) [–]

  

Does anyone have a good recommendation on developing operating systems?

[reply](https://news.ycombinator.com/reply?id=34106873&goto=item%3Fid%3D34105770%2334106873)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106947&how=up&auth=3a06fffc265e51b6f403cc3d7592b2d1596ec37b&goto=item%3Fid%3D34105770#34106947)

[BirAdam](https://news.ycombinator.com/user?id=BirAdam) [2 days ago](https://news.ycombinator.com/item?id=34106947) | [parent](https://news.ycombinator.com/item?id=34105770#34106873) | [next](https://news.ycombinator.com/item?id=34105770#34107960) [–]

  

This depends on how deep you want to go. I honestly recommend starting small.

First, Code by Petzold to make sure you understand how a computer truly works. Then I recommend:

Understanding Microprocessors (Motorola)

8086/8088 Assembly Language Programming

Dissecting DOS

Practical Filesystem Design

FreeBSD Internals

The reason there are so many books is because it’s a big topic. The reason most are old is that things were once much simpler. Modern stuff builds on old, so if you get a good understanding of older stuff, you have an idea of what is happening. From there, I strongly recommend getting a good understanding of UEFI. UEFI will handle most of your system initialization no boot requirements, allowing you to focus on higher level stuff. Prior knowledge of systems will still be a requirement so that you can understand some of the assumptions and choices that UEFI makes.

OSDev is fun. I hope you have a good time.

[reply](https://news.ycombinator.com/reply?id=34106947&goto=item%3Fid%3D34105770%2334106947)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108375&how=up&auth=0f68453468a7f9d0354c44cfde2d421e3afe18d4&goto=item%3Fid%3D34105770#34108375)

[rramadass](https://news.ycombinator.com/user?id=rramadass) [2 days ago](https://news.ycombinator.com/item?id=34108375) | [root](https://news.ycombinator.com/item?id=34105770#34106873) | [parent](https://news.ycombinator.com/item?id=34105770#34106947) | [next](https://news.ycombinator.com/item?id=34105770#34107960) [–]

  

Great recommendations; a couple i had not come across.

>Understanding Microprocessors (Motorola)

>Dissecting DOS

Can you share or point me to where i can get the above from?

[reply](https://news.ycombinator.com/reply?id=34108375&goto=item%3Fid%3D34105770%2334108375)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34111589&how=up&auth=60939a7126bac7ea3ebf2447c9f5821ad8f11974&goto=item%3Fid%3D34105770#34111589)

[BirAdam](https://news.ycombinator.com/user?id=BirAdam) [1 day ago](https://news.ycombinator.com/item?id=34111589) | [root](https://news.ycombinator.com/item?id=34105770#34106873) | [parent](https://news.ycombinator.com/item?id=34105770#34108375) | [next](https://news.ycombinator.com/item?id=34105770#34107960) [–]

  

For Dissecting DOS, I recommend eBay, but it is available on Amazon:

[https://www.amazon.com/Dissecting-DOS-Code-Level-Operating-S...](https://www.amazon.com/Dissecting-DOS-Code-Level-Operating-System/dp/020162687X/ref=sr_1_1?crid=1BGSDZ592N8NC&keywords=dissecting+dos&qid=1671836287&s=books&sprefix=dissecting+dos%2Cstripbooks%2C93&sr=1-1)

The original companion floppy for that is available on archive.org:

[https://archive.org/details/dissecting-dos-companion-floppy-...](https://archive.org/details/dissecting-dos-companion-floppy-disk)

Understanding Microprocessors is a small book, and currently unavailable anywhere I've seen except for eBay, but it's also normally quite inexpensive.

[reply](https://news.ycombinator.com/reply?id=34111589&goto=item%3Fid%3D34105770%2334111589)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107960&how=up&auth=bb4c69126b28f04acc3fbf42a03f6400595d06af&goto=item%3Fid%3D34105770#34107960)

[theusus](https://news.ycombinator.com/user?id=theusus) [2 days ago](https://news.ycombinator.com/item?id=34107960) | [prev](https://news.ycombinator.com/item?id=34105770#34106873) | [next](https://news.ycombinator.com/item?id=34105770#34105944) [–]

  

The C Programming Language by Brian K is still the best resource.

[reply](https://news.ycombinator.com/reply?id=34107960&goto=item%3Fid%3D34105770%2334107960)

![](https://news.ycombinator.com/s.gif)

![](https://news.ycombinator.com/s.gif)

[pettycashstash2](https://news.ycombinator.com/user?id=pettycashstash2) [2 days ago](https://news.ycombinator.com/item?id=34105944) | [prev](https://news.ycombinator.com/item?id=34105770#34107960) | [next](https://news.ycombinator.com/item?id=34105770#34106556) [3 more]

  

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106556&how=up&auth=5ce1a28ec23d6ac3005774941ab8728606e3fb80&goto=item%3Fid%3D34105770#34106556)

[astrodust](https://news.ycombinator.com/user?id=astrodust) [2 days ago](https://news.ycombinator.com/item?id=34106556) | [prev](https://news.ycombinator.com/item?id=34105770#34105944) | [next](https://news.ycombinator.com/item?id=34105770#34107115) [–]

  

If you've learned Rust then you've learned a lot about the importance of ownership and lifetimes, something C will never teach you directly. It will only punish you for failing to understand.

[reply](https://news.ycombinator.com/reply?id=34106556&goto=item%3Fid%3D34105770%2334106556)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107115&how=up&auth=f2989f4c6d7fd474408aa2a55803e2460288d448&goto=item%3Fid%3D34105770#34107115)

[slingnow](https://news.ycombinator.com/user?id=slingnow) [2 days ago](https://news.ycombinator.com/item?id=34107115) | [prev](https://news.ycombinator.com/item?id=34105770#34106556) | [next](https://news.ycombinator.com/item?id=34105770#34106570) [–]

  

How in the world in this day and age of nearly infinite resources to learn ANYTHING on the internet is this a question that gets voted up to the top of HN?

Here's an idea: pick a project that interests you and start. For any question or issue you run into, there will be 1000 Youtube videos to choose from, and 10000 online tutorials. There will be an associated sub reddit to ask questions.

[reply](https://news.ycombinator.com/reply?id=34107115&goto=item%3Fid%3D34105770%2334107115)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107577&how=up&auth=30d26b3d7f300abf0edc37e93cf3c4342273ce91&goto=item%3Fid%3D34105770#34107577)

[tzs](https://news.ycombinator.com/user?id=tzs) [2 days ago](https://news.ycombinator.com/item?id=34107577) | [parent](https://news.ycombinator.com/item?id=34105770#34107115) | [next](https://news.ycombinator.com/item?id=34105770#34106570) [–]

  

> How in the world in this day and age of nearly infinite resources to learn ANYTHING on the internet is this a question that gets voted up to the top of HN?

I'd guess it is because Sturgeon's Law[1] applies to most things on the internet. In fact there is so much crap when it specifically comes to instructional material on the internet that it reaches a point where we find that Sturgeon's Law was only a linear approximation to crappiness that is only valid with the relatively limited crap levels they were capable of producing 60 years ago when Sturgeon's Law was formulated.

The internet is capable of achieving crap harmonics, and we find that something like 90% of the stuff that doesn't seem to be crap is actually crap too.

The HN answers to this question are going to have a much lower crap level which will keep them well within the linear region of the Law. Many here recognize that having someplace that answers this question with only a 90% crap level is good and so upvote it.

[1] 90% of everything is crap

[reply](https://news.ycombinator.com/reply?id=34107577&goto=item%3Fid%3D34105770%2334107577)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106570&how=up&auth=02c2d8645b064aea72daf8f9163ba3e0bebdc247&goto=item%3Fid%3D34105770#34106570)

[svnt](https://news.ycombinator.com/user?id=svnt) [2 days ago](https://news.ycombinator.com/item?id=34106570) | [prev](https://news.ycombinator.com/item?id=34105770#34107115) [–]

  

I like firearms as an analogy here.

Rust is like a modern rifle, unloaded, with six safeties and three locked triggers in a gun safe. The ammunition will harmlessly self-destruct if you don’t focus properly before firing.

C is sort of like Rust except the rifle is on the couch, loaded, one in the chamber, all the safeties are off, the numerous open triggers are held by RNGs, and the business end is pointed right at your most prized possession.

It is fun, in a way. Use it in an environment that you don’t care too much about.

[reply](https://news.ycombinator.com/reply?id=34106570&goto=item%3Fid%3D34105770%2334106570)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106788&how=up&auth=8fb4f35c6c5eb5909ff40d57c53a1e789a1577db&goto=item%3Fid%3D34105770#34106788)

[notwokeno](https://news.ycombinator.com/user?id=notwokeno) [2 days ago](https://news.ycombinator.com/item?id=34106788) | [parent](https://news.ycombinator.com/item?id=34105770#34106570) | [next](https://news.ycombinator.com/item?id=34105770#34107643) [–]

  

Meh. People love to hate on C but the invariants you have to keep track of are relatively simple (things like 0<=i<L.) Most languages replace these with multi-dimensional string based hashmaps that have much more complex invariants to keep track of and similar consequences for failing to do so.

Some languages like Rust and Haskell do a little better by having the compiler keep track of most of them (if you use it right) but the consequence of misunderstanding them here is accidentally coding yourself into a corner which is extremely unpleasant.

There are other things people like to talk about like automatic memory management and package management and again I'm not convinced these really improve much. The pain of allocating memory in C means people often find ways to avoid doing it altogether which results in much more predictable behavior. Same with pulling in third party libraries. It's rare using more than a couple is actually beneficial but when the package manager is there the temptation is way too much for most people and the result is both unpleasant and very vulnerable to supply chain attacks.

I'm still not convinced any of this is the improvement people think it is.

[reply](https://news.ycombinator.com/reply?id=34106788&goto=item%3Fid%3D34105770%2334106788)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106931&how=up&auth=67a171b7f8bd4f54ef8252db8c2ef7e7ba0e4824&goto=item%3Fid%3D34105770#34106931)

[svnt](https://news.ycombinator.com/user?id=svnt) [2 days ago](https://news.ycombinator.com/item?id=34106931) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34106788) | [next](https://news.ycombinator.com/item?id=34105770#34106928) [–]

  

I mean I’ve worked with C for a long time. It’s love-hate. I think it’s hard to see the value from the other side of having internalized C, and Rust definitely encumbers you.

The number of things a beginner doesn’t have to know or worry about to code something as essential as a function that reads a file, searches for a string, and interacts in a serious way with the network in Rust vs C is substantial.

Probably I’ll get downvoted again but it’s fine.

[reply](https://news.ycombinator.com/reply?id=34106931&goto=item%3Fid%3D34105770%2334106931)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107184&how=up&auth=649b424c20d95442ffa5461d7e6226294a667e96&goto=item%3Fid%3D34105770#34107184)

[mturmon](https://news.ycombinator.com/user?id=mturmon) [2 days ago](https://news.ycombinator.com/item?id=34107184) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34106931) | [next](https://news.ycombinator.com/item?id=34105770#34106928) [–]

  

But you do learn from such things (opening a file, reading from the network). All the stuff that can go wrong, and that it all has to be handled. Like writing a device driver.

It’s useful to write that C code, and then try to break it (e.g., to force a segfault rather than a proper error-then-exit).

Is knowing all that necessary? No, as you are kind of implying. ;-)

[reply](https://news.ycombinator.com/reply?id=34107184&goto=item%3Fid%3D34105770%2334107184)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34106928&how=up&auth=b2b5f22001c873df9504bf58a8cec333805b1a61&goto=item%3Fid%3D34105770#34106928)

[quotemstr](https://news.ycombinator.com/user?id=quotemstr) [2 days ago](https://news.ycombinator.com/item?id=34106928) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34106788) | [prev](https://news.ycombinator.com/item?id=34105770#34106931) | [next](https://news.ycombinator.com/item?id=34105770#34107643) [–]

  

> the invariants you have to keep track of are relatively simple (things like 0<=i<L.)

If maintaining these invariants were so simple, why would the smartest people in the industry have so much trouble doing so? Humans cannot consistently write correct C, not even the ones who believe they can. We have decades of evidence.

> Most languages replace these with multi-dimensional string based hashmaps that have much more complex invariants to keep track of and similar consequences for failing to do so.

This take reflects a profound ignorance of the relative severity of different kinds of software defect.

When you get an array bound wrong in Java or Rust, you get an exception or a panic. Your program exits and restarts. You get a ticket for fixing the bug. Life goes on.

When you get an array bound wrong in C, you open the door for a nefarious actor to gain control over everything that your program can access.

These aren't similar consequences. They're not even on the same planet of consequences. C apologists are like apologists for asbestos and lead paint.

[reply](https://news.ycombinator.com/reply?id=34106928&goto=item%3Fid%3D34105770%2334106928)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107016&how=up&auth=29e0a1343db51e9f68b10574054e3370d4ef5469&goto=item%3Fid%3D34105770#34107016)

[notwokeno](https://news.ycombinator.com/user?id=notwokeno) [2 days ago](https://news.ycombinator.com/item?id=34107016) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34106928) | [next](https://news.ycombinator.com/item?id=34105770#34119898) [–]

  

My point is that they can't consistently write correct software in any language. Javascript for example has similar issues where if you're not careful about which keys in a hashmap you write to you can create situations that allow for all kinds of things including arbitrary code execution.

Modern C compilers/runtimes also have bounds checking (although I've never used it) like java if you want to use it.

Sure there's some improvement but it's not nearly as extreme as it's made out to be.

[reply](https://news.ycombinator.com/reply?id=34107016&goto=item%3Fid%3D34105770%2334107016)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107044&how=up&auth=c5bb38349827515928c505b722b98cf6406499c7&goto=item%3Fid%3D34105770#34107044)

[quotemstr](https://news.ycombinator.com/user?id=quotemstr) [2 days ago](https://news.ycombinator.com/item?id=34107044) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34107016) | [next](https://news.ycombinator.com/item?id=34105770#34107113) [–]

  

It's true that humans will inevitably produce software with defects. There are differences in the impacts of these defects.

> Javascript for example has similar issues where if you're not careful about which keys in a hashmap you write to you can create situations that allow for all kinds of things including arbitrary code execution.

Not without eval() you can't.

> Modern C compilers/runtimes also have bounds checking

No, they don't. There are C-like languages (e.g. C++/CLI) that do, and you can use containers, but no, plain C cannot do bounds checking. There's no room in the ABI for it, at least not without CHERI or ASAN.

> although I've never used it

One of the most dangerous things in the world is giving powerful tools to people who don't know that they shouldn't use them.

[reply](https://news.ycombinator.com/reply?id=34107044&goto=item%3Fid%3D34105770%2334107044)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107113&how=up&auth=faab2348d48a050f22a7d4a2d14e032fb0e6e71b&goto=item%3Fid%3D34105770#34107113)

[svnt](https://news.ycombinator.com/user?id=svnt) [2 days ago](https://news.ycombinator.com/item?id=34107113) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34107016) | [prev](https://news.ycombinator.com/item?id=34105770#34107044) | [next](https://news.ycombinator.com/item?id=34105770#34119898) [–]

  

> My point is that they can’t consistently write correct software in any language. Javascript for example

I don’t think Javascript is the counterexample you want to use as a good language. It is like C where safety is traded for flexibility.

If people can’t write correct software in any language, the language designers can help, which is exactly what the Rust designers did.

It’s not great for everyone because now all the scar tissue that great C devs have around what near-fatal errors they’ve made previously has lost some of its value.

[reply](https://news.ycombinator.com/reply?id=34107113&goto=item%3Fid%3D34105770%2334107113)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34119898&how=up&auth=297e1b580e6d2c4d7c0efc6636f1a5e855758ea1&goto=item%3Fid%3D34105770#34119898)

[lelanthran](https://news.ycombinator.com/user?id=lelanthran) [23 hours ago](https://news.ycombinator.com/item?id=34119898) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34106928) | [prev](https://news.ycombinator.com/item?id=34105770#34107016) | [next](https://news.ycombinator.com/item?id=34105770#34107643) [–]

  

> Humans cannot consistently write correct C,

So? Humans cannot consistently write correct code in any language.

> When you get an array bound wrong in C, you open the door for a nefarious actor to gain control over everything that your program can access.

What's really interesting to me is how Heartbleed (the largest example of getting an array bound wrong) was orders of magnitude less expensive to fix than Log4j, and yet it was Log4j that resulted in gauranteed RCE, not Heartbleed.

> C apologists are like apologists for asbestos and lead paint.

Let's not go spilling hyperbole all down our shirt.

I get the feeling that a lot of people comparing C to things that kill have never worked in life-critical situations, the systems of which ran C for decades and decades, and never appeared to have lost very many lives in the process.

Military, munitions, healthcare machines ... I can't remember any of them losing lives because of a misfeature in the C programming language, and yet there you are, proudly proclaiming to the word that C is like Asbestos, which _has_ taken peoples lives.

Get a grip.

[reply](https://news.ycombinator.com/reply?id=34119898&goto=item%3Fid%3D34105770%2334119898)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34107643&how=up&auth=dd1abadddc839bb3fe811541c1f126d7f164afe6&goto=item%3Fid%3D34105770#34107643)

[weakfortress](https://news.ycombinator.com/user?id=weakfortress) [2 days ago](https://news.ycombinator.com/item?id=34107643) | [parent](https://news.ycombinator.com/item?id=34105770#34106570) | [prev](https://news.ycombinator.com/item?id=34105770#34106788) | [next](https://news.ycombinator.com/item?id=34105770#34108614) [–]

  

This analogy is really more like dogmatic Rustacean-ism than actual truth. Rust has it's merits. However, it's a tired language. It doesn't maintain very good stability due to it's ever-evolving compiler which makes it almost unusable in anything but a "move fast and break things" environment. It's too particular, making development a burden and interrupting flow. Most importantly, the community is absolutely exhausting to listen to and this post is a perfect example.

C can be dangerous, yes. No one is denying that. C is also extremely stable, fast, and a "known-known". With proper discipline that any experienced developer of ANY language will have C can be very safe. Most, if not all, dangerous errors can be found and fixed simply by understanding how to use valgrind. With the advent of modern CI/CD it's never been easier to remember to valgrind your code.

C is probably the first language I would use in an environment I cared a lot about. For example, if binary size, availability of libraries, competency of programmer, etc would be in my list of concerns I would use C. I take particular offense to the "RNG" part of your analogy. C errors are not random at all. They are _always_ the developer's fault and are often very predictable (use-after-free, poor array hygiene, etc). Despite the ever-evolving condition of post-C religions (Go, Rust, etc) there are only a handful of languages that fit the bill for a so-called serious systems project. C, C++, Java, and C#. You simply will not find a better library ecosystem and community of highly tenured developers to draw from for at least another decade in other languages.

[reply](https://news.ycombinator.com/reply?id=34107643&goto=item%3Fid%3D34105770%2334107643)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109540&how=up&auth=fc06a65c24afc352646fa5e897c77a9c400699f1&goto=item%3Fid%3D34105770#34109540)

[svnt](https://news.ycombinator.com/user?id=svnt) [1 day ago](https://news.ycombinator.com/item?id=34109540) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34107643) | [next](https://news.ycombinator.com/item?id=34105770#34108614) [–]

  

I don’t even prefer Rust to C personally, and I’m not a dogmatic sort. My analogy wasn’t particularly complimentary to either language.

If some people are too sensitive to take a little caricaturization of C in stride, maybe those people should examine those sensitivities.

Edit because I didn’t read your comment all the way through: uninitialized memory is RNG, functionally.

[reply](https://news.ycombinator.com/reply?id=34109540&goto=item%3Fid%3D34105770%2334109540)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34108614&how=up&auth=7f5cfb726c4cfa9fb9aac0bbb9d554f918f00b77&goto=item%3Fid%3D34105770#34108614)

[julianeon](https://news.ycombinator.com/user?id=julianeon) [2 days ago](https://news.ycombinator.com/item?id=34108614) | [parent](https://news.ycombinator.com/item?id=34105770#34106570) | [prev](https://news.ycombinator.com/item?id=34105770#34107643) [–]

  

Ha.

Confession time: I came to this thread to Ctrl-F for "Rust" to find the Rust-C fisticuffs. And this answer is like what I would expect if I said "ChatGPT, write me a reddit-style Rust defense that would rustle some jimmies."

I'm about equally capable at both Rust and C - I could pass a 1st year undergrad CS course using either, but I'm no better than that.

I know C has its place - its marketshare of embedded devices and gamedev is huuuge - and Rust... may eventually reach a similar place of importance, but isn't there yet.

I also know the very best C and Rust devs are pretty much geniuses - Ph.D. level understanding of math, that kind of thing - and would squash me in any type of fair competition.

Maybe in the year 2050, an attitude of condescending acceptance towards C might be appropriate. And also... maybe not, because C might not slip. Who knows?

And really, beyond this pretty crude but accurate-ish description, I think we're getting into unnecessary hair-splitting, among people who don't work with either and wouldn't understand the tradeoffs anyway.

[reply](https://news.ycombinator.com/reply?id=34108614&goto=item%3Fid%3D34105770%2334108614)

![](https://news.ycombinator.com/s.gif)

[

](https://news.ycombinator.com/vote?id=34109601&how=up&auth=e18220f46defba367b9e5f20d9cd471b725217da&goto=item%3Fid%3D34105770#34109601)

[svnt](https://news.ycombinator.com/user?id=svnt) [1 day ago](https://news.ycombinator.com/item?id=34109601) | [root](https://news.ycombinator.com/item?id=34105770#34106570) | [parent](https://news.ycombinator.com/item?id=34105770#34108614) [–]

  

There is a lot of nuance in this discussion that someone inexperienced could misinterpret, you’re definitely right about that.

[reply](https://news.ycombinator.com/reply?id=34109601&goto=item%3Fid%3D34105770%2334109601)
