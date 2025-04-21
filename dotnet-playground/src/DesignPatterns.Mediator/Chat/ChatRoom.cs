namespace DesignPatterns.Mediator.Chat;

public class ChatRoom : IChatMediator
{
    private readonly List<IUser> _users = new();

    public void RegisterUser(IUser user)
    {
        _users.Add(user);
    }

    public void SendMessage(string message, IUser sender)
    {
        foreach (var user in _users)
        {
            if (user != sender)
            {
                user.ReceiveMessage($"{sender.Name}: {message}");
            }
        }
    }
}
