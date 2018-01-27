const { readFileSync } = require("fs");
const path = require("path");

const { knotHash } = require("../Day10_JS/lib");

const input = readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

let totalCount = 0;

function hashToBinary(hash) {
  const line = hash
    .split("")
    .map(c =>
      parseInt(c, 16)
        .toString(2)
        .padStart(4, "0"),
    )
    .join("");

  return line;
}

const map = {};

for (let i = 0; i < 128; i++) {
  const s = `${input}-${i}`;
  const line = hashToBinary(knotHash(s));

  if (line.length !== 128) {
    throw new Error("WTF");
  }

  const count = line.split("").reduce((acc, c) => {
    return acc + (c === "1" ? 1 : 0);
  }, 0);

  line
    .split("")
    .map(c => {
      if (c === "1") {
        return "#";
      }
      return ".";
    })
    .forEach((c, j) => {
      map[`x${j}y${i}`] = c;
    });

  totalCount += count;
}

let currentZone = 0;

function go(x, y, zone) {
  const coord = `x${x}y${y}`;

  if (typeof map[coord] === "undefined") {
    return;
  }

  if (map[coord] === ".") {
    return;
  }

  if (typeof map[coord] === "number") {
    return;
  }

  if (map[coord] === "#") {
    let z;
    if (typeof zone === "undefined") {
      z = currentZone;
      currentZone += 1;
    } else {
      z = zone;
    }
    map[coord] = z;
    go(x - 1, y, z);
    go(x + 1, y, z);
    go(x, y - 1, z);
    go(x, y + 1, z);
  }
}

for (let x = 0; x < 128; x++) {
  for (let y = 0; y < 128; y++) {
    go(x, y);
  }
}

console.assert(currentZone === 1180, "Part 1 is incorrect!");
console.assert(totalCount === 8148, "Part 2 is incorrect!");
