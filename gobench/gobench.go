package bench
func map_add(){
	m := map[int]int{1:2,3:4}
	_ = m[1] + m[3]
}

func struct_add(){
	type mystruct struct{
		a int
		b int
	}
	s:= mystruct{a:2, b:4}
	_ = s.a + s.b 
}