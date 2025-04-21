namespace DesignPatterns.Adapter.Payment;

internal sealed class PayPalPayment : IPayPalPayment
{
    public string MakePayment(decimal amount)
    {
        return $"Paid {amount} using PayPal";
    }
}
