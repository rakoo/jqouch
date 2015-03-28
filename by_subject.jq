{
  "_id": "_design/by_subject_jq",
  "language": "jq",
  "views": {
    "by_subject": {
      "map": "[[.Classifications | .\"Subject matter:\" | .[0] |.text, null]]"
    }
  }
}

