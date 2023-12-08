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

const gcd = (a, b) => (a ? gcd(b % a, a) : b);
const lcm = (a, b) => (a * b) / gcd(a, b);

function solve() {
	const startingNodes = Object.keys(dataMap).filter((node) =>
		node.endsWith("A"),
	);

	const solutions = [];

	for (let curr of startingNodes) {
		let steps = 0;
		let i = 0;
		while (true) {
			if (curr.endsWith("Z")) break;

			const move = pattern[i] === "L" ? 0 : 1;
			curr = dataMap[curr][move];

			steps += 1;

			i += 1;
			if (i >= pattern.length) {
				i = 0;
			}
		}
		solutions.push(steps);
	}

	console.log(solutions.reduce(lcm));
}

solve();
