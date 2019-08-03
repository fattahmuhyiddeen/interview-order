package model

import (
	"log"
	"strconv"
)

//EmailMessage is a model
type EmailMessage struct {
	ID           int    `json:"id" form:"id"`
	UserID       int    `json:"user_id" form:"user_id"`
	Title        string `json:"title" form:"title"`
	Recipients   string `json:"recipients" form:"recipients"`
	Sender       string `json:"sender" form:"sender"`
	Status       string `json:"status" form:"status"`
	IP           string `json:"ip" form:"ip"`
	UserAgent    string `json:"user_agent" form:"user_agent"`
	ShouldNotify bool   `json:"should_notify" form:"should_notify"`
	CreatedAt    string `json:"created_at" form:"created_at"`
	UpdatedAt    string `json:"updated_at" form:"updated_at"`
}

const emailMessageTable = "email_messages"
const emailMessageFields = "user_id, title, recipients, sender, status, should_notify, ip, user_agent, created_at, updated_at"

//InsertEmailMessage a
func InsertEmailMessage(emailMessage *EmailMessage) {
	connectDB()
	defer disconnectDB()

	emailMessage.CreatedAt = DateTimeNow()
	emailMessage.UpdatedAt = emailMessage.CreatedAt

	err := db.QueryRow(
		"INSERT INTO "+emailMessageTable+" ("+emailMessageFields+") VALUES ("+tablePlaceholder(emailMessageFields)+") RETURNING id",
		emailMessage.UserID,
		emailMessage.Title,
		emailMessage.Recipients,
		emailMessage.Sender,
		emailMessage.Status,
		emailMessage.ShouldNotify,
		emailMessage.IP,
		emailMessage.UserAgent,
		emailMessage.CreatedAt,
		emailMessage.UpdatedAt,
	).Scan(&emailMessage.ID)

	log.Println("Error while inserting Email Message to Database")
	log.Println(err)
}

//UpdateEmailMessage is to
func UpdateEmailMessage(emailMessage *EmailMessage) {
	connectDB()
	defer disconnectDB()

	emailMessage.CreatedAt = DateTimeNow()
	emailMessage.UpdatedAt = emailMessage.CreatedAt

	err := db.QueryRow(
		"UPDATE " + emailMessageTable + " SET status='" +
			emailMessage.Status + "', title='" + emailMessage.Title +
			"', recipients='" + emailMessage.Recipients +
			"' WHERE id=" + strconv.Itoa(emailMessage.ID))

	log.Println(err)
	log.Println(err)
}

//GetEmail is to search user by email
func GetEmail(id int) (emailMessage EmailMessage) {
	connectDB()
	defer disconnectDB()

	row := db.QueryRow("SELECT id, "+emailMessageFields+" FROM "+emailMessageTable+" WHERE id=$1", id)
	row.Scan(&emailMessage.ID, &emailMessage.UserID, &emailMessage.Title, &emailMessage.Recipients, &emailMessage.Sender, &emailMessage.Status, &emailMessage.ShouldNotify, &emailMessage.IP, &emailMessage.UserAgent, &emailMessage.CreatedAt, &emailMessage.UpdatedAt)
	return
}

//DeleteEmailMessage a
func DeleteEmailMessage(emailID int) {
	connectDB()
	defer disconnectDB()

	db.QueryRow("DELETE FROM " + emailMessageTable + " WHERE id=" + strconv.Itoa(emailID))
}
