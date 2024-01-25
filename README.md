[![Go](https://github.com/njitacm/jerseyctf-registration-site/actions/workflows/go.yml/badge.svg)](https://github.com/njitacm/jerseyctf-registration-site/actions/workflows/go.yml)
# jerseyctf-website


NJIT ACM's website for registration and information on the JerseyCTF event. ([jerseyctf.com](https://jerseyctf.com/))


<!-- NJIT ACM LOGO -->
<p align="center">
    <img src="https://raw.githubusercontent.com/NJIT-ACM/NJIT-ACM/main/ACMLOGO.png" alt="NJIT ACM Logo" data-canonical-src="" width="400">
</p>

## Technical Details
* A dynamic site that uses HTML templates, from the **`go`** standard library 
* The back-end (web server) is written in **golang-go**
* The front-end is written using **Bootstrapv5**

## Usage
```bash
# To test JCTF site locally
# Required: go (golang) 
# Optional: make
$ git clone https://github.com/jerseyctf/jerseyctf-website.git
$ cd jerseyctf-website/src
$ go build main.go
# Optionally: use `make`
-----
Running:
# If MacOS / Linux:
$ ./main  
# If Windows:
$ ./main.exe
- Go to http://localhost:9990
```

## Purpose
* Made to be deployed onto a Droplet (cloud server on DigitalOcean) easily
* Check '[How To Deploy a Go Web Application Using Nginx on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04)' for further instructions on deploying the site to web.
