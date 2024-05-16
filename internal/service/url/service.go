package url

type Repository interface {
	SaveURL(urlToSave string, alias string) error
	GetURL(alias string) (string, error)
}

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) GetURL(alias string) (string, error) {
	alias, err := s.Repository.GetURL(alias)
	return alias, err
}
func (s *Service) SaveURL(url string, alias string) error {
	err := s.Repository.SaveURL(url, alias)
	if err != nil {
		return err
	}
	return nil

}
