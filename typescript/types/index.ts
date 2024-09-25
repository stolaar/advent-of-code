export interface ISolution {
    processInput?: <T,>(input: string[]) => T
    partOne: (args: string[]) => string
    partTwo: (args: string[]) => string
}
