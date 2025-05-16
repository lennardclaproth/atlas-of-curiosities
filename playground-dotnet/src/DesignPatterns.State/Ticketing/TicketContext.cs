namespace DesignPatterns.State;

public class TicketContext
{
    private ITicketState _currentState;

    public TicketContext(ITicketState initialState)
    {
        _currentState = initialState;
    }

    public void SetState(ITicketState state)
    {
        _currentState = state;
        Console.WriteLine($"State changed to: {_currentState.GetType().Name}");
    }

    public void Handle()
    {
        _currentState.Handle(this);
    }

    public void Close()
    {
        _currentState.Close(this);
    }
}
