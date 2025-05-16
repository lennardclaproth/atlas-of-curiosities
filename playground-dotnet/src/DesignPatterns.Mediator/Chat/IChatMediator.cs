namespace DesignPatterns.Mediator.Chat;

public interface IChatMediator
{
    void SendMessage(string message, IUser sender);
}
