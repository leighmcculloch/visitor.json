# visitor.json

Get information about your visitor in a simple JavaScript Object.

```json
{
  "acceptedLanguages": [
    "pt-BR",
    "pt",
    "en-US",
    "en"
  ]
}
```

## Usage

### Simple

```html
<script src="https://example.com/js"></script>
<script>
  console.log(window.visitor.acceptedLanguages);
</script>
```

### Simple Async

```html
<script>
  function visitorLoaded(visitor) {
    console.log(visitor.acceptedLanguages);
  }
</script>
<script src="https://example.com/jsonp?callback=visitorLoaded" async="async"></script>
```

### Ajax

Using jQuery it would look like the following, but you can use whichever
framework or non-framework you please.

```javascript
$.ajax('https://example.com/json').done(function(data) {
  console.log(data.acceptedLanguages)
});
```

## Why

Originally created for [greatstories.org](https://greatstories.org), a static
website that has no backend and but is viewable in multiple languages. The
website needed to know the user's preferred languages to make accurate
recommendations.

## Why Abandoned

Abandoned when I realized I could use a series of simple tests to get the same
information in JavaScript.

## Deploy to PivotalWS

The project contains a `manifest.yml` and can be deployed to
CloudFounder/[PivotalWS](https://run.pivotal.io) using `cf push`.

## Deploy to Heroku

The project contains a `Procfile` and can be deployed to
[Heroku](https://heroku.com) out-of-the-box.
