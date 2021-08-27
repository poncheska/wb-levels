package shell

import "io"

type Shell struct{
	StdIn  io.ReadWriter
	StdOut io.ReadWriter
	StdErr io.ReadWriter
}
