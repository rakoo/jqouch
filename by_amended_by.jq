{
  "_id": "_design/by_amended_by_jq",
  "language": "jq",
  "views": {
    "by_amended_by": {
        "map": "[[.title, (.\"Relationship between documents\" | .\"Amended by:\"  | reduce .[] as $item (0; . + ($item.text | length))) ]]"
    }
  }
}
