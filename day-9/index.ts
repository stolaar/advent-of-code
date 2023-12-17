export const processInput = (input: string[]) => {
  return input.map((line) =>
    line
      .split(/(\s+)/)
      .filter((x) => {
        return x !== " "
      })
      .map((x) => Number(x)),
  )
}

const reduceSequence = (
  sequences: number[][],
  current: number[],
): number[][] => {
  const lastSequence = sequences[sequences.length - 1]
  if (lastSequence) {
    const allZeros = lastSequence.every((x) => x === 0)
    if (allZeros) {
      lastSequence.push(0)
      return sequences
    }
  }

  const nextSequence = current.reduce((x, current, currentIndex, self) => {
    if (currentIndex === self.length - 1) {
      return x
    }
    const next = self[currentIndex + 1]
    const diff = next - current
    x.push(diff)
    return x
  }, [] as number[])

  sequences.push(nextSequence)
  return reduceSequence(sequences, nextSequence)
}

export const partOne = (input: number[][]) => {
  return input
    .map((line) => {
      let sequences = [line].reduce(reduceSequence, [line])
      for (let x = sequences.length - 1; x > 0; x--) {
        const previous = sequences[x - 1]
        const lastValue = sequences[x][sequences[x].length - 1]

        previous.push(lastValue + previous[previous.length - 1])
      }

      const first = sequences[0]
      return first[first.length - 1]
    })
    .reduce((x, y) => x + y, 0)
}

export const partTwo = (input: number[][]) => {
  return input
    .map((line) => {
      let sequences = [line].reduce(reduceSequence, [line])
      for (let x = sequences.length - 1; x > 0; x--) {
        const previous = sequences[x - 1]
        const firstValue = sequences[x][0]

        previous.unshift(previous[0] - firstValue)
      }

      const first = sequences[0]
      return first[0]
    })
    .reduce((x, y) => x + y, 0)
}
