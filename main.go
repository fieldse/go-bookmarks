package main

import (
		"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
		"fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Bookmarks")
    myWindow.Resize(fyne.NewSize(440, 700))


    title := canvas.NewText("Bookmark App", theme.PrimaryColor())
    title.TextStyle = fyne.TextStyle{Bold: true}
    title.TextSize = 32

    // Horizontal row of buttons
    buttonRow := container.NewHBox(
        widget.NewButton("Example 1", func() {}),
        widget.NewButton("Example 2", func() {}),
    )

    // Centered vertically and horizontally
    centered := container.New(layout.NewCenterLayout(), buttonRow)
    filler := container.NewMax(centered)

    content := container.NewVBox(
        title,
        layout.NewSpacer(),
        filler,
        layout.NewSpacer(),
    )

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
