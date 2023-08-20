run:
	go run app/services/sales-api/main.go

tidy:
	go mod tidy
	go mod vendor

test-endpoint-local:
	curl -il localhost:3000/test

test-flyio:
	curl -il https://new-service.fly.dev/test

query-local:
	@curl -s "http://localhost:3000/users?page=1&rows=2&orderBy=name,ASC" | jq

query-flyio:
	@curl -s "https://new-service.fly.dev/users?page=1&rows=2&orderBy=name,ASC" | jq

run-admin:
	go run app/tooling/admin/main.go

metrics-view:
	expvarmon -ports="https://new-service.fly.dev:4000/" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

load:
	hey -m GET -c 100 -n 10000 https://new-service.fly.dev/users?page=1&rows=2