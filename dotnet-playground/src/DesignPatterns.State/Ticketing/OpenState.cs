namespace DesignPatterns.State;

public class OpenState : ITicketState
{
    public void Handle(TicketContext context)
    {
        Console.WriteLine("Ticket is now InProgress.");
        context.SetState(new InProgressState());
    }

    public void Close(TicketContext context)
    {
        Console.WriteLine("Ticket cannot be closed directly from Open state.");
    }
}
