package src

//AutoMigration is the func for update the database schema
func AutoMigration(){
	db := DBConnect()
	defer db.Close()

	db.AutoMigrate(user{}, debt{})
	db.Model(&debt{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}