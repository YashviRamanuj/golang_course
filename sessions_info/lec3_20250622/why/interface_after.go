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




// if u wanna add a new logger 

type SlackLogger struct{}

func (s SlackLogger) Log(msg string) {
    // send to Slack API
}
