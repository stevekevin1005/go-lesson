package mock

type Retreiver struct {
	Contents string
}

func (r Retreiver) String() string {
	return "test"
}

func (r *Retreiver) Post(url string, form map[string]string) string {
	r.Contents = form["name"] + form["course"]
	return "ok"
}

func (r Retreiver) Get(url string) string {
	return r.Contents
}
