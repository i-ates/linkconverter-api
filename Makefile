check-swagger:
	which swagger || (go get -d github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor && swagger generate spec -o ./swagger.json --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger -p 3001 swagger.json