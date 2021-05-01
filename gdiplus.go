// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
)

const LANG_NEUTRAL = 0x00

const (
	Ok                        GpStatus = 0
	GenericError              GpStatus = 1
	InvalidParameter          GpStatus = 2
	OutOfMemory               GpStatus = 3
	ObjectBusy                GpStatus = 4
	InsufficientBuffer        GpStatus = 5
	NotImplemented            GpStatus = 6
	Win32Error                GpStatus = 7
	WrongState                GpStatus = 8
	Aborted                   GpStatus = 9
	FileNotFound              GpStatus = 10
	ValueOverflow             GpStatus = 11
	AccessDenied              GpStatus = 12
	UnknownImageFormat        GpStatus = 13
	FontFamilyNotFound        GpStatus = 14
	FontStyleNotFound         GpStatus = 15
	NotTrueTypeFont           GpStatus = 16
	UnsupportedGdiplusVersion GpStatus = 17
	GdiplusNotInitialized     GpStatus = 18
	PropertyNotFound          GpStatus = 19
	PropertyNotSupported      GpStatus = 20
	ProfileNotFound           GpStatus = 21
)

const (
	Color_AliceBlue            = 0xFFF0F8FF
	Color_AntiqueWhite         = 0xFFFAEBD7
	Color_Aqua                 = 0xFF00FFFF
	Color_Aquamarine           = 0xFF7FFFD4
	Color_Azure                = 0xFFF0FFFF
	Color_Beige                = 0xFFF5F5DC
	Color_Bisque               = 0xFFFFE4C4
	Color_Black                = 0xFF000000
	Color_BlanchedAlmond       = 0xFFFFEBCD
	Color_Blue                 = 0xFF0000FF
	Color_BlueViolet           = 0xFF8A2BE2
	Color_Brown                = 0xFFA52A2A
	Color_BurlyWood            = 0xFFDEB887
	Color_CadetBlue            = 0xFF5F9EA0
	Color_Chartreuse           = 0xFF7FFF00
	Color_Chocolate            = 0xFFD2691E
	Color_Coral                = 0xFFFF7F50
	Color_CornflowerBlue       = 0xFF6495ED
	Color_Cornsilk             = 0xFFFFF8DC
	Color_Crimson              = 0xFFDC143C
	Color_Cyan                 = 0xFF00FFFF
	Color_DarkBlue             = 0xFF00008B
	Color_DarkCyan             = 0xFF008B8B
	Color_DarkGoldenrod        = 0xFFB8860B
	Color_DarkGray             = 0xFFA9A9A9
	Color_DarkGreen            = 0xFF006400
	Color_DarkKhaki            = 0xFFBDB76B
	Color_DarkMagenta          = 0xFF8B008B
	Color_DarkOliveGreen       = 0xFF556B2F
	Color_DarkOrange           = 0xFFFF8C00
	Color_DarkOrchid           = 0xFF9932CC
	Color_DarkRed              = 0xFF8B0000
	Color_DarkSalmon           = 0xFFE9967A
	Color_DarkSeaGreen         = 0xFF8FBC8B
	Color_DarkSlateBlue        = 0xFF483D8B
	Color_DarkSlateGray        = 0xFF2F4F4F
	Color_DarkTurquoise        = 0xFF00CED1
	Color_DarkViolet           = 0xFF9400D3
	Color_DeepPink             = 0xFFFF1493
	Color_DeepSkyBlue          = 0xFF00BFFF
	Color_DimGray              = 0xFF696969
	Color_DodgerBlue           = 0xFF1E90FF
	Color_Firebrick            = 0xFFB22222
	Color_FloralWhite          = 0xFFFFFAF0
	Color_ForestGreen          = 0xFF228B22
	Color_Fuchsia              = 0xFFFF00FF
	Color_Gainsboro            = 0xFFDCDCDC
	Color_GhostWhite           = 0xFFF8F8FF
	Color_Gold                 = 0xFFFFD700
	Color_Goldenrod            = 0xFFDAA520
	Color_Gray                 = 0xFF808080
	Color_Green                = 0xFF008000
	Color_GreenYellow          = 0xFFADFF2F
	Color_Honeydew             = 0xFFF0FFF0
	Color_HotPink              = 0xFFFF69B4
	Color_IndianRed            = 0xFFCD5C5C
	Color_Indigo               = 0xFF4B0082
	Color_Ivory                = 0xFFFFFFF0
	Color_Khaki                = 0xFFF0E68C
	Color_Lavender             = 0xFFE6E6FA
	Color_LavenderBlush        = 0xFFFFF0F5
	Color_LawnGreen            = 0xFF7CFC00
	Color_LemonChiffon         = 0xFFFFFACD
	Color_LightBlue            = 0xFFADD8E6
	Color_LightCoral           = 0xFFF08080
	Color_LightCyan            = 0xFFE0FFFF
	Color_LightGoldenrodYellow = 0xFFFAFAD2
	Color_LightGray            = 0xFFD3D3D3
	Color_LightGreen           = 0xFF90EE90
	Color_LightPink            = 0xFFFFB6C1
	Color_LightSalmon          = 0xFFFFA07A
	Color_LightSeaGreen        = 0xFF20B2AA
	Color_LightSkyBlue         = 0xFF87CEFA
	Color_LightSlateGray       = 0xFF778899
	Color_LightSteelBlue       = 0xFFB0C4DE
	Color_LightYellow          = 0xFFFFFFE0
	Color_Lime                 = 0xFF00FF00
	Color_LimeGreen            = 0xFF32CD32
	Color_Linen                = 0xFFFAF0E6
	Color_Magenta              = 0xFFFF00FF
	Color_Maroon               = 0xFF800000
	Color_MediumAquamarine     = 0xFF66CDAA
	Color_MediumBlue           = 0xFF0000CD
	Color_MediumOrchid         = 0xFFBA55D3
	Color_MediumPurple         = 0xFF9370DB
	Color_MediumSeaGreen       = 0xFF3CB371
	Color_MediumSlateBlue      = 0xFF7B68EE
	Color_MediumSpringGreen    = 0xFF00FA9A
	Color_MediumTurquoise      = 0xFF48D1CC
	Color_MediumVioletRed      = 0xFFC71585
	Color_MidnightBlue         = 0xFF191970
	Color_MintCream            = 0xFFF5FFFA
	Color_MistyRose            = 0xFFFFE4E1
	Color_Moccasin             = 0xFFFFE4B5
	Color_NavajoWhite          = 0xFFFFDEAD
	Color_Navy                 = 0xFF000080
	Color_OldLace              = 0xFFFDF5E6
	Color_Olive                = 0xFF808000
	Color_OliveDrab            = 0xFF6B8E23
	Color_Orange               = 0xFFFFA500
	Color_OrangeRed            = 0xFFFF4500
	Color_Orchid               = 0xFFDA70D6
	Color_PaleGoldenrod        = 0xFFEEE8AA
	Color_PaleGreen            = 0xFF98FB98
	Color_PaleTurquoise        = 0xFFAFEEEE
	Color_PaleVioletRed        = 0xFFDB7093
	Color_PapayaWhip           = 0xFFFFEFD5
	Color_PeachPuff            = 0xFFFFDAB9
	Color_Peru                 = 0xFFCD853F
	Color_Pink                 = 0xFFFFC0CB
	Color_Plum                 = 0xFFDDA0DD
	Color_PowderBlue           = 0xFFB0E0E6
	Color_Purple               = 0xFF800080
	Color_Red                  = 0xFFFF0000
	Color_RosyBrown            = 0xFFBC8F8F
	Color_RoyalBlue            = 0xFF4169E1
	Color_SaddleBrown          = 0xFF8B4513
	Color_Salmon               = 0xFFFA8072
	Color_SandyBrown           = 0xFFF4A460
	Color_SeaGreen             = 0xFF2E8B57
	Color_SeaShell             = 0xFFFFF5EE
	Color_Sienna               = 0xFFA0522D
	Color_Silver               = 0xFFC0C0C0
	Color_SkyBlue              = 0xFF87CEEB
	Color_SlateBlue            = 0xFF6A5ACD
	Color_SlateGray            = 0xFF708090
	Color_Snow                 = 0xFFFFFAFA
	Color_SpringGreen          = 0xFF00FF7F
	Color_SteelBlue            = 0xFF4682B4
	Color_Tan                  = 0xFFD2B48C
	Color_Teal                 = 0xFF008080
	Color_Thistle              = 0xFFD8BFD8
	Color_Tomato               = 0xFFFF6347
	Color_Transparent          = 0x00FFFFFF
	Color_Turquoise            = 0xFF40E0D0
	Color_Violet               = 0xFFEE82EE
	Color_Wheat                = 0xFFF5DEB3
	Color_White                = 0xFFFFFFFF
	Color_WhiteSmoke           = 0xFFF5F5F5
	Color_Yellow               = 0xFFFFFF00
	Color_YellowGreen          = 0xFF9ACD32
)

// Unit
const (
	UnitWorld      = 0 // 0 -- World coordinate (non-physical unit)
	UnitDisplay    = 1 // 1 -- Variable -- for PageTransform only
	UnitPixel      = 2 // 2 -- Each unit is one device pixel.
	UnitPoint      = 3 // 3 -- Each unit is a printer's point, or 1/72 inch.
	UnitInch       = 4 // 4 -- Each unit is 1 inch.
	UnitDocument   = 5 // 5 -- Each unit is 1/300 inch.
	UnitMillimeter = 6 // 6 -- Each unit is 1 millimeter.
)

const (
	AlphaShift = 24
	RedShift   = 16
	GreenShift = 8
	BlueShift  = 0
)

const (
	AlphaMask = 0xff000000
	RedMask   = 0x00ff0000
	GreenMask = 0x0000ff00
	BlueMask  = 0x000000ff
)

// FontStyle
const (
	FontStyleRegular    = 0
	FontStyleBold       = 1
	FontStyleItalic     = 2
	FontStyleBoldItalic = 3
	FontStyleUnderline  = 4
	FontStyleStrikeout  = 8
)

// QualityMode
const (
	QualityModeInvalid = iota - 1
	QualityModeDefault
	QualityModeLow  // Best performance
	QualityModeHigh // Best rendering quality
)

// Alpha Compositing mode
const (
	CompositingModeSourceOver = iota // 0
	CompositingModeSourceCopy        // 1
)

