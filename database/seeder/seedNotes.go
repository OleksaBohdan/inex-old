package seeder

// FIX CREATE NOTE

//func seedNotes(pg *postgres.Postgres) error {
//	inex := repository.New(pg)
//	fake := faker.New()
//
//	users, err := inex.ReadAllUsers(context.Background())
//	if err != nil {
//		return fmt.Errorf("seeder - seedNotes - inex.ReadUser: %w", err)
//	}
//
//	for _, user := range users {
//		var note domain.Note
//		note.Text = fake.Lorem().Sentence(15)
//		err = inex.CreateNote(context.Background(), note)
//		if err != nil {
//			return fmt.Errorf("seeder - seedNotes - inex.CreateNote: %w", err)
//		}
//	}
//
//	return nil
//}

//var note domain.Note
//note.Text = "Lorem lorem lorem text updated"

//err = inex.CreateNote(context.Background(), note, user)
//if err != nil {
//	l.Fatal(fmt.Errorf("app - Rux - inex.Delete: %w", err))
//}

//err = inex.UpdateNote(context.Background(), note, user)
//if err != nil {
//	l.Fatal(fmt.Errorf("app - Rux - inex.UpdateNote: %w", err))
//}

// read Note
//n, err := inex.ReadNote(context.Background(), user)
//if err != nil {
//	l.Fatal(fmt.Errorf("app - Rux - inex.ReadNote: %w", err))
//}
//fmt.Println("note:", n.Text)

//err = inex.DeleteNote(context.Background(), user)
//if err != nil {
//	l.Fatal(fmt.Errorf("app - Rux - inex.DeleteNote: %w", err))
//}
