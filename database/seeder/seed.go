package seeder

import (
	"fmt"
	"inex/main/config"
	"inex/main/pkg/logger"
	"inex/main/pkg/postgres"
)

func SeedDatabase(cfg *config.Config) {
	fmt.Println("Start seeding database")
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("seeder - seedDatabase - postgres.New: %w", err))

	}
	defer pg.Close()

	err = seedUsers(pg)
	if err != nil {
		l.Fatal(err)
	}

	//err = seedNotes(pg)
	//if err != nil {
	//	l.Fatal(err)
	//}

	err = seedIncomeItems(pg)
	if err != nil {
		l.Fatal(err)
	}

	err = seedCostItems(pg)
	if err != nil {
		l.Fatal(err)
	}

	err = seedIncome(pg)
	if err != nil {
		l.Fatal(err)
	}

	err = seedCost(pg)
	if err != nil {
		l.Fatal(err)
	}

}
