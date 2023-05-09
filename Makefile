# Colors used in this Makefile
escape=$(shell printf '\033')
RESET_COLOR=$(escape)[0m
COLOR_YELLOW=$(escape)[38;5;220m
COLOR_RED=$(escape)[91m
COLOR_BLUE=$(escape)[94m

COLOR_LEVEL_TRACE=$(escape)[38;5;87m
COLOR_LEVEL_DEBUG=$(escape)[38;5;87m
COLOR_LEVEL_INFO=$(escape)[92m
COLOR_LEVEL_WARN=$(escape)[38;5;208m
COLOR_LEVEL_ERROR=$(escape)[91m
COLOR_LEVEL_FATAL=$(escape)[91m

define COLORIZE
sed -u -e "s/\\\\\"/'/g; \
s/Method\([^ ]*\)/Method$(COLOR_BLUE)\1$(RESET_COLOR)/g;        \
s/ERROR\"\([^\"]*\)\"/error=\"$(COLOR_RED)\1$(RESET_COLOR)\"/g;  \
s/ProductID:\s\([^\"]*\)/$(COLOR_YELLOW)ProductID: \1$(RESET_COLOR)/g;   \
s/\[TRACE\]/$(COLOR_LEVEL_TRACE)\[TRACE\]$(RESET_COLOR)/g;    \
s/\[DEBUG\]/$(COLOR_LEVEL_DEBUG)DEBUG$(RESET_COLOR)/g;    \
s/\[INFO\]/$(COLOR_LEVEL_INFO)[INFO]$(RESET_COLOR)/g;       \
s/\[WARNING\]/$(COLOR_LEVEL_WARN)[WARNING]$(RESET_COLOR)/g; \
s/\[ERROR\]/$(COLOR_LEVEL_ERROR)[ERROR]$(RESET_COLOR)/g;    \
s/\[FATAL\]/level=$(COLOR_LEVEL_FATAL)[FATAL]$(RESET_COLOR)/g"
endef


#####################
# Help targets      #
#####################

.PHONY: help.highlevel help.all

#help help.highlevel: show help for high level targets. Use 'make help.all' to display all help messages
help.highlevel:
	@grep -hE '^[a-z_-]+:' $(MAKEFILE_LIST) | LANG=C sort -d | \
	awk 'BEGIN {FS = ":"}; {printf("$(COLOR_YELLOW)%-25s$(RESET_COLOR) %s\n", $$1, $$2)}'

#help help.all: display all targets' help messages
help.all:
	@grep -hE '^#help|^[a-z_-]+:' $(MAKEFILE_LIST) | sed "s/#help //g" | LANG=C sort -d | \
	awk 'BEGIN {FS = ":"}; {if ($$1 ~ /\./) printf("    $(COLOR_BLUE)%-21s$(RESET_COLOR) %s\n", $$1, $$2); else printf("$(COLOR_YELLOW)%-25s$(RESET_COLOR) %s\n", $$1, $$2)}'


#####################
# Build targets     #
#####################
.PHONY: build run

NAME=airzone
GIT_COMMIT=$(shell git rev-list -1 HEAD --abbrev-commit)

IMAGE_TAG=$(VERSION)-$(GIT_COMMIT)
IMAGE_NAME=$(NAME)

#help build.prepare: prepare target/ folder
build.prepare:
	@mkdir -p $(CURDIR)/target
	@rm -f $(CURDIR)/target/$(NAME)

#help build.vendor: retrieve all the dependencies used for the project
build.vendor:
	go mod vendor

#help build.vendor.full: retrieve all the dependencies after cleaning the go.sum
build.vendor.full:
	@rm -fr $(CURDIR)/vendor
	go mod vendor

#help build.docker: build a docker image
build.docker:
	podman build --build-arg build_args="$(BUILD_ARGS)"  -t $(IMAGE_NAME):$(IMAGE_TAG) -f Containerfile .

build.local:
	go build -o $(CURDIR)/target/hvac *.go

ZONEID=0
SYSTEMID=1
AIRZONE_URL="airzone:3000"
run:
	$(CURDIR)/target/hvac --url $(AIRZONE_URL) --system-id $(SYSTEMID) --zone-id $(ZONEID)
