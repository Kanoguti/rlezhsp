export CGO_ENABLED=1

ifeq ($(OS),Windows_NT)
SHELL:=powershell.exe
.SHELLFLAGS:=-Command
make-folder=if (-not (Test-Path -Path $(1))) { New-Item -ItemType Directory -Path $(1) }
copy-file=Copy-Item -Path $(1) -Destination $(2) -Force
OUTPUT_FILE=rlez.dll
else
make-folder=mkdir -p $(1)
copy-file=cp -f $(1) $(2)
UNAME_S:=$(shell uname -s)
OUTPUT_FILE=librlez.so
ifeq ($(UNAME_S),Darwin)
OUTPUT_FILE=librlez.dylib
endif
endif

build_opengl33:
	$(call make-folder,"./build/opengl33")
	go build -buildmode=c-shared -ldflags="-s -w" -o $(OUTPUT_FILE) .
	$(call copy-file,"./$(OUTPUT_FILE)","./build/opengl33/$(OUTPUT_FILE)")

build_angle:
	$(call make-folder,"./build/angle")
	go build -tags es2 -buildmode=c-shared -ldflags="-s -w -extldflags '-L. -lGLESv2 -lEGL -Wl,--allow-multiple-definition'" -o $(OUTPUT_FILE) .
	$(call copy-file,"./$(OUTPUT_FILE)","./build/angle/$(OUTPUT_FILE)")

build_all: build_opengl33 build_angle

debug:
	go run .

download_raylib:
	go clean -modcache
	go get -v -u github.com/gen2brain/raylib-go/raylib@v0.60.1
	go mod tidy