package interfaces

type AuthUseCase interface {
	VerifyUser(email string, password string) error
	VerifyAdmin(email string, password string) error
	VerifyWorker(email string, password string) error
	SendVerificationEmail(email string) error
}
