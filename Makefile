TESTARGS:=-short
SUBDIRS := $(wildcard */.)

Gopkg.lock: Gopkg.toml
	dep ensure -v

.deps: Gopkg.lock
	dep ensure -v

build: .deps
	go build -o hello-skeleton.exe ./cmd/hello-skeleton


# The grep command allows the test target to stop upon error (which
# does not happen using the gocov command). The tee is to allow the
# developer to see the test results live.
test: .deps
	cd data && go test
	@grep "FAIL" test.out > /dev/null; \
	if [ $$? -eq 0 ] ; then \
		exit 1; \
	fi