// filepath: c:\Users\Cheng\Documents\f_admin\f_admin_go\internal\config\config_test.go
package test

import (
	"f_admin_go/internal/config"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("should load config when all environment variables are set", func(t *testing.T) {
		// Arrange
		os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname")
		os.Setenv("PORT", "8080")
		os.Setenv("ENVIRONMENT", "development")

		// Act
		cfg := config.LoadConfig()

		// Assert
		if cfg.DBURL != "postgres://user:password@localhost:5432/dbname" {
			t.Errorf("expected DBURL to be 'postgres://user:password@localhost:5432/dbname', got '%s'", cfg.DBURL)
		}
		if cfg.Port != "8080" {
			t.Errorf("expected Port to be '8080', got '%s'", cfg.Port)
		}
		if cfg.Environment != "development" {
			t.Errorf("expected Environment to be 'development', got '%s'", cfg.Environment)
		}
	})

	t.Run("should log fatal error when DATABASE_URL is not set", func(t *testing.T) {
		// Arrange
		os.Unsetenv("DATABASE_URL")
		os.Setenv("PORT", "8080")
		os.Setenv("ENVIRONMENT", "development")

		// Act & Assert
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected log.Fatal to be called, but it was not")
			}
		}()
		config.LoadConfig()
	})
}
