# VERSION ?= $(shell git describe --long --tags)
VERSION ?= "0.0"
GOURL ?= github.com/jspreddy/go-algo


##############################################################################
# Make file tutorials: https://makefiletutorial.com/#getting-started
##############################################################################
# Default shell that make will use.
SHELL=/bin/bash

##############################################################################
# Terminal Colors:
# to see all colors, run
# bash -c 'for c in {0..255}; do tput setaf $c; tput setaf $c | cat -v; echo =$c; done'
# the first 15 entries are the 8-bit colors

# define standard colors
ifneq (,$(findstring xterm,${TERM}))
	BOLD         := $(shell tput -Txterm bold)
	UNDERLINE    := $(shell tput -Txterm smul)
	STANDOUT     := $(shell tput -Txterm smso)
	BLACK        := $(shell tput -Txterm setaf 0)
	RED          := $(shell tput -Txterm setaf 1)
	GREEN        := $(shell tput -Txterm setaf 2)
	YELLOW       := $(shell tput -Txterm setaf 3)
	BLUE         := $(shell tput -Txterm setaf 4)
	PURPLE       := $(shell tput -Txterm setaf 5)
	CYAN         := $(shell tput -Txterm setaf 6)
	WHITE        := $(shell tput -Txterm setaf 7)
	NORMAL := $(shell tput -Txterm sgr0)
else
	BOLD         := ""
	UNDERLINE    := ""
	STANDOUT     := ""
	BLACK        := ""
	RED          := ""
	GREEN        := ""
	YELLOW       := ""
	BLUE         := ""
	PURPLE       := ""
	CYAN         := ""
	WHITE        := ""
	NORMAL       := ""
endif

##############################################################################
# Makefile TARGETS:
##############################################################################
.PHONY: build install help helper-line clean
.DEFAULT_GOAL := help


##############################################################################
# Auto document targets by adding comment next to target name
# starting with double pound and space like "## doc text"
##############################################################################
helper-line:
	@echo "${BOLD}${BLUE}--------------------------------------------------${NORMAL}"

help: ## Show help documentation.
	@make -s helper-line
	@echo "${BOLD}${BLUE}Here are all the targets available with make command.${NORMAL}"
	@make -s helper-line
	@echo ""
	@echo "Start by running ${YELLOW}'make install'${NORMAL}"
	@echo ""
	@grep -E '^[a-zA-Z_0-9%-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    ${YELLOW}%-30s${NORMAL} %s\n", $$1, $$2}'

##############################################################################



build: ## Will build the go code and output binary to project root folder.
	go build $(GOURL)

install: build ## Will build and install binary into the go bin path.
	go install $(GOURL)

clean: ## Will clean the go outputs
	go clean $(GOURL)