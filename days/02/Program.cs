using day_2;

var lines = File.ReadLines("input.txt");
if (lines == null || !lines.Any()) {
    Console.WriteLine("file/lines not found");
    return -1;
}
lines = lines.Select(l => l.Trim().ToLower());

var part1 = Part1.Run(lines);
var part2 = Part2.Run(lines);

Console.WriteLine($"{new string('=', 10)} SOLUTIONS {new string('=', 10)}");
Console.WriteLine($"Part 1: {part1}");
Console.WriteLine($"Part 2: {part2}");

return 0;
