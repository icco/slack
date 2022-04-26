module github.com/slack-go/slack/examples/custom_logger

go 1.16

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/slack-go/slack v0.10.3
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0
)

replace github.com/slack-go/slack v0.10.3 => ../../
