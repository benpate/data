package memory

import (
	"context"
	"testing"
	"time"

	"github.com/benpate/criteria"
	"github.com/benpate/derp"
	"github.com/benpate/remote"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPresto(t *testing.T) {

	db := testDB{}
	factory := testFactory{db: &db}
	filter := criteria.Expression{}

	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.NoContent(200)
	})

	NewCollection(e, &factory, "Persons", "/persons", "personId").
		Post().
		Get().
		Put().
		Patch().
		Delete()

	go e.Start(":8080")

	// Verify that the server is running.
	if err := remote.Get("http://localhost:8080/").Send(); err != nil {
		err.Report()
		assert.Fail(t, "Error getting default route", err)
	}

	///////////////////
	// POST a record
	john := testPerson{
		PersonID: "jc123",
		Name:     "John Connor",
		Email:    "john@sky.net",
	}

	person := testPerson{}

	// Post a record to the "remote" server
	t1 := remote.Post("http://localhost:8080/persons").
		JSON(john)

	if err := t1.Send(); err != nil {
		err.Report()
		assert.Fail(t, "Error posting to localhost", err)
	}

	// Confirm that the record was sent correctly.
	filter = criteria.Expression{{"personId", "=", john.PersonID}}
	if err := db.Load("Persons", filter, &person); err != nil {
		err.Report()
		assert.Fail(t, "Error loading new record from db", err)
	}

	assert.Equal(t, john.PersonID, person.PersonID)
	assert.Equal(t, john.Name, person.Name)
	assert.Equal(t, john.Email, person.Email)

	//////////////////////////
	// PUT a record

	sarah := testPerson{
		PersonID: "sc456",
		Name:     "Sarah Connor",
		Email:    "sarah@sky.net",
	}

	t2 := remote.Put("http://localhost:8080/persons/" + sarah.ID()).
		JSON(sarah)

	if err := t2.Send(); err != nil {
		err.Report()
		assert.Fail(t, "Error PUT-ing a record", sarah)
	}

	// Confirm that the record was sent correctly.
	filter = criteria.Expression{{"personId", "=", sarah.PersonID}}
	if err := db.Load("Persons", filter, &person); err != nil {
		err.Report()
		assert.Fail(t, "Error loading new record", err)
	}

	assert.Equal(t, sarah.PersonID, person.PersonID)
	assert.Equal(t, sarah.Name, person.Name)
	assert.Equal(t, sarah.Email, person.Email)

	////////////////////////
	// GET records

	// Load John
	t3 := remote.Get("http://localhost:8080/persons/"+john.PersonID).
		Response(&person, nil)

	if err := t3.Send(); err != nil {
		err.Report()
		assert.Fail(t, "Error retrieving person from REST service")
	}

	assert.Equal(t, john.PersonID, person.PersonID)
	assert.Equal(t, john.Name, person.Name)
	assert.Equal(t, john.Email, person.Email)

	// Load Sarah
	t4 := remote.Get("http://localhost:8080/persons/"+sarah.PersonID).
		Response(&person, nil)

	if err := t4.Send(); err != nil {
		err.Report()
		assert.Fail(t, "Error retrieving person from REST service")
	}

	assert.Equal(t, sarah.PersonID, person.PersonID)
	assert.Equal(t, sarah.Name, person.Name)
	assert.Equal(t, sarah.Email, person.Email)
}

// MODEL OBJECT

type testPerson struct {
	PersonID   string `json:"personId"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	CreateDate int64  `json:"createDate"`
	UpdateDate int64  `json:"updateDate"`
}

func (person *testPerson) ID() string {
	return person.PersonID
}

func (person *testPerson) IsNew() bool {
	return person.CreateDate == 0
}

func (person *testPerson) ETag() string {
	return ""
}

func (person *testPerson) SetCreated(comment string) {
	person.CreateDate = time.Now().Unix()
}

func (person *testPerson) SetUpdated(comment string) {
	person.UpdateDate = time.Now().Unix()
}

func (person *testPerson) SetDeleted(comment string) {
}

// SERVICE OBJECT
type testPersonService struct {
	session *testDB
}

func (service *testPersonService) NewObject() Object {
	return &testPerson{}
}

func (service *testPersonService) LoadObject(filter criteria.Expression) (Object, *derp.Error) {

	person := service.NewObject()

	if err := service.session.Load("Persons", filter, person); err != nil {
		return nil, derp.Wrap(err, "testPersonService.Load", "Error Loading Person")
	}

	return person, nil
}

func (service *testPersonService) SaveObject(person Object, note string) *derp.Error {

	if err := service.session.Save("Persons", person, note); err != nil {
		return derp.Wrap(err, "testPersonService.Save", "Error Saving Person", person)
	}

	return nil
}

func (service *testPersonService) DeleteObject(person Object, note string) *derp.Error {

	if err := service.session.Delete("Persons", person, note); err != nil {
		return derp.Wrap(err, "testPersonService.Delete", "Error Deleting Person", person)
	}

	return nil
}

func (service *testPersonService) Close() {}

// FACTORY OBJECT
type testFactory struct {
	db *testDB
}

func (factory *testFactory) Service(string) Service {

	ctx := context.TODO()

	return &testPersonService{
		session: factory.db.Session(ctx),
	}
}
