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

const partNumber = (number: INumber, currentLine: string, lineIndex: number, lines: string[]) => {
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
                if (index === numbersArr.length - 1) {
                    const number = {
                        value: Number(`${state.currentNumber}${char}`),
                        start: state.start ?? 0,
                        end: index
                    }
                    return {...state, sum: state.sum + partNumber(number, line, lineIndex, self)}
                }
                return {
                    ...state,
                    currentNumber: `${state.currentNumber}${char}`,
                    start: Number.isInteger(state.start) ? state.start : index
                }
            }
            if (state.currentNumber) {
                return {
                    ...initialNumbersReduce, sum: state.sum + partNumber({
                        value: Number(state.currentNumber),
                        start: state.start ?? 0,
                        end: index - 1
                    }, line, lineIndex, self)
                }
            }
            return {...initialNumbersReduce, sum: state.sum}
        }, initialNumbersReduce).sum
    }, 0)
}

export interface IPartTwoNumber extends INumber {
    isPartNumber?: boolean
}

export interface IGearSymbol {
    x: number
}

export interface ILine {
    numbers: IPartTwoNumber[]
    gearSymbols: IGearSymbol[]
    start?: number
    end?: number
    currentNumber: string
}

const isAdjacentToGearSymbol = (x: number, isSameLine = false) => (num: IPartTwoNumber) => {
    if (!num.isPartNumber) return false

    if (!isSameLine) {
        let isAdjacent = false
        for (let i = num.start - 1; i <= num.end + 1; i++) {
           if(i === x) {
               isAdjacent = true
               break
           }
        }
        return isAdjacent
    }

    return (x === num.end + 1 || x === num.start - 1)
}

export const partTwo = (input: string[]) => {
    return input.reduce<ILine[]>((acc, line, lineIndex, self) => {
        const lineDetails = line.split('').reduce<ILine>((state, char, index, numbersArr) => {
            if (Number.isInteger(parseInt(char))) {
                if (index === numbersArr.length - 1) {
                    const number: IPartTwoNumber = {
                        value: Number(`${state.currentNumber}${char}`),
                        start: state.start ?? 0,
                        end: index,
                    }
                    number.isPartNumber = !!partNumber(number, line, lineIndex, self)
                    return {...state, numbers: [...state.numbers, number]}
                }
                return {
                    ...state,
                    currentNumber: `${state.currentNumber}${char}`,
                    start: Number.isInteger(state.start) ? state.start : index
                }
            }
            if (state.currentNumber) {
                const number: IPartTwoNumber = {
                    value: Number(state.currentNumber),
                    start: state.start ?? 0,
                    end: index - 1
                }
                number.isPartNumber = !!partNumber(number, line, lineIndex, self)
                return {
                    ...initialNumbersReduce, numbers: [...state.numbers, number],
                    gearSymbols: char === '*' ? [...state.gearSymbols, {x: index}] : state.gearSymbols
                }
            }
            if (char === '*') {
                return {
                    ...initialNumbersReduce,
                    ...state,
                    gearSymbols: [...state.gearSymbols, {x: index}]
                }
            }
            return state
        }, {numbers: [], gearSymbols: [], currentNumber: ''})
        acc.push(lineDetails)
        return acc
    }, []).reduce<number>((acc, line, index, self) => {
        line.gearSymbols.forEach((gearSymbol) => {
            let adjacentNumbers = []
            const horizontallyAdjacentNumbers = line.numbers.filter(isAdjacentToGearSymbol(gearSymbol.x, true))

            adjacentNumbers.push(...horizontallyAdjacentNumbers)
            if (index > 0) {
                const prevLine = self[index - 1]
                const topAdjacentNumbers = prevLine.numbers.filter(isAdjacentToGearSymbol(gearSymbol.x))
                adjacentNumbers.push(...topAdjacentNumbers)
            }

            if (index < self.length - 1) {
                const nextLine = self[index + 1]
                const bottomAdjacentNumbers = nextLine.numbers.filter(isAdjacentToGearSymbol(gearSymbol.x))
                adjacentNumbers.push(...bottomAdjacentNumbers)
            }

            if (adjacentNumbers.length === 2) {
                acc += adjacentNumbers.reduce((prod, curr) => prod * curr.value, 1)
            }
        })

        return acc
    }, 0)
}
