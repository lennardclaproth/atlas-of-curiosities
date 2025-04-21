namespace DesignPatterns.AbstractFactory.Devices;

public class AppleMobile : IDevice
{
    public void GetDetails()
    {
        Console.WriteLine("Apple Mobile");
    }
}

public class AppleLaptop : IDevice
{
    public void GetDetails()
    {
        Console.WriteLine("Apple Laptop");
    }
}

public class AppleDesktop : IDevice
{
    public void GetDetails()
    {
        Console.WriteLine("Apple Desktop");
    }
}
