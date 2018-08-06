package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	tracker "github.com/remerge/go-tracker"
	"github.com/remerge/sarama"
	users "mod/github.com/remerge/users-proto-go@v0.0.0-20180518181310-98b61f308deb"
)

func newConsumerGroup(kafkaID string, kafkaBrokers string, topic string) (*sarama.ConsumerGroup, error) {
	kafkaClient, err := sarama.NewClientWithID(kafkaID, kafkaBrokers)
	if err != nil {
		return nil, err
	}
	kafkaClient.Config().Version = sarama.V0_10_0_0
	kafkaClient.Config().Group.AutoCommit = false
	kafkaClient.Config().Group.Return.Notifications = true

	group, err := sarama.NewConsumerGroupFromClient(kafkaClient, kafkaID, []string{topic})
	if err != nil {
		return nil, err
	}
	return &group, err
}

func readUsersUs6(brokers []string) (error, sarama.ConsumerGroup) {
	consumerName := "gb_test_01"
	consumerGroup, err := newConsumerGroup(consumerName, brokers[0], "users_us5")
	if err != nil {
		return err, nil
	}
	_, ok := <-(*consumerGroup).Notifications()
	if !ok {
		return errors.New("Error rebalancing"), nil
	}
	fmt.Println("Inited")

	return nil, (*consumerGroup)
}

func forwardFnc(cg sarama.ConsumerGroup, tracker *tracker.KafkaTracker, topic string) {
	for {
		select {
		case message, ok := <-cg.Messages():
			if !ok {
				fmt.Println("FUP")
				return
			}
			key := message.Key
			value := message.Value

			newID := users.User{}
			err := proto.Unmarshal(message.Value, &newID)
			if err != nil {
				fmt.Println(err)
				return
			}
			var jb []byte
			jb, err = json.Marshal(newID)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(jb))

			// Now encode to strings
			str := hex.EncodeToString(message.Value)
			fmt.Println(str)

			err = tracker.SafeMessageWithKey(topic, value, key)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("forwarded")
			cg.MarkMessage(message, "")
			cg.Commit()

			time.Sleep(10 * time.Second)
		case err, ok := <-cg.Errors():
			if !ok {
				fmt.Println("FUP")
				return
			}
			fmt.Println(err)
			return
		}
	}
}

func main() {
	brokers := []string{"kafka:9092"}
	topic := "users1"

	//brokers := []string{"fh1.dw1.remerge.io:9092"}
	//topic := "users_us6"
	metadata := tracker.EventMetadata{}
	tracker, err := tracker.NewKafkaTracker(brokers, &metadata)
	if err != nil {
		fmt.Println(err)
		return
	}

	//err, consumerGroup := readUsersUs6(brokers)
	if err != nil {
		fmt.Println(err)
		return
	}

	//forwardFnc(consumerGroup, tracker, topic)

	user := users.User{}
	user.Id = "e2e99785-d252-4de1-8e37-b4d845c80a79"
	user.IdType = users.User_AAID
	user.AudienceIds = []int32{1743, 1234}

	bytes, err := proto.Marshal(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	key := []byte(user.GetId())
	err = tracker.SafeMessageWithKey(topic, bytes, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
	// Now encode to strings
	str := hex.EncodeToString(bytes)
	fmt.Println(str)

	fmt.Printf("Message sent to %s and %s \n", brokers[0], topic)
}
