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
	"image/color"
	"strings"

	"github.com/minio/minio/pkg/env"
)

// image dimensions and margins
var (
	height       = float64(1920)
	width        = float64(1080)
	marginTop    = float64(250)
	marginBottom = float64(250)
	marginLeft   = float64(100)
	marginRight  = float64(100)
	fontSize     = float64(96)
)

// colors, font-type, etc ...
var palettes = []palette{
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0x4C, G: 0xF1, B: 0xE1, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0xDB, G: 0xCF, B: 0xB0, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0xBF, G: 0xC8, B: 0xAD, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0x90, G: 0xB4, B: 0x94, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 0x71, G: 0x8F, B: 0x94, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 0x54, G: 0x57, B: 0x75, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 232, G: 30, B: 99, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 156, G: 39, B: 175, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 33, G: 150, B: 242, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 0x4C, G: 0xAE, B: 0x50, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0xCC, G: 0xDB, B: 0x39, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0xFE, G: 0x98, B: 0x00, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0xFE, G: 0xEA, B: 0x3B, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 0xc2, G: 0x18, B: 0x5b, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 0x21, G: 0x21, B: 0x21, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 0xff, G: 0x80, B: 0xab, A: 0xff},
	},
	{
		fontColor:       color.RGBA{A: 0xff},
		backgroundColor: color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
	},
	{
		fontColor:       color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
		backgroundColor: color.RGBA{R: 0xab, G: 0x47, B: 0xbc, A: 0xff},
	},
}

var fonts = []string{
	"assets/fonts/Roboto-Black.ttf",
}

var publicationsMock = []publication{
	{
		title:  "New Framework Released to Protect Machine Learning Systems From Adversarial Attacks",
		url:    "https://thehackernews.com/2020/10/adversarial-ml-threat-matrix.html",
		source: "The Hacker News",
	},
	{
		title:  "The RIAA is coming for the YouTube downloaders",
		url:    "https://thehackernews.com/2020/10/adversarial-ml-threat-matrix.html",
		source: "TechCrunch",
	},
	{
		title:  "Daily Crunch: Uber and Lyft defeated again in court",
		url:    "https://thehackernews.com/2020/10/adversarial-ml-threat-matrix.html",
		source: "TechCrunch",
	},
	{
		title:  "U.S. Levies Sanctions Against Russian Research Institution Linked to Triton Malware",
		url:    "https://thehackernews.com/2020/10/adversarial-ml-threat-matrix.html",
		source: "Threatpost",
	},
	{
		title:  "Cybercriminals Could be Coming After Your Coffee",
		url:    "https://thehackernews.com/2020/10/adversarial-ml-threat-matrix.html",
		source: "Dark Reading",
	},
	{
		title:  "Iran-Linked Seedworm APT target orgs in the Middle East",
		url:    "https://thehackernews.com/2020/10/adversarial-ml-threat-matrix.html",
		source: "Security Affairs",
	},
}

// assets constants
const (
	avatarPath = "assets/images/avatar.png"
	lightFont  = "assets/fonts/Rajdhani-Regular.ttf"
)

// ENV variables for additional configuration
const (
	authorBioDescription = "AUTHOR_BIO_DESCRIPTION"
)

// author description, it's a long text separated by comma
func getAuthorBioDescription() []string {
	description := strings.TrimSpace(env.Get(authorBioDescription, "If you want to view paradise, Simply look around and view it, Follow me @alevskey"))
	return strings.Split(description, ",")
}
