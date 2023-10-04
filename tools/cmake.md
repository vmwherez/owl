## cmake

https://riptutorial.com/cmake/example/7501/simple--hello-world--project

##### `main.cpp`

```c++
// main.cpp (C++ Hello World Example)

#include <iostream> // io std lib
// #include <,,,> 

int main()
{
    std::cout << "Hello World!\n";
    return 0;
}


```

##### `CMakeLists.txt`

```
cmake_minimum_required(VERSION 2.4)

project(hello_world)

add_executable(app main.cpp)
```

1. [`cmake_minimum_required(VERSION 2.4)`](https://cmake.org/cmake/help/latest/command/cmake_minimum_required.html) sets a minimum CMake version required to evaluate the current script.
2. [`project(hello_world)`](https://cmake.org/cmake/help/latest/command/project.html) starts a new CMake project. This will trigger a lot of internal CMake logic, especially the detection of the default C and C++ compiler.
3. With [`add_executable(app main.cpp)`](https://cmake.org/cmake/help/latest/command/add_executable.html) a build target `app` is created, which will invoke the configured compiler with some default flags for the current setting to compile an executable `app` from the given source file `main.cpp`.

**Command Line** *(In-Source-Build, not recommended)*

```
> cmake .
...
> cmake --build .
```

[`cmake .`](https://cmake.org/cmake/help/latest/manual/cmake.1.html) does the compiler detection, evaluates the `CMakeLists.txt` in the given `.` directory and generates the build environment in the current working directory.

The `cmake --build .` command is an abstraction for the necessary build/make call.

**Command Line** *(Out-of-Source, recommended)*

To keep your source code clean from any build artifacts you should do "out-of-source" builds.

```
> mkdir build
> cd build
> cmake ..
> cmake --build .
```

Or CMake can also abstract your platforms shell's basic commands from above's example:

```
> cmake -E make_directory build
> cmake -E chdir build cmake .. 
> cmake --build build 
```


---


- https://iamsorush.com/posts/cpp-cmake-essential/
- https://iamsorush.com/posts/cpp-cmake-build/
- https://marketplace.visualstudio.com/items?itemName=ms-vscode.cmake-tools
- https://code.visualstudio.com/docs/cpp/cmake-linux