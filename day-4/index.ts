export interface IGame {
    id: string
    winningCards: number[]
    ownCards: number[]
}

export const processInput = (input: string[]) => {
    return input.map((line) => {
        const [gameName, cards] = line.split(': ')
        const id = gameName.split(' ').pop()
        const [winningCards, ownCards] = cards.split(' | ')

        return ({
            id,
            winningCards: winningCards.split(' ').filter(Boolean).map((card) => +card),
            ownCards: ownCards.split(' ').filter(Boolean).map((card) => +card),
        })
    })
}
export const partOne = (games: IGame[]) => {
    return games.reduce((acc, curr) => {
        let points = 0
        curr.winningCards.forEach((card) => {
            if (curr.ownCards.some((ownCard) => ownCard === card)) {
                points = points ? (points * 2) : 1
            }
        })
        return acc + points
    }, 0)
}


export const partTwo = (input: string[]) => {
    return ''
}
