// The bottom line about migrating to Golang is it needs to be done 
// if you’re anticipating a surge in service demands 
//that will crash existing infrastructure capacity


import (
	"fmt"
	"strings"
	"io"
	"os"
	"time"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"sync"
)


/**************************  Go's basic operation *********************************/

/* 
go mod init <mod-name>: create a go.mod file that identifies the code as a module
go mod edit -replace=example.com/greetings=../greetings
go mod tidy: Add new module requirements and sums
go add: 
go run test.go

module collects one or more related packages for a discrete and useful set of functions
a function whose name starts with a capital letter can be called by a function not in the same package


compile: go build wiki.go
run: ./wiki.go
 */



func swap(x, y string) (string, string) {
	return y, x
}

// x, y share the type int

var i, j int = 1, 2

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. k := 3



/**************************  Go's basic types *********************************/

/* bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128 */


/**************************  Go's basic types *********************************/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/* default values of variables:
0 for numeric
false for boolean type
"" for strings
 */

var i = 42
i := 42
const Pi = 3.14


/**************************  Loop construct *********************************/

// go has only one looping construct for

for i := 0; i < 10; i++ {
	sum += i
}

for sum < 1000 {
	sum += sum
}

for {
}

/**************************  switch *********************************/

switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.\n", os)
}
		
		
today := time.Now().Weekday()
switch time.Saturday {
case today + 0:
	fmt.Println("Today.")
case today + 1:
	fmt.Println("Tomorrow.")
case today + 2:
	fmt.Println("In two days.")
default:
	fmt.Println("Too far away.")
}
	
	
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}

/**************************  pointer *********************************/
// var p *int

i, j := 42, 2701

p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j

/**************************  struct *********************************/

// struct is a collection of field
type Vertex struct {
	X int
	Y int
}

v := Vertex{1, 2}
v.X = 4
// allow using p.X instead of (*p).X
p := &v
p.X = 4

v1 = Vertex{1, 2}  // has type Vertex
v2 = Vertex{X: 1}  // Y:0 is implicit
v3 = Vertex{}      // X:0 and Y:0
p  = &Vertex{1, 2} // has type *Vertex


/**************************  Arrays *********************************/
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Printf(a)

primes := [6]int{2, 3, 5, 7, 11, 13}

// slice: Slices are like references to arrays
var s []int = primes[1:4]  //include 1, exclude 4

// slice literal is like an array literal without the length
q := []int{2, 3, 5, 7, 11}
r := []bool{true, false, true, true, false}
s := []struct {
	i int
	b bool
} {
	{2, true},
	{3, false},
}

// slice defaults
primes[:10]
primes[0:]
primes[:]

// slice length and capacity
s := primes[2:]
len(s) // slice length 
cap(s) // slice capacity

// create a slice with make

// allocate a zeroed array and returns a slice that refers to that array
a := make([]int, 5) // len(a)=5, the defualt capacity equals the specified length
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

// slices can contain any type, including other slices
board := [][]sting {
	[]string{"-", "-", "-"},
	[]string{"-", "-", "-"},
	[]string{"-", "-", "-"},
}

// append to a slice

var s []int
s = append(s, 0)
s = append(s, 1, 2, 3)

// range, the first is the index, the second is a copy of element
var pow =[]int{1, 2, 4, 8, 16, 32, 64, 128}

for i, v := range pow {
}

for i, _ := range pow {
}

for _, v := range pow {
}

for i := range pow {

}


// implement the pic using slices
func Pic(dx, dy int) [][]uint8 {
	sl := make([][]uint8, dy)
	for i := range sl {
		sl[i] = make([]uint8, dx)
		for j :=0; j < dx; j++ {
			sl[i][j] = uint8((i+j)/2 + i ^ j)
		}
	}
	return sl
	
}

pic.Show(Pic)

/********************************** maps **************************************/
type Vertex struct {
	Lat, Long float64	
}

var m map[string]Vertex

// make function returns a map of the given type, initialized and ready for use
m = make(map[string]Vertex)
m["Bell Labs"] = Vertex {40.68433, -74.39967}


