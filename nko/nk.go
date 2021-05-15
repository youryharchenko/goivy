package nk

import (
	"github.com/spaolacci/murmur3"
)

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

type PanelFlags uint32

const (
	NK_WINDOW_BORDER           PanelFlags = 1 << iota
	NK_WINDOW_MOVABLE                     = 1 << iota
	NK_WINDOW_SCALABLE                    = 1 << iota
	NK_WINDOW_CLOSABLE                    = 1 << iota
	NK_WINDOW_MINIMIZABLE                 = 1 << iota
	NK_WINDOW_NO_SCROLLBAR                = 1 << iota
	NK_WINDOW_TITLE                       = 1 << iota
	NK_WINDOW_SCROLL_AUTO_HIDE            = 1 << iota
	NK_WINDOW_BACKGROUND                  = 1 << iota
	NK_WINDOW_SCALE_LEFT                  = 1 << iota
	NK_WINDOW_NO_INPUT                    = 1 << iota
)

type WindowInsertLocation uint32

const (
	// inserts window into the back of list (front of screen)
	NK_INSERT_BACK WindowInsertLocation = iota
	// inserts window into the front of list (back of screen)
	NK_INSERT_FRONT
)

type CommandClipping bool

const (
	NK_CLIPPING_OFFCommandClipping = false
	NK_CLIPPING_ON                 = true
)

type WindowFlags uint32

const (
	NK_WINDOW_PRIVATE WindowFlags = 1 << 11 //NK_FLAG(11),
	// special window type growing up in height while being filled to a certain maximum height
	NK_WINDOW_DYNAMIC = NK_WINDOW_PRIVATE
	/* sets window widgets into a read only mode and does not allow input changes */
	NK_WINDOW_ROM = 1 << 12 //NK_FLAG(12),
	/* prevents all interaction caused by input to either window or widgets inside */
	NK_WINDOW_NOT_INTERACTIVE = NK_WINDOW_ROM | NK_WINDOW_NO_INPUT
	/* Hides window and stops any window interaction and drawing */
	NK_WINDOW_HIDDEN = 1 << 13 //NK_FLAG(13),
	/* Directly closes and frees the window at the end of the frame */
	NK_WINDOW_CLOSED = 1 << 14 //NK_FLAG(14),
	/* marks the window as minimized */
	NK_WINDOW_MINIMIZED = 1 << 15 //NK_FLAG(15),
	/* Removes read only mode at the end of the window */
	NK_WINDOW_REMOVE_ROM = 1 << 16 //NK_FLAG(16)

)

var NullRect = Rect{-8192.0, -8192.0, 16384, 16384}

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

func (ctx *Context) Begin(title string, bounds *Rect, flags Flags) bool {
	return ctx.BeginTitled(title, title, bounds, flags)
}

