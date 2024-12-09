package main

import (
	"strconv"
	"strings"
)

func stringToBlocks(str string) ([]*fileblock, []*fileblock) {
	blocks, freeSpaces, id, file := []*fileblock{}, []*fileblock{}, 0, true
	index := 0

	for idx, part := range str {
		if idx == len(str)-1 {
			break
		}
		num, _ := strconv.Atoi(string(part))
		if file {
			blocks = append(blocks, &fileblock{
				id:         id,
				startIndex: index,
				size:       num,
			})
			id += 1
			index += num
			file = false
		} else {
			block := &fileblock{
				size:       num,
				startIndex: index,
				freeSpace:  true,
			}
			index += num
			blocks = append(blocks, block)
			freeSpaces = append(freeSpaces, block)

			file = true
		}
	}

	return blocks, freeSpaces
}

func ProcessInput(input []string) interface{} {
	return strings.Join(input, "\n")
}

type fileblock struct {
	size, id, startIndex int
	freeSpace            bool
}

func PartOne(input interface{}) interface{} {
	blocks, _ := stringToBlocks(input.(string))
	id := 0

	i, j, sum, id := 0, len(blocks)-1, 0, 0

	for i <= j {
		if !blocks[i].freeSpace {
			for s := 0; s < blocks[i].size; s++ {
				sum += (id * blocks[i].id)
				id++
			}
			i++
			continue
		}

		if blocks[j].freeSpace {
			j--
			continue
		}

		if blocks[j].size == blocks[i].size {
			for s := 0; s < blocks[i].size; s++ {
				sum += (id * blocks[j].id)
				id++
			}

			i++
			j--
			continue
		}

		if blocks[j].size > blocks[i].size {
			for s := 0; s < blocks[i].size; s++ {
				sum += (id * blocks[j].id)
				id++
			}

			blocks[j].size -= blocks[i].size

			i++
			continue
		}

		if blocks[j].size < blocks[i].size {
			for s := 0; s < blocks[j].size; s++ {
				sum += (id * blocks[j].id)
				id++
			}

			blocks[i].size -= blocks[j].size

			j--
			continue
		}
	}

	return sum
}

func PartTwo(input interface{}) interface{} {
	blocks, freeSpaces := stringToBlocks(input.(string))

	j, sum := len(blocks)-1, 0

	for j >= 0 {
		if blocks[j].freeSpace {
			j--
			continue
		}

		foundFree := false
		for _, freeSpace := range freeSpaces {
			if freeSpace.size >= blocks[j].size && freeSpace.size > 0 && freeSpace.startIndex < blocks[j].startIndex {
				foundFree = true
				for s := 0; s < blocks[j].size; s++ {
					sum += ((freeSpace.startIndex + s) * blocks[j].id)
				}

				freeSpace.size -= blocks[j].size
				freeSpace.startIndex += blocks[j].size

				j--
				break
			}
		}

		if !foundFree {
			for s := 0; s < blocks[j].size; s++ {
				sum += ((blocks[j].startIndex + s) * blocks[j].id)
			}

			j--
			continue
		}

	}

	return sum
}

