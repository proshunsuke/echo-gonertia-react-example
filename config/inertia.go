package config

import (
	"echo-gonertia-react-example/lib"
	"echo-gonertia-react-example/provider"
	"encoding/json"
	"fmt"
	inertia "github.com/romsar/gonertia"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func InitInertia() *inertia.Inertia {
	rootDir := lib.ProjectRoot()
	viteHotFile := filepath.Join(rootDir, "public", "hot")
	rootViewFile := filepath.Join(rootDir, "resources", "views", "root.html")
	manifestPath := filepath.Join(rootDir, "public", "build", "manifest.json")
	viteManifestPath := filepath.Join(rootDir, "public", "build", ".vite", "manifest.json")
	flashProvider := provider.NewSessionFlashProvider()

	// check if laravel-vite-plugin is running in dev mode (it puts a "hot" file in the public folder)
	url, err := viteHotFileUrl(viteHotFile)
	if err != nil {
		log.Fatal(err)
	}
	if url != "" {
		i, err := inertia.NewFromFile(
			rootViewFile,
			inertia.WithFlashProvider(flashProvider),
		)
		if err != nil {
			log.Fatal(err)
		}

		i.ShareTemplateFunc("vite", func(entry string) (template.HTML, error) {
			if entry != "" && !strings.HasPrefix(entry, "/") {
				entry = "/" + entry
			}
			htmlTag := fmt.Sprintf(`<script type="module" src="%s%s"></script>`, url, entry)
			return template.HTML(htmlTag), nil
		})
		i.ShareTemplateFunc("viteReactRefresh", viteReactRefresh(url))
		return i
	}

	// laravel-vite-plugin not running in dev mode, use build manifest file
	// check if the manifest file exists, if not, rename it
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		// move the manifest from ./public/build/.vite/manifest.json to ./public/build/manifest.json
		// so that the vite function can find it
		err := os.Rename(viteManifestPath, manifestPath)
		if err != nil {
			return nil
		}
	}

	i, err := inertia.NewFromFile(
		rootViewFile,
		inertia.WithVersionFromFile(manifestPath),
		inertia.WithFlashProvider(flashProvider),
	)
	if err != nil {
		log.Fatal(err)
	}

	i.ShareTemplateFunc("vite", vite(manifestPath, "/public/build/"))
	i.ShareTemplateFunc("viteReactRefresh", viteReactRefresh(url))

	return i
}

// vite Generate Vite script tag and css link for production
func vite(manifestPath, buildDir string) func(path string) (template.HTML, error) {
	f, err := os.Open(manifestPath)
	if err != nil {
		log.Fatalf("cannot open provided vite manifest file: %s", err)
	}
	defer f.Close()

	viteAssets := make(map[string]*struct {
		File   string   `json:"file"`
		Source string   `json:"src"`
		CSS    []string `json:"css"`
	})
	err = json.NewDecoder(f).Decode(&viteAssets)
	// print content of viteAssets
	for k, v := range viteAssets {
		log.Printf("%s: %s\n", k, v.File)
	}

	if err != nil {
		log.Fatalf("cannot unmarshal vite manifest file to json: %s", err)
	}

	return func(p string) (template.HTML, error) {
		if val, ok := viteAssets[p]; ok {
			cssLinks := ""
			for _, css := range val.CSS {
				cssLinks += fmt.Sprintf(`<link rel="stylesheet" href="%s%s">`, buildDir, css)
			}
			htmlTag := fmt.Sprintf(
				`%s<script type="module" src="%s%s"></script>`,
				cssLinks,
				buildDir,
				val.File,
			)
			return template.HTML(htmlTag), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}

// viteReactRefresh Generate React refresh runtime script
func viteReactRefresh(url string) func() (template.HTML, error) {
	return func() (template.HTML, error) {
		if url == "" {
			return "", nil
		}
		script := fmt.Sprintf(`
<script type="module">
    import RefreshRuntime from '%s/@react-refresh'
    RefreshRuntime.injectIntoGlobalHook(window)
    window.$RefreshReg$ = () => {}
    window.$RefreshSig$ = () => (type) => type
    window.__vite_plugin_react_preamble_installed__ = true
</script>`, url)

		return template.HTML(script), nil
	}
}

// viteHotFileUrl Get the vite hot file url
func viteHotFileUrl(viteHotFile string) (string, error) {
	_, err := os.Stat(viteHotFile)
	if err != nil {
		return "", nil
	}
	content, err := os.ReadFile(viteHotFile)
	if err != nil {
		return "", err
	}
	url := strings.TrimSpace(string(content))
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		url = url[strings.Index(url, ":")+1:]
	} else {
		url = "//localhost:1323"
	}
	return url, nil
}
