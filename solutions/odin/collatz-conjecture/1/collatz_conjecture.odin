package collatz_conjecture


import "core:fmt"

steps :: proc(start: int) -> (result: int, ok: bool) {

	if start <= 0 {
		return 0, false
	}
	
	s := 0 
	i := start

	for i!=1 {
		
		if i%2==0 {
			i = i/2 
		} else {
			i = (i*3)+1
		}

		s = s + 1 
		
	}

	return s, true

}




