namespace DesignPatterns.Adapter.Payment;

internal sealed class StripePayment : IStripePayment
{
    public string ProcessPayment(decimal amount)
    {
        // Simulate making a payment via PayPal
        return $"Paid {amount} using Stripe";
    }
}
