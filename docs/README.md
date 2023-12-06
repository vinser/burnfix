![appIcon](../burnfix.png)


**Try to fix screen image retentions and burn in, look for screen defects**

After turning off or moving the taskbar, you suddenly find that the logo in the corner of the screen and the search magnifier are still visible on your monitor or TV screen, although not as brightly.

This is the so-called **image retention** or **burn-in**.

Burnfix can attempt to get rid of this effect on your LCD, LED or plasma screen by showing a special moving Signal Pattern<sup>1</sup> for a period of time.

Moreover, burnfix can show full-screen images of various colors to help you identify defects in the display matrix. 

Enjoy!  

## Ready to use

For Linux users, Burnfix is avaliable as a Flatpak on [Flathub](https://flathub.org/apps/io.github.vinser.burnfix):

<a href='https://flathub.org/apps/details/io.github.vinser.burnfix'><img width='200' alt='Download on Flathub' src='https://dl.flathub.org/assets/badges/flathub-badge-en.svg'/></a>

Pre-built binaries are available for Linux (`x86-64`,`arm64` and `arm`), macOS (`x86-64`) and Windows (`x86-64`). Please visit the [release page](https://github.com/vinser/burnfix/releases) to download the latest release.

## Building

Burnfix compiles into a statically linked binary with no explicit runtime dependencies. 

Compiling requires a [Go](https://go.dev) compiler (v1.18 or later is required) and the [prerequisites for Fyne](https://developer.fyne.io/started/)<sup>2</sup>.

On systems with the above compile-time requirements met one can build the project using `go build` in the project root:
```
git clone https://github.com/vinser/burnfix.git
go build
```
---
<sup>1</sup> Similar Signal Pattern was used to remove after images on Samsung Plasma Display Panels and they clamed it to be more effective then All White signal.   
<sup>2</sup> Burnfix was created using [Fyne](https://fyne.io) toolkit.  

