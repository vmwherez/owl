package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Hello, world!")
}

// This is a comment

// This is another comment

// Cody, write a simple HTTP server in golang:

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", nil)
}

// Cody, write a simple TCP server in golang

func tcp_server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handle_conn(conn)
	}
}

func handle_conn(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
		conn.Write([]byte("Hello, world!\n"))
	}
}

// Cody, write a simple UDP server in golang:

func udp_server() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
		conn.WriteToUDP([]byte("Hello, world!\n"), addr)
	}
}

// Cody, write a simple UDP client in golang:

func udp_client() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("Hello, world!\n"))
	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}

// Cody, write a simple websocket server in golang

func websocket_server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer ws.Close()
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}
			log.Printf("recv: %s", message)
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println(err)
				break
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Cody, write a simple websocket client in golang

func websocket_client() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/",
	}
	log.Printf("connecting to %s", u.String())
	err := websocket.Dial(u.String(), "", "http://localhost/")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

// Cody, write a simple TCP client in golang

func tcp_client() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("Hello, world!\n"))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}

// Cody, write a simple HTTP client in golang

func http_client() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}


// Cody, what are the data types in golang?
/*
string "" // string
int // int
float64 // float64
bool // bool
[]byte // []byte
map[string]int // map[string]int
func() // func()

// Cody, what are the data structures in golang?

struct // struct
array // array
slice // slice

// Cody, what are the operators in golang?

+ - * / %
& | ^
<< >>
==!= < > <= >=
&& ||
= += -= *= /= %=
&= |= ^= <<= >>=

// Cody, what are the keywords in golang?

break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var

// Cody, what are the built-in functions in golang?

make len cap new append copy close delete complex imag len panic print println real recover

// Cody, what are the built-in types in golang?

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64
complex64 complex128
string
bool

// Cody, what are the built-in constants in golang?

true false iota nil

// Cody, what are the built-in packages in golang?

fmt io net http os reflect runtime strconv strings sync time

// Cody, what are the built-in interfaces in golang?

io.Reader io.Writer io.Closer io.ReaderAt io.WriterAt io.Seeker io.ByteReader io.ByteScanner io

// Cody, what are the built-in errors in golang?

error os.PathError os.LinkError os.SyscallError os.Error os.ErrorString

*/

// Cody, how to define a struct in golang?

type Person struct {
	Name string
	Age  int
}

// Cody, how to define a function in golang?

func say_hello(name string) {
	fmt.Println("Hello, " + name)
}

// Cody, how to define a method in golang?

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello, " + p.Name)
}

// Cody, how to define a function with multiple return values in golang?

func get_info() (string, int) {
	return "Cody", 25
}

// Cody, how do I define a function with multiple return values and a named return value in golang?

func get_info() (name string, age int) {
	return "Cody", 25
}

// Cody, how do I add routes to my server in golang?

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, world!")
// })

// Cody, how do I read a file in golang?

//  data, err := ioutil.ReadFile("file.txt")

// Cody, how do I write a file in golang?

//  ioutil.WriteFile("file.txt", []byte("Hello, world!"), 0644)

// Cody, how do I create a directory in golang?

//  os.Mkdir("dir", 0755)

// Cody, how do I remove a directory in golang?

//  os.RemoveAll("dir")

// Cody, how do I create a file in golang?

//  os.Create("file.txt")

// Cody, how do I remove a file in golang?

//  os.Remove("file.txt")

// Cody, how do I rename a file in golang?

//  os.Rename("file.txt", "file2.txt")

// Cody, how do I copy a file in golang?

//  os.Copy("file.txt", "file2.txt")

// Cody, how do I get the current working directory in golang?

//  os.Getwd()

// Cody, how do I get the current user in golang?

//  os.Getuid()

// Cody, how do I get the current group in golang?

//  os.Getgid()

// Cody, how do I get the current process id in golang?

//  os.Getpid()

// Cody, how do I get the current process name in golang?

//  os.Args

// Cody, how do I get the current process environment in golang?

//  os.Environ()

// Cody, how do I get the current process executable in golang?

//  os.Executable()

// Cody, how do I get the current process executable directory in golang?

//  os.Executable()

// Cody, how I use pointers in Golang:


// Declaration: To declare a pointer, you use the var keyword followed by the variable name and the type it points to.
var ptr *int // Declares a pointer to an integer

// Initialization: You can initialize a pointer by using 
// the address-of operator (&) followed by a variable of the same type.
x := 42
ptr = &x // Initializes ptr to point to the memory address of x

// Dereferencing: To access the value a pointer points to, you use the dereference operator (*).

fmt.Println(*ptr) // Prints the value of x (which is 42)

/*
Pointer Arithmetic: Unlike some languages, Go does not support pointer arithmetic. 
You can't perform arithmetic operations directly on pointers.
Passing by Reference: Go passes arguments to functions by value, 
but you can achieve pass-by-reference behavior by passing pointers.
*/

func modifyValue(ptr *int) {
    *ptr = 100
}
modifyValue(&x)
fmt.Println(x) // x is now 100

/*
Nil Pointers: If a pointer does not point to anything (not initialized), 
it holds a special value called nil.
*/

var ptr *int
if ptr == nil {
    fmt.Println("Pointer is nil")
}

// New Function: You can use the new function 
// to allocate memory for a new value 
// and return a pointer to it.

ptr := new(int) // Allocates memory for an integer and returns a pointer
*ptr = 10

/*
Returning Pointers: 
Be cautious when returning pointers from functions. 
It's generally a good practice to return pointers 
to values that have a longer lifetime 
than the function itself.
*/

func createPointer() *int {
    x := 42
    return &x // This can be risky, as x will go out of scope when the function exits
}

/*
Remember to handle pointers with care, 
as improper use can lead to runtime errors 
like null pointer dereferences 
or memory leaks.
*/


