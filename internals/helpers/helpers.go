package helpers

import (
	"nextgen/templates/components"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

var AlertId int

type FlashMessage struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	Id        string `json:"id"`
	DismissId string `json:"dismiss_id"`
}

func GetAlerts(c *gin.Context) ([]components.AlertProps, error) {
	var alerts []components.AlertProps
	flashes, _ := GetFlash(c, "flash")
	for _, f := range flashes {
		alerts = append(alerts, components.AlertProps{Type: f.Type, Message: f.Message, Id: f.Id, DismissId: f.DismissId})
	}
	return alerts, nil

}
func SetFlash(c *gin.Context, key string, value FlashMessage) {
	strAlertId := fmt.Sprintf("%d", AlertId)
	value.Id = "alert-" + strAlertId
	value.DismissId = "#" + value.Id
	AlertId = AlertId + 1
	fmt.Println(AlertId)
	session := sessions.Default(c)
	flashes := session.Flashes()
	if len(flashes) == 0 {
		var flashMessages []FlashMessage

		flashMessages = append(flashMessages, value)

		flashMessageStr, err := json.Marshal(flashMessages)
		if err != nil {
			log.Printf("Error marshalling flash message: %v\n", err)
			return
		}

		session.AddFlash(string(flashMessageStr))
		log.Println(string(flashMessageStr))
		if err := session.Save(); err != nil {
			log.Printf("Error saving session: %v\n", err)
		}

	} else {

		var flashMessages []FlashMessage
		err := json.Unmarshal([]byte(flashes[0].(string)), &flashMessages)

		flashMessages = append(flashMessages, value)

		flashMessageStr, err := json.Marshal(flashMessages)
		if err != nil {
			log.Printf("Error marshalling flash message: %v\n", err)
			return
		}

		session.AddFlash(string(flashMessageStr))
		fmt.Println(flashMessages)
		if err := session.Save(); err != nil {
			log.Printf("Error saving session: %v\n", err)
		}

	}

}

func GetFlash(c *gin.Context, key string) ([]FlashMessage, error) {
	session := sessions.Default(c)
	if session == nil {
		log.Println("Session is not initialized")
		return nil, fmt.Errorf("session not initialized")
	}

	flashes := session.Flashes()
	var flashMessages []FlashMessage
	if len(flashes) == 0 {
		return nil, nil

	}
	err := json.Unmarshal([]byte(flashes[0].(string)), &flashMessages)
	if err != nil {
		log.Printf("Error unmarshalling flash message: %v\n", err)
		return nil, err

	}

	fmt.Println(flashMessages)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	return flashMessages, nil
}
