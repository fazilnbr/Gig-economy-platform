package interfaces

type AuthUseCase interface {
	VerifyUser(email string, password string) error
	VerifyAdmin(email string, password string) error
	VerifyWorker(email string, password string) error
	SendOTP(phone string) error
	SendVerificationEmail(email string) error
	UserVerifyAccount(email string, code string) error
	WorkerVerifyAccount(email string, code int) error
}
