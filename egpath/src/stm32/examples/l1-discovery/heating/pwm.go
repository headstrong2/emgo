package main

import (
	"mmio"

	"stm32/hal/raw/tim"
)

// setupPulsePWM setups the timer t to perform pulse PWM on all 4 channels.
//
// pclk is peripheral clock (use system.APBx.Clock() to obtain correct value,
// check out doc to find out to which bus the timer t is connected), periodms
// is PWM period in millisecond, maxCCR is CCR value that should correspond to
// 100% duty cycle (PWM full on). The PWM resolution is equal to maxCCR+1.
// Because pclk is usually multiple of 10^6 Hz and it is multipled by period and
// divided by resolution to calculate the prescaler value, it is good to use
// 10^n * periodms/m - 1 as maxCCR to avoid rounding errors.
//
// The timer is configured in center-aligned PWM, one-pulse mode. Pulses should
// be generated by software using [store CCRx, set UG, set CEN] sequence. This
// allows the timer to only counting up and stops count at ARR value without
// reseting CNT (in edge-aligned mode CNT is reset to 0 after overflow which
// means that for CCRx > 0 channel x remains in active state after pulse
// generation). Ensure CCRx value is always less than maxCCR then the x channel
// will be always set inactive at the end of the pulse and stays inactive until
// the next pulse will be generated. This ensures automatic disable of heater
// after most software crashes.
func setupPulsePWM(t *tim.TIM_Periph, pclk uint, periodms, maxCCR int) {
	t.PSC.U16.Store(uint16(int(pclk/1000)*periodms/(maxCCR+1) - 1))
	t.ARR.U32.Store(uint32(maxCCR))
	t.CCMR1.StoreBits(
		tim.OC1M|tim.OC2M|tim.OC1PE|tim.OC2PE,
		6<<tim.OC1Mn|6<<tim.OC2Mn|1<<tim.OC1PEn|1<<tim.OC2PEn,
	)
	t.CCMR2.StoreBits(
		tim.OC3M|tim.OC4M|tim.OC3PE|tim.OC4PE,
		6<<tim.OC3Mn|6<<tim.OC4Mn|1<<tim.OC3PEn|1<<tim.OC4PEn,
	)
	t.CCER.SetBits(tim.CC1E | tim.CC2E | tim.CC3E)
	t.CR1.StoreBits(
		tim.OPM|tim.CMS|tim.URS,
		1<<tim.OPMn|1<<tim.CMSn|1<<tim.URSn,
	)
	t.DIER.Store(tim.UIE)
}

type pulsePWM3 struct {
	t   *tim.TIM_Periph
	ccr [3]*mmio.U32
}

func (p *pulsePWM3) Init(t *tim.TIM_Periph) {
	p.t = t
	p.ccr[0] = &t.CCR1.U32
	p.ccr[1] = &t.CCR2.U32
	p.ccr[2] = &t.CCR3.U32
}

func (p *pulsePWM3) Timer() *tim.TIM_Periph {
	return p.t
}

func (p *pulsePWM3) Next() {
	p.ccr[0], p.ccr[1], p.ccr[2] = p.ccr[1], p.ccr[2], p.ccr[0]
}

func (p *pulsePWM3) Set(v int) {
	max := p.t.ARR.U32.Load() - 1
	v32 := uint32(v)
	switch {
	case v32 <= max:
		p.ccr[0].Store(v32)
		p.ccr[1].Store(0)
		p.ccr[2].Store(0)
	case v32 <= 2*max:
		p.ccr[0].Store(max)
		p.ccr[1].Store(v32 - max)
		p.ccr[2].Store(0)
	case v32 <= 3*max:
		p.ccr[0].Store(max)
		p.ccr[1].Store(max)
		p.ccr[2].Store(v32 - 2*max)
	default:
		p.ccr[0].Store(max)
		p.ccr[1].Store(max)
		p.ccr[2].Store(max)
	}
}

func (p *pulsePWM3) Pulse() {
	p.t.EGR.Store(tim.UG)
	p.t.CEN().Set()
}

func (p *pulsePWM3) ClearIF() {
	p.t.SR.Store(0)
}
