PARAMS = $(filter-out $@,$(MAKECMDGOALS))

# ignore `No rule to make target errors`
%:
	@echo ""

# Install development tools
tools:
	@go get -u github.com/golang/dep/cmd/dep
.PHONY: tools

# Install vendor dependencies
deps.install:
	@dep ensure
.PHONY: deps.install

# Update vendor dependencies
deps.update:
	@dep ensure -update ${PARAMS}
.PHONY: deps.update

# Install aws.env cmd
install:
	@go install .
.PHONY: install