// Alpha Compositing quality
const (
	CompositingQualityInvalid = iota + QualityModeInvalid
	CompositingQualityDefault
	CompositingQualityHighSpeed
	CompositingQualityHighQuality
	CompositingQualityGammaCorrected
	CompositingQualityAssumeLinear
)

// InterpolationMode
const (
	InterpolationModeInvalid = iota + QualityModeInvalid
	InterpolationModeDefault
	InterpolationModeLowQuality
	InterpolationModeHighQuality
	InterpolationModeBilinear
	InterpolationModeBicubic
	InterpolationModeNearestNeighbor
	InterpolationModeHighQualityBilinear
	InterpolationModeHighQualityBicubic
)

// SmoothingMode
const (
	SmoothingModeInvalid = iota + QualityModeInvalid
	SmoothingModeDefault
	SmoothingModeHighSpeed
	SmoothingModeHighQuality
	SmoothingModeNone
	SmoothingModeAntiAlias

/*
#if (GDIPVER >= 0x0110)
    SmoothingModeAntiAlias8x4 = SmoothingModeAntiAlias,
    SmoothingModeAntiAlias8x8
#endif //(GDIPVER >= 0x0110)
*/
)

// Pixel Format Mode
const (
	PixelOffsetModeInvalid = iota + QualityModeInvalid
	PixelOffsetModeDefault
	PixelOffsetModeHighSpeed
	PixelOffsetModeHighQuality
	PixelOffsetModeNone // No pixel offset
	PixelOffsetModeHalf // Offset by -0.5, -0.5 for fast anti-alias perf
)

// Text Rendering Hint
const (
	TextRenderingHintSystemDefault            = iota // Glyph with system default rendering hint
	TextRenderingHintSingleBitPerPixelGridFit        // Glyph bitmap with hinting
	TextRenderingHintSingleBitPerPixel               // Glyph bitmap without hinting
	TextRenderingHintAntiAliasGridFit                // Glyph anti-alias bitmap with hinting
	TextRenderingHintAntiAlias                       // Glyph anti-alias bitmap without hinting
	TextRenderingHintClearTypeGridFit                // Glyph CT bitmap with hinting
)

// Fill mode constants
const (
	FillModeAlternate = iota // 0
	FillModeWinding          // 1
)

// BrushType
const (
	BrushTypeSolidColor GpBrushType = iota
	BrushTypeHatchFill
	BrushTypeTextureFill
	BrushTypePathGradient
	BrushTypeLinearGradient
)

// LineCap
const (
	LineCapFlat GpLineCap = iota
	LineCapSquare
	LineCapRound
	LineCapTriangle
	LineCapNoAnchor
	LineCapSquareAnchor
	LineCapRoundAnchor
	LineCapDiamondAnchor
	LineCapArrowAnchor
	LineCapCustom
	LineCapAnchorMask
)

// LineJoin
const (
	LineJoinMiter GpLineJoin = iota
	LineJoinBevel
	LineJoinRound
	LineJoinMiterClipped
)

// DashCap
const (
	DashCapFlat GpDashCap = iota
	DashCapRound
	DashCapTriangle
)

// DashStyle
const (
	DashStyleSolid GpDashStyle = iota
	DashStyleDash
	DashStyleDot
	DashStyleDashDot
	DashStyleDashDotDot
	DashStyleCustom
)

// PenAlignment
const (
	PenAlignmentCenter GpPenAlignment = iota
	PenAlignmentInset
)

// MatrixOrder
const (
	MatrixOrderPrepend GpMatrixOrder = iota
	MatrixOrderAppend
)

// PenType
const (
	PenTypeSolidColor GpPenType = iota
	PenTypeHatchFill
	PenTypeTextureFill
	PenTypePathGradient
	PenTypeLinearGradient
	PenTypeUnknown
)

func (s GpStatus) String() string {
	switch s {
	case Ok:
		return "Ok"

	case GenericError:
		return "GenericError"

	case InvalidParameter:
		return "InvalidParameter"

	case OutOfMemory:
		return "OutOfMemory"

	case ObjectBusy:
		return "ObjectBusy"

	case InsufficientBuffer:
		return "InsufficientBuffer"

	case NotImplemented:
		return "NotImplemented"

	case Win32Error:
		return "Win32Error"

	case WrongState:
		return "WrongState"

	case Aborted:
		return "Aborted"

	case FileNotFound:
		return "FileNotFound"

	case ValueOverflow:
		return "ValueOverflow"

	case AccessDenied:
		return "AccessDenied"

	case UnknownImageFormat:
		return "UnknownImageFormat"

	case FontFamilyNotFound:
		return "FontFamilyNotFound"

	case FontStyleNotFound:
		return "FontStyleNotFound"

	case NotTrueTypeFont:
		return "NotTrueTypeFont"

	case UnsupportedGdiplusVersion:
		return "UnsupportedGdiplusVersion"

	case GdiplusNotInitialized:
		return "GdiplusNotInitialized"

	case PropertyNotFound:
		return "PropertyNotFound"

	case PropertyNotSupported:
		return "PropertyNotSupported"

	case ProfileNotFound:
		return "ProfileNotFound"
	}

	return "Unknown Status Value"
}

type ARGB uint32

type GpStatus int32
type GpGraphics struct{}
type GpImage struct{}
type GpPen struct{}
type GpBrush struct{}
type GpSolidFill struct{ GpBrush }
type GpStringFormat struct{}
type GpFont struct{}
type GpFontFamily struct{}
type GpFontCollection struct{}
type GpRegion struct{}
type GpPath struct{}
type GpUnit int32
type GpBitmap GpImage
type GpMatrix struct{}
type GpCustomLineCap struct{}

// Enum types
type GpBrushType int32
type GpPenType int32
type GpLineCap int32
type GpLineJoin int32
type GpDashCap int32
type GpDashStyle int32
type GpPenAlignment int32
type GpMatrixOrder int32

type BrushType GpBrushType
type PenType GpPenType
type LineCap GpLineCap
type LineJoin GpLineJoin
type DashCap GpDashCap
type DashStyle GpDashStyle
type PenAlignment GpPenAlignment
type MatrixOrder GpMatrixOrder

type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread int32
	SuppressExternalCodecs   int32
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type RectF struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type PointF struct {
	X float32
	Y float32
}

type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type Point struct {
	X int32
	Y int32
}

type EncoderParameter struct {
	Guid           ole.GUID
	NumberOfValues uint32
	TypeAPI        uint32
	Value          uintptr
}

type EncoderParameters struct {
	Count     uint32
	Parameter [1]EncoderParameter
}

// In-memory pixel data formats:
// bits 0-7 = format index
// bits 8-15 = pixel size (in bits)
// bits 16-23 = flags
// bits 24-31 = reserved

type PixelFormat int32

