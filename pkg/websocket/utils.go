package websocket

import (
	"bytes"
	"html/template"
	"log"
)

func getTemplate(templatePath string, msg *Message) []byte {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	var renderedMessage bytes.Buffer
	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderedMessage.Bytes()
}
