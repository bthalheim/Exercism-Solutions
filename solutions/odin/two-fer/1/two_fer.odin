package two_fer

import "core:strings"

two_fer :: proc(name: string = "you") -> string {
	return strings.concatenate({"One for ", name, ", one for me."})
}
