# Katsu-Don: A Firefox Extension Downloader
This is a tool that allows you to download and save any firefox extension located at addons.mozilla.org

# Build requirements:
- [go](https://golang.org)

# Building:
Run the following to build:
`go build katsu-don.go`

# Usage:
The first thing you need to know is what the extension number is. To find out, go to the page on addons.mozilla.org. On the button that says "Add to firefox, right click and press "Copy Link Location". You should get a link similar to this:

`https://addons.mozilla.org/firefox/downloads/file/#######/extension.xpi`

The `#######` is the extension number.

After building katsu-don, the arguments go as follows:

`./katsu-don.exe -extNum ####### -fileName extension`

where `#######` is the extension number and `extension` is the file name you want.

### TODO:
I will need to change how it works so all you need to type is `./katsu-don.exe -ext extension` but that'll come as I learn.