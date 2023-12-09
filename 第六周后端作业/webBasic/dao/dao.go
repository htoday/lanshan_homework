package dao

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

var Database = map[string]string{
	"lzh": "666666",
}

type Message struct {
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var Messages []Message

func WriteMessagesToFile() {
	file, err := json.MarshalIndent(Messages, "", "  ")
	if err != nil {
		//fmt.Println("Error encoding JSON:", err)
		return
	}

	err = os.WriteFile("messages.json", file, 0644)
	if err != nil {
		//fmt.Println("Error writing to file:", err)
		return
	}
}
func ReadMessagesFromFile() {
	file, err := os.ReadFile("messages.json")
	if err != nil {
		//fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(file, &Messages)
	if err != nil {
		//fmt.Println("Error decoding JSON:", err)
		return
	}
}
func SaveToJSON(filename string, data map[string]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
func LoadFormJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	jsonData, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &Database)
	if err != nil {
		return err
	}
	return nil
}
func AddUser(userName string, password string) {
	Database[userName] = password

}
func CheckUser(userName string) bool {
	if Database[userName] == "" {
		return false
	}
	return true
}
func FindPasswordFormUserName(username string) string {
	return Database[username]
}
func ChangePassword(username string, newPassword string) {
	delete(Database, username)
	Database[username] = newPassword
}
