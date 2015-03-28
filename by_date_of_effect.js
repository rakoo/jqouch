{
  "_id": "_design/by_date_js",
  "views": {
    "by_date": {
      "map": "function(doc) { for (var i = 0; i < doc.dates.length; i++) {emit([doc.dates[i].type, doc.dates[i].date], null)} }"
    }
  }
}
