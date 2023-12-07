export interface INumber {
    value: number
    start: number
    end: number
}

export interface INumbersReducer {
    currentNumber: string
    sum: number
    start?: number
}
const isVerticallyAdjacent = (start: number, end: number) => (_: string, index: number) => index >= start - 1 && index <= end + 1

const partNumber = (number: INumber,currentLine: string,lineIndex: number, lines: string[]) => {
    let checks = []
    if (number.start > 0) {
        checks.push(currentLine[number.start - 1])
    }
    if (number.end < currentLine.length - 1 && number.end >= 0) {
        checks.push(currentLine[number.end + 1])
    }
    if (lineIndex > 0) {
        const sliced = lines[lineIndex - 1].substring(number.start - 1, number.end + 2)
        checks.push(...sliced)
    }
    if (lineIndex < lines.length - 1) {
        const sliced = lines[lineIndex + 1].substring(number.start - 1, number.end + 2)
        checks.push(...sliced)
    }

    const hasAdjacentSymbol = checks.some((value) => {
        return !Number.isInteger(parseInt(value)) && value !== '.'
    })

    if (hasAdjacentSymbol) {
        return number.value
    }

    return 0
}

const initialNumbersReduce: INumbersReducer = {
    currentNumber: '',
    sum: 0,
}

export const partOne = (input: string[]) => {
    return input.reduce((acc, line, lineIndex, self) => {
       return acc + line.split('').reduce<INumbersReducer>((state, char, index, numbersArr) => {
            if (Number.isInteger(parseInt(char))) {
                if(index === numbersArr.length - 1) {
                    const number ={
                        value: Number(`${state.currentNumber}${char}`),
                        start: state.start ?? 0,
                        end: index
                    }
                    return { ...state, sum: state.sum + partNumber(number, line, lineIndex, self) }
                }
                return {
                    ...state,
                    currentNumber: `${state.currentNumber}${char}`,
                    start: Number.isInteger(state.start) ? state.start : index
                }
            }
            if (state.currentNumber) {
                return {...initialNumbersReduce, sum: state.sum + partNumber({
                        value: Number(state.currentNumber),
                        start: state.start ?? 0,
                        end: index - 1
                    }, line, lineIndex, self) }
            }
            return {...initialNumbersReduce, sum: state.sum}
        }, initialNumbersReduce).sum
    }, 0)
}

export const partTwo = (input: string[]) => {
    return ''
}
