package kafka

import (
        "fmt"
        "Libraries/libspark-msil/constants"
        "github.com/Shopify/sarama"
        log "github.com/sirupsen/logrus"
        "strings"
        "encoding/json"
        "Libraries/libspark-msil/entities"
        "time"
)

var Producer sarama.AsyncProducer
var err error

func KClient() {
        fmt.Println("start method - KClient")
        servers := strings.Split(constants.KakfaServer, ",")
        config := sarama.NewConfig()
        config.Producer.Timeout = 2 * 1000
        config.Producer.Partitioner = sarama.NewRandomPartitioner
        Producer, err = sarama.NewAsyncProducer(servers, config)
        if err != nil {
                log.WithFields(log.Fields{"error": err}).Errorln(constants.KafkaConnectErrorMsg)
                panic(err)
        }
        log.WithFields(log.Fields{"method": "KClient", "server": servers}).Infoln(constants.KafkaConnect)
        fmt.Println("end method - KClient")
}

func Produce(channel chan struct{}, topic string, payload []byte) {

        select {
        case <-channel:
                return
                case Producer.Input() <- &sarama.ProducerMessage{
			 Topic: topic,
                        Key:   nil,
                        Value: sarama.ByteEncoder(payload),
                }:
                log.WithFields(log.Fields{"method": "KafkaProduce", "topic": topic, "payload": string(payload)}).Debugln(constants.PostedToKafka)
        case err := <-Producer.Errors():
                log.WithFields(log.Fields{"error": err, "method": "KafkaProduce"}).Errorln(constants.KafkaProducerErrorMsg)
        }
}

func CloseProducer() {
        if err := Producer.Close(); err != nil {
                log.WithFields(log.Fields{"error": err}).Errorln(constants.KafkaCloseProducerErrorMsg)
        }
}

func SendLogsToKafka( topic string,  message string , messageID string ,path string ,payload interface{}) {
        channel := make(chan struct{})
        //topic := constants.NotificationLog
        LogStatement := entities.Log{
                Time : time.Now(),
                MessageID : messageID,
                Service : path,
                PayloadType : message,
                Payload : payload,
        }

        LogStatement_,_ := json.Marshal(LogStatement)

        Produce(channel, topic , LogStatement_ )
}

