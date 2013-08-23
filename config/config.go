package config

import (
	"github.com/longears/pixelslinger/midi"
)

// midi pads
const (
	FLASH_PAD        = midi.LPD8_PAD1
	TWINKLE_PAD      = midi.LPD8_PAD2
	RIPPLE_PAD       = midi.LPD8_PAD3 // todo
	SLOWMO_PAD       = midi.LPD8_PAD4 // todo
	BLINK_CIRCLE_PAD = midi.LPD8_PAD5
	BLINK_ARCH_PAD   = midi.LPD8_PAD6
	BLINK_BACK_PAD   = midi.LPD8_PAD7
)

// midi knobs
const (
	GAIN_KNOB   = midi.LPD8_KNOB1
	EYELID_KNOB = midi.LPD8_KNOB2
	SPEED_KNOB  = midi.LPD8_KNOB3
	SWITCH_KNOB = midi.LPD8_KNOB4
	MORPH_KNOB  = midi.LPD8_KNOB5
	HUE_KNOB    = midi.LPD8_KNOB6
)

// knob starting values before they have been moved
//  (because the midi hardware only sends us values when the knobs move)
var DEFAULT_KNOB_VALUES map[byte]byte

func init() {
	DEFAULT_KNOB_VALUES = map[byte]byte{
		GAIN_KNOB:   127,
		EYELID_KNOB: 127,
		SPEED_KNOB:  63,
		SWITCH_KNOB: 0,
		MORPH_KNOB:  0,
		HUE_KNOB:    0,
	}
}
