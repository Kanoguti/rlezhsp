package main

import "C"

import (
	"bytes"
	"encoding/binary"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"strings"
	"time"
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type eface struct {
	_type unsafe.Pointer
	data  unsafe.Pointer
}

type Resource struct {
	original_data interface{}
	data          unsafe.Pointer
	life          bool
	type_name     string
	type_size     int32
	group         string
}

type StringMemory struct {
	original_string string
	string_pointer  unsafe.Pointer
	string_uintptr  uintptr
}

type RLEZMesh struct {
	Mesh            rl.Mesh
	Material        rl.Material
	IsFromMesh      bool
	IsModel         bool
	Model           rl.Model
	ModelTextures   [][]rl.Texture2D
	ModelAnimations []rl.ModelAnimation
}

type RLEZFont struct {
	Codepoints      []rune
	GlyphInfosCount []int32
	GlyphInfos      []rl.GlyphInfo
	FontSize        int32
	Fonts           []rl.Font
}

type RLEZSound struct {
	IsMusic         bool
	MusicBytes      []byte
	MusicUpdateTime float64
	SoundData       rl.Sound
	MusicData       rl.Music
	TickerFunc      func(data *RLEZSound)
	TickerTime      int64
	TickerStatus    bool
	ChannelSeek     chan float32
	ChannelMessage  chan string
	Time            float32
}

type System struct {
	window_status        bool
	window_config        uint32
	shape                int32
	shape_auto_normal    bool
	shape_vertex         [][]float32
	shape_texture        int32
	shape_texture_width  int32
	shape_texture_height int32
	shape_texture_flip   bool
	resource             []Resource
	return_type          string
	return_data          interface{}
	return_size          int32
	draw_mode            string
	default_material     rl.Material
	color                rl.Color
	pixels               interface{} //rl.Image
	pixels_resource      int32
}

var system System

func toString(p uintptr) string {
	return C.GoString((*C.char)(unsafe.Pointer(p)))
}

func fromString(str string) StringMemory {
	return_data := StringMemory{}
	return_data.original_string = "" + str + "\x00"
	return_data.string_pointer = unsafe.Pointer(unsafe.StringData(return_data.original_string))
	return_data.string_uintptr = uintptr(return_data.string_pointer)
	return return_data
}

func getColor(r, g, b, a int32) rl.Color {
	return rl.GetColor(uint(uint32(r<<24) | uint32(g<<16) | uint32(b<<8) | uint32(a<<0)))
}

func getTypeInfo(data interface{}, type_data reflect.Type) (string, int32) {
	return_name := ""
	return_size := int32(0)

	var data_check reflect.Type

	if type_data == nil {
		data_check = reflect.ValueOf(data).Type()
	} else {
		data_check = type_data
	}
	return_name = data_check.Name()
	switch return_name {
	case "string":
		return_size = int32(len(data.(string)))
	default:
		return_size = int32(data_check.Size())
	}
	return return_name, return_size
}

func setReturnData(data interface{}, type_data reflect.Type) {
	system.return_data = data
	system.return_type, system.return_size = getTypeInfo(data, type_data)
}

func clearResource(id int32) {
	if checkResource(id) == true {
		system.resource[id].data = nil
		system.resource[id].life = false
		system.resource[id].type_name = ""
		system.resource[id].type_size = 0
		system.resource[id].group = ""
	}
	check := int32(-1)
	for cnt := int32(len(system.resource) - 1); cnt >= int32(0); cnt-- {
		if system.resource[cnt].life == true {
			check = cnt
			break
		}
	}
	if check >= 0 {
		system.resource = slices.Delete(system.resource, int(check)+1, len(system.resource))
	} else {
		system.resource = []Resource{}
	}
}

func checkResource(id int32) bool {
	return_data := false
	if 0 <= id && id < int32(len(system.resource)) {
		if system.resource[id].life == true {
			return_data = true
		}
	}
	return return_data
}

func getResource(id int32) unsafe.Pointer {
	var return_data unsafe.Pointer
	if checkResource(id) == true {
		return_data = (*eface)(system.resource[id].data).data
	}
	return return_data
}

func interfaceToPointer(data interface{}) unsafe.Pointer {
	var return_data unsafe.Pointer
	return_data = (*eface)(unsafe.Pointer(reflect.ValueOf(&data).Pointer())).data
	return return_data
}

func registerResource(data interface{}, type_data reflect.Type) int32 {
	return_data := len(system.resource)
	append_flag := true
	for cnt := 0; cnt < len(system.resource); cnt++ {
		if system.resource[cnt].life == false {
			return_data = cnt
			append_flag = false
			break
		}
	}
	var target_resource Resource
	target_resource.original_data = data
	target_resource.data = interfaceToPointer(target_resource.original_data)
	target_resource.life = true
	target_resource.type_name, target_resource.type_size = type_data.Name(), int32(type_data.Size())
	target_resource.group = ""
	if append_flag == true {
		system.resource = append(system.resource, target_resource)
	} else {
		system.resource[return_data] = target_resource
	}
	return int32(return_data)
}

func getTexture(id int32) *rl.Texture2D {
	return_data := &rl.Texture2D{}
	if checkResource(id) == true {
		switch system.resource[id].type_name {
		case "Texture2D":
			return_data = (*rl.Texture2D)(system.resource[id].data)
		case "RenderTexture2D":
			return_data = &((*rl.RenderTexture2D)(system.resource[id].data).Texture)
		}
	}
	return return_data
}

func getWindowFlag(name string) uint32 {
	return_data := uint32(0)
	switch strings.ToUpper(name) {
	case "FLAG_VSYNC_HINT":
		return_data = rl.FlagVsyncHint
	case "FLAG_FULLSCREEN_MODE":
		return_data = rl.FlagFullscreenMode
	case "FLAG_WINDOW_RESIZABLE":
		return_data = rl.FlagWindowResizable
	case "FLAG_WINDOW_UNDECORATED":
		return_data = rl.FlagWindowUndecorated
	case "FLAG_WINDOW_HIDDEN":
		return_data = rl.FlagWindowHidden
	case "FLAG_WINDOW_MINIMIZED":
		return_data = rl.FlagWindowMinimized
	case "FLAG_WINDOW_MAXIMIZED":
		return_data = rl.FlagWindowMaximized
	case "FLAG_WINDOW_UNFOCUSED":
		return_data = rl.FlagWindowUnfocused
	case "FLAG_WINDOW_TOPMOST":
		return_data = rl.FlagWindowTopmost
	case "FLAG_WINDOW_ALWAYS_RUN":
		return_data = rl.FlagWindowAlwaysRun
	case "FLAG_WINDOW_TRANSPARENT":
		return_data = rl.FlagWindowTransparent
	case "FLAG_WINDOW_HIGHDPI":
		return_data = rl.FlagWindowHighdpi
	case "FLAG_WINDOW_MOUSE_PASSTHROUGH":
		return_data = rl.FlagWindowMousePassthrough
	case "FLAG_BORDERLESS_WINDOWED_MODE":
		return_data = rl.FlagBorderlessWindowedMode
	case "FLAG_MSAA_4X_HINT":
		return_data = rl.FlagMsaa4xHint
	case "FLAG_INTERLACED_HINT":
		return_data = rl.FlagInterlacedHint
	}
	return return_data
}

func getMaterialMap(name string) int32 {
	return_data := int32(-1)
	switch strings.ToUpper(name) {
	case "MATERIAL_MAP_ALBEDO":
		return_data = int32(rl.MapAlbedo)
	case "MATERIAL_MAP_METALNESS":
		return_data = int32(rl.MapMetalness)
	case "MATERIAL_MAP_NORMAL":
		return_data = int32(rl.MapNormal)
	case "MATERIAL_MAP_ROUGHNESS":
		return_data = int32(rl.MapRoughness)
	case "MATERIAL_MAP_OCCLUSION":
		return_data = int32(rl.MapOcclusion)
	case "MATERIAL_MAP_EMISSION":
		return_data = int32(rl.MapEmission)
	case "MATERIAL_MAP_HEIGHT":
		return_data = int32(rl.MapHeight)
	case "MATERIAL_MAP_CUBEMAP":
		return_data = int32(rl.MapCubemap)
	case "MATERIAL_MAP_IRRADIANCE":
		return_data = int32(rl.MapIrradiance)
	case "MATERIAL_MAP_PREFILTER":
		return_data = int32(rl.MapPrefilter)
	case "MATERIAL_MAP_BRDF":
		return_data = int32(rl.MapBrdf)

	case "MATERIAL_MAP_DIFFUSE":
		return_data = int32(rl.MapDiffuse)
	case "MATERIAL_MAP_SPECULAR":
		return_data = int32(rl.MapSpecular)
	}
	return return_data
}

func getShapeMode(name string) int32 {
	return_data := int32(-1)
	switch strings.ToUpper(name) {
	case "TRIANGLES":
		return_data = 0
	case "TRIANGLE_FAN":
		return_data = 1
	case "TRIANGLE_STRIP":
		return_data = 2
	case "LINES":
		return_data = 3
	case "LINE_LOOP":
		return_data = 4
	case "LINE_STRIP":
		return_data = 5
	default:
		return_data = -1
	}
	return return_data
}

func getBlendMode(name string) int32 {
	return_data := int32(rl.BlendAlpha)
	switch strings.ToUpper(name) {
	case "BLEND_ALPHA":
		return_data = int32(rl.BlendAlpha)
	case "BLEND_ADDITIVE":
		return_data = int32(rl.BlendAdditive)
	case "BLEND_MULTIPLIED":
		return_data = int32(rl.BlendMultiplied)
	case "BLEND_ADD_COLORS":
		return_data = int32(rl.BlendAddColors)
	case "BLEND_SUBTRACT_COLORS":
		return_data = int32(rl.BlendSubtractColors)
	case "BLEND_ALPHA_PREMULTIPLY":
		return_data = int32(rl.BlendAlphaPremultiply)
	}
	return return_data
}

func getBlendFactor(name string) int32 {
	return_data := int32(rl.Zero)
	switch strings.ToUpper(name) {
	case "ZERO":
		return_data = rl.Zero
	case "ONE":
		return_data = rl.One
	case "SRC_COLOR":
		return_data = rl.SrcColor
	case "ONE_MINUS_SRC_COLOR":
		return_data = rl.OneMinusSrcColor
	case "SRC_ALPHA":
		return_data = rl.SrcAlpha
	case "ONE_MINUS_SRC_ALPHA":
		return_data = rl.OneMinusSrcAlpha
	case "DST_ALPHA":
		return_data = rl.DstAlpha
	case "ONE_MINUS_DST_ALPHA":
		return_data = rl.OneMinusDstAlpha
	case "DST_COLOR":
		return_data = rl.DstColor
	case "ONE_MINUS_DST_COLOR":
		return_data = rl.OneMinusDstColor
	case "SRC_ALPHA_SATURATE":
		return_data = rl.SrcAlphaSaturate
	case "CONSTANT_COLOR":
		return_data = rl.ConstantColor
	case "ONE_MINUS_CONSTANT_COLOR":
		return_data = rl.OneMinusConstantColor
	case "CONSTANT_ALPHA":
		return_data = rl.ConstantAlpha
	case "ONE_MINUS_CONSTANT_ALPHA":
		return_data = rl.OneMinusConstantAlpha
	}
	return return_data
}

func getBlendEquation(name string) int32 {
	return_data := int32(rl.FuncAdd)
	switch strings.ToUpper(name) {
	case "FUNC_ADD":
		return_data = rl.FuncAdd
	case "MIN":
		return_data = rl.Min
	case "MAX":
		return_data = rl.Max
	case "FUNC_SUBTRACT":
		return_data = rl.FuncSubtract
	case "FUNC_REVERSE_SUBTRACT":
		return_data = rl.FuncReverseSubtract
	case "BLEND_EQUATION":
		return_data = rl.BlendEquation
	case "BLEND_EQUATION_RGB":
		return_data = rl.BlendEquationRgb
	case "BLEND_EQUATION_ALPHA":
		return_data = rl.BlendEquationAlpha
	case "BLEND_DST_RGB":
		return_data = rl.BlendDstRgb
	case "BLEND_SRC_RGB":
		return_data = rl.BlendSrcRgb
	case "BLEND_DST_ALPHA":
		return_data = rl.BlendDstAlpha
	case "BLEND_SRC_ALPHA":
		return_data = rl.BlendSrcAlpha
	case "BLEND_COLOR":
		return_data = rl.BlendColor
	}
	return return_data
}

func getKey(name string) int32 {
	return_data := int32(rl.KeyNull)
	switch strings.ToUpper(name) {
	case "KEY_NULL":
		return_data = rl.KeyNull
	case "KEY_APOSTROPHE":
		return_data = rl.KeyApostrophe
	case "KEY_COMMA":
		return_data = rl.KeyComma
	case "KEY_MINUS":
		return_data = rl.KeyMinus
	case "KEY_PERIOD":
		return_data = rl.KeyPeriod
	case "KEY_SLASH":
		return_data = rl.KeySlash
	case "KEY_ZERO":
		return_data = rl.KeyZero
	case "KEY_ONE":
		return_data = rl.KeyOne
	case "KEY_TWO":
		return_data = rl.KeyTwo
	case "KEY_THREE":
		return_data = rl.KeyThree
	case "KEY_FOUR":
		return_data = rl.KeyFour
	case "KEY_FIVE":
		return_data = rl.KeyFive
	case "KEY_SIX":
		return_data = rl.KeySix
	case "KEY_SEVEN":
		return_data = rl.KeySeven
	case "KEY_EIGHT":
		return_data = rl.KeyEight
	case "KEY_NINE":
		return_data = rl.KeyNine
	case "KEY_SEMICOLON":
		return_data = rl.KeySemicolon
	case "KEY_EQUAL":
		return_data = rl.KeyEqual
	case "KEY_A":
		return_data = rl.KeyA
	case "KEY_B":
		return_data = rl.KeyB
	case "KEY_C":
		return_data = rl.KeyC
	case "KEY_D":
		return_data = rl.KeyD
	case "KEY_E":
		return_data = rl.KeyE
	case "KEY_F":
		return_data = rl.KeyF
	case "KEY_G":
		return_data = rl.KeyG
	case "KEY_H":
		return_data = rl.KeyH
	case "KEY_I":
		return_data = rl.KeyI
	case "KEY_J":
		return_data = rl.KeyJ
	case "KEY_K":
		return_data = rl.KeyK
	case "KEY_L":
		return_data = rl.KeyL
	case "KEY_M":
		return_data = rl.KeyM
	case "KEY_N":
		return_data = rl.KeyN
	case "KEY_O":
		return_data = rl.KeyO
	case "KEY_P":
		return_data = rl.KeyP
	case "KEY_Q":
		return_data = rl.KeyQ
	case "KEY_R":
		return_data = rl.KeyR
	case "KEY_S":
		return_data = rl.KeyS
	case "KEY_T":
		return_data = rl.KeyT
	case "KEY_U":
		return_data = rl.KeyU
	case "KEY_V":
		return_data = rl.KeyV
	case "KEY_W":
		return_data = rl.KeyW
	case "KEY_X":
		return_data = rl.KeyX
	case "KEY_Y":
		return_data = rl.KeyY
	case "KEY_Z":
		return_data = rl.KeyZ
	case "KEY_LEFT_BRACKET":
		return_data = rl.KeyLeftBracket
	case "KEY_BACKSLASH":
		return_data = rl.KeyBackSlash
	case "KEY_RIGHT_BRACKET":
		return_data = rl.KeyRightBracket
	case "KEY_GRAVE":
		return_data = rl.KeyGrave
	case "KEY_SPACE":
		return_data = rl.KeySpace
	case "KEY_ESCAPE":
		return_data = rl.KeyEscape
	case "KEY_ENTER":
		return_data = rl.KeyEnter
	case "KEY_TAB":
		return_data = rl.KeyTab
	case "KEY_BACKSPACE":
		return_data = rl.KeyBackspace
	case "KEY_INSERT":
		return_data = rl.KeyInsert
	case "KEY_DELETE":
		return_data = rl.KeyDelete
	case "KEY_RIGHT":
		return_data = rl.KeyRight
	case "KEY_LEFT":
		return_data = rl.KeyLeft
	case "KEY_DOWN":
		return_data = rl.KeyDown
	case "KEY_UP":
		return_data = rl.KeyUp
	case "KEY_PAGE_UP":
		return_data = rl.KeyPageUp
	case "KEY_PAGE_DOWN":
		return_data = rl.KeyPageDown
	case "KEY_HOME":
		return_data = rl.KeyHome
	case "KEY_END":
		return_data = rl.KeyEnd
	case "KEY_CAPS_LOCK":
		return_data = rl.KeyCapsLock
	case "KEY_SCROLL_LOCK":
		return_data = rl.KeyScrollLock
	case "KEY_NUM_LOCK":
		return_data = rl.KeyNumLock
	case "KEY_PRINT_SCREEN":
		return_data = rl.KeyPrintScreen
	case "KEY_PAUSE":
		return_data = rl.KeyPause
	case "KEY_F1":
		return_data = rl.KeyF1
	case "KEY_F2":
		return_data = rl.KeyF2
	case "KEY_F3":
		return_data = rl.KeyF3
	case "KEY_F4":
		return_data = rl.KeyF4
	case "KEY_F5":
		return_data = rl.KeyF5
	case "KEY_F6":
		return_data = rl.KeyF6
	case "KEY_F7":
		return_data = rl.KeyF7
	case "KEY_F8":
		return_data = rl.KeyF8
	case "KEY_F9":
		return_data = rl.KeyF9
	case "KEY_F10":
		return_data = rl.KeyF10
	case "KEY_F11":
		return_data = rl.KeyF11
	case "KEY_F12":
		return_data = rl.KeyF12
	case "KEY_LEFT_SHIFT":
		return_data = rl.KeyLeftShift
	case "KEY_LEFT_CONTROL":
		return_data = rl.KeyLeftControl
	case "KEY_LEFT_ALT":
		return_data = rl.KeyLeftAlt
	case "KEY_LEFT_SUPER":
		return_data = rl.KeyLeftSuper
	case "KEY_RIGHT_SHIFT":
		return_data = rl.KeyRightShift
	case "KEY_RIGHT_CONTROL":
		return_data = rl.KeyRightControl
	case "KEY_RIGHT_ALT":
		return_data = rl.KeyRightAlt
	case "KEY_RIGHT_SUPER":
		return_data = rl.KeyRightSuper
	case "KEY_KB_MENU":
		return_data = rl.KeyKbMenu
	case "KEY_KP_0":
		return_data = rl.KeyKp0
	case "KEY_KP_1":
		return_data = rl.KeyKp1
	case "KEY_KP_2":
		return_data = rl.KeyKp2
	case "KEY_KP_3":
		return_data = rl.KeyKp3
	case "KEY_KP_4":
		return_data = rl.KeyKp4
	case "KEY_KP_5":
		return_data = rl.KeyKp5
	case "KEY_KP_6":
		return_data = rl.KeyKp6
	case "KEY_KP_7":
		return_data = rl.KeyKp7
	case "KEY_KP_8":
		return_data = rl.KeyKp8
	case "KEY_KP_9":
		return_data = rl.KeyKp9
	case "KEY_KP_DECIMAL":
		return_data = rl.KeyKpDecimal
	case "KEY_KP_DIVIDE":
		return_data = rl.KeyKpDivide
	case "KEY_KP_MULTIPLY":
		return_data = rl.KeyKpMultiply
	case "KEY_KP_SUBTRACT":
		return_data = rl.KeyKpSubtract
	case "KEY_KP_ADD":
		return_data = rl.KeyKpAdd
	case "KEY_KP_ENTER":
		return_data = rl.KeyKpEnter
	case "KEY_KP_EQUAL":
		return_data = rl.KeyKpEqual
	}
	return return_data
}

func getMouseButton(name string) rl.MouseButton {
	return_data := rl.MouseButtonLeft
	switch strings.ToUpper(name) {
	case "MOUSE_BUTTON_LEFT":
		return_data = rl.MouseButtonLeft
	case "MOUSE_BUTTON_RIGHT":
		return_data = rl.MouseButtonRight
	case "MOUSE_BUTTON_MIDDLE":
		return_data = rl.MouseButtonMiddle
	case "MOUSE_BUTTON_SIDE":
		return_data = rl.MouseButtonSide
	case "MOUSE_BUTTON_EXTRA":
		return_data = rl.MouseButtonExtra
	case "MOUSE_BUTTON_FORWARD":
		return_data = rl.MouseButtonForward
	case "MOUSE_BUTTON_BACK":
		return_data = rl.MouseButtonBack
	}
	return return_data
}

func getGamepadButton(name string) int32 {
	return_data := int32(rl.GamepadButtonUnknown)
	switch strings.ToUpper(name) {
	case "GAMEPAD_BUTTON_UNKNOWN":
		return_data = rl.GamepadButtonUnknown
	case "GAMEPAD_BUTTON_LEFT_FACE_UP":
		return_data = rl.GamepadButtonLeftFaceUp
	case "GAMEPAD_BUTTON_LEFT_FACE_RIGHT":
		return_data = rl.GamepadButtonLeftFaceRight
	case "GAMEPAD_BUTTON_LEFT_FACE_DOWN":
		return_data = rl.GamepadButtonLeftFaceDown
	case "GAMEPAD_BUTTON_LEFT_FACE_LEFT":
		return_data = rl.GamepadButtonLeftFaceLeft
	case "GAMEPAD_BUTTON_RIGHT_FACE_UP":
		return_data = rl.GamepadButtonRightFaceUp
	case "GAMEPAD_BUTTON_RIGHT_FACE_RIGHT":
		return_data = rl.GamepadButtonRightFaceRight
	case "GAMEPAD_BUTTON_RIGHT_FACE_DOWN":
		return_data = rl.GamepadButtonRightFaceDown
	case "GAMEPAD_BUTTON_RIGHT_FACE_LEFT":
		return_data = rl.GamepadButtonRightFaceLeft
	case "GAMEPAD_BUTTON_LEFT_TRIGGER_1":
		return_data = rl.GamepadButtonLeftTrigger1
	case "GAMEPAD_BUTTON_LEFT_TRIGGER_2":
		return_data = rl.GamepadButtonLeftTrigger2
	case "GAMEPAD_BUTTON_RIGHT_TRIGGER_1":
		return_data = rl.GamepadButtonRightTrigger1
	case "GAMEPAD_BUTTON_RIGHT_TRIGGER_2":
		return_data = rl.GamepadButtonRightTrigger2
	case "GAMEPAD_BUTTON_MIDDLE_LEFT":
		return_data = rl.GamepadButtonMiddleLeft
	case "GAMEPAD_BUTTON_MIDDLE":
		return_data = rl.GamepadButtonMiddle
	case "GAMEPAD_BUTTON_MIDDLE_RIGHT":
		return_data = rl.GamepadButtonMiddleRight
	case "GAMEPAD_BUTTON_LEFT_THUMB":
		return_data = rl.GamepadButtonLeftThumb
	case "GAMEPAD_BUTTON_RIGHT_THUMB":
		return_data = rl.GamepadButtonRightThumb
	}
	return return_data
}

func getTextureFilter(name string) rl.TextureFilterMode {
	return_data := rl.TextureFilterMode(0)
	switch strings.ToUpper(name) {
	case "TEXTURE_FILTER_POINT":
		return_data = rl.TextureFilterMode(0)
	case "TEXTURE_FILTER_BILINEAR":
		return_data = rl.TextureFilterMode(1)
	case "TEXTURE_FILTER_TRILINEAR":
		return_data = rl.TextureFilterMode(2)
	case "TEXTURE_FILTER_ANISOTROPIC_4X":
		return_data = rl.TextureFilterMode(3)
	case "TEXTURE_FILTER_ANISOTROPIC_8X":
		return_data = rl.TextureFilterMode(4)
	case "TEXTURE_FILTER_ANISOTROPIC_16X":
		return_data = rl.TextureFilterMode(5)
	}
	return return_data
}

func memcopy(dest unsafe.Pointer, src unsafe.Pointer, size int) {
	copy(unsafe.Slice((*byte)(dest), size), unsafe.Slice((*byte)(src), size))
}

func halfToFloat(x uint16) float32 {
	result := float32(0.0)
	uni_fm := float32(0.0)
	uni_ui := uint32(0)
	e := uint32((x & 0x7c00) >> 10)
	m := uint32((x & 0x03ff) << 13)
	uni_fm = float32(m)
	memcopy(unsafe.Pointer(&uni_ui), unsafe.Pointer(&uni_fm), 4)
	v := uint32(uni_ui >> 23)
	value := [3]uint32{}
	if e != 0 {
		value[0] = 1
	}
	if e == 0 {
		value[1] = 1
	}
	if m != 0 {
		value[2] = 1
	}
	uni_ui = uint32(x&0x8000)<<16 | value[0]*((e+112)<<23|m) | (value[1]&value[2])*((v-37)<<23|((m<<(150-v))&0x007fe000))
	memcopy(unsafe.Pointer(&uni_fm), unsafe.Pointer(&uni_ui), 4)
	result = uni_fm
	return result
}

//export RlezInit
func RlezInit() {
	system.window_status = false
	system.window_config = 0
	system.shape = -1
	system.shape_auto_normal = true
	system.shape_vertex = [][]float32{}
	system.shape_texture = -1
	system.shape_texture_width = 1
	system.shape_texture_height = 1
	system.shape_texture_flip = false
	system.resource = []Resource{}
	system.return_type = ""
	system.return_data = nil
	system.return_size = 0
	system.draw_mode = "NONE"
	//system.default_material
	system.color = getColor(255, 255, 255, 255)
	system.pixels = nil
	system.pixels_resource = -1
}

//export RlezEnd
func RlezEnd() {
	if system.window_status == true {
		RlezEndDraw()

		RlezDeleteAll()

		rl.UnloadMaterial(system.default_material)
		RlezUnloadPixels()

		rl.CloseWindow()

		rl.CloseAudioDevice()

		RlezInit()
	}
}

//export RlezGetReturnSize
func RlezGetReturnSize() int32 {
	return system.return_size
}

//export RlezGetReturnTypeNameSize
func RlezGetReturnTypeNameSize() int32 {
	return int32(len(system.return_type))
}

//export RlezGetReturnType
func RlezGetReturnType(p uintptr) {
	src := []byte(system.return_type)
	for cnt := int32(0); cnt < RlezGetReturnTypeNameSize(); cnt++ {
		dest := (*byte)(unsafe.Pointer(p + uintptr(cnt)))
		*dest = src[cnt]
	}
}

//export RlezGetReturnData
func RlezGetReturnData(p uintptr) {
	data_bytes := make([]byte, RlezGetReturnTypeNameSize())

	switch system.return_type {
	case "uint8":
		data, _ := system.return_data.(uint8)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "uint16":
		data, _ := system.return_data.(uint16)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "uint32":
		data, _ := system.return_data.(uint32)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "int8":
		data, _ := system.return_data.(int8)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "int16":
		data, _ := system.return_data.(int16)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "int32":
		data, _ := system.return_data.(int32)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "float32":
		data, _ := system.return_data.(float32)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "byte":
		data, _ := system.return_data.(byte)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "rune":
		data, _ := system.return_data.(rune)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "bool":
		data, _ := system.return_data.(bool)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "uint64":
		data, _ := system.return_data.(uint64)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "int64":
		data, _ := system.return_data.(int64)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "float64":
		data, _ := system.return_data.(float64)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, data)
		data_bytes = buffer.Bytes()
	case "uint":
		data, _ := system.return_data.(uint)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, uint32(data))
		data_bytes = buffer.Bytes()
	case "int":
		data, _ := system.return_data.(int)
		buffer := new(bytes.Buffer)
		binary.Write(buffer, binary.LittleEndian, int32(data))
		data_bytes = buffer.Bytes()
	case "uintptr":
		data, _ := system.return_data.(uintptr)
		buffer := new(bytes.Buffer)
		switch unsafe.Sizeof(data) {
		case 4:
			binary.Write(buffer, binary.LittleEndian, uint32(data))
		case 8:
			binary.Write(buffer, binary.LittleEndian, uint64(data))
		}
		data_bytes = buffer.Bytes()
	case "string":
		data, _ := system.return_data.(string)
		data_bytes = []byte(data)
	}
	for cnt := int32(0); cnt < RlezGetReturnSize(); cnt++ {
		dest := (*byte)(unsafe.Pointer(p + uintptr(cnt)))
		*dest = data_bytes[cnt]
	}
}

