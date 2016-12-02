package services

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/mono83/slf/wd"
)

var Router *mux.Router

var RedisPool *pool.Pool

var RabbitMQChannel *amqp.Channel

var RootFolder string

var Logger wd.Watchdog
