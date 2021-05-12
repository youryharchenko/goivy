package nk

type ButtonBehavior uint32

const (
	NK_BUTTON_DEFAULT ButtonBehavior = iota
	NK_BUTTON_REPEATER
)

type AllocationType uint32

const (
	NK_BUFFER_FIXED AllocationType = iota
	NK_BUFFER_DYNAMIC
)

type StyleItemType uint32

const (
	NK_STYLE_ITEM_COLOR StyleItemType = iota
	NK_STYLE_ITEM_IMAGE
)

type AntiAliasing uint32

const (
	NK_ANTI_ALIASING_OFF AntiAliasing = iota
	NK_ANTI_ALIASING_ON
)

type DrawVertexLayoutAttribute uint32

const (
	NK_VERTEX_POSITION DrawVertexLayoutAttribute = iota
	NK_VERTEX_COLOR
	NK_VERTEX_TEXCOORD
	NK_VERTEX_ATTRIBUTE_COUNT
)

type DrawVertexLayoutFormat uint32

const (
	NK_FORMAT_SCHAR = iota
	NK_FORMAT_SSHORT
	NK_FORMAT_SINT
	NK_FORMAT_UCHAR
	NK_FORMAT_USHORT
	NK_FORMAT_UINT
	NK_FORMAT_FLOAT
	NK_FORMAT_DOUBLE

	NK_FORMAT_COLOR_BEGIN
	NK_FORMAT_R8G8B8    = NK_FORMAT_COLOR_BEGIN
	NK_FORMAT_R16G15B16 = iota - 1
	NK_FORMAT_R32G32B32

	NK_FORMAT_R8G8B8A8
	NK_FORMAT_B8G8R8A8
	NK_FORMAT_R16G15B16A16
	NK_FORMAT_R32G32B32A32
	NK_FORMAT_R32G32B32A32_FLOAT
	NK_FORMAT_R32G32B32A32_DOUBLE

	NK_FORMAT_RGB32
	NK_FORMAT_RGBA32
	NK_FORMAT_COLOR_END = NK_FORMAT_RGBA32
	NK_FORMAT_COUNT     = iota - 2
)

type SymbolType uint32

const (
	NK_SYMBOL_NONE SymbolType = iota
	NK_SYMBOL_X
	NK_SYMBOL_UNDERSCORE
	NK_SYMBOL_CIRCLE_SOLID
	NK_SYMBOL_CIRCLE_OUTLINE
	NK_SYMBOL_RECT_SOLID
	NK_SYMBOL_RECT_OUTLINE
	NK_SYMBOL_TRIANGLE_UP
	NK_SYMBOL_TRIANGLE_DOWN
	NK_SYMBOL_TRIANGLE_LEFT
	NK_SYMBOL_TRIANGLE_RIGHT
	NK_SYMBOL_PLUS
	NK_SYMBOL_MINUS
	NK_SYMBOL_MAX
)

type StyleHeaderAlign uint32

const (
	NK_HEADER_LEFT StyleHeaderAlign = iota
	NK_HEADER_RIGHT
)

type PanelType uint32

const (
	NK_PANEL_NONE       PanelType = 0
	NK_PANEL_WINDOW               = 1 << (iota - 1) //NK_FLAG(0)
	NK_PANEL_GROUP                = 1 << (iota - 1) //NK_FLAG(1)
	NK_PANEL_POPUP                = 1 << (iota - 1) //NK_FLAG(2)
	NK_PANEL_CONTEXTUAL           = 1 << (iota)     //NK_FLAG(4)
	NK_PANEL_COMBO                = 1 << (iota)     //NK_FLAG(5)
	NK_PANEL_MENU                 = 1 << (iota)     //NK_FLAG(6)
	NK_PANEL_TOOLTIP              = 1 << (iota)     //NK_FLAG(7)
)

type PanelRowLayoutType uint32

const (
	NK_LAYOUT_DYNAMIC_FIXED PanelRowLayoutType = iota
	NK_LAYOUT_DYNAMIC_ROW
	NK_LAYOUT_DYNAMIC_FREE
	NK_LAYOUT_DYNAMIC
	NK_LAYOUT_STATIC_FIXED
	NK_LAYOUT_STATIC_ROW
	NK_LAYOUT_STATIC_FREE
	NK_LAYOUT_STATIC
	NK_LAYOUT_TEMPLATE
	NK_LAYOUT_COUNT
)

type ChartType uint32

const (
	NK_CHART_LINES ChartType = iota
	NK_CHART_COLUMN
	NK_CHART_MAX
)

