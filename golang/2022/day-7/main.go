package main

import (
	"math"
	"strconv"
	"strings"
)

type File struct {
	name     string
	size     int
	fileType string
}

type Dir struct {
	totalSize int
	name      string
	dirs      map[string]*Dir
	files     map[string]*File
	parent    *Dir
	level     int
}

type Filesystem struct {
	root       *Dir
	currentDir *Dir
}

func (dir *Dir) addDir(name string) {
	if _, ok := dir.dirs[name]; !ok {
		dir.dirs[name] = &Dir{
			name:      name,
			totalSize: 0,
			dirs:      map[string]*Dir{},
			files:     map[string]*File{},
			parent:    dir,
			level:     dir.level + 1,
		}
	}
}

func (dir *Dir) addFile(filename string, size int) {
	if _, ok := dir.files[filename]; !ok {
		filenameParts := strings.Split(filename, ".")
		var fileType string
		name := filenameParts[0]
		if len(filenameParts) > 1 {
			fileType = filenameParts[1]
		}

		dir.files[filename] = &File{
			name:     name,
			fileType: fileType,
			size:     size,
		}

		dir.totalSize += size

		parent := dir.parent

		for parent != nil {
			parent.totalSize += size
			parent = parent.parent
		}
	}
}

func (dir *Dir) list(lines []string, i int) int {
	for i < len(lines) {
		line := lines[i]
		if line == "" {
			i++
			continue
		}
		if strings.HasPrefix(line, "$") {
			return i
		}

		parts := strings.Split(line, " ")

		if strings.HasPrefix(line, "dir") {
			name := parts[1]
			dir.addDir(name)
			i++
			continue
		}

		sizeString, filename := parts[0], parts[1]
		size, _ := strconv.Atoi(sizeString)

		dir.addFile(filename, size)
		i++
	}
	return i
}

func (fs *Filesystem) cd(to string) {
	if to == ".." {
		fs.currentDir = fs.currentDir.parent
		return
	}

	if _, ok := fs.currentDir.dirs[to]; !ok {
		fs.currentDir.dirs[to] = &Dir{
			name:      to,
			totalSize: 0,
			dirs:      map[string]*Dir{},
			files:     map[string]*File{},
			parent:    fs.currentDir,
			level:     fs.currentDir.level + 1,
		}
	}
	fs.currentDir = fs.currentDir.dirs[to]
}

func (fs *Filesystem) Exec(line string, lines []string, i int) int {
	parts := strings.Split(line, " ")
	command := parts[1]

	if command == "cd" {
		dir := parts[2]
		fs.cd(dir)
		return i
	}

	if command == "ls" {
		i = fs.currentDir.list(lines, i)
		return i
	}
	return i
}

func ProcessInput(lines []string) interface{} {
	rootDir := &Dir{
		totalSize: 0,
		name:      "/",
		dirs:      map[string]*Dir{},
		files:     map[string]*File{},
		level:     0,
	}

	fs := Filesystem{
		root:       rootDir,
		currentDir: rootDir,
	}

	i := 0

	for i < len(lines) {
		line := lines[i]
		if strings.HasPrefix(line, "$") {
			i = fs.Exec(line, lines, i+1)
			continue
		}
		i++
	}

	return fs
}

func countSize(root *Dir, count int) int {
	if root.totalSize <= 100000 {
		count += root.totalSize
	}

	for _, dir := range root.dirs {
		count = countSize(dir, count)
	}
	return count
}

func PartOne(input interface{}) interface{} {
	fs := input.(Filesystem)
	return countSize(fs.root, 0)
}

func findDirToRemove(minToRemove int, dir *Dir, currentMin int) int {
	if dir.totalSize > minToRemove && dir.totalSize < currentMin {
		currentMin = dir.totalSize
	}

	for _, subdir := range dir.dirs {
		currentMin = findDirToRemove(minToRemove, subdir, currentMin)
	}
	return currentMin
}

func PartTwo(input interface{}) interface{} {
	fs, totalAvailable, minUnused := input.(Filesystem), 70_000_000, 30_000_000
	requiredToRemove := minUnused - (totalAvailable - fs.root.totalSize)

	return findDirToRemove(requiredToRemove, fs.root, math.MaxInt64)
}
