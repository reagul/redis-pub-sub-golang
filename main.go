package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//

func main() {
	/* // main function

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	print(pong)

	pubsub := rdb.Subscribe("PODLIST")

	// Wait for confirmation that subscription is created before publishing anything.
	for i := 0; i < 2; i++ {
		// ReceiveTimeout is a low level API. Use ReceiveMessage instead.
		msgi, err := pubsub.ReceiveTimeout(time.Second)
		if err != nil {
			break
		}

		switch msg := msgi.(type) {
		case *redis.Subscription:
			fmt.Println("subscribed to", msg.Channel)

			/*
				_, err := rdb.Publish("PODLIST", "hello").Result()
				if err != nil {
					panic(err)
				}
		case *redis.Message:
			fmt.Println("received", msg.Payload, "from", msg.Channel)
		default:
			panic("unreached")
		}

	} */

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	podlistchannel := client.Subscribe("PODLIST")
	somemsg, err := podlistchannel.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	//for {
	fmt.Println(somemsg.Payload)
	if somemsg.Payload == "blog" {
		fmt.Println("kickoff new workglow")
	}
	//time.Sleep(time.Second)
	//}

}
