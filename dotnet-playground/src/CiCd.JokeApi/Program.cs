var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

var jokes = new List<string>
{
    "Why don't skeletons fight each other? They don't have the guts.",
    "What do you call fake spaghetti? An impasta!",
    "Why don't some couples go to the gym? Because some relationships don't work out.",
    "I'm reading a book about anti-gravity. It's impossible to put down!",
    "I would tell you a construction pun, but I'm still working on it."
};

app.MapGet("/random", () =>
    {
        Random.Shared.Next(-20, 55);
        var joke = jokes[Random.Shared.Next(jokes.Count)];
        return joke;
    })
    .WithName("GetRandomJoke")
    .WithOpenApi();

app.MapGet("/all", () => jokes)
.WithName("GetAllJokes")
.WithOpenApi();

app.Run();

public partial class Program { }
