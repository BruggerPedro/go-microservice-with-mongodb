package main

var {
	local bool
}

func init() {
	flag.BoolVar(&local, "local", true, "Run server in local mode")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	
	defer conn.Close()
}