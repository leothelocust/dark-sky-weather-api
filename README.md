# Dark Sky API Interface
A Golang Server that interfaces with the Dark Sky API so my Chrome Extension can get the weather.

## Firewall
    sudo ufw status
    sudo ufw allow OpenSSH
    sudo ufw allow https
    sudo ufw enable

## Install Go

    curl -C - https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz -o go1.11.4.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz
    echo "export PATH=$PATH:/usr/local/go/bin"
    go version

## Letsencrypt
    sudo add-apt-repository ppa:certbot/certbot
    sudo apt-get update
    sudo apt-get install python-certbot-nginx
    certbot certonly --standalone -d weather.l3vi.co

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
