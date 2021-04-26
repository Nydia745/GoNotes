import (
	"fmt"
	"io"
	"os"
)

// Println: spaces are always added
const name, age = "Kim", 22
fmt.Println(name, "is", age, "years old.")

// need to add spaces, standard output
fmt.Print(name, " is ", age, " years old.\n")

// return the resulting string
fmt.Sprintln(name, "is", age, "years old.")

// Sprintf formats according to a format specifier and return the resulting string
s := fmt.Sprintf("%s is %d years old.\n", name, age)

// Fprintf formats according to a format specifier and writes to w
fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)