//export RlezGroup
func RlezGroup(id int32, name uintptr) {
	if 0 <= id && id < int32(len(system.resource)) {
		if system.resource[id].life == true {
			system.resource[id].group = toString(name)
		}
	}
}

//export RlezDelete
func RlezDelete(id int32) {
	if 0 <= id && id < int32(len(system.resource)) {
		if system.resource[id].life == true {
			switch system.resource[id].type_name {
			case "Image":
				rl.UnloadImage((*rl.Image)(system.resource[id].data))
			case "Texture2D":
				rl.UnloadTexture(*(*rl.Texture2D)(system.resource[id].data))
			case "RenderTexture2D":
				rl.UnloadRenderTexture(*(*rl.RenderTexture2D)(system.resource[id].data))
			case "Font":
				rl.UnloadFont(*(*rl.Font)(system.resource[id].data))
			case "Shader":
				rl.UnloadShader(*(*rl.Shader)(system.resource[id].data))
			case "Material":
				rl.UnloadMaterial(*(*rl.Material)(system.resource[id].data))
			case "Mesh":
				rl.UnloadMesh((*rl.Mesh)(system.resource[id].data))
			case "Model":
				rl.UnloadModel(*(*rl.Model)(system.resource[id].data))
			case "ModelAnimation":
				//rl.UnloadModelAnimations
			case "Wave":
				rl.UnloadWave(*(*rl.Wave)(system.resource[id].data))
			case "AudioStream":
				rl.UnloadAudioStream(*(*rl.AudioStream)(system.resource[id].data))
			case "Sound":
				rl.UnloadSound(*(*rl.Sound)(system.resource[id].data))
			case "Music":
				rl.UnloadMusicStream(*(*rl.Music)(system.resource[id].data))
			case "VrStereoConfig":
				rl.UnloadVrStereoConfig(*(*rl.VrStereoConfig)(system.resource[id].data))
			case "AutomationEventList":
				rl.UnloadAutomationEventList((*rl.AutomationEventList)(system.resource[id].data))
			case "RLEZMesh":
				data := (*RLEZMesh)(system.resource[id].data)
				if data.IsFromMesh == true {
					maps := unsafe.Slice(data.Material.Maps, rl.MaxMaterialMaps)
					default_maps := unsafe.Slice(system.default_material.Maps, rl.MaxMaterialMaps)
					for cnt := 0; cnt < rl.MaxMaterialMaps; cnt++ {
						if maps[cnt].Texture != (rl.Texture2D{}) && default_maps[cnt].Texture != (rl.Texture2D{}) {
							maps[cnt].Texture = default_maps[cnt].Texture
						}
					}
					rl.UnloadMaterial(data.Material)
				}
				if data.IsModel == true {
					if data.IsFromMesh == true {
						data.Model.Meshes.Vertices = nil
						data.Model.Meshes.Texcoords = nil
						data.Model.Meshes.Colors = nil
						data.Model.Meshes.Normals = nil
					}
					materials := unsafe.Slice(data.Model.Materials, data.Model.MaterialCount)
					for cnt1 := int32(0); cnt1 < data.Model.MaterialCount; cnt1++ {
						material_maps := unsafe.Slice(materials[cnt1].Maps, rl.MaxMaterialMaps)
						for cnt2 := int32(0); cnt2 < rl.MaxMaterialMaps; cnt2++ {
							material_maps[cnt2].Texture = data.ModelTextures[cnt1][cnt2]
						}
					}
					for cnt := 0; cnt < len(data.ModelAnimations); cnt++ {
						rl.UnloadModelAnimations(data.ModelAnimations)
					}
					rl.UnloadModel(data.Model)
				} else {
					rl.UnloadMesh(&data.Mesh)
				}
			case "RLEZFont":
				data := (*RLEZFont)(system.resource[id].data)
				for cnt := 0; cnt < len(data.Fonts); cnt++ {
					rl.UnloadTexture(data.Fonts[cnt].Texture)
				}
				rl.UnloadFontData(data.GlyphInfos)
			case "RLEZSound":
				data := (*RLEZSound)(system.resource[id].data)
				RlezStopSound(id)
				if data.IsMusic == false {
					rl.UnloadSound(data.SoundData)
				} else {
					data.ChannelMessage <- "delete"
					waiter := time.NewTicker(time.Millisecond)
					for {
						<-waiter.C
						if data.TickerStatus == false {
							break
						}
					}
					waiter.Stop()
					close(data.ChannelSeek)
					close(data.ChannelMessage)
					rl.UnloadMusicStream(data.MusicData)
				}
			}
		}
		clearResource(id)
	}
}

