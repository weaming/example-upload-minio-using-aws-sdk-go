# Example of upload image to minio using `aws-sdk-go` 
## Steps

1. `./docker-run-minio.sh`
2. open [http://localhost:9001](http://localhost:9001), create bucket `test`, and set "Access Policy" to "public"
3. `go run .`
4. open [http://localhost:9000/test/img/example.jpg](http://localhost:9000/test/img/example.jpg)