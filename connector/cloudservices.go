package connector

import (
	"context"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/awssnssqs"
	//servicebus "github.com/Azure/azure-service-bus-go"
)

/*
	----------------------
	For AWS SNS && AWS SQS
	----------------------
*/
var AWSSession *aws.Session

type MessagePrviderDetail struct {
	Qurl string `json:"qurl"`
	TopicArn string `json:"topic_arn"`
	PubSubTopic string `json:"pub_sub_topic"`
	MessageBusAddr string `json:"message_bus_addr"`
}

func SendMessageToSQS(MessageBody []byte, ctx context.Context) error {
	var msgDetail *MessagePrviderDetail
	ope  := awssnssqs.OpenSQSTopic(ctx,AWSSession, msgDetail.Qurl, nil)
	if err :=ope.Send(ctx, &pubsub.Message{
		Body:       MessageBody,
		Metadata:   nil,
		BeforeSend: nil,
	});err != nil {
			return err
	}
	defer ope.Shutdown(ctx)
	return nil
}
func SendMessageToSNS(MessageBody []byte, ctx context.Context) error  {
	var msgDetail *MessagePrviderDetail
	open := awssnssqs.OpenSNSTopic(ctx, AWSSession, msgDetail.TopicArn, nil)
	if err := open.Send(ctx, &pubsub.Message{
		Body:       MessageBody,
		Metadata:   nil,
		BeforeSend: nil,
	});err != nil {
			return err
	}
	defer open.Shutdown(ctx)
	return nil
}
/*
	-----------------------
	For Google Cloud PubSub
	-----------------------
*/
func SendMessageToPubsub(MessageBody []byte, ctx context.Context) error {
	var msgDetail *MessagePrviderDetail
	open, err := pubsub.OpenTopic(ctx, msgDetail.PubSubTopic)
	if err != nil {
		return err
	}
	if err := open.Send(ctx, &pubsub.Message{
		Body:       MessageBody,
		Metadata:   nil,
		BeforeSend: nil,
	});err != nil {
			return err
	}
	defer open.Shutdown(ctx)
	return nil
}
/*
	-------------------------------
	For Microsoft Azure ServicesBus
	TODO: Send To Azure Service Bus
	-------------------------------
*/
//func SendMessageToServicesBus(MessageBody []byte, ctx context.Context) error {
//	var msgDetail *MessagePrviderDetail
//	open, err := azuresb.OpenTopic(ctx, &servicebus.Topic{
//		nil, &servicebus.Sender{Name: msgDetail.MessageBusAddr}, nil
//	},nil )
//	if err != nil {
//		return err
//	}
//	if err := open.Send(ctx, &pubsub.Message{
//		Body:       MessageBody,
//		Metadata:   nil,
//		BeforeSend: nil,
//	});err != nil {
//		return err
//	}
//	return nil
//}