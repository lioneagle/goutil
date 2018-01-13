package sdl

import (
	"unsafe"
)

/**
 *  \brief General event structure
 */
type SDL_EventType struct {
	Type uint32
}

type SDL_Event struct {
	SDL_EventType
	padding [56]uint8
}

func (this *SDL_Event) SDL_CommonEvent() *SDL_CommonEvent {
	return (*SDL_CommonEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_WindowEvent() *SDL_WindowEvent {
	return (*SDL_WindowEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_KeyboardEvent() *SDL_KeyboardEvent {
	return (*SDL_KeyboardEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_TextEditingEvent() *SDL_TextEditingEvent {
	return (*SDL_TextEditingEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_TextInputEvent() *SDL_TextInputEvent {
	return (*SDL_TextInputEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_MouseMotionEvent() *SDL_MouseMotionEvent {
	return (*SDL_MouseMotionEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_MouseButtonEvent() *SDL_MouseButtonEvent {
	return (*SDL_MouseButtonEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_MouseWheelEvent() *SDL_MouseWheelEvent {
	return (*SDL_MouseWheelEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_JoyAxisEvent() *SDL_JoyAxisEvent {
	return (*SDL_JoyAxisEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_JoyBallEvent() *SDL_JoyBallEvent {
	return (*SDL_JoyBallEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_JoyHatEvent() *SDL_JoyHatEvent {
	return (*SDL_JoyHatEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_JoyButtonEvent() *SDL_JoyButtonEvent {
	return (*SDL_JoyButtonEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_JoyDeviceEvent() *SDL_JoyDeviceEvent {
	return (*SDL_JoyDeviceEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_ControllerAxisEvent() *SDL_ControllerAxisEvent {
	return (*SDL_ControllerAxisEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_ControllerButtonEvent() *SDL_ControllerButtonEvent {
	return (*SDL_ControllerButtonEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_ControllerDeviceEvent() *SDL_ControllerDeviceEvent {
	return (*SDL_ControllerDeviceEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_AudioDeviceEvent() *SDL_AudioDeviceEvent {
	return (*SDL_AudioDeviceEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_TouchFingerEvent() *SDL_TouchFingerEvent {
	return (*SDL_TouchFingerEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_MultiGestureEvent() *SDL_MultiGestureEvent {
	return (*SDL_MultiGestureEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_DollarGestureEvent() *SDL_DollarGestureEvent {
	return (*SDL_DollarGestureEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_DropEvent() *SDL_DropEvent {
	return (*SDL_DropEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_QuitEvent() *SDL_QuitEvent {
	return (*SDL_QuitEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_OSEvent() *SDL_OSEvent {
	return (*SDL_OSEvent)(unsafe.Pointer(this))
}

func (this *SDL_Event) SDL_UserEvent() *SDL_UserEvent {
	return (*SDL_UserEvent)(unsafe.Pointer(this))
}

/**
 *  \brief Fields shared by every event
 */
type SDL_CommonEvent struct {
	SDL_EventType
	Timestamp uint32
}

/**
 *  \brief Window state change event data (event.window.*)
 */
type SDL_WindowEvent struct {
	SDL_EventType          /**< ::SDL_WINDOWEVENT */
	Timestamp     uint32   /*  */
	WindowID      uint32   /**< The associated window */
	Event         uint8    /**< ::SDL_WindowEventID */
	padding       [3]uint8 /* */
	Data1         int32    /**< event dependent data */
	Data2         int32    /**< event dependent data */
}

/**
 *  \brief Keyboard button event structure (event.key.*)
 */
type SDL_KeyboardEvent struct {
	SDL_EventType            /**< ::SDL_KEYDOWN or ::SDL_KEYUP */
	Timestamp     uint32     /*  */
	WindowID      uint32     /**< The window with keyboard focus, if any */
	State         uint8      /**< ::SDL_PRESSED or ::SDL_RELEASED */
	Repeat        uint8      /**< Non-zero if this is a key repeat */
	padding       [2]uint8   /*  */
	Keysym        SDL_Keysym /**< The key that was pressed or released */
}

const SDL_TEXTEDITINGEVENT_TEXT_SIZE = 32

/**
 *  \brief Keyboard text editing event structure (event.edit.*)
 */
type SDL_TextEditingEvent struct {
	SDL_EventType                                      /**< ::SDL_TEXTEDITING */
	Timestamp     uint32                               /*  */
	WindowID      uint32                               /**< The window with keyboard focus, if any */
	Text          [SDL_TEXTEDITINGEVENT_TEXT_SIZE]byte /**< The editing text */
	Start         int32                                /**< The start cursor of selected editing text */
	Length        int32                                /**< The length of selected editing text */
}

const SDL_TEXTINPUTEVENT_TEXT_SIZE = 32

/**
 *  \brief Keyboard text input event structure (event.text.*)
 */
type SDL_TextInputEvent struct {
	SDL_EventType                                    /**< ::SDL_TEXTINPUT */
	Timestamp     uint32                             /*  */
	WindowID      uint32                             /**< The window with keyboard focus, if any */
	Text          [SDL_TEXTINPUTEVENT_TEXT_SIZE]byte /**< The input text */
}

/**
 *  \brief Mouse motion event structure (event.motion.*)
 */
type SDL_MouseMotionEvent struct {
	SDL_EventType        /**< ::SDL_MOUSEMOTION */
	Timestamp     uint32 /*  */
	WindowID      uint32 /**< The window with mouse focus, if any */
	Which         uint32 /**< The mouse instance id, or SDL_TOUCH_MOUSEID */
	State         uint32 /**< The current button state */
	X             int32  /**< X coordinate, relative to window */
	Y             int32  /**< Y coordinate, relative to window */
	Xrel          int32  /**< The relative motion in the X direction */
	Yrel          int32  /**< The relative motion in the Y direction */
}

/**
 *  \brief Mouse button event structure (event.button.*)
 */
type SDL_MouseButtonEvent struct {
	SDL_EventType        /**< ::SDL_MOUSEBUTTONDOWN or ::SDL_MOUSEBUTTONUP */
	Timestamp     uint32 /*  */
	WindowID      uint32 /**< The window with mouse focus, if any */
	Which         uint32 /**< The mouse instance id, or SDL_TOUCH_MOUSEID */
	Button        uint8  /**< The mouse button index */
	State         uint8  /**< ::SDL_PRESSED or ::SDL_RELEASED */
	Clicks        uint8  /**< X coordinate, relative to window */
	padding       uint8  /*  */
	X             int32  /**< X coordinate, relative to window */
	Y             int32  /**< Y coordinate, relative to window */
}

/**
 *  \brief Mouse wheel event structure (event.wheel.*)
 */
type SDL_MouseWheelEvent struct {
	SDL_EventType        /**< ::SDL_MOUSEWHEEL */
	Timestamp     uint32 /*  */
	WindowID      uint32 /**< The window with mouse focus, if any */
	Which         uint32 /**< The mouse instance id, or SDL_TOUCH_MOUSEID */
	X             int32  /**< The amount scrolled horizontally, positive to the right and negative to the left */
	Y             int32  /**< The amount scrolled vertically, positive away from the user and negative toward the user */
	Direction     uint32 /**< Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in X and Y will be opposite. Multiply by -1 to change them back */
}

/**
 *  \brief Joystick axis motion event structure (event.jaxis.*)
 */
type SDL_JoyAxisEvent struct {
	SDL_EventType                /**< ::SDL_JOYAXISMOTION */
	Timestamp     uint32         /*  */
	Which         SDL_JoystickID /**< The joystick instance id */
	Axis          uint8          /**< The joystick axis index */
	padding       [3]uint8       /*  */
	Value         int16          /**< The axis value (range: -32768 to 32767) */
	padding2      uint16         /*  */
}

/**
 *  \brief Joystick trackball motion event structure (event.jball.*)
 */
type SDL_JoyBallEvent struct {
	SDL_EventType                /**< ::SDL_JOYBALLMOTION */
	Timestamp     uint32         /*  */
	Which         SDL_JoystickID /**< The joystick instance id */
	Ball          uint8          /**< The joystick trackball index */
	padding       [3]uint8       /*  */
	Xrel          int16          /**< The relative motion in the X direction */
	Yrel          int16          /**< The relative motion in the Y direction */
}

/**
 *  \brief Joystick hat position change event structure (event.jhat.*)
 */
type SDL_JoyHatEvent struct {
	SDL_EventType                /**< ::SDL_JOYHATMOTION */
	Timestamp     uint32         /*  */
	Which         SDL_JoystickID /**< The joystick instance id */
	Hat           uint8          /**< The joystick hat index */
	Value         uint8          /**< The hat position value.
	 *   \sa ::SDL_HAT_LEFTUP ::SDL_HAT_UP ::SDL_HAT_RIGHTUP
	 *   \sa ::SDL_HAT_LEFT ::SDL_HAT_CENTERED ::SDL_HAT_RIGHT
	 *   \sa ::SDL_HAT_LEFTDOWN ::SDL_HAT_DOWN ::SDL_HAT_RIGHTDOWN
	 *
	 *   Note that zero means the POV is centered.
	 */
	padding [2]uint8
}

/**
 *  \brief Joystick button event structure (event.jbutton.*)
 */
type SDL_JoyButtonEvent struct {
	SDL_EventType                /**< ::SDL_JOYBUTTONDOWN or ::SDL_JOYBUTTONUP */
	Timestamp     uint32         /*  */
	Which         SDL_JoystickID /**< The joystick instance id */
	Button        uint8          /**< The joystick button index */
	State         uint8          /**< ::SDL_PRESSED or ::SDL_RELEASED */
	padding       [2]uint8
}

/**
 *  \brief Joystick device event structure (event.jdevice.*)
 */
type SDL_JoyDeviceEvent struct {
	SDL_EventType        /**< ::SDL_JOYDEVICEADDED or ::SDL_JOYDEVICEREMOVED */
	Timestamp     uint32 /*  */
	Which         int32  /**< The joystick device index for the ADDED event, instance id for the REMOVED event */
}

/**
 *  \brief Game controller axis motion event structure (event.caxis.*)
 */
type SDL_ControllerAxisEvent struct {
	SDL_EventType                /**< ::SDL_CONTROLLERAXISMOTION */
	Timestamp     uint32         /*  */
	Which         SDL_JoystickID /**< The joystick instance id */
	Axis          uint8          /**< The joystick axis index */
	padding       [3]uint8       /*  */
	Value         int16          /**< The axis value (range: -32768 to 32767) */
	padding2      uint16         /*  */
}

/**
 *  \brief Game controller button event structure (event.cbutton.*)
 */
type SDL_ControllerButtonEvent struct {
	SDL_EventType                /**< ::SDL_CONTROLLERBUTTONDOWN or ::SDL_CONTROLLERBUTTONUP */
	Timestamp     uint32         /*  */
	Which         SDL_JoystickID /**< The joystick instance id */
	Button        uint8          /**< The controller button (SDL_GameControllerButton) */
	State         uint8          /**< ::SDL_PRESSED or ::SDL_RELEASED */
	Sat           [2]uint8       /*  */
}

/**
 *  \brief Controller device event structure (event.cdevice.*)
 */
type SDL_ControllerDeviceEvent struct {
	SDL_EventType        /**< ::SDL_CONTROLLERDEVICEADDED, ::SDL_CONTROLLERDEVICEREMOVED, or ::SDL_CONTROLLERDEVICEREMAPPED */
	Timestamp     uint32 /*  */
	Which         int32  /**< The joystick device index for the ADDED event, instance id for the REMOVED or REMAPPED event */
}

/**
 *  \brief Audio device event structure (event.adevice.*)
 */
type SDL_AudioDeviceEvent struct {
	SDL_EventType          /**< ::SDL_AUDIODEVICEADDED, or ::SDL_AUDIODEVICEREMOVED */
	Timestamp     uint32   /*  */
	Which         uint32   /**< The audio device index for the ADDED event (valid until next SDL_GetNumAudioDevices() call), SDL_AudioDeviceID for the REMOVED event */
	IsCapture     uint8    /**< zero if an output device, non-zero if a capture device. */
	padding       [3]uint8 /*  */
}

/**
 *  \brief Touch finger event structure (event.tfinger.*)
 */
type SDL_TouchFingerEvent struct {
	SDL_EventType              /**< ::SDL_FINGERMOTION or ::SDL_FINGERDOWN or ::SDL_FINGERUP */
	Timestamp     uint32       /*  */
	TouchId       SDL_TouchID  /**< The touch device id */
	FingerId      SDL_FingerID /*  */
	X             float32      /**< Normalized in the range 0...1 */
	Y             float32      /**< Normalized in the range 0...1 */
	DX            float32      /**< Normalized in the range -1...1 */
	DY            float32      /**< Normalized in the range -1...1 */
	Pressure      float32      /**< Normalized in the range 0...1 */
}

/**
 *  \brief Multiple Finger Gesture Event (event.mgesture.*)
 */
type SDL_MultiGestureEvent struct {
	SDL_EventType             /**< ::SDL_MULTIGESTURE */
	Timestamp     uint32      /*  */
	TouchId       SDL_TouchID /**< The touch device id */
	DTheta        float32
	DDist         float32
	X             float32
	Y             float32
	NumFingers    uint16
	padding       uint16
}

/**
 * \brief Dollar Gesture Event (event.dgesture.*)
 */
type SDL_DollarGestureEvent struct {
	SDL_EventType               /**< ::SDL_DOLLARGESTURE or ::SDL_DOLLARRECORD */
	Timestamp     uint32        /*  */
	TouchId       SDL_TouchID   /**< The touch device id */
	GestureId     SDL_GestureID /*  */
	NumFingers    uint32        /*  */
	Error         float32       /*  */
	X             float32       /**< Normalized center of gesture */
	Y             float32       /**< Normalized center of gesture */
}

/**
 *  \brief An event used to request a file open by the system (event.drop.*)
 *         This event is enabled by default, you can disable it with SDL_EventState().
 *  \note If this event is enabled, you must free the filename in the event.
 */
type SDL_DropEvent struct {
	SDL_EventType        /**< ::SDL_DROPBEGIN or ::SDL_DROPFILE or ::SDL_DROPTEXT or ::SDL_DROPCOMPLETE */
	Timestamp     uint32 /*  */
	File          *byte  /**< The file name, which should be freed with SDL_free(), is NULL on begin/complete */
	WindowID      uint32 /**< The window that was dropped on, if any */
}

/**
 *  \brief The "quit requested" event
 */
type SDL_QuitEvent struct {
	SDL_EventType        /**< ::SDL_QUIT */
	Timestamp     uint32 /*  */
}

/**
 *  \brief OS Specific event
 */
type SDL_OSEvent struct {
	SDL_EventType        /**< ::SDL_QUIT */
	Timestamp     uint32 /*  */
}

/**
 *  \brief A user-defined event type (event.user.*)
 */
type SDL_UserEvent struct {
	SDL_EventType        /**< ::SDL_USEREVENT through ::SDL_LASTEVENT-1 */
	Timestamp     uint32 /*  */
	WindowID      uint32 /**< The window that was dropped on, if any */
	Code          int32  /**< User defined event code */
	Data1         *byte  /**< User defined data pointer */
	Data2         *byte  /**< User defined data pointer */
}

/**
 * \brief The types of events that can be delivered.
 */
const (
	SDL_FIRSTEVENT = 0 /**< Unused (do not remove) */

	/* Application events */
	SDL_QUIT = 0x100 /**< User-requested quit */

	/* These application events have special meaning on iOS see README-ios.md for details */
	SDL_APP_TERMINATING = 0x101 /**< The application is being terminated by the OS
	  Called on iOS in applicationWillTerminate()
	  Called on Android in onDestroy()
	*/
	SDL_APP_LOWMEMORY = 0x102 /**< The application is low on memory free memory if possible.
	  Called on iOS in applicationDidReceiveMemoryWarning()
	  Called on Android in onLowMemory()
	*/
	SDL_APP_WILLENTERBACKGROUND = 0x103 /**< The application is about to enter the background
	  Called on iOS in applicationWillResignActive()
	  Called on Android in onPause()
	*/
	SDL_APP_DIDENTERBACKGROUND = 0x104 /**< The application did enter the background and may not get CPU for some time
	  Called on iOS in applicationDidEnterBackground()
	  Called on Android in onPause()
	*/
	SDL_APP_WILLENTERFOREGROUND = 0x105 /**< The application is about to enter the foreground
	  Called on iOS in applicationWillEnterForeground()
	  Called on Android in onResume()
	*/
	SDL_APP_DIDENTERFOREGROUND = 0x106 /**< The application is now interactive
	  Called on iOS in applicationDidBecomeActive()
	  Called on Android in onResume()
	*/

	/* Window events */
	SDL_WINDOWEVENT = 0x200 /**< Window state change */
	SDL_SYSWMEVENT  = 0x201 /**< System specific event */

	/* Keyboard events */
	SDL_KEYDOWN       = 0x300 /**< Key pressed */
	SDL_KEYUP         = 0x301 /**< Key released */
	SDL_TEXTEDITING   = 0x302 /**< Keyboard text editing (composition) */
	SDL_TEXTINPUT     = 0x303 /**< Keyboard text input */
	SDL_KEYMAPCHANGED = 0x304 /**< Keymap changed due to a system event such as an
	  input language or keyboard layout change.
	*/

	/* Mouse events */
	SDL_MOUSEMOTION     = 0x400 /**< Mouse moved */
	SDL_MOUSEBUTTONDOWN = 0x401 /**< Mouse button pressed */
	SDL_MOUSEBUTTONUP   = 0x402 /**< Mouse button released */
	SDL_MOUSEWHEEL      = 0x403 /**< Mouse wheel motion */

	/* Joystick events */
	SDL_JOYAXISMOTION    = 0x600 /**< Joystick axis motion */
	SDL_JOYBALLMOTION    = 0x601 /**< Joystick trackball motion */
	SDL_JOYHATMOTION     = 0x602 /**< Joystick hat position change */
	SDL_JOYBUTTONDOWN    = 0x603 /**< Joystick button pressed */
	SDL_JOYBUTTONUP      = 0x604 /**< Joystick button released */
	SDL_JOYDEVICEADDED   = 0x605 /**< A new joystick has been inserted into the system */
	SDL_JOYDEVICEREMOVED = 0x606 /**< An opened joystick has been removed */

	/* Game controller events */
	SDL_CONTROLLERAXISMOTION     = 0x650 /**< Game controller axis motion */
	SDL_CONTROLLERBUTTONDOWN     = 0x651 /**< Game controller button pressed */
	SDL_CONTROLLERBUTTONUP       = 0x652 /**< Game controller button released */
	SDL_CONTROLLERDEVICEADDED    = 0x653 /**< A new Game controller has been inserted into the system */
	SDL_CONTROLLERDEVICEREMOVED  = 0x654 /**< An opened Game controller has been removed */
	SDL_CONTROLLERDEVICEREMAPPED = 0x655 /**< The controller mapping was updated */

	/* Touch events */
	SDL_FINGERDOWN   = 0x700
	SDL_FINGERUP     = 0x701
	SDL_FINGERMOTION = 0x702

	/* Gesture events */
	SDL_DOLLARGESTURE = 0x800
	SDL_DOLLARRECORD  = 0x801
	SDL_MULTIGESTURE  = 0x802

	/* Clipboard events */
	SDL_CLIPBOARDUPDATE = 0x900 /**< The clipboard changed */

	/* Drag and drop events */
	SDL_DROPFILE     = 0x1000 /**< The system requests a file open */
	SDL_DROPTEXT     = 0x1001 /**< text/plain drag-and-drop event */
	SDL_DROPBEGIN    = 0x1002 /**< A new set of drops is beginning (NULL filename) */
	SDL_DROPCOMPLETE = 0x1003 /**< Current set of drops is now complete (NULL filename) */

	/* Audio hotplug events */
	SDL_AUDIODEVICEADDED   = 0x1100 /**< A new audio device is available */
	SDL_AUDIODEVICEREMOVED = 0x1101 /**< An audio device has been removed. */

	/* Render events */
	SDL_RENDER_TARGETS_RESET = 0x2000 /**< The render targets have been reset and their contents need to be updated */
	SDL_RENDER_DEVICE_RESET  = 0x2001 /**< The device has been reset and all textures need to be recreated */

	/** Events ::SDL_USEREVENT through ::SDL_LASTEVENT are for your use
	 *  and should be allocated with SDL_RegisterEvents()
	 */
	SDL_USEREVENT = 0x8000

	/**
	 *  This last event is only for bounding internal arrays
	 */
	SDL_LASTEVENT = 0xFFFF
)