const (
	PixelFormatIndexed   = 0x00010000 // Indexes into a palette
	PixelFormatGDI       = 0x00020000 // Is a GDI-supported format
	PixelFormatAlpha     = 0x00040000 // Has an alpha component
	PixelFormatPAlpha    = 0x00080000 // Pre-multiplied alpha
	PixelFormatExtended  = 0x00100000 // Extended color 16 bits/channel
	PixelFormatCanonical = 0x00200000

	PixelFormatUndefined = 0
	PixelFormatDontCare  = 0

	PixelFormat1bppIndexed    = (1 | (1 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat4bppIndexed    = (2 | (4 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat8bppIndexed    = (3 | (8 << 8) | PixelFormatIndexed | PixelFormatGDI)
	PixelFormat16bppGrayScale = (4 | (16 << 8) | PixelFormatExtended)
	PixelFormat16bppRGB555    = (5 | (16 << 8) | PixelFormatGDI)
	PixelFormat16bppRGB565    = (6 | (16 << 8) | PixelFormatGDI)
	PixelFormat16bppARGB1555  = (7 | (16 << 8) | PixelFormatAlpha | PixelFormatGDI)
	PixelFormat24bppRGB       = (8 | (24 << 8) | PixelFormatGDI)
	PixelFormat32bppRGB       = (9 | (32 << 8) | PixelFormatGDI)
	PixelFormat32bppARGB      = (10 | (32 << 8) | PixelFormatAlpha | PixelFormatGDI | PixelFormatCanonical)
	PixelFormat32bppPARGB     = (11 | (32 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatGDI)
	PixelFormat48bppRGB       = (12 | (48 << 8) | PixelFormatExtended)
	PixelFormat64bppARGB      = (13 | (64 << 8) | PixelFormatAlpha | PixelFormatCanonical | PixelFormatExtended)
	PixelFormat64bppPARGB     = (14 | (64 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatExtended)
	PixelFormat32bppCMYK      = (15 | (32 << 8))
	PixelFormatMax            = 16
)

func NewRect(x, y, width, height int32) *Rect {
	return &Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func NewRectF(x, y, width, height float32) *RectF {
	return &RectF{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (rect *Rect) Left() int32 {
	return rect.X
}

func (rect *Rect) Top() int32 {
	return rect.Y
}

func (rect *RectF) Left() float32 {
	return rect.X
}

func (rect *RectF) Top() float32 {
	return rect.Y
}

func (rect *Rect) Right() int32 {
	return rect.X + rect.Width
}

func (rect *Rect) Bottom() int32 {
	return rect.Y + rect.Height
}

func (rect *RectF) Right() float32 {
	return rect.X + rect.Width
}

func (rect *RectF) Bottom() float32 {
	return rect.Y + rect.Height
}

var (
	// Library
	libgdiplus *windows.LazyDLL

	// Functions
	gdiplusShutdown *windows.LazyProc
	gdiplusStartup  *windows.LazyProc
	// Graphics
	gdipCreateFromHDC          *windows.LazyProc
	gdipCreateFromHDC2         *windows.LazyProc
	gdipCreateFromHWND         *windows.LazyProc
	gdipCreateFromHWNDICM      *windows.LazyProc
	gdipDeleteGraphics         *windows.LazyProc
	gdipGetDC                  *windows.LazyProc
	gdipReleaseDC              *windows.LazyProc
	gdipSetInterpolationMode   *windows.LazyProc
	gdipSetSmoothingMode       *windows.LazyProc
	gdipSetPixelOffsetMode     *windows.LazyProc
	gdipSetCompositingQuality  *windows.LazyProc
	gdipSetCompositingMode     *windows.LazyProc
	gdipSetRenderingOrigin     *windows.LazyProc
	gdipSetTextRenderingHint   *windows.LazyProc
	gdipGraphicsClear          *windows.LazyProc
	gdipDrawLine               *windows.LazyProc
	gdipDrawLineI              *windows.LazyProc
	gdipDrawArc                *windows.LazyProc
	gdipDrawArcI               *windows.LazyProc
	gdipDrawBezier             *windows.LazyProc
	gdipDrawBezierI            *windows.LazyProc
	gdipDrawRectangle          *windows.LazyProc
	gdipDrawRectangleI         *windows.LazyProc
	gdipDrawEllipse            *windows.LazyProc
	gdipDrawEllipseI           *windows.LazyProc
	gdipDrawPie                *windows.LazyProc
	gdipDrawPieI               *windows.LazyProc
	gdipDrawPolygon            *windows.LazyProc
	gdipDrawPolygonI           *windows.LazyProc
	gdipDrawPath               *windows.LazyProc
	gdipDrawString             *windows.LazyProc
	gdipDrawImage              *windows.LazyProc
	gdipDrawImageI             *windows.LazyProc
	gdipDrawImageRect          *windows.LazyProc
	gdipDrawImageRectI         *windows.LazyProc
	gdipFillRectangle          *windows.LazyProc
	gdipFillRectangleI         *windows.LazyProc
	gdipFillPolygon            *windows.LazyProc
	gdipFillPolygonI           *windows.LazyProc
	gdipFillPath               *windows.LazyProc
	gdipFillEllipse            *windows.LazyProc
	gdipFillEllipseI           *windows.LazyProc
	gdipMeasureString          *windows.LazyProc
	gdipMeasureCharacterRanges *windows.LazyProc
	// Pen
	gdipCreatePen1            *windows.LazyProc
	gdipCreatePen2            *windows.LazyProc
	gdipClonePen              *windows.LazyProc
	gdipDeletePen             *windows.LazyProc
	gdipSetPenWidth           *windows.LazyProc
	gdipGetPenWidth           *windows.LazyProc
	gdipSetPenLineCap197819   *windows.LazyProc
	gdipSetPenStartCap        *windows.LazyProc
	gdipSetPenEndCap          *windows.LazyProc
	gdipSetPenDashCap197819   *windows.LazyProc
	gdipGetPenStartCap        *windows.LazyProc
	gdipGetPenEndCap          *windows.LazyProc
	gdipGetPenDashCap197819   *windows.LazyProc
	gdipSetPenLineJoin        *windows.LazyProc
	gdipGetPenLineJoin        *windows.LazyProc
	gdipSetPenCustomStartCap  *windows.LazyProc
	gdipGetPenCustomStartCap  *windows.LazyProc
	gdipSetPenCustomEndCap    *windows.LazyProc
	gdipGetPenCustomEndCap    *windows.LazyProc
	gdipSetPenMiterLimit      *windows.LazyProc
	gdipGetPenMiterLimit      *windows.LazyProc
	gdipSetPenMode            *windows.LazyProc
	gdipGetPenMode            *windows.LazyProc
	gdipSetPenTransform       *windows.LazyProc
	gdipGetPenTransform       *windows.LazyProc
	gdipResetPenTransform     *windows.LazyProc
	gdipMultiplyPenTransform  *windows.LazyProc
	gdipTranslatePenTransform *windows.LazyProc
	gdipScalePenTransform     *windows.LazyProc
	gdipRotatePenTransform    *windows.LazyProc
	gdipSetPenColor           *windows.LazyProc
	gdipGetPenColor           *windows.LazyProc
	gdipSetPenBrushFill       *windows.LazyProc
	gdipGetPenBrushFill       *windows.LazyProc
	gdipGetPenFillType        *windows.LazyProc
	gdipGetPenDashStyle       *windows.LazyProc
	gdipSetPenDashStyle       *windows.LazyProc
	gdipGetPenDashOffset      *windows.LazyProc
	gdipSetPenDashOffset      *windows.LazyProc
	gdipGetPenDashCount       *windows.LazyProc
	gdipSetPenDashArray       *windows.LazyProc
	gdipGetPenDashArray       *windows.LazyProc
	gdipGetPenCompoundCount   *windows.LazyProc
	gdipSetPenCompoundArray   *windows.LazyProc
	gdipGetPenCompoundArray   *windows.LazyProc
	// Brush
	gdipCloneBrush   *windows.LazyProc
	gdipDeleteBrush  *windows.LazyProc
	gdipGetBrushType *windows.LazyProc
	// Solid Brush
	gdipCreateSolidFill   *windows.LazyProc
	gdipSetSolidFillColor *windows.LazyProc
	gdipGetSolidFillColor *windows.LazyProc
	// Image
	gdipLoadImageFromFile       *windows.LazyProc
	gdipSaveImageToFile         *windows.LazyProc
	gdipGetImageWidth           *windows.LazyProc
	gdipGetImageHeight          *windows.LazyProc
	gdipGetImageGraphicsContext *windows.LazyProc
	gdipDisposeImage            *windows.LazyProc
	// Bitmap
	gdipCreateBitmapFromScan0   *windows.LazyProc
	gdipCreateBitmapFromFile    *windows.LazyProc
	gdipCreateBitmapFromHBITMAP *windows.LazyProc
	gdipCreateHBITMAPFromBitmap *windows.LazyProc
	// Font
	gdipCreateFontFromDC           *windows.LazyProc
	gdipCreateFont                 *windows.LazyProc
	gdipDeleteFont                 *windows.LazyProc
	gdipNewInstalledFontCollection *windows.LazyProc
	gdipCreateFontFamilyFromName   *windows.LazyProc
	gdipDeleteFontFamily           *windows.LazyProc
	// StringFormat
	gdipCreateStringFormat                *windows.LazyProc
	gdipDeleteStringFormat                *windows.LazyProc
	gdipStringFormatGetGenericTypographic *windows.LazyProc
	// Path
	gdipCreatePath       *windows.LazyProc
	gdipDeletePath       *windows.LazyProc
	gdipAddPathArc       *windows.LazyProc
	gdipAddPathArcI      *windows.LazyProc
	gdipAddPathLine      *windows.LazyProc
	gdipAddPathLineI     *windows.LazyProc
	gdipClosePathFigure  *windows.LazyProc
	gdipClosePathFigures *windows.LazyProc
)

var (
	token uintptr
)

func init() {
	// Library
	libgdiplus = windows.NewLazySystemDLL("gdiplus.dll")
	// Functions
	gdiplusShutdown = libgdiplus.NewProc("GdiplusShutdown")
	gdiplusStartup = libgdiplus.NewProc("GdiplusStartup")

	// Functions
	gdipCreateBitmapFromFile = libgdiplus.NewProc("GdipCreateBitmapFromFile")
	gdipCreateBitmapFromHBITMAP = libgdiplus.NewProc("GdipCreateBitmapFromHBITMAP")
	gdipCreateHBITMAPFromBitmap = libgdiplus.NewProc("GdipCreateHBITMAPFromBitmap")
	gdipDisposeImage = libgdiplus.NewProc("GdipDisposeImage")

	gdipCreateFromHDC = libgdiplus.NewProc("GdipCreateFromHDC")
	gdipCreateFromHDC2 = libgdiplus.NewProc("GdipCreateFromHDC2")
	gdipCreateFromHWND = libgdiplus.NewProc("GdipCreateFromHWND")
	gdipCreateFromHWNDICM = libgdiplus.NewProc("GdipCreateFromHWNDICM")
	gdipDeleteGraphics = libgdiplus.NewProc("GdipDeleteGraphics")
	gdipGetDC = libgdiplus.NewProc("GdipGetDC")
	gdipReleaseDC = libgdiplus.NewProc("GdipReleaseDC")
	gdipSetCompositingMode = libgdiplus.NewProc("GdipSetCompositingMode")
	gdipSetRenderingOrigin = libgdiplus.NewProc("GdipSetRenderingOrigin")
	gdipSetCompositingQuality = libgdiplus.NewProc("GdipSetCompositingQuality")
	gdipSetSmoothingMode = libgdiplus.NewProc("GdipSetSmoothingMode")
	gdipSetPixelOffsetMode = libgdiplus.NewProc("GdipSetPixelOffsetMode")
	gdipSetInterpolationMode = libgdiplus.NewProc("GdipSetInterpolationMode")
	gdipSetTextRenderingHint = libgdiplus.NewProc("GdipSetTextRenderingHint")
	gdipGraphicsClear = libgdiplus.NewProc("GdipGraphicsClear")
	gdipDrawLine = libgdiplus.NewProc("GdipDrawLine")
	gdipDrawLineI = libgdiplus.NewProc("GdipDrawLineI")
	gdipDrawArc = libgdiplus.NewProc("GdipDrawArc")
	gdipDrawArcI = libgdiplus.NewProc("GdipDrawArcI")
	gdipDrawBezier = libgdiplus.NewProc("GdipDrawBezier")
	gdipDrawBezierI = libgdiplus.NewProc("GdipDrawBezierI")
	gdipDrawRectangle = libgdiplus.NewProc("GdipDrawRectangle")
	gdipDrawRectangleI = libgdiplus.NewProc("GdipDrawRectangleI")
	gdipDrawEllipse = libgdiplus.NewProc("GdipDrawEllipse")
	gdipDrawEllipseI = libgdiplus.NewProc("GdipDrawEllipseI")
	gdipDrawPie = libgdiplus.NewProc("GdipDrawPie")
	gdipDrawPieI = libgdiplus.NewProc("GdipDrawPieI")
	gdipDrawPolygonI = libgdiplus.NewProc("GdipDrawPolygonI")
	gdipDrawPolygon = libgdiplus.NewProc("GdipDrawPolygon")
	gdipDrawPath = libgdiplus.NewProc("GdipDrawPath")
	gdipDrawString = libgdiplus.NewProc("GdipDrawString")
	gdipDrawImage = libgdiplus.NewProc("GdipDrawImage")
	gdipDrawImageI = libgdiplus.NewProc("GdipDrawImageI")
	gdipDrawImageRect = libgdiplus.NewProc("GdipDrawImageRect")
	gdipDrawImageRectI = libgdiplus.NewProc("GdipDrawImageRectI")
	gdipFillRectangle = libgdiplus.NewProc("GdipFillRectangle")
	gdipFillRectangleI = libgdiplus.NewProc("GdipFillRectangleI")
	gdipFillPolygon = libgdiplus.NewProc("GdipFillPolygon")
	gdipFillPolygonI = libgdiplus.NewProc("GdipFillPolygonI")
	gdipFillPath = libgdiplus.NewProc("GdipFillPath")
	gdipFillEllipse = libgdiplus.NewProc("GdipFillEllipse")
	gdipFillEllipseI = libgdiplus.NewProc("GdipFillEllipseI")
	gdipMeasureString = libgdiplus.NewProc("GdipMeasureString")
	gdipMeasureCharacterRanges = libgdiplus.NewProc("GdipMeasureCharacterRanges")
	// Pen
	gdipCreatePen1 = libgdiplus.NewProc("GdipCreatePen1")
	gdipCreatePen2 = libgdiplus.NewProc("GdipCreatePen2")
	gdipClonePen = libgdiplus.NewProc("GdipClonePen")
	gdipDeletePen = libgdiplus.NewProc("GdipDeletePen")
	gdipSetPenWidth = libgdiplus.NewProc("GdipSetPenWidth")
	gdipGetPenWidth = libgdiplus.NewProc("GdipGetPenWidth")
	gdipSetPenLineCap197819 = libgdiplus.NewProc("GdipSetPenLineCap197819")
	gdipSetPenStartCap = libgdiplus.NewProc("GdipSetPenStartCap")
	gdipSetPenEndCap = libgdiplus.NewProc("GdipSetPenEndCap")
	gdipSetPenDashCap197819 = libgdiplus.NewProc("GdipSetPenDashCap197819")
	gdipGetPenStartCap = libgdiplus.NewProc("GdipGetPenStartCap")
	gdipGetPenEndCap = libgdiplus.NewProc("GdipGetPenEndCap")
	gdipGetPenDashCap197819 = libgdiplus.NewProc("GdipGetPenDashCap197819")
	gdipSetPenLineJoin = libgdiplus.NewProc("GdipSetPenLineJoin")
	gdipGetPenLineJoin = libgdiplus.NewProc("GdipGetPenLineJoin")
	gdipSetPenCustomStartCap = libgdiplus.NewProc("GdipSetPenCustomStartCap")
	gdipGetPenCustomStartCap = libgdiplus.NewProc("GdipGetPenCustomStartCap")
	gdipSetPenCustomEndCap = libgdiplus.NewProc("GdipSetPenCustomEndCap")
	gdipGetPenCustomEndCap = libgdiplus.NewProc("GdipGetPenCustomEndCap")
	gdipSetPenMiterLimit = libgdiplus.NewProc("GdipSetPenMiterLimit")
	gdipGetPenMiterLimit = libgdiplus.NewProc("GdipGetPenMiterLimit")
	gdipSetPenMode = libgdiplus.NewProc("GdipSetPenMode")
	gdipGetPenMode = libgdiplus.NewProc("GdipGetPenMode")
	gdipSetPenTransform = libgdiplus.NewProc("GdipSetPenTransform")
	gdipGetPenTransform = libgdiplus.NewProc("GdipGetPenTransform")
	gdipResetPenTransform = libgdiplus.NewProc("GdipResetPenTransform")
	gdipMultiplyPenTransform = libgdiplus.NewProc("GdipMultiplyPenTransform")
	gdipTranslatePenTransform = libgdiplus.NewProc("GdipTranslatePenTransform")
	gdipScalePenTransform = libgdiplus.NewProc("GdipScalePenTransform")
	gdipRotatePenTransform = libgdiplus.NewProc("GdipRotatePenTransform")
	gdipSetPenColor = libgdiplus.NewProc("GdipSetPenColor")
	gdipGetPenColor = libgdiplus.NewProc("GdipGetPenColor")
	gdipSetPenBrushFill = libgdiplus.NewProc("GdipSetPenBrushFill")
	gdipGetPenBrushFill = libgdiplus.NewProc("GdipGetPenBrushFill")
	gdipGetPenFillType = libgdiplus.NewProc("GdipGetPenFillType")
	gdipGetPenDashStyle = libgdiplus.NewProc("GdipGetPenDashStyle")
	gdipSetPenDashStyle = libgdiplus.NewProc("GdipSetPenDashStyle")
	gdipGetPenDashOffset = libgdiplus.NewProc("GdipGetPenDashOffset")
	gdipSetPenDashOffset = libgdiplus.NewProc("GdipSetPenDashOffset")
	gdipGetPenDashCount = libgdiplus.NewProc("GdipGetPenDashCount")
	gdipSetPenDashArray = libgdiplus.NewProc("GdipSetPenDashArray")
	gdipGetPenDashArray = libgdiplus.NewProc("GdipGetPenDashArray")
	gdipGetPenCompoundCount = libgdiplus.NewProc("GdipGetPenCompoundCount")
	gdipSetPenCompoundArray = libgdiplus.NewProc("GdipSetPenCompoundArray")
	gdipGetPenCompoundArray = libgdiplus.NewProc("GdipGetPenCompoundArray")
	// Brush
	gdipCloneBrush = libgdiplus.NewProc("GdipCloneBrush")
	gdipDeleteBrush = libgdiplus.NewProc("GdipDeleteBrush")
	gdipGetBrushType = libgdiplus.NewProc("GdipGetBrushType")
	// Solid Brush
	gdipCreateSolidFill = libgdiplus.NewProc("GdipCreateSolidFill")
	gdipSetSolidFillColor = libgdiplus.NewProc("GdipSetSolidFillColor")
	gdipGetSolidFillColor = libgdiplus.NewProc("GdipGetSolidFillColor")
	// Image
	gdipLoadImageFromFile = libgdiplus.NewProc("GdipLoadImageFromFile")
	gdipSaveImageToFile = libgdiplus.NewProc("GdipSaveImageToFile")
	gdipGetImageWidth = libgdiplus.NewProc("GdipGetImageWidth")
	gdipGetImageHeight = libgdiplus.NewProc("GdipGetImageHeight")
	gdipGetImageGraphicsContext = libgdiplus.NewProc("GdipGetImageGraphicsContext")
	gdipDisposeImage = libgdiplus.NewProc("GdipDisposeImage")
	// Bitmap
	gdipCreateBitmapFromScan0 = libgdiplus.NewProc("GdipCreateBitmapFromScan0")
	gdipCreateBitmapFromFile = libgdiplus.NewProc("GdipCreateBitmapFromFile")
	gdipCreateBitmapFromHBITMAP = libgdiplus.NewProc("GdipCreateBitmapFromHBITMAP")
	gdipCreateHBITMAPFromBitmap = libgdiplus.NewProc("GdipCreateHBITMAPFromBitmap")
	// Font
	gdipCreateFontFromDC = libgdiplus.NewProc("GdipCreateFontFromDC")
	gdipCreateFont = libgdiplus.NewProc("GdipCreateFont")
	gdipDeleteFont = libgdiplus.NewProc("GdipDeleteFont")
	gdipNewInstalledFontCollection = libgdiplus.NewProc("GdipNewInstalledFontCollection")
	gdipCreateFontFamilyFromName = libgdiplus.NewProc("GdipCreateFontFamilyFromName")
	gdipDeleteFontFamily = libgdiplus.NewProc("GdipDeleteFontFamily")
	// StringFormat
	gdipCreateStringFormat = libgdiplus.NewProc("GdipCreateStringFormat")
	gdipDeleteStringFormat = libgdiplus.NewProc("GdipDeleteStringFormat")
	gdipStringFormatGetGenericTypographic = libgdiplus.NewProc("GdipStringFormatGetGenericTypographic")
	// Path
	gdipCreatePath = libgdiplus.NewProc("GdipCreatePath")
	gdipDeletePath = libgdiplus.NewProc("GdipDeletePath")
	gdipAddPathArc = libgdiplus.NewProc("GdipAddPathArc")
	gdipAddPathArcI = libgdiplus.NewProc("GdipAddPathArcI")
	gdipAddPathLine = libgdiplus.NewProc("GdipAddPathLine")
	gdipAddPathLineI = libgdiplus.NewProc("GdipAddPathLineI")
	gdipClosePathFigure = libgdiplus.NewProc("GdipClosePathFigure")
	gdipClosePathFigures = libgdiplus.NewProc("GdipClosePathFigures")

}

func GdiplusShutdown() {
	syscall.Syscall(gdiplusShutdown.Addr(), 1,
		token,
		0,
		0)
}

func GdiplusStartup(input *GdiplusStartupInput, output *GdiplusStartupOutput) GpStatus {
	ret, _, _ := syscall.Syscall(gdiplusStartup.Addr(), 3,
		uintptr(unsafe.Pointer(&token)),
		uintptr(unsafe.Pointer(input)),
		uintptr(unsafe.Pointer(output)))

	return GpStatus(ret)
}

// Graphics
func GdipCreateFromHDC(hdc HDC, graphics **GpGraphics) GpStatus {
	ret, _, _ := gdipCreateFromHDC.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(graphics)))
	return GpStatus(ret)
}

func GdipCreateFromHDC2(hdc HDC, hDevice HANDLE, graphics **GpGraphics) GpStatus {
	ret, _, _ := gdipCreateFromHDC2.Call(
		uintptr(hdc),
		uintptr(hDevice),
		uintptr(unsafe.Pointer(graphics)))
	return GpStatus(ret)
}

func GdipCreateFromHWND(hwnd HWND, graphics **GpGraphics) GpStatus {
	ret, _, _ := gdipCreateFromHWND.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(graphics)))
	return GpStatus(ret)
}

func GdipCreateFromHWNDICM(hwnd HWND, graphics **GpGraphics) GpStatus {
	ret, _, _ := gdipCreateFromHWNDICM.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(graphics)))
	return GpStatus(ret)
}

func GdipDeleteGraphics(graphics *GpGraphics) GpStatus {
	ret, _, _ := gdipDeleteGraphics.Call(uintptr(unsafe.Pointer(graphics)))
	return GpStatus(ret)
}

func GdipGetDC(graphics *GpGraphics, hdc *HDC) GpStatus {
	ret, _, _ := gdipGetDC.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(hdc)))
	return GpStatus(ret)
}

func GdipReleaseDC(graphics *GpGraphics, hdc HDC) GpStatus {
	ret, _, _ := gdipReleaseDC.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(hdc))
	return GpStatus(ret)
}

