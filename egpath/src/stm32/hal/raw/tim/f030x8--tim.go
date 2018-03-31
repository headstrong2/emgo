// +build f030x8

// Peripheral: TIM_Periph  TIM.
// Instances:
//  TIM3   mmap.TIM3_BASE
//  TIM6   mmap.TIM6_BASE
//  TIM14  mmap.TIM14_BASE
//  TIM1   mmap.TIM1_BASE
//  TIM15  mmap.TIM15_BASE
//  TIM16  mmap.TIM16_BASE
//  TIM17  mmap.TIM17_BASE
// Registers:
//  0x00 32  CR1   Control register 1.
//  0x04 32  CR2   Control register 2.
//  0x08 32  SMCR  Slave Mode Control register.
//  0x0C 32  DIER  DMA/interrupt enable register.
//  0x10 32  SR    Status register.
//  0x14 32  EGR   Event generation register.
//  0x18 32  CCMR1 Capture/compare mode register 1.
//  0x1C 32  CCMR2 Capture/compare mode register 2.
//  0x20 32  CCER  Capture/compare enable register.
//  0x24 32  CNT   Counter register.
//  0x28 32  PSC   Prescaler register.
//  0x2C 32  ARR   Auto-reload register.
//  0x30 32  RCR   Repetition counter register.
//  0x34 32  CCR1  Capture/compare register 1.
//  0x38 32  CCR2  Capture/compare register 2.
//  0x3C 32  CCR3  Capture/compare register 3.
//  0x40 32  CCR4  Capture/compare register 4.
//  0x44 32  BDTR  Break and dead-time register.
//  0x48 32  DCR   DMA control register.
//  0x4C 32  DMAR  DMA address for full transfer register.
//  0x50 32  OR    Option register.
// Import:
//  stm32/o/f030x8/mmap
package tim

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CEN  CR1 = 0x01 << 0 //+ Counter enable.
	UDIS CR1 = 0x01 << 1 //+ Update disable.
	URS  CR1 = 0x01 << 2 //+ Update request source.
	OPM  CR1 = 0x01 << 3 //+ One pulse mode.
	DIR  CR1 = 0x01 << 4 //+ Direction.
	CMS  CR1 = 0x03 << 5 //+ CMS[1:0] bits (Center-aligned mode selection).
	ARPE CR1 = 0x01 << 7 //+ Auto-reload preload enable.
	CKD  CR1 = 0x03 << 8 //+ CKD[1:0] bits (clock division).
)

const (
	CENn  = 0
	UDISn = 1
	URSn  = 2
	OPMn  = 3
	DIRn  = 4
	CMSn  = 5
	ARPEn = 7
	CKDn  = 8
)

const (
	CCPC  CR2 = 0x01 << 0  //+ Capture/Compare Preloaded Control.
	CCUS  CR2 = 0x01 << 2  //+ Capture/Compare Control Update Selection.
	CCDS  CR2 = 0x01 << 3  //+ Capture/Compare DMA Selection.
	MMS   CR2 = 0x07 << 4  //+ MMS[2:0] bits (Master Mode Selection).
	TI1S  CR2 = 0x01 << 7  //+ TI1 Selection.
	OIS1  CR2 = 0x01 << 8  //+ Output Idle state 1 (OC1 output).
	OIS1N CR2 = 0x01 << 9  //+ Output Idle state 1 (OC1N output).
	OIS2  CR2 = 0x01 << 10 //+ Output Idle state 2 (OC2 output).
	OIS2N CR2 = 0x01 << 11 //+ Output Idle state 2 (OC2N output).
	OIS3  CR2 = 0x01 << 12 //+ Output Idle state 3 (OC3 output).
	OIS3N CR2 = 0x01 << 13 //+ Output Idle state 3 (OC3N output).
	OIS4  CR2 = 0x01 << 14 //+ Output Idle state 4 (OC4 output).
)

const (
	CCPCn  = 0
	CCUSn  = 2
	CCDSn  = 3
	MMSn   = 4
	TI1Sn  = 7
	OIS1n  = 8
	OIS1Nn = 9
	OIS2n  = 10
	OIS2Nn = 11
	OIS3n  = 12
	OIS3Nn = 13
	OIS4n  = 14
)

