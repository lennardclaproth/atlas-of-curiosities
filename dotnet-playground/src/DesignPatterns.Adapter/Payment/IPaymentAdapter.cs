namespace DesignPatterns.Adapter.Payment;

public interface IPaymentAdapter
{
    string Pay(decimal amount);
}
