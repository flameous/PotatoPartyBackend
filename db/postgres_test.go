package db

import (
	"testing"
	"log"
	"github.com/stretchr/testify/require"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestStorage(t *testing.T) {
	s, err := ConnectToDatabase(os.Getenv("DB_HOST"))
	require.NoError(t, err)

	// get user by id test
	user, err := s.GetUserByID(1)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotEmpty(t, user.Interests.Sport)
	require.Empty(t, user.Interests.Outdoor)
	require.Empty(t, user.Interests.Films)

	// get user by nickname
	userByNickname, err := s.GetUserByNickname(user.Nickname)
	require.NoError(t, err)
	require.NotNil(t, userByNickname)

	// get non-existing user by nickname
	userNotExist, err := s.GetUserByNickname("Rigfox_Not_Exists")
	require.NoError(t, err)
	require.Nil(t, userNotExist)

	userValidToken, err := s.GetUserByToken("123")
	require.NoError(t, err)
	require.NotNil(t, userValidToken)

	userInvalidToken, err := s.GetUserByToken("00000")
	require.NoError(t, err)
	require.Nil(t, userInvalidToken)

	events, err := s.GetAllEvents(user)
	require.NoError(t, err)
	require.NotEmpty(t, events)
	log.Printf("%#v", events.OwnEvents)
	log.Printf("%#v", events.Suggested)
	log.Printf("%#v", events.Attended)
	log.Printf("%#v", events.Private)
}
