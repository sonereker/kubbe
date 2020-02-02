package config

import (
	"github.com/stretchr/testify/assert"
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

	assert.EqualValues(t, c.DB.Host, dbHost)
	assert.EqualValues(t, c.DB.Host, dbHost)
	assert.EqualValues(t, c.DB.Port, dbPort)
	assert.EqualValues(t, c.DB.Username, dbUsername)
	assert.EqualValues(t, c.DB.Password, dbPassword)
	assert.EqualValues(t, c.DB.Name, dbName)
	assert.EqualValues(t, c.App.Title, appTitle)
	assert.EqualValues(t, c.App.Host, appHost)
}
