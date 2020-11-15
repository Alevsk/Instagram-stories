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

package restapi

import (
	image_generator "github.com/Alevsk/instagram-stories/internal/image-generator"
	"github.com/Alevsk/instagram-stories/models"
	"github.com/Alevsk/instagram-stories/restapi/operations"
	"github.com/Alevsk/instagram-stories/restapi/operations/user_api"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func registerStoriesHandler(api *operations.InstagramStoriesAPI) {
	// create new story
	api.UserAPIStoryCreateHandler = user_api.StoryCreateHandlerFunc(func(params user_api.StoryCreateParams) middleware.Responder {
		if err := createStoryResponse(params); err != nil {
			return user_api.NewStoryCreateDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return user_api.NewStoryCreateOK()
	})
}

func createStory(title, source, url string) error {
	image_generator.RenderNewImage(title, source, url)
	return nil
}

func createStoryResponse(params user_api.StoryCreateParams) error {
	return createStory(params.Body.Title, params.Body.Source, params.Body.URL)
}