func GdipSetCompositingMode(graphics *GpGraphics, mode int32) GpStatus {
	ret, _, _ := gdipSetCompositingMode.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(mode))
	return GpStatus(ret)
}

func GdipSetRenderingOrigin(graphics *GpGraphics, x, y int32) GpStatus {
	ret, _, _ := gdipSetRenderingOrigin.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(x),
		uintptr(y))
	return GpStatus(ret)
}

func GdipSetCompositingQuality(graphics *GpGraphics, quality int32) GpStatus {
	ret, _, _ := gdipSetCompositingQuality.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(quality))
	return GpStatus(ret)
}

func GdipSetInterpolationMode(graphics *GpGraphics, mode int32) GpStatus {
	ret, _, _ := gdipSetInterpolationMode.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(mode))
	return GpStatus(ret)
}

func GdipSetPixelOffsetMode(graphics *GpGraphics, mode int32) GpStatus {
	ret, _, _ := gdipSetPixelOffsetMode.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(mode))
	return GpStatus(ret)
}

func GdipSetSmoothingMode(graphics *GpGraphics, mode int32) GpStatus {
	ret, _, _ := gdipSetSmoothingMode.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(mode))
	return GpStatus(ret)
}

func GdipSetTextRenderingHint(graphics *GpGraphics, hint int32) GpStatus {
	ret, _, _ := gdipSetTextRenderingHint.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(hint))
	return GpStatus(ret)
}

