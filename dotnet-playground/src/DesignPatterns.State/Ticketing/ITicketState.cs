namespace DesignPatterns.State;

public interface ITicketState
{
    void Handle(TicketContext context);
    void Close(TicketContext context);
}
