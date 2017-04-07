# songo
[![Build Status](https://travis-ci.org/WindomZ/songo.svg?branch=master)](https://travis-ci.org/WindomZ/songo)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

实现songo协议的Go库

## Documents

[songo协议文档](https://github.com/suboat/songo/blob/master/desc.md)

## Features

- [ ] support MySQL
- [ ] support PostgreSQL
- [x] support MongoDB

## [MySQL](https://github.com/WindomZ/songo/tree/master/mysql)

![v0.0.0](https://img.shields.io/badge/version-v0.0.0-orange.svg)

* 开发中...

## [PostgreSQL](https://github.com/WindomZ/songo/tree/master/postgre)

![v0.0.0](https://img.shields.io/badge/version-v0.0.0-orange.svg)

* 开发中...

## [MongoDB](https://github.com/WindomZ/songo/tree/master/mongo)

![v1.1.0](https://img.shields.io/badge/version-v1.1.0-blue.svg)

### Example
* 实例1

  URL地址：
  ```
  http(s)://127.0.0.1/demo?_limit=50&_page=2
    &_sort=created,money,-level
    &year=$eq$2016&month=$bt$8,11&date=$eq$1&day=$in$0,6
  ```
  对应JSON：
  ```
  {
      "limit":50,
      "page":2,
      "sort":[
          "created",
          "money",
          "-level"
      ],
      "year":"$eq$2016",
      "month":"$bt$8,11",
      "date":"$eq$1",
      "day":"$in$0,6"
  }
  ```

## License

The [MIT License](https://github.com/WindomZ/songo/blob/master/LICENSE)
