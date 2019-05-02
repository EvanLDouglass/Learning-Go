// Evan Douglass
// Created May 01 2019
// A basic approach to templating html

package main

import "fmt"

func main() {
	name := "Evan Douglass"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="utf-f">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>` + name + `</h1>
	</body>
	</html>
	`
	// Use go run 01-intro.go > index.html in bash to output to file
	fmt.Println(tpl)
}
