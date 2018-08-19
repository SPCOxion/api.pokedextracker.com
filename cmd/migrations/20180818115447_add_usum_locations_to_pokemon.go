package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE pokemon
				ADD COLUMN us_location TEXT,
				ADD COLUMN um_location TEXT
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE pokemon
				DROP COLUMN us_location,
				DROP COLUMN um_location
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20180818115447_add_usum_locations_to_pokemon", up, down, opts)
}