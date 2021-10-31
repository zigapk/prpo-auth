// Code generated for package config by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.ini
// configs/migrations/01_initial.up.sql
package config

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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x5f\x4f\xf2\x30\x18\xc5\xef\x9f\x4f\xf1\x7c\x02\xd8\xfb\xa2\xc6\x9b\x5d\x94\xad\x8e\x86\xd2\x2e\x6b\xc1\x3f\x64\x69\x06\xd4\x41\x9c\x74\x69\x07\x89\xdf\xde\x14\x35\x9a\x68\x9a\x5e\x9c\x5f\x4e\xcf\x39\xe9\xba\x73\x6d\xa8\x81\xcb\xc2\x70\xba\xa2\x1c\x53\x4c\xe0\x8e\x71\x6a\xb8\x2c\x0a\x26\x0a\x4c\x71\xf0\x27\x0b\x99\x14\x4a\xfe\xc6\xf1\x61\x49\xf4\x0c\x53\x8c\x49\xe3\xe6\x34\xec\x47\x9d\x6b\x61\x41\x1e\x8c\x62\x4f\x14\x53\xfc\x97\x5c\x14\x29\x7e\x88\x29\xc9\xe6\xcb\x52\x7d\x00\x58\xef\x9a\xa1\xd9\x34\xc1\xd6\x30\x93\x4a\x5f\xc2\xb6\x4d\xb7\x77\x61\x80\x52\x56\x11\x5c\x5f\x4d\xfe\xc3\x52\xd1\x0a\x53\xec\x7d\xef\xa0\x24\x4a\xdd\xcb\x2a\xc7\x14\xbd\x73\x43\xbc\x90\x4f\x8d\x20\x8b\x58\xd3\xbb\x30\xb4\xde\x06\x50\x8a\x9b\x85\xcc\x23\xdb\x1d\x42\xb3\xe9\x2c\xc0\x3a\x58\x7f\xb6\xbe\x06\xce\x94\xa6\xc2\x90\x3c\xaf\xa8\x8a\x63\x92\xd1\xe5\x7c\x95\xde\x26\x49\x5c\xd7\xb9\xf6\x70\xac\x41\xcb\x39\x15\x46\xeb\xf8\x49\x93\x9b\x24\x01\xc5\x0a\xc1\x44\x61\xca\x8a\xad\x88\xa6\x66\x4e\x1f\x0d\x97\x19\xd1\x4c\x0a\x4c\x71\xeb\x8e\xcf\xe3\xde\x1f\xce\x2f\xf6\x6d\xd4\xdb\xd7\x6f\xff\x72\xca\x59\xf6\xb7\xfd\xb4\xf9\x74\xbf\x07\x00\x00\xff\xff\x99\x07\x6a\x59\x9a\x01\x00\x00")

func configIniBytes() ([]byte, error) {
	return bindataRead(
		_configIni,
		"config.ini",
	)
}

func configIni() (*asset, error) {
	bytes, err := configIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.ini", size: 410, mode: os.FileMode(420), modTime: time.Unix(1635673246, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xcf\x4a\xf3\x40\x14\xc5\xf7\xf3\x14\x67\xd9\x42\xdf\xa0\xab\xf9\x92\x5b\xbe\x60\x32\xa9\x93\x1b\xb4\x6e\xc2\x90\xb9\x60\xb0\x4d\x25\x33\xd1\xd7\x97\xa6\xb6\x52\x41\x44\xc8\xdd\x5d\x38\xfc\x38\x7f\x12\x4b\x9a\x09\xac\xff\xe5\x84\x31\xc8\x10\xd4\x42\x01\xc0\xd8\x79\x5c\x8f\xe9\x91\xb1\xb5\x59\xa1\xed\x0e\x77\xb4\x5b\xa9\x49\x23\x07\xd7\xed\x6f\x34\x7f\x3b\x53\x32\x4c\x9d\xe7\xa8\x4d\x76\x5f\xd3\x6a\x82\xf6\xee\x20\x98\x01\xfa\xe9\xf1\xd5\x85\xf0\x7e\x1c\xfc\x4c\x38\xef\xa2\x34\xed\x20\x2e\x8a\x07\x67\x05\x55\xac\x8b\x2d\x1e\x32\xfe\x3f\xbd\x78\x2a\x0d\x21\xa5\x8d\xae\x73\x46\x52\x5b\x4b\x86\x9b\x2f\xe1\x05\xa7\x96\x6b\xa5\x6e\xca\x77\x63\x7c\x96\x3e\x76\xed\x09\xdd\x78\x79\xeb\x5a\xb9\x8e\x11\x64\x68\x2e\x83\x4c\x21\x2c\x6d\xc8\x92\x49\xa8\x3a\xaf\x86\xc5\xd8\xf9\x25\x4a\x83\x94\x72\x62\x42\xa2\xab\x44\xa7\xf4\x3d\xc4\xe9\x8b\xc7\x17\xe9\xe7\x6b\x78\xae\x4a\xce\xe6\xf6\x2e\xc4\x66\x0c\x32\x85\xfd\x91\xf6\xbb\x39\xb5\x5c\x7f\x04\x00\x00\xff\xff\xe3\x94\x2c\x26\xdc\x02\x00\x00")

func migrations01_initialUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initialUpSql,
		"migrations/01_initial.up.sql",
	)
}

func migrations01_initialUpSql() (*asset, error) {
	bytes, err := migrations01_initialUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_initial.up.sql", size: 732, mode: os.FileMode(420), modTime: time.Unix(1635670826, 0)}
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
	"config.ini":                   configIni,
	"migrations/01_initial.up.sql": migrations01_initialUpSql,
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
	"config.ini": &bintree{configIni, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_initial.up.sql": &bintree{migrations01_initialUpSql, map[string]*bintree{}},
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
