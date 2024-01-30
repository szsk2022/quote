# Using SZSK Daily Quote API in HTML Pages

#### Add the script code at any desired location
```html
<script>
    document.addEventListener("DOMContentLoaded", function () {
        getRandomQuote();

        function getRandomQuote() {
            fetch('http://localhost:8080/?lang=cn') // Replace this with your deployed API URL; change to 'en' for English quotes
                .then(response => response.json())
                .then(data => {
                    if (data.error) {
                        document.getElementById('quote').innerText = 'Loading failed, please try again later';
                    } else {
                        document.getElementById('quote').innerText = data.quote;
                    }
                })
                .catch(error => console.error('Error:', error));
        }
    });
</script>
```
### Call the API at the appropriate place

```html
<div id="quote">Loading...</div>
```

In this section, you are guided on how to integrate SZSK's Daily Quote API into an HTML page. When the page is loaded, the `getRandomQuote` function will be executed, which sends a request to the API (with the language parameter set as Chinese). The fetched quote is then displayed within a `div` element with the ID `'quote'`. If there's an error while fetching the quote, it displays a message saying 'Loading failed, please try again later'. To request an English quote, simply replace `'lang=cn'` with `'lang=en'` in the fetch URL.