type Uint uint32
type Float float64

type Size uint64
type Handle uint64

type Hash Uint
type Flags Uint
type Rune rune

type Str string

//typedef void*(*nk_plugin_alloc)(nk_handle, void *old, nk_size);
type PluginAlloc func(Handle, interface{}, Size) interface{}

//typedef void (*nk_plugin_free)(nk_handle, void *old);
type PluginFree func(Handle, interface{}, Size)

//typedef int(*nk_plugin_filter)(const struct nk_text_edit*, nk_rune unicode);
type PluginFilter func(*TextEdit, Rune) int

//typedef void(*nk_plugin_paste)(nk_handle, struct nk_text_edit*);
type PluginPaste func(Handle, *TextEdit)

//typedef void(*nk_plugin_copy)(nk_handle, const char*, int len);
type PluginCopy func(Handle, string, int)

//typedef float(*nk_text_width_f)(nk_handle, float h, const char*, int len)
type TextWidthF func(Handle, float64, string, int) float64

//typedef void(*nk_query_font_glyph_f)(nk_handle handle, float font_height, struct nk_user_font_glyph *glyph, nk_rune codepoint, nk_rune next_codepoint);
type QueryFontGlyphF func(Handle, float64, *UserFontGlyph, Rune, Rune)

type Context struct {
	// public: can be accessed freely
	Input            *Input
	Style            *Style
	Memory           *Buffer
	Clip             *Clipboard
	LastWidgetState  Flags
	ButtonBehavior   ButtonBehavior
	Stacks           *ConfigurationStacks
	DeltaTimeSeconds float64

	// private: hould only be accessed if you know what you are doing
	drawList *DrawList
	//nk_handle userdata;

	// text editor objects are quite big because of an internal
	// undo/redo stack. Therefore it does not make sense to have one for
	// each window for temporary use cases, so I only provide *one* instance
	// for all windows. This works because the content is cleared anyway
	textEdit *TextEdit
	// draw buffer used for overlay drawing operation like cursor
	overlay *CommandBuffer

	// windows
	build    int
	usePool  int
	pool     *Pool
	begin    *Window
	end      *Window
	active   *Window
	current  *Window
	freeList *PageElement
	count    uint
	seq      uint
}

type Vec2 struct {
	x, y float64
}

type Rect struct {
	x, y, w, h float64
}

type Color struct {
	r, g, b, a byte
}

type Scroll struct {
	x, y Uint
}

type Input struct {
	Keyboard *Keyboard
	Mouse    *Mouse
}

type Style struct {
	Font          *UserFont
	Cursors       []Cursor
	CursorActive  *Cursor
	CursorLast    *Cursor
	cursorVisible int
	//
	Text             *StyleText
	Button           *StyleButton
	ContextualButton *StyleButton
	MenuButton       *StyleButton
	Option           *StyleToggle
	Checkbox         *StyleToggle
	Selectable       *StyleSelectable
	Slider           *StyleSlider
	Progress         *StyleProgress
	Property         *StyleProperty
	Edit             *StyleEdit
	Chart            *StyleChart
	ScrollH          *StyleScrollbar
	ScrollV          *StyleScrollbar
	Tab              *StyleTab
	Combo            *StyleCombo
	Window           *StyleWindow
}

type Buffer struct {
	// buffer marker to free a buffer to a certain offset
	Marker []BufferMarker
	// allocator callback for dynamic buffers
	Pool *Allocator
	// memory management type
	Type AllocationType
	// memory and size of the current memory block
	Memory *Memory
	// growing factor for dynamic memory management
	GrowFactor float64
	// total amount of memory allocated
	Allocated Size
	// totally consumed memory given that enough memory is present
	Needed Size
	// number of allocation calls
	Calls Size
	// current size of the buffer
	Size Size
}

type Clipboard struct {
	Userdata Handle
	Paste    PluginPaste
	Copy     PluginCopy
}

type ConfigurationStacks struct {
	StyleItems      *ConfigStackStyleItem
	Floats          *ConfigStackFloat
	Vectors         *ConfigStackVec2
	Flags           *ConfigStackFlags
	Colors          *ConfigStackColor
	Fonts           *ConfigStackUserFont
	ButtonBehaviors *ConfigStackButtonBehavior
}

