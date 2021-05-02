package errors

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  400,
		Error:   "Bad Request",
	}
}

func NewNotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  404,
		Error:   "Not Found",
	}
}

func NewAlreadyExistsError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  409,
		Error:   "Already Exists",
	}
}
