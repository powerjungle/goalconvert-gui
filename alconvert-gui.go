package main

import (
	"fmt"
	"image/color"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/powerjungle/goalconvert/alconvert"
)

var Version = "v1.0.0"

func alcCalcAll(alcval *alconvert.Alcovalues) {
	alcval.CalcGotUnits()
	alcval.CalcTargetUnits()
	alcval.CalcTargetPercent()
	alcval.CalcTargetMl()
}

type allLabels struct {
	unitsLabel      *widget.Label
	finAmountLabel  *widget.Label
	finalRemAmLabel *widget.Label
	finalTarPerc    *widget.Label
	finalTarPercA   *widget.Label
	finalTarMlP     *widget.Label
	finalTarMlD     *widget.Label
	notSetWarn      *widget.Label
}

func initAllLabels() *allLabels {
	return &allLabels{
		unitsLabel:      widget.NewLabel("0"),
		finAmountLabel:  widget.NewLabel("0"),
		finalRemAmLabel: widget.NewLabel("0"),
		finalTarPerc:    widget.NewLabel("0"),
		finalTarPercA:   widget.NewLabel("0"),
		finalTarMlP:     widget.NewLabel("0"),
		finalTarMlD:     widget.NewLabel("0"),
		notSetWarn:      widget.NewLabel("Initial ml and percentage are not set in the \"Units\" section!"),
	}
}

func resetAllLabels(alcoval *alconvert.Alcovalues, aLa *allLabels) {
	aLa.unitsLabel.SetText(strconv.FormatFloat(float64(alcoval.GotUnits()), 'f', -1, 32))
	aLa.finAmountLabel.SetText(strconv.FormatFloat(float64(alcoval.GotTargUnitsFinalAmount()), 'f', -1, 32))
	aLa.finalRemAmLabel.SetText(strconv.FormatFloat(float64(alcoval.GotTargUnitsRemAmount()), 'f', -1, 32))
	aLa.finalTarPerc.SetText(strconv.FormatFloat(float64(alcoval.GotTargPercAddWater()), 'f', -1, 32))
	aLa.finalTarPercA.SetText(strconv.FormatFloat(float64(alcoval.GotTargPercAlcLeft()), 'f', -1, 32))
	aLa.finalTarMlP.SetText(strconv.FormatFloat(float64(alcoval.GotTargMlNewAlcPerc()), 'f', -1, 32))
	aLa.finalTarMlD.SetText(strconv.FormatFloat(float64(alcoval.GotTargMlNeededWater()), 'f', -1, 32))
	if alcoval.GotUnits() == 0 {
		aLa.notSetWarn.Show()
	} else {
		aLa.notSetWarn.Hide()
	}
}

type inputWidgets struct {
	amountLabel  *widget.Label
	amountSlider *widget.Slider
	amountEntry  *widget.Entry
}

func fromInputToAlcv(label string, alcoval *alconvert.Alcovalues, changeTo float64) {
	switch label {
	case "Milliliters":
		alcoval.UserSet.Milliliters = float32(changeTo)
	case "Percentage":
		alcoval.UserSet.Percent = float32(changeTo)
	case "Unit Target":
		alcoval.UserSet.UnitTarget = float32(changeTo)
	case "Percentage Target":
		alcoval.UserSet.PercenTarget = float32(changeTo)
	case "Milliliter Target":
		alcoval.UserSet.TargetMl = float32(changeTo)
	}
}

func newInputWidgets(label string, rangeMin float64, rangeMax float64,
	alcoval *alconvert.Alcovalues, aLa *allLabels) *inputWidgets {

	iw := inputWidgets{
		amountLabel:  widget.NewLabel(label),
		amountSlider: widget.NewSlider(rangeMin, rangeMax),
		amountEntry:  widget.NewEntry(),
	}

	iw.amountEntry.SetText("0")

	iw.amountSlider.Step = 1
	iw.amountSlider.OnChanged = func(slideVal float64) {
		iw.amountEntry.SetText(strconv.FormatFloat(slideVal, 'f', -1, 64))
		fromInputToAlcv(label, alcoval, slideVal)
		alcCalcAll(alcoval)
		resetAllLabels(alcoval, aLa)
	}

	iw.amountEntry.OnChanged = func(inputStr string) {
		if f, err := strconv.ParseFloat(inputStr, 32); err == nil {
			fromInputToAlcv(label, alcoval, f)
		} else if inputStr != "" {
			fmt.Println(err)
		}

		if f, err := strconv.ParseFloat(inputStr, 64); err == nil {
			iw.amountSlider.SetValue(f)
		} else if inputStr != "" {
			fmt.Println(err)
		}
	}

	return &iw
}

