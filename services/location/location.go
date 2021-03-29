package location

import (
	"math"

	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/operacion-fuego-quasar/database/models"
	satellite "github.com/lmartinezsch/operacion-fuego-quasar/database/models"
	"github.com/lmartinezsch/operacion-fuego-quasar/services"
)

// ServiceName: nombre sugerido para el servicio
var ServiceName string = "location"

type Service interface {
	services.Service
	GetLocation(distances ...float32) (x, y float32)
}

type locationService struct {
	serviceName  string
	dbConnection *gorm.DB
}

// NewService devuelve la implementación del servicio locationService
func NewService(db *gorm.DB) Service {

	return &locationService{
		dbConnection: db,
	}
}

// Deregister Se realizan las operaciones necesarias para quitar el
// servicio del registry
func (service *locationService) Deregister() {
}

// Register Establece el nombre con el que fue registrado el servicio
func (service *locationService) Register(serviceName string) {
	service.serviceName = serviceName
}

func (service *locationService) GetLocation(distances ...float32) (x, y float32) {

	var r1 float32 = distances[0]
	var r2 float32 = distances[1]
	var r3 float32 = distances[2]

	//get satellites
	satellites := getSatellites(service)

	var a1 float32 = satellites[0].Position.X
	var a2 float32 = satellites[1].Position.X
	var a3 float32 = satellites[2].Position.X

	var b1 float32 = satellites[0].Position.Y
	var b2 float32 = satellites[1].Position.Y
	var b3 float32 = satellites[2].Position.Y

	var (
		r1Sq = r1 * r1
		r2Sq = r2 * r2
		r3Sq = r3 * r3
		a1Sq = a1 * a1
		a2Sq = a2 * a2
		a3Sq = a3 * a3
		b1Sq = b1 * b1
		b2Sq = b2 * b2
		b3Sq = b3 * b3
	)

	// Ax + By = C
	// Dx + Ey = F
	A := (a2 - a1)
	B := (b2 - b1)
	C := (r1Sq - r2Sq - a1Sq + a2Sq - b1Sq + b2Sq) / 2
	D := (a3 - a2)
	E := (b3 - b2)
	F := (r2Sq - r3Sq - a2Sq + a3Sq - b2Sq + b3Sq) / 2

	d := createMatrix(A, B, D, E)
	d1 := createMatrix(C, B, F, E)
	d2 := createMatrix(A, C, D, F)

	det := getDeterminant(d)
	det1 := getDeterminant(d1)
	det2 := getDeterminant(d2)

	if det != 0 {
		val1 := float32(math.Round(float64(det1 / det)))
		val2 := float32(math.Round(float64(det2 / det)))
		return val1, val2
	}
	return 0, 0
}

// Create matrix for determinant
func createMatrix(x1, y1, x2, y2 float32) [2][2]float32 {
	return [2][2]float32{{x1, y1}, {x2, y2}}
}

// Determinant from 2x2 matrix
func getDeterminant(mat [2][2]float32) float32 {
	var ans float32
	ans = mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	return ans
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func getSatellites(service *locationService) []satellite.Satellite {

	db := service.dbConnection
	satellites := []models.Satellite{}
	db.Preload("Position").Find(&satellites)

	return satellites
}
