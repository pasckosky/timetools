
ANDROID_TARGET=Pasckosky_TimezoneToolbox.apk
LINUX_TARGET=times


.PHONY: all android linux clean run


all: android linux

clean:
	rm -f ${ANDROID_TARGET} ${LINUX_TARGET}

android:
	fyne package --target android/arm64

linux:
	go build .

run:
	go run .
