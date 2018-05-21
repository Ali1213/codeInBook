package main

func main() {

}
func eliminateString(a []string) []string {
	s := a[:0]
	str := a[0]
	for i := 1; i < len(a); i++ {
		if str == "" {
			str = a[i]
			continue
		}
		if str == a[i] {
			str = ""
			continue
		} else {
			str = a[i]
			s := append(s, str)
		}
	}
	return s
}
