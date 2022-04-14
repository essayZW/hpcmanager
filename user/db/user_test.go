package db

import (
	"context"
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

func getDBConn() *db.DB {
	dbConn, err := db.NewDB()
	if err != nil {
		logger.Fatal(err)
	}
	return dbConn
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
			res, err := userDb.LoginQuery(context.Background(), example.Username, example.Password)
			if err != nil {
				t.Fatal(err)
			}
			if res != example.Except {
				t.Error(res)
			}
		})

	}
}

func TestQueryByGroupID(t *testing.T) {
	tests := []struct {
		Name string

		GroupID int

		ExceptCount int
		Error       bool
	}{
		{
			Name:        "test success",
			GroupID:     0,
			ExceptCount: 3,
			Error:       false,
		},
		{
			Name:        "test success2",
			GroupID:     1,
			ExceptCount: 1,
			Error:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ids, err := userDb.QueryUserByGroupID(context.Background(), test.GroupID)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v, Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v, Except: %v", err, test.Error)
			}
			if len(ids) != test.ExceptCount {
				t.Errorf("Get: %v, Except: %v", ids, test.ExceptCount)
			}
		})
	}
}

func TestUserDB_QueryCountWithPYNamePrefix(t *testing.T) {
	type fields struct {
		conn *db.DB
	}
	type args struct {
		ctx    context.Context
		prefix string
	}
	conn := getDBConn()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test shenqingzhe",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:    context.Background(),
				prefix: "shenqingzhe",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "test dalao",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:    context.Background(),
				prefix: "dalao",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "test shenqingzu",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:    context.Background(),
				prefix: "shenqingzu",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "test none",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx:    context.Background(),
				prefix: "none",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &UserDB{
				conn: tt.fields.conn,
			}
			got, err := db.QueryCountWithPYNamePrefix(tt.args.ctx, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UserDB.QueryCountWithPYNamePrefix() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if got != tt.want {
				t.Errorf("UserDB.QueryCountWithPYNamePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
