const part1Config : Record<string, number>= {
    green: 13,
    red: 12,
    blue: 14,
}

export const partOne = (input: string[]) => {
    return input.filter((line) => {
       const [, results] = line.split(': ')
         const reveals = results.split('; ')
        return reveals.every((reveal) => {
              const revealedColor= reveal.split(', ')
            return revealedColor.every((color) => {
                const [colorValue, colorName] = color.split(' ')
                return (part1Config[colorName] ?? 0) >= parseInt(colorValue)
            })
        })
    }).reduce((acc, line) => {
        const [name] = line.split(': ')
        return acc + parseInt(name.split(' ').pop() ?? '0')
    }, 0)
}

export const partTwo = (input: string[]) => {
    return input.reduce((result, line) => {
       const fewestColors= (line.split(': ').pop() as string).split('; ').reduce((acc, set) => {
          const setColors = set.split(', ')
           setColors.forEach((color) => {
            const [value, colorName] = color.split(' ')
               if(acc[colorName] < parseInt(value)) {
                   acc[colorName] = parseInt(value)
               }
           })
            return acc
       }, {red: 0, green: 0, blue: 0} as Record<string, number>)
        const powerOfFewestColors = Object.values(fewestColors).reduce((acc, value) => acc * value, 1)
            return result + powerOfFewestColors
    }, 0)
}
