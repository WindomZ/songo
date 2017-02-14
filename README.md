# songo
实现songo协议的Golang库

## Documents

[songo协议文档](https://github.com/suboat/songo)

## Usage

开发中...

## Example
* 实例1

  URL地址：
  ```
  http://x.x.x/xxx?_limit=50&_page=2
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

## TODO

- [ ] sql support SQL
- [ ] mongo support MongoDB

## License

The MIT License ([MIT](https://github.com/WindomZ/songo/blob/master/LICENSE))
