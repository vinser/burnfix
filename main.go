package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&Theme{})
	w := a.NewWindow("Burn Fix")
	w.CenterOnScreen()
	w.SetPadded(false)
	w.SetMaster()
	buttonsRow := container.NewGridWithRows(3,
		newButtonWithGradient("Clear retensions", colorOf(black), colorOf(white), func() {
			go clearRetensions(a)
		}),
		newButtonWithColor("Defects search", colorOf(red), func() {
			go solidColor(a)
		}),
		newButtonWithColor("About", colorOf(green), func() {
			go about(a)
		}),
	)
	w.SetContent(buttonsRow)
	w.SetFixedSize(true)
	w.Show()

	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyEscape:
			w.Close()
		}
	})
	a.Run()
}

type Color int

const (
	white Color = iota
	black
	red
	green
	blue
	grey
)

func colorOf(c Color) color.Color {
	switch c {
	case white:
		return color.White
	case black:
		return color.Black
	case red:
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	case green:
		return color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	case blue:
		return color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	case grey:
		return color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	}
	return color.White
}

func newButtonWithGradient(label string, startColor, endColor color.Color, tapped func()) *fyne.Container {
	g := canvas.NewHorizontalGradient(startColor, endColor)
	b := widget.NewButton(label, tapped)
	return container.NewStack(g, b)
}
func newButtonWithColor(label string, color color.Color, tapped func()) *fyne.Container {
	r := canvas.NewRectangle(color)
	r.CornerRadius = theme.InputRadiusSize()
	b := widget.NewButton(label, tapped)
	return container.NewStack(r, b)
}

func clearRetensions(a fyne.App) {
	w := a.NewWindow("")
	defer w.Close()
	w.SetFullScreen(true)
	w.SetPadded(false)
	w.Show()
	time.Sleep(time.Second)
	content := container.NewWithoutLayout()
	col := 2
	winW := w.Canvas().Size().Width
	winH := w.Canvas().Size().Height
	colW := winW / float32(col)
	var lgs []fyne.CanvasObject
	for i := 0; i < col+2; i++ {
		obj := canvas.NewHorizontalGradient(color.Black, color.White)
		obj.Resize(fyne.NewSize(winW/2, winH))
		lgs = append(lgs, obj)
		content.Add(lgs[i])
	}

	var colAnm []*fyne.Animation
	for i := 0; i < col+2; i++ {
		start := fyne.NewPos(colW*float32(i), 0)
		stop := fyne.NewPos(-winW+colW*float32(i), 0)
		colAnm = append(colAnm, canvas.NewPositionAnimation(start, stop, 30*time.Second, lgs[i].Move))
		colAnm[i].Curve = fyne.AnimationLinear
		colAnm[i].RepeatCount = fyne.AnimationRepeatForever
		colAnm[i].Start()
	}
	msgText := `Clearing Retentions... Press Esc to cancel`
	msg := canvas.NewText(msgText, colorOf(green))
	msg.TextSize = fyne.CurrentApp().Settings().Theme().Size("text") * 3
	content.Add(msg)
	msgAnm := canvas.NewPositionAnimation(fyne.NewPos(winW, winH/2), fyne.NewPos(-msg.MinSize().Width, winH/2), 10*time.Second, msg.Move)
	msgAnm.Curve = fyne.AnimationEaseOut
	msgAnm.RepeatCount = 1
	msgAnm.Start()
	w.SetContent(content)

	cancelCh := make(chan struct{})
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if ev.Name == fyne.KeyEscape {
			cancelCh <- struct{}{}
		}
	})
	<-cancelCh
}

type Pallete struct {
	fg color.Color
	bg color.Color
}

func solidColor(a fyne.App) {
	w := a.NewWindow("")
	defer w.Close()
	w.SetFullScreen(true)
	w.SetPadded(false)
	w.Show()
	time.Sleep(time.Second)
	content := container.NewWithoutLayout()
	winW := w.Canvas().Size().Width
	winH := w.Canvas().Size().Height
	palette := []Pallete{
		{fg: colorOf(white), bg: colorOf(red)},
		{fg: colorOf(white), bg: colorOf(green)},
		{fg: colorOf(white), bg: colorOf(blue)},
		{fg: colorOf(black), bg: colorOf(white)},
		{fg: colorOf(white), bg: colorOf(black)},
		{fg: colorOf(white), bg: colorOf(grey)},
	}

	iColor := 0
	obj := canvas.NewRectangle(palette[iColor].bg)
	obj.Resize(fyne.NewSize(winW, winH))
	content.Add(obj)

	msgText := `Look for bad pixels... Press Space to change color or Esc to cancel`
	msg := canvas.NewText(msgText, palette[iColor].fg)
	msg.TextSize = fyne.CurrentApp().Settings().Theme().Size("text") * 3
	content.Add(msg)
	msgAnm := canvas.NewPositionAnimation(fyne.NewPos(winW, winH/2), fyne.NewPos(-msg.MinSize().Width, winH/2), 10*time.Second, msg.Move)
	msgAnm.Curve = fyne.AnimationEaseOut
	msgAnm.RepeatCount = 1
	msgAnm.Start()
	w.SetContent(content)

	cancelCh := make(chan struct{})
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeySpace:
			iColor++
			if iColor >= len(palette) {
				iColor = 0
			}
			msg.Color = palette[iColor].fg
			obj.FillColor = palette[iColor].bg
			canvas.Refresh(obj)
		case fyne.KeyEscape:
			cancelCh <- struct{}{}
		}
	})
	<-cancelCh
}

