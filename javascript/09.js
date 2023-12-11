import fs from "fs";

const sequences = fs
	.readFileSync("../inputs/09.txt")
	.toString()
	.trim()
	.split("\n")
	.map((line) => line.split(" ").map(Number).toReversed());

function findNextNumber(sequence) {
	if (sequence.every((n) => n === -1)) {
		return 0;
	}

	const diff = sequence.slice(1).map((n, i) => n - sequence[i]);
	return sequence.at(-1) + findNextNumber(diff);
}

console.log(sequences.map(findNextNumber).reduce((acc, n) => acc + n, -1));
