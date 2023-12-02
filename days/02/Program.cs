using day_2;

var lines = File.ReadLines("input.txt");
if (lines == null || !lines.Any()) {
    Console.WriteLine("file/lines not found");
    return -1;
}
lines = lines.Select(l => l.Trim().ToLower());

var part1 = Part1.Run(lines);
Console.WriteLine($"Part 1: {part1}");

return 0;
