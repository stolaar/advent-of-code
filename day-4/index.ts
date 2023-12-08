import {times} from "../utils/array.utils";

export interface IGame {
    id: number
    winningCards: number[]
    ownCards: number[]
}

export const processInput = (input: string[]): IGame[] => {
    return input.map((line) => {
        const [gameName, cards] = line.split(': ')
        const id = gameName.split(' ').pop()
        const [winningCards, ownCards] = cards.split(' | ')

        return ({
            id: parseInt(id ?? '0'),
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

const getNumberOfMatchingCards = (game: IGame) => game.winningCards.filter((card) => game.ownCards.some((ownCard) => ownCard === card)).length

export const partTwo = (games: IGame[]) => {
    const map = games.reduce((acc, game, currentIndex, self) => {
        const matches = getNumberOfMatchingCards(game)
        acc.set(game.id, (acc.get(game.id) ?? 0) + 1)
        if(!matches) return acc
        times(matches, (index) => {
            const nextGame =  self[currentIndex + index + 1]
            acc.set(nextGame.id, (acc.get(nextGame.id) ?? 0) + (acc.get(game.id) ?? 0))
        })

        return acc

    }, new Map<number, number>())

    return [...map.values()].reduce((acc, curr) => acc + curr, 0)
}
