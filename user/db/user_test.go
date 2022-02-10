package db

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

var userDb *UserDB

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")
	if err != nil {
		logger.Fatal(err)
	}
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	userDb = NewUser(dbConn)
}

func TestLoginQuery(t *testing.T) {

	md5Pass := fmt.Sprintf("%x", md5.Sum([]byte("essay")))
	md5Pass = strings.ToUpper(md5Pass)
	examples := []struct {
		Username string
		Password string

		Except bool
	}{
		{
			Username: "121121121",
			Password: md5Pass,
			Except:   true,
		},
		{
			Username: "123123123",
			Password: md5Pass,
			Except:   true,
		},
		{
			Username: "no",
			Password: md5Pass,
			Except:   false,
		},
	}

	for index, example := range examples {
		t.Run("LoginQuery"+strconv.Itoa(index), func(t *testing.T) {
			res, err := userDb.LoginQuery(example.Username, example.Password)
			if err != nil {
				t.Fatal(err)
			}
			if res != example.Except {
				t.Error(res)
			}
		})

	}
}
