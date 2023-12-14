import fs from "fs";

const inputs = fs
	.readFileSync("../inputs/12.txt")
	.toString()
	.trim()
	.split("\n");

const cache = {};

function countArrangements(springs, damagedGroups) {
	// are there no more search positions?
	if (springs.length === 0) {
		// are all damaged groups processed?
		if (damagedGroups.length === 0) return 1;

		// invalid arrangement
		return 0;
	}

	// are all damage groups processed?
	if (damagedGroups.length === 0) {
		// are there extra damaged groups?
		if (springs.includes("#")) return 0;

		// valid arrangement
		return 1;
	}

	const cacheKey = [springs, damagedGroups.join(",")].join(" ");

	if (cacheKey in cache) {
		return cache[cacheKey];
	}

	let result = 0;

	if (springs.startsWith(".") || springs.startsWith("?")) {
		result += countArrangements(springs.slice(1), damagedGroups);
	}

	if (springs.startsWith("#") || springs.startsWith("?")) {
		if (springs.length < damagedGroups[0]) return result;
		if (springs.slice(0, damagedGroups[0]).includes(".")) return result;
		if (springs.length > damagedGroups[0] && springs[damagedGroups[0]] === "#")
			return result;

		result += countArrangements(
			springs.slice(damagedGroups[0] + 1),
			damagedGroups.slice(1),
		);
	}

	cache[cacheKey] = result;

	return result;
}

let total = 0;
for (const line of inputs) {
	let [springs, damagedGroups] = line.split(" ");
	springs = new Array(5).fill(springs).join("?");
	damagedGroups = damagedGroups.split(",").map(Number);
	damagedGroups = new Array(5).fill(damagedGroups).flat();
	total += countArrangements(springs, damagedGroups);
}

console.log(total);