type ConfigStackStyleItemElement struct {
	Address  *StyleItem
	OldValue StyleItem
}
type ConfigStackStyleItem struct {
	Head     int
	Elements []ConfigStackStyleItemElement
}
type ConfigStackFloatElement struct {
	Address  *Float
	OldValue Float
}
type ConfigStackFloat struct {
	Head     int
	Elements []ConfigStackFloatElement
}
type ConfigStackFlagsElement struct {
	Address  *Flags
	OldValue Flags
}
type ConfigStackFlags struct {
	Head     int
	Elements []ConfigStackFlagsElement
}
type ConfigStackVec2Element struct {
	Address  *Vec2
	OldValue Vec2
}
type ConfigStackVec2 struct {
	Head     int
	Elements []ConfigStackVec2Element
}
type ConfigStackColorElement struct {
	Address  *Color
	OldValue Color
}
type ConfigStackColor struct {
	Head     int
	Elements []ConfigStackColorElement
}
type ConfigStackUserFontElement struct {
	Address  *UserFont
	OldValue UserFont
}
type ConfigStackUserFont struct {
	Head     int
	Elements []ConfigStackUserFontElement
}
type ConfigStackButtonBehaviorElement struct {
	Address  *ButtonBehavior
	OldValue ButtonBehavior
}
type ConfigStackButtonBehavior struct {
	Head     int
	Elements []ConfigStackButtonBehaviorElement
}

type ConvertConfig struct {
	// global alpha value
	GlobalAlpha float64
	// line anti-aliasing flag can be turned off if you are tight on memory
	LineAA AntiAliasing
	// shape anti-aliasing flag can be turned off if you are tight on memory
	ShapeAA AntiAliasing
	// number of segments used for circles: default to 22
	CircleSegmentCount uint
	// number of segments used for arcs: default to 22
	ArcSegmentCount uint
	// number of segments used for curves: default to 22
	CurveSegmentCount uint
	// handle to texture with a white pixel for shape drawing
	Null *DrawNullTexture
	// describes the vertex output format and packing
	VertexLayout *DrawVertexLayoutElement
	// sizeof one vertex for vertex packing
	VertexSize Size
	// vertex alignment: Can be obtained by NK_ALIGNOF
	VertexAlignment Size
}

type DrawNullTexture struct {
	// texture handle to a texture with a white pixel
	Texture Handle
	// coordinates to a white pixel in the texture
	uv *Vec2
}

type DrawVertexLayoutElement struct {
	Attribute DrawVertexLayoutAttribute
	Format    DrawVertexLayoutFormat
	Offset    Size
}

type DrawList struct {
	ClipRect  *Rect
	CircleVtx [12]Vec2
	Config    *ConvertConfig
	//
	Buffer   *Buffer
	Vertices *Buffer
	Elements *Buffer
	//
	ElementCount uint
	VertexCount  uint
	CmdCount     uint
	CmdOffset    Size
	//
	PathCount  uint
	PathOffset uint
	//
	LineAA  AntiAliasing
	ShapeAA AntiAliasing
	//
	// nk_handle userdata;
}

type TextEdit struct {
	Clip      *Clipboard
	String    Str
	Filter    PluginFilter
	Scrollbar *Vec2
	//
	cursor            int
	SelectStart       int
	SelectEnd         int
	Mode              byte
	CursorAtEndOfLine byte
	Initialized       byte
	HasPreferred_x    byte
	SingleLine        byte
	Active            byte
	Padding1          byte
	PreferredX        float64
	Undo              *TextUndoState
}

type CommandBuffer struct {
	Base             *Buffer
	Clip             *Rect
	UseClipping      int
	Userdata         Handle
	begin, end, last Size
}

type Pool struct {
	Alloc     *Allocator
	Type      AllocationType
	PageCount int
	Pages     *Page
	FreeList  *PageElement
	Capacity  uint
	Size      Size
	Cap       Size
}

type Window struct {
	Seq        uint
	Name       Hash
	NameString string
	Flags      Flags
	//
	Bounds               *Rect
	Scrollbar            *Scroll
	Buffer               *CommandBuffer
	Layout               *Panel
	ScrollbarHidingTimer float64
	// persistent widget state
	Property *PropertyState
	Popup    *PopupState
	Edit     *EditState
	Scrolled uint
	//
	Tables     *Table
	TableCount uint
	// window list hooks
	Next   *Window
	Prev   *Window
	Parent *Window
}

type PageElement struct {
	Data *PageData
	Next *PageElement
	Prev *PageElement
}

type Keyboard struct {
	Keys    []Key
	Text    string
	TextLen int
}

