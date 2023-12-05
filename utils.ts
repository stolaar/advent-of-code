import * as fs from 'fs'

export const readInput = (inputPath: string) => {
    const input = fs.readFileSync(inputPath, 'utf8')
    const result = input.split('\n')
    result.pop()
    return result
}
