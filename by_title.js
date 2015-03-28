{
  "_id": "_design/by_title_js",
  "views": {
    "by_title": {
      "map": "function(doc) { emit(doc.title, null) }"
    }
  }
}

