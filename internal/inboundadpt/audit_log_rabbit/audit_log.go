package audit_log_rabbit

import (
	"context"
	"log"
	"time"

	audit_log_grpc "auditLog/internal/inboundadpt/audit_log_grpc"

	"github.com/streadway/amqp"
)

type RabbitURI string
type RabbitExchangeName string
type RabbitExchangeType string
type RabbitQueueName string

type RabbitMQ struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	consumer   <-chan amqp.Delivery
	errorChan  chan *amqp.Error
	msgHandler audit_log_grpc.Service
}

func New(uri RabbitURI, auditSrv audit_log_grpc.Service) (*RabbitMQ, error) {
	conn, err := amqp.Dial(string(uri))
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	exchange := "logs"
	if exchange == "" {
		exchange = "default_exchange"
	}

	queue, err := channel.QueueDeclare(string("queueName"), true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	if err := channel.ExchangeDeclare(
		exchange,
		amqp.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}
	//relationship between exchange and a queue
	if err := channel.QueueBind(queue.Name, "", exchange, false, nil); err != nil {
		return nil, err
	}

	consumer, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	r := &RabbitMQ{
		conn:       conn,
		channel:    channel,
		queue:      queue,
		consumer:   consumer,
		errorChan:  make(chan *amqp.Error),
		msgHandler: auditSrv,
	}

	go r.consume()

	return r, nil
}

func (r *RabbitMQ) consume() {
	var cancel context.CancelFunc

	for msg := range r.consumer {
		var ctx context.Context
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		if err := r.msgHandler.LogAudit(ctx, msg.Body); err != nil {
			log.Printf("Failed to process message: %s", err.Error())
		} else {
			if err := msg.Ack(false); err != nil {
				log.Printf("Failed to acknowledge message: %s", err.Error())
			}
		}
	}

	if cancel != nil {
		cancel()
	}
}

func (r *RabbitMQ) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}

	if err := r.conn.Close(); err != nil {
		return err
	}

	return nil
}
