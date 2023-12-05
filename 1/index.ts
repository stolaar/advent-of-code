import {readInput} from "../utils";
import path from "path";

const wordDigitsMap: Record<string, number> = {
    'zero': 0,
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9,
}
const getSumOfCalibrationValues = (input: string[]) => {
    return input.reduce((acc, curr) => {
        const [firstNumber, ...numberValues] = curr.split('').filter(char => Number.isInteger(parseInt(char)))

        if (!firstNumber) return acc
        return acc + (parseInt(`${firstNumber}${numberValues[numberValues.length - 1] ?? firstNumber}`))
    }, 0)
}

const getSumOfCalibrationValues2 = (input: string[]) => {
    return input.reduce((acc, curr) => {
        const [firstNumber, ...restNumber] = curr.split('').reduce((validNumbers, currentChar, index, self) => {

            const word = self.slice(index).join('')
            const foundWord = Object.keys(wordDigitsMap).find(wordKey=> word.startsWith(wordKey))
            if(foundWord) {
                validNumbers.push(wordDigitsMap[foundWord as string])
                return validNumbers
            }
            if (Number.isInteger(parseInt(currentChar))) {
                validNumbers.push(parseInt(currentChar))
                return validNumbers
            }
            return validNumbers
        }, [] as number[])

        if (!firstNumber) return acc
        const sum =  (parseInt(`${firstNumber}${restNumber[restNumber.length - 1] ?? firstNumber}`))
        return acc + sum
    }, 0)
}

export const main = () => {
    const input = readInput(path.resolve(__dirname, 'input.txt'))
    const input2 = readInput(path.resolve(__dirname, 'input2.txt'))

    console.log("Solution part 1", getSumOfCalibrationValues(input))
    console.log("Solution part 2", getSumOfCalibrationValues2(input2))
}

main()

