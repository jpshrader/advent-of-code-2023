var lines = File.ReadLines("input.txt");
if (lines == null || !lines.Any()) {
    Console.WriteLine("file/lines not found");
    return -1;
}
lines = lines.Select(l => l.Trim().ToLower());

var part1 = Part1(lines);
var part2 = Part2(lines);

Console.WriteLine($"{new string('=', 10)} SOLUTIONS {new string('=', 10)}");
Console.WriteLine($"Part 1: {part1}");
Console.WriteLine($"Part 2: {part2}");

return 0;

static int Part1(IEnumerable<string> lines) {
    var numbers = new List<int>();
    foreach (var line in lines) {
        var lineNums = new List<int>();
        foreach (var c in line) {
            if (int.TryParse(c.ToString(), out var number)) {
                lineNums.Add(number);
            }
        }
        if (lineNums.Count == 0) {
            continue;
        }
        numbers.Add((lineNums.First() * 10) + lineNums.Last());
    }

    return numbers.Sum();
}

static int Part2(IEnumerable<string> lines) {
    var numberMap = new Dictionary<string, int>{
        {"one", 1},
        {"two", 2},
        {"three", 3},
        {"four", 4},
        {"five", 5},
        {"six", 6},
        {"seven", 7},
        {"eight", 8},
        {"nine", 9},
    };
    var numbers = new List<int>();
    foreach (var line in lines) {
        var lineNums = new List<(int index, int digit)>();
        foreach (var (word, number) in numberMap) {
            var idx = 0;
            while ((idx = IndexOf(line, word, idx)) != -1) {
                lineNums.Add((idx, number));
            }
        }
        for (var i = 0; i < line.Length; i++) {
            if (int.TryParse(line[i].ToString(), out var number)) {
                lineNums.Add((i, number));
            }
        }
        if (lineNums.Count == 0) {
            continue;
        }
        lineNums.Sort((a, b) => a.index.CompareTo(b.index));

        Console.WriteLine($"{line} => {string.Join(", ", lineNums.Select(n => n.digit))}");

        numbers.Add((lineNums.First().digit * 10) + lineNums.Last().digit);
    }

    return numbers.Sum();
}

static int IndexOf(string line, string word, int startIndex = 0) {
    var index = line.IndexOf(word, startIndex);
    if (index == -1) {
        return -1;
    }
    return index + word.Length;
}