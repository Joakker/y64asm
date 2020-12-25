package assembler

const PlaceHolder = 1024

const (
	Halt  = 0x00
	Nop   = 0x10
	IRMov = 0x30
	RMMov = 0x40
	MRMov = 0x50
	Call  = 0x80
	Ret   = 0x90
)

const (
	FnUn = iota
	FnLe
	FnLs
	FnEq
	FnNe
	FnGe
	FnGr
)

func Jmp(s string) uint16 {
	fn := getFn(s)
	return 0x70 + fn
}

func CMov(s string) uint16 {
	fn := getFn(s)
	return 0x20 + fn
}
