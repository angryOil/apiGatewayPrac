package cafe

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

var baseUrl = "http://localhost:8083/cafe"
