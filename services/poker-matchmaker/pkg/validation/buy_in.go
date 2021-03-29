package validation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/serviceconfig"
)

//VerifyBuyIn verifies the buy in amount against the user wallet using the bank service
func VerifyBuyIn(userId string, buyIn int) error {
	client := &http.Client{
		Timeout: 25 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/api/bank/verifyAmount", viper.GetString(serviceconfig.BankServiceUrl)), nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("userId", userId)
	q.Add("amount", strconv.Itoa(buyIn))
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		text, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			return err
		}
		return fmt.Errorf("bank service responded with a non 200 response. %v", text)
	}

	var result dtos.VerifyAmount
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return err
	}

	err = res.Body.Close()
	if err != nil {
		return err
	}
	if !result.VerificationResult {
		return &errors.InvalidBuyIn{}
	}

	return nil
}
