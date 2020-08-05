.PHONY: all deps rpm docker docker-deps docker-cgo docker-push clean docs test test-race test-integration fmt lint install

GO_FLAGS =

test:
	@go test $(GO_FLAGS) -short ./...