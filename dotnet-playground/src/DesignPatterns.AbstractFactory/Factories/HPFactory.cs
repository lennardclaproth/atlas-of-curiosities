using DesignPatterns.AbstractFactory.Devices;

namespace DesignPatterns.AbstractFactory.Factories;

public class HPFactory : IDeviceFactory
{
    public IDevice CreateMobile()
    {
        return new HPMobile();
    }

    public IDevice CreateLaptop()
    {
        return new HPLaptop();
    }

    public IDevice CreateDesktop()
    {
        return new HPDesktop();
    }
}
