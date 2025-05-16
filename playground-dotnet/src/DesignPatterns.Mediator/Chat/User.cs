namespace DesignPatterns.Mediator.Chat;

public class User : IUser
{
    private readonly IChatMediator _chatMediator;
    public string Name { get; }

    public User(IChatMediator chatMediator, string name)
    {
        _chatMediator = chatMediator;
        Name = name;
    }

    public void SendMessage(string message)
    {
        Console.WriteLine($"{Name} sends: {message}");
        _chatMediator.SendMessage(message, this);
    }

    public void ReceiveMessage(string message)
    {
        Console.WriteLine($"{Name} receives: {message}");
    }
}
