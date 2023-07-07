



var a [4]int
a[0] = 1
i := a[0]
// i == 1
// a[2] == 0, the zero value of the int type
b := [2]string{"Penn", "Teller"} // literal 
b := [...]string{"penn","Teller"} // the compiler count the mount of element

// Slice
letters := []string{"a","b","c","d"}// you leave out the element count
func make([]T,len,cap)[]T // where T stands for the element type of the slice to be created. The make function takes a type, a length, and an optional capacity. When called, make allocates an array and returns a slice that refers to that array.
var s []byte
s = make([]byte,5,5)
//$>_ s == []byte{0,0,0,0,0}
//$>_ len(s) == 5
//$>_ cap(s) == 5
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
// b[:2] == []byte{'g', 'o'}
// b[2:] == []byte{'l', 'a', 'n', 'g'}
// b[:] == b
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // a slice referencing the storage of x
// points
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:]
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
