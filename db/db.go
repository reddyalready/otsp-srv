package db

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/reddyalready/otsp-srv/model"
)

var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		User:     "otsp",
		Database: "otsp",
	})
	if err := createSchema(); err != nil {
		panic(err)
	}
	if err := createFakeData(); err != nil {
		panic(err)
	}
}

func ProjectList() []model.Project {
	var projects []model.Project
	err := DB.Model(&projects).Select()
	if err != nil {
		panic(err)
	}
	return projects
}

func ProjectGet(id int) model.Project {
	project := model.Project{ID: id}
	err := DB.Select(&project)
	if err != nil {
		panic(err)
	}
	return project
}

func createSchema() error {
	for _, model := range []interface{}{&model.Project{}} {
		err := DB.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func createFakeData() error {
	return DB.Insert(&model.Project{Name: "Hi"})
}
