package statuspage

import (
	"fmt"
	"github.com/broadinstitute/revere/internal/shared"
	"github.com/broadinstitute/revere/internal/statuspage/statuspagetypes"
	"github.com/go-resty/resty/v2"
)

// GetComponents provides a slice of all components on the remote page
func GetComponents(client *resty.Client, pageID string) (*[]statuspagetypes.Component, error) {
	resp, err := client.R().
		SetResult([]statuspagetypes.Component{}).
		Get(fmt.Sprintf("/pages/%s/components", pageID))
	if err = shared.CheckResponse(resp, err); err != nil {
		return nil, err
	}
	return resp.Result().(*[]statuspagetypes.Component), nil
}

// PostComponent creates a new component on the remote page
func PostComponent(client *resty.Client, pageID string, component statuspagetypes.Component) (*statuspagetypes.Component, error) {
	resp, err := client.R().
		SetResult(statuspagetypes.Component{}).
		SetBody(map[string]interface{}{"component": component.ToRequest()}).
		Post(fmt.Sprintf("/pages/%s/components", pageID))
	if err = shared.CheckResponse(resp, err); err != nil {
		return nil, err
	}
	return resp.Result().(*statuspagetypes.Component), nil
}

// PatchComponent updates an existing component on the remote page by the component's ID, not name
func PatchComponent(client *resty.Client, pageID string, componentID string, component statuspagetypes.Component) (*statuspagetypes.Component, error) {
	resp, err := client.R().
		SetResult(statuspagetypes.Component{}).
		SetBody(map[string]interface{}{"component": component.ToRequest()}).
		Patch(fmt.Sprintf("/pages/%s/components/%s", pageID, componentID))
	if err = shared.CheckResponse(resp, err); err != nil {
		return nil, err
	}
	return resp.Result().(*statuspagetypes.Component), nil
}

// DeleteComponent deletes an existing component on the remote page by the component's ID, not name
func DeleteComponent(client *resty.Client, pageID string, componentID string) error {
	resp, err := client.R().
		Delete(fmt.Sprintf("/pages/%s/components/%s", pageID, componentID))
	if err = shared.CheckResponse(resp, err); err != nil {
		return err
	}
	return nil
}
