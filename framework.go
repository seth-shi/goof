// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// framework/commands/console.go
// framework/config/app.go
// framework/database/data/users.json
// framework/database/migrations/2006_01_02_15_04_05__create_users_table.go
// framework/database/seeds/seeder.go
// framework/http/controllers/user.go
// framework/http/requests/user.go
// framework/main.go
// framework/models/user.go
// framework/resources/views/users/index.go
// framework/routes/api.go
// framework/routes/web.go
// framework/services/user.go
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _commandsConsoleGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\xce\xcf\xcd\x4d\xcc\x4b\x29\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\x05\x04\x00\x00\xff\xff\xa9\x25\xec\x49\x23\x00\x00\x00")

func commandsConsoleGoBytes() ([]byte, error) {
	return bindataRead(
		_commandsConsoleGo,
		"commands/console.go",
	)
}

func commandsConsoleGo() (*asset, error) {
	bytes, err := commandsConsoleGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/console.go", size: 35, mode: os.FileMode(438), modTime: time.Unix(1592100146, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configAppGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\xce\xcf\x4b\xcb\x4c\xe7\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\x05\x04\x00\x00\xff\xff\x3f\xe5\xdc\xe6\x21\x00\x00\x00")

func configAppGoBytes() ([]byte, error) {
	return bindataRead(
		_configAppGo,
		"config/app.go",
	)
}

func configAppGo() (*asset, error) {
	bytes, err := configAppGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/app.go", size: 33, mode: os.FileMode(438), modTime: time.Unix(1592100146, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _databaseDataUsersJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\x8a\x05\x04\x00\x00\xff\xff\x80\xdd\x9d\x70\x03\x00\x00\x00")

func databaseDataUsersJsonBytes() ([]byte, error) {
	return bindataRead(
		_databaseDataUsersJson,
		"database/data/users.json",
	)
}

