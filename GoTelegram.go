package GoTelegram

import(
        "time"
        //log "./golang/src/github.com/r4y/GoTelegram/log"
        logging "github.com/r4y/GoTelegram/logging"
        "fmt"
)
type Token struct {

     Value string

}

func (t Token) BuildBot() {
 logger := logging.New(time.RFC3339, true)
 logger.Log("info", "starting up service")
 return fmt.Sprintf("here is the token: %d ", t.Value)
}

func Hello() {
    fmt.Println("Hello, World!")
}
