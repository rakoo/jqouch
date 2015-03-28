{
  "_id": "_design/by_date_jq",
  "language": "jq",
  "views": {
    "by_date": {
    "map": "[ .dates[] | [[.type, .date], null] ]"
    }
  }
}

