package unsafe

import "unsafe"

func StringToBytes(s string) []byte {
	d := unsafe.StringData(s)
	return unsafe.Slice(d, len(s))
}
