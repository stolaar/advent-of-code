import path from "path";
import * as fs from "fs";
import { ISolution } from "./types";

const [, , day] = process.argv;

if (!Number.isInteger(parseInt(day))) {
  throw new Error("Please provide the day");
}

const readMultilineInput = (inputPath: string) => {
  const result = fs.readFileSync(inputPath, "utf-8").split("\n");
  result.pop();

  return result;
};

(async () => {
  const modulePath = path.resolve(__dirname, `day-${day}`);
  const dayModule: ISolution = await import(
    path.resolve(modulePath, "index.ts")
  );
  const input = readMultilineInput(path.resolve(modulePath, "input.txt"));

  console.time("Input processing");
  const args: string[] = dayModule.processInput?.(input) || input;
  console.timeEnd("Input processing");

  const partOneStart = performance.now();
  const solutionPartOne = dayModule.partOne(args);
  const partOneEnd = performance.now();
  console.log("Part one solution", solutionPartOne);
  console.log(
    `Part one execution time ${Math.round(partOneEnd - partOneStart)}ms`,
  );

  const partTwoStart = performance.now();
  const solutionPartTwo = dayModule.partTwo(args);
  const partTwoEnd = performance.now();
  console.log("Part two solution", solutionPartTwo);
  console.log(
    `Part two execution time ${Math.round(partTwoEnd - partTwoStart)}ms`,
  );
})();
