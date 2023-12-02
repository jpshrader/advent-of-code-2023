namespace day_2 {
    public static class Part1 {
        public static int Run(IEnumerable<string> lines) {
            var games = new List<Game>();
            foreach (var line in lines) {
                games.Add(new Game(line));
            }

            var possibleGames = games.Where(g => g.IsPossible(redLimit: 12, greenLimit: 13, blueLimit: 14));
            return possibleGames.Sum(g => g.GameId);
        }
    }

    class Game {
        public int GameId { get; set; }
        public IEnumerable<Set> Sets { get; set; }
        public Game(string line) {
            var parts = line.Split(':');

            GameId = int.Parse(parts[0].Replace("game", "").Trim());
            Sets = parts[1].Split(';').Select(s => new Set(s));
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
}