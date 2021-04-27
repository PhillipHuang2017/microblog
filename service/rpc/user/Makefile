build:
	go build -o ./build/user-rpc-server .

win:build
	cp ./build/user-rpc-server ./build/user-rpc-server.exe

proto:
	goctl rpc proto --src user.proto --dir .

sql:
	goctl model mysql ddl --src ./sql/user.sql --dir ./sql/model --cache

docker:
	goctl docker -go user.go