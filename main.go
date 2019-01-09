package main

import "./lib/swagger"

func main() {
	swagger.HelloWorld()

	swagger.ReadFile("asset/apartment-api.yaml")
}
