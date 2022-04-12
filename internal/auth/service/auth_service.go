package authServicePkg

type AuthService interface {
	CreateJWT() (string, error)
}
