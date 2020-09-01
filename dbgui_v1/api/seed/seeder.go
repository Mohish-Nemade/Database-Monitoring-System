package seed

import (
	"bitbucket.org/mohishnemade/dbgui_v1/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Mohish Nemade",
		Email:    "mohish@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Tanish ",
		Email:    "tanish@gmail.com",
		Password: "Tanish@123",
	},
}

var rdsdatas = []models.Rdsdata{
	models.Rdsdata{
		Instanceidentifier:    "2019-09-20T13",
		Createdon:             "sqlserver-ex",
		Storageengine:         "default.sqlserver-ex-14.0",
		Pgname:                "20",
		Storageallocated:      "7",
		BackupRetentionPeriod: "14.00.3049.1.v1",
		DBName:                "ap-southeast-1b",
		Engineversion:         "gp2",
		RDSzone:               "false",
		Volumetype:            "db.t2.micro",
		Storageencrypted:      "criffdb",
		RDSclass:              "mysql",
		Monitored:             1,
	},
}

//"mysql","stg-d",ops-spring-ms80",400,1,"dev_dbops_db01","5.7.25","ap-southeast-1b","gp2",true,"db.t2.medium","dev-dbops-db01"-"2019-11-23T08

//"2019-09-20T13","sqlserver-ex","default.sqlserver-ex-14.0",20,7,"14.00.3049.1.v1","ap-southeast-1b","gp2",false,"db.t2.micro","criffdb"-"2018-05-25T06
//"mysql","stg-d",ops-spring-ms80",400,1,"dev_dbops_db01","5.7.25","ap-southeast-1b","gp2",true,"db.t2.medium","dev-dbops-db01"-"2019-11-23T08

func Load(db *gorm.DB) {

	/*	err := db.Debug().DropTableIfExists(&models.User{}, &models.Rdsdata{}).Error
			if err != nil {
				log.Fatalf("cannot drop table: %v", err)
			}


			err = db.Debug().AutoMigrate(&models.User{}).Error
			if err != nil {
				log.Fatalf("cannot migrate table: %v", err)
			}

			err = db.Debug().AutoMigrate(&models.Rdsdata{}).Error
			if err != nil {
				log.Fatalf("cannot migrate table: %v", err)
			}


		err := db.Debug().AutoMigrate(&models.Pmmdata{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}

		err = db.Debug().AutoMigrate(&models.Tempdata{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}*/
	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	/*	for i, _ := range users {
				err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
				if err != nil {
					log.Fatalf("cannot seed users table: %v", err)
				}

			}

		for i, _ := range rdsdatas {
			err = db.Debug().Model(&models.Rdsdata{}).Create(&rdsdatas[i]).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)
			}

		}
	*/
}
