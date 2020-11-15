// Instagram Stories Generator, (C) 2020 Lenin Alevski Huerta Arias.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package image_generator

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/otiai10/opengraph"
)

func downloadFile(URL string) ([]byte, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}
	image, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func getImageFromURL(url string) ([]byte, error) {
	og, err := opengraph.Fetch(url)
	if err != nil {
		return nil, err
	}
	if len(og.Image) > 0 {
		imageURL := og.Image[0].URL
		return downloadFile(imageURL)
	}
	return nil, errors.New("no opengraph image found")
}