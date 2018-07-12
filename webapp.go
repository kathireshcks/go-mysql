package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
		
	)
type User struct{
	name string `json:"name"`
	username string `json:"username"`
}

//var db *sql.DB;

var  connect = "kathirgo:12439361@tcp(127.0.0.1:3306)/newdata";

func init() {
	db,err:= sql.Open("mysql",connect)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Database Connected");
		fmt.Println(db);

	}
}



func index_handler(w http.ResponseWriter,r *http.Request){
	
	fmt.Fprintf(w,`wow, Go is neat  Go is fast`)
	

}
func about_handler(w http.ResponseWriter,r *http.Request){
	
	db,err:= sql.Open("mysql",connect)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	results, err := db.Query("SELECT name FROM users")
	if err != nil {
		fmt.Println("error");
		panic(err.Error()) ;
	}else{
		for results.Next(){
			var user User

			err=results.Scan(&user.name);
			if err!=nil{
				panic(err.Error());
			}
			fmt.Println(user.name);
		}

	}
	//fmt.Fprintf(w,string(results));

	
	fmt.Fprintf(w,"about go")
	


}


func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
	//	results = append(results, string(body))
	fmt.Println(string(body));
		//fmt.Fprint(w, body)
		w.Write(body);
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main(){
/*
	db, err := sql.Open("mysql",
		connect)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("database connected");	
	}
	
	results, err := db.Query("SELECT name username FROM users")
	if err != nil {
		fmt.Println("error");
		panic(err.Error()) 
	}else{
		fmt.Println(results);
		for results.Next(){
			var user User

			err=results.Scan(&user.name);
			if err!=nil{
				panic(err.Error());
			}
			fmt.Println(user.name);
		}

	}
*/
	
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about",about_handler)
	http.HandleFunc("/post",PostHandler)
	http.ListenAndServe(":8000",nil)
	fmt.Println("running 8000");
}