func GdipGraphicsClear(graphics *GpGraphics, color ARGB) GpStatus {
	ret, _, _ := gdipGraphicsClear.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(color))
	return GpStatus(ret)
}

func GdipDrawLine(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2 float32) GpStatus {
	ret, _, _ := gdipDrawLine.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x1)),
		uintptr(math.Float32bits(y1)),
		uintptr(math.Float32bits(x2)),
		uintptr(math.Float32bits(y2)))
	return GpStatus(ret)
}

func GdipDrawLineI(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2 int32) GpStatus {
	ret, _, _ := gdipDrawLineI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
	return GpStatus(ret)
}

func GdipDrawArc(graphics *GpGraphics, pen *GpPen, x, y, width, height, startAngle, sweepAngle float32) GpStatus {
	ret, _, _ := gdipDrawArc.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return GpStatus(ret)
}

func GdipDrawArcI(graphics *GpGraphics, pen *GpPen, x, y, width, height int32, startAngle, sweepAngle float32) GpStatus {
	ret, _, _ := gdipDrawArcI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return GpStatus(ret)
}

func GdipDrawBezier(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2, x3, y3, x4, y4 float32) GpStatus {
	ret, _, _ := gdipDrawBezier.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x1)),
		uintptr(math.Float32bits(y1)),
		uintptr(math.Float32bits(x2)),
		uintptr(math.Float32bits(y2)),
		uintptr(math.Float32bits(x3)),
		uintptr(math.Float32bits(y3)),
		uintptr(math.Float32bits(x4)),
		uintptr(math.Float32bits(y4)))
	return GpStatus(ret)
}

func GdipDrawBezierI(graphics *GpGraphics, pen *GpPen, x1, y1, x2, y2, x3, y3, x4, y4 int32) GpStatus {
	ret, _, _ := gdipDrawBezierI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(x3),
		uintptr(y3),
		uintptr(x4),
		uintptr(y4))
	return GpStatus(ret)
}

func GdipDrawRectangle(graphics *GpGraphics, pen *GpPen, x, y, width, height float32) GpStatus {
	ret, _, _ := gdipDrawRectangle.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return GpStatus(ret)
}

func GdipDrawRectangleI(graphics *GpGraphics, pen *GpPen, x, y, width, height int32) GpStatus {
	ret, _, _ := gdipDrawRectangleI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return GpStatus(ret)
}

func GdipDrawEllipse(graphics *GpGraphics, pen *GpPen, x, y, width, height float32) GpStatus {
	ret, _, _ := gdipDrawEllipse.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return GpStatus(ret)
}

func GdipDrawEllipseI(graphics *GpGraphics, pen *GpPen, x, y, width, height int32) GpStatus {
	ret, _, _ := gdipDrawEllipseI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return GpStatus(ret)
}

func GdipDrawPie(graphics *GpGraphics, pen *GpPen, x, y, width, height, startAngle, sweepAngle float32) GpStatus {
	ret, _, _ := gdipDrawPie.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return GpStatus(ret)
}

func GdipDrawPieI(graphics *GpGraphics, pen *GpPen, x, y, width, height int32, startAngle, sweepAngle float32) GpStatus {
	ret, _, _ := gdipDrawPieI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return GpStatus(ret)
}

func GdipDrawPolygon(graphics *GpGraphics, pen *GpPen, points *PointF, count int32) GpStatus {
	ret, _, _ := gdipDrawPolygon.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return GpStatus(ret)
}

func GdipDrawPolygonI(graphics *GpGraphics, pen *GpPen, points *Point, count int32) GpStatus {
	ret, _, _ := gdipDrawPolygonI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return GpStatus(ret)
}

func GdipDrawPath(graphics *GpGraphics, pen *GpPen, path *GpPath) GpStatus {
	ret, _, _ := gdipDrawPath.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(path)))
	return GpStatus(ret)
}

func GdipDrawString(graphics *GpGraphics, text *uint16, length int32, font *GpFont, layoutRect *RectF, stringFormat *GpStringFormat, brush *GpBrush) GpStatus {
	ret, _, _ := gdipDrawString.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(text)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(stringFormat)),
		uintptr(unsafe.Pointer(brush)))
	return GpStatus(ret)
}

func GdipDrawImage(graphics *GpGraphics, image *GpImage, x, y float32) GpStatus {
	ret, _, _ := gdipDrawImage.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)))
	return GpStatus(ret)
}

func GdipDrawImageI(graphics *GpGraphics, image *GpImage, x, y int32) GpStatus {
	ret, _, _ := gdipDrawImageI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(x),
		uintptr(y))
	return GpStatus(ret)
}

func GdipDrawImageRect(graphics *GpGraphics, image *GpImage, x, y, width, height float32) GpStatus {
	ret, _, _ := gdipDrawImageRect.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return GpStatus(ret)
}

func GdipDrawImageRectI(graphics *GpGraphics, image *GpImage, x, y, width, height int32) GpStatus {
	ret, _, _ := gdipDrawImageRectI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return GpStatus(ret)
}

func GdipFillRectangle(graphics *GpGraphics, brush *GpBrush, x, y, width, height float32) GpStatus {
	ret, _, _ := gdipFillRectangle.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return GpStatus(ret)
}

func GdipFillRectangleI(graphics *GpGraphics, brush *GpBrush, x, y, width, height int32) GpStatus {
	ret, _, _ := gdipFillRectangleI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return GpStatus(ret)
}

func GdipFillEllipse(graphics *GpGraphics, brush *GpBrush, x, y, width, height float32) GpStatus {
	ret, _, _ := gdipFillEllipse.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return GpStatus(ret)
}

func GdipFillEllipseI(graphics *GpGraphics, brush *GpBrush, x, y, width, height int32) GpStatus {
	ret, _, _ := gdipFillEllipseI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return GpStatus(ret)
}

func GdipFillPolygon(graphics *GpGraphics, brush *GpBrush, points *PointF, count int32, fillMode int32) GpStatus {
	ret, _, _ := gdipFillPolygon.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(fillMode))
	return GpStatus(ret)
}

func GdipFillPolygonI(graphics *GpGraphics, brush *GpBrush, points *Point, count int32, fillMode int32) GpStatus {
	ret, _, _ := gdipFillPolygonI.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(fillMode))
	return GpStatus(ret)
}

func GdipFillPath(graphics *GpGraphics, brush *GpBrush, path *GpPath) GpStatus {
	ret, _, _ := gdipFillPath.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(path)))
	return GpStatus(ret)
}

func GdipMeasureString(
	graphics *GpGraphics, text *uint16,
	length int32, font *GpFont, layoutRect *RectF,
	stringFormat *GpStringFormat, boundingBox *RectF,
	codepointsFitted *int32, linesFilled *int32) GpStatus {

	ret, _, _ := gdipMeasureString.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(text)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(stringFormat)),
		uintptr(unsafe.Pointer(boundingBox)),
		uintptr(unsafe.Pointer(codepointsFitted)),
		uintptr(unsafe.Pointer(linesFilled)))
	return GpStatus(ret)
}

func GdipMeasureCharacterRanges(
	graphics *GpGraphics, text *uint16,
	length int32, font *GpFont, layoutRect *RectF,
	stringFormat *GpStringFormat, regionCount int32,
	regions **GpRegion) GpStatus {

	ret, _, _ := gdipMeasureCharacterRanges.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(text)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(stringFormat)),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(regions)))
	return GpStatus(ret)
}

// Pen
func GdipCreatePen1(color ARGB, width float32, unit GpUnit, pen **GpPen) GpStatus {
	ret, _, _ := gdipCreatePen1.Call(
		uintptr(color),
		uintptr(math.Float32bits(width)),
		uintptr(unit),
		uintptr(unsafe.Pointer(pen)))
	return GpStatus(ret)
}

func GdipCreatePen2(brush *GpBrush, width float32, unit GpUnit, pen **GpPen) GpStatus {
	ret, _, _ := gdipCreatePen2.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(width)),
		uintptr(unit),
		uintptr(unsafe.Pointer(pen)))
	return GpStatus(ret)
}

func GdipClonePen(pen *GpPen, clonepen **GpPen) GpStatus {
	ret, _, _ := gdipClonePen.Call(uintptr(unsafe.Pointer(pen)), uintptr(unsafe.Pointer(clonepen)))
	return GpStatus(ret)
}

func GdipDeletePen(pen *GpPen) GpStatus {
	ret, _, _ := gdipDeletePen.Call(uintptr(unsafe.Pointer(pen)))
	return GpStatus(ret)
}

func GdipSetPenWidth(pen *GpPen, width float32) GpStatus {
	ret, _, _ := gdipSetPenWidth.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(width)))
	return GpStatus(ret)
}

