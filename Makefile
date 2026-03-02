BINARY      := curve-generate
COLLECT_BIN := curve-collect
SERVER_BIN  := curve-server

.PHONY: build build-collect generate collect collect-codegen collect-contentmod collect-algotrade collect-support collect-credit collect-medicaldx collect-legalai collect-hire collect-education pipeline frontend server run-server dev deploy clean test

build:
	go build -o $(BINARY) ./cmd/generate

build-collect:
	go build -o $(COLLECT_BIN) ./cmd/collect

COLLECT_CMD := ./$(COLLECT_BIN) -seed seed/seed.json -overrides seed/overrides.yaml -log seed/collect.log.json

collect: build-collect
	$(COLLECT_CMD)
	@echo "seed.json updated. Run 'make generate' to rebuild Parquet."

collect-contentmod: build-collect
	$(COLLECT_CMD) -domain content-mod
collect-algotrade: build-collect
	$(COLLECT_CMD) -domain algo-trade
collect-codegen: build-collect
	$(COLLECT_CMD) -domain code-gen
collect-support: build-collect
	$(COLLECT_CMD) -domain support
collect-credit: build-collect
	$(COLLECT_CMD) -domain credit
collect-medicaldx: build-collect
	$(COLLECT_CMD) -domain medical-dx
collect-legalai: build-collect
	$(COLLECT_CMD) -domain legal-ai
collect-hire: build-collect
	$(COLLECT_CMD) -domain hire
collect-education: build-collect
	$(COLLECT_CMD) -domain education

pipeline: collect generate frontend

generate: build
	./$(BINARY) -input seed/seed.json -output frontend/static/data

frontend: generate
	cd frontend && npm ci && npm run build

server: frontend
	rm -rf cmd/server/static
	cp -r frontend/build cmd/server/static
	go build -o $(SERVER_BIN) ./cmd/server

run-server: server
	./$(SERVER_BIN) -port 8080

dev: generate
	cd frontend && npm run dev

deploy:
	fly deploy --remote-only

test:
	go test ./...

clean:
	rm -f $(BINARY) $(COLLECT_BIN) $(SERVER_BIN)
	rm -f frontend/static/data/*.parquet
	rm -rf frontend/build
	rm -rf cmd/server/static
