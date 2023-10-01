#include <stdio.h>

/* int fibonacci(int n)
{
    if (n <= 1)
    {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

*/

int pointers()
{
    int num = 10; // Declare an integer variable
    int *ptr;     // Declare a pointer variable

    ptr = &num; // Assign the address of 'num' to the pointer

//    printf("Value of num: %d\n", num);
//    printf("Address of num: %p\n", &num);
//    printf("Value of ptr: %p\n", ptr);
//    printf("Value pointed by ptr: %d\n", *ptr);

    *ptr = 20; // Change the value of 'num' indirectly through the pointer

    printf("New value of num: %d\n", num);

    return 0;
}

int main()
{
    int n = 10; // Change this value to calculate a different number of Fibonacci sequence terms
/*$ r2 /bin/ls   # open file in read-only
 * > aaa          # analyse the program (r2 -A)
 * > afl          # list all functions (try aflt, aflm)
 * > px 32        # print 32 byte hexdump current block
 * > s sym.main   # seek to main (using flag name)
 * > f~foo        # filter flags matching 'foo' (internal |grep)
 * > iS;is        # list sections and symbols (rabin2 -Ss)
 * > pdf; agf     # disassembly and ascii-art function graph
 * > oo+;w hello  # reopen in read-write and write a string
 * > ?*~...       # interactive filter in all command help
 * > q            # quit
 *
    printf("Fibonacci Sequence up to %d terms:\n", n);
    for (int i = 0; i < n; i++)
    {
        printf("%d ", fibonacci(i));
    }
 
 */
    pointers();

    return 0;
}
