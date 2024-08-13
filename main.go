package main

import (
	"context"
	"fmt"
	"github.com/percyjw-2/tinygo-keyboard/keycodes/jp"
	"github.com/percyjw-2/tinygo-keyboard/rgbanimations"
	"image/color"
	"machine"
	"machine/usb"
	"time"
	"tinygo.org/x/drivers/mcp23017"
	"tinygo.org/x/drivers/ws2812"
	"tinygo.org/x/tinydraw"

	"tinygo.org/x/drivers/ssd1306"

	keyboard "github.com/percyjw-2/tinygo-keyboard"
)

type RCS struct {
	row, col int
	state    keyboard.State
}

func main() {
	usb.Product = "popcorne"

	time.Sleep(3 * time.Second)

	err := machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
		SCL:       machine.P0_29,
		SDA:       machine.P1_13,
	})

	if err != nil {
		fmt.Printf("Failed to configure machine: %s\n", err)
	}

	ch := make(chan RCS, 16)

	expander, err := mcp23017.NewI2C(machine.I2C0, 0x20)

	if err != nil {
		fmt.Printf("Failed to create I2C device IO Expander: %s\n", err)
	}

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Width:   128,
		Height:  32,
		Address: ssd1306.Address_128_32,
	})
	display.ClearDisplay()

	ledPin := machine.P1_12
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led := ws2812.NewWS2812(ledPin)

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

	colPinsRight := []int{
		0,
		1,
		2,
		3,
		4,
		5,
	}

	rowPinsRight := []int{
		8,
		9,
		10,
		11,
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
		fmt.Printf("left row: %d, col: %d state: %d\n", row, col, state)
		select {
		case ch <- RCS{row: row, col: col, state: state}:
		}
	})

	mkRight, err := d.AddExpanderKeyboard(expander, colPinsRight, rowPinsRight, [][]keyboard.Keycode{
		{
			jp.KeyZ, jp.KeyU, jp.KeyI, jp.KeyO, jp.KeyP, jp.KeyMinus,
			jp.KeyH, jp.KeyJ, jp.KeyK, jp.KeyL, jp.KeySemicolon, jp.KeyColon,
			jp.KeyN, jp.KeyM, jp.KeyComma, jp.KeyPeriod, jp.KeySlash, jp.KeyRightShift,
			jp.KeyEnter, jp.KeyMod3, jp.KeyBackspace, jp.KeyDelete, 0, 0,
		},
		{
			jp.Key6, jp.Key7, jp.Key8, jp.Key9, jp.Key0, jp.KeyHat,
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
		},
		{
			jp.KeyF7, jp.KeyF8, jp.KeyF9, jp.KeyF10, jp.KeyF11, jp.KeyF12,
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0,
		},
	}, keyboard.InvertDiode(false))
	if err != nil {
		fmt.Printf("Error adding expander keyboard: %s\n", err)
	}
	mkRight.SetCallback(func(layer, index int, state keyboard.State) {
		row := index / len(colPinsRight)
		col := index % len(colPinsRight)
		fmt.Printf("right row: %d, col: %d state: %d\n", row, col, state)
		select {
		case ch <- RCS{row: row, col: col, state: state}:
		}
	})

	mkEncoder := d.AddRotaryKeyboard(machine.P0_23, machine.P0_21, [][]keyboard.Keycode{
		{
			jp.KeyMediaVolumeDec, jp.KeyMediaVolumeInc,
		},
	})
	mkEncoder.SetCallback(func(layer, index int, state keyboard.State) {
		row := index / len(colPinsRight)
		col := index % len(colPinsRight)
		fmt.Printf("encoder row: %d, col: %d state: %d\n", row, col, state)
		select {
		case ch <- RCS{row: row, col: col, state: state}:
		}
	})

	go func() {
		for {
			select {
			case x := <-ch:
				c := color.RGBA{R: 255, G: 255, B: 255, A: 255}
				if x.state == keyboard.PressToRelease {
					c = color.RGBA{A: 255}
				}
				_ = tinydraw.FilledRectangle(&display, 10+20*int16(x.col), 8*int16(x.row), 18, 6, c)
				_ = display.Display()
			}
		}
	}()

	d.AddRGBMatrix(0x40, 44, rgbLedMap, []keyboard.RgbAnimation{
		rgbanimations.GetDirectAnim(),
		rgbanimations.GetSolidColorAnim(),
		rgbanimations.GetAlphasModsAnim(),
	}, &led)

	loadKeyboardDef()

	fmt.Printf("Starting main loop\n")

	d.Debug = true
	err = d.Loop(context.Background())

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Loop exit without Error\n")
}
