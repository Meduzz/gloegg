package common

type (
	Tag struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	Tags []*Tag
)
