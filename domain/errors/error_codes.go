package errors

// Os 10 primeiros códigos de erro serão reservados para tipos de erros genéricos
const (
	UnexpectedError  = 0
	InvalidGroupRole = 10
)

var errorCodeToMessage = map[uint16]string{
	UnexpectedError:  "Erro inesperado.",
	InvalidGroupRole: "O cargo informado não representa um grupo válido.",
}
