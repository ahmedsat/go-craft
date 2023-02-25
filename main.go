package main

import gocraft "github.com/ahmedsat/go-craft/go-craft"

func main() {
	test()
}

func test() {
	g := gocraft.NewGame()
	g.Run()
}
