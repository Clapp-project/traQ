package imagemagick

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"
)

// ErrUnsupportedType 未サポートのファイルタイプです
var ErrUnsupportedType = errors.New("unsupported file type")

// ErrUnavailable ImageMagickが使用できません
var ErrUnavailable = errors.New("imagemagick is unavailable")

// ConvertToPNG srcをimagemagickでPNGに変換します。5秒以内に変換できなかった場合はエラーとなります
func ConvertToPNG(ctx context.Context, execPath string, src io.Reader, maxWidth, maxHeight int) (*bytes.Buffer, error) {
	if len(execPath) == 0 {
		return nil, ErrUnavailable
	}

	if maxHeight <= 0 || maxWidth <= 0 {
		return nil, errors.New("maxWidth or maxHeight is wrong")
	}

	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(c, execPath, "-resize", fmt.Sprintf("%dx%d", maxWidth, maxHeight), "-background", "none", "-", "png:-")

	b, err := cmdPipe(cmd, src)
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			return nil, ErrUnsupportedType
		default:
			return nil, err
		}
	}

	return b, nil
}

// ResizeAnimationGIF Animation GIF画像をimagemagickでリサイズします
// expandがfalseの場合、縮小は行いますが拡大は行いません
func ResizeAnimationGIF(ctx context.Context, execPath string, src io.Reader, maxWidth, maxHeight int, expand bool) (*bytes.Buffer, error) {
	if len(execPath) == 0 {
		return nil, ErrUnavailable
	}

	if maxHeight <= 0 || maxWidth <= 0 {
		return nil, errors.New("maxWidth or maxHeight is wrong")
	}

	sizer := fmt.Sprintf("%dx%d", maxWidth, maxHeight)
	if !expand {
		sizer += ">"
	}
	cmd := exec.CommandContext(ctx, execPath, "-", "-coalesce", "-repage", "0x0", "-resize", sizer, "-layers", "Optimize", "gif:-")

	b, err := cmdPipe(cmd, src)
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			return nil, ErrUnsupportedType
		default:
			return nil, err
		}
	}

	return b, nil
}

func cmdPipe(cmd *exec.Cmd, input io.Reader) (output *bytes.Buffer, err error) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		_, _ = io.Copy(stdin, input)
	}()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	_, _ = io.Copy(b, stdout)

	return b, cmd.Wait()
}