type Mouse struct {
	Buttons     []MouseButton
	Pos         *Vec2
	Prev        *Vec2
	Delta       *Vec2
	ScrollDelta *Vec2
	Grab        byte
	Grabbed     byte
	Ungrab      byte
}

type UserFont struct {
	// user provided font handle
	Userdata Handle
	// max height of the font
	Height float64
	// font string width in pixel callback
	Width TextWidthF
	//
	//#ifdef NK_INCLUDE_VERTEX_BUFFER_OUTPUT
	/* font glyph callback to query drawing info */
	Query QueryFontGlyphF
	/* texture handle to the used font atlas or texture */
	Texture Handle
	//#endif
}

type Cursor struct {
	Img          *Image
	Size, Offset *Vec2
}

type StyleItemData struct {
	Image *Image
	Color *Color
}

type StyleItem struct {
	Type StyleItemType
	Data StyleItemData
}

type StyleText struct {
	Color   *Color
	Padding *Vec2
}

type StyleButton struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	// text
	TextBackground *Color
	TextNormal     *Color
	TextHover      *Color
	TextActive     *Color
	TextAlignment  Flags
	// properties
	Border       float64
	Rounding     float64
	Padding      *Vec2
	ImagePadding *Vec2
	TouchPadding *Vec2
	// optional user callbacks
	Userdata Handle
	//void(*draw_begin)(struct nk_command_buffer*, nk_handle userdata);
	DrawBegin func(*CommandBuffer, Handle)
	//void(*draw_end)(struct nk_command_buffer*, nk_handle userdata);
	DrawEnd func(*CommandBuffer, Handle)
}

type StyleToggle struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	// cursor
	CursorNormal *StyleItem
	CursorHover  *StyleItem
	// text
	TextBackground *Color
	TextNormal     *Color
	TextHover      *Color
	TextActive     *Color
	TextAlignment  Flags
	// properties
	Padding      *Vec2
	TouchPadding *Vec2
	Spacing      float64
	Border       float64
	// optional user callbacks
	Userdata Handle
	//void(*draw_begin)(struct nk_command_buffer*, nk_handle userdata);
	DrawBegin func(*CommandBuffer, Handle)
	//void(*draw_end)(struct nk_command_buffer*, nk_handle userdata);
	DrawEnd func(*CommandBuffer, Handle)
}

type StyleSelectable struct {
	// background (inactive)
	Normal  *StyleItem
	Hover   *StyleItem
	Pressed *StyleItem
	// background (active)
	NormalActive  *StyleItem
	HoverActive   *StyleItem
	PressedActive *StyleItem
	// text color (inactive)
	TextNormal  *Color
	TextHover   *Color
	TextPressed *Color
	// text color (active)
	TextNormalActive     *Color
	TextHoverActive      *Color
	TextPressedActive    *Color
	TextBackgroundActive *Color
	TextAlignment        Flags
	// properties
	Rounding     float64
	Padding      *Vec2
	ImagePadding *Vec2
	TouchPadding *Vec2
	// optional user callbacks
	Userdata Handle
	//void(*draw_begin)(struct nk_command_buffer*, nk_handle userdata);
	DrawBegin func(*CommandBuffer, Handle)
	//void(*draw_end)(struct nk_command_buffer*, nk_handle userdata);
	DrawEnd func(*CommandBuffer, Handle)
}

type StyleSlider struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	// background bar
	BarNormal *Color
	BarHover  *Color
	BarActive *Color
	BarFilled *Color
	// cursor
	CursorNormal *StyleItem
	CursorHover  *StyleItem
	CursorActive *StyleItem
	// properties
	Border     float64
	Rounding   float64
	BarHeight  float64
	Padding    *Vec2
	Spacing    *Vec2
	CursorSize *Vec2
	// optional buttons
	ShowButtons int
	IncButton   *StyleButton
	DecButton   *StyleButton
	IncSymbol   SymbolType
	DecSymbol   SymbolType
	// optional user callbacks
	Userdata Handle
	//void(*draw_begin)(struct nk_command_buffer*, nk_handle userdata);
	DrawBegin func(*CommandBuffer, Handle)
	//void(*draw_end)(struct nk_command_buffer*, nk_handle userdata);
	DrawEnd func(*CommandBuffer, Handle)
}