func (ctx *Context) BeginTitled(name string, title string, bounds *Rect, flags Flags) bool {
	if ctx.current != nil || len(title) == 0 || len(name) == 0 {
		return false
	}

	ret := false
	nameHash := murmurHash(name, NK_WINDOW_TITLE)
	//style := ctx.Style

	// find or create window
	win := ctx.findWindow(nameHash, name)
	if win == nil {
		// create new window
		win = ctx.createWindow()
		if flags&NK_WINDOW_BACKGROUND != 0 {
			ctx.insertWindow(win, NK_INSERT_FRONT)
		} else {
			ctx.insertWindow(win, NK_INSERT_BACK)
		}
		commandBufferInit(win.Buffer, ctx.Memory, NK_CLIPPING_ON)
		win.Flags = flags
		win.Bounds = bounds
		win.Name = nameHash
	} else {
		// update window
		win.Flags &= ^Flags(NK_WINDOW_PRIVATE - 1)
		win.Flags |= flags
		if win.Flags&(NK_WINDOW_MOVABLE|NK_WINDOW_SCALABLE) == 0 {
			win.Bounds = bounds
		}

		// If this assert triggers you either:
		//
		// I.) Have more than one window with the same name or
		// II.) You forgot to actually draw the window.
		//      More specific you did not call `nk_clear` (nk_clear will be
		//      automatically called for you if you are using one of the
		//      provided demo backends). */

		//NK_ASSERT(win->seq != ctx->seq);

		win.Seq = ctx.seq
		if ctx.active == nil && win.Flags&NK_WINDOW_HIDDEN == 0 {
			ctx.active = win
			ctx.end = win
		}
	}
	if win.Flags&NK_WINDOW_HIDDEN != 0 {
		ctx.current = win
		win.Layout = nil
		return false
	} else {
		ctx.start(win)
	}
	// window overlapping
	/*
		    if win.Flags & NK_WINDOW_HIDDEN==0 && win.Flags & NK_WINDOW_NO_INPUT==0    {
		        var inpanel, ishovered bool
		        iter := win
		        h := ctx.style.font.height + 2.0 * style.window.header.padding.y +
		            (2.0 * style.window.header.label_padding.y)

		        var winBounds Rect
				if win.Flags & NK_WINDOW_MINIMIZED == 0 {
		            winBounds = win.bounds
				} else {
					winBounds = Rect{win.bounds.x, win.bounds.y, win.bounds.w, h}
				}

		        // activate window if hovered and no other window is overlapping this window
		        inpanel = nk_input_has_mouse_click_down_in_rect(&ctx->input, NK_BUTTON_LEFT, win_bounds, nk_true);
		        inpanel = inpanel && ctx->input.mouse.buttons[NK_BUTTON_LEFT].clicked;
		        ishovered = nk_input_is_mouse_hovering_rect(&ctx->input, win_bounds);
		        if ((win != ctx->active) && ishovered && !ctx->input.mouse.buttons[NK_BUTTON_LEFT].down) {
		            iter = win->next;
		            while (iter) {
		                struct nk_rect iter_bounds = (!(iter->flags & NK_WINDOW_MINIMIZED))?
		                    iter->bounds: nk_rect(iter->bounds.x, iter->bounds.y, iter->bounds.w, h);
		                if (NK_INTERSECT(win_bounds.x, win_bounds.y, win_bounds.w, win_bounds.h,
		                    iter_bounds.x, iter_bounds.y, iter_bounds.w, iter_bounds.h) &&
		                    (!(iter->flags & NK_WINDOW_HIDDEN)))
		                    break;

		                if (iter->popup.win && iter->popup.active && !(iter->flags & NK_WINDOW_HIDDEN) &&
		                    NK_INTERSECT(win->bounds.x, win_bounds.y, win_bounds.w, win_bounds.h,
		                    iter->popup.win->bounds.x, iter->popup.win->bounds.y,
		                    iter->popup.win->bounds.w, iter->popup.win->bounds.h))
		                    break;
		                iter = iter->next;
		            }
		        }

		        // activate window if clicked
		        if (iter && inpanel && (win != ctx->end)) {
		            iter = win->next;
		            while (iter) {
		                // try to find a panel with higher priority in the same position
		                struct nk_rect iter_bounds = (!(iter->flags & NK_WINDOW_MINIMIZED))?
		                iter->bounds: nk_rect(iter->bounds.x, iter->bounds.y, iter->bounds.w, h);
		                if (NK_INBOX(ctx->input.mouse.pos.x, ctx->input.mouse.pos.y,
		                    iter_bounds.x, iter_bounds.y, iter_bounds.w, iter_bounds.h) &&
		                    !(iter->flags & NK_WINDOW_HIDDEN))
		                    break;
		                if (iter->popup.win && iter->popup.active && !(iter->flags & NK_WINDOW_HIDDEN) &&
		                    NK_INTERSECT(win_bounds.x, win_bounds.y, win_bounds.w, win_bounds.h,
		                    iter->popup.win->bounds.x, iter->popup.win->bounds.y,
		                    iter->popup.win->bounds.w, iter->popup.win->bounds.h))
		                    break;
		                iter = iter->next;
		            }
		        }
		        if (iter && !(win->flags & NK_WINDOW_ROM) && (win->flags & NK_WINDOW_BACKGROUND)) {
		            win->flags |= (nk_flags)NK_WINDOW_ROM;
		            iter->flags &= ~(nk_flags)NK_WINDOW_ROM;
		            ctx->active = iter;
		            if (!(iter->flags & NK_WINDOW_BACKGROUND)) {
		                // current window is active in that position so transfer to top
		                // at the highest priority in stack
		                nk_remove_window(ctx, iter);
		                nk_insert_window(ctx, iter, NK_INSERT_BACK);
		            }
		        } else {
		            if (!iter && ctx->end != win) {
		                if (!(win->flags & NK_WINDOW_BACKGROUND)) {
		                    // current window is active in that position so transfer to top
		                    // at the highest priority in stack
		                    nk_remove_window(ctx, win);
		                    nk_insert_window(ctx, win, NK_INSERT_BACK);
		                }
		                win->flags &= ~(nk_flags)NK_WINDOW_ROM;
		                ctx->active = win;
		            }
		            if (ctx->end != win && !(win->flags & NK_WINDOW_BACKGROUND))
		                win->flags |= NK_WINDOW_ROM;
		        }
		    }
	*/
	return ret
}

