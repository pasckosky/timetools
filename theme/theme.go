package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// panelTheme is a simple demonstration of a bespoke theme loaded by a Fyne app.
type panelTheme struct {
}

func (panelTheme) Color(c fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return color.Gray{Y: 64}
	case theme.ColorNameButton, theme.ColorNameDisabled:
		return color.White
	case theme.ColorNamePlaceHolder, theme.ColorNameScrollBar:
		return color.RGBA{R: 255, G: 0, B: 0, A: 255}
	case theme.ColorNamePrimary, theme.ColorNameHover, theme.ColorNameFocus:
		return color.Gray{Y: 128}
	case theme.ColorNameShadow:
		return color.RGBA{R: 0xcc, G: 0xcc, B: 0xcc, A: 0xcc}
	default:
		return color.White
	}
}

func (panelTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DarkTheme().Font(style)
}

func (panelTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (panelTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNameScrollBar:
		return 10
	case theme.SizeNameScrollBarSmall:
		return 5
	case theme.SizeNameText:
		return 32
	case theme.SizeNameHeadingText:
		return 30
	case theme.SizeNameSubHeadingText:
		return 25
	case theme.SizeNameCaptionText:
		return 15
	case theme.SizeNameInputBorder:
		return 1
	default:
		return 0
	}
}

func PanelTheme() fyne.Theme {
	return &panelTheme{}
}
