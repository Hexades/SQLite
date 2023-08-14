package sqlite

var localbus = sqliteBus{}

type sqliteBus struct {
	repositoryListenerChannels []chan Event
}

func AddListener(EventListener EventListener) {
	listenerChannel := make(chan Event, 10)
	localbus.repositoryListenerChannels = append(localbus.repositoryListenerChannels, listenerChannel)
	go EventListener.OnEvent(listenerChannel)
}

func SendEvent(Event Event) {
	for _, channel := range localbus.repositoryListenerChannels {
		channel <- Event
	}
}

type EventListener interface {
	OnEvent(Event <-chan Event)
}
