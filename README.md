# Dark Sky API Interface
A Golang Server that interfaces with the Dark Sky API so my Chrome Extension can get the weather.

## Systemd
    sudo useradd weatherapi -s /sbin/nologin -M
    sudo cp weather-api.service /lib/systemd/system/.
    ls -al /lib/systemd/system
    sudo chmod 755 /lib/systemd/system/weather-api.service
    # add the DARK_SKY_API_KEY to the weather-api.service file
    sudo systemctl enable weather-api.service
    sudo systemctl start weather-api.service
    sudo journalctl -f -u weather-api
    # or 
    sudo systemctl status weather-api.service

## Thanks
Much of the source I obtained from [github.com/shawntoffel/darksky](https://github.com/shawntoffel/darksky).  So thank you [shawntoffel](https://github.com/shawntoffel) and [contributors](https://github.com/shawntoffel/darksky/graphs/contributors) for that amazing codebase!
