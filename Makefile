APP = erlcmd
VERSION = $(shell cat VERSION)
ARCH = $(shell uname -m)
OS = $(shell uname -s|tr '[:upper:]' '[:lower:]')

DVCS_HOST = github.com
ORG = geomyidia
PROJ = $(APP)
FQ_PROJ = $(DVCS_HOST)/$(ORG)/$(PROJ)

GO_VERSION_STRING = $(shell go version)
GO_VERSION = $(strip $(subst go, , $(word 3, $(GO_VERSION_STRING))))
GO_ARCH = $(strip $(word 4, $(GO_VERSION_STRING)))
LD_VERSION = -X $(FQ_PROJ)/pkg/version.version=$(VERSION)
LD_BUILDDATE = -X $(FQ_PROJ)/pkg/version.buildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LD_GITCOMMIT = -X $(FQ_PROJ)/pkg/version.gitCommit=$(shell git rev-parse --short HEAD)
LD_GITBRANCH = -X $(FQ_PROJ)/pkg/version.gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
LD_GITSUMMARY = -X $(FQ_PROJ)/pkg/version.gitSummary=$(shell git describe --tags --dirty --always)
LD_GO_VERSION = -X $(FQ_PROJ)/pkg/version.goVersion=$(GO_VERSION)
LD_GO_ARCH = -X $(FQ_PROJ)/pkg/version.goArch=$(GO_ARCH)

LDFLAGS = -w -s $(LD_VERSION) $(LD_BUILDDATE) $(LD_GITBRANCH) $(LD_GITSUMMARY) $(LD_GITCOMMIT) $(LD_GO_VERSION) $(LD_GO_ARCH)

default: test

goversion:
	@echo $(GO_VERSION)

test:
	@echo ">> Running unit tests ..."
	@export PATH=$$PATH:~/go/bin && \
	richgo test -race -v ./... || echo "Uh-oh ... 🔥"

.PHONY: test
