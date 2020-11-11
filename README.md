# go-rss-box

go-rss-box generates a private RSS feed from the newsletters in your inbox.

Essentially my own implementation of [Kill the Newsletter](https://kill-the-newsletter.com/).

## Usage

Start by copying ```config.example.yml``` to ```config.yml``` and filling in the missing details.

```bash
git clone https://github.com/qcasey/go-rss-box
cd go-rss-box
cp config.example.yml config.yml

go get .
go build

> ./go-rss-box
2020/11/10 16:45:45 Starting server...
2020/11/10 16:45:45 Your feed:
2020/11/10 16:45:45     JSON: http://localhost:8000/entries/
2020/11/10 16:45:45     RSS: http://localhost:8000/feed/
2020/11/10 16:45:45     Atom: http://localhost:8000/atom.xml
```

go-rss-box fetches the latest emails on demand. This allows your Feed Reader to have the final say in the refresh interval. 