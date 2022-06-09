.SILENT:

# generate api contracts for all proto files
gen:
	for filename in api/*.proto; do \
		fname=$$(basename -- "$$filename" .proto); \
		echo $$fname; \
		mkdir -p pkg/$$fname; \
		protoc --go_out=pkg/$$fname --go-grpc_out=pkg/$$fname --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/$$fname.proto; \
	done \

run:
	docker-compose up

rebuild:
	docker-compose up --build --remove-orphans

stop:
	docker-compose down