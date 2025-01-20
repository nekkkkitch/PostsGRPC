buildauthpb: 
	protoc --proto_path=proto/auth --go_out=pb/auth --go-grpc_out=pb/auth proto/auth/auth.proto
buildnotificatorpb: 
	protoc --proto_path=proto/notificator --go_out=pb/notificator --go-grpc_out=pb/notificator proto/notificator/*.proto
buildpostspb: 
	protoc --proto_path=proto/posts --go_out=pb/posts --go-grpc_out=pb/posts proto/posts/*.proto
buildsubspb: 
	protoc --proto_path=proto/subs --go_out=pb/subs --go-grpc_out=pb/subs proto/subs/*.proto
