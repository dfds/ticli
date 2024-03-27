package selfservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://api.hellman.oxygen.dfds.cloud/ssu/api"

// dev environment URL for testing
const devURL = "https://api.hellman.oxygen.dfds.cloud/ssu/dev-env/develop/api"

type SelfServiceClient struct {
	AccessToken string
}

func NewSelfServiceClient(accessToken string) *SelfServiceClient {
	return &SelfServiceClient{AccessToken: accessToken}
}

func (sc *SelfServiceClient) queryServer(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sc.AccessToken))
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		log.Fatal("Unexpected status code: ", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseBody
}

func (sc *SelfServiceClient) queryServerPost(url string, body []byte) []byte {

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sc.AccessToken))
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		log.Fatal("Unexpected status code: ", response.StatusCode)

	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseBody
}

func (sc *SelfServiceClient) GetCapabilities() CapabilitiesResponse {
	url := fmt.Sprintf("%s/capabilities", baseURL)
	responseBody := sc.queryServer(url)
	var capabilitiesResponse CapabilitiesResponse

	err := json.Unmarshal(responseBody, &capabilitiesResponse)
	if err != nil {
		log.Fatal(err)
	}

	return capabilitiesResponse
}

func (sc *SelfServiceClient) GetCapabilityByID(id string) CapabilityByIDResponse {

	url := fmt.Sprintf("%s/capabilities/%s", baseURL, id)
	responseBody := sc.queryServer(url)
	var capabilitiesResponse CapabilityByIDResponse

	err := json.Unmarshal(responseBody, &capabilitiesResponse)
	if err != nil {
		log.Fatal(err)
	}

	return capabilitiesResponse
}

func (sc *SelfServiceClient) GetECRRepositories() []EcrResponse {
	url := fmt.Sprintf("%s/ecr/repositories", baseURL)
	responseBody := sc.queryServer(url)
	var ecrResponse []EcrResponse

	err := json.Unmarshal(responseBody, &ecrResponse)
	if err != nil {
		log.Fatal(err)
	}

	return ecrResponse
}

func (sc *SelfServiceClient) GetTopics() TopicsResponse {
	url := fmt.Sprintf("%s/kafkatopics", baseURL)
	responseBody := sc.queryServer(url)
	var topicsResponse TopicsResponse

	err := json.Unmarshal(responseBody, &topicsResponse)
	if err != nil {
		log.Fatal(err)
	}

	return topicsResponse

}

func (sc *SelfServiceClient) CreateCapability(body CapabilityRequest) CapabilitiesResponse {
	url := fmt.Sprintf("%s/capabilities", baseURL)

	payload, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	responseBody := sc.queryServerPost(url, payload)
	var capabilitiesResponse CapabilitiesResponse

	err = json.Unmarshal(responseBody, &capabilitiesResponse)
	if err != nil {
		log.Fatal(err)
	}

	return capabilitiesResponse
}

func (sc *SelfServiceClient) CreateECRRepo(body CapabilityRequest) EcrResponse {
	url := fmt.Sprintf("%s/ecr/repositories", devURL)

	payload, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	responseBody := sc.queryServerPost(url, payload)
	var ecrResponse EcrResponse

	err = json.Unmarshal(responseBody, &ecrResponse)
	if err != nil {
		log.Fatal(err)
	}

	return ecrResponse
}

type CapabilitiesResponse struct {
	Items []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Status       string `json:"status"`
		Description  string `json:"description"`
		JSONMetadata string `json:"jsonMetadata"`
		AwsAccountID string `json:"awsAccountId"`
		Links        struct {
			Self struct {
				Href  string   `json:"href"`
				Rel   string   `json:"rel"`
				Allow []string `json:"allow"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"items"`
	Links struct {
		Self struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"self"`
	} `json:"_links"`
}

type EcrResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	RequestedAt time.Time `json:"requestedAt"`
	Id          struct {
	} `json:"id"`
}

type TopicsResponse struct {
	Items []struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		CapabilityId   string `json:"capabilityId"`
		KafkaClusterId string `json:"kafkaClusterId"`
		Partitions     int    `json:"partitions"`
		Retention      string `json:"retention"`
		Status         string `json:"status"`
		Links          struct {
			Self struct {
				Href  string   `json:"href"`
				Rel   string   `json:"rel"`
				Allow []string `json:"allow"`
			} `json:"self"`
			MessageContracts struct {
				Href  string   `json:"href"`
				Rel   string   `json:"rel"`
				Allow []string `json:"allow"`
			} `json:"messageContracts"`
			Consumers struct {
				Href  string   `json:"href"`
				Rel   string   `json:"rel"`
				Allow []string `json:"allow"`
			} `json:"consumers"`
			UpdateDescription *struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"updateDescription"`
		} `json:"_links"`
	} `json:"items"`
	Embedded struct {
		KafkaClusters struct {
			Items []struct {
				Id          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Links       struct {
					Self struct {
						Href  string   `json:"href"`
						Rel   string   `json:"rel"`
						Allow []string `json:"allow"`
					} `json:"self"`
				} `json:"_links"`
			} `json:"items"`
			Links struct {
				Self struct {
					Href  string   `json:"href"`
					Rel   string   `json:"rel"`
					Allow []string `json:"allow"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"kafkaClusters"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"self"`
	} `json:"_links"`
}

type CapabilityByIDResponse struct {
	Id                        string `json:"id"`
	Name                      string `json:"name"`
	Status                    string `json:"status"`
	Description               string `json:"description"`
	JsonMetadata              string `json:"jsonMetadata"`
	JsonMetadataSchemaVersion int    `json:"jsonMetadataSchemaVersion"`
	Links                     struct {
		Self struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"self"`
		Members struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"members"`
		Clusters struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"clusters"`
		MembershipApplications struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"membershipApplications"`
		LeaveCapability struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"leaveCapability"`
		AwsAccount struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"awsAccount"`
		RequestCapabilityDeletion struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"requestCapabilityDeletion"`
		CancelCapabilityDeletionRequest struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"cancelCapabilityDeletionRequest"`
		Metadata struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"metadata"`
		SetRequiredMetadata struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"setRequiredMetadata"`
		GetLinkedTeams struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"getLinkedTeams"`
		JoinCapability struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"joinCapability"`
		SendInvitations struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"sendInvitations"`
		ConfigurationLevel struct {
			Href  string   `json:"href"`
			Rel   string   `json:"rel"`
			Allow []string `json:"allow"`
		} `json:"configurationLevel"`
	} `json:"_links"`
}
type CapabilityRequest struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Invitees     []string `json:"invitees"`
	JsonMetadata string   `json:"jsonMetadata"`
}
type ECRRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
