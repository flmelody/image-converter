/*
 * Copyright 2023 Flmelody
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package converter

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

const (
	PNG string = "png"
)

type PngConverter struct {
}

func (converter *PngConverter) Encode(dest string, m image.Image) error {
	f, err := os.Create(dest)
	if err != nil {
		fmt.Print("failed to create png file", err)
		return err
	}
	return png.Encode(f, m)
}

func (converter *PngConverter) Decode(source string) (image.Image, error) {
	f, err := os.Open(source)
	if err != nil {
		fmt.Println("failed to open png file", err)
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("failed to close file", err)
		}
	}(f)
	return png.Decode(f)
}
