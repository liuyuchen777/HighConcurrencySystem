<!--
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 09:59:49
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-12 15:05:07
 * @Description: 
 * @FilePath: /Local_Lab/spike_system/README.md
 * @GitHub: https://github.com/liuyuchen777
-->
# Super-High Concurrency System

## Intro

This is a project base on Golang Gin framework which is able to withstand extremely high network pressure. Using Redis betweem MySQL database and inquire. Finally, it can realize more than 3000 QPS overload. And because my developemnt environment is on VMware, I believe it can be better on higher performance signle point machine.

## Technology Stack

* Frone End: React.js

* Back End: Golang + Gin

* API Style: RESTful API

* Database: Redis + MySQL

* ORM: GORM

* DevOps: Docker + Docker Compose (Under Construct)

## System Structure

![System Architecture](./README_assets/sa.jpg)

## Performance

[Performance Data](https://liuyuchen777.github.io/blog/blog6/index.html)
