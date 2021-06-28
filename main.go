package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const app_name string = "Mine Fyne GUI"

func app_setup() fyne.App {
	fmt.Printf("Launching \"%s\" ...\n", app_name)
	an_app := app.New()

	return an_app
}

func app_cleanup() {
	fmt.Println("Exiting...")
}

func show_another(a fyne.App) {
	time.Sleep(time.Second * 5)
	win := a.NewWindow("About")
	win.SetContent(widget.NewLabel("... 5 seconds later"))
	win.Resize(fyne.NewSize(320, 240))

	win.Show()

	time.Sleep(time.Second * 2)
	win.Close()
}

func main() {
	const output_label_prefix_str string = "Output: "

	main_app := app_setup()

	main_window := main_app.NewWindow(app_name)

	header_label := widget.NewLabel("* Welcome to my \"\" *")
	output_label := widget.NewLabel(output_label_prefix_str)
	footer_label := widget.NewLabel("* copyright 2021 - Gabri Botha / School of Death *")

	example_button := widget.NewButton("Test!", func() {
		output_label.SetText(output_label_prefix_str + "Showing about window...")
		show_another(main_app)
	})
	root_container := container.NewVBox(header_label, output_label, example_button, footer_label)
	main_window.SetContent(root_container)

	main_window.Show()
	main_app.Run()

	app_cleanup()
}
