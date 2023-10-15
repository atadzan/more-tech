package handler

type response struct {
	Message string
}

func newResponse(msg string) response {
	return response{Message: msg}
}
