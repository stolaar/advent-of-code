interface IHand {
  cards: string[]
  bid: number
}

interface IHandValue {
  firstEvalStrength: number
  secondEvalStrengths: number[]
}

const cardStrengthsMap: Record<string | number, number> = {
  A: 14,
  K: 13,
  Q: 12,
  J: 11,
  T: 10,
}

const evaluateHand = (cards: string[]): IHandValue => {
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
    secondEvalStrengths: cards.map(
      ([card, count]) => cardStrengthsMap[card] || +card,
    ),
  }
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

  evaluatedHands.sort((a, b) => {
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
  })

  return evaluatedHands.reduce((acc, curr, index) => {
    return acc + curr.bid * (index + 1)
  }, 0)
}

export const partTwo = (input: string[]) => {
  return ""
}