func GdipGetPenWidth(pen *GpPen, width *float32) GpStatus {
	var penWidth uint32
	ret, _, _ := gdipGetPenWidth.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(&penWidth)))
	*width = math.Float32frombits(penWidth)
	return GpStatus(ret)
}

func GdipSetPenLineCap197819(pen *GpPen, startCap, endCap GpLineCap, dashCap GpDashCap) GpStatus {
	ret, _, _ := gdipSetPenLineCap197819.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(startCap),
		uintptr(endCap),
		uintptr(dashCap))
	return GpStatus(ret)
}
func GdipSetPenStartCap(pen *GpPen, startCap GpLineCap) GpStatus {
	ret, _, _ := gdipSetPenStartCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(startCap))
	return GpStatus(ret)
}
func GdipSetPenEndCap(pen *GpPen, endCap GpLineCap) GpStatus {
	ret, _, _ := gdipSetPenEndCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(endCap))
	return GpStatus(ret)
}
func GdipSetPenDashCap197819(pen *GpPen, dashCap GpDashCap) GpStatus {
	ret, _, _ := gdipSetPenDashCap197819.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(dashCap))
	return GpStatus(ret)
}
func GdipGetPenStartCap(pen *GpPen, startCap *GpLineCap) GpStatus {
	ret, _, _ := gdipGetPenStartCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(startCap)))
	return GpStatus(ret)
}
func GdipGetPenEndCap(pen *GpPen, endCap *GpLineCap) GpStatus {
	ret, _, _ := gdipGetPenEndCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(endCap)))
	return GpStatus(ret)
}
func GdipGetPenDashCap197819(pen *GpPen, dashCap *GpDashCap) GpStatus {
	ret, _, _ := gdipGetPenDashCap197819.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dashCap)))
	return GpStatus(ret)
}
func GdipSetPenLineJoin(pen *GpPen, lineJoin GpLineJoin) GpStatus {
	ret, _, _ := gdipSetPenLineJoin.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(lineJoin))
	return GpStatus(ret)
}
func GdipGetPenLineJoin(pen *GpPen, lineJoin *GpLineJoin) GpStatus {
	ret, _, _ := gdipGetPenLineJoin.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(lineJoin)))
	return GpStatus(ret)
}
func GdipSetPenCustomStartCap(pen *GpPen, customCap *GpCustomLineCap) GpStatus {
	ret, _, _ := gdipSetPenCustomStartCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return GpStatus(ret)
}
func GdipGetPenCustomStartCap(pen *GpPen, customCap **GpCustomLineCap) GpStatus {
	ret, _, _ := gdipGetPenCustomStartCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return GpStatus(ret)
}
func GdipSetPenCustomEndCap(pen *GpPen, customCap *GpCustomLineCap) GpStatus {
	ret, _, _ := gdipSetPenCustomEndCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return GpStatus(ret)
}
func GdipGetPenCustomEndCap(pen *GpPen, customCap **GpCustomLineCap) GpStatus {
	ret, _, _ := gdipGetPenCustomEndCap.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return GpStatus(ret)
}
func GdipSetPenMiterLimit(pen *GpPen, miterLimit float32) GpStatus {
	ret, _, _ := gdipSetPenMiterLimit.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(miterLimit)))
	return GpStatus(ret)
}
func GdipGetPenMiterLimit(pen *GpPen, miterLimit *float32) GpStatus {
	var iMiterLimit uint32
	ret, _, _ := gdipGetPenMiterLimit.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(&iMiterLimit)))
	*miterLimit = math.Float32frombits(iMiterLimit)
	return GpStatus(ret)
}
func GdipSetPenMode(pen *GpPen, penMode GpPenAlignment) GpStatus {
	ret, _, _ := gdipSetPenMode.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(penMode))
	return GpStatus(ret)
}
func GdipGetPenMode(pen *GpPen, penMode *GpPenAlignment) GpStatus {
	ret, _, _ := gdipGetPenMode.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(penMode)))
	return GpStatus(ret)
}
func GdipSetPenTransform(pen *GpPen, matrix *GpMatrix) GpStatus {
	ret, _, _ := gdipSetPenTransform.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)))
	return GpStatus(ret)
}
func GdipGetPenTransform(pen *GpPen, matrix *GpMatrix) GpStatus {
	ret, _, _ := gdipGetPenTransform.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)))
	return GpStatus(ret)
}
func GdipResetPenTransform(pen *GpPen) GpStatus {
	ret, _, _ := gdipResetPenTransform.Call(uintptr(unsafe.Pointer(pen)))
	return GpStatus(ret)
}
func GdipMultiplyPenTransform(pen *GpPen, matrix *GpMatrix, order GpMatrixOrder) GpStatus {
	ret, _, _ := gdipMultiplyPenTransform.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
	return GpStatus(ret)
}
func GdipTranslatePenTransform(pen *GpPen, dx, dy float32, order GpMatrixOrder) GpStatus {
	ret, _, _ := gdipTranslatePenTransform.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)),
		uintptr(order))
	return GpStatus(ret)
}
func GdipScalePenTransform(pen *GpPen, sx, sy float32, order GpMatrixOrder) GpStatus {
	ret, _, _ := gdipScalePenTransform.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(sx)),
		uintptr(math.Float32bits(sy)),
		uintptr(order))
	return GpStatus(ret)
}
func GdipRotatePenTransform(pen *GpPen, angle float32, order GpMatrixOrder) GpStatus {
	ret, _, _ := gdipRotatePenTransform.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(angle)),
		uintptr(order))
	return GpStatus(ret)
}
func GdipSetPenColor(pen *GpPen, argb ARGB) GpStatus {
	ret, _, _ := gdipSetPenColor.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(argb))
	return GpStatus(ret)
}
func GdipGetPenColor(pen *GpPen, argb *ARGB) GpStatus {
	ret, _, _ := gdipGetPenColor.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(argb)))
	return GpStatus(ret)
}
func GdipSetPenBrushFill(pen *GpPen, brush *GpBrush) GpStatus {
	ret, _, _ := gdipSetPenBrushFill.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(brush)))
	return GpStatus(ret)
}
func GdipGetPenBrushFill(pen *GpPen, brush **GpBrush) GpStatus {
	ret, _, _ := gdipGetPenBrushFill.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(brush)))
	return GpStatus(ret)
}
func GdipGetPenFillType(pen *GpPen, penType *GpPenType) GpStatus {
	ret, _, _ := gdipGetPenFillType.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(penType)))
	return GpStatus(ret)
}
func GdipGetPenDashStyle(pen *GpPen, dashStyle *GpDashStyle) GpStatus {
	ret, _, _ := gdipGetPenDashStyle.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dashStyle)))
	return GpStatus(ret)
}
func GdipSetPenDashStyle(pen *GpPen, dashStyle GpDashStyle) GpStatus {
	ret, _, _ := gdipSetPenDashStyle.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(dashStyle))
	return GpStatus(ret)
}
func GdipGetPenDashOffset(pen *GpPen, offset *float32) GpStatus {
	var iOffset uint32
	ret, _, _ := gdipGetPenDashOffset.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(&iOffset)))
	*offset = math.Float32frombits(iOffset)
	return GpStatus(ret)
}
func GdipSetPenDashOffset(pen *GpPen, offset float32) GpStatus {
	ret, _, _ := gdipSetPenDashOffset.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(offset)))
	return GpStatus(ret)
}
func GdipGetPenDashCount(pen *GpPen, count *int32) GpStatus {
	ret, _, _ := gdipGetPenDashCount.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(count)))
	return GpStatus(ret)
}
func GdipSetPenDashArray(pen *GpPen, dash *float32, count int32) GpStatus {
	ret, _, _ := gdipSetPenDashArray.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return GpStatus(ret)
}
func GdipGetPenDashArray(pen *GpPen, dash *float32, count int32) GpStatus {
	ret, _, _ := gdipGetPenDashArray.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return GpStatus(ret)
}

func GdipGetPenCompoundCount(pen *GpPen, count *int32) GpStatus {
	ret, _, _ := gdipGetPenCompoundCount.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(count)))
	return GpStatus(ret)
}

func GdipSetPenCompoundArray(pen *GpPen, dash *float32, count int32) GpStatus {
	ret, _, _ := gdipSetPenCompoundArray.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return GpStatus(ret)
}

func GdipGetPenCompoundArray(pen *GpPen, dash *float32, count int32) GpStatus {
	ret, _, _ := gdipGetPenCompoundArray.Call(
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return GpStatus(ret)
}

// Brush

func GdipCloneBrush(brush *GpBrush, clone **GpBrush) GpStatus {
	ret, _, _ := gdipCloneBrush.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(clone)))
	return GpStatus(ret)
}

func GdipDeleteBrush(brush *GpBrush) GpStatus {
	ret, _, _ := gdipDeleteBrush.Call(uintptr(unsafe.Pointer(brush)))
	return GpStatus(ret)
}

func GdipGetBrushType(brush *GpBrush, brushType *GpBrushType) GpStatus {
	ret, _, _ := gdipGetBrushType.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(brushType)))
	return GpStatus(ret)
}

// Solid Brush

func GdipCreateSolidFill(color ARGB, brush **GpSolidFill) GpStatus {
	ret, _, _ := gdipCreateSolidFill.Call(
		uintptr(color),
		uintptr(unsafe.Pointer(brush)))
	return GpStatus(ret)
}

func GdipSetSolidFillColor(brush *GpBrush, color ARGB) GpStatus {
	ret, _, _ := gdipSetSolidFillColor.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(color))
	return GpStatus(ret)
}

func GdipGetSolidFillColor(brush *GpBrush, color *ARGB) GpStatus {
	ret, _, _ := gdipGetSolidFillColor.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(color)))
	return GpStatus(ret)
}

// Font
func GdipCreateFontFromDC(hdc HDC, font **GpFont) GpStatus {
	ret, _, _ := gdipCreateFontFromDC.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(font)))
	return GpStatus(ret)
}

func GdipCreateFont(fontFamily *GpFontFamily, emSize float32, style int32, unit GpUnit, font **GpFont) GpStatus {
	ret, _, _ := gdipCreateFont.Call(
		uintptr(unsafe.Pointer(fontFamily)),
		uintptr(math.Float32bits(emSize)),
		uintptr(style),
		uintptr(unit),
		uintptr(unsafe.Pointer(font)))
	return GpStatus(ret)
}

