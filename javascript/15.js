import fs from "fs";

const inputs = fs.readFileSync("../inputs/15.txt").toString().trim();
const words = inputs.split(",");

function getHashValue(word) {
    let value = 0;
    for (const ch of word) {
        value += ch.charCodeAt(0);
        value *= 17;
        value %= 256;
    }
    return value;
}

const boxes = Array.from(Array(256), () => []);
const lenses = {};

for (const word of words) {
    if (word.endsWith("-")) {
        const [label,] = word.split("-");

        if (!(label in lenses)) continue;

        boxes[lenses[label][0]] = boxes[lenses[label][0]].filter(lens => lens !== label);
        delete lenses[label];
    } else {
        const [label, value] = word.split("=");

        if (!(label in lenses)) {
            const hashValue = getHashValue(label);
            boxes[hashValue].push(label);
            lenses[label] = [hashValue, Number(value)];
        }

        lenses[label][1] = value;
    }
}

let sum = 0;
for (const [b, box] of boxes.entries()) {
    for (const [l, label] of box.entries()) {
        sum += (1 + b) * (1 + l) * (lenses[label][1]);
    }
}
console.log(sum)