# ping-service
.PHONY: run-ping-service
# run service :-->: run ping-service
run-ping-service:
	go run ./testdata/ping-service/cmd/... -conf=./testdata/ping-service/configs

# ping-service
.PHONY: run-service
# run service :-->: run ping-service
run-service:
	#@$(MAKE) run-ping-service
	go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs

.PHONY: testing-ping-service
# testing service :-->: testing ping-service
testing-ping-service:
	curl http://127.0.0.1:10101/api/v1/ping/logger && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/error && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/panic && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/say_hello && echo "\n"
	#curl http://127.0.0.1:10101/api/v1/ping/http_and_grpc && echo "\n"

.PHONY: testing-service
# testing service :-->: testing ping-service
testing-service:
	@$(MAKE) testing-ping-service

# ping-service
.PHONY: run-all-in-one
# run service :-->: run all-in-one
run-all-in-one:
	go run ./testdata/ping-service/cmd/all-in-one/main.go -conf=./testdata/ping-service/configs

.PHONY: testing-all-in-one
# testing service :-->: testing all-in-one
testing-all-in-one:
	curl http://127.0.0.1:10101/api/v1/ping/logger && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/error && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/panic && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/say_hello && echo "\n"


