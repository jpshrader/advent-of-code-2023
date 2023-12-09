
var lines = File.ReadLines("input.txt").ToList();

var times = lines[0]
    .Replace("Time:", "")
    .Split(" ")
    .Where(s => !string.IsNullOrWhiteSpace(s))
    .Select(int.Parse).ToList();

var distances = lines[1]
    .Replace("Distance:", "")
    .Split(" ")
    .Where(s => !string.IsNullOrWhiteSpace(s))
    .Select(int.Parse).ToList();

var races = times.Select((time, idx) => new Race {
    Time = time,
    Distance = distances[idx]
}).ToList();

foreach (var race in races) {
    Console.WriteLine($"race: {race.Time} | {race.Distance}");
}

return 0;

class Race {
    public required int Time { get; set; }

    public required int Distance { get; set; }

    public static IEnumerable<int> GetWinConditions() {
        return Enumerable.Empty<int>();
    }
}