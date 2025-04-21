using DesignPatterns.AbstractFactory.Devices;

namespace DesignPatterns.AbstractFactory.Factories;

public interface IDeviceFactory
{
    IDevice CreateMobile();
    IDevice CreateLaptop();
    IDevice CreateDesktop();
}
