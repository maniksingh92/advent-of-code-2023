import fs from "fs";

const inputs = fs.readFileSync("../inputs/15.txt").toString().trim();

const words = inputs.split(",");

let sum = 0;

for (const word of words) {
    let value = 0;
    for (const ch of word) {
        value += ch.charCodeAt(0);
        value *= 17;
        value %= 256;
    }
    sum += value;
}

console.log(sum);