package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Site    string
	Url     string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
