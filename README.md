# Go Alcohol Converter (GUI)

<img src="AlcGuIcon.svg" alt="drawing" width="150"/>

Using this you can convert
[alcohol (drinkable)](https://en.wikipedia.org/wiki/Alcohol_(drug))
[millilitres](https://en.wikipedia.org/wiki/Litre#SI_prefixes_applied_to_the_litre),
[percentage](https://en.wikipedia.org/wiki/Alcohol_by_volume) and
[units](https://en.wikipedia.org/wiki/Unit_of_alcohol).

The units are the UK definition which you can checkout here:
https://www.nhs.uk/live-well/alcohol-advice/calculating-alcohol-units/

From the above link:
"One unit equals 10ml or 8g of pure alcohol, which is around the amount of
alcohol the average adult can process in an hour."

This project is not for the endorsement of alcohol consumption.
The aim is to help people get a perspective of what different amounts mean.

Stay safe!

## Dependencies

If you have an already built binary from the "Releases" page,
you don't need these.

For Go compilation/installation you'll need this.

You will need Go 1.17 or later.

Packages required by Fyne:
https://developer.fyne.io/started/#prerequisites

For installing or compiling on 32 bit or a different architecture or OS, you
need to install the appropriate packages, but it depends on the
operating system. You need to research for your own.
Some information can be found on the [Releases](#Releases) section.

## Installation

If you don't want or can't use the already built binaries
in the "Releases" page.

GUI: `go install github.com/powerjungle/goalconvert-gui@latest`

If you want to use as a module for your code.

Module: `go get github.com/powerjungle/goalconvert/alconvert`

## Compilation

`go build .`

Run the command where the `README.md` file is!

## Build for Android

You will need this:

https://developer.android.com/ndk/downloads

### Prepare for Android on Linux

After download, extract and put the folder in the home directory, then
add this to the end of your `.profile` file in your home directory:

`export ANDROID_NDK_HOME=/home/YOURusername/android-ndk-THEversion`

Change `YOURusername` and `THEversion` appropriately!

You'll need to logout and login again to your operating system.

### Building for Android

Install the fyne CLI utility: `go install fyne.io/fyne/v2/cmd/fyne@latest`

Run inside the `alconvgui` directory:

`fyne package -os android -appID testing.alconvert -icon AlcGuIcon.png`

This will create an [APK](https://en.wikipedia.org/wiki/Apk_(file_format)). 

### Installing on the phone

You'll need to use "Android Debug Bridge".

https://developer.android.com/studio/command-line/adb

#### Preparing ADB on Linux

Most distros should have an "Android tools" package.

OpenSUSE Tumbleweed:
`sudo zypper install android-tools`

Arch Linux:
`sudo pacman -S android-tools`

#### Installing on Android

You need to turn on "USB debugging" in your "Developer options" on your phone!

After connecting your phone to your PC with an USB cable, run:
`adb devices -l`

Then a window should show up on the phone asking you to allow access.

If you allow, then run:
`adb install alconvgui.apk`

## Generating icon.go

You'll need the Fyne CLI tool from [Building for Android](#building-for-android).

Run this command in the "alconvgui" directory:

`fyne bundle --output icon.go AlcGuIcon.png`

## Windows icons

When changing the icons, the appropriate .ico files need to be created.

Afterwards this tool needs to be installed:
`go install github.com/akavel/rsrc@latest` 

Then the old rsrc file needs to be replaced with the following command:

```
rsrc -ico AlcGuIcon-64x64.ico -arch amd64 && \
rsrc -ico AlcGuIcon-64x64.ico -arch 386
```

Change the .ico files to the appropriate names and change the command
in this readme if needed.

## Releases

The approprate packages need to be installed.
Checkout [Dependencies](#dependencies)!

To install Goreleaser: https://goreleaser.com/install/

To do a release: https://goreleaser.com/cmd/goreleaser/

The `.goreleaser.yaml` file is already done and is in the repo.

If you don't want to release for all OSs and architecture or want for more,
edit the `.goreleaser.yaml` file! Info on how to edit here:

https://goreleaser.com/customization/build/

All Go OS and architecture combos here:

https://go.dev/doc/install/source#environment

This needs to be tested every time it's changed, as not all builds work
without some preparation.

### Dependencies for 32 bit builds (GUI)

OpenSUSE Tumbleweed:

`sudo zypper install gcc-32bit glibc-32bit glibc-devel-32bit
libXcursor-devel-32bit libXrandr-devel-32bit Mesa-libGL-devel-32bit
libXi-devel-32bit libXinerama-devel-32bit libXxf86vm-devel-32bit`

### Dependencies for Windows builds (GUI)

OpenSUSE Tumbleweed (64 bit build):
`sudo zypper install mingw64-cross-gcc`

OpenSUSE Tumbleweed (32 bit build):
`sudo zypper install mingw32-cross-gcc`

## Tested running GUI binaries

The GUI release binaries have been tested on
Windows 10, OpenSUSE Tumbleweed, Fedora 35 and
should be working on similar distros.

### Running GUI on Debian/Ubuntu and similar

It's possible that the libraries are too old to run in such distros.

The "Fyne" module can't be completely statically linked:
https://github.com/fyne-io/fyne/issues/512

### Running GUI on Windows without proper OpenGL drivers

For example when running in a Windows virtual machine, without
the proper OpenGL support.

You'll need to install this:

https://www.microsoft.com/en-us/p/opencl-and-opengl-compatibility-pack/9nqpsl29bfff#activetab=pivot:overviewtab

## Image

Thanks to https://github.com/P3600 for the image!
