package checkins_e2e_test

import (
	"api-gym-on-go/tests/e2e/checkins/seed"
	"api-gym-on-go/tests/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckinsCountHistoryE2E(t *testing.T) {
	preCreateCheckin := true
	utils.ResetDb()
	token := utils.CreateAndAuthenticateUser("MEMBER")
	app := utils.SetupTestApp("checkins")
	seed.SeedCheckins(preCreateCheckin)
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))

	defer server.Close()

	t.Run("should be able to count history", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/checkin/history/count", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, _ := app.Test(req)

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("error reading response body: %v", err)
		}

		var responseData map[string]interface{}
		err = json.Unmarshal(respBody, &responseData)
		if err != nil {
			t.Fatalf("Erro ao parsear JSON: %v", err)
		}

		fmt.Println("#######", responseData, token)

		assert.Equal(t, 1, int(responseData["count"].(float64)), "count does not match")
		assert.Equalf(t, 200, resp.StatusCode, "get HTTP status 200")
	})
}
