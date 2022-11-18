package service

import (
	"GinTemplate/models"
	"GinTemplate/service/dbs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var dbColl = "user"

func FindPasswd(c *models.Claims) bool {
	u := models.Claims{}
	err := dbs.Mongo.Find(dbColl, bson.M{"name": c.Name, "password": c.Password, "status": true}).One(&u)
	if err != nil {
		return false
	}
	c.Role = u.Role
	if u.Password == c.Password {
		_ = dbs.Mongo.Update(dbColl, bson.M{"name": c.Name}, bson.M{"$set": bson.M{"lasttime": time.Now().Unix()}})
		return true
	}

	return false
}

func UpdatePasswd(name, pwd string) error {
	return dbs.Mongo.Update(dbColl, bson.M{"name": name}, bson.M{"$set": bson.M{"password": pwd}})
}

func GetUser() []models.User {
	var users []models.User
	if err := dbs.Mongo.Find(dbColl, bson.M{}).Select(bson.M{"lasttime": 0, "password": 0}).All(&users); err != nil {
		return nil
	}
	return users
}

func CheckUser(name string) bool {
	if name == "" {
		return false
	}
	if n, err := dbs.Mongo.Find(dbColl, bson.M{"name": name}).Count(); err != nil || n == 1 {
		return false
	}
	return true
}

func CreateUser(user models.User) bool {
	return dbs.Mongo.Insert(dbColl, user) == nil
}

func UpdateUser(user models.User) bool {
	if user.Name == "" {
		return false
	}
	if user.Password == "" {
		return dbs.Mongo.Update(dbColl, bson.M{"name": user.Name}, bson.M{"$set": bson.M{"role": user.Role}}) == nil
	} else {
		return dbs.Mongo.Update(dbColl, bson.M{"name": user.Name}, bson.M{"$set": bson.M{"password": user.Password, "role": user.Role}}) == nil
	}
}

func DeleteUser(name string) bool {
	if name == "" {
		return false
	}
	return dbs.Mongo.Remove(dbColl, bson.M{"name": name}) == nil
}

func UpdateStatus(name string, status bool) bool {
	if name == "" {
		return false
	}
	return dbs.Mongo.Update(dbColl, bson.M{"name": name}, bson.M{"$set": bson.M{"status": status}}) == nil
}
