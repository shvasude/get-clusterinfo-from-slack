package main

import (
	"encoding/json"
	"strings"
)

//SlackResponse is a struct based on the output of sbr request sent
type SlackResponse struct {
	Ok       bool `json:"ok"`
	Channels []struct {
		ID                      string        `json:"id"`
		Name                    string        `json:"name"`
		IsChannel               bool          `json:"is_channel"`
		IsGroup                 bool          `json:"is_group"`
		IsIm                    bool          `json:"is_im"`
		Created                 int           `json:"created"`
		IsArchived              bool          `json:"is_archived"`
		IsGeneral               bool          `json:"is_general"`
		Unlinked                int           `json:"unlinked"`
		NameNormalized          string        `json:"name_normalized"`
		IsShared                bool          `json:"is_shared"`
		ParentConversation      interface{}   `json:"parent_conversation"`
		Creator                 string        `json:"creator"`
		IsExtShared             bool          `json:"is_ext_shared"`
		IsOrgShared             bool          `json:"is_org_shared"`
		SharedTeamIds           []string      `json:"shared_team_ids"`
		PendingShared           []interface{} `json:"pending_shared"`
		PendingConnectedTeamIds []interface{} `json:"pending_connected_team_ids"`
		IsPendingExtShared      bool          `json:"is_pending_ext_shared"`
		IsMember                bool          `json:"is_member"`
		IsPrivate               bool          `json:"is_private"`
		IsMpim                  bool          `json:"is_mpim"`
		Topic                   struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"topic"`
		Purpose struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"purpose"`
		PreviousNames []string `json:"previous_names"`
		NumMembers    int      `json:"num_members"`
	} `json:"channels"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

//UnmarshalJSONData unmarshall the data in form of json to a struct
func UnmarshalJSONData(jsonData string) SlackResponse {
	res := SlackResponse{}
	if strings.Contains(jsonData, "'") {
		jsonData = strings.Trim(jsonData, "'")
	}
	json.Unmarshal([]byte(jsonData), &res)
	return res
}
