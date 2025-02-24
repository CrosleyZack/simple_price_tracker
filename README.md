# Price Tracker

Gets price events from websites and stores them into a json file.

This requires two configuration files:

## data/sites.json

Defines how to parse prices out of specific websites

```json
{
  "type": "object",
  "description": "sites.json schema",
  "properties": {
    "name": {
      "type": "string",
      "description": "String name for the website",
      "exclusiveMinimum": 0,
    },
    "url": {
      "type": "string",
      "description": "Top level url for the website",
      "exclusiveMinimum": 0,
    },
    "price_path": {
      "type": "string",
      "description": "Defintion of how to pull prices from the website. See ...",
    }
  },
  "required": [ "name", "url", "price_path" ]
}
```

## data/items.json

Defines the specific items on the site to track prices of by their path extension

```json
{
  "type": "object",
  "description": "quotes.json schema",
  "properties": {
    "name": {
      "type": "string",
      "description": "Name to identify the item by",
    },
    "website": {
      "type": "string",
      "description": "Name of the website in sites.json it is on",
    },
    "path": {
      "type": "string",
      "description": "path extension following the base website url to the product",
    }
  },
  "required": [ "name", "website", "path" ],
}
```

## Path Defintions

Websites are parsed using [Soup](github.com/anaskhan96/soup). The path definition is a string that:

1. Separates [FindAllStrict](https://pkg.go.dev/github.com/anaskhan96/soup#Root.FindAllStrict) requests using pipe `|` character
2. Separates individual arguments to FindAllStrict using a period `.`
3. Indexes into the results of FindAllStrict either using index 0 or the number at the end of the arguments in step 2. Negative numbers index from the end of the list.

For example, to find price into Patagonia we use `div.class.price|span.itemprop.price.-1` which indicates:

1. First do a FindAllStrict on tags `<div class="price">` and return the first result (index 0)
2. Then do a FindAllStrict on tags `<span itemprop="price">` and return the last result (index -1)

## Cron

A daily cron is run using the workflow in `.github/workflows/process.yml` to fetch the prices of the items in `data/items.json` and store changes in `data/events.json`. This requires read-write and commit permissions to the workflow under repository settings > actions > general > workflow permissions.
