package mock

type Retiever struct {
	Content string
}

func (r Retiever) GetFrom(url string) string{
	return r.Content
}