func about(a fyne.App) {
	w := a.NewWindow("")
	defer w.Close()
	w.SetFullScreen(true)
	w.SetPadded(false)
	w.Show()
	time.Sleep(time.Second)
	winW := w.Canvas().Size().Width
	winH := w.Canvas().Size().Height
	logo := canvas.NewImageFromResource(appIcon)
	logo.FillMode = canvas.ImageFillOriginal
	logoRow := container.NewVBox(widget.NewLabel(""), container.NewGridWithColumns(10, logo))

	aboutRow := widget.NewRichTextFromMarkdown(`
# burnfix

After turning off or moving the taskbar in Windows, you suddenly discover that the window corner logo
and the search magnifying glass are still visible on the screen of your monitor or TV, albeit not so brightly.

This is the so-called **image retention** or **burn in**.  

Once a similar effect was manifested on plasma panels, but as it turned out, LCD and LED displays can also be affected by it.
In some cases, such changes are irreversible, especially when a static image is displayed continuously for many days.
But you can try to get rid of this effect by showing a special pattern on the screen.
This pattern was built into my old Samsung plasma TV more than 10 years ago.
Therefore, having discovered this effect on my modern 4K IPS monitor, I decided to code this simple application that might help you too.  

In addition, using this application you can identify defects of display matrix by looking at solid images in different colors.  

Enjoy!
`)
	aboutRow.Wrapping = fyne.TextWrapWord

	tributeRow := widget.NewRichTextFromMarkdown(`
---

Powered by [fyne.io](https://fyne.io/) GUI toolkit    
`)
	tributeRow.Wrapping = fyne.TextWrapWord

	licenseRow := widget.NewRichTextFromMarkdown(`

---

MIT License

Copyright (c) 2023 vinser

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`)

	licenseRow.Wrapping = fyne.TextWrapWord

	left := canvas.NewRectangle(colorOf(black))
	center := container.NewVBox(logoRow, aboutRow, tributeRow, licenseRow)
	right := canvas.NewRectangle(colorOf(black))

	left.Resize(fyne.NewSize(winW/4, winH))
	center.Resize(fyne.NewSize(winW/4*2, winH))
	right.Resize(fyne.NewSize(winW/4, winH))

	left.Move(fyne.NewPos(0, 0))
	center.Move(fyne.NewPos(winW/4, 0))
	right.Move(fyne.NewPos(winW/4*3, 0))
	content := container.NewWithoutLayout(left, center, right)

	msgText := `Press Esc to close`
	msg := canvas.NewText(msgText, colorOf(green))
	msg.TextSize = fyne.CurrentApp().Settings().Theme().Size("text") * 3
	content.Add(msg)
	msgAnm := canvas.NewPositionAnimation(fyne.NewPos(winW, winH/2), fyne.NewPos(-msg.MinSize().Width, winH/2), 10*time.Second, msg.Move)
	msgAnm.Curve = fyne.AnimationEaseOut
	msgAnm.RepeatCount = 1
	msgAnm.Start()
	w.SetContent(content)

	cancelCh := make(chan struct{})
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyEscape:
			cancelCh <- struct{}{}
		}
	})
	<-cancelCh
}

var appIcon = &fyne.StaticResource{
	StaticName: "burnfix.svg",
	StaticContent: []byte(`
	<svg width="96" height="96" xmlns="http://www.w3.org/2000/svg">
 <g id="Layer_1">
  <title>Layer 1</title>
  <rect id="svg_1" height="28" width="84" y="6"  x="6" stroke-width="0" stroke="#000" fill="#e5e5e5"/>
  <rect id="svg_2" height="28" width="84" y="34" x="6" stroke-width="0" stroke="#000" fill="#ff0000"/>
  <rect id="svg_3" height="28" width="84" y="62" x="6" stroke-width="0" stroke="#000" fill="#00ff00"/>
 </g>
</svg>
`),
}

// Application custom theme and interface inplementation
type Theme struct{}

func (t *Theme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch {
	case name == theme.ColorNameButton:
		return color.Transparent
		// case name == theme.ColorNameForeground:
		// 	return colorOf(red)
	}
	return theme.DefaultTheme().Color(name, theme.VariantDark)
}

func (t *Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *Theme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding, theme.SizeNameInputRadius:
		return 0
	}
	return theme.DefaultTheme().Size(name) * 1.5
}
