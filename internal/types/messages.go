package types

type Messages []Message

func (m *Messages) AddMessage(message Message) Message {
	*m = append(*m, message)
	return message
}

func (m *Messages) FindMessage(name string) *Message {
	for _, message := range *m {
		if message.Name == name {
			return &message
		}
	}
	return nil
}

func (m *Messages) DelMessage(name string) {
	for i, message := range *m {
		if message.Name == name {
			*m = append((*m)[:i], (*m)[i+1:]...)
			return
		}
	}
}
