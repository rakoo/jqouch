{
  "_id": "_design/by_subject_js",
  "views": {
    "by_subject": {
      "map": "function(doc) { emit(doc.Classifications['Subject matter:'][0].text, null) }"
    }
  }
}
