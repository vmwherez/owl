a.out ELF was not always the standard; original UNIX systems used a file format called `a.out`. We can see the vestiges of this if you compile a program without the -o option to specify an output file name; the executable will be created with a default name of `a.out`. 

`a.out` is a very simple header format that only allows a single data, code and BSS section. As you will come to see, this is insufficient for modern systems with [[dynamic libraries]].


