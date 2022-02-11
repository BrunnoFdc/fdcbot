package errors

type BotError struct {
	ErrorType   uint16
	SourceError error
}

func (botError BotError) Error() string {
	sourceError := botError.SourceError

	if sourceError != nil {
		return sourceError.Error()
	} else {
		return errorCodeToMessage[botError.ErrorType]
	}
}

func (botError BotError) Message() string {
	return errorCodeToMessage[botError.ErrorType]
}
