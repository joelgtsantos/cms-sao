package model

const (
	NesoMessageEntryKind = "entry"
	NesoMessageDraftKind = "draft"
)

type NesoMessage struct {
	Kind string `json:"kind"`
	Auth struct {
		Cookies []string `json:"cookies"`
	} `json:"auth"`
	Transaction struct {
		ID string `json:"id"`
	} `json:"transaction"`
	EntryPayload struct {
		ContestSlug string `json:"contestSlug"`
		TaskSlug    string `json:"taskSlug"`
		Token       bool   `json:"token"`
		Sources     []struct {
			Filename string `json:"filename"`
			FileID   string `json:"fileid"`
			Language string `json:"language"`
			Content  string `json:"content"`
		} `json:"sources"`
	} `json:"entry"`
}
