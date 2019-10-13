# GoServe - HTTP File Server

Quickly download files and folders from computers on the local network

## Why

I found Python's quick local network file server to be useful for transferring files in a pinch. The only problem is you cannot download folders. GoServe quickly zips folders before sending them. It also looks a bit nicer than the default Python file server.

## Usage
```
Usage of ./file_server:
  -home
        Sets the base file server directory to your home directory
  -port string
        Change the server's port (default "8080")
  -showDots
        Enable to make dotfiles visible
  -unsafe
        Sets the base file server directory to the root '/' directory. Use with caution.
```

## Screenshot

<img src="https://raw.githubusercontent.com/evan-buss/go-serve/master/screenshot/screenshot.png" alt="screenshot"/>
