<h1 align="center">GAuth</h1>

<p align="center">
它是一个第三方授权登录的工具类库，可以让我们脱离繁琐的第三方登录 SDK，让登录变得更加简单。
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


## 内容目录

- [内容目录](#内容目录)
- [项目介绍](#项目介绍)
- [已支持的第三方](#已支持的第三方)
- [快速入门](#快速入门)
- [灵感](#灵感)

## 项目介绍

为 Go 程序提供第三方授权登录的工具类库。
将集成第三方登录的繁琐流程封装成一个简单的函数调用，让登录变得更加简单。
更多第三方登录正在开发中。

## 已支持的第三方

* Github
* Google
* Facebook
* Slack
* Linkedin

## 快速入门

```go
// 创建授权 request
authRequest, err := gauth.New().
    Source("github").
    ClientId("your_client_id").
    ClientSecret("your_client_secret").
    RedirectUrl("http://localhost:8080/auth/github/callback").
    Build()

// 生成授权 URL		
authorizeUrl, err := authRequest.Authorize("state")

// 授权登录后会返回 code
response, err := authRequest.Login(callback)
```

更多文档将更新在[网站](https://gauth.dev)

## 灵感

项目灵感主要来自：
- [JustAuth](https://github.com/justauth/JustAuth)