type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}


// ----------------------------
type FileReaderWriter struct {
    filename string
}

func (f FileReaderWriter) Read(p []byte) (n int, err error) {
    // read from file
    return len(p), nil
}

func (f FileReaderWriter) Write(p []byte) (n int, err error) {
    // write to file
    return len(p), nil
}

// ----------------------------
func main() {
    f := FileReaderWriter{"myfile.txt"}
    ReadAndWrite(f)
}

func ReadAndWrite(rw ReadWriter) {
    // read and write using the same interface
    data := make([]byte, 1024)
    rw.Read(data)
    rw.Write(data)
}