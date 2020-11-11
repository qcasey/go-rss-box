# go-rss-box

![output](https://quinncasey.com/post/go-rss-box/output_huf35c8b4dd353f57a289c069d281b0a18_137180_1000x0_resize_lanczos_2.png)

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
./go-rss-box
```

go-rss-box fetches the latest emails on demand. This allows your Feed Reader to have the final say in the refresh interval. 