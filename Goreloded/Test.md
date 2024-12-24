var i interface{} = x
fmt.Printf("%v\n", i) \\will print the x

/dev/stdin   /dev/stdout  /dev/stderr (0,1,2)
/dev/pts/0

# /////////////////( stdin ) /////////////////////////////

package os // import "os"
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
    Stdin, Stdout, and Stderr are open Files pointing to the standard input,
    standard output, and standard error file descriptors.

    Note that the Go runtime writes to standard error for panics and crashes;
    closing Stderr may cause those messages to go elsewhere, perhaps to a file
    opened later.

# /////////////////( stdout ) /////////////////////////////
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
    Stdin, Stdout, and Stderr are open Files pointing to the standard input,
    standard output, and standard error file descriptors.

    Note that the Go runtime writes to standard error for panics and crashes;
    closing Stderr may cause those messages to go elsewhere, perhaps to a file
    opened later.
# /////////////////( stderr ) /////////////////////////////
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
    Stdin, Stdout, and Stderr are open Files pointing to the standard input,
    standard output, and standard error file descriptors.

    Note that the Go runtime writes to standard error for panics and crashes;
    closing Stderr may cause those messages to go elsewhere, perhaps to a file
    opened later.

# /////////////////( io.Writer ) /////////////////////////////
package io // import "io"

type Writer interface {
	Write(p []byte) (n int, err error)
}
    Writer is the interface that wraps the basic Write method.

    Write writes len(p) bytes from p to the underlying data stream. It returns
    the number of bytes written from p (0 <= n <= len(p)) and any error
    encountered that caused the write to stop early. Write must return a non-nil
    error if it returns n < len(p). Write must not modify the slice data,
    even temporarily.

    Implementations must not retain p .

var Discard Writer = discard{}
func MultiWriter(writers ...Writer) Writer








#####
uintptr - NewFile - syscall - fileDescriptor - io.Writer - bufio 
- the deffirent between os.ReadFile and bufio.NewScanner
- 