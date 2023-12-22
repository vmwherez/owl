[obsidian://open?vault=C&file=pdf%2Fbgc_usl_c_1.pdf](obsidian://open?vault=C&file=pdf%2Fbgc_usl_c_1.pdf)



#### C++
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