namespace CiCd.JokeApi.Tests
{
    public class JokeApiIntegrationTests : IClassFixture<TestWebApplicationFactory<Program>>
    {
        private readonly HttpClient _client;
        private readonly TestWebApplicationFactory<Program> _factory;

        public JokeApiIntegrationTests(TestWebApplicationFactory<Program> factory)
        {
            _factory = factory;
            _client = factory.CreateClient();
        }

        [Fact]
        public async Task etRandomJoke_ReturnsJoke()
        {
            // Act
            var response = await _client.GetAsync("/random");

            // Assert
            response.EnsureSuccessStatusCode(); // 200-299 status code
            var joke = await response.Content.ReadAsStringAsync();
            Assert.False(string.IsNullOrWhiteSpace(joke));
        }

        [Fact]
        public async Task GetAllJokes_ReturnsAllJokes()
        {
            // Act
            var response = await _client.GetAsync("/all");

            // Assert
            response.EnsureSuccessStatusCode(); // 200-299 status code
            var jokes = await response.Content.ReadAsStringAsync();
            Assert.Contains("Why don't skeletons fight each other? They don't have the guts.", jokes);
            Assert.Contains("What do you call fake spaghetti? An impasta!", jokes);
        }
    }
}
