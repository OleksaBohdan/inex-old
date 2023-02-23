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

func seedCost(pg *postgres.Postgres) error {
	inex := repository.New(pg)
	fake := faker.New()

	users, err := inex.ReadAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("seeder - seedCost - inex.ReadUser: %w", err)
	}

	for _, user := range users {
		costItems, err := inex.ReadCostItems(context.Background(), user)
		if err != nil {
			return fmt.Errorf("seeder - seedCost - inex.ReadIncomeItems: %w", err)
		}

		for _, costItem := range costItems {

			rand.Seed(time.Now().UnixNano())
			numRandom := rand.Intn(10) + 1

			for i := 0; i < numRandom; i++ {
				var cost domain.Cost
				cost.Value = fake.Float64(2, 1, 200)
				err = inex.CreateCost(context.Background(), cost, costItem, user)
				if err != nil {
					return fmt.Errorf("seeder - seedCost - inex.CreateCost: %w", err)
				}
			}

		}

	}

	return nil
}
