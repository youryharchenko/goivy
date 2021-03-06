// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 15 May 2021 15:29:56 EEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package nk

/*
#include "nk.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// IncludeFixedTypes as defined in goivy/<predefine>:24
	IncludeFixedTypes = 1
	// IncludeStandardIo as defined in goivy/<predefine>:25
	IncludeStandardIo = 1
	// IncludeDefaultAllocator as defined in goivy/<predefine>:26
	IncludeDefaultAllocator = 1
	// IncludeFontBaking as defined in goivy/<predefine>:27
	IncludeFontBaking = 1
	// IncludeDefaultFont as defined in goivy/<predefine>:28
	IncludeDefaultFont = 1
	// IncludeVertexBufferOutput as defined in goivy/<predefine>:29
	IncludeVertexBufferOutput = 1
	// Undefined as defined in nk/nuklear.h:233
	Undefined = (-1.0)
	// UtfInvalid as defined in nk/nuklear.h:234
	UtfInvalid = 0xFFFD
	// UtfSize as defined in nk/nuklear.h:235
	UtfSize = 4
	// InputMax as defined in nk/nuklear.h:237
	InputMax = 16
	// MaxNumberBuffer as defined in nk/nuklear.h:240
	MaxNumberBuffer = 64
	// ScrollbarHidingTimeout as defined in nk/nuklear.h:243
	ScrollbarHidingTimeout = 4.0
	// Lib as defined in nk/nuklear.h:267
	Lib = 0
	// TexteditUndostatecount as defined in nk/nuklear.h:4239
	TexteditUndostatecount = 99
	// TexteditUndocharcount as defined in nk/nuklear.h:4243
	TexteditUndocharcount = 999
	// MaxLayoutRowTemplateColumns as defined in nk/nuklear.h:5246
	MaxLayoutRowTemplateColumns = 16
	// ChartMaxSlot as defined in nk/nuklear.h:5249
	ChartMaxSlot = 4
	// WindowMaxName as defined in nk/nuklear.h:5348
	WindowMaxName = 64
	// ButtonBehaviorStackSize as defined in nk/nuklear.h:5461
	ButtonBehaviorStackSize = 8
	// FontStackSize as defined in nk/nuklear.h:5465
	FontStackSize = 8
	// StyleItemStackSize as defined in nk/nuklear.h:5469
	StyleItemStackSize = 16
	// FloatStackSize as defined in nk/nuklear.h:5473
	FloatStackSize = 32
	// VectorStackSize as defined in nk/nuklear.h:5477
	VectorStackSize = 16
	// FlagsStackSize as defined in nk/nuklear.h:5481
	FlagsStackSize = 32
	// ColorStackSize as defined in nk/nuklear.h:5485
	ColorStackSize = 32
	// Float as defined in nk/nuklear.h:5499
	Float = 0
	// Pi as defined in nk/nuklear.h:5613
	Pi = 3.141592654
	// MaxFloatPrecision as defined in nk/nuklear.h:5615
	MaxFloatPrecision = 2
)

// Heading as declared in nk/nuklear.h:469
type Heading int32

// Heading enumeration from nk/nuklear.h:469
const (
	Up    = iota
	Right = 1
	Down  = 2
	Left  = 3
)

// ButtonBehavior as declared in nk/nuklear.h:470
type ButtonBehavior int32

// ButtonBehavior enumeration from nk/nuklear.h:470
const (
	ButtonDefault  = iota
	ButtonRepeater = 1
)

// Modify as declared in nk/nuklear.h:471
type Modify int32

// Modify enumeration from nk/nuklear.h:471
const (
	Fixed      = False
	Modifiable = True
)

// Orientation as declared in nk/nuklear.h:472
type Orientation int32

// Orientation enumeration from nk/nuklear.h:472
const (
	Vertical   = iota
	Horizontal = 1
)

// CollapseStates as declared in nk/nuklear.h:473
type CollapseStates int32

// CollapseStates enumeration from nk/nuklear.h:473
const (
	Minimized = False
	Maximized = True
)

// ShowStates as declared in nk/nuklear.h:474
type ShowStates int32

// ShowStates enumeration from nk/nuklear.h:474
const (
	Hidden = False
	Shown  = True
)

// ChartType as declared in nk/nuklear.h:475
type ChartType int32

// ChartType enumeration from nk/nuklear.h:475
const (
	ChartLines  = iota
	ChartColumn = 1
	ChartMax    = 2
)

// ChartEvent as declared in nk/nuklear.h:476
type ChartEvent int32

// ChartEvent enumeration from nk/nuklear.h:476
const (
	ChartHovering = 0x01
	ChartClicked  = 0x02
)

// ColorFormat as declared in nk/nuklear.h:477
type ColorFormat int32

// ColorFormat enumeration from nk/nuklear.h:477
const (
	ColorFormatRGB  = iota
	ColorFormatRGBA = 1
)

// PopupType as declared in nk/nuklear.h:478
type PopupType int32

// PopupType enumeration from nk/nuklear.h:478
const (
	PopupStatic  = iota
	PopupDynamic = 1
)

// LayoutFormat as declared in nk/nuklear.h:479
type LayoutFormat int32

// LayoutFormat enumeration from nk/nuklear.h:479
const (
	Dynamic = iota
	Static  = 1
)

// TreeType as declared in nk/nuklear.h:480
type TreeType int32

// TreeType enumeration from nk/nuklear.h:480
const (
	TreeNode = iota
	TreeTab  = 1
)

// SymbolType as declared in nk/nuklear.h:493
type SymbolType int32

// SymbolType enumeration from nk/nuklear.h:493
const (
	SymbolNone          = iota
	SymbolX             = 1
	SymbolUnderscore    = 2
	SymbolCircleSolid   = 3
	SymbolCircleOutline = 4
	SymbolRectSolid     = 5
	SymbolRectOutline   = 6
	SymbolTriangleUp    = 7
	SymbolTriangleDown  = 8
	SymbolTriangleLeft  = 9
	SymbolTriangleRight = 10
	SymbolPlus          = 11
	SymbolMinus         = 12
	SymbolMax           = 13
)

// Keys as declared in nk/nuklear.h:735
type Keys int32

// Keys enumeration from nk/nuklear.h:735
const (
	KeyNone            = iota
	KeyShift           = 1
	KeyCtrl            = 2
	KeyDel             = 3
	KeyEnter           = 4
	KeyTab             = 5
	KeyBackspace       = 6
	KeyCopy            = 7
	KeyCut             = 8
	KeyPaste           = 9
	KeyUp              = 10
	KeyDown            = 11
	KeyLeft            = 12
	KeyRight           = 13
	KeyTextInsertMode  = 14
	KeyTextReplaceMode = 15
	KeyTextResetMode   = 16
	KeyTextLineStart   = 17
	KeyTextLineEnd     = 18
	KeyTextStart       = 19
	KeyTextEnd         = 20
	KeyTextUndo        = 21
	KeyTextRedo        = 22
	KeyTextSelectAll   = 23
	KeyTextWordLeft    = 24
	KeyTextWordRight   = 25
	KeyScrollStart     = 26
	KeyScrollEnd       = 27
	KeyScrollDown      = 28
	KeyScrollUp        = 29
	KeyMax             = 30
)

// Buttons as declared in nk/nuklear.h:770
type Buttons int32

// Buttons enumeration from nk/nuklear.h:770
const (
	ButtonLeft   = iota
	ButtonMiddle = 1
	ButtonRight  = 2
	ButtonDouble = 3
	ButtonMax    = 4
)

// AntiAliasing as declared in nk/nuklear.h:1142
type AntiAliasing int32

// AntiAliasing enumeration from nk/nuklear.h:1142
const (
	AntiAliasingOff = iota
	AntiAliasingOn  = 1
)

// ConvertResult as declared in nk/nuklear.h:1143
type ConvertResult int32

// ConvertResult enumeration from nk/nuklear.h:1143
const (
	ConvertSuccess           = iota
	ConvertInvalidParam      = 1
	ConvertCommandBufferFull = (1 << (1))
	ConvertVertexBufferFull  = (1 << (2))
	ConvertElementBufferFull = (1 << (3))
)

// PanelFlags as declared in nk/nuklear.h:1450
type PanelFlags int32

// PanelFlags enumeration from nk/nuklear.h:1450
const (
	WindowBorder         = (1 << (0))
	WindowMovable        = (1 << (1))
	WindowScalable       = (1 << (2))
	WindowClosable       = (1 << (3))
	WindowMinimizable    = (1 << (4))
	WindowNoScrollbar    = (1 << (5))
	WindowTitle          = (1 << (6))
	WindowScrollAutoHide = (1 << (7))
	WindowBackground     = (1 << (8))
	WindowScaleLeft      = (1 << (9))
	WindowNoInput        = (1 << (10))
)

// WidgetLayoutStates as declared in nk/nuklear.h:3041
type WidgetLayoutStates int32

// WidgetLayoutStates enumeration from nk/nuklear.h:3041
const (
	WidgetInvalid = iota
	WidgetValid   = 1
	WidgetRom     = 2
)

// WidgetStates as declared in nk/nuklear.h:3046
type WidgetStates int32

// WidgetStates enumeration from nk/nuklear.h:3046
const (
	WidgetStateModified = (1 << (1))
	WidgetStateInactive = (1 << (2))
	WidgetStateEntered  = (1 << (3))
	WidgetStateHover    = (1 << (4))
	WidgetStateActived  = (1 << (5))
	WidgetStateLeft     = (1 << (6))
	WidgetStateHovered  = WidgetStateHover | WidgetStateModified
	WidgetStateActive   = WidgetStateActived | WidgetStateModified
)

// TextAlign as declared in nk/nuklear.h:3072
type TextAlign int32

// TextAlign enumeration from nk/nuklear.h:3072
const (
	TextAlignLeft     = 0x01
	TextAlignCentered = 0x02
	TextAlignRight    = 0x04
	TextAlignTop      = 0x08
	TextAlignMiddle   = 0x10
	TextAlignBottom   = 0x20
)

// TextAlignment as declared in nk/nuklear.h:3080
type TextAlignment int32

// TextAlignment enumeration from nk/nuklear.h:3080
const (
	TextLeft     = TextAlignMiddle | TextAlignLeft
	TextCentered = TextAlignMiddle | TextAlignCentered
	TextRight    = TextAlignMiddle | TextAlignRight
)

// EditFlags as declared in nk/nuklear.h:3415
type EditFlags int32

// EditFlags enumeration from nk/nuklear.h:3415
const (
	EditDefault            = iota
	EditReadOnly           = (1 << (0))
	EditAutoSelect         = (1 << (1))
	EditSigEnter           = (1 << (2))
	EditAllowTab           = (1 << (3))
	EditNoCursor           = (1 << (4))
	EditSelectable         = (1 << (5))
	EditClipboard          = (1 << (6))
	EditCtrlEnterNewline   = (1 << (7))
	EditNoHorizontalScroll = (1 << (8))
	EditAlwaysInsertMode   = (1 << (9))
	EditMultiline          = (1 << (10))
	EditGotoEndOnActivate  = (1 << (11))
)

// EditTypes as declared in nk/nuklear.h:3430
type EditTypes int32

// EditTypes enumeration from nk/nuklear.h:3430
const (
	EditSimple = EditAlwaysInsertMode
	EditField  = EditSimple | EditSelectable | EditClipboard
	EditBox    = EditAlwaysInsertMode | EditSelectable | EditMultiline | EditAllowTab | EditClipboard
	EditEditor = EditSelectable | EditMultiline | EditAllowTab | EditClipboard
)

// EditEvents as declared in nk/nuklear.h:3436
type EditEvents int32

// EditEvents enumeration from nk/nuklear.h:3436
const (
	EditActive      = (1 << (0))
	EditInactive    = (1 << (1))
	EditActivated   = (1 << (2))
	EditDeactivated = (1 << (3))
	EditCommited    = (1 << (4))
)

// StyleColors as declared in nk/nuklear.h:3561
type StyleColors int32

// StyleColors enumeration from nk/nuklear.h:3561
const (
	ColorText                  = iota
	ColorWindow                = 1
	ColorHeader                = 2
	ColorBorder                = 3
	ColorButton                = 4
	ColorButtonHover           = 5
	ColorButtonActive          = 6
	ColorToggle                = 7
	ColorToggleHover           = 8
	ColorToggleCursor          = 9
	ColorSelect                = 10
	ColorSelectActive          = 11
	ColorSlider                = 12
	ColorSliderCursor          = 13
	ColorSliderCursorHover     = 14
	ColorSliderCursorActive    = 15
	ColorProperty              = 16
	ColorEdit                  = 17
	ColorEditCursor            = 18
	ColorCombo                 = 19
	ColorChart                 = 20
	ColorChartColor            = 21
	ColorChartColorHighlight   = 22
	ColorScrollbar             = 23
	ColorScrollbarCursor       = 24
	ColorScrollbarCursorHover  = 25
	ColorScrollbarCursorActive = 26
	ColorTabHeader             = 27
	ColorCount                 = 28
)

// StyleCursor as declared in nk/nuklear.h:3592
type StyleCursor int32

// StyleCursor enumeration from nk/nuklear.h:3592
const (
	CursorArrow                  = iota
	CursorText                   = 1
	CursorMove                   = 2
	CursorResizeVertical         = 3
	CursorResizeHorizontal       = 4
	CursorResizeTopLeftDownRight = 5
	CursorResizeTopRightDownLeft = 6
	CursorCount                  = 7
)

// FontCoordType as declared in nk/nuklear.h:3930
type FontCoordType int32

// FontCoordType enumeration from nk/nuklear.h:3930
const (
	CoordUv    = iota
	CoordPixel = 1
)

// FontAtlasFormat as declared in nk/nuklear.h:4004
type FontAtlasFormat int32

// FontAtlasFormat enumeration from nk/nuklear.h:4004
const (
	FontAtlasAlpha8 = iota
	FontAtlasRgba32 = 1
)

// AllocationType as declared in nk/nuklear.h:4101
type AllocationType int32

// AllocationType enumeration from nk/nuklear.h:4101
const (
	BufferFixed   = iota
	BufferDynamic = 1
)

// BufferAllocationType as declared in nk/nuklear.h:4106
type BufferAllocationType int32

// BufferAllocationType enumeration from nk/nuklear.h:4106
const (
	BufferFront = iota
	BufferBack  = 1
	BufferMax   = 2
)

// TextEditType as declared in nk/nuklear.h:4269
type TextEditType int32

// TextEditType enumeration from nk/nuklear.h:4269
const (
	TextEditSingleLine = iota
	TextEditMultiLine  = 1
)

// TextEditMode as declared in nk/nuklear.h:4274
type TextEditMode int32

// TextEditMode enumeration from nk/nuklear.h:4274
const (
	TextEditModeView    = iota
	TextEditModeInsert  = 1
	TextEditModeReplace = 2
)

// CommandType as declared in nk/nuklear.h:4374
type CommandType int32

// CommandType enumeration from nk/nuklear.h:4374
const (
	CommandTypeNop            = iota
	CommandTypeScissor        = 1
	CommandTypeLine           = 2
	CommandTypeCurve          = 3
	CommandTypeRect           = 4
	CommandTypeRectFilled     = 5
	CommandTypeRectMultiColor = 6
	CommandTypeCircle         = 7
	CommandTypeCircleFilled   = 8
	CommandTypeArc            = 9
	CommandTypeArcFilled      = 10
	CommandTypeTriangle       = 11
	CommandTypeTriangleFilled = 12
	CommandTypePolygon        = 13
	CommandTypePolygonFilled  = 14
	CommandTypePolyline       = 15
	CommandTypeText           = 16
	CommandTypeImage          = 17
	CommandTypeCustom         = 18
)

// CommandClipping as declared in nk/nuklear.h:4557
type CommandClipping int32

// CommandClipping enumeration from nk/nuklear.h:4557
const (
	ClippingOff = False
	ClippingOn  = True
)

// DrawListStroke as declared in nk/nuklear.h:4670
type DrawListStroke int32

// DrawListStroke enumeration from nk/nuklear.h:4670
const (
	StrokeOpen   = False
	StrokeClosed = True
)

// DrawVertexLayoutAttribute as declared in nk/nuklear.h:4677
type DrawVertexLayoutAttribute int32

// DrawVertexLayoutAttribute enumeration from nk/nuklear.h:4677
const (
	VertexPosition       = iota
	VertexColor          = 1
	VertexTexcoord       = 2
	VertexAttributeCount = 3
)

// DrawVertexLayoutFormat as declared in nk/nuklear.h:4684
type DrawVertexLayoutFormat int32

// DrawVertexLayoutFormat enumeration from nk/nuklear.h:4684
const (
	FormatSchar              = iota
	FormatSshort             = 1
	FormatSint               = 2
	FormatUchar              = 3
	FormatUshort             = 4
	FormatUint               = 5
	FormatFloat              = 6
	FormatDouble             = 7
	FormatColorBegin         = 8
	FormatR8g8b8             = FormatColorBegin
	FormatR16g15b16          = 9
	FormatR32g32b32          = 10
	FormatR8g8b8a8           = 11
	FormatB8g8r8a8           = 12
	FormatR16g15b16a16       = 13
	FormatR32g32b32a32       = 14
	FormatR32g32b32a32Float  = 15
	FormatR32g32b32a32Double = 16
	FormatRgb32              = 17
	FormatRgba32             = 18
	FormatColorEnd           = FormatRgba32
	FormatCount              = 19
)

// StyleItemType as declared in nk/nuklear.h:4805
type StyleItemType int32

// StyleItemType enumeration from nk/nuklear.h:4805
const (
	StyleItemColor = iota
	StyleItemImage = 1
)

// StyleHeaderAlign as declared in nk/nuklear.h:5146
type StyleHeaderAlign int32

// StyleHeaderAlign enumeration from nk/nuklear.h:5146
const (
	HeaderLeft  = iota
	HeaderRight = 1
)

// PanelType as declared in nk/nuklear.h:5252
type PanelType int32

// PanelType enumeration from nk/nuklear.h:5252
const (
	PanelNone       = iota
	PanelWindow     = (1 << (0))
	PanelGroup      = (1 << (1))
	PanelPopup      = (1 << (2))
	PanelContextual = (1 << (4))
	PanelCombo      = (1 << (5))
	PanelMenu       = (1 << (6))
	PanelTooltip    = (1 << (7))
)

// PanelSet as declared in nk/nuklear.h:5262
type PanelSet int32

// PanelSet enumeration from nk/nuklear.h:5262
const (
	PanelSetNonblock = PanelContextual | PanelCombo | PanelMenu | PanelTooltip
	PanelSetPopup    = PanelSetNonblock | PanelPopup
	PanelSetSub      = PanelSetPopup | PanelGroup
)

// PanelRowLayoutType as declared in nk/nuklear.h:5284
type PanelRowLayoutType int32

// PanelRowLayoutType enumeration from nk/nuklear.h:5284
const (
	LayoutDynamicFixed = iota
	LayoutDynamicRow   = 1
	LayoutDynamicFree  = 2
	LayoutDynamic      = 3
	LayoutStaticFixed  = 4
	LayoutStaticRow    = 5
	LayoutStaticFree   = 6
	LayoutStatic       = 7
	LayoutTemplate     = 8
	LayoutCount        = 9
)

// WindowFlags as declared in nk/nuklear.h:5352
type WindowFlags int32

// WindowFlags enumeration from nk/nuklear.h:5352
const (
	WindowPrivate        = (1 << (11))
	WindowDynamic        = WindowPrivate
	WindowRom            = (1 << (12))
	WindowNotInteractive = WindowRom | WindowNoInput
	WindowHidden         = (1 << (13))
	WindowClosed         = (1 << (14))
	WindowMinimized      = (1 << (15))
	WindowRemoveRom      = (1 << (16))
)

const (
	// False as declared in nk/nuklear.h:456
	False = iota
	// True as declared in nk/nuklear.h:456
	True = 1
)
