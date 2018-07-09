func main() {
	db, err := sql.Open("mysql", "root:root123@/test")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id int
		name string
		email string
		mobile string
		)
	result,err := db.Query("select * from lol")
}