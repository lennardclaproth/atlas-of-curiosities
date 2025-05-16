namespace DesignPatterns.Adapter.Payment;

public interface IPayPalPayment
{
    string MakePayment(decimal amount);
}
