// +build f10x_md f10x_hd f40_41xxx f411xe l1xx_md

package sdmmc

import (
	"unsafe"

	"stm32/hal/raw/mmap"
	"stm32/hal/raw/sdio"
)

type Periph struct {
	raw sdio.SDIO_Periph
}

//emgo:const
var SDMMC = (*Periph)(unsafe.Pointer(mmap.SDIO_BASE))
