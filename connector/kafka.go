package connector

import (
	"context"
	"crypto/tls"
	"golang.org/x/net/proxy"
	"net"
	"time"
	"strings"
	"github.com/Shopify/sarama"
)
import "gocloud.dev/pubsub/kafkapubsub"

type MsgProvider struct {
	AWSSNS interface{}
	AWSSQS interface{}
	GCPPUBSUB interface{}
	AZURESBUS interface{}
}
func KafkaMessageBusConnector()  {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()
	var mProvider *MsgProvider
	switch mProvider {
	case mProvider.AWSSNS:
		msgBody, err := SubscribeKafkaTopic(ctx)
		if err != nil {
			log.Printf("SubscribeMessageFailed [ERROR]:%s", err.Error())
			return
		}
		if err := SendMessageToSNS(msgBody, ctx); err != nil {
			log.Printf("SendMessageToSNSFailed [ERROR]:%s", err.Error())
			return
		}
	case mProvider.AWSSQS:
		msgBody, err := SubscribeKafkaTopic(ctx)
		if err != nil {
			log.Printf("SubscribeMessageFailed [ERROR]:%s", err.Error())
			return
		}
		if err := SendMessageToSQS(msgBody, ctx); err != nil {
			log.Printf("SendMessageToSNSFailed [ERROR]:%s", err.Error())
			return
		}
	case mProvider.AZURESBUS:
		// TODO: Pending with Azure ServiceBus
	case mProvider.GCPPUBSUB:
		msgBody, err := SubscribeKafkaTopic(ctx)
		if err != nil {
			log.Printf("SubscribeMessageFailed [ERROR]:%s", err.Error())
			return
		}
		if err := SendMessageToPubsub(msgBody, ctx); err != nil {
			log.Printf("SendMessageToSNSFailed [ERROR]:%s", err.Error())
			return
		}
	default:
		log.Fatalln("Provider not support")
	}
}

type KafkaConnectionDetail struct {
	Broker string
	Brokers []string
	Zookeeper string
	Zookeepers []string
	Topic string
	Topics []string
	Partition int64
}
func (bsc *BaseConfig)ConfigConvert() (kcd *KafkaConnectionDetail)  {
	kcd.Broker = bsc.Broker
	kcd.Zookeeper = bsc.Zookeeper
	kcd.Topic = bsc.Topic
	return
}
func (kc *KafkaConnectionDetail)SplitZKBK (kf *KafkaConnectionDetail) {
	ffbks:= strings.Split(kc.Broker, ";")
	ffzks:= strings.Split(kc.Zookeeper, ";")
	fftps:= strings.Split(kc.Topic, ";")
	kf.Brokers = ffbks
	kf.Zookeepers = ffzks
	kf.Topics = fftps
	return
}
func SubscribeKafkaTopic(ctx context.Context) ([]byte,error) {
	var kakfkDetail *KafkaConnectionDetail
	open, err := kafkapubsub.OpenSubscription(kakfkDetail.Brokers, &sarama.Config{
		Admin: struct {
			Retry struct {
				Max     int
				Backoff time.Duration
			}
			Timeout time.Duration
		}{},
		Net: struct {
			MaxOpenRequests int
			DialTimeout     time.Duration
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			TLS             struct {
				Enable bool
				Config *tls.Config
			}
			SASL struct {
				Enable                   bool
				Mechanism                sarama.SASLMechanism
				Version                  int16
				Handshake                bool
				AuthIdentity             string
				User                     string
				Password                 string
				SCRAMAuthzID             string
				SCRAMClientGeneratorFunc func() sarama.SCRAMClient
				TokenProvider            sarama.AccessTokenProvider
				GSSAPI                   sarama.GSSAPIConfig
			}
			KeepAlive time.Duration
			LocalAddr net.Addr
			Proxy     struct {
				Enable bool
				Dialer proxy.Dialer
			}
		}{},
		Metadata: struct {
			Retry struct {
				Max         int
				Backoff     time.Duration
				BackoffFunc func(retries int, maxRetries int) time.Duration
			}
			RefreshFrequency time.Duration
			Full             bool
			Timeout          time.Duration
		}{},
		Producer: struct {
			MaxMessageBytes  int
			RequiredAcks     sarama.RequiredAcks
			Timeout          time.Duration
			Compression      sarama.CompressionCodec
			CompressionLevel int
			Partitioner      sarama.PartitionerConstructor
			Idempotent       bool
			Return           struct {
				Successes bool
				Errors    bool
			}
			Flush struct {
				Bytes       int
				Messages    int
				Frequency   time.Duration
				MaxMessages int
			}
			Retry struct {
				Max         int
				Backoff     time.Duration
				BackoffFunc func(retries int, maxRetries int) time.Duration
			}
		}{},
		Consumer: struct {
			Group struct {
				Session struct {
					Timeout time.Duration
				}
				Heartbeat struct {
					Interval time.Duration
				}
				Rebalance struct {
					Strategy sarama.BalanceStrategy
					Timeout  time.Duration
					Retry    struct {
						Max     int
						Backoff time.Duration
					}
				}
				Member struct {
					UserData []byte
				}
			}
			Retry struct {
				Backoff     time.Duration
				BackoffFunc func(retries int) time.Duration
			}
			Fetch struct {
				Min     int32
				Default int32
				Max     int32
			}
			MaxWaitTime       time.Duration
			MaxProcessingTime time.Duration
			Return            struct {
				Errors bool
			}
			Offsets struct {
				CommitInterval time.Duration
				AutoCommit     struct {
					Enable   bool
					Interval time.Duration
				}
				Initial   int64
				Retention time.Duration
				Retry     struct {
					Max int
				}
			}
			IsolationLevel sarama.IsolationLevel
		}{},
		ClientID:          "",
		RackID:            "",
		ChannelBufferSize: 0,
		Version:           sarama.KafkaVersion{},
		MetricRegistry:    nil,
	}, "",kakfkDetail.Topics, nil )
	if err != nil {
		return nil, err
	}
	msg, err := open.Receive(ctx)
	if err != nil {
		return nil, err
	}
	return msg.Body, nil
}