const (
	SMS  SMCR = 0x07 << 0  //+ SMS[2:0] bits (Slave mode selection).
	OCCS SMCR = 0x01 << 3  //+ OCREF clear selection.
	TS   SMCR = 0x07 << 4  //+ TS[2:0] bits (Trigger selection).
	MSM  SMCR = 0x01 << 7  //+ Master/slave mode.
	ETF  SMCR = 0x0F << 8  //+ ETF[3:0] bits (External trigger filter).
	ETPS SMCR = 0x03 << 12 //+ ETPS[1:0] bits (External trigger prescaler).
	ECE  SMCR = 0x01 << 14 //+ External clock enable.
	ETP  SMCR = 0x01 << 15 //+ External trigger polarity.
)

const (
	SMSn  = 0
	OCCSn = 3
	TSn   = 4
	MSMn  = 7
	ETFn  = 8
	ETPSn = 12
	ECEn  = 14
	ETPn  = 15
)

const (
	UIE   DIER = 0x01 << 0  //+ Update interrupt enable.
	CC1IE DIER = 0x01 << 1  //+ Capture/Compare 1 interrupt enable.
	CC2IE DIER = 0x01 << 2  //+ Capture/Compare 2 interrupt enable.
	CC3IE DIER = 0x01 << 3  //+ Capture/Compare 3 interrupt enable.
	CC4IE DIER = 0x01 << 4  //+ Capture/Compare 4 interrupt enable.
	COMIE DIER = 0x01 << 5  //+ COM interrupt enable.
	TIE   DIER = 0x01 << 6  //+ Trigger interrupt enable.
	BIE   DIER = 0x01 << 7  //+ Break interrupt enable.
	UDE   DIER = 0x01 << 8  //+ Update DMA request enable.
	CC1DE DIER = 0x01 << 9  //+ Capture/Compare 1 DMA request enable.
	CC2DE DIER = 0x01 << 10 //+ Capture/Compare 2 DMA request enable.
	CC3DE DIER = 0x01 << 11 //+ Capture/Compare 3 DMA request enable.
	CC4DE DIER = 0x01 << 12 //+ Capture/Compare 4 DMA request enable.
	COMDE DIER = 0x01 << 13 //+ COM DMA request enable.
	TDE   DIER = 0x01 << 14 //+ Trigger DMA request enable.
)

const (
	UIEn   = 0
	CC1IEn = 1
	CC2IEn = 2
	CC3IEn = 3
	CC4IEn = 4
	COMIEn = 5
	TIEn   = 6
	BIEn   = 7
	UDEn   = 8
	CC1DEn = 9
	CC2DEn = 10
	CC3DEn = 11
	CC4DEn = 12
	COMDEn = 13
	TDEn   = 14
)

const (
	UIF   SR = 0x01 << 0  //+ Update interrupt Flag.
	CC1IF SR = 0x01 << 1  //+ Capture/Compare 1 interrupt Flag.
	CC2IF SR = 0x01 << 2  //+ Capture/Compare 2 interrupt Flag.
	CC3IF SR = 0x01 << 3  //+ Capture/Compare 3 interrupt Flag.
	CC4IF SR = 0x01 << 4  //+ Capture/Compare 4 interrupt Flag.
	COMIF SR = 0x01 << 5  //+ COM interrupt Flag.
	TIF   SR = 0x01 << 6  //+ Trigger interrupt Flag.
	BIF   SR = 0x01 << 7  //+ Break interrupt Flag.
	CC1OF SR = 0x01 << 9  //+ Capture/Compare 1 Overcapture Flag.
	CC2OF SR = 0x01 << 10 //+ Capture/Compare 2 Overcapture Flag.
	CC3OF SR = 0x01 << 11 //+ Capture/Compare 3 Overcapture Flag.
	CC4OF SR = 0x01 << 12 //+ Capture/Compare 4 Overcapture Flag.
)

const (
	UIFn   = 0
	CC1IFn = 1
	CC2IFn = 2
	CC3IFn = 3
	CC4IFn = 4
	COMIFn = 5
	TIFn   = 6
	BIFn   = 7
	CC1OFn = 9
	CC2OFn = 10
	CC3OFn = 11
	CC4OFn = 12
)

