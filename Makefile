target_dir := dist
target := $(target_dir)/anki-immersion-reader
src := main.go ankiConnect.go | $(target_dir)

all := $(target)_windows_x86_64.exe $(target)_windows_arm64.exe $(target)_darwin_x86_64 $(target)_darwin_arm64 $(target)_linux_x86_64 $(target)_linux_arm64

.PHONY: all clean

all: $(all)

clean:
	rm -f $(all)

$(target_dir):
	mkdir $(target_dir)

# Windows
$(target)_windows_x86_64.exe: $(src)
	GOOS=windows GOARCH=amd64 go build -ldflags "-s" -o $(target)_windows_x86_64.exe

$(target)_windows_arm64.exe: $(src)
	GOOS=windows GOARCH=arm64 go build -ldflags "-s" -o $(target)_windows_arm64.exe

# MacOS
$(target)_darwin_x86_64: $(src)
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s" -o $(target)_darwin_x86_64

$(target)_darwin_arm64: $(src)
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s" -o $(target)_darwin_arm64

# Linux
$(target)_linux_x86_64: $(src)
	GOOS=linux GOARCH=amd64 go build -ldflags "-s" -o $(target)_linux_x86_64

$(target)_linux_arm64: $(src)
	GOOS=linux GOARCH=arm64 go build -ldflags "-s" -o $(target)_linux_arm64
