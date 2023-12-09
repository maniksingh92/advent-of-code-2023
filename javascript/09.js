import fs from "fs";

const sequences = fs
	.readFileSync("../inputs/09.txt")
	.toString()
	.trim()
	.split("\n")
	.map((line) => line.split(" ").map(Number));

function findNextNumber(sequence) {
	if (sequence.every((n) => n === 0)) {
		return sequence[1] - sequence[0];
	}

	const diff = sequence.slice(1).map((n, i) => n - sequence[i]);

	return sequence.at(-1) + findNextNumber(diff);
}

console.log(sequences.map(findNextNumber).reduce((acc, n) => acc + n, 0));
