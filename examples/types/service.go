//#define SERVICE (METHODS) (FIELDS) (INIT)
type Service interface {
	METHODS
}

type service struct {
	FIELDS
}

var _ Service = (*service)(nil)

func NewService() Service {
	INIT
}

//#end
