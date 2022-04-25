package main

// import (
// 	"fmt"
// 	"time"
// )

// type myerr struct {
// 	when time.Time
// 	what string
// }

// func (e *myerr) Error() string {
// 	return fmt.Sprintf("at %v, %s", e.when, e.what)
// }
// func run() error {
// 	return &myerr{
// 		time.Now(),
// 		"it not work",
// 	}
// }
// func main() {

// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }
