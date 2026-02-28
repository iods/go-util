#
# Makefile for `string` utility
#
PKGS := github.com/iods/go-util
SRCDIRS := $(shell go list -f '{{.Dir}}' $(PKGS))

string\:pkg:
	@echo "This is the string package."

packages:
	# @echo $(SRCDIRS)
