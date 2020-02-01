package DI

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "world")
// }
//
// func main() {
// 	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
// }
