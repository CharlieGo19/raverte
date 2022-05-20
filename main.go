package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"raverte/raverte"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte
func main() {
	// Create an instance of the app structure
	raverte := raverte.RaverteInit()
	app := NewApp(raverte)

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "raverte",
		Width:             1024,
		Height:            768,
		MinWidth:          720,
		MinHeight:         570,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false, // true,
		//RGBA:              &options.RGBA{255, 255, 255, 1},
		Assets:     assets,
		LogLevel:   logger.DEBUG,
		OnStartup:  app.startup,
		OnDomReady: app.domReady,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			raverte,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHidden(),
			About: &mac.AboutInfo{
				Title:   "Oh, we have menus now!?",
				Message: "Thanks @wails ☺️\np.s. Please send graphics help. @raverte_",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
