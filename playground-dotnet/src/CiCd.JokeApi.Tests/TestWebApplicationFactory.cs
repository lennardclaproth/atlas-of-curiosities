using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.Extensions.Hosting;

namespace CiCd.JokeApi.Tests;

public class TestWebApplicationFactory<TProgram> : WebApplicationFactory<TProgram> where TProgram : class 
{
    protected override IHost CreateHost(IHostBuilder builder)
    {
        builder.ConfigureHostConfiguration(config =>
        {
        
        });

        builder.ConfigureServices(services =>
        {
            
        });

        return base.CreateHost(builder);
    }
}