//export RlezDeleteGroup
func RlezDeleteGroup(name uintptr) {
	get_name := toString(name)
	for {
		check := true
		for cnt := int32(0); cnt < int32(len(system.resource)); cnt++ {
			if checkResource(cnt) == true {
				if system.resource[cnt].group == get_name {
					check = false
					RlezDelete(cnt)
					break
				}
			}
		}
		if check == true {
			break
		}
	}
}

//export RlezDeleteAll
func RlezDeleteAll() {
	for cnt := int32(len(system.resource)) - 1; cnt >= 0; cnt-- {
		RlezDelete(cnt)
	}
}

//export RlezLoadTempDir
func RlezLoadTempDir(return_pointer uintptr) {
	dir, err := os.MkdirTemp("", "*")
	if err != nil {
		panic(err)
	}
	return_pointer_slice := unsafe.Slice((*byte)(unsafe.Pointer(return_pointer)), len(dir))
	dir_slice := []byte(dir)
	copy(return_pointer_slice, dir_slice)
}

//export RlezDeleteDir
func RlezDeleteDir(path uintptr) {
	err := os.RemoveAll(toString(path))
	if err != nil {
		panic(err)
	}
}

//export RlezOpenWindow
func RlezOpenWindow(width, height int32) {
	if system.window_status == false {
		system.window_status = true

		runtime.LockOSThread()

		rl.InitAudioDevice()

		set_width := width
		set_height := height
		if set_width <= 0 || set_height <= 0 {
			set_width = 0
			set_height = 0
		}

		rl.SetConfigFlags(system.window_config)

		rl.InitWindow(set_width, set_height, "")

		RlezSetWindowFocus()

		rl.SetTargetFPS(60)
		rl.DisableBackfaceCulling()
		rl.SetExitKey(0)

		system.default_material = rl.LoadMaterialDefault()
	}
}

//export RlezSetWindowFocus
func RlezSetWindowFocus() {
	if system.window_status == true {
		rl.SetWindowFocused()
	}
}

//export RlezCheckWindowFocus
func RlezCheckWindowFocus() int32 {
	return_data := int32(0)
	if system.window_status == true {
		if rl.IsWindowFocused() == true {
			return_data = int32(1)
		}
	}
	return return_data
}

//export RlezSetWindowSize
func RlezSetWindowSize(width, height int32) {
	if system.window_status == true {
		if width >= 0 && height >= 0 {
			rl.SetWindowSize(int(width), int(height))
		}
	}
}

//export RlezSetWindowPosition
func RlezSetWindowPosition(x, y int32) {
	if system.window_status == true {
		rl.SetWindowPosition(int(x), int(y))
	}
}

//export RlezGetWindowX
func RlezGetWindowX() int32 {
	return_data := int32(0)
	if system.window_status == true {
		return_data = int32(rl.GetWindowPosition().X)
	}
	return return_data
}

//export RlezGetWindowY
func RlezGetWindowY() int32 {
	return_data := int32(0)
	if system.window_status == true {
		return_data = int32(rl.GetWindowPosition().Y)
	}
	return return_data
}

//export RlezGetWindowHandle
func RlezGetWindowHandle() uintptr {
	return_data := uintptr(0)
	if system.window_status == true {
		return_data = uintptr(rl.GetWindowHandle())
	}
	setReturnData(return_data, reflect.TypeFor[uintptr]())
	return return_data
}

//export RlezSetWindowIcon
func RlezSetWindowIcon(texture int32) {
	if system.window_status == true {
		if checkResource(texture) == true {
			if system.resource[texture].type_name == "RenderTexture2D" || system.resource[texture].type_name == "Texture2D" {
				image := rl.LoadImageFromTexture(*getTexture(texture))
				if system.resource[texture].type_name == "RenderTexture2D" {
					rl.ImageFlipVertical(image)
				}
				rl.SetWindowIcon(*image)
				rl.UnloadImage(image)
			} else {
				rl.SetWindowIcon(rl.Image{})
			}
		} else {
			rl.SetWindowIcon(rl.Image{})
		}
	}
}

//export RlezGetWindowWidth
func RlezGetWindowWidth() int32 {
	return int32(rl.GetScreenWidth())
}

//export RlezGetWindowHeight
func RlezGetWindowHeight() int32 {
	return int32(rl.GetScreenHeight())
}

//export RlezGetWidth
func RlezGetWidth(resource int32) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if checkResource(resource) == true {
			if system.resource[resource].type_name == "RenderTexture2D" || system.resource[resource].type_name == "Texture2D" {
				return_data = getTexture(resource).Width
			}
		} else {
			return_data = int32(rl.GetRenderWidth())
		}
	}
	return return_data
}

//export RlezGetHeight
func RlezGetHeight(resource int32) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if checkResource(resource) == true {
			if system.resource[resource].type_name == "RenderTexture2D" || system.resource[resource].type_name == "Texture2D" {
				return_data = getTexture(resource).Height
			}
		} else {
			return_data = int32(rl.GetRenderHeight())
		}
	}
	return return_data
}

