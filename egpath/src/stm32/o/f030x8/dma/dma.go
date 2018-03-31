// Peripheral: DMA_Periph  DMA Controller.
// Instances:
//  DMA1  mmap.DMA1_BASE
// Registers:
//  0x00 32  ISR  Interrupt status register.
//  0x04 32  IFCR Interrupt flag clear register.
// Import:
//  stm32/o/f030x8/mmap
package dma

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	GIF1  ISR = 0x01 << 0  //+ Channel 1 Global interrupt flag.
	TCIF1 ISR = 0x01 << 1  //+ Channel 1 Transfer Complete flag.
	HTIF1 ISR = 0x01 << 2  //+ Channel 1 Half Transfer flag.
	TEIF1 ISR = 0x01 << 3  //+ Channel 1 Transfer Error flag.
	GIF2  ISR = 0x01 << 4  //+ Channel 2 Global interrupt flag.
	TCIF2 ISR = 0x01 << 5  //+ Channel 2 Transfer Complete flag.
	HTIF2 ISR = 0x01 << 6  //+ Channel 2 Half Transfer flag.
	TEIF2 ISR = 0x01 << 7  //+ Channel 2 Transfer Error flag.
	GIF3  ISR = 0x01 << 8  //+ Channel 3 Global interrupt flag.
	TCIF3 ISR = 0x01 << 9  //+ Channel 3 Transfer Complete flag.
	HTIF3 ISR = 0x01 << 10 //+ Channel 3 Half Transfer flag.
	TEIF3 ISR = 0x01 << 11 //+ Channel 3 Transfer Error flag.
	GIF4  ISR = 0x01 << 12 //+ Channel 4 Global interrupt flag.
	TCIF4 ISR = 0x01 << 13 //+ Channel 4 Transfer Complete flag.
	HTIF4 ISR = 0x01 << 14 //+ Channel 4 Half Transfer flag.
	TEIF4 ISR = 0x01 << 15 //+ Channel 4 Transfer Error flag.
	GIF5  ISR = 0x01 << 16 //+ Channel 5 Global interrupt flag.
	TCIF5 ISR = 0x01 << 17 //+ Channel 5 Transfer Complete flag.
	HTIF5 ISR = 0x01 << 18 //+ Channel 5 Half Transfer flag.
	TEIF5 ISR = 0x01 << 19 //+ Channel 5 Transfer Error flag.
)

const (
	GIF1n  = 0
	TCIF1n = 1
	HTIF1n = 2
	TEIF1n = 3
	GIF2n  = 4
	TCIF2n = 5
	HTIF2n = 6
	TEIF2n = 7
	GIF3n  = 8
	TCIF3n = 9
	HTIF3n = 10
	TEIF3n = 11
	GIF4n  = 12
	TCIF4n = 13
	HTIF4n = 14
	TEIF4n = 15
	GIF5n  = 16
	TCIF5n = 17
	HTIF5n = 18
	TEIF5n = 19
)

const (
	CGIF1  IFCR = 0x01 << 0  //+ Channel 1 Global interrupt clear.
	CTCIF1 IFCR = 0x01 << 1  //+ Channel 1 Transfer Complete clear.
	CHTIF1 IFCR = 0x01 << 2  //+ Channel 1 Half Transfer clear.
	CTEIF1 IFCR = 0x01 << 3  //+ Channel 1 Transfer Error clear.
	CGIF2  IFCR = 0x01 << 4  //+ Channel 2 Global interrupt clear.
	CTCIF2 IFCR = 0x01 << 5  //+ Channel 2 Transfer Complete clear.
	CHTIF2 IFCR = 0x01 << 6  //+ Channel 2 Half Transfer clear.
	CTEIF2 IFCR = 0x01 << 7  //+ Channel 2 Transfer Error clear.
	CGIF3  IFCR = 0x01 << 8  //+ Channel 3 Global interrupt clear.
	CTCIF3 IFCR = 0x01 << 9  //+ Channel 3 Transfer Complete clear.
	CHTIF3 IFCR = 0x01 << 10 //+ Channel 3 Half Transfer clear.
	CTEIF3 IFCR = 0x01 << 11 //+ Channel 3 Transfer Error clear.
	CGIF4  IFCR = 0x01 << 12 //+ Channel 4 Global interrupt clear.
	CTCIF4 IFCR = 0x01 << 13 //+ Channel 4 Transfer Complete clear.
	CHTIF4 IFCR = 0x01 << 14 //+ Channel 4 Half Transfer clear.
	CTEIF4 IFCR = 0x01 << 15 //+ Channel 4 Transfer Error clear.
	CGIF5  IFCR = 0x01 << 16 //+ Channel 5 Global interrupt clear.
	CTCIF5 IFCR = 0x01 << 17 //+ Channel 5 Transfer Complete clear.
	CHTIF5 IFCR = 0x01 << 18 //+ Channel 5 Half Transfer clear.
	CTEIF5 IFCR = 0x01 << 19 //+ Channel 5 Transfer Error clear.
)

const (
	CGIF1n  = 0
	CTCIF1n = 1
	CHTIF1n = 2
	CTEIF1n = 3
	CGIF2n  = 4
	CTCIF2n = 5
	CHTIF2n = 6
	CTEIF2n = 7
	CGIF3n  = 8
	CTCIF3n = 9
	CHTIF3n = 10
	CTEIF3n = 11
	CGIF4n  = 12
	CTCIF4n = 13
	CHTIF4n = 14
	CTEIF4n = 15
	CGIF5n  = 16
	CTCIF5n = 17
	CHTIF5n = 18
	CTEIF5n = 19
)
