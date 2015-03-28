{
  "_id": "_design/by_amended_by_js",
  "views": {
    "by_amended_by": {
      "map": "function(doc) { var count = 0; var field = doc[\"Relationship between documents\"][\"Amended by:\"]; for (var i = 0; i < field.length; i++) { count += field[i].text.length } emit(doc.title, count) }"
    }
  }
}
