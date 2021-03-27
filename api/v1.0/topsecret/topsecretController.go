package topsecret

import (
	"encoding/json"
	"sort"
	"strings"

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

	// Bind Request
	if err := c.BindJSON(&r); err != nil {
		log.Error("El request enviado no es correcto")
		c.AbortWithStatus(404)
		return
	}

	// Save Satellites
	for _, satellite := range r.Satellites {
		saveSatelliteContact(c, satellite, db)
	}

	response := getTopSecretResponse(c, db)

	c.JSON(200, response)
}

// CreatTopSecretSplit godoc
// @Summary Create Top Secret Split
// @Description Create SatelliteContact
// @Accept  json
// @Produce  json
// @Success 201 {object} TopSecretResponse
// @Failure 404
// @Failure 500
// @Failure default
// @Router /topsecret_split/:name [post]
func createTopSecretSplit(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var r SatelliteRequest
	name := c.Param("id")
	// Add name to request
	r.Name = strings.ToLower(name)

	if err := c.BindJSON(&r); err != nil {
		log.Error("El request enviado no es correcto")
		c.AbortWithStatus(404)
		return
	}

	saveSatelliteContact(c, r, db)

	c.JSON(201, "")
}

// GetTopSecretSplit godoc
// @Summary Get Top Secret Split
// @Description get position and message of ship
// @Accept  json
// @Produce  json
// @Success 200 {object} TopSecretResponse
// @Header 200 {string} Token "qwerty"
// @Failure 404
// @Failure 500
// @Failure default
// @Router /topsecret_split/:name [post]
func getTopSecretSplit(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	response := getTopSecretResponse(c, db)

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
		log.Error("No se encontró el satelite: " + satelliteRequest.Name)
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
		log.Error("No se pudo convertir a json el mensaje")
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

func getSatellitesContacts(c *gin.Context, db *gorm.DB) []models.SatelliteContact {

	var satelliteContacts []models.SatelliteContact

	//Find all satellites
	if err := db.Find(&satelliteContacts).Error; err != nil {
		c.AbortWithStatus(500)
		return nil
	}

	return satelliteContacts
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

func getTopSecretResponse(c *gin.Context, db *gorm.DB) *TopSecretResponse {

	var response TopSecretResponse
	satellites := getSatellitesContacts(c, db)

	if len(satellites) < 3 {
		log.Error("Deben existir 3 satelites por lo menos")
		c.AbortWithStatus(404)
		return nil
	}

	x, y := locationService.GetLocation(satellites[0].Distance, satellites[1].Distance, satellites[2].Distance)

	//TODO: se debe pasar de otra forma los parámetros por si llegan a ser más salites
	message := messageService.GetMessage(convertMessageString(satellites[0].Message), convertMessageString(satellites[1].Message), convertMessageString(satellites[2].Message))

	response.Position.X = x
	response.Position.Y = y
	response.Message = message

	return &response
}
