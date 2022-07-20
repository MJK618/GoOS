package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
	"strings"
)

func showCalculatorApp() {
	//fmt.Print("GoKalkulator")
	//a := app.New()
	//w := a.NewWindow("GoKalkulator")

	output := ""
	input := widget.NewLabel(output)
	logString := ""
	prevLogEntry := widget.NewLabel(logString)
	isLog := false

	var logsArray []string
	logsButton := widget.NewButton("Logs", func() {
		if isLog {
			logString = ""
		} else {
			logString = strings.Join(logsArray, "\n")
		}
		isLog = !isLog
		prevLogEntry.SetText(logString)
	})

	backspaceButton := widget.NewButton("Backspace", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})

	CButton := widget.NewButton("©", func() {
		output = ""
		input.SetText(output)
	})

	leftBracketButton := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})

	rightBracketButton := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})

	divideButton := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})

	sevenButton := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})

	eigthButton := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})

	nineButton := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})

	multiplyButton := widget.NewButton("x", func() {
		output = output + "*"
		input.SetText(output)
	})

	fourButton := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})

	fiveButton := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})

	sixButton := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})

	minusButton := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})

	oneButton := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})

	twoButton := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})

	threeButton := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})

	plusButton := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})

	zeroButton := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})

	decimalButton := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})

	equalsButton := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				answer := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				temp := output + " = " + answer
				logsArray = append(logsArray, temp)
				output = answer

			} else {
				output = "Woah Dude Relax! Type Valid Expression"
			}
		} else {
			output = "Woah Dude Relax! Type Valid Expression"

		}
		input.SetText(output)
	})

	calculatorContainer := container.NewVBox(input,
			prevLogEntry,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2, logsButton, backspaceButton),
				container.NewGridWithColumns(4, CButton, leftBracketButton, rightBracketButton, divideButton),
				container.NewGridWithColumns(4, sevenButton, eigthButton, nineButton, multiplyButton),
				container.NewGridWithColumns(4, fourButton, fiveButton, sixButton, minusButton),
				container.NewGridWithColumns(4, oneButton, twoButton, threeButton, plusButton),
				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2, zeroButton, decimalButton), equalsButton),
			))

	w := myOS.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500,400))
	w.SetContent(container.NewBorder(nil,nil,nil,nil,calculatorContainer))
	w.Show()

}