type StyleProgress struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	// cursor
	CursorNormal      *StyleItem
	CursorHover       *StyleItem
	CursorActive      *StyleItem
	CursorBorderColor *Color
	// properties
	Border         float64
	Rounding       float64
	CursorBorder   float64
	CursorRounding float64
	Padding        *Vec2
	// optional user callbacks
	Userdata Handle
	//void(*draw_begin)(struct nk_command_buffer*, nk_handle userdata);
	DrawBegin func(*CommandBuffer, Handle)
	//void(*draw_end)(struct nk_command_buffer*, nk_handle userdata);
	DrawEnd func(*CommandBuffer, Handle)
}

type StyleProperty struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	// text
	LabelNormal  *Color
	LabelHover   *Color
	LabelPressed *Color
	// symbols
	SymLeft  SymbolType
	SymRight SymbolType
	// properties
	Border   float64
	Rounding float64
	Padding  *Vec2
	// edit
	Edit      *StyleEdit
	IncButton *StyleButton
	DecButton *StyleButton
	// optional user callbacks
	Userdata Handle
	//void(*draw_begin)(struct nk_command_buffer*, nk_handle userdata);
	DrawBegin func(*CommandBuffer, Handle)
	//void(*draw_end)(struct nk_command_buffer*, nk_handle userdata);
	DrawEnd func(*CommandBuffer, Handle)
}

type StyleEdit struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	Scrollbar   *StyleScrollbar
	// cursor
	CursorNormal     *Color
	CursorHover      *Color
	CursorTextNormal *Color
	CursorTextHover  *Color
	// text color (unselected)
	TextNormal  *Color
	TextHover   *Color
	TextPressed *Color
	// text color (selected)
	SelectedNormal     *Color
	SelectedHover      *Color
	SelectedTextNormal *Color
	SelectedTextHover  *Color
	// properties
	Border        float64
	Rounding      float64
	RowPadding    float64
	CursorSize    float64
	Padding       *Vec2
	ScrollbarSize *Vec2
}

type StyleChart struct {
	// colors
	Background    *StyleItem
	BorderColor   *Color
	SelectedColor *Color
	Color         *Color
	// properties
	Border   float64
	Rounding float64
	Padding  *Vec2
}

type StyleScrollbar struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	// cursor
	CursorNormal      *StyleItem
	CursorHover       *StyleItem
	CursorActive      *StyleItem
	CursorBorderColor *Color
	// properties
	Border         float64
	Rounding       float64
	BorderCursor   float64
	RoundingCursor float64
	Padding        *Vec2
	// optional buttons
	ShowButtons int
	IncButton   *StyleButton
	DecButton   *StyleButton
	IncSymbol   SymbolType
	DecSymbol   SymbolType
	// optional user callbacks
	Userdata Handle
	//void(*draw_begin)(struct nk_command_buffer*, nk_handle userdata);
	DrawBegin func(*CommandBuffer, Handle)
	//void(*draw_end)(struct nk_command_buffer*, nk_handle userdata);
	DrawEnd func(*CommandBuffer, Handle)
}

type StyleTab struct {
	// background
	Background  *StyleItem
	BorderColor *Color
	TextColor   *Color
	// button
	TabMaximizeButton  *StyleButton
	TabMinimizeButton  *StyleButton
	NodeMaximizeButton *StyleButton
	NodeMinimizeButton *StyleButton
	SymMaximize        SymbolType
	SymMinimize        SymbolType
	// properties
	Border   float64
	Rounding float64
	Indent   float64
	Padding  *Vec2
	Spacing  *Vec2
}

type StyleCombo struct {
	// background
	Normal      *StyleItem
	Hover       *StyleItem
	Active      *StyleItem
	BorderColor *Color
	// label
	LabelNormal *Color
	LabelHover  *Color
	LabelActive *Color
	// symbol
	SymbolNormal *Color
	SymbolHover  *Color
	SymbolActive *Color
	// button
	Button    *StyleButton
	SymNormal SymbolType
	SymHover  SymbolType
	SymActive SymbolType
	// properties
	Border         float64
	Rounding       float64
	ContentPadding *Vec2
	ButtonPadding  *Vec2
	Spacing        *Vec2
}

type StyleWindowHeader struct {
	// background
	Normal *StyleItem
	Hover  *StyleItem
	Active *StyleItem
	// button
	CloseButton    *StyleButton
	MinimizeButton *StyleButton
	MaximizeButton *StyleButton
	CloseSymbol    SymbolType
	MaximizeSymbol SymbolType
	MinimizeSymbol SymbolType
	// title
	LabelNormal *Color
	LabelHover  *Color
	LabelActive *Color
	// properties
	Align        StyleHeaderAlign
	Padding      *Vec2
	LabelPadding *Vec2
	Spacing      *Vec2
}

