package models

type Certificate struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Timestamp int64  `json:"timestamp"`
	Signature []byte `json:"signature"`
}
