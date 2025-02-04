package db

import (
	"gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}

type conn struct {
	session *mgo.Session
	dataBase *mgo.Database
}

func NewConnection(cfg Config) (Connection, error) {
	fmt.Println("Connecting to database...")
	fmt.Println("DNS: ", cfg.Dns())

	session, err := mgo.Dial(c.Dns())
	if err != nil {
		log.Fatalln("Error on connect to database: ", err.Error())
		return nil, err
	}

	return &conn{session: session, dataBase: session.DB(cfg.DbName())}, nil
}

func (c *conn) Close() {
	fmt.Println("Closing connection...")
	c.session.Close()
}

func (c *conn) DB() *mgo.Database {
	return c.dataBase
}