//export RlezCheckWindowState
func RlezCheckWindowState(flag uintptr) int32 {
	return_data := int32(0)
	if system.window_status == true {
		get_state := getWindowFlag(toString(flag))
		if rl.IsWindowState(get_state) == true {
			return_data = 1
		}
	}
	return return_data
}

//export RlezSetWindowState
func RlezSetWindowState(flag uintptr, value int32) {
	get_state := getWindowFlag(toString(flag))
	if system.window_status == true {
		if value == 0 {
			rl.ClearWindowState(get_state)
		} else {
			rl.SetWindowState(get_state)
		}
	} else {
		if value == 0 {
			if system.window_config&get_state != 0 {
				system.window_config &^= get_state
			}
		} else {
			if system.window_config&get_state == 0 {
				system.window_config |= get_state
			}
		}
	}
}

//export RlezSetWindowTitle
func RlezSetWindowTitle(title uintptr) {
	if system.window_status == true {
		rl.SetWindowTitle(toString(title))
	}
}

//export RlezCheckWindowClose
func RlezCheckWindowClose() int32 {
	return_data := int32(0)
	if system.window_status == true {
		if rl.WindowShouldClose() == true {
			return_data = 1
		}
	}
	return return_data
}

//export RlezGetFrameRate
func RlezGetFrameRate() int32 {
	return_data := int32(0)
	if system.window_status == true {
		return_data = rl.GetFPS()
	}
	return return_data
}

//export RlezSetFrameRate
func RlezSetFrameRate(fps int32) {
	if system.window_status == true {
		rl.SetTargetFPS(fps)
	}
}

//export RlezGetDisplayCount
func RlezGetDisplayCount() int32 {
	return_data := int32(0)
	if system.window_status == true {
		return_data = int32(rl.GetMonitorCount())
	}
	return return_data
}

//export RlezGetCurrentDisplay
func RlezGetCurrentDisplay() int32 {
	return_data := int32(-1)
	if system.window_status == true {
		return_data = int32(rl.GetCurrentMonitor())
	}
	return return_data
}

//export RlezGetDisplayX
func RlezGetDisplayX(display int32) int32 {
	return_data := int32(0)
	if 0 <= display && display < RlezGetDisplayCount() {
		return_data = int32(rl.GetMonitorPosition(int(display)).X)
	}
	return return_data
}

//export RlezGetDisplayY
func RlezGetDisplayY(display int32) int32 {
	return_data := int32(0)
	if 0 <= display && display < RlezGetDisplayCount() {
		return_data = int32(rl.GetMonitorPosition(int(display)).Y)
	}
	return return_data
}

//export RlezGetDisplayWidth
func RlezGetDisplayWidth(display int32) int32 {
	return_data := int32(0)
	if 0 <= display && display < RlezGetDisplayCount() {
		return_data = int32(rl.GetMonitorWidth(int(display)))
	}
	return return_data
}

//export RlezGetDisplayHeight
func RlezGetDisplayHeight(display int32) int32 {
	return_data := int32(0)
	if 0 <= display && display < RlezGetDisplayCount() {
		return_data = int32(rl.GetMonitorHeight(int(display)))
	}
	return return_data
}

//export RlezGetBackend
func RlezGetBackend(return_pointer uintptr) {
	return_pointer_slice := unsafe.Slice((*byte)(unsafe.Pointer(return_pointer)), len(GetBackend))
	backend_slice := []byte(GetBackend)
	copy(return_pointer_slice, backend_slice)
}

//export RlezGetTime
func RlezGetTime() float64 {
	return_data := 0.0
	if system.window_status == true {
		return_data = rl.GetTime()
	}
	setReturnData(return_data, reflect.TypeFor[float64]())
	return return_data
}

//export RlezBeginDraw
func RlezBeginDraw(resource int32) {
	if system.window_status == true {
		if system.draw_mode == "NONE" {
			if checkResource(resource) == true {
				if system.resource[resource].type_name == "RenderTexture2D" {
					system.draw_mode = "RENDER_TEXTURE"
					get_render_texture := (*rl.RenderTexture2D)(system.resource[resource].data)
					rl.BeginTextureMode(*get_render_texture)
				} else {
					system.draw_mode = "WINDOW"
					rl.BeginDrawing()
				}
			} else {
				system.draw_mode = "WINDOW"
				rl.BeginDrawing()
			}
		}
	}
}

//export RlezBeginWindow
func RlezBeginWindow() {
	if system.window_status == true {
		RlezBeginDraw(-1)
	}
}

//export RlezEndDraw
func RlezEndDraw() {
	if system.window_status == true {
		if system.draw_mode != "NONE" {
			switch system.draw_mode {
			case "RENDER_TEXTURE":
				rl.EndTextureMode()
			case "WINDOW":
				rl.EndDrawing()
			}

			system.draw_mode = "NONE"
		}
	}
}

//export RlezBegin2D
func RlezBegin2D(offset_x, offset_y float64, target_x, target_y float64, rotation float64, zoom float64) {
	if system.window_status == true {
		camera := rl.Camera2D{}
		camera.Offset = rl.Vector2{X: float32(offset_x), Y: float32(offset_y)}
		camera.Target = rl.Vector2{X: float32(target_x), Y: float32(target_y)}
		camera.Rotation = float32(rotation)
		camera.Zoom = float32(zoom)
		rl.BeginMode2D(camera)
	}
}

//export RlezEnd2D
func RlezEnd2D() {
	if system.window_status == true {
		rl.EndMode2D()
	}
}

//export RlezBegin3D
func RlezBegin3D(position_x, position_y, position_z float64, target_x, target_y, target_z float64, up_x, up_y, up_z float64, fovy float64, projection int32) {
	if system.window_status == true {
		camera := rl.Camera3D{}
		camera.Position = rl.Vector3{X: float32(position_x), Y: float32(position_y), Z: float32(position_z)}
		camera.Target = rl.Vector3{X: float32(target_x), Y: float32(target_y), Z: float32(target_z)}
		camera.Up = rl.Vector3{X: float32(up_x), Y: float32(up_y), Z: float32(up_z)}
		camera.Fovy = float32(fovy)
		if projection == 0 {
			camera.Projection = rl.CameraPerspective
		} else {
			camera.Projection = rl.CameraOrthographic
		}
		rl.BeginMode3D(camera)
	}
}

//export RlezEnd3D
func RlezEnd3D() {
	if system.window_status == true {
		rl.EndMode3D()
	}
}

//export RlezBeginBlend
func RlezBeginBlend(mode uintptr) {
	if system.window_status == true {
		rl.BeginBlendMode(rl.BlendMode(getBlendMode(toString(mode))))
	}
}

//export RlezBeginBlendCustom
func RlezBeginBlendCustom(src_factor, dest_factor, equation uintptr) {
	if system.window_status == true {
		rl.SetBlendFactors(getBlendFactor(toString(src_factor)), getBlendFactor(toString(dest_factor)), getBlendEquation(toString(equation)))
		rl.BeginBlendMode(rl.BlendCustom)
	}
}

//export RlezBeginBlendCustomSeparate
func RlezBeginBlendCustomSeparate(src_rgb, dest_rgb, src_alpha, dest_alpha, eq_rgb, eq_alpha uintptr) {
	if system.window_status == true {
		rl.SetBlendFactorsSeparate(getBlendFactor(toString(src_rgb)), getBlendFactor(toString(dest_rgb)), getBlendFactor(toString(src_alpha)), getBlendFactor(toString(dest_alpha)), getBlendEquation(toString(eq_rgb)), getBlendEquation(toString(eq_alpha)))
		rl.BeginBlendMode(rl.BlendCustomSeparate)
	}
}

//export RlezEndBlend
func RlezEndBlend() {
	if system.window_status == true {
		rl.EndBlendMode()
	}
}

//export RlezBeginShader
func RlezBeginShader(resource int32) {
	if system.window_status == true {
		if checkResource(resource) == true {
			if system.resource[resource].type_name == "Shader" {
				rl.BeginShaderMode(*(*rl.Shader)(system.resource[resource].data))
			}
		}
	}
}

//export RlezEndShader
func RlezEndShader() {
	if system.window_status == true {
		rl.EndShaderMode()
	}
}

//export RlezPerspective
func RlezPerspective(fovy, aspect, near, far float64) {
	if system.window_status == true {
		top := near * math.Tan(fovy*0.5*rl.Deg2rad)
		right := top * aspect
		RlezFrustum(-right, right, -top, top, near, far)
	}
}

//export RlezFrustum
func RlezFrustum(left, right, bottom, top, near, far float64) {
	if system.window_status == true {
		rl.MatrixMode(rl.Projection)
		rl.LoadIdentity()
		rl.Frustum(left, right, bottom, top, near, far)
		rl.MatrixMode(rl.Modelview)
	}
}

//export RlezOrtho
func RlezOrtho(left, right, bottom, top, near, far float64) {
	if system.window_status == true {
		rl.MatrixMode(rl.Projection)
		rl.LoadIdentity()
		rl.Ortho(left, right, bottom, top, near, far)
		rl.MatrixMode(rl.Modelview)
	}
}

//export RlezPushMatrix
func RlezPushMatrix() {
	if system.window_status == true {
		rl.PushMatrix()
	}
}

//export RlezPopMatrix
func RlezPopMatrix() {
	if system.window_status == true {
		rl.PopMatrix()
	}
}

//export RlezTranslate
func RlezTranslate(x, y, z float64) {
	if system.window_status == true {
		rl.Translatef(float32(x), float32(y), float32(z))
	}
}

//export RlezScale
func RlezScale(x, y, z float64) {
	if system.window_status == true {
		rl.Scalef(float32(x), float32(y), float32(z))
	}
}

//export RlezRotateAxis
func RlezRotateAxis(angle float64, x, y, z float64) {
	if system.window_status == true {
		rl.Rotatef(float32(angle), float32(x), float32(y), float32(z))
	}
}

//export RlezRotateX
func RlezRotateX(angle float64) {
	if system.window_status == true {
		RlezRotateAxis(angle, 1, 0, 0)
	}
}

//export RlezRotateY
func RlezRotateY(angle float64) {
	if system.window_status == true {
		RlezRotateAxis(angle, 0, 1, 0)
	}
}

//export RlezRotateZ
func RlezRotateZ(angle float64) {
	if system.window_status == true {
		RlezRotateAxis(angle, 0, 0, 1)
	}
}

//export RlezRotate
func RlezRotate(angle float64) {
	if system.window_status == true {
		RlezRotateZ(angle)
	}
}

//export RlezLocalToWorld
func RlezLocalToWorld(sx, sy, sz float64, dx, dy, dz uintptr) {
	if system.window_status == true {
		source_vec := rl.Vector3{X: float32(sx), Y: float32(sy), Z: float32(sz)}
		current_matrix := rl.GetMatrixTransform()
		get_vec := rl.Vector3Transform(source_vec, current_matrix)
		ptr_dx := (*float64)(unsafe.Pointer(dx))
		ptr_dy := (*float64)(unsafe.Pointer(dy))
		ptr_dz := (*float64)(unsafe.Pointer(dz))
		*ptr_dx = float64(get_vec.X)
		*ptr_dy = float64(get_vec.Y)
		*ptr_dz = float64(get_vec.Z)
	}
}

//export RlezWorldToLocal
func RlezWorldToLocal(sx, sy, sz float64, dx, dy, dz uintptr) {
	if system.window_status == true {
		source_vec := rl.Vector3{X: float32(sx), Y: float32(sy), Z: float32(sz)}
		current_matrix := rl.MatrixInvert(rl.GetMatrixTransform())
		get_vec := rl.Vector3Transform(source_vec, current_matrix)
		ptr_dx := (*float64)(unsafe.Pointer(dx))
		ptr_dy := (*float64)(unsafe.Pointer(dy))
		ptr_dz := (*float64)(unsafe.Pointer(dz))
		*ptr_dx = float64(get_vec.X)
		*ptr_dy = float64(get_vec.Y)
		*ptr_dz = float64(get_vec.Z)
	}
}

