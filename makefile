gen:
	protoc --proto_path=model --go_out=plugins=grpc:model todo-service.proto
	protoc-go-inject-tag -input=model/todo-service.pb.go


clean:
	rm model/*.pb.go
