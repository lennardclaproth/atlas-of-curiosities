using DesignPatterns.AbstractFactory.Devices;

namespace DesignPatterns.AbstractFactory.Factories;

public class AppleFactory : IDeviceFactory
{
    public IDevice CreateMobile()
    {
        return new AppleMobile();
    }

    public IDevice CreateLaptop()
    {
        return new AppleLaptop();
    }

    public IDevice CreateDesktop()
    {
        return new AppleDesktop();
    }
}
