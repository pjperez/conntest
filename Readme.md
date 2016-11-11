# Conntest
## Simple webserver

Conntest is a very simple web server built in Go that returns server hostname, client ip address and current time for any request in json format.



### Install from source

    go get github.com/pjperez/conntest
    go install github.com/pjperez/conntest

### Usage

    ./conntest

### Notes

It listens by default on port 8000, so you can run it without root / admin privileges. If you'd like to use a different port just change the value for the serverPORT constant.

### Examples

Ideally you would run a client for a certain amount of time and log the results in a file. In the near future, [httping](https://github.com/pjperez/httping) will support recording [Conntest](https://github.com/pjperez/conntest)'s output to a json file, along with its own RTT measures.    
Both combined would result in a powerful network assesment tool.


#### PowerShell client

    PS C:\temp> $result = ((Invoke-WebRequest -Uri http://localhost:8000).content | ConvertFrom-Json)
    PS C:\temp> $result.Hostname
    DESKTOP-NVM2739
    PS C:\temp> $result.ClientIP
    [::1]
    PS C:\temp> $result.Time
    2016-11-11T00:27:32.4100165Z
    PS C:\temp> $result

    Hostname        ClientIP Time
    --------        -------- ----
    DESKTOP-NVM2739 [::1]    2016-11-11T00:27:32.4100165Z