// Map literals
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40, -74,
	},
	"Google": Vertex{
		37, -122,
	},
}

// the literals can be simplified
var m = map[string]Vertex{
	"Bell Labs": {40, -74},
	"Google": {37, -122},
}

delete(m, "Bell Labs")

// if key in m, ok is true, if not, ok is false
v, ok := m["Bell Labs"]

// count the words using map
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := string.Split(s, " ")
	for _, v := range words {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	return m
}


/***************** Function values ***********************************/


// functions are passed as values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

compute(math.Pow)

/***************************** function closures *************************/
// closure is a function value that references variables from outside its body

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

pos, neg := adder(), adder()
for i := 0; i < 10; i++ {
	fmt.Println(
		pos(i),
		neg(-2*i),
	)
}

func fibonacci() func() int {
	prev := 0
	curr := 1
	return func() int {
		prev, curr = curr, prev + curr
		return curr
	}
}

/******************* method ******************************************/

// A method is a function with a special receiver argument

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

v := Vertex{3, 4}
fmt.Println(v.abs())

// you can only declare a method with a receiver whose type is defined in the same package

type MyFloat float64

func (f MyFloat) Abs() float {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

f := MyFloat(-math.Sqrt2)
f.Abs()

// pointer receivers are more common than value receivers, because they can modify the value to which the receiver points

func (v *Vertex) Scale(f float) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/********************** interfaces **************************/
// An interface type is defined as a set of method signatures

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// in Go we need to handle the nil gracefully

// I is the name of the interface, M is the name of the function
// An interface value holds a value of a specific underlying concrete type. 
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

var i I

var t *T
i = t // interface value is nil
i.M()


i = &T{"hello"} // interface value is type struct
i.M()

// type assertion 

var i interface{} = "hello"

s := i.(string) // assert that i holds the concrete type string, return "hello"
s, ok := i.(string) // return hello, true

s = i.(float64) // trigger a panic
s, ok := i.(float64) // return 0, false

// type switches

switch v := i.(type) {
case int:
	fmt.Printf("Twice %v is % %v \n", v, v * 2)
case string:
	fmt.Printf("%q is %v bytes long\n", v, len(v))
default:
	fmt.Printf("I don't know about the type %T!\n", v)
}

/************************************** Stringers ****************************/

/* Stringer is implemented by any value that has a String method, 
which defines the “native” format for that value. 
The String method is used to print values passed as an operand 
to any format that accepts a string 
or to an unformatted printer such as Print. */


// example 1
type Person struct {
	Name string
	Aget int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years", p.Name, p.age)
}

a := Person("Arthur Dent", 42)
z := Person("Zaphod Beeblebrox", 9901)
fmt.Println(a, z)

// example 2
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
	
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/************************************** Errors ****************************/


// example 1

type MyError struct {
	When time.time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError {
		time.Now(),
		"it didn't work",
	}
}

if err := run(); err != nil {
	fmt.Println(err)
}

// example 2

type ErrNegativeSqrt Error() string {
	return fmt.Sprintf("No sqrt for the negative number %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1.5)
	val := float64(0)
	for {
		z = z - (z*z - x)/(2*z)
		if math.Abs(val-z) < 1e-10 {
			break
		}
		val = z
	}
	return val, nil
}

fmt.Println(Sqrt(2))
fmt.Println(Sqrt(-2))


/************************************** Readers ****************************/

//Read populates the given byte slice with data
// and returns the number of bytes populated and an error value
// it returns an io.EOF when the stream ends

r := strings.NewReader("Hello, Reader!")
b := make([]byte, 8)
for {
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
	if err == io.EOF {
		break
	}
}

// implements a Reader type that emits an infinite stream of the ASCII character 'A'
type MyReader struct{}

func (MyReader) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		b[i] = 65
	}
	return len(b), nil
}


func main() {
	reader.Validate(MyReader{})
}

// rot13Reader
type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return
}

