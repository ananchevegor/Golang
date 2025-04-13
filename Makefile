ifeq ($(OS),Windows_NT)
    BINPATH=bin/myapp.exe
else
    BINPATH=bin/myapp
endif

.PHONY: build
build:
	templ generate 
	go build -o $(BINPATH) cmd/myapp/main.go

.PHONY: run
run: build build-css build-js
	$(BINPATH)

.PHONY: build-css
build-css:
	npm --prefix web run build:css

.PHONY: build-js
build-css:
	npm --prefix web run build:js

.PHONY: watch-app
watch-app:
	go run github.com/air-verse/air@latest \
	--build.cmd "$(MAKE) build" \
	--build.bin "/bin/myapp" \
	--build.include_ext "go" \
	--build.exclude_dir "bin"
	air

.PHONY: fmt
fmt:
	templ fmt internal/view