func databaseDataUsersJson() (*asset, error) {
	bytes, err := databaseDataUsersJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "database/data/users.json", size: 3, mode: os.FileMode(438), modTime: time.Unix(1592100146, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _databaseMigrations2006_01_02_15_04_05__create_users_tableGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\x2b\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\x05\x04\x00\x00\xff\xff\xae\xc1\x74\x84\x25\x00\x00\x00")

func databaseMigrations2006_01_02_15_04_05__create_users_tableGoBytes() ([]byte, error) {
	return bindataRead(
		_databaseMigrations2006_01_02_15_04_05__create_users_tableGo,
		"database/migrations/2006_01_02_15_04_05__create_users_table.go",
	)
}

func databaseMigrations2006_01_02_15_04_05__create_users_tableGo() (*asset, error) {
	bytes, err := databaseMigrations2006_01_02_15_04_05__create_users_tableGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "database/migrations/2006_01_02_15_04_05__create_users_table.go", size: 37, mode: os.FileMode(438), modTime: time.Unix(1592100558, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _databaseSeedsSeederGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\x4d\x4d\x29\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\x3f\x03\x45\x76\x21\x00\x00\x00")

func databaseSeedsSeederGoBytes() ([]byte, error) {
	return bindataRead(
		_databaseSeedsSeederGo,
		"database/seeds/seeder.go",
	)
}

func databaseSeedsSeederGo() (*asset, error) {
	bytes, err := databaseSeedsSeederGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "database/seeds/seeder.go", size: 33, mode: os.FileMode(438), modTime: time.Unix(1592100146, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _httpControllersUserGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\xce\xcf\x2b\x29\xca\xcf\xc9\x49\x2d\x2a\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\xa8\xe6\xe2\xe4\xaa\xe5\x02\x04\x00\x00\xff\xff\x5c\xf4\x43\xb6\x27\x00\x00\x00")

func httpControllersUserGoBytes() ([]byte, error) {
	return bindataRead(
		_httpControllersUserGo,
		"http/controllers/user.go",
	)
}

func httpControllersUserGo() (*asset, error) {
	bytes, err := httpControllersUserGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "http/controllers/user.go", size: 39, mode: os.FileMode(438), modTime: time.Unix(1592100558, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _httpRequestsUserGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4a\x2d\x2c\x4d\x2d\x2e\x29\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\xcf\x33\x02\x63\x24\x00\x00\x00")

func httpRequestsUserGoBytes() ([]byte, error) {
	return bindataRead(
		_httpRequestsUserGo,
		"http/requests/user.go",
	)
}

func httpRequestsUserGo() (*asset, error) {
	bytes, err := httpRequestsUserGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "http/requests/user.go", size: 36, mode: os.FileMode(438), modTime: time.Unix(1592100558, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\x31\x4e\x03\x31\x10\x46\xe1\xda\x73\x8a\x1f\x57\x76\xe3\xf4\x48\x69\xa9\xb8\x84\xb1\x67\xbd\x03\x5e\x4f\xe4\x4c\x56\xa0\xd5\xde\x1d\x81\x94\xfa\xbd\xef\x96\xcb\x57\x6e\x8c\x2d\xcb\x20\x92\xed\xa6\xd3\x10\xc8\xf9\x26\xb6\x3e\x3e\x52\xd1\xed\xf2\xa9\xab\x5e\x9a\x56\x35\x1e\xbb\x27\xe7\xbb\x36\x4f\x91\x68\x79\x8c\xf2\x2f\x43\x04\x0e\x22\xc7\x73\xe2\xf5\x8a\xe7\x9b\xde\x35\xd7\x10\xc9\xc9\x82\xbf\xf4\x72\xc5\x90\x8e\x83\x9c\xeb\xda\xd2\x5b\xb6\xdc\x83\xb7\x95\x31\x55\x0d\x55\x26\x17\xd3\xf9\x83\xaa\x7c\xc7\x50\x03\x7f\xcb\xdd\x90\x78\xec\x58\xa4\xb3\x8f\xe4\x4e\x3a\x7f\x03\x00\x00\xff\xff\x28\xec\xa2\xd9\xb5\x00\x00\x00")

func mainGoBytes() ([]byte, error) {
	return bindataRead(
		_mainGo,
		"main.go",
	)
}

func mainGo() (*asset, error) {
	bytes, err := mainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go", size: 181, mode: os.FileMode(438), modTime: time.Unix(1592034921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelsUserGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\xcd\x4f\x49\xcd\x29\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\x5d\x2f\xe7\x81\x22\x00\x00\x00")

func modelsUserGoBytes() ([]byte, error) {
	return bindataRead(
		_modelsUserGo,
		"models/user.go",
	)
}

func modelsUserGo() (*asset, error) {
	bytes, err := modelsUserGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "models/user.go", size: 34, mode: os.FileMode(438), modTime: time.Unix(1592100558, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesViewsUsersIndexGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x2d\x4e\x2d\x2a\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\x35\x52\xce\x09\x21\x00\x00\x00")

func resourcesViewsUsersIndexGoBytes() ([]byte, error) {
	return bindataRead(
		_resourcesViewsUsersIndexGo,
		"resources/views/users/index.go",
	)
}

func resourcesViewsUsersIndexGo() (*asset, error) {
	bytes, err := resourcesViewsUsersIndexGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/views/users/index.go", size: 33, mode: os.FileMode(438), modTime: time.Unix(1592100934, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routesApiGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\xca\x2f\x2d\x49\x2d\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\xfd\x16\x7b\x6c\x22\x00\x00\x00")

func routesApiGoBytes() ([]byte, error) {
	return bindataRead(
		_routesApiGo,
		"routes/api.go",
	)
}

func routesApiGo() (*asset, error) {
	bytes, err := routesApiGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "routes/api.go", size: 34, mode: os.FileMode(438), modTime: time.Unix(1592100934, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routesWebGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\xca\x2f\x2d\x49\x2d\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\xcc\xcb\x2c\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\x05\x04\x00\x00\xff\xff\x78\x60\xe9\x72\x21\x00\x00\x00")

func routesWebGoBytes() ([]byte, error) {
	return bindataRead(
		_routesWebGo,
		"routes/web.go",
	)
}

func routesWebGo() (*asset, error) {
	bytes, err := routesWebGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "routes/web.go", size: 33, mode: os.FileMode(438), modTime: time.Unix(1592100934, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _servicesUserGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\x2d\x2a\xcb\x4c\x4e\x2d\xe6\xe2\x4a\x2b\xcd\x4b\x56\xc8\x4d\xcc\xcc\xd1\xd0\x54\x50\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\xae\xab\xff\x27\x24\x00\x00\x00")

func servicesUserGoBytes() ([]byte, error) {
	return bindataRead(
		_servicesUserGo,
		"services/user.go",
	)
}

func servicesUserGo() (*asset, error) {
	bytes, err := servicesUserGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "services/user.go", size: 36, mode: os.FileMode(438), modTime: time.Unix(1592100934, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"commands/console.go":                                            commandsConsoleGo,
	"config/app.go":                                                  configAppGo,
	"database/data/users.json":                                       databaseDataUsersJson,
	"database/migrations/2006_01_02_15_04_05__create_users_table.go": databaseMigrations2006_01_02_15_04_05__create_users_tableGo,
	"database/seeds/seeder.go":                                       databaseSeedsSeederGo,
	"http/controllers/user.go":                                       httpControllersUserGo,
	"http/requests/user.go":                                          httpRequestsUserGo,
	"main.go":                                                        mainGo,
	"models/user.go":                                                 modelsUserGo,
	"resources/views/users/index.go":                                 resourcesViewsUsersIndexGo,
	"routes/api.go":                                                  routesApiGo,
	"routes/web.go":                                                  routesWebGo,
	"services/user.go":                                               servicesUserGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"commands": &bintree{nil, map[string]*bintree{
		"console.go": &bintree{commandsConsoleGo, map[string]*bintree{}},
	}},
	"config": &bintree{nil, map[string]*bintree{
		"app.go": &bintree{configAppGo, map[string]*bintree{}},
	}},
	"database": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"users.json": &bintree{databaseDataUsersJson, map[string]*bintree{}},
		}},
		"migrations": &bintree{nil, map[string]*bintree{
			"2006_01_02_15_04_05__create_users_table.go": &bintree{databaseMigrations2006_01_02_15_04_05__create_users_tableGo, map[string]*bintree{}},
		}},
		"seeds": &bintree{nil, map[string]*bintree{
			"seeder.go": &bintree{databaseSeedsSeederGo, map[string]*bintree{}},
		}},
	}},
	"http": &bintree{nil, map[string]*bintree{
		"controllers": &bintree{nil, map[string]*bintree{
			"user.go": &bintree{httpControllersUserGo, map[string]*bintree{}},
		}},
		"requests": &bintree{nil, map[string]*bintree{
			"user.go": &bintree{httpRequestsUserGo, map[string]*bintree{}},
		}},
	}},
	"main.go": &bintree{mainGo, map[string]*bintree{}},
	"models": &bintree{nil, map[string]*bintree{
		"user.go": &bintree{modelsUserGo, map[string]*bintree{}},
	}},
	"resources": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"users": &bintree{nil, map[string]*bintree{
				"index.go": &bintree{resourcesViewsUsersIndexGo, map[string]*bintree{}},
			}},
		}},
	}},
	"routes": &bintree{nil, map[string]*bintree{
		"api.go": &bintree{routesApiGo, map[string]*bintree{}},
		"web.go": &bintree{routesWebGo, map[string]*bintree{}},
	}},
	"services": &bintree{nil, map[string]*bintree{
		"user.go": &bintree{servicesUserGo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
