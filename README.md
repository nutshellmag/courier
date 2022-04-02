# Courier Service

This tool checks your RSS feed for any new additions, and sends out emails accordingly.

## Usage
Enter the Buttondown dashboard and select `"Programming"`. Your API key will the first item on the page. Set this as an environment variables called `BUTTONDOWN_KEY`. You can tell Courier Service to send emails itself, rather than "cache" it for manual approval by using the `COURIER_DIRECT` variable and setting it to `"true"`.

Run the tool everytime a Git push is made with a Git hook or by some other recurring request to ensure Courier Service can see changes as they come: `./courier-service https://maatt.ch/atom.xml`

## Acknowledgements
Courier Service is licensed under the Apache 2.0 license. A plain-text copy can be found with the source of this tool (`/LICENSE`).

- [Buttondown](https://buttondown.email) for an amazing service
- [mmcdole](https://github.com/mmcdole/gofeed) for the feed parser