package types

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (m *Message) SetName(name string) {
	m.Name = name
}

func (m *Message) SetMessage(message string) {
	m.Message = message
}
