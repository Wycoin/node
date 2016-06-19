package main
//go full ham on go,  all in http request
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "os/exec"
    "os"
//    "encoding/json"
    "math/rand"
    "time"
)



func main () {

    koen := 100
    matti := 20


    app := gin.Default()
    fmt.Print("you've made it heree.. your life sucks")
    sendLog("success", "Registration completed",string(random(2000,4000)))


    app.GET("transaction/koen/matti/:value", func(c *gin.Context){
        value := c.Param("value")
        value = int(value)
        if (koen - value) > 0 {
            koen = koen - value
            matti = matti + value
        }
        for i := 1;  i<=5; i++ {
                sendLog("verifing transaction {348>0675}", "success", string(random(2000,400)))
        }
        c.JSON(200, "{status:success}")
    })

    app.GET("transaction/matti/koen/:value", func(c *gin.Context){
        value := c.Param("value")
        value = int(value)
        if (matti - value)> 0 {
            koen = koen + value
            matti = matti - value
        }
        for i := 1;  i<=5; i++ {
                sendLog("verifing transaction {0675>348}", "success", string(random(2000,400)))
        }

        c.JSON(200, "{status:success}")
    })

    app.Run(":"+os.Getenv("PORT"))
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}


func sendLog(status,message, port string){
    exec.Command("curl -XPOST -F 'source=http://192.168.99.100:"+port+"' -F 'event="+message+"' -F 'status="+status+"' "+os.Getenv("LOGGER"))
}