s := strings.NewReader("lbn frgfdslkjhgf ffs!")
r := rot13Reader{s}
io.Copy(os.Stdout, &r)


/************************************** Images ****************************/
type Image struct{
	pixels [][]color.Color
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0, len(i.pixels[0]), len(i.pixels))
}

func (i Image) At(x, y int) color.Color {
	return i.pixels[y][x]
}

func Pic(dx, dy int) [][]color.Color {
	sl := make([][]color.Color, dy)
	for i := range sl {
		sl[i] = make([]color.Color, dx)
		for j := 0; j < dx; j++ {
			sl[i][j] = color.RGBA{uint8((i+j)/2 + i ^ j), uint8((i+j)/2 + i ^ j), 255, 255}
		}
	}
	return sl
}

func main() {
	m := Image{Pic(50,100)}
	pic.ShowImage(m)
}

/************************************** goroutine ****************************/

// go starts a new goroutine
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

go say("world")
say("hello")

/************************************** channels ****************************
channels are a typed concuit through which you can send and receive values
by default channels are unbuffered, meaning that they will only accept sends
(chan <-) if there is a corresponding receive ready to receive the sent value(<-chan)
******************************************************************************/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

s := []int{7, 2, 8, -9, 4, 0}
c := make(chan int)
go sum(s[:len(s)/2], c)
go sum(s[len(s)/2:], c)
x, y := <-c, <-c
fmt.Println(x, y, x + y)

// buffered channels, the buffer length as the second argument to make to initialize a buffered channel
// sends to a buffered channel block only when the buffer is full
// if there are concurrent receive, the number of send is not limited

ch := make(chan int, 2)
go func() { 
	ch <- 11
	ch <- 12
	ch <- 13
}()
go func() { ch <- 2 }()
go func() { ch <- 3 }()
ch <- 4
ch <- 5

fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)	
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)

// range and close

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x 
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// The loop for i := range c receives values from the channel repeatedly until it is closed. 
	for i := range c {
		fmt.Println(i)
	}
}
 
// select statement lets a goroutine wait on multiple communication operations

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x: // send x to channel(there is a receiver waiting for the signal)
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
		}
	}
}

c := make(chan int)
quit := make(chan int)
go func() {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

fibonacci(c, quit)

// default selection
// use a default case to try a send or receive without blocking

tick := time.Tick(100 * time.Millisecond)
boom := time.After(500 * time.Millisecond)
for {
	select {
	case <-tick:
		fmt.Println("tick.")
	case <-boom:
		fmt.Println("BOOM!")
		return
	default:
		fmt.Println("  .")
		time.Sleep(50 * time.Millisecond)
	}
}


// Exercise: Equivalent Binary Trees using Go's concurrency and channels

type Tree struct {
	left *Tree
	Value int
	Right *Tree
}


func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkRecursive(t.left, ch)
		ch <- t.value
		WalkRecursive(t.Right, ch)
	}
}


func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int) // need to declare twice here, unlike in the function parameters
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		if ok1 != ok2 || n1 != n2 {
			return false
		}
		if !ok1 {
			break
		}
	}

	return true
}

/*********************************** sync.Mutex ****************************
mutual exclusives
use defer to ensure the mutex will be unlocked
******************************************************************************/

//SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// lock so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// lock so only one goroutine at a time can access the map c.v
	defer c.mu.Unlock()
	return c.v[key]
}

c := SafeCounter{v: make(map[string]int)}
for i := 0; i < 100; i++ {
	go c.Inc("somekey")
}

time.Sleep(time.Second)
fmt.Println(c.Value("somekey"))

/***************************************************************************
**************         Exercise Web Crawler with mutual exclusion
****************************************************************************/
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCounter struct {
	v map[string]bool
	mux sync.Mutex
}

var c SafeCounter = SafeCounter{v: make(map[string]bool)}
func (s SafeCounter)checkvisited(url string)bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.v[url]
	if ok == false {
		s.v[url] = true
		return false
	}
	return true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	
	if c.checkvisited(url) {
		return;
	}
	
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(5*time.Second)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
