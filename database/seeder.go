package database

import (
	"fmt"
	"go-template/entity"
	"os"
)

func runSeeder() {
	if os.Getenv("ENV") == "dev" {
		dropTable()
		createTable()
		seeding()
	}
}

func dropTable() {
	sql := `
		drop table if exists users;
	`

	err := db.Exec(sql).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully delete tables")
}

func createTable() {
	err := db.AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully migrate tables")
}

func seeding() {
	err := db.Exec(`
		insert into
			users(name, email, password, created_at)
		values
			('name1', 'name1@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name2', 'name2@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name3', 'name3@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name4', 'name4@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name5', 'name5@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name6', 'name6@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name7', 'name7@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name8', 'name8@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name9', 'name9@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01'),
			('name10', 'name10@mail.com',  '$2a$05$JBpQrAG3FltyVi1J3yClCOfn/41kawtx1LLn1k0MmhougJhdVFyCK', '2023-01-01');
	`).Error

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully seed tables")
}
