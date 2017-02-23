# songo

实现songo协议的Go库

## Documents

[songo协议文档](https://github.com/suboat/songo/blob/master/desc.md)

## Features

- [ ] support SQL
- [ ] support Postgres
- [x] support Mongo

## Usage

#### SQL

* 开发中...

#### Postgres

* 开发中...

#### Mongo

## Example
* 实例1

  URL地址：
  ```
  http(s)://x.x.x/xxx?_limit=50&_page=2
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
