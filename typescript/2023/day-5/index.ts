interface IGardeningMapEntry {
  sourceStart: number
  destinationStart: number
  rangeLength: number
  sourceEnd: number
}

export interface IGardeningMap {
  entries: IGardeningMapEntry[]
  get: (value: number) => number
}

interface Interval {
  start: number
  end: number
}

interface OffsetInterval extends Interval {
  offset: number
}

const createGardeningMap = (values: string[]): IGardeningMap => {
  const entries = values.map((entry) => {
    const [destinationStart, sourceStart, rangeLength] = entry
      .split(" ")
      .map(Number)
    return {
      sourceStart,
      destinationStart,
      rangeLength,
      sourceEnd: sourceStart + rangeLength - 1,
    }
  })
  return {
    entries,
    get: (value: number) => {
      const fromEntry = entries.find(
        (entry) =>
          entry.sourceStart <= value &&
          value <= entry.rangeLength + entry.sourceStart,
      )
      if (fromEntry) {
        const sourcePosition = value - fromEntry.sourceStart
        return fromEntry.destinationStart + sourcePosition
      }
      return value
    },
  }
}

interface IInput {
  seeds: number[]
  seedToSoilMap: IGardeningMap
  soilToFertilizerMap: IGardeningMap
  fertilizerToWaterMap: IGardeningMap
  waterToLightMap: IGardeningMap
  lightToTemperatureMap: IGardeningMap
  temperatureToHumidityMap: IGardeningMap
  humidityToLocationMap: IGardeningMap
}

const inputMapKeys: (keyof Omit<IInput, "seeds">)[] = [
  "seedToSoilMap",
  "soilToFertilizerMap",
  "fertilizerToWaterMap",
  "waterToLightMap",
  "lightToTemperatureMap",
  "temperatureToHumidityMap",
  "humidityToLocationMap",
]

export const processInput = (input: string[]): IInput => {
  return input
    .join("\n")
    .split("\n\n")
    .reduce((acc, curr, currentIndex) => {
      if (currentIndex === 0) {
        const seedsMap = (curr.split(": ").pop() ?? "").split(" ").map(Number)
        return {
          ...acc,
          seeds: seedsMap,
        }
      }
      const values = curr.split("\n").slice(1)
      acc[inputMapKeys[currentIndex - 1]] = createGardeningMap(values)

      return acc
    }, {} as IInput)
}

const getLocation = (seed: number, input: IInput) => {
  const soil = input.seedToSoilMap.get(seed)
  const fertilizer = input.soilToFertilizerMap.get(soil!) ?? soil
  const water = input.fertilizerToWaterMap.get(fertilizer!) ?? fertilizer
  const light = input.waterToLightMap.get(water!) ?? water
  const temperature = input.lightToTemperatureMap.get(light!) ?? light
  const humidity =
    input.temperatureToHumidityMap.get(temperature!) ?? temperature
  const location = input.humidityToLocationMap.get(humidity!) ?? humidity

  return location ?? 0
}

export const partOne = (input: IInput) => {
  return Math.min(
    ...input.seeds.map((seed) => {
      return getLocation(seed, input) ?? 0
    }),
  )
}

const processMapperInterval = (
  interval: Interval,
  mapper: OffsetInterval[],
) => {
  const current = { ...interval }
  const result = [] as [Interval, Interval][]
  for (const mapping of mapper) {
    if (mapping.end < current.start) {
      continue
    }

    const end = Math.min(mapping.end, current.end)
    result.push([
      { start: current.start, end },
      { start: current.start + mapping.offset, end: end + mapping.offset },
    ])
    current.start = end + 1

    if (current.start > current.end) {
      break
    }
  }

  return result.map(([_, mapped]) => mapped)
}

const processSeedIntervals = (seed: Interval, mappers: OffsetInterval[][]) => {
  let result = [seed]
  for (const mapper of mappers) {
    const newResult = [] as Interval[]
    for (const interval of result) {
      newResult.push(...processMapperInterval(interval, mapper))
    }
    result = newResult
  }

  return result
}

export const partTwo = (input: IInput) => {
  const groupedInputs = input.seeds.reduce((acc, curr) => {
    if (!acc.length) {
      acc.push({ start: curr, end: 0 })
      return acc
    }
    const last = acc[acc.length - 1]
    if (!last.end) {
      last.end = last.start + curr - 1
      return acc
    }
    acc.push({ start: curr, end: 0 })
    return acc
  }, [] as Interval[])

  const mapperIntervals: OffsetInterval[][] = [
    input.seedToSoilMap.entries,
    input.soilToFertilizerMap.entries,
    input.fertilizerToWaterMap.entries,
    input.waterToLightMap.entries,
    input.lightToTemperatureMap.entries,
    input.temperatureToHumidityMap.entries,
    input.humidityToLocationMap.entries,
  ].map((mapper) => {
    const offsetIntervals = mapper.map((entry) => ({
      start: entry.sourceStart,
      end: entry.sourceStart + entry.rangeLength - 1,
      offset: entry.destinationStart - entry.sourceStart,
    }))
    offsetIntervals.sort((a, b) => a.start - b.start)
    if (offsetIntervals[0].start !== 0) {
      offsetIntervals.unshift({
        start: 0,
        end: offsetIntervals[0].start - 1,
        offset: 0,
      })
    }
    for (let i = 0; i < offsetIntervals.length - 1; i++) {
      const current = offsetIntervals[i]
      const next = offsetIntervals[i + 1]
      if (current.end + 1 !== next.start) {
        offsetIntervals.splice(i + 1, 0, {
          start: current.end + 1,
          end: next.start - 1,
          offset: 0,
        })
      }
    }

    const last = offsetIntervals[offsetIntervals.length - 1]
    offsetIntervals.push({
      start: last.end + 1,
      end: Number.MAX_SAFE_INTEGER,
      offset: 0,
    })
    return offsetIntervals
  })

  const mappedIntervals = groupedInputs.flatMap((seed) =>
    processSeedIntervals(seed, mapperIntervals),
  )

  return Math.min(...mappedIntervals.map((interval) => interval.start))
}
