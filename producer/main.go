package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	tracker "github.com/remerge/go-tracker"
	users "mod/github.com/remerge/users-proto-go@v0.0.0-20180518181310-98b61f308deb"
)

func main() {
	user := users.User{}
	user.Id = "12345-6789-01234"
	user.IdType = users.User_AAID
	user.AudienceIds = []int32{42, 34, 5, 6}

	brokers := []string{"kafka:9092"}
	metadata := tracker.EventMetadata{}
	tracker, err := tracker.NewKafkaTracker(brokers, &metadata)
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := proto.Marshal(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	key := []byte(user.GetId())
	topic := "users_us6"
	err = tracker.SafeMessageWithKey(topic, bytes, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Message sent to %s and %s \n", brokers[0], topic)
}
