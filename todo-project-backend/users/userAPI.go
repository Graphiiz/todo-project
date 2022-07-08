package users

import (
	"encoding/binary"
	"fmt"
	"net/http"
	db "todo-project-backend/db"

	"github.com/labstack/echo/v4"
)

// ex http://localhost:1234/users/5
func GetUserLogin(c echo.Context) error {
	// link with http request from frontend -> get value from request -> connect to db
	email := c.Param("email")

	userDB, err := db.ReadEmail(email)
	if err != nil {
		fmt.Println(err)
		// not found
		return c.String(http.StatusNotFound, "Not found")
	}

	// convert []byte to []string
	TodoString := Decode(userDB.Todos)
	user := User{
		// Id:    userDB.Code,
		// Name:  userDB.Name,
		Email: userDB.Email,
		Todos: TodoString,
	}
	fmt.Println("GetUserLogin Success!!")
	return c.JSON(http.StatusOK, user)
}

func GetUserTodos(c echo.Context) error {
	// link with http request from frontend -> get value from request -> connect to db
	email := c.Param("email")

	userDB, err := db.ReadEmail(email)
	if err != nil {
		// not found
		return c.String(http.StatusNotFound, "Not found")
	}

	// convert []byte to []string
	TodoString := Decode(userDB.Todos)
	user := User{
		// Id:    userDB.Code,
		// Name:  userDB.Name,
		Email: userDB.Email,
		Todos: TodoString,
	}
	fmt.Println("GetUserTodo Success!!")
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {

	user := User{}
	if err := c.Bind(&user); err != nil {
		// &user => user variable outside (line 23) if statement is modified
		// user in line 29 has value from c already
		return err
	} // convert json to object user

	// write to db

	// convert []string to []byte
	TodoByte := Encode(user.Todos)

	userDB := db.UserDB{
		// Code:  user.Id,
		// Name:  user.Name,
		Email: user.Email,
		Todos: TodoByte,
	}

	if err := db.Create(userDB); err != nil {
		return c.String(http.StatusExpectationFailed, "Create Fail")
	}

	fmt.Println("CreateUser Success!!")
	return c.JSON(http.StatusCreated, user)
}

func Save(c echo.Context) error {

	user := User{}
	if err := c.Bind(&user); err != nil {
		// &user => user variable outside (line 23) if statement is modified
		// user in line 29 has value from c already
		return err
	} // convert json to object user

	// write to db

	// convert []string to []byte
	TodoByte := Encode(user.Todos)
	userDB := db.UserDB{
		// Code:  user.Id,
		// Name:  user.Name,
		Email: user.Email,
		Todos: TodoByte,
	}

	if err := db.Save(userDB); err != nil {
		return c.String(http.StatusExpectationFailed, "Create Fail")
	}

	fmt.Println("Save Success!!")
	return c.JSON(http.StatusCreated, user)
}

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// for []string <-> []byte
const maxInt32 = 1<<(32-1) - 1

func writeLen(b []byte, l int) []byte {
	if 0 > l || l > maxInt32 {
		panic("writeLen: invalid length")
	}
	var lb [4]byte
	binary.BigEndian.PutUint32(lb[:], uint32(l))
	return append(b, lb[:]...)
}

func readLen(b []byte) ([]byte, int) {
	if len(b) < 4 {
		panic("readLen: invalid length")
	}
	l := binary.BigEndian.Uint32(b)
	if l > maxInt32 {
		panic("readLen: invalid length")
	}
	return b[4:], int(l)
}

func Decode(b []byte) []string {
	b, ls := readLen(b)
	s := make([]string, ls)
	for i := range s {
		b, ls = readLen(b)
		s[i] = string(b[:ls])
		b = b[ls:]
	}
	return s
}

func Encode(s []string) []byte {
	var b []byte
	b = writeLen(b, len(s))
	for _, ss := range s {
		b = writeLen(b, len(ss))
		b = append(b, ss...)
	}
	return b
}

// func codecEqual(s []string) bool {
// 	return fmt.Sprint(s) == fmt.Sprint(Decode(Encode(s)))
// }
