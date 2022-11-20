package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwt_key = []byte("secret_key")
var db *gorm.DB
var err error

const DNS = "root:root@123@tcp(127.0.0.1:3306)/user_db?charset=utf8&parseTime=True"

func InitialMigration() {
	db, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the database")
	} else {

		fmt.Println("Connected to database successfully")

	}
	db.AutoMigrate(&User{})

}
/*func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}*/
func GetUser(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("token")
	fmt.Println(tokenStr)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return jwt_key, nil })
	fmt.Println(err)
	if err != nil {
		//fmt.Println("inside1")
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			//fmt.Println("inside2")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	if !tkn.Valid {
		//fmt.Println("inside3")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//fmt.Println("inside4")
	//w.Write([]byte(fmt.Sprintf("Hello, %s", claims.User)))
	w.Header().Set("Content-Type", "application/json")
	var user User
	fmt.Println(claims.Email)
	db.Where("email = ?", claims.Email).First(&user)
	json.NewEncoder(w).Encode(&user)

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)

}
/*func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(r)
	db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The user was deleted successfully")
}*/

/*func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(r)
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}*/
func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	var dBuser User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pass1 := user.Password
	email := user.Email

	db.Where("email = ?", email).First(&dBuser)
	//if condition to check if any user exist with the email

	//

	expectedPassword := dBuser.Password
	fmt.Println(pass1, expectedPassword)

	if expectedPassword != pass1 {
		w.WriteHeader(http.StatusUnauthorized)
	}
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	/*http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires : expirationTime,
	})*/
	result := make(map[string]string)

	result["token"] = tokenString
	result["expires"] = expirationTime.String()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}


	

func Refresh(w http.ResponseWriter, r *http.Request) {

	tokenStr := r.Header.Get("token")

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwt_key)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
	&http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	result := make(map[string]string)

	result["refresh_token"] = tokenString
	result["expires"] = expirationTime.String()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
