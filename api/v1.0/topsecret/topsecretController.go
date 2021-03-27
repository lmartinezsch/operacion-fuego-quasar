package topsecret

import (
	"encoding/json"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/operacion-fuego-quasar/database/models"
	log "github.com/lmartinezsch/operacion-fuego-quasar/logger"
)

// TopSecret godoc
// @Summary Get Top Secret
// @Description get position and message of ship
// @Accept  json
// @Produce  json
// @Success 200 {object} TopSecretResponse
// @Header 200 {string} Token "qwerty"
// @Failure 404
// @Failure 500
// @Failure default
// @Router /topsecret [post]
func topSecret(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var r RequestBody
	var response TopSecretResponse

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatus(404)
		return
	}

	for _, satelliteRequest := range r.Satellites {
		saveSatelliteContact(c, satelliteRequest, db)
	}

	x, y := locationService.GetLocation(r.Satellites[0].Distance, r.Satellites[1].Distance, r.Satellites[2].Distance)

	//TODO: se debe pasar de otra forma los parámetros por si llegan a ser más salites
	message := messageService.GetMessage(r.Satellites[0].Message, r.Satellites[1].Message, r.Satellites[2].Message)

	response.Position.X = x
	response.Position.Y = y
	response.Message = message

	c.JSON(200, response)
}

// TopSecret godoc
// @Summary Get Top Secret
// @Description get position and message of ship
// @Accept  json
// @Produce  json
// @Success 200 {object} TopSecretResponse
// @Header 200 {string} Token "qwerty"
// @Failure 404
// @Failure 500
// @Failure default
// @Router /topsecret_split/:name [post]
func topSecretSplit(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var r SatelliteRequest
	var response TopSecretResponse
	var err error
	name := c.Param("id")
	// Add name to request
	r.Name = name

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatus(404)
		return
	}

	saveSatelliteContact(c, r, db)
	satellites, err := getSatellitesContacts(c, db)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	if len(satellites) < 3 {
		log.Error("Deben existir 3 satelites por lo menos")
		c.AbortWithStatus(404)
		return
	}

	x, y := locationService.GetLocation(satellites[0].Distance, satellites[1].Distance, satellites[2].Distance)

	//TODO: se debe pasar de otra forma los parámetros por si llegan a ser más salites
	message := messageService.GetMessage(convertMessageString(satellites[0].Message), convertMessageString(satellites[1].Message), convertMessageString(satellites[2].Message))

	response.Position.X = x
	response.Position.Y = y
	response.Message = message

	c.JSON(200, response)
}

type MessageStruct struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func saveSatelliteContact(c *gin.Context, satelliteRequest SatelliteRequest, db *gorm.DB) {

	Satellite := models.Satellite{}

	// Find Satellite by name
	if err := db.Set("gorm:auto_preload", true).Where("name = ?", satelliteRequest.Name).First(&Satellite).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	// Convert message in json
	datas := make(map[int]string)
	for i, fragment := range satelliteRequest.Message {
		datas[i] = fragment
	}
	jsonString, err := json.Marshal(datas)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	// Create SatelliteContact
	satelliteContact := SatelliteContact{
		Distance:  satelliteRequest.Distance,
		Message:   string(jsonString),
		Satellite: Satellite,
	}

	// Update SatelliteContact
	if db.Model(&satelliteContact).Where("satellite_id = ?", Satellite.ID).Updates(&satelliteContact).RowsAffected == 0 {
		// Create SatelliteContact if it was not found
		db.NewRecord(satelliteContact)
		db.Create(&satelliteContact)
	}
}

func getSatellitesContacts(c *gin.Context, db *gorm.DB) ([]models.SatelliteContact, error) {

	var satelliteContacts []models.SatelliteContact

	//Find all satellites
	if err := db.Find(&satelliteContacts).Error; err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}

	return satelliteContacts, nil
}

func convertMessageString(messageString string) []string {
	var messageMap map[int]string

	if err := json.Unmarshal([]byte(messageString), &messageMap); err != nil {
		panic(err)
	}

	// To store the keys in slice in sorted order
	keys := make([]int, len(messageMap))
	i := 0
	for k := range messageMap {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	// To perform the opertion you want
	var messageArray []string
	for _, k := range keys {
		messageArray = append(messageArray, messageMap[k])
	}

	return messageArray
}
