package telemetrytags

type app struct {
	DefaultSourceName string
	Consumer          string
	EventHandler      string
	Subscription      string
	Stream            string
	Partition         string
	Request           string
	RequestName       string
	RequestResultName string
	RequestResult     string
	Command           string
	CommandName       string
	CommandResult     string
	CommandResultName string
	Query             string
	QueryName         string
	QueryResult       string
	QueryResultName   string
	Event             string
	EventName         string
	EventResult       string
	EventResultName   string
}

// https://opentelemetry.io/docs/specs/otel/trace/semantic_conventions/messaging/
type messaging struct {
	System          string
	Destination     string
	DestinationKind string
	Url             string
	MessageId       string
	ConversationId  string
	CorrelationId   string
	CausationId     string
	Operation       string
}

type exceptions struct {
	EventName  string
	Type       string
	Message    string
	Stacktrace string
}

type general struct{}

// https://opentelemetry.io/docs/specs/otel/trace/semantic_conventions/rpc/
type grpc struct{}

type netHttp struct {
	Transport string
	PeerIp    string
	PeerPort  string
	PeerName  string
	HostIp    string
	HostPort  string
	HostName  string
}

var App = app{
	DefaultSourceName: "app",
	Consumer:          "app.consumer",
	EventHandler:      "app.event-handler",
	Subscription:      "app.subscription",
	Stream:            "app.stream",
	Partition:         "app.partition",
	Request:           "app.request",
	RequestName:       "app.request_name",
	RequestResultName: "app.request_result_name",
	RequestResult:     "app.request_result",
	Command:           "app.command",
	CommandName:       "app.command_name",
	CommandResult:     "app.command_result",
	CommandResultName: "app.command_result_name",
	Query:             "app.query",
	QueryName:         "app.query_name",
	QueryResult:       "app.query_result",
	QueryResultName:   "app.query_result_name",
	Event:             "app.event",
	EventName:         "app.event_name",
	EventResult:       "app.event_result",
	EventResultName:   "app.event_result_name",
}

var Exceptions = exceptions{
	EventName:  "exception",
	Type:       "exception.type",
	Message:    "exception.message",
	Stacktrace: "exception.stacktrace",
}

var General = general{}

var Grpc = grpc{}

var Messaging = messaging{
	System:          "messaging.system",
	Destination:     "messaging.destination",
	DestinationKind: "messaging.destination_kind",
	Url:             "messaging.url",
	MessageId:       "messaging.message_id",
	ConversationId:  "messaging.conversation_id",
	CorrelationId:   "messaging.correlation_id",
	CausationId:     "messaging.causation_id",
	Operation:       "messaging.operation",
}

var NetHttp = netHttp{
	Transport: "net.transport",
	PeerIp:    "net.peer.ip",
	PeerPort:  "net.peer.port",
	PeerName:  "net.peer.name",
	HostIp:    "net.host.ip",
	HostPort:  "net.host.port",
	HostName:  "net.host.name",
}
