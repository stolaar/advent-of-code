# Advent of Code 🎄

Welcome to my **Advent of Code** solutions repository! This repository contains my solutions to the annual [Advent of Code](https://adventofcode.com/) programming challenge. Each year, Advent of Code provides a fun and challenging set of puzzles for the 25 days leading up to Christmas.

---

## 📂 Directory Structure

The solutions are organized by programming language and year, with each day having its own subdirectory.

```
advent-of-code/
├── golang/ # Solutions written in Go
│ ├── 2023/ # Solutions for the 2023 Advent of Code
│ │ ├── day01/
│ │ ├── day02/
│ │ └── ...
│ └── ...
├── typescript/ # Solutions written in TypeScript
│ ├── 2023/
│ │ ├── day01/
│ │ ├── day02/
│ │ └── ...
│ └── ...
└── ...
```

### Language Details

- **Go (Golang):** Solutions can be found in the `golang` directory.
- **TypeScript:** Solutions can be found in the `typescript` directory.

Each day's directory typically includes:
- The solution code for the day.
- Input files (if applicable).
- Any additional notes or files required for the solution.
---

## 🛠️ Usage

### Running Go Solutions
1. Build the cli
```bash
go build -o aoc main.go
```

2. Run the solution using:

```bash
# For running day 1 from year 2022
aoc 1 2022
```

```bash
# For running day 1 the current year
aoc 1
```

3. Create new day
```
# generates day 1 for the current year
aoc generate 1

# generates day 1 for year 2023
aoc generate 1 2023
```

### Running TypeScript Solutions

Navigate to the `typescript` directory

Install dependencies if required:
```
yarn
```

Run the solution using:
```
yarn start 1
```

Initialize new day files
```
yarn generate 1
```

🚀 Getting Started

Clone the repository:
```
git clone https://github.com/stolaar/advent-of-code.git
```

Navigate to the desired language and year folder.
Run the solutions as described above.

📝 Notes
Advent of Code Website: https://adventofcode.com/
Solutions are tailored to the inputs provided for my account; they may not work directly with other inputs.
If you find any issues or want to suggest improvements, feel free to open an issue or a pull request.

📜 License
This repository is licensed under the MIT License.

Happy coding! 🎅

Feel free to modify the details to match your specific setup.