//export RlezWorldToScreen
func RlezWorldToScreen(screen_x, screen_y, screen_w, screen_h float64, sx, sy, sz float64, dx, dy, dz, dw uintptr) {
	if system.window_status == true {
		source_vec := rl.Vector4{X: float32(sx), Y: float32(sy), Z: float32(sz), W: float32(1)}
		current_matrix := rl.MatrixMultiply(rl.GetMatrixProjection(), rl.GetMatrixModelview())
		get_vec := rl.Vector4{}
		get_vec.X = current_matrix.M0*source_vec.X + current_matrix.M4*source_vec.Y + current_matrix.M8*source_vec.Z + current_matrix.M12*source_vec.W
		get_vec.Y = current_matrix.M1*source_vec.X + current_matrix.M5*source_vec.Y + current_matrix.M9*source_vec.Z + current_matrix.M13*source_vec.W
		get_vec.Z = current_matrix.M2*source_vec.X + current_matrix.M6*source_vec.Y + current_matrix.M10*source_vec.Z + current_matrix.M14*source_vec.W
		get_vec.W = current_matrix.M3*source_vec.X + current_matrix.M7*source_vec.Y + current_matrix.M11*source_vec.Z + current_matrix.M15*source_vec.W
		get_vec.X /= get_vec.W
		get_vec.Y /= get_vec.W
		get_vec.Z /= get_vec.W
		get_vec.X = (get_vec.X+1.0)/2.0*float32(screen_w) + float32(screen_x)
		get_vec.Y = (get_vec.Y+1.0)/2.0*float32(screen_h) + float32(screen_y)
		ptr_dx := (*float64)(unsafe.Pointer(dx))
		ptr_dy := (*float64)(unsafe.Pointer(dy))
		ptr_dz := (*float64)(unsafe.Pointer(dz))
		ptr_dw := (*float64)(unsafe.Pointer(dw))
		*ptr_dx = float64(get_vec.X)
		*ptr_dy = float64(get_vec.Y)
		*ptr_dz = float64(get_vec.Z)
		*ptr_dw = float64(get_vec.W)
	}
}

//export RlezLoadRenderTexture
func RlezLoadRenderTexture(width, height int32) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		if width >= 1 && height >= 1 {
			return_data = registerResource(rl.LoadRenderTexture(width, height), reflect.TypeFor[rl.RenderTexture2D]())
		}
	}
	return return_data
}

//export RlezLoadTexture
func RlezLoadTexture(path uintptr) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		image := rl.LoadImage(toString(path))
		rl.ImageFormat(image, rl.UncompressedR8g8b8a8)
		rl.ImageFlipVertical(image)
		return_data = RlezLoadRenderTexture((*image).Width, (*image).Height)
		pixels := unsafe.Slice((*color.RGBA)((*image).Data), ((*image).Width)*((*image).Height))
		rl.UpdateTexture((*rl.RenderTexture2D)(system.resource[return_data].data).Texture, pixels)
		rl.UnloadImage(image)
	}
	return return_data
}

//export RlezLoadTextureFromMemory
func RlezLoadTextureFromMemory(file_type uintptr, data uintptr, size int32) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		bytes := unsafe.Slice((*byte)(unsafe.Pointer(data)), size)
		image := rl.LoadImageFromMemory(toString(file_type), bytes, size)
		rl.ImageFormat(image, rl.UncompressedR8g8b8a8)
		rl.ImageFlipVertical(image)
		return_data = RlezLoadRenderTexture((*image).Width, (*image).Height)
		pixels := unsafe.Slice((*color.RGBA)((*image).Data), ((*image).Width)*((*image).Height))
		rl.UpdateTexture((*rl.RenderTexture2D)(system.resource[return_data].data).Texture, pixels)
		rl.UnloadImage(image)
	}
	return return_data
}

//export RlezSetTextureFilter
func RlezSetTextureFilter(texture int32, filter_type uintptr) {
	if system.window_status == true {
		if system.resource[texture].type_name == "RenderTexture2D" || system.resource[texture].type_name == "Texture2D" {
			rl.SetTextureFilter(*getTexture(texture), getTextureFilter(toString(filter_type)))
		}
	}
}

//export RlezSetTextureMipmaps
func RlezSetTextureMipmaps(texture int32) {
	if system.window_status == true {
		if system.resource[texture].type_name == "RenderTexture2D" || system.resource[texture].type_name == "Texture2D" {
			rl.GenTextureMipmaps(getTexture(texture))
		}
	}
}

//export RlezLoadFont
func RlezLoadFont(path uintptr, font_size int32, target_string uintptr, target_image_width, target_image_height int32) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		file, err := os.ReadFile(toString(path))
		if err == nil {
			return_data = RlezLoadFontFromMemory((uintptr)(unsafe.Pointer(unsafe.SliceData(file))), int32(len(file)), font_size, target_string, target_image_width, target_image_height)
		}
	}
	return return_data
}

//export RlezLoadFontFromMemory
func RlezLoadFontFromMemory(font_data uintptr, font_data_size int32, font_size int32, target_string uintptr, target_image_width, target_image_height int32) int32 {
	return_data := int32(-1)
	get_string := toString(target_string)
	//get_string = strings.ReplaceAll(get_string, " ", "")
	if len(get_string) == 0 {
		get_string = " ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890\"!`?'.,;:()[]{}<>|/@\\^$-%+=#_&~*"
	}
	get_string = strings.ReplaceAll(get_string, " ", "")
	if system.window_status == true && font_data_size > 0 && font_size > 0 && len(get_string) > 0 && target_image_width > 0 && target_image_height > 0 {
		return_font := RLEZFont{}
		get_codepoints := []rune(get_string)
		for cnt1 := 0; cnt1 < len(get_codepoints); cnt1++ {
			flag := true
			get_codepoint := get_codepoints[cnt1]
			for cnt2 := 0; cnt2 < len(return_font.Codepoints); cnt2++ {
				if get_codepoint == return_font.Codepoints[cnt2] {
					flag = false
					break
				}
			}
			if flag == true {
				return_font.Codepoints = append(return_font.Codepoints, get_codepoint)
			}
		}
		return_font.FontSize = font_size
		return_font.GlyphInfos = rl.LoadFontData(unsafe.Slice((*byte)(unsafe.Pointer(font_data)), font_data_size), font_size, return_font.Codepoints, int32(len(return_font.Codepoints)), rl.FontDefault)
		get_max_size := [2]int32{0, 0}
		for cnt := 0; cnt < len(return_font.GlyphInfos); cnt++ {
			if return_font.GlyphInfos[cnt].Image.Width > get_max_size[0] {
				get_max_size[0] = return_font.GlyphInfos[cnt].Image.Width
			}
			if return_font.GlyphInfos[cnt].Image.Height > get_max_size[1] {
				get_max_size[1] = return_font.GlyphInfos[cnt].Image.Height
			}
		}
		get_index := 0
		get_length := len(return_font.Codepoints)
		count := 0
		flag := true
		if len(return_font.GlyphInfos) == 0 {
			flag = false
		}
		for {
			if flag == false || get_index >= get_length {
				break
			}
			add_count := int(float32(target_image_width)/float32(get_max_size[0])*(float32(target_image_height)/float32(get_max_size[1]))) - 1
			if add_count < 0 {
				add_count = 0
			}
			get_last := get_index + add_count
			if get_last >= get_length {
				get_last = get_length - 1
			}
			for {
				get_font := rl.Font{}
				get_glyphinfos_part := return_font.GlyphInfos[get_index : get_last+1]
				get_glyphinfos_part_count := (get_last + 1) - get_index
				get_image := rl.GenImageFontAtlas(get_glyphinfos_part, unsafe.Slice(&get_font.Recs, get_glyphinfos_part_count), return_font.FontSize, 4, 1)
				if get_image.Width <= 0 || get_image.Height <= 0 {
					flag = false
					break
				}
				if (get_image.Width >= target_image_width && get_image.Height >= target_image_height) || (get_last >= get_length-1) {
					get_font.BaseSize = return_font.FontSize
					get_font.CharsCount = int32(len(get_glyphinfos_part))
					get_font.Chars = &(get_glyphinfos_part[0])
					get_font.Texture = rl.LoadTextureFromImage(&get_image)
					rl.UnloadImage(&get_image)
					return_font.GlyphInfosCount = append(return_font.GlyphInfosCount, int32(get_glyphinfos_part_count))
					return_font.Fonts = append(return_font.Fonts, get_font)

					get_index = get_last + 1
					count++
					break
				}
				rl.UnloadImage(&get_image)
				get_last++
				if get_last >= get_length {
					get_last = get_length - 1
				}
			}
		}
		if flag == true {
			return_data = registerResource(return_font, reflect.TypeFor[RLEZFont]())
		} else {
			for cnt := 0; cnt < count; cnt++ {
				rl.UnloadTexture(return_font.Fonts[cnt].Texture)
			}
			rl.UnloadFontData(return_font.GlyphInfos)
		}
	}
	return return_data
}

//export RlezLoadPixels
func RlezLoadPixels(texture int32, format uintptr) {
	if system.window_status == true {
		RlezUnloadPixels()
		if texture >= 0 {
			if checkResource(texture) == true {
				if system.resource[texture].type_name == "RenderTexture2D" || system.resource[texture].type_name == "Texture2D" {
					system.pixels = rl.LoadImageFromTexture(*getTexture(texture))
					if system.resource[texture].type_name == "RenderTexture2D" {
						rl.ImageFlipVertical(system.pixels.(*rl.Image))
					}
					system.pixels_resource = texture
				} else {
					system.pixels = rl.LoadImageFromScreen()
					system.pixels_resource = -1
				}
			} else {
				system.pixels = rl.LoadImageFromScreen()
				system.pixels_resource = -1
			}
		} else {
			system.pixels = rl.LoadImageFromScreen()
			system.pixels_resource = -1
		}
		get_image := system.pixels.(*rl.Image)
		switch strings.ToUpper(toString(format)) {
		case "RGBA":
			rl.ImageFormat(get_image, rl.UncompressedR8g8b8a8)
		case "RGB":
			rl.ImageFormat(get_image, rl.UncompressedR8g8b8)
		case "BGR":
			rl.ImageFormat(get_image, rl.UncompressedR8g8b8)
			get_pixels := unsafe.Slice((*byte)(get_image.Data), 3*(get_image.Width*get_image.Height))
			for cnt := 0; cnt < int(3*(get_image.Width*get_image.Height)); cnt += 3 {
				get_r := get_pixels[cnt+0]
				get_b := get_pixels[cnt+2]
				get_pixels[cnt+0] = get_b
				get_pixels[cnt+2] = get_r
			}
		}
	}
}

//export RlezCopyPixels
func RlezCopyPixels(src_offset, src_length int32, dest_pointer uintptr, dest_offset int32) {
	if system.window_status == true && system.pixels != nil {
		get_image := system.pixels.(*rl.Image)
		src := unsafe.Slice((*byte)(unsafe.Pointer((uintptr(get_image.Data) + uintptr(src_offset)))), src_length)
		dest := unsafe.Slice((*byte)(unsafe.Pointer(dest_pointer+uintptr(dest_offset))), src_length)
		for cnt := int32(0); cnt < src_length; cnt++ {
			dest[cnt] = src[cnt]
		}
	}
}

//export RlezRestorePixels
func RlezRestorePixels(src_pointer uintptr, src_offset, src_length int32, dest_offset int32) {
	if system.window_status == true && system.pixels != nil {
		get_image := system.pixels.(*rl.Image)
		src := unsafe.Slice((*byte)(unsafe.Pointer(src_pointer+uintptr(src_offset))), src_length)
		dest := unsafe.Slice((*byte)(unsafe.Pointer((uintptr(get_image.Data) + uintptr(dest_offset)))), src_length)
		for cnt := int32(0); cnt < src_length; cnt++ {
			dest[cnt] = src[cnt]
		}
	}
}

//export RlezGetPixel
func RlezGetPixel(x, y int32) {
	if system.window_status == true && system.pixels != nil {
		image := system.pixels.(*rl.Image)
		if x < (*image).Width && y < (*image).Height {
			color := rl.GetImageColor(*image, x, y)
			system.color.R = color.R
			system.color.G = color.G
			system.color.B = color.B
			system.color.A = color.A
		}
	}
	return
}

//export RlezSetPixel
func RlezSetPixel(x, y int32) {
	if system.window_status == true && system.pixels != nil {
		image := system.pixels.(*rl.Image)
		if x < (*image).Width && y < (*image).Height {
			rl.ImageDrawPixel(image, x, y, system.color)
		}
	}
}

//export RlezUpdatePixels
func RlezUpdatePixels() {
	if system.window_status == true && system.pixels != nil && system.pixels_resource >= 0 {
		image := system.pixels.(*rl.Image)
		if system.resource[system.pixels_resource].type_name == "RenderTexture2D" {
			rl.ImageFlipVertical(image)
		}
		pixels := unsafe.Slice((*color.RGBA)((*image).Data), ((*image).Width)*((*image).Height))
		rl.UpdateTexture(*getTexture(system.pixels_resource), pixels)
	}
}

//export RlezSavePixels
func RlezSavePixels(path uintptr) {
	if system.window_status == true && system.pixels != nil {
		image := system.pixels.(*rl.Image)
		rl.ExportImage(*image, toString(path))
	}
}

