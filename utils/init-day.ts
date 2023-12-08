import path from "path";
import * as fs from "fs";

const day = process.argv[2]

if (!day) {
  console.error('Missing day argument')
  process.exit(1)
}

const folderName = path.join(__dirname, '..', `day-${day}`)

try {
    fs.mkdirSync(folderName)
} catch (err) {
    console.error('Folder already exists', err)
    process.exit(1)
}

const template = `
export const processInput = (input: string[]) => {
    return input
}

export const partOne = (input: string[]) => {
    return ''
}

export const partTwo = (input: string[]) => {
    return ''
}
`

fs.writeFileSync(path.join(folderName, 'index.ts'), template)
fs.writeFileSync(path.join(folderName, 'input.txt'), 'demo input')
