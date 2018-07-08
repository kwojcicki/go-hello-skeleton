Gopkg.lock: Gopkg.toml
	dep ensure -v

.deps: Gopkg.lock
	dep ensure -v

build: .deps
	go build -o hello-skeleton.exe ./cmd/hello-skeleton