func (ctx *Context) createWindow() *Window {

	elem := ctx.createPageElement()
	if elem == nil {
		return nil
	}
	elem.Data.Win.Seq = ctx.seq
	return elem.Data.Win
}

func (ctx *Context) createPageElement() *PageElement {
	var elem *PageElement
	if ctx.freeList != nil {
		// unlink page element from free list
		elem = ctx.freeList
		ctx.freeList = elem.Next
	} else if ctx.usePool != 0 {
		// allocate page element from memory pool
		//elem = poolAlloc(ctx.pool)
		//NK_ASSERT(elem);
		if elem == nil {
			return nil
		}
	} else {
		// allocate new page element from back of fixed size memory buffer
		//NK_STORAGE const nk_size size = sizeof(struct nk_page_element);
		//NK_STORAGE const nk_size align = NK_ALIGNOF(struct nk_page_element);
		//elem = (struct nk_page_element*)nk_buffer_alloc(&ctx->memory, NK_BUFFER_BACK, size, align);
		//NK_ASSERT(elem);
		//if (!elem) return 0;
	}
	//nk_zero_struct(*elem)
	elem.Next = nil
	elem.Prev = nil
	return elem
}

func (ctx *Context) insertWindow(win *Window, loc WindowInsertLocation) {
	//NK_ASSERT(ctx);
	//NK_ASSERT(win);
	if win == nil {
		return
	}
	iter := ctx.begin
	for iter != nil {
		//NK_ASSERT(iter != iter->next);
		//NK_ASSERT(iter != win);
		if iter == win {
			return
		}
		iter = iter.Next
	}
	if ctx.begin == nil {
		win.Next = nil
		win.Prev = nil
		ctx.begin = win
		ctx.end = win
		ctx.count = 1
		return
	}
	if loc == NK_INSERT_BACK {
		end := ctx.end
		end.Flags |= NK_WINDOW_ROM
		end.Next = win
		win.Prev = ctx.end
		win.Next = nil
		ctx.end = win
		ctx.active = ctx.end
		ctx.end.Flags &= ^Flags(NK_WINDOW_ROM)
	} else {
		/*ctx->end->flags |= NK_WINDOW_ROM;*/
		ctx.begin.Prev = win
		win.Next = ctx.begin
		win.Prev = nil
		ctx.begin = win
		ctx.begin.Flags &= ^Flags(NK_WINDOW_ROM)
	}
	ctx.count++
}

func (ctx *Context) start(win *Window) {
	//NK_ASSERT(ctx);
	//NK_ASSERT(win);
	ctx.startBuffer(win.Buffer)
}

func (ctx *Context) startBuffer(buffer *CommandBuffer) {
	//NK_ASSERT(ctx);
	//NK_ASSERT(buffer);

	if buffer == nil {
		return
	}

	buffer.begin = ctx.Memory.Allocated
	buffer.end = buffer.begin
	buffer.last = buffer.begin
	buffer.Clip = &NullRect
}

func (ctx *Context) findWindow(hash Hash, name string) *Window {
	iter := ctx.begin
	for iter != nil {
		//NK_ASSERT(iter != iter->next);
		if iter.Name == hash {
			if iter.NameString == name {
				return iter
			}
		}
		iter = iter.Next
	}
	return nil
}

type Vec2 struct {
	X, Y float64
}

type Rect struct {
	X, Y, W, H float64
}

type Color struct {
	r, g, b, a byte
}

type Scroll struct {
	X, Y Uint
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
	UseClipping      bool
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

func murmurHash(name string, seed Hash) Hash {
	return Hash(murmur3.Sum32WithSeed([]byte(name), uint32(seed)))
}

func commandBufferInit(cb *CommandBuffer, b *Buffer, clip CommandClipping) {
	//NK_ASSERT(cb);
	//NK_ASSERT(b);
	if cb == nil || b == nil {
		return
	}
	cb.Base = b
	cb.UseClipping = bool(clip)
	cb.begin = b.Allocated
	cb.end = b.Allocated
	cb.last = b.Allocated
}
