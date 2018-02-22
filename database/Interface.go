package database

import "gopkg.in/mgo.v2/bson"

//Crud interface is used to implement the CRUD principle on a database
type Crud interface {
	Create()
	Read() []Device
	Update()
	Delete()
	Validate() (bool, error)
}

///TESTING PULLING A BRANCH
///TESTING A SECOND PULL

//Device struct is used as a return vaule for necessary Database device converstions.
type Device struct {
	Hostname   string        `bson:"hostname,omitempty"`
	IPAddress  string        `bson:"ipaddress,omitempty"`
	DeviceType string        `bson:"devicetype,omitempty"`
	ID         bson.ObjectId `bson:"_id,omitempty"`
}

//StorageCreate  takes an argument of type Crud and runs the Create() method.
func StorageCreate(i Crud) {
	i.Create()
}

//StorageRead  takes an argument of type Crud and runs the Read() method.
func StorageRead(i Crud) []Device {
	return i.Read()
}

//StorageUpdate  takes an argument of type Crud and runs the Update() method.
func StorageUpdate(i Crud) {
	i.Update()
}

//StorageDelete  takes an argument of type Crud and runs the Delete() method.
func StorageDelete(i Crud) {
	i.Delete()
}

//StorageValidate  takes an argument of type Crud and runs the Delete() method.
func StorageValidate(i Crud) (bool, error) {
	return i.Validate()
}
