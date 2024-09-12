package internal

var WEBHOOK string = "https://prod-02.australiasoutheast.logic.azure.com:443/workflows/e60c2e89b72b458f9a2ff0b37affec4a/triggers/manual/paths/invoke?api-version=2016-06-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=y8XAivf3HU1EXm7dQK0WNss78eM8wBGGIdOZ9HHaTrM"

type MessageBody struct {
	Title string `json:"message_title,omitempty"`
	Body  string `json:"message_body,omitempty"`
}
