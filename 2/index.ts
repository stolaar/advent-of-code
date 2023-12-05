import {readInput} from "../utils";
import path from "path";

const part1Config : Record<string, number>= {
    green: 13,
    red: 12,
    blue: 14,
}

export const part1 = (input: string[]) => {
    return input.filter((line) => {
       const [, results] = line.split(': ')
         const reveals = results.split('; ')
        return reveals.every((reveal) => {
              const revealedColor= reveal.split(', ')
            return revealedColor.every((color) => {
                const [colorValue, colorName] = color.split(' ')
                return (part1Config[colorName] ?? 0) >= parseInt(colorValue)
            })
        })
    }).reduce((acc, line) => {
        const [name] = line.split(': ')
        return acc + parseInt(name.split(' ').pop() ?? '0')
    }, 0)
}

export const part2 = (input: string[]) => {
    const configMultiplied = Object.values(part1Config).reduce((acc, value) => acc * value, 1)
    return input.reduce((result, line) => {
       const fewestColors= (line.split(': ').pop() as string).split('; ').reduce((acc, set) => {
          const setColors = set.split(', ')
           setColors.forEach((color) => {
            const [value, colorName] = color.split(' ')
               if(acc[colorName] < parseInt(value)) {
                   acc[colorName] = parseInt(value)
               }
           })
            return acc
       }, {red: 0, green: 0, blue: 0} as Record<string, number>)
        const powerOfFewestColors = Object.values(fewestColors).reduce((acc, value) => acc * value, 1)
            return result + powerOfFewestColors
    }, 0)
}

export const main = () => {
    const example = readInput(path.resolve(__dirname, 'example.txt'))
    const input = readInput(path.resolve(__dirname, 'input.txt'))

    console.log("Solution example", part1(example))
    console.log("Solution part 1", part1(input))
    console.log("Solution part 2 example", part2(example))
    console.log("Solution part 2", part2(input))
}

main()

