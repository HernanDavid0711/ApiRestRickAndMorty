package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"prubarickmorti/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func PostgresConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return db, err
}

func MakeDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("DB"), os.Getenv("PORT"))
}

func LlenarBD() error {
	dbConn, err := PostgresConnection()
	if err != nil {
		return err
	}

	err = getAllCharacter(dbConn)
	if err != nil {
		return err
	}

	err = getAllepisodes(dbConn)
	if err != nil {
		return err
	}

	return nil
}

func CharactersLista() ([]models.Character, error) {
	dbConn, err := PostgresConnection()
	if err != nil {
		return nil, err
	}

	var data []models.Character

	result := dbConn.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func Characters(id int) (models.Character, error) {
	var data models.Character

	dbConn, err := PostgresConnection()
	if err != nil {
		return data, err
	}

	result := dbConn.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

func SyncCharacters() error {
	dbConn, err := PostgresConnection()
	if err != nil {
		return err
	}

	dbConn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Character{})

	err = getAllCharacter(dbConn)
	if err != nil {
		return err
	}

	return nil
}

func SyncEpisodes() error {
	dbConn, err := PostgresConnection()
	if err != nil {
		return err
	}

	dbConn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Episode{})

	err = getAllepisodes(dbConn)
	if err != nil {
		return err
	}

	return nil
}

func DelCharacters() error {
	dbConn, err := PostgresConnection()
	if err != nil {
		return err
	}

	dbConn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Character{})

	return nil
}

func getAllCharacter(dbConn *gorm.DB) error {
	response, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		return err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseObject models.ResponseCharacters
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return err
	}

	for i := 0; i < len(responseObject.Results); i++ {
		current := models.Character{
			Id:           responseObject.Results[i].ID,
			Name:         responseObject.Results[i].Name,
			Status:       responseObject.Results[i].Status,
			Species:      responseObject.Results[i].Species,
			Type:         responseObject.Results[i].Type,
			Gender:       responseObject.Results[i].Gender,
			OriginName:   responseObject.Results[i].Origin.Name,
			LocationName: responseObject.Results[i].Location.Name,
			ImageUrl:     responseObject.Results[i].Image,
		}

		result := dbConn.Create(&current)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func getAllepisodes(dbConn *gorm.DB) error {
	response, err := http.Get("https://rickandmortyapi.com/api/episode")
	if err != nil {
		return err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseObject models.ResponseEpisodes
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return err
	}

	for i := 0; i < len(responseObject.Results); i++ {
		current := models.Episode{
			Id:          responseObject.Results[i].ID,
			Name:        responseObject.Results[i].Name,
			EpisodeCode: responseObject.Results[i].Episode,
		}

		result := dbConn.Create(&current)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
