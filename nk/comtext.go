package nk

type ButtonBehavior uint32

const (
	NK_BUTTON_DEFAULT ButtonBehavior = iota
	NK_BUTTON_REPEATER
)

type Uint uint32

type Hash Uint
type Flags Uint
type Rune Uint

type Str string

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
	scrollbar *Vec2
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
	Name       nk_hash
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
