package main

import (
	"encoding/base64"
	"fmt"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Base64 Encoder/Decoder (go-gtk)")
	window.SetDefaultSize(600, 400)

	window.Connect("destroy", func() { gtk.MainQuit() })

	hBox := gtk.NewHBox(false, 5)
	window.Add(hBox)

	vBoxLeft := gtk.NewVBox(false, 5)
	hBox.PackStart(vBoxLeft, true, true, 5)

	textViewLeft := gtk.NewTextView()
	scrollLeft := gtk.NewScrolledWindow(nil, nil)
	scrollLeft.Add(textViewLeft)
	vBoxLeft.PackStart(scrollLeft, true, true, 5)

	encodeBtn := gtk.NewButtonWithLabel("Encode")
	vBoxLeft.PackStart(encodeBtn, false, false, 5)

	vBoxRight := gtk.NewVBox(false, 5)
	hBox.PackStart(vBoxRight, true, true, 5)

	textViewRight := gtk.NewTextView()
	scrollRight := gtk.NewScrolledWindow(nil, nil)
	scrollRight.Add(textViewRight)
	vBoxRight.PackStart(scrollRight, true, true, 5)

	decodeBtn := gtk.NewButtonWithLabel("Decode")
	vBoxRight.PackStart(decodeBtn, false, false, 5)

	encodeBtn.Clicked(func() {
		bufferLeft := textViewLeft.GetBuffer()

		startLeft := &gtk.TextIter{}
		endLeft := &gtk.TextIter{}
		bufferLeft.GetStartIter(startLeft)
		bufferLeft.GetEndIter(endLeft)

		leftText := bufferLeft.GetText(startLeft, endLeft, true)

		encoded := base64.StdEncoding.EncodeToString([]byte(leftText))

		bufferRight := textViewRight.GetBuffer()
		bufferRight.SetText(encoded)
	})

	decodeBtn.Clicked(func() {
		bufferRight := textViewRight.GetBuffer()

		startRight := &gtk.TextIter{}
		endRight := &gtk.TextIter{}
		bufferRight.GetStartIter(startRight)
		bufferRight.GetEndIter(endRight)

		rightText := bufferRight.GetText(startRight, endRight, true)

		decodedBytes, err := base64.StdEncoding.DecodeString(rightText)
		if err != nil {
			bufferLeft := textViewLeft.GetBuffer()
			bufferLeft.SetText(fmt.Sprintf("Ошибка декодирования: %v", err))
		} else {
			bufferLeft := textViewLeft.GetBuffer()
			bufferLeft.SetText(string(decodedBytes))
		}
	})

	window.ShowAll()
	gtk.Main()
}
