package seeder

import (
	"context"
	"fmt"
	"github.com/jaswdr/faker"
	"inex/main/domain"
	"inex/main/internal/repository"
	"inex/main/pkg/postgres"
	"math/rand"
	"time"
)

func seedCostItems(pg *postgres.Postgres) error {
	inex := repository.New(pg)
	fake := faker.New()

	users, err := inex.ReadAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("seeder - seedCostItems - inex.ReadUser: %w", err)
	}

	for _, user := range users {
		rand.Seed(time.Now().UnixNano())
		numRandom := rand.Intn(10) + 1

		for i := 0; i < numRandom; i++ {
			var costItem domain.CostItem
			costItem.Name = fake.Internet().User()
			err = inex.CreateCostItem(context.Background(), costItem, user)
			if err != nil {
				return fmt.Errorf("seeder - seedIncomeItems - inex.CreateCostItem: %w", err)
			}
		}

	}

	return nil
}
