function grab(arr, current, length) {
  let i = 0;

  const out = [];

  while (i < length) {
    out.push(arr[(current + i) % arr.length]);
    i++;
  }

  return out;
}

function perform(arr, current, length) {
  let i = current;

  grab(arr, current, length)
    .reverse()
    .forEach(item => {
      arr[i] = item;

      i = (i + 1) % arr.length;
    });
}

function partition(arr, partitionSize) {
  if (arr.length === 0) {
    return [];
  }

  return [
    arr.slice(0, partitionSize),
    ...partition(arr.slice(partitionSize), partitionSize),
  ];
}

function run(lengthSequence, times) {
  for (let k = 0; k < times; k++) {
    lengthSequence.forEach(length => {
      perform(arr, current, length);

      current = (current + length + skip) % arr.length;
      skip += 1;
    });
  }
}

module.exports = function knotHash(input) {
  const arr = [];

  const LIST_SIZE = 256;

  for (let i = 0; i < LIST_SIZE; i++) {
    arr.push(i);
  }

  let current = 0;
  let skip = 0;

  const lengthSequence = input
    .split("")
    .map(char => char.charCodeAt(0))
    .concat([17, 31, 73, 47, 23]);

  for (let k = 0; k < 64; k++) {
    lengthSequence.forEach(length => {
      perform(arr, current, length);

      current = (current + length + skip) % arr.length;
      skip += 1;
    });
  }

  const out = partition(arr, 16)
    .map(arr => arr.reduce((acc, n) => acc ^ n, 0))
    .map(n => n.toString(16).padStart(2, "0"))
    .join("");

  return out;
};

// if (process.argv[2] !== "2") {
//   const part1LengthSequence = [
//     230,
//     1,
//     2,
//     221,
//     97,
//     252,
//     168,
//     169,
//     57,
//     99,
//     0,
//     254,
//     181,
//     255,
//     235,
//     167,
//   ];

//   run(part1LengthSequence, 1);
//   console.log(arr[0] * arr[1]);
// } else {
//   const part2Input = `230,1,2,221,97,252,168,169,57,99,0,254,181,255,235,167`;

//   console.log(knotHash(part2Input));
// }