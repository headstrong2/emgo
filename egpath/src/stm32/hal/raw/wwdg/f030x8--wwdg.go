// +build f030x8

// Peripheral: WWDG_Periph  Window WATCHDOG.
// Instances:
//  WWDG  mmap.WWDG_BASE
// Registers:
//  0x00 32  CR  Control register.
//  0x04 32  CFR Configuration register.
//  0x08 32  SR  Status register.
// Import:
//  stm32/o/f030x8/mmap
package wwdg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	T    CR = 0x7F << 0 //+ T[6:0] bits (7-Bit counter (MSB to LSB)).
	WDGA CR = 0x01 << 7 //+ Activation bit.
)

const (
	Tn    = 0
	WDGAn = 7
)

const (
	W     CFR = 0x7F << 0 //+ W[6:0] bits (7-bit window value).
	WDGTB CFR = 0x03 << 7 //+ WDGTB[1:0] bits (Timer Base).
	EWI   CFR = 0x01 << 9 //+ Early Wakeup Interrupt.
)

const (
	Wn     = 0
	WDGTBn = 7
	EWIn   = 9
)

const (
	EWIF SR = 0x01 << 0 //+ Early Wakeup Interrupt Flag.
)

const (
	EWIFn = 0
)
