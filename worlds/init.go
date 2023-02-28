package worlds

import "github.com/ahmedsat/go-craft/app"

// Sets the category and name of your test in the demos.Map
// The category name choosen here starts with a "|" so it shows as the
// last category in list. Change "model" to the name of your test.
func init() {
	app.AddWorld(&mainWorld{worldID: "worlds.main"})
	app.AddWorld(&testingWorld{worldID: "worlds.test"})
}