const (
	UG   EGR = 0x01 << 0 //+ Update Generation.
	CC1G EGR = 0x01 << 1 //+ Capture/Compare 1 Generation.
	CC2G EGR = 0x01 << 2 //+ Capture/Compare 2 Generation.
	CC3G EGR = 0x01 << 3 //+ Capture/Compare 3 Generation.
	CC4G EGR = 0x01 << 4 //+ Capture/Compare 4 Generation.
	COMG EGR = 0x01 << 5 //+ Capture/Compare Control Update Generation.
	TG   EGR = 0x01 << 6 //+ Trigger Generation.
	BG   EGR = 0x01 << 7 //+ Break Generation.
)

const (
	UGn   = 0
	CC1Gn = 1
	CC2Gn = 2
	CC3Gn = 3
	CC4Gn = 4
	COMGn = 5
	TGn   = 6
	BGn   = 7
)

const (
	CC1S   CCMR1 = 0x03 << 0  //+ CC1S[1:0] bits (Capture/Compare 1 Selection).
	OC1FE  CCMR1 = 0x01 << 2  //+ Output Compare 1 Fast enable.
	OC1PE  CCMR1 = 0x01 << 3  //+ Output Compare 1 Preload enable.
	OC1M   CCMR1 = 0x07 << 4  //+ OC1M[2:0] bits (Output Compare 1 Mode).
	OC1CE  CCMR1 = 0x01 << 7  //+ Output Compare 1Clear Enable.
	CC2S   CCMR1 = 0x03 << 8  //+ CC2S[1:0] bits (Capture/Compare 2 Selection).
	OC2FE  CCMR1 = 0x01 << 10 //+ Output Compare 2 Fast enable.
	OC2PE  CCMR1 = 0x01 << 11 //+ Output Compare 2 Preload enable.
	OC2M   CCMR1 = 0x07 << 12 //+ OC2M[2:0] bits (Output Compare 2 Mode).
	OC2CE  CCMR1 = 0x01 << 15 //+ Output Compare 2 Clear Enable.
	IC1PSC CCMR1 = 0x03 << 2  //+ IC1PSC[1:0] bits (Input Capture 1 Prescaler).
	IC1F   CCMR1 = 0x0F << 4  //+ IC1F[3:0] bits (Input Capture 1 Filter).
	IC2PSC CCMR1 = 0x03 << 10 //+ IC2PSC[1:0] bits (Input Capture 2 Prescaler).
	IC2F   CCMR1 = 0x0F << 12 //+ IC2F[3:0] bits (Input Capture 2 Filter).
)

const (
	CC1Sn   = 0
	OC1FEn  = 2
	OC1PEn  = 3
	OC1Mn   = 4
	OC1CEn  = 7
	CC2Sn   = 8
	OC2FEn  = 10
	OC2PEn  = 11
	OC2Mn   = 12
	OC2CEn  = 15
	IC1PSCn = 2
	IC1Fn   = 4
	IC2PSCn = 10
	IC2Fn   = 12
)

const (
	CC3S   CCMR2 = 0x03 << 0  //+ CC3S[1:0] bits (Capture/Compare 3 Selection).
	OC3FE  CCMR2 = 0x01 << 2  //+ Output Compare 3 Fast enable.
	OC3PE  CCMR2 = 0x01 << 3  //+ Output Compare 3 Preload enable.
	OC3M   CCMR2 = 0x07 << 4  //+ OC3M[2:0] bits (Output Compare 3 Mode).
	OC3CE  CCMR2 = 0x01 << 7  //+ Output Compare 3 Clear Enable.
	CC4S   CCMR2 = 0x03 << 8  //+ CC4S[1:0] bits (Capture/Compare 4 Selection).
	OC4FE  CCMR2 = 0x01 << 10 //+ Output Compare 4 Fast enable.
	OC4PE  CCMR2 = 0x01 << 11 //+ Output Compare 4 Preload enable.
	OC4M   CCMR2 = 0x07 << 12 //+ OC4M[2:0] bits (Output Compare 4 Mode).
	OC4CE  CCMR2 = 0x01 << 15 //+ Output Compare 4 Clear Enable.
	IC3PSC CCMR2 = 0x03 << 2  //+ IC3PSC[1:0] bits (Input Capture 3 Prescaler).
	IC3F   CCMR2 = 0x0F << 4  //+ IC3F[3:0] bits (Input Capture 3 Filter).
	IC4PSC CCMR2 = 0x03 << 10 //+ IC4PSC[1:0] bits (Input Capture 4 Prescaler).
	IC4F   CCMR2 = 0x0F << 12 //+ IC4F[3:0] bits (Input Capture 4 Filter).
)

