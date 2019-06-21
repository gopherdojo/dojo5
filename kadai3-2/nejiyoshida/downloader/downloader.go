package downloader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"sync"
)

type dlunit struct {
	buf    *bytes.Buffer //DLしたデータ
	offset int64         //どの部分のデータなのか
	err    error
}

func CheckHead(url string) (int64, string, error) {

	// 分割できるか＆ファイルサイズ確認するためHead要求
	resp, err := http.Head(url)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", err
	}

	//responseのContent-Lengthからサイズを確認
	size := resp.ContentLength

	//分割ダウンロードできるかどうか。bytesなら可能
	dltype := resp.Header.Get("Accept-Ranges")

	return size, dltype, nil
}

//一つのダウンロード単位
func download(ctx context.Context, url string, from, to int64) <-chan dlunit {
	ch := make(chan dlunit)

	go func() {
		defer close(ch)

		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			ch <- dlunit{buf: nil, offset: 0, err: err}
			return
		}

		//fromからtoまでリクエストするよう指定。RangeHeaderが利用できないと０から最後まで
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", from, to))
		req = req.WithContext(ctx)

		cli := http.DefaultClient
		resp, err := cli.Do(req)
		if err != nil {
			ch <- dlunit{buf: nil, offset: 0, err: err}
			return
		}

		defer resp.Body.Close()

		var buf bytes.Buffer

		//ゴルーチンで戻すためにDLしたパーツをバッファにコピー
		_, err = io.Copy(&buf, resp.Body)

		if err != nil {
			ch <- dlunit{buf: nil, offset: 0, err: err}
			return
		}

		ch <- dlunit{buf: &buf, offset: from, err: nil}

	}()

	return ch
}

//分割DLできないときの普通のDL
func Download(ctx context.Context, fp *os.File, url string, size int64) error {

	//分割しないので０から最後（ファイルサイズ分）まで
	p := <-download(ctx, url, 0, size)
	if p.err != nil {
		return p.err
	}

	//DLしたファイルを書き込む
	_, err := io.Copy(fp, p.buf)
	if err != nil {
		return err
	}

	_, err = fp.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}

//並行DL
func ParallelDownload(ctx context.Context, fp *os.File, url string, size int64) error {

	numcpu := runtime.NumCPU()
	partsize := size / int64(numcpu)

	//並行してダウンロードするためのスライス
	dlunits := make([]<-chan dlunit, numcpu)

	for i := 0; i < numcpu; i++ {
		var from, to int64

		if i == 0 {
			from = 0
		} else {
			from = partsize*int64(i) + 1
		}

		if i == numcpu-1 {
			to = size
		} else {
			to = from + partsize
		}

		dlunits[i] = download(ctx, url, from, to)
	}

	for p := range merge(dlunits...) {
		if p.err != nil {
			return p.err
		}
		//offsetの地点からそれぞれ書き込むことで分割ダウンロードしたものを組み合わせる
		fp.WriteAt(p.buf.Bytes(), p.offset)
	}

	_, err := fp.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}

func merge(chs ...<-chan dlunit) <-chan dlunit {
	var wg sync.WaitGroup
	merged := make(chan dlunit)

	wg.Add(len(chs))

	for _, ch := range chs {
		go func(ch <-chan dlunit) {
			defer wg.Done()

			p := <-ch
			merged <- p

		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
