run:
	go run app/main.go 

tidy:
	go mod tidy
	go mod vendor

test-endpoint-local:
	curl -il localhost:3000/test

query-local:
	@curl -s "http://localhost:3000/users?page=1&rows=2&orderBy=name,ASC" | jq