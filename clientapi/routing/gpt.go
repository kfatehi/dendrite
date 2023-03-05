// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package routing

import (
	"errors"
	"net/http"

	"github.com/kfatehi/go-openai"
	"github.com/matrix-org/dendrite/clientapi/httputil"
	"github.com/matrix-org/util"
)

func Gpt(req *http.Request, openapi_key string) util.JSONResponse {

	body := openai.ChatRequest{}

	jsonErr := httputil.UnmarshalJSONRequest(req, &body)
	if jsonErr != nil {
		return util.JSONResponse{
			Code: http.StatusBadRequest,
			JSON: errorResponse(req.Context(), errors.New("error unmarshalling user request body"), "error unmarshalling user request body"),
		}
	}

	resp, err := openai.ChatCompletion(openapi_key, body)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusInternalServerError,
			JSON: errorResponse(req.Context(), err, "Error communicating with OpenAI"),
		}
	}

	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: resp,
	}
}
