package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	gt "github.com/bas24/googletranslatefree"
	"golang.design/x/clipboard"
)

func MainForm() ui.Control {
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	group := ui.NewGroup("")
	group.SetMargined(false)
	box.Append(group, false)

	form := ui.NewForm()
	form.SetPadded(false)
	group.SetChild(form)

	emp := ui.NewLabel("")
	form.Append("", emp, false)

	input := ui.NewEntry()
	form.Append("Input:", input, false)

	lang := ui.NewCombobox()
	lang.Append("Auto")
	lang.Append("English")
	lang.Append("Chinese(Simplified)")
	lang.Append("Chinese(Traditional)")
	lang.Append("Malay")
	lang.Append("Indonesia")
	lang.Append("Sanskrit")
	lang.Append("Spanish")
	lang.Append("Japanese")
	lang.Append("German")
	lang.Append("Russian")
	lang.SetSelected(0)
	form.Append("Translate From: ", lang, true)

	langto := ui.NewCombobox()
	langto.Append("English")
	langto.Append("Chinese(Simplified)")
	langto.Append("Chinese(Traditional)")
	langto.Append("Malay")
	langto.Append("Indonesia")
	langto.Append("Sanskrit")
	langto.Append("Spanish")
	langto.Append("Japanese")
	langto.Append("German")
	langto.Append("Russian")
	langto.SetSelected(0)
	form.Append("Translate To: ", langto, true)

	output := ui.NewEntry()
	output.SetReadOnly(true)
	form.Append("Output: ", output, false)

	proceed := ui.NewButton("Proceed")
	box.Append(proceed, false)

	proceed.OnClicked(func(button *ui.Button) {
		from := ""
		to := ""
		langfromselec := lang.Selected()
		langtoselec := langto.Selected()

		switch langfromselec {
		case 0:
			from = "Auto"
		case 1:
			from = "en_US"
		case 2:
			from = "zh_CN"
		case 3:
			from = "zh_TW"
		case 4:
			from = "ms"
		case 5:
			from = "id"
		case 6:
			from = "sa"
		case 7:
			from = "es"
		case 8:
			from = "ja"
		case 9:
			from = "de"
		case 10:
			from = "ru"
		}

		switch langtoselec {
		case 0:
			to = "en_US"
		case 1:
			to = "zh_CN"
		case 2:
			to = "zh_TW"
		case 3:
			to = "ms"
		case 4:
			to = "id"
		case 5:
			to = "sa"
		case 6:
			to = "es"
		case 7:
			to = "ja"
		case 8:
			to = "de"
		case 9:
			to = "ru"

		}

		inputtxt := input.Text()
		result, _ := gt.Translate(inputtxt, from, to)

		output.SetText(result)
	})

	clip := ui.NewButton("Grab From Clipboard")
	box.Append(clip, false)
	clip.OnClicked(func(button *ui.Button) {
		clipb := clipboard.Read(clipboard.FmtText)
		input.SetText(string(clipb))
	})

	clrall := ui.NewButton("Clear All")
	box.Append(clrall, false)

	clrall.OnClicked(func(button *ui.Button) {
		input.SetText("")
		output.SetText("")
		lang.SetSelected(0)
		langto.SetSelected(0)
	})

	return box
}

func AboutForm() ui.Control {
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	var Empty = ui.NewLabel("")
	box.Append(Empty, false)

	var Name = ui.NewLabel("Name: Translator")
	box.Append(Name, false)

	var Author = ui.NewLabel("Author: MoQi")
	box.Append(Author, false)

	var Version = ui.NewLabel("Version: v1.0.0")
	box.Append(Version, false)

	return box
}

func Translator() {
	w := ui.NewWindow("Translator", 420, 550, false)
	w.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		w.Destroy()
		return true
	})

	tabs := ui.NewTab()
	tabs.Append("Translate", MainForm())
	tabs.SetMargined(0, true)
	tabs.Append("About", AboutForm())
	tabs.SetMargined(0, true)
	w.SetChild(tabs)
	w.SetMargined(true)

	w.Show()
}

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	ui.Main(Translator)
}
