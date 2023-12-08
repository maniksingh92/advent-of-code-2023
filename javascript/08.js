import fs from "fs";

const inputs = fs
	.readFileSync("../inputs/08.txt")
	.toString()
	.trim()
	.split("\n");

const pattern = inputs[0];

const dataMap = inputs
	.slice(2)
	.map((line) => {
		const l = line.split(" = (");
		return [l[0], l[1].slice(0, -1).split(", ")];
	})
	.reduce((acc, curr) => {
		acc[curr[0]] = curr[1];
		return acc;
	}, {});

let curr = "AAA";
let steps = 0;
let i = 0;

function solvePart1() {
	while (true) {
		if (curr === "ZZZ") break;

		if (pattern[i] === "L") {
			curr = dataMap[curr][0];
		} else {
			curr = dataMap[curr][1];
		}

		steps += 1;

		i += 1;
		if (i >= pattern.length) {
			i = 0;
		}
	}
}

solvePart1();

console.log(steps);
