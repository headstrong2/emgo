// +build f030x8

// Peripheral: SPI_Periph  Serial Peripheral Interface.
// Instances:
//  SPI1  mmap.SPI1_BASE
//  SPI2  mmap.SPI2_BASE
// Registers:
//  0x00 32  CR1     Control register 1 (not used in I2S mode).
//  0x04 32  CR2     Control register 2.
//  0x08 32  SR      Status register.
//  0x0C 32  DR      Data register.
//  0x10 32  CRCPR   CRC polynomial register (not used in I2S mode).
//  0x14 32  RXCRCR  Rx CRC register (not used in I2S mode).
//  0x18 32  TXCRCR  Tx CRC register (not used in I2S mode).
//  0x1C 32  I2SCFGR SPI_I2S configuration register.
// Import:
//  stm32/o/f030x8/mmap
package spi

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CPHA     CR1 = 0x01 << 0  //+ Clock Phase.
	CPOL     CR1 = 0x01 << 1  //+ Clock Polarity.
	MSTR     CR1 = 0x01 << 2  //+ Master Selection.
	BR       CR1 = 0x07 << 3  //+ BR[2:0] bits (Baud Rate Control).
	SPE      CR1 = 0x01 << 6  //+ SPI Enable.
	LSBFIRST CR1 = 0x01 << 7  //+ Frame Format.
	SSI      CR1 = 0x01 << 8  //+ Internal slave select.
	SSM      CR1 = 0x01 << 9  //+ Software slave management.
	RXONLY   CR1 = 0x01 << 10 //+ Receive only.
	CRCL     CR1 = 0x01 << 11 //+ CRC Length.
	CRCNEXT  CR1 = 0x01 << 12 //+ Transmit CRC next.
	CRCEN    CR1 = 0x01 << 13 //+ Hardware CRC calculation enable.
	BIDIOE   CR1 = 0x01 << 14 //+ Output enable in bidirectional mode.
	BIDIMODE CR1 = 0x01 << 15 //+ Bidirectional data mode enable.
)

const (
	CPHAn     = 0
	CPOLn     = 1
	MSTRn     = 2
	BRn       = 3
	SPEn      = 6
	LSBFIRSTn = 7
	SSIn      = 8
	SSMn      = 9
	RXONLYn   = 10
	CRCLn     = 11
	CRCNEXTn  = 12
	CRCENn    = 13
	BIDIOEn   = 14
	BIDIMODEn = 15
)

const (
	RXDMAEN CR2 = 0x01 << 0  //+ Rx Buffer DMA Enable.
	TXDMAEN CR2 = 0x01 << 1  //+ Tx Buffer DMA Enable.
	SSOE    CR2 = 0x01 << 2  //+ SS Output Enable.
	NSSP    CR2 = 0x01 << 3  //+ NSS pulse management Enable.
	FRF     CR2 = 0x01 << 4  //+ Frame Format Enable.
	ERRIE   CR2 = 0x01 << 5  //+ Error Interrupt Enable.
	RXNEIE  CR2 = 0x01 << 6  //+ RX buffer Not Empty Interrupt Enable.
	TXEIE   CR2 = 0x01 << 7  //+ Tx buffer Empty Interrupt Enable.
	DS      CR2 = 0x0F << 8  //+ DS[3:0] Data Size.
	FRXTH   CR2 = 0x01 << 12 //+ FIFO reception Threshold.
	LDMARX  CR2 = 0x01 << 13 //+ Last DMA transfer for reception.
	LDMATX  CR2 = 0x01 << 14 //+ Last DMA transfer for transmission.
)

const (
	RXDMAENn = 0
	TXDMAENn = 1
	SSOEn    = 2
	NSSPn    = 3
	FRFn     = 4
	ERRIEn   = 5
	RXNEIEn  = 6
	TXEIEn   = 7
	DSn      = 8
	FRXTHn   = 12
	LDMARXn  = 13
	LDMATXn  = 14
)

const (
	RXNE   SR = 0x01 << 0  //+ Receive buffer Not Empty.
	TXE    SR = 0x01 << 1  //+ Transmit buffer Empty.
	CRCERR SR = 0x01 << 4  //+ CRC Error flag.
	MODF   SR = 0x01 << 5  //+ Mode fault.
	OVR    SR = 0x01 << 6  //+ Overrun flag.
	BSY    SR = 0x01 << 7  //+ Busy flag.
	FRE    SR = 0x01 << 8  //+ TI frame format error.
	FRLVL  SR = 0x03 << 9  //+ FIFO Reception Level.
	FTLVL  SR = 0x03 << 11 //+ FIFO Transmission Level.
)

const (
	RXNEn   = 0
	TXEn    = 1
	CRCERRn = 4
	MODFn   = 5
	OVRn    = 6
	BSYn    = 7
	FREn    = 8
	FRLVLn  = 9
	FTLVLn  = 11
)

const (
	CRCPOLY CRCPR = 0xFFFFFFFF << 0 //+ CRC polynomial register.
)

const (
	CRCPOLYn = 0
)

const (
	RXCRC RXCRCR = 0xFFFFFFFF << 0 //+ Rx CRC Register.
)

const (
	RXCRCn = 0
)

const (
	TXCRC TXCRCR = 0xFFFFFFFF << 0 //+ Tx CRC Register.
)

const (
	TXCRCn = 0
)

const (
	I2SMOD I2SCFGR = 0x01 << 11 //+ Keep for compatibility.
)

const (
	I2SMODn = 11
)
