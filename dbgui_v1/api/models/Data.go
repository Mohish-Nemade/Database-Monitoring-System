package models

import (
	"errors"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Rdsdata struct {
	ID                    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Instanceidentifier    string `gorm:"size:255;not null" json:"instanceidentifier"`
	Createdon             string `gorm:"size:255;not null" json:"createdon"`
	Storageengine         string `gorm:"size:255;not null" json:"storageengine"`
	Pgname                string `gorm:"size:255;not null" json:"pgname"`
	Storageallocated      string `gorm:"size:255;not null" json:"storageallocated"`
	BackupRetentionPeriod string `gorm:"size:255;not null" json:"BackupRetentionPeriod"`
	DBName                string `gorm:"size:255;not null" json:"DBName"`
	Engineversion         string `gorm:"size:255;not null" json:"engineversion"`
	RDSzone               string `gorm:"size:255;not null" json:"RDSzone"`
	Volumetype            string `gorm:"size:255;not null" json:"volumetype"`
	Storageencrypted      string `gorm:"size:255;not null" json:"storageencrypted"`
	RDSclass              string `gorm:"size:255;not null" json:"RDSclass"`
	Monitored             uint32 `gorm:"not null" json:"monitored"`
}

type Pmmdata struct {
	ID                    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Instanceidentifier    string `gorm:"size:255;not null" json:"instanceidentifier"`
	Createdon             string `gorm:"size:255;not null" json:"createdon"`
	Storageengine         string `gorm:"size:255;not null" json:"storageengine"`
	Pgname                string `gorm:"size:255;not null" json:"pgname"`
	Storageallocated      string `gorm:"size:255;not null" json:"storageallocated"`
	BackupRetentionPeriod string `gorm:"size:255;not null" json:"BackupRetentionPeriod"`
	DBName                string `gorm:"size:255;not null" json:"DBName"`
	Engineversion         string `gorm:"size:255;not null" json:"engineversion"`
	RDSzone               string `gorm:"size:255;not null" json:"RDSzone"`
	Volumetype            string `gorm:"size:255;not null" json:"volumetype"`
	Storageencrypted      string `gorm:"size:255;not null" json:"storageencrypted"`
	RDSclass              string `gorm:"size:255;not null" json:"RDSclass"`
	Monitored             uint32 `gorm:"not null" json:"monitored"`
}

type Tempdata struct {
	ID                    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Instanceidentifier    string `gorm:"size:255;not null" json:"instanceidentifier"`
	Createdon             string `gorm:"size:255;not null" json:"createdon"`
	Storageengine         string `gorm:"size:255;not null" json:"storageengine"`
	Pgname                string `gorm:"size:255;not null" json:"pgname"`
	Storageallocated      string `gorm:"size:255;not null" json:"storageallocated"`
	BackupRetentionPeriod string `gorm:"size:255;not null" json:"BackupRetentionPeriod"`
	DBName                string `gorm:"size:255;not null" json:"DBName"`
	Engineversion         string `gorm:"size:255;not null" json:"engineversion"`
	RDSzone               string `gorm:"size:255;not null" json:"RDSzone"`
	Volumetype            string `gorm:"size:255;not null" json:"volumetype"`
	Storageencrypted      string `gorm:"size:255;not null" json:"storageencrypted"`
	RDSclass              string `gorm:"size:255;not null" json:"RDSclass"`
	Monitored             uint32 `gorm:"not null" json:"monitored"`
}

func VerifyPasswordData(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *Rdsdata) PrepareData() {
	u.ID = 0
	/*	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
		u.Email = html.EscapeString(strings.TrimSpace(u.Email))
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()

	*/

}

func (u *User) ValidateData(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *Rdsdata) SaveData(db *gorm.DB) (*Rdsdata, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Rdsdata{}, err
	}
	return u, nil
}

func (u *Rdsdata) PullRdsData(db *gorm.DB) (*Rdsdata, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Rdsdata{}, err
	}
	return u, nil
}

//func (u *Rdsdata) compare(db *gorm.DB) (*Rdsdata, error) {

func (u *Rdsdata) FindAllData(db *gorm.DB) (*[]Rdsdata, error) {
	var err error
	datas := []Rdsdata{}
	err = db.Debug().Model(&Rdsdata{}).Limit(100).Find(&datas).Error
	if err != nil {
		return &[]Rdsdata{}, err
	}
	return &datas, err
}

func (u *Pmmdata) FindAllPmmData(db *gorm.DB) (*[]Pmmdata, error) {
	var err error
	datas := []Pmmdata{}
	err = db.Debug().Model(&Pmmdata{}).Limit(100).Find(&datas).Error
	if err != nil {
		return &[]Pmmdata{}, err
	}
	return &datas, err
}
func (u *Tempdata) SimplyGetTempData(db *gorm.DB) (*[]Tempdata, error) {
	var err error
	datas := []Tempdata{}
	err = db.Debug().Model(&Tempdata{}).Limit(100).Find(&datas).Error
	if err != nil {
		return &[]Tempdata{}, err
	}
	return &datas, err

}
func (u *Tempdata) FindAllTempData(db *gorm.DB) (*[]Tempdata, error) {

	var err error
	datas := []Rdsdata{}
	err = db.Debug().Model(&Rdsdata{}).Limit(100).Where("monitored = ?", 0).Find(&datas).Error
	if err != nil {
		return &[]Tempdata{}, err
	}
	/*	rows, err := db.Model(&Rdsdata{}).Where("monitored = ?", 0).Select("instanceidentifier,createdon,storageengine,pgname,storageallocated,BackupRetentionPeriod,DBName,engineversion,RDSzone,volumetype,storageencrypted,RDSclasss").Rows() // (*sql.Rows, error)
		defer rows.Close()
		for rows.Next() {

			//rows.Scan(&name, &age, &email)
			err = db.Debug().Model(&Tempdata{}).Create(PrepareData(datas[i])).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)

			}
		}*/

	err = db.Debug().DropTableIfExists(&Tempdata{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&Tempdata{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	tempvar := []Tempdata{}

	for i := range datas {

		tempvar = append(tempvar, Tempdata(datas[i]))
		tempvar[i].ID = 0
		err = db.Debug().Model(&Tempdata{}).Create(&tempvar[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)

		}
	}

	//	tempvar := []Tempdata{Tempdata(datas)}
	//	tempvar = (datas)

	return &tempvar, err
}

func (u *Rdsdata) FindUserByID(db *gorm.DB, uid uint32) (*Rdsdata, error) {
	var err error
	err = db.Debug().Model(Rdsdata{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Rdsdata{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Rdsdata{}, errors.New("User Not Found")
	}
	return u, err
}

/*
func (u *Rdsdata) UpdateAUser(db *gorm.DB, uid uint32) (*Rdsdata, error) {

	db = db.Debug().Model(&Rdsdata{}).Where("id = ?", uid).Take(&Rdsdata{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Rdsdata{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Rdsdata{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Rdsdata{}, err
	}
	return u, nil
}

// Delete data
func (u *Rdsdata) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Rdsdata{}).Where("id = ?", uid).Take(&Rdsdata{}).Delete(&Rdsdata{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
*/
