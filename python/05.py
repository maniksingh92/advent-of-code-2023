from collections.abc import Iterator
from typing import List

def day_05_puzzle_1(inputs: Iterator[str]):
    seeds = list(map(int, next(inputs).rstrip().split(": ")[1].split()))

    dataMaps: List[List[List[int]]] = []
    for line in inputs:
        line = line.rstrip()

        if len(line) == 0:
            dataMaps.append([])
            continue

        if line.endswith("map:"):
            continue

        dataMaps[-1].append(list(map(int, line.split())))

    locations: List[int] = []
    for seed in seeds:
        curr = seed
        for dataMap in dataMaps:
            for data in dataMap:
                destination, source, length = data
                if source <= curr <= source + length:
                    curr = destination + curr - source
                    break
        locations.append(curr)

    print(min(locations))

if __name__ == "__main__":
    with open("../inputs/05.txt") as inputs:
        day_05_puzzle_1(inputs)