//export RlezUnloadPixels
func RlezUnloadPixels() {
	if system.window_status == true && system.pixels != nil {
		rl.UnloadImage(system.pixels.(*rl.Image))
		system.pixels = nil
		system.pixels_resource = -1
	}
}

//export RlezBackground
func RlezBackground(r, g, b, a int32) {
	if system.window_status == true {
		rl.ClearBackground(getColor(r, g, b, a))
	}
}

//export RlezBeginShape
func RlezBeginShape(mode uintptr, auto_normal int32, resource int32) {
	if system.window_status == true {
		system.shape = -1
		system.shape_auto_normal = false
		system.shape_vertex = [][]float32{}
		system.shape_texture = -1
		system.shape_texture_width = 1
		system.shape_texture_height = 1
		system.shape_texture_flip = false
		if checkResource(resource) == true {
			if system.resource[resource].type_name == "RenderTexture2D" || system.resource[resource].type_name == "Texture2D" {
				system.shape_texture = resource
				temp_texture := *getTexture(resource)
				system.shape_texture_width = temp_texture.Width
				system.shape_texture_height = temp_texture.Height
				if system.shape_texture_width < 1 {
					system.shape_texture_width = 1
				}
				if system.shape_texture_height < 1 {
					system.shape_texture_height = 1
				}
				if system.resource[resource].type_name == "RenderTexture2D" {
					system.shape_texture_flip = true
				}
			}
		}

		system.shape = getShapeMode(toString(mode))

		if auto_normal > 0 {
			system.shape_auto_normal = true
		}
	}
}

//export RlezVertex
func RlezVertex(x, y, z float64, u, v float64, r, g, b, a int32, nx, ny, nz float64) {
	if system.window_status == true {
		get_v := float32(v / float64(system.shape_texture_height))
		if system.shape_texture_flip == true {
			get_v *= -1.0
		}
		system.shape_vertex = append(system.shape_vertex, []float32{float32(x), float32(y), float32(z), float32(u / float64(system.shape_texture_width)), get_v, float32(r) / 255.0, float32(g) / 255.0, float32(b) / 255.0, float32(a) / 255.0, float32(nx), float32(ny), float32(nz)})
	}
}

func endShape(return_mode int32) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		if system.shape != -1 {
			vertex_count := 1
			loop_start := 0
			loop_goal := 0
			loop_add := 0
			switch system.shape {
			case 0:
				vertex_count = 3
				loop_start = 0
				loop_goal = len(system.shape_vertex) / 3 * 3
				loop_add = 3
			case 1:
				vertex_count = 3
				loop_start = 2
				loop_goal = len(system.shape_vertex)
				loop_add = 1
			case 2:
				vertex_count = 3
				loop_start = 2
				loop_goal = len(system.shape_vertex)
				loop_add = 1
			case 3:
				vertex_count = 2
				loop_start = 0
				loop_goal = len(system.shape_vertex) / 2 * 2
				loop_add = 2
			case 4:
				vertex_count = 2
				loop_start = 1
				loop_goal = len(system.shape_vertex) + 1
				loop_add = 1
			case 5:
				vertex_count = 2
				loop_start = 1
				loop_goal = len(system.shape_vertex)
				loop_add = 1
			}
			if len(system.shape_vertex) >= vertex_count {
				mesh := rl.Mesh{}
				var vertices, texcoords, normals []float32
				var colors []uint8
				var texture *rl.Texture2D
				if return_mode == 0 {
					if system.shape_texture >= 0 {
						texture = getTexture(system.shape_texture)
						rl.SetTexture(texture.ID)
					}
				}
				switch vertex_count {
				case 2:
					if return_mode == 0 {
						rl.Begin(rl.Lines)
					}
				case 3:
					if return_mode == 0 {
						rl.Begin(rl.Triangles)
					} else {
						mesh.TriangleCount = 0
						for cnt := loop_start; cnt < loop_goal; cnt += loop_add {
							mesh.TriangleCount++
						}
						mesh.VertexCount = mesh.TriangleCount * 3
					}
				}
				for cnt := loop_start; cnt < loop_goal; cnt += loop_add {
					index := make([]int, vertex_count)
					switch system.shape {
					case 0:
						index[0] = cnt
						index[1] = cnt + 1
						index[2] = cnt + 2
					case 1:
						index[0] = 0
						index[1] = cnt - 1
						index[2] = cnt
					case 2:
						if (cnt-loop_start)%2 == 0 {
							index[0] = cnt - 2
							index[1] = cnt - 1
							index[2] = cnt
						} else {
							index[0] = cnt - 1
							index[1] = cnt - 2
							index[2] = cnt
						}
					case 3:
						index[0] = cnt
						index[1] = cnt + 1
					case 4:
						index[0] = cnt - 1
						index[1] = cnt % (loop_goal - 1)
					case 5:
						index[0] = cnt - 1
						index[1] = cnt
					}
					vec := [3]rl.Vector3{}
					normal := rl.Vector3{}
					if system.shape_auto_normal == true && vertex_count == 3 {
						for cnt2 := 0; cnt2 < vertex_count; cnt2++ {
							vec[cnt2].X = system.shape_vertex[index[cnt2]][0]
							vec[cnt2].Y = system.shape_vertex[index[cnt2]][1]
							vec[cnt2].Z = system.shape_vertex[index[cnt2]][2]
						}
						normal = rl.Vector3CrossProduct(rl.Vector3Subtract(vec[1], vec[0]), rl.Vector3Subtract(vec[2], vec[0]))
					}
					for cnt2 := 0; cnt2 < vertex_count; cnt2++ {
						if return_mode == 0 {
							rl.Color4f(system.shape_vertex[index[cnt2]][5], system.shape_vertex[index[cnt2]][6], system.shape_vertex[index[cnt2]][7], system.shape_vertex[index[cnt2]][8])
							rl.TexCoord2f(system.shape_vertex[index[cnt2]][3], system.shape_vertex[index[cnt2]][4])
							if system.shape_auto_normal == false && vertex_count != 3 {
								rl.Normal3f(system.shape_vertex[index[cnt2]][9], system.shape_vertex[index[cnt2]][10], system.shape_vertex[index[cnt2]][11])
							} else {
								rl.Normal3f(normal.X, normal.Y, normal.Z)
							}
							rl.Vertex3f(system.shape_vertex[index[cnt2]][0], system.shape_vertex[index[cnt2]][1], system.shape_vertex[index[cnt2]][2])
						} else {
							colors = append(colors, uint8(system.shape_vertex[index[cnt2]][5]*255.0), uint8(system.shape_vertex[index[cnt2]][6]*255.0), uint8(system.shape_vertex[index[cnt2]][7]*255.0), uint8(system.shape_vertex[index[cnt2]][8]*255.0))
							texcoords = append(texcoords, system.shape_vertex[index[cnt2]][3], system.shape_vertex[index[cnt2]][4])
							if system.shape_auto_normal == false && vertex_count != 3 {
								normals = append(normals, system.shape_vertex[index[cnt2]][9], system.shape_vertex[index[cnt2]][10], system.shape_vertex[index[cnt2]][11])
							} else {
								normals = append(normals, normal.X, normal.Y, normal.Z)
							}
							vertices = append(vertices, system.shape_vertex[index[cnt2]][0], system.shape_vertex[index[cnt2]][1], system.shape_vertex[index[cnt2]][2])
						}
					}
				}
				if return_mode == 0 {
					rl.End()

					if system.shape_texture >= 0 {
						rl.SetTexture(0)
					}
				} else {
					mesh.Vertices = unsafe.SliceData(vertices)
					mesh.Texcoords = unsafe.SliceData(texcoords)
					mesh.Colors = unsafe.SliceData(colors)
					mesh.Normals = unsafe.SliceData(normals)
					rl.UploadMesh(&mesh, true)

					return_mesh := RLEZMesh{}
					return_mesh.Mesh = mesh
					return_mesh.Material = rl.LoadMaterialDefault()
					return_mesh.IsModel = false
					return_mesh.IsFromMesh = true

					return_data = registerResource(return_mesh, reflect.TypeFor[RLEZMesh]())

					if system.shape_texture >= 0 {
						RlezSetMeshTexture(return_data, system.shape_texture)
					}
				}
			}
		}

		system.shape = -1
		system.shape_auto_normal = true
		system.shape_vertex = [][]float32{}
		system.shape_texture = -1
	}
	return return_data
}

//export RlezEndShape
func RlezEndShape() {
	if system.window_status == true {
		endShape(int32(0))
	}
}

//export RlezEndMesh
func RlezEndMesh() int32 {
	return_data := int32(-1)
	if system.window_status == true {
		return_data = endShape(int32(1))
	}
	return return_data
}

//export RlezSetMeshTexture
func RlezSetMeshTexture(mesh int32, texture int32) {
	if system.window_status == true {
		if checkResource(mesh) == true {
			if system.resource[mesh].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[mesh].data)
				maps := unsafe.Slice(get_mesh.Material.Maps, rl.MaxMaterialMaps)
				default_maps := unsafe.Slice(system.default_material.Maps, rl.MaxMaterialMaps)
				if texture < 0 {
					maps[rl.MapAlbedo].Texture = default_maps[rl.MapAlbedo].Texture
				} else {
					if system.resource[texture].type_name == "RenderTexture2D" || system.resource[texture].type_name == "Texture2D" {
						maps[rl.MapAlbedo].Texture = *getTexture(texture)
					} else {
						maps[rl.MapAlbedo].Texture = default_maps[rl.MapAlbedo].Texture
					}
				}
			}
		}
	}
}

//export RlezLoadModel
func RlezLoadModel(path uintptr, load_animation int32) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		get_mesh := RLEZMesh{}
		get_mesh.Model = rl.LoadModel(toString(path))
		get_mesh.IsFromMesh = false
		get_mesh.IsModel = true
		get_materials := unsafe.Slice(get_mesh.Model.Materials, get_mesh.Model.MaterialCount)
		get_mesh.ModelTextures = make([][]rl.Texture2D, get_mesh.Model.MaterialCount)
		for cnt1 := int32(0); cnt1 < get_mesh.Model.MaterialCount; cnt1++ {
			get_mesh.ModelTextures[cnt1] = make([]rl.Texture2D, rl.MaxMaterialMaps)
			get_maps := unsafe.Slice(get_materials[cnt1].Maps, rl.MaxMaterialMaps)
			for cnt2 := int32(0); cnt2 < get_mesh.Model.MaterialCount; cnt2++ {
				get_mesh.ModelTextures[cnt1][cnt2] = get_maps[cnt2].Texture
			}
		}
		if load_animation != 0 {
			get_mesh.ModelAnimations = rl.LoadModelAnimations(toString(path))
		}
		return_data = registerResource(get_mesh, reflect.TypeFor[RLEZMesh]())
	}
	return return_data
}

//export RlezFromMeshToModel
func RlezFromMeshToModel(mesh int32) {
	if system.window_status == true {
		if checkResource(mesh) == true {
			if system.resource[mesh].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[mesh].data)
				if get_mesh.IsModel == false {
					get_mesh.IsModel = true
					get_mesh.Model = rl.LoadModelFromMesh(get_mesh.Mesh)
					get_materials := unsafe.Slice(get_mesh.Model.Materials, get_mesh.Model.MaterialCount)
					get_mesh.ModelTextures = make([][]rl.Texture2D, get_mesh.Model.MaterialCount)
					for cnt1 := int32(0); cnt1 < get_mesh.Model.MaterialCount; cnt1++ {
						get_mesh.ModelTextures[cnt1] = make([]rl.Texture2D, rl.MaxMaterialMaps)
						get_maps := unsafe.Slice(get_materials[cnt1].Maps, rl.MaxMaterialMaps)
						for cnt2 := int32(0); cnt2 < get_mesh.Model.MaterialCount; cnt2++ {
							get_mesh.ModelTextures[cnt1][cnt2] = get_maps[cnt2].Texture
						}
					}
				}
			}
		}
	}
}

//export RlezSetModelTexture
func RlezSetModelTexture(model int32, material_index int32, map_name uintptr, texture int32) {
	if system.window_status == true {
		if checkResource(model) == true {
			if system.resource[model].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[model].data)
				if get_mesh.IsModel == true {
					if 0 <= material_index && material_index < get_mesh.Model.MaterialCount {
						materials := unsafe.Slice(get_mesh.Model.Materials, get_mesh.Model.MaterialCount)
						material_maps_index := getMaterialMap(toString(map_name))
						var get_texture rl.Texture2D
						if texture >= 0 {
							if checkResource(texture) == true {
								if system.resource[texture].type_name == "Texture2D" || system.resource[texture].type_name == "RenderTexture2D" {
									get_texture = *getTexture(texture)
								} else {
									get_texture = get_mesh.ModelTextures[material_index][material_maps_index]
								}
							} else {
								get_texture = get_mesh.ModelTextures[material_index][material_maps_index]
							}
						} else {
							get_texture = get_mesh.ModelTextures[material_index][material_maps_index]
						}

						material_maps := unsafe.Slice(materials[material_index].Maps, rl.MaxMaterialMaps)
						if material_maps_index >= 0 {
							material_maps[material_maps_index].Texture = get_texture
						}
					}
				}
			}
		}
	}
}

