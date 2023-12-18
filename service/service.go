package service

type Service struct {
}

type IWord interface {
	GetAllWords(filePath string) (map[string]bool, error)
}

func NewService() IWord {
	return &Service{}
}
