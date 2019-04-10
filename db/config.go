package db

import mgo "gopkg.in/mgo.v2"

type Config struct {
	URI         string
	SSL         bool
	Creds       mgo.Credential
	Database    string
	Collections []string
}

func (p *Config) Load() error {
	return nil
}

func (p *Config) validate() error {
	return nil
}
