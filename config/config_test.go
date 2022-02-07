package config

import "testing"

func TestLoadDatabaseConfig(t *testing.T) {
	database, err := LoadDatabase()
	if err != nil {
		t.Error(err)
	}
	excepet := Database{
		Host:     "172.17.0.3",
		Database: "hpcmanager",
		Port:     3306,
		Username: "root",
		Password: "beihai",
	}
	if excepet != *database {
		t.Errorf("Except: %#v, Get: %#v", excepet, database)
	}
}