func makeIOCanvasObjects(iWs1 *inputWidgets, iWs2 *inputWidgets,
	calcValueLabel1 *widget.Label, calcValueLabel2 *widget.Label,
	firstResultLabel string, secondResultLabel string,
	warnLabel *widget.Label) *fyne.Container {

	co := []fyne.CanvasObject{}

	label1 := widget.NewLabel(firstResultLabel)
	label2 := widget.NewLabel(secondResultLabel)

	if fyne.CurrentDevice().IsMobile() {
		label1.Wrapping = fyne.TextWrapWord
		label2.Wrapping = fyne.TextWrapWord
	}

	if firstResultLabel == "nope" {
		label1.Hide()
	}

	if secondResultLabel == "nope" {
		label2.Hide()
	}

	cont1 := container.NewVBox()
	cont2 := container.NewVBox()

	if calcValueLabel1 != nil {
		calcValueLabel1.TextStyle.Bold = true
		cont1.Objects = append(cont1.Objects, label1, calcValueLabel1)
		co = append(co, cont1)
	}

	if calcValueLabel2 != nil {
		calcValueLabel2.TextStyle.Bold = true
		cont2.Objects = append(cont2.Objects, label2, calcValueLabel2)
		co = append(co, cont2)
	}

	co = append(co, canvas.NewLine(color.White))

	if iWs1 != nil {
		co = append(co, iWs1.amountLabel, iWs1.amountEntry, iWs1.amountSlider)
	}

	if iWs2 != nil {
		co = append(co, iWs2.amountLabel, iWs2.amountEntry, iWs2.amountSlider)
	}

	if warnLabel != nil {
		if fyne.CurrentDevice().IsMobile() {
			warnLabel.Wrapping = fyne.TextWrapWord
		}
		co = append(co, warnLabel)
	}

	return container.New(layout.NewVBoxLayout(), co...)
}

func makeMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	helpMenu := fyne.NewMenu("About",
		fyne.NewMenuItem("GitHub Page", func() {
			u, _ := url.Parse("https://github.com/powerjungle/goalconvert")
			_ = app.OpenURL(u)
		}))
	return fyne.NewMainMenu(
		helpMenu,
	)
}

func main() {
	av := alconvert.NewAV()

	alcApp := app.NewWithID("com.github.goalconvert")
	alcApp.SetIcon(resourceAlcGuIconPng)
	alcWindow := alcApp.NewWindow("goalconvert " + Version)
	alcWindow.Resize(fyne.NewSize(600, 0))

	alcWindow.SetMainMenu(makeMenu(alcApp, alcWindow))
	alcWindow.SetMaster()

	aLa := initAllLabels()

	// Init input widgets
	mlInputWidgets := newInputWidgets("Milliliters", 0, 2000, av, aLa)

	percInputWidgets := newInputWidgets("Percentage", 0, 100, av, aLa)

	unitTarInputWidgets := newInputWidgets("Unit Target", 0, 20, av, aLa)
	unitTarInputWidgets.amountSlider.Step = 0.1

	percTarInputWidgets := newInputWidgets("Percentage Target", 0, 100, av, aLa)

	mlTarInputWidgets := newInputWidgets("Milliliter Target", 0, 2000, av, aLa)
	///////////////////////

	// Init input CanvasObjects
	gotUnitsObjects := makeIOCanvasObjects(
		mlInputWidgets,
		percInputWidgets,
		aLa.unitsLabel,
		nil,
		"Calculated Units using ml and %.",
		"nope",
		nil,
	)

	gotUnitTargetObjects := makeIOCanvasObjects(
		unitTarInputWidgets,
		nil,
		aLa.finAmountLabel,
		aLa.finalRemAmLabel,
		"Amount in ml alcohol left to reach target.",
		"Amount in ml alcohol to be removed to reach target.",
		aLa.notSetWarn,
	)

	gotPercTargetObjects := makeIOCanvasObjects(
		percTarInputWidgets,
		nil,
		aLa.finalTarPerc,
		aLa.finalTarPercA,
		"Amount of water in ml to add, to reach the target.",
		"Amount of diluted alcohol left after adding water, to reach the target.",
		aLa.notSetWarn,
	)

	gotMlTargetObjects := makeIOCanvasObjects(
		mlTarInputWidgets,
		nil,
		aLa.finalTarMlP,
		aLa.finalTarMlD,
		"Alcohol becomes this percentage, after adding water, to reach the target.",
		"The amount of water that needs to be added, in order to reach the target.",
		aLa.notSetWarn,
	)
	///////////////////////

	tabs := container.NewAppTabs(
		container.NewTabItem("Units", gotUnitsObjects),
		container.NewTabItem("Target Units", gotUnitTargetObjects),
		container.NewTabItem("Percent Target", gotPercTargetObjects),
		container.NewTabItem("Ml Target", gotMlTargetObjects),
	)

	if fyne.CurrentDevice().IsMobile() {
		scrollCont := container.NewVScroll(tabs)
		alcWindow.SetContent(scrollCont)
	} else {
		alcWindow.SetContent(tabs)
	}

	alcWindow.ShowAndRun()
}
