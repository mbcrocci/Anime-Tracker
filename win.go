package main

import "gopkg.in/qml.v1"

// For now just shows a black rectangle
func RunQmlApp() error {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("app.qml")
	if err != nil {
		return err
	}

	context := engine.Context()
	context.SetVar("animeList", &Anime{})
	context.SetVar("db", db)

	window := component.CreateWindow(nil)
	window.Show()
	window.Wait()

	return nil
}
