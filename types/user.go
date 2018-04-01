package types

import (
	"time"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID         int64         `json:"id"`
	Nickname   string        `json:"nickname"`
	Password   string        `json:"-"`
	Email      string        `json:"email"`
	LastSeen   time.Time     `json:"last_seen"`
	Lat        float32       `json:"lat"`
	Lon        float32       `json:"lon"`
	Interests  *AllInterests `json:"interests"`
	Token      string        `json:"token,omitempty"`
	PictureURL string        `json:"picture_url"`
}

func (u *User) GenerateNewToken() {
	hash := md5.New()
	hash.Write([]byte(time.Now().String() + u.Nickname))
	u.Token = hex.EncodeToString(hash.Sum(nil))
}
