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
  const arr = [];

  const LIST_SIZE = 256;

  for (let i = 0; i < LIST_SIZE; i++) {
    arr.push(i);
  }

  let current = 0;
  let skip = 0;

  for (let k = 0; k < times; k++) {
    lengthSequence.forEach(length => {
      perform(arr, current, length);

      current = (current + length + skip) % arr.length;
      skip += 1;
    });
  }

  return arr;
}

module.exports.part1 = function(input) {
  const lengthSequence = input
    .trim()
    .split(",")
    .map(Number);

  return run(lengthSequence, 1);
};

module.exports.knotHash = function(input) {
  const lengthSequence = input
    .split("")
    .map(char => char.charCodeAt(0))
    .concat([17, 31, 73, 47, 23]);

  const out = partition(run(lengthSequence, 64), 16)
    .map(arr => arr.reduce((acc, n) => acc ^ n, 0))
    .map(n => n.toString(16).padStart(2, "0"))
    .join("");

  return out;
};
