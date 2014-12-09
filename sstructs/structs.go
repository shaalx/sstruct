package sstructs

type Content struct {
	Contents map[string]Content
	Stuff    []string
}

func Init() Content {
	cnt := Content{}
	cnt.Stuff = []string{"Ipad-mini", "23"}
	cnt.Contents = nil
	return cnt
}
