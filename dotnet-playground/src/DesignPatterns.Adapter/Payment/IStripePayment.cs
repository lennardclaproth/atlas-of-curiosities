namespace DesignPatterns.Adapter.Payment;

public interface IStripePayment
{
    string ProcessPayment(decimal amount);

}
