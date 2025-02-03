package mnist

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"os"
)

func ReadLabels(name string) ([][]float64, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	r, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	var magic, nLabel int32
	if err := binary.Read(r, binary.BigEndian, &magic); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &nLabel); err != nil {
		return nil, err
	}

	if magic != 2049 {
		return nil, fmt.Errorf("not allowed label file (magic: %d)", magic)
	}

	labels := make([]byte, nLabel)
	if err := binary.Read(r, binary.BigEndian, &labels); err != nil {
		return nil, err
	}

	output := make([][]float64, len(labels))
	for i, label := range labels {
		output[i] = make([]float64, 10)
		output[i][int(label)] = 1
	}

	return output, r.Close()
}

func ReadImages(name string) ([][]float64, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	r, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	var magic, nImage, rows, cols int32
	if err := binary.Read(r, binary.BigEndian, &magic); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &nImage); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &rows); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &cols); err != nil {
		return nil, err
	}

	if magic != 2051 {
		return nil, fmt.Errorf("not allowed image file (magic: %d)", magic)
	}

	images := make([][]float64, nImage)
	for i := range images {
		img := make([]byte, rows*cols)
		if err := binary.Read(r, binary.BigEndian, img); err != nil {
			return nil, err
		}

		images[i] = make([]float64, rows*cols)
		for j, pixel := range img {
			images[i][j] = float64(pixel) / 255.0
		}
	}

	return images, r.Close()
}
