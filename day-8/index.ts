interface IGraph {
  addVertex: (v: string) => void
  addEdge: (v: string, w: string) => void
  getVertex: (v: string) => string[]
  getVertices: () => string[]
}

interface IMapDocument {
  instructions: string
  graph: IGraph
  start: string
}

const graph = (): IGraph => {
  const adjList = new Map()

  return {
    addVertex: (v: string) => {
      adjList.set(v, [])
    },
    addEdge: (v: string, w: string) => {
      adjList.get(v).push(w)
    },
    getVertex: (v: string) => adjList.get(v),
    getVertices: () => [...adjList.keys()],
  }
}
const gcd = (a: number, b: number): number => (a ? gcd(b % a, a) : b)

const lcm = (a: number, b: number) => (a * b) / gcd(a, b)

export const processInput = (input: string[]) => {
  const [instructions, nodesMap] = input.join("\n").split("\n\n")
  const g = graph()
  let start = ""
  nodesMap
    .split("\n")
    .map((node) => {
      const [name, children] = node.split(" = ")
      if (!start) {
        start = name
      }
      g.addVertex(name)
      return { name, children: children.replace(/[()]/g, "").split(", ") }
    })
    .forEach(({ name, children }) => {
      children.forEach((child) => g.addEdge(name, child))
    })
  return { instructions: instructions, graph: g, start }
}

export const partOne = (mapDocuments: IMapDocument) => {
  let instructionIndex = 0
  let iterations = 0
  let currentVertex = "AAA"

  if (!mapDocuments.graph.getVertex(currentVertex)) {
    return 0
  }

  while (currentVertex !== "ZZZ") {
    const [left, right] = mapDocuments.graph.getVertex(currentVertex)
    currentVertex =
      mapDocuments.instructions[instructionIndex] === "L" ? left : right
    iterations++
    instructionIndex =
      instructionIndex + 1 < mapDocuments.instructions.length
        ? instructionIndex + 1
        : 0
  }
  return iterations
}

export const partTwo = (mapDocuments: IMapDocument) => {
  const { graph, instructions } = mapDocuments

  let currentVertices = graph
    .getVertices()
    .filter((v) => v.endsWith("A"))
    .map((v) => ({ vertex: v, i: 0 }))

  let iterations = 0
  let ended = false
  let instructionIndex = 0

  while (!ended) {
    let allEnd = true
    currentVertices = currentVertices.map(({ vertex, i }) => {
      const [left, right] = graph.getVertex(vertex)
      const next = instructions[instructionIndex] === "L" ? left : right
      if (!vertex.endsWith("Z")) {
        allEnd = false
      } else {
        return {
          vertex,
          i,
        }
      }
      return {
        vertex: next,
        i: i + 1,
      }
    })

    iterations++
    if (allEnd) {
      ended = true
      break
    }
    instructionIndex =
      instructionIndex + 1 < instructions.length ? instructionIndex + 1 : 0
  }

  return currentVertices.reduce((acc, { i }) => lcm(acc, i), 1)
}
