package main

import (
	"changeme/lib"
	"context"
	"github.com/mitchellh/go-homedir"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path"
)

type FileWithContent struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
	HTML    string `json:"html,omitempty"`
}

// App struct
type App struct {
	ctx          context.Context
	filename     string
	markdown     string
	renderedHTML string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// SendMarkdownToRenderer sends markdown to the renderer
func (a *App) SendMarkdownToRenderer(content string) string {
	a.markdown = content
	a.renderedHTML = lib.RenderMD(content)
	return a.renderedHTML
}

func (a *App) Save() (string, error) {
	hd, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	chosenPath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:           path.Join(hd, "Documents"),
		DefaultFilename:            a.filename,
		Title:                      "Save file",
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: true,
	})
	if err != nil {
		return "", err
	}
	err = os.WriteFile(chosenPath, []byte(a.markdown), 0644)
	if err != nil {
		return "", err
	}
	return chosenPath, nil
}

func (a *App) Open() (FileWithContent, error) {
	hd, err := homedir.Dir()
	if err != nil {
		return FileWithContent{}, err
	}
	openedFile, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: path.Join(hd, "Documents"),
		Title:            "Open file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Markdown files",
				Pattern:     "*.md",
			},
			{
				DisplayName: "All files",
				Pattern:     "*",
			},
		},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: true,
	})
	fileContent, err := os.ReadFile(openedFile)
	if err != nil {
		return FileWithContent{}, err
	}
	a.filename = openedFile
	a.markdown = string(fileContent)
	a.renderedHTML = lib.RenderMD(a.markdown)
	return FileWithContent{
		Name:    a.filename,
		Content: a.markdown,
		HTML:    a.renderedHTML,
	}, nil
}
