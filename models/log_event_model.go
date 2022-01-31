package models

type LogEventModel struct {
	RequestUrl  string
	ResponseUrl string
}

func NewEvent(requestUrl string, responseUrl string) LogEventModel {
	return LogEventModel{
		RequestUrl:  requestUrl,
		ResponseUrl: responseUrl,
	}
}
