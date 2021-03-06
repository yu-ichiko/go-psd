package psd

import (
	"errors"
	"image"
)

// header error
var (
	headerLens = []int{4, 2, 6, 2, 4, 4, 2, 2}
	headerLen  = 0
	headerSig  = []byte("8BPS")

	ErrHeaderFormat    = errors.New("psd: invalid header format")
	ErrHeaderVersion   = errors.New("psd: invalid header version")
	ErrHeaderChannels  = errors.New("psd: invalid header channels")
	ErrHeaderHeight    = errors.New("psd: invalid header height")
	ErrHeaderWidth     = errors.New("psd: invalid header width")
	ErrHeaderDepth     = errors.New("psd: invalid header depth")
	ErrHeaderColorMode = errors.New("psd: invalid header colorMode")
)

const (
	ColorModeBitmap       = ColorMode(0)
	ColorModeGrayScale    = ColorMode(1)
	ColorModeIndexed      = ColorMode(2)
	ColorModeRGB          = ColorMode(3)
	ColorModeCMYK         = ColorMode(4)
	ColorModeMultichannel = ColorMode(7)
	ColorModeDuotone      = ColorMode(8)
	ColorModeLab          = ColorMode(9)
)

type Header struct {
	Version   int
	Channels  int
	Height    int
	Width     int
	Depth     int
	ColorMode ColorMode
}

func (h *Header) IsPSB() bool {
	return h.Version == 2
}

func (h *Header) Rect() image.Rectangle {
	return image.Rect(0, 0, h.Width, h.Height)
}

type ColorMode int

func (c ColorMode) String() string {
	switch c {
	case ColorModeBitmap:
		return "Bitmap"
	case ColorModeGrayScale:
		return "Grayscale"
	case ColorModeIndexed:
		return "Indexed"
	case ColorModeRGB:
		return "RGB"
	case ColorModeCMYK:
		return "CMYK"
	case ColorModeMultichannel:
		return "Multichannel"
	case ColorModeDuotone:
		return "Duotone"
	case ColorModeLab:
		return "Lab"
	}
	return ""
}

func (c ColorMode) Channels() int {
	switch c {
	case ColorModeBitmap, ColorModeGrayScale, ColorModeIndexed:
		return 1
	case ColorModeRGB:
		return 3
	case ColorModeCMYK:
		return 4
	}
	return -1
}

func init() {
	for _, n := range headerLens {
		headerLen += n
	}
}