//export RlezGetAnimationCount
func RlezGetAnimationCount(model int32) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if checkResource(model) == true {
			if system.resource[model].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[model].data)
				if get_mesh.IsModel == true {
					return_data = int32(len(get_mesh.ModelAnimations))
				}
			}
		}
	}
	return return_data
}

//export RlezGetAnimationId
func RlezGetAnimationId(model int32, name uintptr) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		if checkResource(model) == true {
			if system.resource[model].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[model].data)
				if get_mesh.IsModel == true {
					get_name := toString(name)
					for cnt := 0; cnt < len(get_mesh.ModelAnimations); cnt++ {
						if string(get_mesh.ModelAnimations[cnt].Name[:]) == get_name {
							return_data = int32(cnt)
							break
						}
					}
				}
			}
		}
	}
	return return_data
}

//export RlezGetAnimationFrames
func RlezGetAnimationFrames(model int32, id int32) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		if checkResource(model) == true {
			if system.resource[model].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[model].data)
				if get_mesh.IsModel == true {
					if int32(0) <= id && id < int32(len(get_mesh.ModelAnimations)) {
						return_data = get_mesh.ModelAnimations[id].KeyframeCount
					}
				}
			}
		}
	}
	return return_data
}

//export RlezSetModelAnimation
func RlezSetModelAnimation(model int32, id int32, frame float64) {
	if system.window_status == true {
		if checkResource(model) == true {
			if system.resource[model].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[model].data)
				if get_mesh.IsModel == true {
					if int32(0) <= id && id < int32(len(get_mesh.ModelAnimations)) {
						rl.UpdateModelAnimation(get_mesh.Model, get_mesh.ModelAnimations[id], float32(frame))
					}
				}
			}
		}
	}
}

//export RlezSetModelAnimationBlend
func RlezSetModelAnimationBlend(model int32, a_id int32, a_frame float64, b_id int32, b_frame float64, blend float64) {
	if system.window_status == true {
		if checkResource(model) == true {
			if system.resource[model].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[model].data)
				if get_mesh.IsModel == true {
					if int32(0) <= a_id && a_id < int32(len(get_mesh.ModelAnimations)) && int32(0) <= b_id && b_id < int32(len(get_mesh.ModelAnimations)) {
						rl.UpdateModelAnimationEx(get_mesh.Model, get_mesh.ModelAnimations[a_id], float32(a_frame), get_mesh.ModelAnimations[b_id], float32(b_frame), float32(blend))
					}
				}
			}
		}
	}
}

//export RlezColor
func RlezColor(r, g, b, a int32) {
	if system.window_status == true {
		system.color.R = uint8(max(0, min(r, 255)))
		system.color.G = uint8(max(0, min(g, 255)))
		system.color.B = uint8(max(0, min(b, 255)))
		system.color.A = uint8(max(0, min(a, 255)))
	}
}

//export RlezGetColorR
func RlezGetColorR() int32 {
	return int32(system.color.R)
}

//export RlezGetColorG
func RlezGetColorG() int32 {
	return int32(system.color.G)
}

//export RlezGetColorB
func RlezGetColorB() int32 {
	return int32(system.color.B)
}

//export RlezGetColorA
func RlezGetColorA() int32 {
	return int32(system.color.A)
}

//export RlezDrawMesh
func RlezDrawMesh(mesh int32) {
	if system.window_status == true {
		if checkResource(mesh) == true {
			if system.resource[mesh].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[mesh].data)
				if get_mesh.IsModel == false {
					maps := unsafe.Slice(get_mesh.Material.Maps, rl.MaxMaterialMaps)
					maps[rl.MapAlbedo].Color = system.color
					rl.DrawMesh(get_mesh.Mesh, get_mesh.Material, rl.MatrixIdentity())
				}
			}
		}
	}
}

//export RlezDrawModel
func RlezDrawModel(model int32) {
	if system.window_status == true {
		if checkResource(model) == true {
			if system.resource[model].type_name == "RLEZMesh" {
				get_mesh := (*RLEZMesh)(system.resource[model].data)
				if get_mesh.IsModel == true {
					rl.DrawModel(get_mesh.Model, rl.Vector3{X: float32(0), Y: float32(0), Z: float32(0)}, float32(1.0), system.color)
				}
			}
		}
	}
}

//export RlezDrawText
func RlezDrawText(font int32, text uintptr, x, y float64, size, spacing float64) {
	if system.window_status == true {
		check_font := false
		if checkResource(font) == true {
			if system.resource[font].type_name == "RLEZFont" {
				check_font = true
			}
		}
		position := rl.Vector2{}
		position.X = float32(x)
		position.Y = float32(y)
		get_font := rl.Font{}
		if check_font == false {
			get_font = rl.GetFontDefault()
			rl.DrawTextEx(get_font, toString(text), position, float32(size), float32(spacing), system.color)
		} else {
			get_codepoints := []rune(toString(text))
			get_fonts := (*RLEZFont)(system.resource[font].data).Fonts
			get_fonts_chars := [][]rl.GlyphInfo{}
			for cnt_font := 0; cnt_font < len(get_fonts); cnt_font++ {
				get_fonts_chars = append(get_fonts_chars, unsafe.Slice(get_fonts[cnt_font].Chars, get_fonts[cnt_font].CharsCount))
			}
			for cnt := 0; cnt < len(get_codepoints); cnt++ {
				flag := false
				for cnt_font := 0; cnt_font < len(get_fonts); cnt_font++ {
					get_font = get_fonts[cnt_font]
					for cnt_glyph := 0; cnt_glyph < int(get_font.CharsCount); cnt_glyph++ {
						if get_fonts_chars[cnt_font][cnt_glyph].Value == get_codepoints[cnt] {
							flag = true
							break
						}
					}
					if flag == true {
						break
					}
				}
				if flag == true {
					rl.DrawTextCodepoint(get_font, get_codepoints[cnt], position, float32(size), system.color)
					get_char_size := rl.MeasureTextEx(get_font, string(get_codepoints[cnt]), float32(size), float32(0.0))
					position.X += get_char_size.X
				} else {
					if get_codepoints[cnt] == 32 {
						position.X += float32(size / 2.0)
					}
				}
				if cnt < len(get_codepoints)-1 {
					position.X += float32(spacing)
				}
			}
		}
	}
}

//export RlezDrawLine
func RlezDrawLine(x1, y1, z1 float64, x2, y2, z2 float64) {
	if system.window_status == true {
		rl.DrawLine3D(rl.Vector3{X: float32(x1), Y: float32(y1), Z: float32(z1)}, rl.Vector3{X: float32(x2), Y: float32(y2), Z: float32(z2)}, system.color)
	}
}

//export RlezDrawRect
func RlezDrawRect(x, y float64, width, height float64, fill int32) {
	if system.window_status == true {
		if fill == 0 {
			rl.Begin(rl.Lines)
			rl.Color4ub(system.color.R, system.color.G, system.color.B, system.color.A)
			rl.Vertex2f(float32(x), float32(y))
			rl.Vertex2f(float32(x+width), float32(y))
			rl.Vertex2f(float32(x+width), float32(y))
			rl.Vertex2f(float32(x+width), float32(y+height))
			rl.Vertex2f(float32(x+width), float32(y+height))
			rl.Vertex2f(float32(x), float32(y+height))
			rl.Vertex2f(float32(x), float32(y+height))
			rl.Vertex2f(float32(x), float32(y))
			rl.End()
		} else {
			rl.DrawRectangleV(rl.Vector2{X: float32(x), Y: float32(y)}, rl.Vector2{X: float32(width), Y: float32(height)}, system.color)
		}
	}
}

//export RlezDrawEllipse
func RlezDrawEllipse(x, y float64, width, height float64, fill int32) {
	if system.window_status == true {
		RlezPushMatrix()
		RlezTranslate(x, y, 0)
		if fill == 0 {
			rl.DrawEllipseLines(0, 0, float32(width)/2.0, float32(height)/2.0, system.color)
		} else {
			rl.DrawEllipse(0, 0, float32(width)/2.0, float32(height)/2.0, system.color)
		}
		RlezPopMatrix()
	}
}

//export RlezDrawBox
func RlezDrawBox(x, y, z float64, width, height, depth float64, fill int32) {
	if system.window_status == true {
		if fill == 0 {
			rl.DrawCubeWires(rl.Vector3{X: float32(x), Y: float32(y), Z: float32(z)}, float32(width), float32(height), float32(depth), system.color)
		} else {
			rl.DrawCube(rl.Vector3{X: float32(x), Y: float32(y), Z: float32(z)}, float32(width), float32(height), float32(depth), system.color)
		}
	}
}

//export RlezDrawSphere
func RlezDrawSphere(x, y, z float64, size float64, rings, slices int32, fill int32) {
	if system.window_status == true {
		get_rings, get_slices := rings, slices
		if get_rings <= 0 {
			get_rings = 16
		}
		if get_slices <= 0 {
			get_slices = 16
		}
		if fill == 0 {
			rl.DrawSphereWires(rl.Vector3{X: float32(x), Y: float32(y), Z: float32(z)}, float32(size/float64(2.0)), get_rings, get_slices, system.color)
		} else {
			rl.DrawSphereEx(rl.Vector3{X: float32(x), Y: float32(y), Z: float32(z)}, float32(size/float64(2.0)), get_rings, get_slices, system.color)
		}
	}
}

//export RlezDrawCylinder
func RlezDrawCylinder(x1, y1, z1 float64, x2, y2, z2 float64, size1, size2 float64, sides int32, fill int32) {
	if system.window_status == true {
		get_sides := sides
		if get_sides < 3 {
			get_sides = 3
		}
		if fill == 0 {
			rl.DrawCylinderWiresEx(rl.Vector3{X: float32(x1), Y: float32(y1), Z: float32(z1)}, rl.Vector3{X: float32(x2), Y: float32(y2), Z: float32(z2)}, float32(size1/float64(2.0)), float32(size2/float64(2.0)), get_sides, system.color)
		} else {
			rl.DrawCylinderEx(rl.Vector3{X: float32(x1), Y: float32(y1), Z: float32(z1)}, rl.Vector3{X: float32(x2), Y: float32(y2), Z: float32(z2)}, float32(size1/float64(2.0)), float32(size2/float64(2.0)), get_sides, system.color)
		}
	}
}

//export RlezDrawCapsule
func RlezDrawCapsule(x1, y1, z1 float64, x2, y2, z2 float64, size float64, slices, rings int32, fill int32) {
	if system.window_status == true {
		get_slices, get_rings := slices, rings
		if get_slices <= 0 {
			get_slices = 8
		}
		if get_rings <= 0 {
			get_rings = 16
		}
		if fill == 0 {
			rl.DrawCapsuleWires(rl.Vector3{X: float32(x1), Y: float32(y1), Z: float32(z1)}, rl.Vector3{X: float32(x2), Y: float32(y2), Z: float32(z2)}, float32(size/float64(2.0)), get_slices, get_rings, system.color)
		} else {
			rl.DrawCapsule(rl.Vector3{X: float32(x1), Y: float32(y1), Z: float32(z1)}, rl.Vector3{X: float32(x2), Y: float32(y2), Z: float32(z2)}, float32(size/float64(2.0)), get_slices, get_rings, system.color)
		}
	}
}

//export RlezDrawTexture
func RlezDrawTexture(texture int32, src_x, src_y, src_w, src_h int32, dest_x, dest_y, dest_w, dest_h float64) {
	if system.window_status == true {
		if checkResource(texture) == true {
			v_scale := float32(1.0)
			if system.resource[texture].type_name == "RenderTexture2D" {
				v_scale = float32(-1.0)
			}
			rl.DrawTexturePro(*getTexture(texture), rl.Rectangle{X: float32(src_x), Y: v_scale * float32(src_y), Width: float32(src_w), Height: v_scale * float32(src_h)}, rl.Rectangle{X: float32(dest_x), Y: float32(dest_y), Width: float32(dest_w), Height: float32(dest_h)}, rl.Vector2{X: float32(0), Y: float32(0)}, float32(0), system.color)
		}
	}
}

//export RlezLoadShaderFromMemory
func RlezLoadShaderFromMemory(vertex_code uintptr, fragment_code uintptr) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		return_shader := rl.LoadShader(toString(vertex_code), toString(fragment_code))
		return_data = registerResource(return_shader, reflect.TypeFor[rl.Shader]())
	}
	return return_data
}

