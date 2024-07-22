package main

import auth "chatapp/server/routes"
import _ "chatapp/server/docs"

func main() {
	auth.CreateRoute()
}
