build:
	go build -o ./build/user-api-server .

win:build
	cp ./build/user-api-server ./build/user-api-server.exe

api:
	goctl api go --api user.api --dir .

docker:
	goctl docker -go user.go