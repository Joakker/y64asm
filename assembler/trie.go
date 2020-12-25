package assembler

func getFn(s string) uint16 {
	switch s[0] {
	case 'l':
		if len(s) == 2 {
			return FnLe
		} else if len(s) == 1 {
			return FnLs
		}
	case 'g':
		if len(s) == 2 {
			return FnGe
		} else if len(s) == 1 {
			return FnGr
		}
	case 'e':
		return FnEq
	case 'n':
		return FnNe
	}
	return FnUn
}
