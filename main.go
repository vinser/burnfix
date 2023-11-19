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
	// w.SetPadded(false)
	buttonsRow := container.NewGridWithColumns(7,
		newButtonWithGradient("Clear", color.Black, color.White, func() {
			go clearRetensions(a)
		}),
		newButtonWithColor("", color.NRGBA{R: 255, G: 0, B: 0, A: 255}, func() {}),
		newButtonWithColor("", color.NRGBA{R: 0, G: 255, B: 0, A: 255}, func() {}),
		newButtonWithColor("", color.NRGBA{R: 0, G: 0, B: 255, A: 255}, func() {}),
		newButtonWithColor("", color.NRGBA{R: 255, G: 255, B: 255, A: 255}, func() {}),
		newButtonWithColor("", color.NRGBA{R: 0, G: 0, B: 0, A: 0}, func() {}),
		newButtonWithColor("", color.NRGBA{R: 44, G: 44, B: 44, A: 255}, func() {}),
	)
	w.SetContent(buttonsRow)
	w.Show()
	// keyCtrlEsc := &desktop.CustomShortcut{KeyName: fyne.KeyEscape, Modifier: fyne.KeyModifierControl}
	// w.Canvas().AddShortcut(keyCtrlEsc, func(shortcut fyne.Shortcut) {
	// 	w.SetFullScreen(false)
	// })
	// go clearRetensions(w)
	a.Run()
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
	w.SetFullScreen(true)
	w.SetPadded(false)
	w.Show()
	time.Sleep(time.Second)
	col := float32(2)
	winW := w.Canvas().Size().Width
	winH := w.Canvas().Size().Height
	colW := winW / col
	var lgs []fyne.CanvasObject
	for i := 0; i < 5; i++ {
		lgs = append(lgs, newWhiteGradient(winW/col, winH))
	}
	w.SetContent(container.NewWithoutLayout(lgs...))

	anmTime := 60 * time.Second
	var anm []*fyne.Animation
	for i := 0; i < 5; i++ {
		start := fyne.NewPos(colW*float32(i), 0)
		stop := fyne.NewPos(-winW+colW*float32(i), 0)
		anm = append(anm, canvas.NewPositionAnimation(start, stop, anmTime, lgs[i].Move))
		anm[i].Curve = fyne.AnimationLinear
		anm[i].Start()
	}
	// time.Sleep(anmTime)
	// w.SetFullScreen(false)
	w.Close()
}

func newWhiteGradient(winW, winH float32) fyne.CanvasObject {
	obj := canvas.NewHorizontalGradient(color.Black, color.White)
	obj.Resize(fyne.NewSize(winW, winH))
	return obj
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
