package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/gin-gonic/gin"
	"github.com/hack-caixa/domain"
	"gorm.io/gorm"
)

type EventHubConfig struct {
	CONNECTION_STRING string
	EVENT_HUB_NAME    string
}

func NewEventHub() *EventHubConfig {
	return &EventHubConfig{
		CONNECTION_STRING: os.Getenv("CONNECTION_STRING"),
		EVENT_HUB_NAME:    os.Getenv("EVENT_HUB_NAME"),
	}
}

func PublishInEventHub(ctx *gin.Context, result domain.SimulacaoDTO, db *gorm.DB) {

	env := map[bool]string{true: "test", false: "dev"}[db.Config.Dialector.Name() == "sqlite"]

	eventHub := NewEventHub()
	producerClient, err := azeventhubs.NewProducerClientFromConnectionString(eventHub.CONNECTION_STRING, eventHub.EVENT_HUB_NAME, nil)

	if err != nil {
		panic(err)
	}

	defer producerClient.Close(ctx)

	event := createEvent(result)

	// create a batch object and add sample event to the batch
	newBatchOptions := &azeventhubs.EventDataBatchOptions{}

	batch, err := producerClient.NewEventDataBatch(ctx, newBatchOptions)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = batch.AddEventData(event, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	if env == "test" {
		printOutput(event.Body)
		return
	}

	// send the batch of event to the event hub
	err = producerClient.SendEventDataBatch(ctx, batch, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Message sent to event hub successfully!")
}

func createEvent(result domain.SimulacaoDTO) *azeventhubs.EventData {

	body, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &azeventhubs.EventData{Body: body}

}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("=====> Output: %s\n", string(out))
	}
}