//export RlezLoadShader
func RlezLoadShader(vertex_code_path uintptr, fragment_code_path uintptr) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		vertex_code, vertex_err := os.ReadFile(toString(vertex_code_path))
		fragment_code, fragment_err := os.ReadFile(toString(fragment_code_path))
		if vertex_err == nil && fragment_err == nil {
			vertex_memory := fromString(string(vertex_code))
			fragment_memory := fromString(string(fragment_code))
			return_data = RlezLoadShaderFromMemory(vertex_memory.string_uintptr, fragment_memory.string_uintptr)
		}
	}
	return return_data
}

//export RlezLoadSoundFromMemory
func RlezLoadSoundFromMemory(file_type uintptr, sound_data uintptr, sound_data_size int32, is_music int32, music_update_time float64, music_update_samples int32) int32 {
	return_data := int32(-1)
	if system.window_status == true && sound_data_size > 0 {
		return_sound := RLEZSound{}
		return_sound.MusicBytes = make([]byte, sound_data_size)
		copy(return_sound.MusicBytes, unsafe.Slice((*byte)(unsafe.Pointer(sound_data)), sound_data_size))
		if is_music == 0 {
			return_sound.IsMusic = false
			get_wave := rl.LoadWaveFromMemory(toString(file_type), return_sound.MusicBytes, sound_data_size)
			return_sound.SoundData = rl.LoadSoundFromWave(get_wave)
			rl.UnloadWave(get_wave)
			return_sound.MusicBytes = nil
		} else {
			if music_update_samples <= 0 {
				rl.SetAudioStreamBufferSizeDefault(4096)
			} else {
				rl.SetAudioStreamBufferSizeDefault(music_update_samples)
			}
			return_sound.IsMusic = true
			return_sound.MusicData = rl.LoadMusicStreamFromMemory(toString(file_type), return_sound.MusicBytes, sound_data_size)
			return_sound.MusicUpdateTime = music_update_time
			return_sound.ChannelSeek = make(chan float32)
			return_sound.ChannelMessage = make(chan string)
			return_sound.Time = float32(0.0)
			return_sound.TickerTime = int64(float64(time.Second.Nanoseconds()) * return_sound.MusicUpdateTime)
			if return_sound.TickerTime <= 0 {
				return_sound.TickerTime = int64(float64(time.Second.Nanoseconds()) * (float64(1) / float64(60)))
			}
			return_sound.TickerStatus = false
			return_sound.TickerFunc = func(sound *RLEZSound) {
				runtime.LockOSThread()
				defer runtime.UnlockOSThread()
				sound.TickerStatus = true
				music := &(sound.MusicData)
				ticker_data := time.NewTicker(time.Duration(sound.TickerTime))
				defer ticker_data.Stop()

				for {
					check_break := false

					sound.Time = rl.GetMusicTimePlayed(*music)

					select {
					case <-ticker_data.C:
						if rl.IsMusicValid(*music) == true {
							if rl.IsMusicStreamPlaying(*music) == true {
								rl.UpdateMusicStream(*music)
							}
						}
					case get_seek := <-sound.ChannelSeek:
						if rl.IsMusicValid(*music) == true {
							rl.SeekMusicStream(*music, float32(get_seek))
						}
					case get_message := <-sound.ChannelMessage:
						if rl.IsMusicValid(*music) == true {
							switch get_message {
							case "update":
								rl.UpdateMusicStream(*music)
							case "play":
								rl.PlayMusicStream(*music)
							case "pause":
								rl.PauseMusicStream(*music)
							case "stop":
								rl.StopMusicStream(*music)
							case "resume":
								rl.ResumeMusicStream(*music)
							case "delete":
								check_break = true
							}
						}
					}

					if check_break == true {
						break
					}
				}
				sound.TickerStatus = false
			}
		}
		return_data = registerResource(return_sound, reflect.TypeFor[RLEZSound]())
		if is_music != 0 {
			get_sound := (*RLEZSound)(system.resource[return_data].data)
			go get_sound.TickerFunc(get_sound)
		}
	}
	return return_data
}

//export RlezLoadSound
func RlezLoadSound(path uintptr, is_music int32, music_update_time float64, music_update_samples int32) int32 {
	return_data := int32(-1)
	if system.window_status == true {
		get_path := toString(path)
		file, err := os.ReadFile(get_path)
		if err == nil {
			get_ext := strings.ToLower(filepath.Ext(get_path))
			ext_memory := fromString(get_ext)
			return_data = RlezLoadSoundFromMemory(ext_memory.string_uintptr, (uintptr)(unsafe.Pointer(unsafe.SliceData(file))), int32(len(file)), is_music, music_update_time, music_update_samples)
		}
	}
	return return_data
}

//export RlezStopSound
func RlezStopSound(sound int32) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					rl.StopSound(get_sound.SoundData)
				} else {
					get_sound.ChannelMessage <- "update"
					get_sound.ChannelMessage <- "stop"
				}
			}
		}
	}
}

//export RlezPauseSound
func RlezPauseSound(sound int32) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					rl.PauseSound(get_sound.SoundData)
				} else {
					get_sound.ChannelMessage <- "update"
					get_sound.ChannelMessage <- "pause"
				}
			}
		}
	}
}

//export RlezResumeSound
func RlezResumeSound(sound int32) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					rl.ResumeSound(get_sound.SoundData)
				} else {
					get_sound.ChannelMessage <- "update"
					get_sound.ChannelMessage <- "resume"
				}
			}
		}
	}
}

//export RlezPlaySound
func RlezPlaySound(sound int32, music_loop int32) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					rl.PlaySound(get_sound.SoundData)
				} else {
					current_time := RlezGetSoundTime(sound)

					get_music := &(get_sound.MusicData)
					if music_loop == int32(0) {
						get_music.Looping = false
					} else {
						get_music.Looping = true
					}
					get_sound.ChannelMessage <- "play"

					RlezSetSoundTime(sound, current_time)

					get_sound.ChannelMessage <- "update"
				}
			}
		}
	}
}

//export RlezGetSoundStatus
func RlezGetSoundStatus(sound int32) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					if rl.IsSoundPlaying(get_sound.SoundData) == true {
						return_data = int32(1)
					}
				} else {
					get_sound.ChannelMessage <- "update"

					get_music := &(get_sound.MusicData)
					if rl.IsMusicStreamPlaying(*get_music) == true {
						return_data = int32(1)
					}
				}
			}
		}
	}
	return return_data
}

//export RlezGetSoundTime
func RlezGetSoundTime(sound int32) float64 {
	return_data := float64(0)
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == true {
					get_sound.ChannelMessage <- "update"
					return_data = float64(get_sound.Time)
				}
			}
		}
	}
	setReturnData(return_data, reflect.TypeFor[float64]())
	return return_data
}

//export RlezSetSoundTime
func RlezSetSoundTime(sound int32, set_time float64) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == true {
					target_time := float32(set_time)
					if target_time < float32(0) || float32(RlezGetSoundLength(sound)) <= target_time {
						target_time = float32(0)
					}

					get_status := RlezGetSoundStatus(sound)
					if get_status == int32(1) {
						RlezPauseSound(sound)
					}

					get_sound.ChannelSeek <- target_time

					get_seconds_per_frame := float32(1.0) / float32(get_sound.MusicData.Stream.SampleRate)

					waiter := time.NewTicker(time.Millisecond)
					for {
						<-waiter.C
						get_time := float32(RlezGetSoundTime(sound))
						if target_time-get_seconds_per_frame <= get_time && get_time <= target_time+get_seconds_per_frame {
							break
						}
					}
					waiter.Stop()

					if get_status == int32(1) {
						RlezResumeSound(sound)
					}
				}
			}
		}
	}
}

//export RlezGetSoundLength
func RlezGetSoundLength(sound int32) float64 {
	return_data := float64(0)
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == true {
					return_data = float64(rl.GetMusicTimeLength(get_sound.MusicData))
				}
			}
		}
	}
	setReturnData(return_data, reflect.TypeFor[float64]())
	return return_data
}

//export RlezSetSoundPitch
func RlezSetSoundPitch(sound int32, pitch float64) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					rl.SetSoundPitch(get_sound.SoundData, float32(pitch))
				} else {
					rl.SetMusicPitch(get_sound.MusicData, float32(pitch))
					get_sound.ChannelMessage <- "update"
				}
			}
		}
	}
}

//export RlezSetSoundVolume
func RlezSetSoundVolume(sound int32, volume float64) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					rl.SetSoundVolume(get_sound.SoundData, float32(volume))
				} else {
					rl.SetMusicVolume(get_sound.MusicData, float32(volume))
					get_sound.ChannelMessage <- "update"
				}
			}
		}
	}
}

//export RlezSetSoundPan
func RlezSetSoundPan(sound int32, pan float64) {
	if system.window_status == true {
		if checkResource(sound) == true {
			if system.resource[sound].type_name == "RLEZSound" {
				get_sound := (*RLEZSound)(system.resource[sound].data)
				if get_sound.IsMusic == false {
					rl.SetSoundPan(get_sound.SoundData, float32((-pan+float64(1))/float64(2)))
				} else {
					rl.SetMusicPan(get_sound.MusicData, float32((-pan+float64(1))/float64(2)))
					get_sound.ChannelMessage <- "update"
				}
			}
		}
	}
}

//export RlezGetKey
func RlezGetKey(name uintptr) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if rl.IsKeyDown(getKey(toString(name))) == true {
			return_data = int32(1)
		}
	}
	return return_data
}

//export RlezGetMouseButton
func RlezGetMouseButton(name uintptr) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if rl.IsMouseButtonDown(getMouseButton(toString(name))) == true {
			return_data = int32(1)
		}
	}
	return return_data
}

//export RlezGetMouseX
func RlezGetMouseX() int32 {
	return_data := int32(0)
	if system.window_status == true {
		return_data = rl.GetMouseX()
	}
	return return_data
}

//export RlezGetMouseY
func RlezGetMouseY() int32 {
	return_data := int32(0)
	if system.window_status == true {
		return_data = rl.GetMouseY()
	}
	return return_data
}

//export RlezSetMousePosition
func RlezSetMousePosition(x, y int32) {
	if system.window_status == true {
		rl.SetMousePosition(int(x), int(y))
	}
}

//export RlezSetMouseVisible
func RlezSetMouseVisible(visible int32) {
	if system.window_status == true {
		if visible == int32(0) {
			rl.HideCursor()
		} else {
			rl.ShowCursor()
		}
	}
}

//export RlezCheckMouseInWindow
func RlezCheckMouseInWindow() int32 {
	return_data := int32(0)
	if system.window_status == true {
		if rl.IsCursorOnScreen() == true {
			return_data = int32(1)
		}
	}
	return return_data
}

//export RlezGetMouseWheelX
func RlezGetMouseWheelX() float64 {
	return_data := float64(0.0)
	if system.window_status == true {
		return_data = float64(rl.GetMouseWheelMoveV().X)
	}
	setReturnData(return_data, reflect.TypeFor[float64]())
	return return_data
}

//export RlezGetMouseWheelY
func RlezGetMouseWheelY() float64 {
	return_data := float64(0.0)
	if system.window_status == true {
		return_data = float64(rl.GetMouseWheelMoveV().Y)
	}
	setReturnData(return_data, reflect.TypeFor[float64]())
	return return_data
}

//export RlezCheckGamepad
func RlezCheckGamepad(gamepad int32) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if rl.IsGamepadAvailable(gamepad) == true {
			return_data = int32(1)
		}
	}
	return return_data
}

//export RlezGetGamepadButton
func RlezGetGamepadButton(gamepad int32, button uintptr) int32 {
	return_data := int32(0)
	if system.window_status == true {
		if rl.IsGamepadButtonDown(gamepad, getGamepadButton(toString(button))) == true {
			return_data = int32(1)
		}
	}
	return return_data
}

//export RlezGetAxisCount
func RlezGetAxisCount(gamepad int32) int32 {
	return_data := int32(0)
	if system.window_status == true {
		return_data = rl.GetGamepadAxisCount(gamepad)
	}
	return return_data
}

//export RlezGetGamepadAxis
func RlezGetGamepadAxis(gamepad int32, axis int32) float64 {
	return_data := float64(0.0)
	if system.window_status == true {
		return_data = float64(rl.GetGamepadAxisMovement(gamepad, axis))
	}
	setReturnData(return_data, reflect.TypeFor[float64]())
	return return_data
}

//export RlezSetGamepadVibration
func RlezSetGamepadVibration(gamepad int32, left float64, right float64, duration float64) {
	if system.window_status == true {
		rl.SetGamepadVibration(gamepad, float32(left), float32(right), float32(duration))
	}
}

func main() {
}
