package services

import (
	"net/http"

	"github.com/followthepattern/go-simple-webservice/models"
	"github.com/followthepattern/go-simple-webservice/utils"
)

//go:generate mockgen -destination=../mocks/services/database.go -package=mocks -source=./memory.go Database

type Database interface {
	GetAllUser() []models.User

	GetUserByID(userID string) (*models.User, error)

	GetUserCount() int

	AddUser(user models.User)

	UpdateUser(userID string, user models.User) error

	DeleteAllUser() error

	DeleteUser(userID string) error
}

func NewDatabase() Database {
	m := memory{}
	m.initPredefinedData()
	return m
}

type memory map[string]models.User

func (m memory) initPredefinedData() {
	m["07cc24c8-b52f-4bfa-9557-48ccfdb1dff7"] = models.User{
		UserID: "07cc24c8-b52f-4bfa-9557-48ccfdb1dff7",
		Age:    30,
		Name:   "Vernon Marshall",
	}

	m["96dc26a1-6b4d-4808-aa2b-ab9f04317543"] = models.User{
		UserID: "96dc26a1-6b4d-4808-aa2b-ab9f04317543",
		Age:    22,
		Name:   "Grace Bell",
	}

	m["7f011b92-1184-42ec-b256-5207f87f4975"] = models.User{
		UserID: "7f011b92-1184-42ec-b256-5207f87f4975",
		Age:    47,
		Name:   "Julia Tate",
	}

	m["6f43e075-fab0-4e3b-a303-9d34770b1255"] = models.User{
		UserID: "6f43e075-fab0-4e3b-a303-9d34770b1255",
		Age:    30,
		Name:   "Alysha Vic",
	}

	m["93c9591e-5c4c-4a9b-a73b-37bf178867b7"] = models.User{
		UserID: "93c9591e-5c4c-4a9b-a73b-37bf178867b7",
		Age:    23,
		Name:   "Pablo Ballard",
	}

	m["b21667a1-00df-4bff-b227-c6d9b26e4b91"] = models.User{
		UserID: "b21667a1-00df-4bff-b227-c6d9b26e4b91",
		Age:    46,
		Name:   "Kay Gordon",
	}

	m["7a5f7022-3694-4584-822d-e802c62225f6"] = models.User{
		UserID: "7a5f7022-3694-4584-822d-e802c62225f6",
		Age:    47,
		Name:   "Willard Pittman",
	}

	m["a6b56bb8-1653-4686-bd96-56e6385e4d7d"] = models.User{
		UserID: "a6b56bb8-1653-4686-bd96-56e6385e4d7d",
		Age:    45,
		Name:   "Delores Coleman",
	}

	m["892fdcf9-f103-4b3b-a4b3-95f5bd8b50f6"] = models.User{
		UserID: "892fdcf9-f103-4b3b-a4b3-95f5bd8b50f6",
		Age:    48,
		Name:   "Noah Moss",
	}

	m["0849b379-10d2-48c3-855a-5c3a0791b3d8"] = models.User{
		UserID: "0849b379-10d2-48c3-855a-5c3a0791b3d8",
		Age:    19,
		Name:   "Gustavo Goodman",
	}

	m["5c905c86-a20c-4ad4-a014-900ca14adff9"] = models.User{
		UserID: "5c905c86-a20c-4ad4-a014-900ca14adff9",
		Age:    21,
		Name:   "Denise Hunter",
	}

	m["019da86e-ee88-436c-91ad-dee29d4a0aaf"] = models.User{
		UserID: "019da86e-ee88-436c-91ad-dee29d4a0aaf",
		Age:    39,
		Name:   "Kurt Nichols",
	}

	m["e40826b2-3ae2-4852-a53f-9db554f3abe7"] = models.User{
		UserID: "e40826b2-3ae2-4852-a53f-9db554f3abe7",
		Age:    20,
		Name:   "Chester Allen",
	}

	m["53cee410-60c3-4c2b-9119-2cbe695cb0fa"] = models.User{
		UserID: "53cee410-60c3-4c2b-9119-2cbe695cb0fa",
		Age:    27,
		Name:   "Andres Quinn",
	}

	m["1b9dbf38-5747-4acb-9e58-217f2923d24f"] = models.User{
		UserID: "1b9dbf38-5747-4acb-9e58-217f2923d24f",
		Age:    42,
		Name:   "Lana Gross",
	}

	m["704139dd-7be4-4c59-89e7-7bba16e9847e"] = models.User{
		UserID: "704139dd-7be4-4c59-89e7-7bba16e9847e",
		Age:    50,
		Name:   "Cecelia Long",
	}

	m["883708d4-f8b2-466a-8b85-05fa9ee07e19"] = models.User{
		UserID: "883708d4-f8b2-466a-8b85-05fa9ee07e19",
		Age:    32,
		Name:   "Mable Pearson",
	}

	m["03ea7ba4-fb84-4c08-85b9-bdffe200b182"] = models.User{
		UserID: "03ea7ba4-fb84-4c08-85b9-bdffe200b182",
		Age:    24,
		Name:   "Max Maxwell",
	}
}

func (m memory) GetAllUser() (result []models.User) {
	result = make([]models.User, len(m))
	i := 0
	for _, v := range m {
		result[i] = v
		i++
	}
	return
}

func (m memory) GetUserCount() int {
	return len(m)
}

func (m memory) GetUserByID(userID string) (result *models.User, err error) {
	if _, ok := m[userID]; ok {
		user := m[userID]
		return &user, nil
	} else {
		return nil, utils.New(http.StatusNotFound, "Given userID doesn't exist!")
	}
}

func (m memory) AddUser(user models.User) {
	if _, ok := m[user.UserID]; !ok {
		m[user.UserID] = user
	}
}

func (m memory) UpdateUser(userID string, user models.User) error {
	if _, ok := m[userID]; ok {
		user.UserID = userID
		m[userID] = user
	} else {
		return utils.New(http.StatusNotFound, "Given userID doesn't exist!")
	}
	return nil
}

func (m memory) DeleteUser(userID string) error {
	if _, ok := m[userID]; ok {
		delete(m, userID)
	} else {
		return utils.New(http.StatusNotFound, "Given userID doesn't exist!")
	}
	return nil
}

func (m memory) DeleteAllUser() error {
	m = make(map[string]models.User)
	return nil
}
