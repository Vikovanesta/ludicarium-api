// Code generated by "stringer -type=PlatformCategory"; DO NOT EDIT.

package igdb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PlatformConsole-1]
	_ = x[PlatformArcade-2]
	_ = x[PlatformPlatform-3]
	_ = x[PlatformOperatingSystem-4]
	_ = x[PlatformPortableConsole-5]
	_ = x[PlatformComputer-6]
}

const _PlatformCategory_name = "PlatformConsolePlatformArcadePlatformPlatformPlatformOperatingSystemPlatformPortableConsolePlatformComputer"

var _PlatformCategory_index = [...]uint8{0, 15, 29, 45, 68, 91, 107}

func (i PlatformCategory) String() string {
	i -= 1
	if i < 0 || i >= PlatformCategory(len(_PlatformCategory_index)-1) {
		return "PlatformCategory(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _PlatformCategory_name[_PlatformCategory_index[i]:_PlatformCategory_index[i+1]]
}