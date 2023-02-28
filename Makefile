PKGS = $(shell go list ./...)

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: mod fmt vet check ; @
	$Q echo "done"

.PHONY: mod
mod: ; $(info $(M) collecting modules…) @
	$Q go mod download

.PHONY: fmt
fmt: ; $(info $(M) running go fmt…) @
	$Q go fmt $(PKGS)

.PHONY: vet
vet: ; $(info $(M) running go vet…) @
	$Q go vet $(PKGS) ; exit 0

.PHONY: check
check: ; $(info $(M) running linter…) @
	$Q go run honnef.co/go/tools/cmd/staticcheck@latest $(PKGS)