func GdipDeleteFont(font *GpFont) GpStatus {
	ret, _, _ := gdipDeleteFont.Call(uintptr(unsafe.Pointer(font)))
	return GpStatus(ret)
}

func GdipNewInstalledFontCollection(fontCollection **GpFontCollection) GpStatus {
	ret, _, _ := gdipNewInstalledFontCollection.Call(uintptr(unsafe.Pointer(fontCollection)))
	return GpStatus(ret)
}

func GdipCreateFontFamilyFromName(name *uint16, fontCollection *GpFontCollection, fontFamily **GpFontFamily) GpStatus {
	ret, _, _ := gdipCreateFontFamilyFromName.Call(
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(fontFamily)))
	return GpStatus(ret)
}

func GdipDeleteFontFamily(fontFamily *GpFontFamily) GpStatus {
	ret, _, _ := gdipDeleteFontFamily.Call(uintptr(unsafe.Pointer(fontFamily)))
	return GpStatus(ret)
}

// StringFormat

func GdipCreateStringFormat(formatAttributes int32, language uint16, format **GpStringFormat) GpStatus {
	ret, _, _ := gdipCreateStringFormat.Call(
		uintptr(formatAttributes),
		uintptr(language),
		uintptr(unsafe.Pointer(format)))
	return GpStatus(ret)
}

func GdipStringFormatGetGenericTypographic(format **GpStringFormat) GpStatus {
	ret, _, _ := gdipStringFormatGetGenericTypographic.Call(uintptr(unsafe.Pointer(format)))
	return GpStatus(ret)
}

func GdipDeleteStringFormat(format *GpStringFormat) GpStatus {
	ret, _, _ := gdipDeleteStringFormat.Call(uintptr(unsafe.Pointer(format)))
	return GpStatus(ret)
}

// Path

func GdipCreatePath(brushMode int32, path **GpPath) GpStatus {
	ret, _, _ := gdipCreatePath.Call(uintptr(brushMode), uintptr(unsafe.Pointer(path)))
	return GpStatus(ret)
}

func GdipDeletePath(path *GpPath) GpStatus {
	ret, _, _ := gdipDeletePath.Call(uintptr(unsafe.Pointer(path)))
	return GpStatus(ret)
}

func GdipAddPathArc(path *GpPath, x, y, width, height, startAngle, sweepAngle float32) GpStatus {
	ret, _, _ := gdipAddPathArc.Call(
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return GpStatus(ret)
}

func GdipAddPathArcI(path *GpPath, x, y, width, height int32, startAngle, sweepAngle float32) GpStatus {
	ret, _, _ := gdipAddPathArcI.Call(
		uintptr(unsafe.Pointer(path)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return GpStatus(ret)
}

func GdipAddPathLine(path *GpPath, x1, y1, x2, y2 float32) GpStatus {
	ret, _, _ := gdipAddPathLine.Call(
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x1)),
		uintptr(math.Float32bits(y1)),
		uintptr(math.Float32bits(x2)),
		uintptr(math.Float32bits(y2)))
	return GpStatus(ret)
}

func GdipAddPathLineI(path *GpPath, x1, y1, x2, y2 int32) GpStatus {
	ret, _, _ := gdipAddPathLineI.Call(
		uintptr(unsafe.Pointer(path)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
	return GpStatus(ret)
}

func GdipClosePathFigure(path *GpPath) GpStatus {
	ret, _, _ := gdipClosePathFigure.Call(uintptr(unsafe.Pointer(path)))
	return GpStatus(ret)
}

func GdipClosePathFigures(path *GpPath) GpStatus {
	ret, _, _ := gdipClosePathFigures.Call(uintptr(unsafe.Pointer(path)))
	return GpStatus(ret)
}

// Image

func GdipGetImageGraphicsContext(image *GpImage, graphics **GpGraphics) GpStatus {
	ret, _, _ := gdipGetImageGraphicsContext.Call(
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(graphics)))
	return GpStatus(ret)
}

func GdipLoadImageFromFile(filename *uint16, image **GpImage) GpStatus {
	ret, _, _ := gdipLoadImageFromFile.Call(
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(image)))
	return GpStatus(ret)
}

func GdipSaveImageToFile(image *GpBitmap, filename *uint16, clsidEncoder *ole.GUID, encoderParams *EncoderParameters) GpStatus {
	ret, _, _ := gdipSaveImageToFile.Call(uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(encoderParams)))
	return GpStatus(ret)
}

func GdipGetImageWidth(image *GpImage, width *uint32) GpStatus {
	ret, _, _ := gdipGetImageWidth.Call(uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(width)))
	return GpStatus(ret)
}

func GdipGetImageHeight(image *GpImage, height *uint32) GpStatus {
	ret, _, _ := gdipGetImageHeight.Call(uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(height)))
	return GpStatus(ret)
}

func GdipDisposeImage(image *GpImage) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDisposeImage.Addr(), 1,
		uintptr(unsafe.Pointer(image)),
		0,
		0)

	return GpStatus(ret)
}

// Bitmap

func GdipCreateBitmapFromFile(filename *uint16, bitmap **GpBitmap) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateBitmapFromFile.Addr(), 2,
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(bitmap)),
		0)

	return GpStatus(ret)
}

func GdipCreateBitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE, bitmap **GpBitmap) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateBitmapFromHBITMAP.Addr(), 3,
		uintptr(hbm),
		uintptr(hpal),
		uintptr(unsafe.Pointer(bitmap)))

	return GpStatus(ret)
}

func GdipCreateHBITMAPFromBitmap(bitmap *GpBitmap, hbmReturn *HBITMAP, background ARGB) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateHBITMAPFromBitmap.Addr(), 3,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(hbmReturn)),
		uintptr(background))

	return GpStatus(ret)
}

func GdipCreateBitmapFromScan0(width, height, stride int32, format PixelFormat, scan0 *byte, bitmap **GpBitmap) GpStatus {
	ret, _, _ := gdipCreateBitmapFromScan0.Call(
		uintptr(width),
		uintptr(height),
		uintptr(stride),
		uintptr(format),
		uintptr(unsafe.Pointer(scan0)),
		uintptr(unsafe.Pointer(bitmap)))
	return GpStatus(ret)
}

/*
func SavePNG(fileName string, newBMP win.HBITMAP) error {
	// HBITMAP
	var bmp *win.GpBitmap
	if win.GdipCreateBitmapFromHBITMAP(newBMP, 0, &bmp) != 0 {
		return fmt.Errorf("failed to create HBITMAP")
	}
	defer win.GdipDisposeImage((*GpImage)(bmp))
	clsid, err := ole.CLSIDFromString("{557CF406-1A04-11D3-9A73-0000F81EF32E}")
	if err != nil {
		return err
	}
	fname, err := syscall.UTF16PtrFromString(fileName)
	if err != nil {
		return err
	}
	if GdipSaveImageToFile(bmp, fname, clsid, nil) != 0 {
		return fmt.Errorf("failed to call PNG encoder")
	}
	return nil
}
*/

type Bitmap struct {
	Image
}

