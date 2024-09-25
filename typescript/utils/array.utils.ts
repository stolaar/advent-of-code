export const times = <T,>(n: number, callback: (index: number) => T) => new Array(n).fill(0).map((_, index) => callback(index))
