package db

type Config interface {
	Dns() string
	DbName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn string
}

func NewConfig() Config {
	var cfr Config
	var err error

	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbPort = os.Getenv("DB_PORT")
	cfg.dbName = os.Getenv("DB_NAME")
	
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var: ", err.Error())
	}

	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)
	return &cfg
}

func (c *) Dns string {
	return c.dsn
}

func (c *) DbName string {
	return c.dbName
}