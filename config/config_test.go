package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	const dbHost string = "localhost_test"
	const dbPort string = "1234"
	const dbUsername string = "dummy_user"
	const dbPassword string = "dummy_pass"
	const dbName string = "dummy_db"
	const appTitle string = "Kubbe Test"
	const appHost string = "localhost_test:1234"

	_ = os.Setenv("KUBBE_DB_HOST", dbHost)
	_ = os.Setenv("KUBBE_DB_PORT", dbPort)
	_ = os.Setenv("KUBBE_DB_USERNAME", dbUsername)
	_ = os.Setenv("KUBBE_DB_PASSWORD", dbPassword)
	_ = os.Setenv("KUBBE_DB_NAME", dbName)
	_ = os.Setenv("KUBBE_APP_TITLE", appTitle)
	_ = os.Setenv("KUBBE_APP_HOST", appHost)

	c := GetConfig()

	assertEquals(t, c.DB.Host, dbHost)
	assertEquals(t, c.DB.Port, dbPort)
	assertEquals(t, c.DB.Username, dbUsername)
	assertEquals(t, c.DB.Password, dbPassword)
	assertEquals(t, c.DB.Name, dbName)
	assertEquals(t, c.App.Title, appTitle)
	assertEquals(t, c.App.Host, appHost)
}

func assertEquals(t *testing.T, value string, envVarValue string) {
	if value != envVarValue {
		t.Errorf("Expected %s, but got '%s'", envVarValue, value)
	}
}
