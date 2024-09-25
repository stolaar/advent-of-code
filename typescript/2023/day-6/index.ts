interface IRace {
  duration: number
  distance: number
}

export const processInput = (input: string[]) => {
  const [timesLine, distancesLine] = input
  const times = timesLine.split(": ")[1].split("  ").filter(Boolean).map(Number)
  const distances = distancesLine
    .split(": ")[1]
    .split("  ")
    .filter(Boolean)
    .map(Number)
  return times.map((time, index) => ({
    duration: time,
    distance: distances[index],
  }))
}

const getNumOfWaysToBeatTheRecord = (duration: number, distance: number) => {
  let results = []
  for (let i = 1; i < duration; i++) {
    if (i * (duration - i) > distance) {
      results.push(i)
    }
  }
  const first = results[0]
  const last = results[results.length - 1]
  return last - first + 1
}

export const partOne = (races: IRace[]) => {
  return races
    .map((race) => getNumOfWaysToBeatTheRecord(race.duration, race.distance))
    .reduce((acc, curr) => acc * curr, 1)
}

export const partTwo = (races: IRace[]) => {
  const race = races.reduce(
    (acc, curr) => {
      return {
        duration: Number(`${acc.duration}${curr.duration}`),
        distance: Number(`${acc.distance}${curr.distance}`),
      }
    },
    { duration: 0, distance: 0 } as IRace,
  )
  return getNumOfWaysToBeatTheRecord(race.duration, race.distance)
}
