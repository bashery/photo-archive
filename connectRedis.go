// How to connect redis with golang app 

package main

import (
    "fmt"
    "time"
    "github.com/go-redis/redis"
)

func main() {

    var client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "", // no password set
        DB:  0 , // 0 is default db
    })

    client.Get("h")
    var ping, err = client.Ping().Result()
    fmt.Println(ping, err)


    err = client.Set("hi", "hello", time.Second*20).Err() // we need to use time.Durition
    if err != nil {
        panic(err)
    }
    val, err := client.Get("hi").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

}
