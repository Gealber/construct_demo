.PHONY: all generate_vendor build
all: generate_vendor build clean

.PHONY: up 
up: generate_vendor compose clean

compose:
	@docker-compose up -d

.PHONY: generate_vendor
generate_vendor: ; @go mod vendor

.PHONY: build
build:
	@docker image build -t gealber/construct_demo .

.PHONY: down
down:
	@docker-compose down

.PHONY: ps
ps:
	@docker-compose ps

.PHONY: clean
clean:
	@rm -r vendor
