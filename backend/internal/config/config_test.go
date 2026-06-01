package config

import "testing"

func TestLoadUsesRailwayPortAndDatabaseURL(t *testing.T) {
	t.Setenv("TRAMPLIN_HTTP_ADDR", "")
	t.Setenv("PORT", "4321")
	t.Setenv("TRAMPLIN_DATABASE_URL", "")
	t.Setenv("DATABASE_URL", "postgres://railway:secret@postgres.railway.internal:5432/railway")

	cfg := Load()

	if cfg.HTTPAddr != ":4321" {
		t.Fatalf("HTTPAddr = %q, want %q", cfg.HTTPAddr, ":4321")
	}
	if cfg.DatabaseURL != "postgres://railway:secret@postgres.railway.internal:5432/railway" {
		t.Fatalf("DatabaseURL = %q, want Railway DATABASE_URL", cfg.DatabaseURL)
	}
}

func TestLoadPrefersAppSpecificOverrides(t *testing.T) {
	t.Setenv("TRAMPLIN_HTTP_ADDR", ":9090")
	t.Setenv("PORT", "4321")
	t.Setenv("TRAMPLIN_DATABASE_URL", "postgres://tramplin:secret@db.internal:5432/tramplin")
	t.Setenv("DATABASE_URL", "postgres://railway:secret@postgres.railway.internal:5432/railway")

	cfg := Load()

	if cfg.HTTPAddr != ":9090" {
		t.Fatalf("HTTPAddr = %q, want %q", cfg.HTTPAddr, ":9090")
	}
	if cfg.DatabaseURL != "postgres://tramplin:secret@db.internal:5432/tramplin" {
		t.Fatalf("DatabaseURL = %q, want TRAMPLIN_DATABASE_URL", cfg.DatabaseURL)
	}
}
