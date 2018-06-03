package url

type Values map[string][]string

func (v Values) Get(k string) string {
	if vs := v[k]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(k, vs string) {
	v[k] = append(v[k], vs)
}