func NewBitmap(width, height int32, format PixelFormat) *Bitmap {
	bitmap := &Bitmap{}
	var nativeBitmap *GpBitmap
	status := GdipCreateBitmapFromScan0(width, height, 0, format, nil, &nativeBitmap)
	if status != Ok {
		// log.Panicln(status.String())
	}
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func NewBitmapEx(width, height, stride int32, format PixelFormat, scan0 *byte) *Bitmap {
	bitmap := &Bitmap{}
	var nativeBitmap *GpBitmap
	GdipCreateBitmapFromScan0(width, height, stride, format, scan0, &nativeBitmap)
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func NewBitmapFromHBITMAP(hbitmap HBITMAP) *Bitmap {
	bitmap := &Bitmap{}
	var nativeBitmap *GpBitmap
	GdipCreateBitmapFromHBITMAP(hbitmap, 0, &nativeBitmap)
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func NewBitmapFromFile(fileName string) *Bitmap {
	bitmap := &Bitmap{}
	fileNameUTF16, _ := syscall.UTF16PtrFromString(fileName)
	var nativeBitmap *GpBitmap
	GdipCreateBitmapFromFile(fileNameUTF16, &nativeBitmap)
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func (bitmap *Bitmap) Dispose() {
	GdipDisposeImage(bitmap.nativeImage)
}

type Brush struct {
	nativeBrush *GpBrush
}

type SolidBrush struct {
	Brush
}

func (b *Brush) Dispose() {
	GdipDeleteBrush(b.nativeBrush)
}

func (b *Brush) GetBrushType() (brushType BrushType) {
	GdipGetBrushType(b.nativeBrush, (*GpBrushType)(&brushType))
	return
}

func (b *Brush) Clone() *Brush {
	clone := &Brush{}
	GdipCloneBrush(b.nativeBrush, &clone.nativeBrush)
	return clone
}

func NewSolidBrush(color *Color) *SolidBrush {
	b := &SolidBrush{}
	var solidFill *GpSolidFill
	GdipCreateSolidFill(color.GetValue(), &solidFill)
	b.nativeBrush = &solidFill.GpBrush
	return b
}

func (b *SolidBrush) AsBrush() *Brush {
	return &b.Brush
}

func (b *SolidBrush) SetColor(color *Color) {
	GdipSetSolidFillColor(b.nativeBrush, color.GetValue())
}

func (b *SolidBrush) GetColor() (color Color) {
	GdipGetSolidFillColor(b.nativeBrush, &color.Argb)
	return
}

type Color struct {
	Argb ARGB
}

func MakeARGB(a, r, g, b byte) ARGB {
	return ((ARGB(b) << BlueShift) | (ARGB(g) << GreenShift) | (ARGB(r) << RedShift) | (ARGB(a) << AlphaShift))
}

func NewColor(r, g, b, a byte) *Color {
	c := &Color{}
	c.Argb = MakeARGB(a, r, g, b)
	return c
}

func (c *Color) GetAlpha() byte {
	return byte(c.Argb >> AlphaShift)
}

func (c *Color) GetA() byte {
	return c.GetAlpha()
}

func (c *Color) GetRed() byte {
	return byte(c.Argb >> RedShift)
}

func (c *Color) GetR() byte {
	return c.GetRed()
}

func (c *Color) GetGreen() byte {
	return byte(c.Argb >> GreenShift)
}

func (c *Color) GetG() byte {
	return c.GetGreen()
}

func (c *Color) GetBlue() byte {
	return byte(c.Argb >> BlueShift)
}

func (c *Color) GetB() byte {
	return c.GetBlue()
}

func (c *Color) GetValue() ARGB {
	return c.Argb
}

type Image struct {
	nativeImage *GpImage
}

func NewImageFromFile(fileName string) *Image {
	image := &Image{}
	fileNameUTF16, _ := syscall.UTF16PtrFromString(fileName)
	GdipLoadImageFromFile(fileNameUTF16, &image.nativeImage)
	return image
}

func (image *Image) GetWidth() (width uint32) {
	GdipGetImageWidth(image.nativeImage, &width)
	return

}

func (image *Image) GetHeight() (height uint32) {
	GdipGetImageHeight(image.nativeImage, &height)
	return

}
func (image *Image) Get() *GpImage {
	return image.nativeImage
}
func (image *Image) Dispose() {
	GdipDisposeImage(image.nativeImage)
}

type GraphicsPath struct {
	nativePath *GpPath
}

func NewPath(fillMode int32) *GraphicsPath {
	p := &GraphicsPath{}
	GdipCreatePath(fillMode, &p.nativePath)
	return p
}

func (p *GraphicsPath) AddArcRect(rect *Rect, startAngle, sweepAngle float32) {
	GdipAddPathArcI(p.nativePath, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (p *GraphicsPath) AddArcRectF(rect *RectF, startAngle, sweepAngle float32) {
	GdipAddPathArc(p.nativePath, rect.X, rect.Y, rect.Width, rect.Height, startAngle, sweepAngle)
}

func (p *GraphicsPath) AddArc(x, y, width, height, startAngle, sweepAngle float32) {
	GdipAddPathArc(p.nativePath, x, y, width, height, startAngle, sweepAngle)
}

func (p *GraphicsPath) AddArcI(x, y, width, height int32, startAngle, sweepAngle float32) {
	GdipAddPathArcI(p.nativePath, x, y, width, height, startAngle, sweepAngle)
}

func (p *GraphicsPath) AddLine(x1, y1, x2, y2 float32) {
	GdipAddPathLine(p.nativePath, x1, y1, x2, y2)
}

func (p *GraphicsPath) AddLineI(x1, y1, x2, y2 int32) {
	GdipAddPathLineI(p.nativePath, x1, y1, x2, y2)
}

func (p *GraphicsPath) CloseAllFigures() {
	GdipClosePathFigures(p.nativePath)
}

func (p *GraphicsPath) CloseFigure() {
	GdipClosePathFigure(p.nativePath)
}

func (p *GraphicsPath) Dispose() {
	GdipDeletePath(p.nativePath)
}

type Pen struct {
	nativePen *GpPen
}

func NewPen(color *Color, width float32) *Pen {
	p := &Pen{}
	GdipCreatePen1(color.GetValue(), width, UnitWorld, &p.nativePen)
	return p
}

func NewPenFromBrush(brush *Brush, width float32) *Pen {
	p := &Pen{}
	GdipCreatePen2(brush.nativeBrush, width, UnitWorld, &p.nativePen)
	return p
}

func (p *Pen) Dispose() {
	GdipDeletePen(p.nativePen)
}

func (p *Pen) Clone() *Pen {
	clone := &Pen{}
	GdipClonePen(p.nativePen, &clone.nativePen)
	return clone
}

func (p *Pen) SetWidth(width float32) {
	GdipSetPenWidth(p.nativePen, width)
}

func (p *Pen) GetWidth() (width float32) {
	GdipGetPenWidth(p.nativePen, &width)
	return
}

func (p *Pen) SetLineCap(startCap, endCap LineCap, dashCap DashCap) {
	GdipSetPenLineCap197819(p.nativePen, GpLineCap(startCap), GpLineCap(endCap), GpDashCap(dashCap))
}

func (p *Pen) SetStartCap(startCap LineCap) {
	GdipSetPenStartCap(p.nativePen, GpLineCap(startCap))
}

func (p *Pen) SetEndCap(endCap LineCap) {
	GdipSetPenEndCap(p.nativePen, GpLineCap(endCap))
}

func (p *Pen) SetDashCap(dashCap DashCap) {
	GdipSetPenDashCap197819(p.nativePen, GpDashCap(dashCap))
}

func (p *Pen) GetStartCap() (startCap LineCap) {
	GdipGetPenStartCap(p.nativePen, (*GpLineCap)(&startCap))
	return
}

func (p *Pen) GetEndCap() (endCap LineCap) {
	GdipGetPenEndCap(p.nativePen, (*GpLineCap)(&endCap))
	return
}

func (p *Pen) GetDashCap() (dashCap DashCap) {
	GdipGetPenDashCap197819(p.nativePen, (*GpDashCap)(&dashCap))
	return
}

func (p *Pen) SetLineJoin(lineJoin LineJoin) {
	GdipSetPenLineJoin(p.nativePen, GpLineJoin(lineJoin))
}

func (p *Pen) GetLineJoin() (lineJoin LineJoin) {
	GdipGetPenLineJoin(p.nativePen, (*GpLineJoin)(&lineJoin))
	return
}

func (p *Pen) SetCustomStartCap(customCap *GpCustomLineCap) {
	GdipSetPenCustomStartCap(p.nativePen, customCap)
}

func (p *Pen) GetCustomStartCap() (customCap *GpCustomLineCap) {
	GdipGetPenCustomStartCap(p.nativePen, &customCap)
	return
}

func (p *Pen) SetCustomEndCap(customCap *GpCustomLineCap) {
	GdipSetPenCustomEndCap(p.nativePen, customCap)
}

func (p *Pen) GetCustomEndCap() (customCap *GpCustomLineCap) {
	GdipGetPenCustomEndCap(p.nativePen, &customCap)
	return
}

func (p *Pen) SetMiterLimit(miterLimit float32) {
	GdipSetPenMiterLimit(p.nativePen, miterLimit)
}

func (p *Pen) GetMiterLimit() (miterLimit float32) {
	GdipGetPenMiterLimit(p.nativePen, &miterLimit)
	return
}

func (p *Pen) SetMode(penMode PenAlignment) {
	GdipSetPenMode(p.nativePen, GpPenAlignment(penMode))
}

func (p *Pen) GetMode() (penMode PenAlignment) {
	GdipGetPenMode(p.nativePen, (*GpPenAlignment)(&penMode))
	return
}

func (p *Pen) SetTransform(matrix *GpMatrix) {
	GdipSetPenTransform(p.nativePen, matrix)
}

func (p *Pen) GetTransform(matrix *GpMatrix) {
	GdipGetPenTransform(p.nativePen, matrix)
}

func (p *Pen) ResetTransform() {
	GdipResetPenTransform(p.nativePen)
}

func (p *Pen) MultiplyTransform(matrix *GpMatrix, order MatrixOrder) {
	GdipMultiplyPenTransform(p.nativePen, matrix, GpMatrixOrder(order))
}

func (p *Pen) TranslateTransform(dx, dy float32, order MatrixOrder) {
	GdipTranslatePenTransform(p.nativePen, dx, dy, GpMatrixOrder(order))
}

func (p *Pen) ScaleTransform(sx, sy float32, order MatrixOrder) {
	GdipScalePenTransform(p.nativePen, sx, sy, GpMatrixOrder(order))
}

func (p *Pen) RotateTransform(angle float32, order MatrixOrder) {
	GdipRotatePenTransform(p.nativePen, angle, GpMatrixOrder(order))
}

func (p *Pen) SetColor(color *Color) {
	GdipSetPenColor(p.nativePen, color.GetValue())
}

func (p *Pen) GetColor() (color Color) {
	GdipGetPenColor(p.nativePen, &color.Argb)
	return
}

func (p *Pen) SetBrush(brush *Brush) {
	GdipSetPenBrushFill(p.nativePen, brush.nativeBrush)
}

func (p *Pen) GetBrush() *Brush {
	brush := &Brush{}
	GdipGetPenBrushFill(p.nativePen, &brush.nativeBrush)
	return brush
}

func (p *Pen) GetPenType() (penType PenType) {
	GdipGetPenFillType(p.nativePen, (*GpPenType)(&penType))
	return
}

func (p *Pen) GetDashStyle() (dashStyle DashStyle) {
	GdipGetPenDashStyle(p.nativePen, (*GpDashStyle)(&dashStyle))
	return
}

func (p *Pen) SetDashStyle(dashStyle DashStyle) {
	GdipSetPenDashStyle(p.nativePen, GpDashStyle(dashStyle))
}

func (p *Pen) GetDashOffset() (offset float32) {
	GdipGetPenDashOffset(p.nativePen, &offset)
	return
}

func (p *Pen) SetDashOffset(offset float32) {
	GdipSetPenDashOffset(p.nativePen, offset)
}

func (p *Pen) GetDashCount() (count int32) {
	GdipGetPenDashCount(p.nativePen, &count)
	return
}

func (p *Pen) SetDashArray(dash []float32) {
	GdipSetPenDashArray(p.nativePen, &dash[0], int32(len(dash)))
}

func (p *Pen) GetDashArray(dash *float32, count int32) {
	GdipGetPenDashArray(p.nativePen, dash, count)
}

func (p *Pen) GetCompoundCount() (count int32) {
	GdipGetPenCompoundCount(p.nativePen, &count)
	return
}

func (p *Pen) SetCompoundArray(dash []float32) {
	GdipSetPenCompoundArray(p.nativePen, &dash[0], int32(len(dash)))
}

func (p *Pen) GetCompoundArray(dash *float32, count int32) {
	GdipGetPenCompoundArray(p.nativePen, dash, count)
}

type StringFormat struct {
	nativeFormat *GpStringFormat
}

func NewStringFormat() *StringFormat {
	format := &StringFormat{}
	GdipCreateStringFormat(0, LANG_NEUTRAL, &format.nativeFormat)
	return format
}

func NewGenericTypographicStringFormat() *StringFormat {
	format := &StringFormat{}
	GdipStringFormatGetGenericTypographic(&format.nativeFormat)
	return format
}

func (format *StringFormat) Dispose() {
	GdipDeleteStringFormat(format.nativeFormat)
}
