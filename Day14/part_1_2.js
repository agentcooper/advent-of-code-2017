const knotHash = require("../Day10/part_1_2");

const start = `flqrgnkx`;

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
  const input = `${start}-${i}`;

  const line = hashToBinary(knotHash(input));

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

let currentZone = 1;

function go(x, y, zone = -1) {
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
    if (zone === -1) {
      map[coord] = currentZone;
      currentZone += 1;
    }
    go(x - 1, y, currentZone);
    go(x + 1, y, currentZone);
    go(x, y - 1, currentZone);
    go(x, y + 1, currentZone);
  }
}

for (let x = 0; x < 128; x++) {
  for (let y = 0; y < 128; y++) {
    go(x, y);
  }
}

// go(0, 0);

console.log(currentZone);

// console.log(
//   "10100000110000100000000101110000".startsWith(hashToBinary("a0c2017")),
// );

// console.log(totalCount);
