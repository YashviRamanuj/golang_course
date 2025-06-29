type Logger interface {
    Log(msg string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(msg string) {
    fmt.Println(msg)
}

type FileLogger struct {
    filepath string
}

func (f FileLogger) Log(msg string) {
    // simplified version
    os.WriteFile(f.filepath, []byte(msg), 0644)
}

func PrintReport(l Logger) {
    l.Log("Report generated at " + time.Now().String())
}

func defaultLog(msg string) {
    fmt.Println(msg)
}

func main(a Logger) {
    // PrintReport(ConsoleLogger{})
    var t type
    t = type(a)
    switch t {
        case ConsoleLogger:
            PrintReport(ConsoleLogger{})
        case FileLogger:
            PrintReport(FileLogger{"report.txt"})
        case SlackLogger:
            PrintReport(SlackLogger{})
        case abc:
            PrintReport(abc{})
        case xyx:
            PrintReport(xyx{})
        case sentry:
            PrintReport(sentry{})
        default:
            defaultLog("Report generated at " + time.Now().String())
}



// if u wanna add a new logger 

type SlackLogger struct{}

func (s SlackLogger) Log(msg string) {
    // send to Slack API
}


type abc struct {}

func (a abc) Log(msg string) {
    // logic
}

type xyx

type sentry struct {}

func (s sentry) Log(msg string) {
    // logic
}


type human struct {
    name string
    age int
}

type dog struct {
    name string
    owner string
}

func main() {
    h := human{"John", 30}
    d := dog{"Buddy", "John"}
  
    d.(human)
}

    