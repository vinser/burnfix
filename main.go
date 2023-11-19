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
	buttonsRow := container.NewGridWithColumns(3,
		newButtonWithGradient("Clear retensions", color.Black, color.White, func() {
			go clearRetensions(a)
		}),
		newButtonWithColor("Defects search", color.NRGBA{R: 255, G: 0, B: 0, A: 255}, func() {
			go solidColor(a)
		}),
		newButtonWithColor("About", color.NRGBA{R: 0, G: 255, B: 0, A: 255}, func() {
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

func newButtonWithGradient(label string, startColor, endColor color.Color, tapped func()) *fyne.Container {
	g := canvas.NewHorizontalGradient(startColor, endColor)
	b := widget.NewButtonWithIcon(label, nil, tapped)
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
	msg := canvas.NewText(msgText, color.NRGBA{R: 0, G: 255, B: 0, A: 255})
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
	red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	green := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	blue := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	black := color.NRGBA{R: 0, G: 0, B: 0, A: 0}
	grey := color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	colors := []Pallete{
		{fg: white, bg: red},
		{fg: white, bg: green},
		{fg: white, bg: blue},
		{fg: black, bg: white},
		{fg: white, bg: black},
		{fg: white, bg: grey},
	}

	iColor := 0
	obj := canvas.NewRectangle(colors[iColor].bg)
	obj.Resize(fyne.NewSize(winW, winH))
	content.Add(obj)

	msgText := `Look for bad pixels... Press Space to change color or Esc to cancel`
	msg := canvas.NewText(msgText, colors[iColor].fg)
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
			if iColor >= len(colors) {
				iColor = 0
			}
			msg.Color = colors[iColor].fg
			obj.FillColor = colors[iColor].bg
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
	left := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 255})
	rich := widget.NewRichTextFromMarkdown(`
# burnfix

After turning off or moving the taskbar in Windows, you suddenly discover that the window symbol and the search magnifying glass are still visible on the screen of your monitor or TV, albeit not so brightly. This is the so-called **image retention** or **burn in**. 

Once upon a time, a similar effect appeared on plasma panels, but as it turned out, LCD and LED displays can be susceptible to it. In some cases, such changes are irreversible, especially when a static image is displayed continuously for many days. But you can try to get rid of this effect by showing a special pattern on the screen. This pattern was built into my old Samsung plasma TV more than 10 years ago. Therefore, having discovered this effect on my modern 4K IPS monitor, I decided to write a simple application that might help you too.  

In addition, using this application you can identify defects of display matrix by looking at solid images in different colors.

Enjoy!

-------------------------	
MIT License  

Copyright (c) 2023 vinser  

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is	furnished to do so, subject to the following conditions:  
	
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.  

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
`)
	rich.Wrapping = fyne.TextWrapWord
	right := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 255})

	left.Resize(fyne.NewSize(winW/4, winH))
	rich.Resize(fyne.NewSize(winW/4*2, winH))
	right.Resize(fyne.NewSize(winW/4, winH))

	left.Move(fyne.NewPos(0, 0))
	rich.Move(fyne.NewPos(winW/4, 0))
	right.Move(fyne.NewPos(winW/4*3, 0))

	content := container.NewWithoutLayout(left, rich, right)

	msgText := `Press Esc to close`
	msg := canvas.NewText(msgText, color.NRGBA{R: 0, G: 255, B: 0, A: 255})
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

// Application custom theme and interface inplementation
type Theme struct{}

func (t *Theme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch {
	case name == theme.ColorNameButton:
		return color.Transparent
		// case name == theme.ColorNameForeground:
		// 	return color.NRGBA{R: 100, G: 100, B: 100, A: 255}
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
	if name == theme.SizeNameInputRadius {
		return 0
	}
	return theme.DefaultTheme().Size(name)
}
