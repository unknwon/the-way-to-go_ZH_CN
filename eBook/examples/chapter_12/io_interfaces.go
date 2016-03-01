// interfaces being used in the GO-package fmt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf.
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
}
