const { knotHash, part1 } = require("./lib");
const { readFileSync } = require("fs");
const path = require("path");

const input = readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

const arr = part1(input);
const part1Solution = arr[0] * arr[1];
console.assert(
  part1Solution === 2928,
  "Part 1 solution is incorrect: %s",
  part1Solution,
);

// part 2
const hash = knotHash(input);
console.assert(
  hash === "0c2f794b2eb555f7830766bf8fb65a16",
  "Hash is incorrect: %s",
  hash,
);
