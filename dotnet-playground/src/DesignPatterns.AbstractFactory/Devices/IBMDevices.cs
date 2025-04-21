namespace DesignPatterns.AbstractFactory.Devices;

public class IBMMobile : IDevice
{
    public void GetDetails()
    {
        Console.WriteLine("IBM Mobile");
    }
}

public class IBMLaptop : IDevice
{
    public void GetDetails()
    {
        Console.WriteLine("IBM Laptop");
    }
}

public class IBMDesktop : IDevice
{
    public void GetDetails()
    {
        Console.WriteLine("IBM Desktop");
    }
}
