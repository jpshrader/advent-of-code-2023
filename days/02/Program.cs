﻿var lines = File.ReadLines("input.txt");
if (lines == null || !lines.Any()) {
    Console.WriteLine("file/lines not found");
    return -1;
}
lines = lines.Select(l => l.Trim().ToLower());

var games = new List<Game>();
foreach (var line in lines) {
    games.Add(new Game(line));
}

Console.WriteLine($"{new string('=', 10)} SOLUTIONS {new string('=', 10)}");
Console.WriteLine($"Part 1: {Part1(games)}");
Console.WriteLine($"Part 2: {Part2(games)}");

return 0;

static int Part1(List<Game> games) {
    return games
        .Where(g => g.IsPossible(redLimit: 12, greenLimit: 13, blueLimit: 14))
        .Sum(g => g.GameId);
}

static int Part2(List<Game> games) {
    return games
        .Select(g => g.GetColorMinimums())
        .Sum(c => c.red * c.green * c.blue);
}

class Game {
    public int GameId { get; set; }
    public IEnumerable<Set> Sets { get; set; }
    public Game(string line) {
        var parts = line.Split(':');

        GameId = int.Parse(parts[0].Replace("game", "").Trim());
        Sets = parts[1].Split(';').Select(s => new Set(s));
    }

    public (int red, int green, int blue) GetColorMinimums() {
        var red = Sets.Max(s => s.Red);
        var green = Sets.Max(s => s.Green);
        var blue = Sets.Max(s => s.Blue);

        return (red, green, blue);
    }

    public bool IsPossible(int redLimit, int greenLimit, int blueLimit) {
        return Sets.All(s => s.IsPossible(redLimit, greenLimit, blueLimit));
    }
}

class Set {
    public int Red { get; set; }
    public int Green { get; set; }
    public int Blue { get; set; }

    public Set(string line) {
        var parts = line.Split(',');
        foreach (var part in parts) {
            if (part.Contains("red")) {
                Red = int.Parse(part.Replace("red", "").Trim());
            } else if (part.Contains("green")) {
                Green = int.Parse(part.Replace("green", "").Trim());
            } else if (part.Contains("blue")) {
                Blue = int.Parse(part.Replace("blue", "").Trim());
            }
        }
    }

    public bool IsPossible(int redLimit, int greenLimit, int blueLimit) {
        return Red <= redLimit && Green <= greenLimit && Blue <= blueLimit;
    }
}