package main

import (
	"flag"
	"fmt"
	"restArchitecture/mikail/App/connections/auth"
	"restArchitecture/mikail/App/controllers"
)

func main() {
	env := flag.String("env", "", "")
	flag.Parse()
	ctr := controllers.NewControllers(*env)
	authHandlers := auth.NewAuthHandlers(ctr)
	err := auth.Run(authHandlers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ctr)
}
