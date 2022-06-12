package color_map

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"
)

// ColorMap ...
type ColorMap struct {
	cMap           map[uint32]map[uint32]map[uint32]int
	objNameByIdMap map[int]string
	objIdByNameMap map[string]int
	objCounter     int
}

// NewColorMap ...
func NewColorMap() *ColorMap {
	return &ColorMap{
		cMap:           make(map[uint32]map[uint32]map[uint32]int),
		objNameByIdMap: make(map[int]string),
		objIdByNameMap: make(map[string]int),
	}
}

// LoadDir ...
func (cm *ColorMap) LoadDir(objectsDir string) error {
	files, err := ioutil.ReadDir(objectsDir)
	if err != nil {
		return err
	}

	objNames := make([]string, 0)

	for _, f := range files {
		imgFile := objectsDir + f.Name()
		objName := strings.Split(f.Name(), "_")[0]
		objNames = append(objNames, objName)

		objId, ok := cm.objIdByNameMap[objName]
		if !ok {
			cm.objCounter++
			cm.objNameByIdMap[cm.objCounter] = objName
			cm.objIdByNameMap[objName] = cm.objCounter
			objId = cm.objCounter
		}

		file, e := os.Open(imgFile)
		if e != nil {
			return fmt.Errorf("open file (%s) error: %v", imgFile, e)
		}
		defer file.Close()

		img, _, e := image.Decode(file)
		if e != nil {
			return fmt.Errorf("decode file (%s) error: %v", imgFile, e)
		}

		for x := 0; x < img.Bounds().Max.X; x++ {
			for y := 0; y < img.Bounds().Max.Y; y++ {
				r, g, b, a := img.At(x, y).RGBA()
				if a == 0 {
					continue
				}

				if _, ok := cm.cMap[r][g][b]; ok {
					delete(cm.cMap[r][g], b)
					continue
				}

				if _, ok := cm.cMap[r]; !ok {
					cm.cMap[r] = make(map[uint32]map[uint32]int)
				}
				if _, ok := cm.cMap[r][g]; !ok {
					cm.cMap[r][g] = make(map[uint32]int)
				}

				cm.cMap[r][g][b] = objId
			}
		}
	}

	return nil
}

// Which ...
func (cm *ColorMap) Which(r, g, b uint32) string {
	objIdx, ok := cm.cMap[r][g][b]
	if !ok {
		return ""
	}

	objName, ok := cm.objNameByIdMap[objIdx]
	return objName
}
