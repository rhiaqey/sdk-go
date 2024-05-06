package sdk

type ProducerMessage struct {
	Key       string      `json:"key"`
	Value     interface{} `json:"val"`
	Timestamp *uint64     `json:"tms,omitempty"`
	Tag       *string     `json:"tag,omitempty"`
	Category  *string     `json:"cat,omitempty"`
	Size      *uint       `json:"siz,omitempty"`
	UserIds   *[]string   `json:"uid,omitempty"`
	ClientIds *[]string   `json:"cid,omitempty"`
}

type Producer interface {
	Setup(settings interface{}) (chan ProducerMessage, error)
	Start()
	SetSettings(interface{})
	Schema() []byte
	Metrics() []byte
	Kind() string
}
