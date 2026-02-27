package shared

const (
	MemorySize    uint16 = 4096
	DisplayWidth         = 64
	DisplayHeight        = 32
	PixelSize            = 10
)

type KeyEvent struct {
	Key     uint8
	Pressed bool
}
