package main

type ValNameTest [2]string

// Value State is a mapping from variable names to values
type ValStateTest struct {
	vals map[ValNameTest]int
}

/*
func testDeepCopy() {
	x := ValNameTest{"x", "global"}
	y := ValNameTest{"y", "global"}
	m := make(map[ValNameTest]int)
	s := ValStateTest{m}
	s.vals[x] = 1
	s.vals[y] = 2
	m2 := make(map[ValNameTest]int)
	s2 := ValStateTest{m2}

	for k, v := range s.vals {
		s.vals[k] = v
	}
	s2.vals[x] = 3
	y2 := ValNameTest{"y", "global-if-else"}
	s2.vals[y2] = 4

	for k := range s2.vals {
		_, ok := s.vals[k]
		println(k[0] + k[1])
		if ok {
			s.vals[k] = s2.vals[k]
		}
	}
	for k, _ := range s.vals {
		print(k[0])
		print(s.vals[k])
	}

}
*/
