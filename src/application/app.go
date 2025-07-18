package application

import "svc-subman/src/application/commands"

type Application struct {
	Commands Command
}
type Command struct {
	CreateSubscribe commands.CreateSubscribeCommand
	DeleteSubscribe commands.DeleteSubscribeCommand
	AgreeSubscribe  commands.AgreeSubscriberQuvery
	GetSubscribe    commands.GetSubscriberQuvery
	GetSubscribes   commands.GetSubscribesQuvery
}
