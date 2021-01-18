package database // saat package ini di import maka function init() akan langsung di eksekusi
import "fmt"

var connection string

func init() { // langsung di eksekusi apabila package database telah di import
	fmt.Println("function init() di panggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection // maka ini outpunya adalah MySQL
}
