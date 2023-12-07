export interface INumber {
    value: number
    start: number
    end: number
}

export interface INumbersReducer {
    numbers: INumber[]
    currentNumber: string
    sum: number
    start?: number
    end?: number
}

const isVerticallyAdjacent = (start: number, end: number) => (_: string, index: number) => index >= start - 1 && index <= end + 1

export const partOne = (input: string[]) => {
    return input.reduce((acc, line, lineIndex, self) => {
        const numbers = line.split('').reduce<INumbersReducer>((state, char, index, numbersArr) => {
            if (Number.isInteger(parseInt(char))) {
                if(index === numbersArr.length - 1) {
                   return {
                       ...state,
                       numbers: [...state.numbers, {
                           value: Number(`${state.currentNumber}${char}`),
                           start: state.start ?? 0,
                           end: index
                       }],
                   }
                }
                return {
                    ...state,
                    currentNumber: `${state.currentNumber}${char}`,
                    start: Number.isInteger(state.start) ? state.start : index
                }
            }
            if (state.currentNumber) {
                return {
                    ...state,
                    numbers: [...state.numbers, {
                        value: Number(state.currentNumber),
                        start: state.start ?? 0,
                        end: index - 1
                    }],
                    currentNumber: ''
                }
            }
            return state
        }, {numbers: [] as INumber[], currentNumber: '', sum: 0}).numbers

        const sum = numbers.reduce((lineNumberAcc, number) => {
            let checks = []
            if (number.start > 0) {
                checks.push(line[number.start - 1])
            }
            if (number.end < line.length - 1 && number.end >= 0) {
                checks.push(line[number.end + 1])
            }
            if (lineIndex > 0) {
                const previousLine = self[(lineIndex as number) - 1].split('')
                checks.push(...previousLine.filter(isVerticallyAdjacent(number.start, number.end)))
            }
            if (lineIndex < self.length - 1) {
                const nextLine = self[(lineIndex as number) + 1].split('')
                checks.push(...nextLine.filter(isVerticallyAdjacent(number.start, number.end)))
            }


            const hasAdjacentSymbol = checks.some((value) => {
                return !Number.isInteger(parseInt(value)) && value !== '.'
            })

            if (hasAdjacentSymbol) {
                return lineNumberAcc + number.value
            }


            return lineNumberAcc
        }, 0)
        return acc + sum
    }, 0)
}

export const partTwo = (input: string[]) => {
    return ''
}
