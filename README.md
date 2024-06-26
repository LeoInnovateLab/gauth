<h1 align="center">GAuth</h1>

<p align="center">
It's just a library of third-party authorization tools that make it easy to get rid of the cumbersome third-party login SDKs.
<br/>
<br/>
<a href="https://github.com/LeoInnovateLab/gauth/blob/master/LICENSE">
  <img alt="GitHub" src="https://img.shields.io/github/license/LeoInnovateLab/gauth"/>
</a>  
<a href="https://pkg.go.dev/github.com/LeoInnovateLab/gauth">
    <img src="https://pkg.go.dev/badge/github.com/LeoInnovateLab/gauth.svg" alt="Go Reference">
</a>
<a href="https://goreportcard.com/report/github.com/LeoInnovateLab/gauth">
    <img src="https://goreportcard.com/badge/github.com/LeoInnovateLab/gauth" />
</a>
<a href="https://github.com/LeoInnovateLab/gauth/actions">
    <img src="https://github.com/LeoInnovateLab/gauth/actions/workflows/test.yml/badge.svg" />
</a>
</p>

<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README.zh-Hans.md)

</samp>
</strong>
</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Introduction](#Introduction)
- [Currently Supported](#Currently-Supported)
- [Getting Started](#Getting-Started)
- [Inspiration](#Inspiration)

## Introduction

Providing a tool library for Go programs to enable third-party authorization login. 
This library encapsulates the cumbersome process of integrating third-party logins into a simple function call, 
making login easier.More third-party authorization are under development.

## Currently Supported

* Github
* Google
* Facebook
* Slack
* Linkedin

## Getting Started

```go
// create authorization request
authRequest, err := gauth.New().
Source("github").
ClientId("your_client_id").
ClientSecret("your_client_secret").
RedirectUrl("http://localhost:8080/auth/github/callback").
Build()

// generate authorization URL		
authorizeUrl, err := authRequest.Authorize("state")

// After authorization login, it will return login information.
response, err := authRequest.Login(callback)
```

More documentation will be updated on [website](https://gauth.dev)

## Inspiration

The project inspiration mainly comes from：
- [JustAuth](https://github.com/justauth/JustAuth)