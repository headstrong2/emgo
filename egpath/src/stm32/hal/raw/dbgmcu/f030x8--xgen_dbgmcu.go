// +build f030x8

package dbgmcu

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x8/mmap"
)

type DBGMCU_Periph struct {
	IDCODE RIDCODE
	CR     RCR
	APB1FZ RAPB1FZ
	APB2FZ RAPB2FZ
}

func (p *DBGMCU_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DBGMCU = (*DBGMCU_Periph)(unsafe.Pointer(uintptr(mmap.DBGMCU_BASE)))

type IDCODE uint32

func (b IDCODE) Field(mask IDCODE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDCODE) J(v int) IDCODE {
	return IDCODE(bits.Make32(v, uint32(mask)))
}

type RIDCODE struct{ mmio.U32 }

func (r *RIDCODE) Bits(mask IDCODE) IDCODE  { return IDCODE(r.U32.Bits(uint32(mask))) }
func (r *RIDCODE) StoreBits(mask, b IDCODE) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIDCODE) SetBits(mask IDCODE)      { r.U32.SetBits(uint32(mask)) }
func (r *RIDCODE) ClearBits(mask IDCODE)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIDCODE) Load() IDCODE             { return IDCODE(r.U32.Load()) }
func (r *RIDCODE) Store(b IDCODE)           { r.U32.Store(uint32(b)) }

func (r *RIDCODE) AtomicStoreBits(mask, b IDCODE) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIDCODE) AtomicSetBits(mask IDCODE)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIDCODE) AtomicClearBits(mask IDCODE)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIDCODE struct{ mmio.UM32 }

func (rm RMIDCODE) Load() IDCODE   { return IDCODE(rm.UM32.Load()) }
func (rm RMIDCODE) Store(b IDCODE) { rm.UM32.Store(uint32(b)) }

func (p *DBGMCU_Periph) DEV_ID() RMIDCODE {
	return RMIDCODE{mmio.UM32{&p.IDCODE.U32, uint32(DEV_ID)}}
}

func (p *DBGMCU_Periph) REV_ID() RMIDCODE {
	return RMIDCODE{mmio.UM32{&p.IDCODE.U32, uint32(REV_ID)}}
}

type CR uint32

func (b CR) Field(mask CR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR) J(v int) CR {
	return CR(bits.Make32(v, uint32(mask)))
}

type RCR struct{ mmio.U32 }

func (r *RCR) Bits(mask CR) CR      { return CR(r.U32.Bits(uint32(mask))) }
func (r *RCR) StoreBits(mask, b CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR) SetBits(mask CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR) ClearBits(mask CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR) Load() CR             { return CR(r.U32.Load()) }
func (r *RCR) Store(b CR)           { r.U32.Store(uint32(b)) }

func (r *RCR) AtomicStoreBits(mask, b CR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR) AtomicSetBits(mask CR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR) AtomicClearBits(mask CR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR struct{ mmio.UM32 }

func (rm RMCR) Load() CR   { return CR(rm.UM32.Load()) }
func (rm RMCR) Store(b CR) { rm.UM32.Store(uint32(b)) }

func (p *DBGMCU_Periph) DBG_STOP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_STANDBY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBG_STANDBY)}}
}

type APB1FZ uint32

func (b APB1FZ) Field(mask APB1FZ) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1FZ) J(v int) APB1FZ {
	return APB1FZ(bits.Make32(v, uint32(mask)))
}

type RAPB1FZ struct{ mmio.U32 }

func (r *RAPB1FZ) Bits(mask APB1FZ) APB1FZ  { return APB1FZ(r.U32.Bits(uint32(mask))) }
func (r *RAPB1FZ) StoreBits(mask, b APB1FZ) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1FZ) SetBits(mask APB1FZ)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1FZ) ClearBits(mask APB1FZ)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1FZ) Load() APB1FZ             { return APB1FZ(r.U32.Load()) }
func (r *RAPB1FZ) Store(b APB1FZ)           { r.U32.Store(uint32(b)) }

func (r *RAPB1FZ) AtomicStoreBits(mask, b APB1FZ) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1FZ) AtomicSetBits(mask APB1FZ)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB1FZ) AtomicClearBits(mask APB1FZ)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB1FZ struct{ mmio.UM32 }

func (rm RMAPB1FZ) Load() APB1FZ   { return APB1FZ(rm.UM32.Load()) }
func (rm RMAPB1FZ) Store(b APB1FZ) { rm.UM32.Store(uint32(b)) }

type APB2FZ uint32

func (b APB2FZ) Field(mask APB2FZ) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2FZ) J(v int) APB2FZ {
	return APB2FZ(bits.Make32(v, uint32(mask)))
}

type RAPB2FZ struct{ mmio.U32 }

func (r *RAPB2FZ) Bits(mask APB2FZ) APB2FZ  { return APB2FZ(r.U32.Bits(uint32(mask))) }
func (r *RAPB2FZ) StoreBits(mask, b APB2FZ) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2FZ) SetBits(mask APB2FZ)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2FZ) ClearBits(mask APB2FZ)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2FZ) Load() APB2FZ             { return APB2FZ(r.U32.Load()) }
func (r *RAPB2FZ) Store(b APB2FZ)           { r.U32.Store(uint32(b)) }

func (r *RAPB2FZ) AtomicStoreBits(mask, b APB2FZ) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2FZ) AtomicSetBits(mask APB2FZ)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB2FZ) AtomicClearBits(mask APB2FZ)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB2FZ struct{ mmio.UM32 }

func (rm RMAPB2FZ) Load() APB2FZ   { return APB2FZ(rm.UM32.Load()) }
func (rm RMAPB2FZ) Store(b APB2FZ) { rm.UM32.Store(uint32(b)) }
