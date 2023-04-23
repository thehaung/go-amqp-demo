package domain

type SlackRepository interface {
	PushNotification() error
}

type SlackService interface {
	PushNotification() error
}
