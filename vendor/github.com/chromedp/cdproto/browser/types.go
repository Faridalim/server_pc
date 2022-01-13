package browser

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// WindowID [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-WindowID
type WindowID int64

// Int64 returns the WindowID as int64 value.
func (t WindowID) Int64() int64 {
	return int64(t)
}

// WindowState the state of the browser window.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-WindowState
type WindowState string

// String returns the WindowState as string value.
func (t WindowState) String() string {
	return string(t)
}

// WindowState values.
const (
	WindowStateNormal     WindowState = "normal"
	WindowStateMinimized  WindowState = "minimized"
	WindowStateMaximized  WindowState = "maximized"
	WindowStateFullscreen WindowState = "fullscreen"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t WindowState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t WindowState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *WindowState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch WindowState(in.String()) {
	case WindowStateNormal:
		*t = WindowStateNormal
	case WindowStateMinimized:
		*t = WindowStateMinimized
	case WindowStateMaximized:
		*t = WindowStateMaximized
	case WindowStateFullscreen:
		*t = WindowStateFullscreen

	default:
		in.AddError(errors.New("unknown WindowState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *WindowState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Bounds browser window bounds information.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-Bounds
type Bounds struct {
	Left        int64       `json:"left,omitempty"`        // The offset from the left edge of the screen to the window in pixels.
	Top         int64       `json:"top,omitempty"`         // The offset from the top edge of the screen to the window in pixels.
	Width       int64       `json:"width,omitempty"`       // The window width in pixels.
	Height      int64       `json:"height,omitempty"`      // The window height in pixels.
	WindowState WindowState `json:"windowState,omitempty"` // The window state. Default to normal.
}

// PermissionType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-PermissionType
type PermissionType string

// String returns the PermissionType as string value.
func (t PermissionType) String() string {
	return string(t)
}

// PermissionType values.
const (
	PermissionTypeAccessibilityEvents      PermissionType = "accessibilityEvents"
	PermissionTypeAudioCapture             PermissionType = "audioCapture"
	PermissionTypeBackgroundSync           PermissionType = "backgroundSync"
	PermissionTypeBackgroundFetch          PermissionType = "backgroundFetch"
	PermissionTypeClipboardReadWrite       PermissionType = "clipboardReadWrite"
	PermissionTypeClipboardSanitizedWrite  PermissionType = "clipboardSanitizedWrite"
	PermissionTypeDisplayCapture           PermissionType = "displayCapture"
	PermissionTypeDurableStorage           PermissionType = "durableStorage"
	PermissionTypeFlash                    PermissionType = "flash"
	PermissionTypeGeolocation              PermissionType = "geolocation"
	PermissionTypeMidi                     PermissionType = "midi"
	PermissionTypeMidiSysex                PermissionType = "midiSysex"
	PermissionTypeNfc                      PermissionType = "nfc"
	PermissionTypeNotifications            PermissionType = "notifications"
	PermissionTypePaymentHandler           PermissionType = "paymentHandler"
	PermissionTypePeriodicBackgroundSync   PermissionType = "periodicBackgroundSync"
	PermissionTypeProtectedMediaIdentifier PermissionType = "protectedMediaIdentifier"
	PermissionTypeSensors                  PermissionType = "sensors"
	PermissionTypeVideoCapture             PermissionType = "videoCapture"
	PermissionTypeVideoCapturePanTiltZoom  PermissionType = "videoCapturePanTiltZoom"
	PermissionTypeIdleDetection            PermissionType = "idleDetection"
	PermissionTypeWakeLockScreen           PermissionType = "wakeLockScreen"
	PermissionTypeWakeLockSystem           PermissionType = "wakeLockSystem"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PermissionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PermissionType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PermissionType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PermissionType(in.String()) {
	case PermissionTypeAccessibilityEvents:
		*t = PermissionTypeAccessibilityEvents
	case PermissionTypeAudioCapture:
		*t = PermissionTypeAudioCapture
	case PermissionTypeBackgroundSync:
		*t = PermissionTypeBackgroundSync
	case PermissionTypeBackgroundFetch:
		*t = PermissionTypeBackgroundFetch
	case PermissionTypeClipboardReadWrite:
		*t = PermissionTypeClipboardReadWrite
	case PermissionTypeClipboardSanitizedWrite:
		*t = PermissionTypeClipboardSanitizedWrite
	case PermissionTypeDisplayCapture:
		*t = PermissionTypeDisplayCapture
	case PermissionTypeDurableStorage:
		*t = PermissionTypeDurableStorage
	case PermissionTypeFlash:
		*t = PermissionTypeFlash
	case PermissionTypeGeolocation:
		*t = PermissionTypeGeolocation
	case PermissionTypeMidi:
		*t = PermissionTypeMidi
	case PermissionTypeMidiSysex:
		*t = PermissionTypeMidiSysex
	case PermissionTypeNfc:
		*t = PermissionTypeNfc
	case PermissionTypeNotifications:
		*t = PermissionTypeNotifications
	case PermissionTypePaymentHandler:
		*t = PermissionTypePaymentHandler
	case PermissionTypePeriodicBackgroundSync:
		*t = PermissionTypePeriodicBackgroundSync
	case PermissionTypeProtectedMediaIdentifier:
		*t = PermissionTypeProtectedMediaIdentifier
	case PermissionTypeSensors:
		*t = PermissionTypeSensors
	case PermissionTypeVideoCapture:
		*t = PermissionTypeVideoCapture
	case PermissionTypeVideoCapturePanTiltZoom:
		*t = PermissionTypeVideoCapturePanTiltZoom
	case PermissionTypeIdleDetection:
		*t = PermissionTypeIdleDetection
	case PermissionTypeWakeLockScreen:
		*t = PermissionTypeWakeLockScreen
	case PermissionTypeWakeLockSystem:
		*t = PermissionTypeWakeLockSystem

	default:
		in.AddError(errors.New("unknown PermissionType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PermissionType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PermissionSetting [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-PermissionSetting
type PermissionSetting string

// String returns the PermissionSetting as string value.
func (t PermissionSetting) String() string {
	return string(t)
}

// PermissionSetting values.
const (
	PermissionSettingGranted PermissionSetting = "granted"
	PermissionSettingDenied  PermissionSetting = "denied"
	PermissionSettingPrompt  PermissionSetting = "prompt"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PermissionSetting) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PermissionSetting) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PermissionSetting) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PermissionSetting(in.String()) {
	case PermissionSettingGranted:
		*t = PermissionSettingGranted
	case PermissionSettingDenied:
		*t = PermissionSettingDenied
	case PermissionSettingPrompt:
		*t = PermissionSettingPrompt

	default:
		in.AddError(errors.New("unknown PermissionSetting value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PermissionSetting) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PermissionDescriptor definition of PermissionDescriptor defined in the
// Permissions API:
// https://w3c.github.io/permissions/#dictdef-permissiondescriptor.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-PermissionDescriptor
type PermissionDescriptor struct {
	Name                     string `json:"name"`                               // Name of permission. See https://cs.chromium.org/chromium/src/third_party/blink/renderer/modules/permissions/permission_descriptor.idl for valid permission names.
	Sysex                    bool   `json:"sysex,omitempty"`                    // For "midi" permission, may also specify sysex control.
	UserVisibleOnly          bool   `json:"userVisibleOnly,omitempty"`          // For "push" permission, may specify userVisibleOnly. Note that userVisibleOnly = true is the only currently supported type.
	AllowWithoutSanitization bool   `json:"allowWithoutSanitization,omitempty"` // For "clipboard" permission, may specify allowWithoutSanitization.
	PanTiltZoom              bool   `json:"panTiltZoom,omitempty"`              // For "camera" permission, may specify panTiltZoom.
}

// CommandID browser command ids used by executeBrowserCommand.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-BrowserCommandId
type CommandID string

// String returns the CommandID as string value.
func (t CommandID) String() string {
	return string(t)
}

// CommandID values.
const (
	CommandIDOpenTabSearch  CommandID = "openTabSearch"
	CommandIDCloseTabSearch CommandID = "closeTabSearch"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CommandID) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CommandID) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CommandID) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CommandID(in.String()) {
	case CommandIDOpenTabSearch:
		*t = CommandIDOpenTabSearch
	case CommandIDCloseTabSearch:
		*t = CommandIDCloseTabSearch

	default:
		in.AddError(errors.New("unknown CommandID value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CommandID) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Bucket chrome histogram bucket.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-Bucket
type Bucket struct {
	Low   int64 `json:"low"`   // Minimum value (inclusive).
	High  int64 `json:"high"`  // Maximum value (exclusive).
	Count int64 `json:"count"` // Number of samples.
}

// Histogram chrome histogram.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#type-Histogram
type Histogram struct {
	Name    string    `json:"name"`    // Name.
	Sum     int64     `json:"sum"`     // Sum of sample values.
	Count   int64     `json:"count"`   // Total number of samples.
	Buckets []*Bucket `json:"buckets"` // Buckets.
}

// DownloadProgressState download status.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#event-downloadProgress
type DownloadProgressState string

// String returns the DownloadProgressState as string value.
func (t DownloadProgressState) String() string {
	return string(t)
}

// DownloadProgressState values.
const (
	DownloadProgressStateInProgress DownloadProgressState = "inProgress"
	DownloadProgressStateCompleted  DownloadProgressState = "completed"
	DownloadProgressStateCanceled   DownloadProgressState = "canceled"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DownloadProgressState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DownloadProgressState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DownloadProgressState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch DownloadProgressState(in.String()) {
	case DownloadProgressStateInProgress:
		*t = DownloadProgressStateInProgress
	case DownloadProgressStateCompleted:
		*t = DownloadProgressStateCompleted
	case DownloadProgressStateCanceled:
		*t = DownloadProgressStateCanceled

	default:
		in.AddError(errors.New("unknown DownloadProgressState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DownloadProgressState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SetDownloadBehaviorBehavior whether to allow all or deny all download
// requests, or use default Chrome behavior if available (otherwise deny).
// |allowAndName| allows download and names files according to their dowmload
// guids.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-setDownloadBehavior
type SetDownloadBehaviorBehavior string

// String returns the SetDownloadBehaviorBehavior as string value.
func (t SetDownloadBehaviorBehavior) String() string {
	return string(t)
}

// SetDownloadBehaviorBehavior values.
const (
	SetDownloadBehaviorBehaviorDeny         SetDownloadBehaviorBehavior = "deny"
	SetDownloadBehaviorBehaviorAllow        SetDownloadBehaviorBehavior = "allow"
	SetDownloadBehaviorBehaviorAllowAndName SetDownloadBehaviorBehavior = "allowAndName"
	SetDownloadBehaviorBehaviorDefault      SetDownloadBehaviorBehavior = "default"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetDownloadBehaviorBehavior) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetDownloadBehaviorBehavior) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetDownloadBehaviorBehavior) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SetDownloadBehaviorBehavior(in.String()) {
	case SetDownloadBehaviorBehaviorDeny:
		*t = SetDownloadBehaviorBehaviorDeny
	case SetDownloadBehaviorBehaviorAllow:
		*t = SetDownloadBehaviorBehaviorAllow
	case SetDownloadBehaviorBehaviorAllowAndName:
		*t = SetDownloadBehaviorBehaviorAllowAndName
	case SetDownloadBehaviorBehaviorDefault:
		*t = SetDownloadBehaviorBehaviorDefault

	default:
		in.AddError(errors.New("unknown SetDownloadBehaviorBehavior value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetDownloadBehaviorBehavior) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
