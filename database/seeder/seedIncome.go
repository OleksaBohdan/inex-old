package seeder

//func seedIncome(pg *postgres.Postgres) error {
//	inex := repository.New(pg)
//	fake := faker.New()
//
//	users, err := inex.ReadAllUsers(context.Background())
//	if err != nil {
//		return fmt.Errorf("seeder - seedIncome - inex.ReadUser: %w", err)
//	}
//
//	for _, user := range users {
//		incomeItems, err := inex.ReadIncomeItems(context.Background(), user)
//		if err != nil {
//			return fmt.Errorf("seeder - seedIncome - inex.ReadIncomeItems: %w", err)
//		}
//
//		for _, incomeItem := range incomeItems {
//
//			rand.Seed(time.Now().UnixNano())
//			numRandom := rand.Intn(10) + 1
//
//			for i := 0; i < numRandom; i++ {
//				var income domain.Income
//				income.Value = fake.Float64(2, 1, 200)
//				err = inex.CreateIncome(context.Background(), income, incomeItem, user)
//				if err != nil {
//					return fmt.Errorf("seeder - seedIncome - inex.CreateIncome: %w", err)
//				}
//			}
//
//		}
//
//	}
//
//	return nil
//}
