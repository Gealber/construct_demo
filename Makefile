.PHONY: all
all: generate_vendor build clean

.PHONY: up 
up: generate_vendor compose

compose:
	@docker-compose up -d

.PHONY: generate_vendor
generate_vendor: ; @go mod vendor

.PHONY: build
build:
	@docker image build --rm -t gealber/construct_demo .

.PHONY: down
down:
	@docker-compose down

.PHONY: ps
ps:
	@docker-compose ps

ifdef OS
	RM = del /Q
else
	RM = rm -rf
endif

.PHONY: clean
clean:
	@$(RM) vendor
ifneq "$(shell docker images -f dangling=true -q --no-trunc)" ""
	@docker rmi $(shell docker images -f dangling=true -q --no-trunc)
endif
