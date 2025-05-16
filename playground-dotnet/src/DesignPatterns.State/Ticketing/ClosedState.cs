namespace DesignPatterns.State;

public class ClosedState : ITicketState
{
    public void Handle(TicketContext context)
    {
        Console.WriteLine("Ticket is closed and cannot be handled further.");
    }

    public void Close(TicketContext context)
    {
        Console.WriteLine("Ticket is already closed.");
    }
}
