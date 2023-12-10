
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
    Id = idx + 1,
    Time = time,
    Distance = distances[idx]
}).ToList();

var winConditions = new List<double>();
foreach (var race in races) {
    winConditions.Add(race.GetWinConditions());
}

var marathonTime = double.Parse(races.Select(r => r.Time.ToString()).Aggregate((a, b) => a + b));
var marathonDistance = double.Parse(races.Select(r => r.Distance.ToString()).Aggregate((a, b) => a + b));
var marathonRace = new Race {
    Id = 0,
    Time = marathonTime,
    Distance = marathonDistance
};

var part1 = winConditions.Aggregate((a, b) => a * b);
var part2 = marathonRace.GetWinConditions();

Console.WriteLine($"{new string('=', 10)} SOLUTIONS {new string('=', 10)}");
Console.WriteLine($"Part 1: {part1}");
Console.WriteLine($"Part 2: {part2}");

return 0;

class Race {
    public required int Id { get; set; }

    public required double Time { get; set; }

    public required double Distance { get; set; }

    public double GetWinConditions() {
        var lowerBound = (double)0;
        var upperBound = (double)Time;

        var lowerTask = Task.Run(() => {
            for (var holdLength = 0; holdLength < Time; holdLength++) {
                var movementPotential = (Time - holdLength) * holdLength;

                if (movementPotential > Distance) {
                    lowerBound = holdLength;
                    break;
                }
            }
        });
        var upperTask = Task.Run(() => {
            for (var holdLength = Time; holdLength > 0; holdLength--) {
                var movementPotential = (Time - holdLength) * holdLength;

                if (movementPotential > Distance) {
                    upperBound = holdLength;
                    break;
                }
            }
        });
        Task.WaitAll(lowerTask, upperTask);

        return upperBound - lowerBound + 1;
    }
}