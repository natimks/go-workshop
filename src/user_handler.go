package src

import (
	"github.com/gin-gonic/gin"
	"strconv")

//GetUsers find all users
func GetUsers(c *gin.Context){
	var users []user
	allUsers := selectAll(&users,c)

	c.JSON(200,allUsers)
}

//GetUser find a user by ID
func GetUser(c *gin.Context){
	ID, err := strconv.Atoi(c.Param("id"))
	var selectedUser user
	if err != nil{
		checkErr(err, 400, c)
		return
	}
	selectedUser, db := selectUserID(ID,c)
	var debts [] debt
	db.Model(&selectedUser).Related(&debts)
	selectedUser.Debts = debts
	c.JSON(200,selectedUser)
}

//PostUser create user
func PostUser(c *gin.Context){
	var newUser user
	if err := c.ShouldBindJSON(&newUser); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := DBConnect()
	db.Create(&newUser)

	c.JSON(201, newUser)
}

//PutUser update user by ID
func PutUser(c *gin.Context){
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		checkErr(err, 400, c)
		return
	}

	var updateUser user
	if err := c.ShouldBindJSON(&updateUser); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, db := selectUserID(ID,c)

	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.BirthDate = updateUser.BirthDate

	db.Save(&user)
	c.JSON(200,user)
}

//DeleteUser delete a user by ID
func DeleteUser(c *gin.Context){
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		checkErr(err, 400, c)
		return
	}
	user, db := selectUserID(ID,c)

	if user.ID > 0 {
		db.Delete(&user)
	}
	c.JSON(204, nil)
}

//GetDebtsByUser find debts by ID user
func GetDebtsByUser(c *gin.Context){
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		checkErr(err, 400, c)
		return
	}
	selectedUser, db := selectUserID(ID,c)
	var debts [] debt
	db.Model(&selectedUser).Related(&debts)
	c.JSON(200,debts)
}