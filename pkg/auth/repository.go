package auth

import (
	"nextgen/internals/db/pg"
)

func FindByEmail(email string) (*User, error) {
	var user User
	if err := pg.GDB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) (*User, error) {
	if err := pg.GDB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// func CreateUser(ctx context.Context, user User){
// 	sqlStatement := `
// 	INSERT INTO TBL_USERS(FIRST_NAME, LAST_NAME, USERNAME, EMAIL, PHONE_NUMBER, HASH_PASSWORD)
// 	VALUES ($1, $2, $3, $4, $5, $6)`
//     _, err := pg.Pool.Exec(context.Background(), sqlStatement,user.FirstName, user.LastName, user.Username, user.Email, user.PhoneNumber, user.Password)
//     if err != nil {
//        panic(err)
// 	}

// }

// func GetUserByEmail(ctx context.Context, email string) (*User, error) {
// 	sqlStatement := `SELECT USERNAME,
// 							HASH_PASSWORD
//                        FROM TBL_USERS
//                       WHERE EMAIL = $1`
// 	var user User
// 	row := pg.Pool.QueryRow(context.Background(), sqlStatement, email)
// 	err := row.Scan(&user.Username, &user.Password)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("Error getting user")
// 		return nil, err
// 	}
// 	return &user, nil
// }
