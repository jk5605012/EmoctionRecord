.PHONY: all deps rpm docker docker-deps docker-cgo docker-push clean docs test test-race test-integration fmt lint install

test:
	docker run -d --name database -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_DB=test -e POSTGRES_PASSWOR=mysecretpassword -e POSTGRES_HOST_AUTH_METHOD=trust postgres:9.6
	sleep 30
	@go test -v ./...
	docker stop database
	docker rm database