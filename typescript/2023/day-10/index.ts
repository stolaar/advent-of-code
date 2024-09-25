const breakCondition = (start: ISurface, check: ISurface) => {
  return start.x === check.x && start.y === check.y
}
const verticalPipe = (surface: ISurface, x: number, y: number) => {
  return (
    ((surface.x === x && surface.y === y + 1) ||
      (surface.x === x && surface.y === y - 1)) &&
    surface.type === "pipe"
  )
}

const horizontalPipe = (surface: ISurface, x: number, y: number) => {
  return (
    ((surface.x === x + 1 && surface.y === y) ||
      (surface.x === x - 1 && surface.y === y)) &&
    surface.type === "pipe"
  )
}

const lPipe = (surface: ISurface, x: number, y: number) => {
  return (
    ((surface.x === x + 1 && surface.y === y) ||
      (surface.x === x && surface.y === y + 1)) &&
    surface.type === "pipe"
  )
}

const jPipe = (surface: ISurface, x: number, y: number) => {
  return (
    ((surface.x === x - 1 && surface.y === y) ||
      (surface.x === x && surface.y === y + 1)) &&
    surface.type === "pipe"
  )
}

const sevenPipe = (surface: ISurface, x: number, y: number) => {
  return (
    ((surface.x === x - 1 && surface.y === y) ||
      (surface.x === x && surface.y === y - 1)) &&
    surface.type === "pipe"
  )
}

const fPipe = (surface: ISurface, x: number, y: number) => {
  return (
    ((surface.x === x + 1 && surface.y === y) ||
      (surface.x === x && surface.y === y - 1)) &&
    surface.type === "pipe"
  )
}

const isVisited = (surface: ISurface) => (that: ISurface) =>
  breakCondition(that, surface)
const getConnectedPipes = (
  x: number,
  y: number,
  pipe: string,
  input: ISurface[],
  visited?: ISurface[],
): ISurface[] => {
  if (
    visited?.length &&
    visited.some((el) => breakCondition(el, { x, y, pipe, type: "pipe" }))
  ) {
    return []
  }
  const nextVisited: ISurface[] = visited
    ? [...visited, { x, y, pipe, type: "pipe" }]
    : [{ x, y, pipe, type: "pipe" }]

  if (pipe === "|") {
    return input
      .filter(
        (surface) =>
          verticalPipe(surface, x, y) && !nextVisited?.some(isVisited(surface)),
      )
      .flatMap((surface) => [
        surface,
        ...getConnectedPipes(
          surface.x,
          surface.y,
          surface.pipe ?? "",
          input,
          nextVisited,
        ),
      ])
  }
  if (pipe === "-") {
    return input
      .filter(
        (surface) =>
          horizontalPipe(surface, x, y) &&
          !nextVisited?.some(isVisited(surface)),
      )
      .flatMap((surface) => [
        surface,
        ...getConnectedPipes(
          surface.x,
          surface.y,
          surface.pipe ?? "",
          input,
          nextVisited,
        ),
      ])
  }

  if (pipe === "L") {
    return input
      .filter(
        (surface) =>
          lPipe(surface, x, y) && !nextVisited?.some(isVisited(surface)),
      )
      .flatMap((surface) => [
        surface,
        ...getConnectedPipes(
          surface.x,
          surface.y,
          surface.pipe ?? "",
          input,
          nextVisited,
        ),
      ])
  }

  if (pipe === "J") {
    return input
      .filter(
        (surface) =>
          jPipe(surface, x, y) && !nextVisited?.some(isVisited(surface)),
      )
      .flatMap((surface) => [
        surface,
        ...getConnectedPipes(
          surface.x,
          surface.y,
          surface.pipe ?? "",
          input,
          nextVisited,
        ),
      ])
  }

  if (pipe === "7") {
    return input
      .filter(
        (surface) =>
          sevenPipe(surface, x, y) && !nextVisited?.some(isVisited(surface)),
      )
      .flatMap((surface) => [
        surface,
        ...getConnectedPipes(
          surface.x,
          surface.y,
          surface.pipe ?? "",
          input,
          nextVisited,
        ),
      ])
  }

  if (pipe === "F") {
    return input
      .filter(
        (surface) =>
          fPipe(surface, x, y) && !nextVisited?.some(isVisited(surface)),
      )
      .flatMap((surface) => [
        surface,
        ...getConnectedPipes(
          surface.x,
          surface.y,
          surface.pipe ?? "",
          input,
          nextVisited,
        ),
      ])
  }

  return []
}

interface ISurface {
  x: number
  y: number
  type: "pipe" | "ground" | "start"
  pipe?: string
  connectedPipes?: ISurface[]
}

export const processInput = (input: string[]) => {
  return input.flatMap((line, y) => {
    const row = line.split("")
    return row.map((cell, x) => {
      const isPipe = cell !== "." && cell !== "S"
      const isStart = cell === "S"
      return {
        x,
        y,
        type: !isPipe ? (isStart ? "start" : "ground") : "pipe",
        ...(isPipe && {
          pipe: cell,
        }),
      }
    })
  })
}

export const partOne = (surfaces: ISurface[]) => {
  const start = surfaces.find((surface) => surface.type === "start")
  if (!start) {
    throw new Error("No start found")
  }
  const connections = surfaces.filter((surface, index) => {
    const { x, y, pipe } = surface
    return (
      ((x === start?.x + 1 && y === start.y) ||
        (x === start?.x - 1 && y === start.y) ||
        (y === start?.y + 1 && x === start.x) ||
        (y === start?.y - 1 && x === start.x)) &&
      surface.type === "pipe"
    )
  })

  const connected = connections.flatMap((surface) => ({
    ...surface,
    connectedPipes: getConnectedPipes(
      surface.x,
      surface.y,
      surface.pipe ?? "",
      surfaces,
    ),
  }))

  connected.forEach((surface) => {
    console.log(surface.x, surface.y, surface.connectedPipes)
  })

  const farthestPoints = connected.map((surface, index, self) => {
    const intersectionIndexes = self
      .filter((el, elIndex) => elIndex !== index)
      .map(
        (el) =>
          surface.connectedPipes?.findIndex(
            (el2) => el.x === el2.x && el.y === el2.y,
          ),
      )

    return Math.max(...intersectionIndexes)
  })

  return Math.max(...connected.map((el) => el.connectedPipes?.length ?? 0)) + 1
}

export const partTwo = (input: string[]) => {
  return ""
}
