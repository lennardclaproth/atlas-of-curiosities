using DesignPatterns.AbstractFactory.Devices;

namespace DesignPatterns.AbstractFactory.Factories;

public class LenovoFactory : IDeviceFactory
{
    public IDevice CreateMobile()
    {
        return new LenovoMobile();
    }

    public IDevice CreateLaptop()
    {
        return new LenovoLaptop();
    }

    public IDevice CreateDesktop()
    {
        return new LenovoDesktop();
    }
}
