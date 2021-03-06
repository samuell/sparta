# Constants for cross compilation and packaging.
appID = com.github.jacalz.sparta
icon = internal/assets/icon-512.png
name = Sparta

# Default path to the go binary directory.
GOBIN ?= ~/go/bin/

bundle:
	# Bundle the correct logo into sparta/src/bundled/bundled.go
	${GOBIN}fyne bundle -package assets internal/assets/icon-256.png > internal/assets/bundled.go

	# Modify the variable name to be correct.
	sed -i 's/resourceIcon256Png/AppIcon/g' internal/assets/bundled.go

check:
	# Check the whole codebase for misspellings.
	${GOBIN}misspell -w .

	# Run full formating on the code.
	gofmt -s -w .

	# Check the whole program for security issues.
	${GOBIN}gosec ./...

darwin:
	${GOBIN}fyne-cross darwin -arch amd64 -app-id ${appID} -icon ${icon} -output ${name}

linux:
	${GOBIN}fyne-cross linux -arch amd64 -app-id ${appID} -icon ${icon}

windows:
	${GOBIN}fyne-cross windows -arch amd64 -app-id ${appID} -icon ${icon}

cross-compile: darwin linux windows