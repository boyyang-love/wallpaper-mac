package wallpaper

import (
	"errors"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type Wallpaper struct {
	Workspace *appkit.Workspace
	Screen    *appkit.Screen
}

func NewWallpaper() *Wallpaper {

	workspace := appkit.Workspace_SharedWorkspace()
	screen := appkit.Screen_MainScreen()

	return &Wallpaper{
		Workspace: &workspace,
		Screen:    &screen,
	}
}

func (w *Wallpaper) SetWallpaper(path string) error {
	if ok := w.Workspace.SetDesktopImageURLForScreenOptionsError(
		foundation.URL_FileURLWithPath(path),
		w.Screen,
		map[appkit.WorkspaceDesktopImageOptionKey]objc.IObject{},
		foundation.IError(nil),
	); !ok {
		return errors.New("failed to set wallpaper")
	}

	return nil
}

func (w *Wallpaper) GetWallpaper() string {
	return w.Workspace.DesktopImageURLForScreen(w.Screen).Description()
}
