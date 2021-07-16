package main

// // Init_environment_variables reads the environment variables
// func Init_environment_variables() {
// 	fmt.Println("\n\n\n=================== < INIT  SERVER > ===================")

// 	Env = Environment{}

// 	Env.postgres_user = os.Getenv("POSTGRES_USER")
// 	if len(Env.postgres_user) < 1 {
// 		// If POSTGRES_USER environment is not set
// 		log.Fatalln("init env: ERROR > ", "POSTGRES_USER environment variable is not set")
// 	}
// 	Env.postgres_host = os.Getenv("POSTGRES_HOST")
// 	if len(Env.postgres_host) < 1 {
// 		// If POSTGRES_HOST environment is not set
// 		log.Fatalln("init env: ERROR > ", "POSTGRES_HOST environment variable is not set")
// 	}
// 	Env.postgres_password = os.Getenv("POSTGRES_PASSWORD")
// 	if len(Env.postgres_password) < 1 {
// 		// If POSTGRES_PASSWORD environment is not set
// 		log.Fatalln("init env: ERROR > ", "POSTGRES_PASSWORD environment variable is not set")
// 	}
// 	Env.postgres_db = os.Getenv("POSTGRES_DB")
// 	if len(Env.postgres_db) < 1 {
// 		// If POSTGRES_DB environment is not set
// 		log.Fatalln("init env: ERROR > ", "POSTGRES_DB environment variable is not set")
// 	}
// 	Env.postgres_port = os.Getenv("POSTGRES_PORT")
// 	if len(Env.postgres_port) < 1 {
// 		// If POSTGRES_PORT environment is not set
// 		log.Fatalln("init env: ERROR > ", "POSTGRES_PORT environment variable is not set")
// 	}
// }

// // Init_postgres initializes postgres.
// func Init_postgres() *gorm.DB {
// 	postgres, err := gorm.Open(postgres.Open("host="+Env.postgres_host+" user="+Env.postgres_user+" password="+Env.postgres_password+" dbname="+Env.postgres_db+" port="+Env.postgres_port+" sslmode=disable"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalln("init postgres: ERROR > ", err)
// 	}

// 	Postgres = postgres

// 	if err := Postgres.AutoMigrate(&models.Api_users{}); err != nil {
// 		log.Fatalln("init postgres: ERROR > ", err)
// 	}

// 	if err = Postgres.AutoMigrate(&models.Fizzbuzz_request_statistics{}); err != nil {
// 		log.Fatalln("init postgres: ERROR > ", err)
// 	}

// 	fmt.Println("init postgres: SUCCESS > ", Postgres)

// 	return Postgres
// }

// // Init_redis initializes redis.
// func Init_redis() *redis.Client {
// 	Redis = redis.NewClient(&redis.Options{
// 		Addr:     "redis:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// 	if _, err := Redis.Ping().Result(); err != nil {
// 		log.Fatalln("init redis: ERROR > ", err)
// 	}

// 	fmt.Println("init redis: SUCCESS > ", Redis)

// 	return Redis
// }
