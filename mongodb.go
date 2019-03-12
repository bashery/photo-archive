package main
import (
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Person struct {
    Name string
    Phone string
}

func main() {
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonicbehavior.
    // session.SetMode(mgo.Monotonic, true)
    c := session.DB("test").C("people")
    err = c.Insert(&Person{"Ale", "+55 53 8116" },
                    &Person{"alae", "+9055 666 777"})
    if err != nil {
        panic(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("name:", result.Name)
    fmt.Println("Phone:", result.Phone)
}
