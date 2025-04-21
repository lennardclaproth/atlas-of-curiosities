using DesignPatterns.AbstractFactory.Devices;

namespace DesignPatterns.AbstractFactory.Factories;

public class IBMFactory : IDeviceFactory
{
    public IDevice CreateMobile()
    {
        return new IBMMobile();
    }

    public IDevice CreateLaptop()
    {
        return new IBMLaptop();
    }

    public IDevice CreateDesktop()
    {
        return new IBMDesktop();
    }
}
