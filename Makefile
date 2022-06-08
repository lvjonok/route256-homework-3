.SILENT:

# generate api contracts for all proto files
gen:
	for filename in api/*.proto; do \
		fname=$$(basename -- "$$filename" .proto); \
		echo $$fname; \
		mkdir -p pkg/$$fname; \
		protoc --go_out=pkg/$$fname --go-grpc_out=pkg/$$fname --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/$$fname.proto; \
	done \

srv-marketplace:
	# docker-compose -f docker-compose-utils.yml -f docker-compose-marketplace.yml up
	docker-compose -f docker-compose-marketplace.yml up --build

srv-marketplace-down:
	# docker-compose -f docker-compose-utils.yml -f docker-compose-marketplace.yml down
	docker-compose -f docker-compose-marketplace.yml down
	
# go run cmd/service-marketplace/main.go