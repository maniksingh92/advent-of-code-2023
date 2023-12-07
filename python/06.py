import math

def quadratic(a, t, d):
    m = math.sqrt(t*t - 4*a*d)

    l = math.ceil((t - m)/(2*a))
    u = math.floor((t + m)/(2*a))

    print(t, d,l, u, u - l + 1)
    if l * (t-l) == d:
        l += 1

    if u * (t-u) == d:
        u -= 1
    print(t, d,l, u, u - l + 1)
    return l, u 

with open("../inputs/06.txt") as inputs:
    lines = inputs.readlines()

    times = map(int, lines[0].split(":")[1].split())
    distances = map(int, "".join(lines[1].split(":")[1].split()))

    time_distances = zip(times, distances)

    sum = 1
    for t, d in time_distances:
        l, u = quadratic(1, t, d)
        if l < u:
            sum *= u - l + 1
    print(sum)
