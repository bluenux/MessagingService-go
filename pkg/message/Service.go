package message

type Service interface {
	GetMessage()
	SendMessage()
}

type service struct {
}

func (s service) GetMessage() {
	//TODO implement me
	panic("implement me")
}

func (s service) SendMessage() {
	//TODO implement me
	panic("implement me")
}

func NewService() Service {
	return &service{}
}
