interface IHand {
  cards: string[]
  bid: number
}

interface IHandValue {
  firstEvalStrength: number
  secondEvalStrengths: number[]
}

interface IEvaluatedHand extends IHand {
  value: IHandValue
}

export interface IHandValueWithJokers extends IHandValue {
  numberOfJokers: number
}

const cardStrengthsMap: Record<string | number, number> = {
  A: 14,
  K: 13,
  Q: 12,
  J: 11,
  T: 10,
}

const getNumberOfMatches = (cards: string[]) => {
  const numberOfMatches = cards.reduce((acc, curr) => {
    acc.set(curr, (acc.get(curr) || 0) + 1)
    return acc
  }, new Map<string, number>())

  const mostMatched = [...numberOfMatches.entries()].reduce(
    (acc, [card, count]) => {
      if (count > acc.count) {
        return { card, count }
      }
      return acc
    },
    { card: "", count: 0 },
  )

  return { numberOfMatches, mostMatched }
}

const getCardStrength = (card: string, useJokers = false) => {
  if (card === "J" && useJokers) {
    return 1
  }

  return cardStrengthsMap[card] || +card
}

const evaluateHand = (cards: string[]): IHandValue => {
  const { numberOfMatches, mostMatched } = getNumberOfMatches(cards)

  const cardStrengths = cards.map((card) => cardStrengthsMap[card] || +card)

  if (mostMatched.count > 3) {
    return {
      firstEvalStrength: mostMatched.count + 2,
      secondEvalStrengths: cardStrengths,
    }
  }

  if (mostMatched.count === 3) {
    const isFullHouse =
      [...numberOfMatches.values()].filter((count) => count === 2).length === 1
    return {
      firstEvalStrength: isFullHouse ? 5 : 4,
      secondEvalStrengths: cardStrengths,
    }
  }
  const numberOfPairs = [...numberOfMatches.values()].filter(
    (match) => match === 2,
  ).length
  if (numberOfPairs === 2) {
    return {
      firstEvalStrength: 3,
      secondEvalStrengths: cardStrengths,
    }
  }
  if (numberOfPairs === 1) {
    return {
      firstEvalStrength: 2,
      secondEvalStrengths: cardStrengths,
    }
  }

  return {
    firstEvalStrength: 1,
    secondEvalStrengths: cardStrengths,
  }
}

const sortCardStrengths =
  (useJoker = false) =>
  (a: IEvaluatedHand, b: IEvaluatedHand) => {
    if (a.value.firstEvalStrength > b.value.firstEvalStrength) {
      return 1
    }
    if (a.value.firstEvalStrength < b.value.firstEvalStrength) {
      return -1
    }

    let winner = 0
    for (let i = 0; i < a.value.secondEvalStrengths.length; i++) {
      if (a.value.secondEvalStrengths[i] > b.value.secondEvalStrengths[i]) {
        winner = 1
        break
      }
      if (a.value.secondEvalStrengths[i] < b.value.secondEvalStrengths[i]) {
        winner = -1
        break
      }
    }
    return winner
  }

export const processInput = (input: string[]) => {
  return input.map((line) => {
    const [cards, bid] = line.split(" ")
    return { cards: cards.split(""), bid: Number(bid) }
  })
}

export const partOne = (hands: IHand[]) => {
  const evaluatedHands = hands.map((hand) => {
    return {
      ...hand,
      value: evaluateHand(hand.cards),
    }
  })

  evaluatedHands.sort(sortCardStrengths())

  return evaluatedHands.reduce((acc, curr, index) => {
    return acc + curr.bid * (index + 1)
  }, 0)
}

const evaluateHandWithJokers = (cards: string[]) => {
  const { numberOfMatches, mostMatched } = getNumberOfMatches(cards)
  let mostMatchedCount = mostMatched.count
  const hasJoker = cards.includes("J")

  const cardStrengths = cards.map((card) => getCardStrength(card, true))

  if (mostMatched.card !== "J" && mostMatched.count > 1) {
    mostMatchedCount += numberOfMatches.get("J") || 0
  }

  if (mostMatched.card === "J" && mostMatched.count < 5) {
    mostMatchedCount +=
      [...numberOfMatches.entries()]
        .filter(([card]) => card !== "J")
        .reduce((acc, [, curr]) => (acc < curr ? curr : acc), 0) || 1
  }

  if (mostMatchedCount > 3) {
    return {
      firstEvalStrength: mostMatchedCount + 2,
      secondEvalStrengths: cardStrengths,
    }
  }

  if (mostMatchedCount === 3) {
    const isFullHouse =
      [...numberOfMatches.entries()].filter(
        ([card, count]) => count === 2 && card !== mostMatched.card,
      ).length === 1
    return {
      firstEvalStrength: isFullHouse ? 5 : 4,
      secondEvalStrengths: cardStrengths,
    }
  }
  const numberOfPairs = [...numberOfMatches.values()].filter(
    (match) => match === 2,
  ).length
  if (numberOfPairs === 2) {
    return {
      firstEvalStrength: 3,
      secondEvalStrengths: cardStrengths,
    }
  }
  if (numberOfPairs === 1) {
    return {
      firstEvalStrength: 2,
      secondEvalStrengths: cardStrengths,
    }
  }
  if (cards.includes("J")) {
    return {
      firstEvalStrength: 2,
      secondEvalStrengths: cardStrengths,
    }
  }
  return {
    firstEvalStrength: 1,
    secondEvalStrengths: cardStrengths,
  }
}

export const partTwo = (hands: IHand[]) => {
  const evaluatedHands = hands.map((hand) => {
    return {
      ...hand,
      value: evaluateHandWithJokers(hand.cards),
    }
  })

  evaluatedHands.sort(sortCardStrengths(true))

  return evaluatedHands.reduce((acc, curr, index) => {
    return acc + curr.bid * (index + 1)
  }, 0)
}
