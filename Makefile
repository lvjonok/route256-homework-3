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

s-warehouse:
	docker-compose up --force-recreate --no-deps --build -d srv-warehouse

s-marketplace:
	docker-compose up --force-recreate --no-deps --build -d srv-marketplace

s-orders:
	docker-compose up --force-recreate --no-deps --build -d srv-orders

run-all:
	docker-compose up

rebuild:
	docker-compose up --build

stop:
	docker-compose down