package main

import (
	"log"
	"time"

	"github.com/Ujjawal0619/Go-gRPC/catalog"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/kit/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listning on PORT 8080 ...")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, 8080))
}
