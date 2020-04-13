package main

/*
	There are several ways to handle files in Go.

	Either through helper libraries like io/ioutil, or directly via buffers/scanners

*/

var (
	fileToRead  string = "./read_this.txt"
	fileToWrite string = "./write_this.txt"
)

func main() {
	readFileViaIoutil()
	readFileViaBuffer()

	writeFileViaIoutil()
	writeFileViaBufio()
}