type StyleWindow struct {
	Header          *StyleWindowHeader
	FixedBackground *StyleItem
	Background      *Color
	// border color
	BorderColor           *Color
	PopupBorderColor      *Color
	ComboBorderColor      *Color
	ContextualBorderColor *Color
	MenuBorderColor       *Color
	GroupBorderColor      *Color
	TooltipBorderColor    *Color
	Scaler                *StyleItem
	// border
	Border              float64
	ComboBorder         float64
	ContextualBorder    float64
	MenuBorder          float64
	GroupBorder         float64
	TooltipBorder       float64
	PopupBorder         float64
	MinRowHeightPadding float64
	// properties
	Rounding      float64
	Spacing       *Vec2
	ScrollbarSize *Vec2
	MinSize       *Vec2
	// padding
	GroupPadding      *Vec2
	PopupPadding      *Vec2
	ComboPadding      *Vec2
	ContextualPadding *Vec2
	MenuPadding       *Vec2
	TooltipPadding    *Vec2
	Padding           *Vec2
}

type BufferMarker struct {
	Active int
	Offset Size
}

type Allocator struct {
	Userdata Handle
	Alloc    PluginAlloc
	Free     PluginFree
}

type Memory struct {
	//void *ptr;
	Ptr  interface{}
	Size Size
}

type TextUndoState struct {
	UndoRec       []TextUndoRecord
	UndoChar      []Rune
	UndoPoint     int16
	RedoPoint     int16
	UndoCharPoint int16
	RedoCharPoint int16
}

type TextUndoRecord struct {
	Where        int
	InsertLength int16
	DeleteLength int16
	CharStorage  int16
}

type Page struct {
	Size uint
	Next *Page
	Win  [1]PageElement
}

type Panel struct {
	Type           PanelType
	Flags          Flags
	bounds         *Rect
	OffsetX        *Uint
	OffsetY        *Uint
	AtX, AtY, MaxX float64
	FooterHeight   float64
	HeaderHeight   float64
	border         float64
	HasScrolling   uint
	Clip           *Rect
	Menu           *MenuState
	Row            *RowLayout
	Chart          *Chart
	Buffer         *CommandBuffer
	Parent         *Panel
}

type PropertyState struct {
	Active, Prev int
	Buffer       []byte
	Length       int
	Cursor       int
	SelectStart  int
	SelectEnd    int
	Name         Hash
	Seq          uint
	Old          uint
	State        int
}

type PopupState struct {
	Win              *Window
	Type             PanelType
	Buf              *PopupBuffer
	Name             Hash
	Active           int
	ComboCount       uint
	ConCount, ConOld uint
	ActiveCon        uint
	header           *Rect
}

type EditState struct {
	Name         Hash
	Seq          uint
	Old          uint
	Active, Prev int
	Cursor       int
	SelStart     int
	SelEnd       int
	Scrollbar    *Scroll
	Mode         byte
	SingleLine   byte
}

type MenuState struct {
	X, Y, W, H float64
	Offset     *Scroll
}

type Table struct {
	seq        uint
	size       uint
	Keys       []Hash
	Values     []Uint
	Next, Prev *Table
}

type RowLayout struct {
	Type       PanelRowLayoutType
	Index      int
	Height     float64
	MinHeight  float64
	Columns    int
	Ratio      *float64
	ItemWidth  float64
	ItemHeight float64
	ItemOffset float64
	Filled     float64
	Item       *Rect
	TreeDepth  int
	Templates  []float64
}

type PageData struct {
	Tbl *Table
	Pan *Panel
	Win *Window
}

type Key struct {
	Down    int
	Clicked uint
}

type MouseButton struct {
	Down       int
	Clicked    uint
	ClickedPos *Vec2
}

type UserFontGlyph struct {
	// texture coordinates
	Uv [2]Vec2
	// offset between top left and glyph
	Offset *Vec2
	// size of the glyph
	Width, Height float64
	// offset to the next glyph
	Xadvance float64
}

type Image struct {
	Handle Handle
	W, H   uint16
	region [4]uint16
}

type Chart struct {
	Slot       int
	X, Y, W, H float64
	Slots      []ChartSlot
}

type ChartSlot struct {
	Type            ChartType
	Color           *Color
	Highlight       *Color
	Min, Max, Range float64
	Count           int
	Last            *Vec2
	Index           int
}

type PopupBuffer struct {
	Begin  Size
	Parent Size
	Last   Size
	End    Size
	Active int
}
