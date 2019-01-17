.PHONY: start-services stop-services gen-proto

start-services:
	cd ./docker && docker-compose -f ./services-compose.yml up -d

stop-services:
	cd ./docker && docker-compose -f ./services-compose.yml down

gen-proto:
	protoc -I ../pkg/pb --go_out=plugins=grpc:../pkg/pb ../pkg/pb/pb.proto
