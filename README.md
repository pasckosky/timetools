# Timetools

Simple app for my needings in working with
USA, India, Australia, Italy, UTC

Built with fyne in go, it works for linux and
Android (not providing the apk)

# Linux

## prerequisites

On Fedora, please provides the following dependencies:

+ libglvnd-devel
+ libX11-devel
+ libXcursor-devel
+ libXrandr-devel
+ libXinerama-devel
+ libXi-devel
+ libXxf86vm-devel

More can be required, this should be the ones beyond the standard development-tool 

## how to build

Classinc mode

```sh
go build .
```

Makefile mode

```sh
make linux
```

## direct install

```sh
go install github.com/pasckosky/timetools@latest
```

# Android

## prerequisites

Needing two environment variables

```sh
export ANDROID_NDK_HOME=/home/user/Android/Sdk/ndk/25.2.9519653
export ANDROID_HOME=/home/user/Android
```

Refer to Android studio and fyne project for more details

## how to build

```sh
make android
```

## how to install

```sh
adb install Timezones_toolbox.apk
```
