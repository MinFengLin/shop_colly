package crawler

type momo_j struct {
	Momo []struct {
		Item         string
		Url          string
		Target_price string
	} `json:"momo"`
}