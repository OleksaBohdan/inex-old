package seeder

//func seedIncomeItems(pg *postgres.Postgres) error {
//	inex := repository.New(pg)
//	fake := faker.New()
//
//	users, err := inex.ReadAllUsers(context.Background())
//	if err != nil {
//		return fmt.Errorf("seeder - seedIncomeItems - inex.ReadUser: %w", err)
//	}
//
//	for _, user := range users {
//		rand.Seed(time.Now().UnixNano())
//		numRandom := rand.Intn(10) + 1
//
//		for i := 0; i < numRandom; i++ {
//			var incomeItem domain.IncomeItem
//			incomeItem.Name = fake.Internet().User()
//			err = inex.CreateIncomeItem(context.Background(), incomeItem, user)
//			if err != nil {
//				return fmt.Errorf("seeder - seedIncomeItems - inex.CreateIncomeItem: %w", err)
//			}
//		}
//
//	}
//
//	return nil
//}

// delete income item
//var incomeItem domain.IncomeItem
//id, _ := uuid.FromString("0e29797f-33bd-4391-9e5c-1ee6c8330f9e")
//incomeItem.ID = id
//inex := repository.New(pg)
//err = inex.DeleteIncomeItem(context.Background(), incomeItem)
