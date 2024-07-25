package main

import (
	"context"
	"fmt"
	"github.com/percyjw-2/tinygo-keyboard/keycodes/jp"
	"image/color"
	"machine"
	"machine/usb"
	"tinygo.org/x/tinydraw"

	"tinygo.org/x/drivers/mcp23017"
	"tinygo.org/x/drivers/ssd1306"

	keyboard "github.com/percyjw-2/tinygo-keyboard"
)

type RCS struct {
	row, col int
	state    keyboard.State
}

func main() {
	usb.Product = "popcorne"

	err := machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
		SCL:       machine.P1_13,
		SDA:       machine.P0_29,
	})

	if err != nil {
		fmt.Printf("Failed to configure machine: %s\n", err)
	}

	ch := make(chan RCS, 16)

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Width:  128,
		Height: 32,
	})
	display.ClearDisplay()

	expander, _ := mcp23017.NewI2C(machine.I2C0, 0b0100000)

	d := keyboard.New()

	colPinsLeft := []machine.Pin{
		machine.P0_13,
		machine.P0_14,
		machine.P0_15,
		machine.P0_16,
		machine.P0_24,
		machine.P0_25,
	}

	rowPinsLeft := []machine.Pin{
		machine.P0_03,
		machine.P0_28,
		machine.P0_02,
		machine.P0_31,
	}

	colPinsRight := []mcp23017.Pin{
		expander.Pin(0),
		expander.Pin(1),
		expander.Pin(2),
		expander.Pin(3),
		expander.Pin(4),
		expander.Pin(5),
	}

	rowPinsRight := []mcp23017.Pin{
		expander.Pin(8),
		expander.Pin(9),
		expander.Pin(10),
		expander.Pin(11),
	}

	mkLeft := d.AddMatrixKeyboard(colPinsLeft, rowPinsLeft, [][]keyboard.Keycode{
		{
			jp.KeyTab, jp.KeyQ, jp.KeyW, jp.KeyE, jp.KeyR, jp.KeyT,
			jp.KeyLeftCtrl, jp.KeyA, jp.KeyS, jp.KeyD, jp.KeyF, jp.KeyG,
			jp.KeyLeftShift, jp.KeyY, jp.KeyX, jp.KeyC, jp.KeyV, jp.KeyB,
			0, jp.KeyMod1, jp.KeyMod2, jp.KeyWindows, jp.KeySpace, 0,
		},
		{
			jp.KeyEsc, jp.Key1, jp.Key2, jp.Key3, jp.Key4, jp.Key5,
			0, 0, jp.KeyUp, 0, 0, 0,
			0, jp.KeyLeft, jp.KeyDown, jp.KeyRight, 0, 0,
			0, 0, 0, 0, 0, 0,
		},
		{
			jp.KeyF1, jp.KeyF2, jp.KeyF3, jp.KeyF4, jp.KeyF5, jp.KeyF6,
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
		},
	}, keyboard.InvertDiode(true))
	mkLeft.SetCallback(func(layer, index int, state keyboard.State) {
		row := index / len(colPinsLeft)
		col := index % len(colPinsLeft)
		fmt.Printf("mkLeft: %d %d %d %d\n", layer, row, col, state)
		select {
		case ch <- RCS{row: row, col: col, state: state}:
		}
	})

	mkRight := d.Add

	go func() {
		for {
			select {
			case x := <-ch:
				c := color.RGBA{R: 255, G: 255, B: 255, A: 255}
				if x.state == keyboard.PressToRelease {
					c = color.RGBA{A: 255}
				}
				_ = tinydraw.FilledRectangle(&display, 10+20*int16(x.col), 10+20*int16(x.row), 18, 18, c)
				_ = display.Display()
			}
		}
	}()

	loadKeyboardDef()

	d.Debug = true
	err = d.Loop(context.Background())

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