const (
	CC3Sn   = 0
	OC3FEn  = 2
	OC3PEn  = 3
	OC3Mn   = 4
	OC3CEn  = 7
	CC4Sn   = 8
	OC4FEn  = 10
	OC4PEn  = 11
	OC4Mn   = 12
	OC4CEn  = 15
	IC3PSCn = 2
	IC3Fn   = 4
	IC4PSCn = 10
	IC4Fn   = 12
)

const (
	CC1E  CCER = 0x01 << 0  //+ Capture/Compare 1 output enable.
	CC1P  CCER = 0x01 << 1  //+ Capture/Compare 1 output Polarity.
	CC1NE CCER = 0x01 << 2  //+ Capture/Compare 1 Complementary output enable.
	CC1NP CCER = 0x01 << 3  //+ Capture/Compare 1 Complementary output Polarity.
	CC2E  CCER = 0x01 << 4  //+ Capture/Compare 2 output enable.
	CC2P  CCER = 0x01 << 5  //+ Capture/Compare 2 output Polarity.
	CC2NE CCER = 0x01 << 6  //+ Capture/Compare 2 Complementary output enable.
	CC2NP CCER = 0x01 << 7  //+ Capture/Compare 2 Complementary output Polarity.
	CC3E  CCER = 0x01 << 8  //+ Capture/Compare 3 output enable.
	CC3P  CCER = 0x01 << 9  //+ Capture/Compare 3 output Polarity.
	CC3NE CCER = 0x01 << 10 //+ Capture/Compare 3 Complementary output enable.
	CC3NP CCER = 0x01 << 11 //+ Capture/Compare 3 Complementary output Polarity.
	CC4E  CCER = 0x01 << 12 //+ Capture/Compare 4 output enable.
	CC4P  CCER = 0x01 << 13 //+ Capture/Compare 4 output Polarity.
	CC4NP CCER = 0x01 << 15 //+ Capture/Compare 4 Complementary output Polarity.
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NEn = 2
	CC1NPn = 3
	CC2En  = 4
	CC2Pn  = 5
	CC2NEn = 6
	CC2NPn = 7
	CC3En  = 8
	CC3Pn  = 9
	CC3NEn = 10
	CC3NPn = 11
	CC4En  = 12
	CC4Pn  = 13
	CC4NPn = 15
)

const (
	REP RCR = 0xFF << 0 //+ Repetition Counter Value.
)

const (
	REPn = 0
)

const (
	DTG  BDTR = 0xFF << 0  //+ DTG[0:7] bits (Dead-Time Generator set-up).
	LOCK BDTR = 0x03 << 8  //+ LOCK[1:0] bits (Lock Configuration).
	OSSI BDTR = 0x01 << 10 //+ Off-State Selection for Idle mode.
	OSSR BDTR = 0x01 << 11 //+ Off-State Selection for Run mode.
	BKE  BDTR = 0x01 << 12 //+ Break enable.
	BKP  BDTR = 0x01 << 13 //+ Break Polarity.
	AOE  BDTR = 0x01 << 14 //+ Automatic Output enable.
	MOE  BDTR = 0x01 << 15 //+ Main Output enable.
)

const (
	DTGn  = 0
	LOCKn = 8
	OSSIn = 10
	OSSRn = 11
	BKEn  = 12
	BKPn  = 13
	AOEn  = 14
	MOEn  = 15
)

const (
	DBA DCR = 0x1F << 0 //+ DBA[4:0] bits (DMA Base Address).
	DBL DCR = 0x1F << 8 //+ DBL[4:0] bits (DMA Burst Length).
)

const (
	DBAn = 0
	DBLn = 8
)

const (
	DMAB DMAR = 0xFFFF << 0 //+ DMA register for burst accesses.
)

const (
	DMABn = 0
)
