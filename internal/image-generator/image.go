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
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	urlx "net/url"
	"strings"
	"time"

	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
	uuidx "github.com/google/uuid"
)

// RenderStoryTitle will receive story title and render it using font color over background color
func RenderStoryTitle(dc *gg.Context, title string, fontIndex, colorIndex int) error {
	// Render story title
	dc.SetColor(palettes[colorIndex].fontColor)
	if err := dc.LoadFontFace(fonts[fontIndex], fontSize); err != nil {
		return err
	}
	dc.DrawStringWrapped(title, marginLeft, marginTop, 0, 0, width-marginLeft-marginRight, 2, gg.AlignLeft)
	return nil
}

func RenderStorySource(dc *gg.Context, title, source, url string, fontIndex, colorIndex int) error {
	// Render story source
	//
	// Calculate the position of the source story "by [source]" based on the previously rendered story title
	lines := dc.WordWrap(title, width-marginLeft-marginRight)
	_, titleHeight := dc.MeasureMultilineString(strings.Join(lines, "\n"), 2)
	// decreasing fontSize for story source
	if err := dc.LoadFontFace(fonts[fontIndex], fontSize*0.55); err != nil {
		return err
	}
	sourceText := fmt.Sprintf("By %s", source)
	// calculate width and height for story source
	widthSourceText, heightSourceText := dc.MeasureString(sourceText)
	// draw background rectangle for source adding enough padding for text inside (30 left & right)
	//
	// AWESOME STORY TEST RIGHT HERE!!!
	// PLEASE LIKE AND SUBSCRIBE
	// I LOVE YOU
	// ___________________
	// | By [The Source] |
	// -------------------
	// position the source rectangle relative to the story title: height + 75px and text: height + 150
	dc.DrawRectangle(marginLeft, marginTop+titleHeight+(150/2), widthSourceText+60, heightSourceText+(150/2))
	dc.SetColor(palettes[colorIndex].fontColor)
	dc.Fill()
	dc.SetColor(palettes[colorIndex].backgroundColor)
	dc.DrawString(sourceText, marginLeft+30, marginTop+titleHeight+150)
	// render story domain url
	if url != "" {
		// calculate Domain based on provided URL
		u, err := urlx.Parse(url)
		if err != nil {
			return err
		}
		parts := strings.Split(u.Hostname(), ".")
		domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
		sourceDomain := fmt.Sprintf("READ MORE: %s", domain)
		if err := dc.LoadFontFace(lightFont, fontSize*0.55); err != nil {
			return err
		}
		dc.SetColor(palettes[colorIndex].fontColor)
		dc.DrawString(sourceDomain, marginLeft, marginTop+titleHeight+150+heightSourceText+(150/2))
	}

	return nil
}

func RenderAuthorDescription(dc *gg.Context, colorIndex int) error {
	extraPadding := float64(50)
	im, err := gg.LoadPNG(avatarPath)
	if err != nil {
		return err
	}
	// draw borderline around the author avatar
	iw, ih := im.Bounds().Dx(), im.Bounds().Dy()
	dc.SetColor(palettes[colorIndex].fontColor)
	borderLine := float64(15)
	dc.SetLineWidth(borderLine)
	dc.DrawCircle(marginLeft+float64(iw/2), height-marginBottom-float64(ih/2)-borderLine+extraPadding, float64(iw/2))
	dc.Stroke()
	// author image will be positioned horizontally aligned with text (margin left), and
	// vertically aligned with: total-height (1920) - marginBotton (250) - imageHeight
	// |
	// |  _____
	// |  |   | FirstName LastName
	// |  |   | Description text
	// |  -----
	// |
	// |______________________
	dc.DrawImage(im, int(marginLeft), int(height-marginBottom-float64(ih)-borderLine+extraPadding))
	// renders author description
	description := getAuthorBioDescription()
	if err := dc.LoadFontFace(lightFont, fontSize*0.55); err != nil {
		return err
	}
	dw, _ := dc.MeasureMultilineString(strings.Join(description, "\n"), 2)
	dc.DrawStringWrapped(
		strings.Join(description, "\n"),
		marginLeft+float64(iw)+30,
		height-marginBottom-float64(ih)-borderLine*2+extraPadding,
		0,
		0,
		dw,
		2,
		gg.AlignLeft,
	)
	return nil
}

func RenderBackgroundImage(dc *gg.Context, backgroundImage []byte, backgroundColor color.Color) error {
	im, err := jpeg.Decode(bytes.NewReader(backgroundImage))
	if err != nil {
		return err
	}
	// 1. Create a new filter list and add some filters.
	g := gift.New(
		//gift.Contrast(-80),
		//gift.Colorize(240, 50, 100),
	)
	// 2. Create a new image of the corresponding size.
	// dst is a new target image, src is the original image.
	dst := image.NewRGBA(g.Bounds(im.Bounds()))
	// 3. Use the Draw func to apply the filters to src and store the result in dst.
	g.Draw(dst, im)

	dc.DrawImage(dst, 0, 0)

	return nil
}

func RenderNewImage(title, source, url string) error {
	// randomizing indexes for styles generations
	min := 0
	max := len(palettes) - 1
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(max-min) + min
	colorIndex := randomIndex
	log.Println("colorIndex", colorIndex)
	fontIndex := ((randomIndex % len(fonts)) + len(fonts)) % len(fonts)
	publicationIndex := ((randomIndex % len(publicationsMock)) + len(publicationsMock)) % len(publicationsMock)

	title = strings.TrimSpace(title)
	source = strings.TrimSpace(source)
	url = strings.TrimSpace(url)

	if title == "" {
		title = publicationsMock[publicationIndex].title
	}
	if source == "" {
		source = publicationsMock[publicationIndex].source
	}
	if url == "" {
		url = publicationsMock[publicationIndex].url
	}

	log.Println("title", title)
	log.Println("source", source)
	log.Println("url", url)
	fmt.Println("--------------------")

	// initialize new image
	dc := gg.NewContext(int(width), int(height))

	// experiment
	//image, _ := getImageFromURL(url)
	//if image != nil {
	//	if err := RenderBackgroundImage(dc, image, palettes[colorIndex].backgroundColor); err != nil {
	//		panic(err)
	//	}
	//}
	// end experiment

	// Set background color
	dc.DrawRectangle(0, 0, width, height)
	dc.SetColor(palettes[colorIndex].backgroundColor)
	dc.Fill()

	if err := RenderStoryTitle(dc, title, fontIndex, colorIndex); err != nil {
		panic(err)
	}

	if err := RenderStorySource(dc, title, source, url, fontIndex, colorIndex); err != nil {
		panic(err)
	}

	if err := RenderAuthorDescription(dc, colorIndex); err != nil {
		panic(err)
	}

	uuidWithHyphen := uuidx.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	if err := dc.SavePNG(fmt.Sprintf("assets/stories/%s.png", uuid)); err != nil {
		panic(err)
	}
	return nil
}
