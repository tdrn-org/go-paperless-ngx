// generated from api.ClientWithResponsesInterface
package paperlessngx

import (
	"context"
	"io"

	"github.com/tdrn-org/go-paperless-ngx/api"
)

func (client *Client) AcknowledgeTasksWithBody(ctx context.Context, params *api.AcknowledgeTasksParams, contentType string, body io.Reader) (*api.AcknowledgeTasksResponse, error) {
	response, err := client.apiClient.AcknowledgeTasksWithBodyWithResponse(ctx, params, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) AcknowledgeTasks(ctx context.Context, params *api.AcknowledgeTasksParams, body api.AcknowledgeTasksJSONRequestBody) (*api.AcknowledgeTasksResponse, error) {
	response, err := client.apiClient.AcknowledgeTasksWithResponse(ctx, params, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) BulkEditObjectsWithBody(ctx context.Context, contentType string, body io.Reader) (*api.BulkEditObjectsResponse, error) {
	response, err := client.apiClient.BulkEditObjectsWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) BulkEditObjects(ctx context.Context, body api.BulkEditObjectsJSONRequestBody) (*api.BulkEditObjectsResponse, error) {
	response, err := client.apiClient.BulkEditObjectsWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) BulkEditWithBody(ctx context.Context, contentType string, body io.Reader) (*api.BulkEditResponse, error) {
	response, err := client.apiClient.BulkEditWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) BulkEdit(ctx context.Context, body api.BulkEditJSONRequestBody) (*api.BulkEditResponse, error) {
	response, err := client.apiClient.BulkEditWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigDestroy(ctx context.Context, id int) (*api.ConfigDestroyResponse, error) {
	response, err := client.apiClient.ConfigDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigList(ctx context.Context) (*api.ConfigListResponse, error) {
	response, err := client.apiClient.ConfigListWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.ConfigPartialUpdateResponse, error) {
	response, err := client.apiClient.ConfigPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.ConfigPartialUpdateFormdataRequestBody) (*api.ConfigPartialUpdateResponse, error) {
	response, err := client.apiClient.ConfigPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigPartialUpdate(ctx context.Context, id int, body api.ConfigPartialUpdateJSONRequestBody) (*api.ConfigPartialUpdateResponse, error) {
	response, err := client.apiClient.ConfigPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigRetrieve(ctx context.Context, id int) (*api.ConfigRetrieveResponse, error) {
	response, err := client.apiClient.ConfigRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.ConfigUpdateResponse, error) {
	response, err := client.apiClient.ConfigUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigUpdateWithFormdataBody(ctx context.Context, id int, body api.ConfigUpdateFormdataRequestBody) (*api.ConfigUpdateResponse, error) {
	response, err := client.apiClient.ConfigUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ConfigUpdate(ctx context.Context, id int, body api.ConfigUpdateJSONRequestBody) (*api.ConfigUpdateResponse, error) {
	response, err := client.apiClient.ConfigUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.CorrespondentsCreateResponse, error) {
	response, err := client.apiClient.CorrespondentsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsCreateWithFormdataBody(ctx context.Context, body api.CorrespondentsCreateFormdataRequestBody) (*api.CorrespondentsCreateResponse, error) {
	response, err := client.apiClient.CorrespondentsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsCreate(ctx context.Context, body api.CorrespondentsCreateJSONRequestBody) (*api.CorrespondentsCreateResponse, error) {
	response, err := client.apiClient.CorrespondentsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsDestroy(ctx context.Context, id int) (*api.CorrespondentsDestroyResponse, error) {
	response, err := client.apiClient.CorrespondentsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsList(ctx context.Context, params *api.CorrespondentsListParams) (*api.CorrespondentsListResponse, error) {
	response, err := client.apiClient.CorrespondentsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.CorrespondentsPartialUpdateResponse, error) {
	response, err := client.apiClient.CorrespondentsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.CorrespondentsPartialUpdateFormdataRequestBody) (*api.CorrespondentsPartialUpdateResponse, error) {
	response, err := client.apiClient.CorrespondentsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsPartialUpdate(ctx context.Context, id int, body api.CorrespondentsPartialUpdateJSONRequestBody) (*api.CorrespondentsPartialUpdateResponse, error) {
	response, err := client.apiClient.CorrespondentsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsRetrieve(ctx context.Context, id int, params *api.CorrespondentsRetrieveParams) (*api.CorrespondentsRetrieveResponse, error) {
	response, err := client.apiClient.CorrespondentsRetrieveWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.CorrespondentsUpdateResponse, error) {
	response, err := client.apiClient.CorrespondentsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsUpdateWithFormdataBody(ctx context.Context, id int, body api.CorrespondentsUpdateFormdataRequestBody) (*api.CorrespondentsUpdateResponse, error) {
	response, err := client.apiClient.CorrespondentsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CorrespondentsUpdate(ctx context.Context, id int, body api.CorrespondentsUpdateJSONRequestBody) (*api.CorrespondentsUpdateResponse, error) {
	response, err := client.apiClient.CorrespondentsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.CustomFieldsCreateResponse, error) {
	response, err := client.apiClient.CustomFieldsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsCreateWithFormdataBody(ctx context.Context, body api.CustomFieldsCreateFormdataRequestBody) (*api.CustomFieldsCreateResponse, error) {
	response, err := client.apiClient.CustomFieldsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsCreate(ctx context.Context, body api.CustomFieldsCreateJSONRequestBody) (*api.CustomFieldsCreateResponse, error) {
	response, err := client.apiClient.CustomFieldsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsDestroy(ctx context.Context, id int) (*api.CustomFieldsDestroyResponse, error) {
	response, err := client.apiClient.CustomFieldsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsList(ctx context.Context, params *api.CustomFieldsListParams) (*api.CustomFieldsListResponse, error) {
	response, err := client.apiClient.CustomFieldsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.CustomFieldsPartialUpdateResponse, error) {
	response, err := client.apiClient.CustomFieldsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.CustomFieldsPartialUpdateFormdataRequestBody) (*api.CustomFieldsPartialUpdateResponse, error) {
	response, err := client.apiClient.CustomFieldsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsPartialUpdate(ctx context.Context, id int, body api.CustomFieldsPartialUpdateJSONRequestBody) (*api.CustomFieldsPartialUpdateResponse, error) {
	response, err := client.apiClient.CustomFieldsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsRetrieve(ctx context.Context, id int) (*api.CustomFieldsRetrieveResponse, error) {
	response, err := client.apiClient.CustomFieldsRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.CustomFieldsUpdateResponse, error) {
	response, err := client.apiClient.CustomFieldsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsUpdateWithFormdataBody(ctx context.Context, id int, body api.CustomFieldsUpdateFormdataRequestBody) (*api.CustomFieldsUpdateResponse, error) {
	response, err := client.apiClient.CustomFieldsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) CustomFieldsUpdate(ctx context.Context, id int, body api.CustomFieldsUpdateJSONRequestBody) (*api.CustomFieldsUpdateResponse, error) {
	response, err := client.apiClient.CustomFieldsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentShareLinks(ctx context.Context, id string) (*api.DocumentShareLinksResponse, error) {
	response, err := client.apiClient.DocumentShareLinksWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.DocumentTypesCreateResponse, error) {
	response, err := client.apiClient.DocumentTypesCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesCreateWithFormdataBody(ctx context.Context, body api.DocumentTypesCreateFormdataRequestBody) (*api.DocumentTypesCreateResponse, error) {
	response, err := client.apiClient.DocumentTypesCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesCreate(ctx context.Context, body api.DocumentTypesCreateJSONRequestBody) (*api.DocumentTypesCreateResponse, error) {
	response, err := client.apiClient.DocumentTypesCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesDestroy(ctx context.Context, id int) (*api.DocumentTypesDestroyResponse, error) {
	response, err := client.apiClient.DocumentTypesDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesList(ctx context.Context, params *api.DocumentTypesListParams) (*api.DocumentTypesListResponse, error) {
	response, err := client.apiClient.DocumentTypesListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.DocumentTypesPartialUpdateResponse, error) {
	response, err := client.apiClient.DocumentTypesPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.DocumentTypesPartialUpdateFormdataRequestBody) (*api.DocumentTypesPartialUpdateResponse, error) {
	response, err := client.apiClient.DocumentTypesPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesPartialUpdate(ctx context.Context, id int, body api.DocumentTypesPartialUpdateJSONRequestBody) (*api.DocumentTypesPartialUpdateResponse, error) {
	response, err := client.apiClient.DocumentTypesPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesRetrieve(ctx context.Context, id int, params *api.DocumentTypesRetrieveParams) (*api.DocumentTypesRetrieveResponse, error) {
	response, err := client.apiClient.DocumentTypesRetrieveWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.DocumentTypesUpdateResponse, error) {
	response, err := client.apiClient.DocumentTypesUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesUpdateWithFormdataBody(ctx context.Context, id int, body api.DocumentTypesUpdateFormdataRequestBody) (*api.DocumentTypesUpdateResponse, error) {
	response, err := client.apiClient.DocumentTypesUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentTypesUpdate(ctx context.Context, id int, body api.DocumentTypesUpdateJSONRequestBody) (*api.DocumentTypesUpdateResponse, error) {
	response, err := client.apiClient.DocumentTypesUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsBulkDownloadCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.DocumentsBulkDownloadCreateResponse, error) {
	response, err := client.apiClient.DocumentsBulkDownloadCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsBulkDownloadCreate(ctx context.Context, body api.DocumentsBulkDownloadCreateJSONRequestBody) (*api.DocumentsBulkDownloadCreateResponse, error) {
	response, err := client.apiClient.DocumentsBulkDownloadCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsDestroy(ctx context.Context, id int) (*api.DocumentsDestroyResponse, error) {
	response, err := client.apiClient.DocumentsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsDownloadRetrieve(ctx context.Context, id int, params *api.DocumentsDownloadRetrieveParams) (*api.DocumentsDownloadRetrieveResponse, error) {
	response, err := client.apiClient.DocumentsDownloadRetrieveWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsEmailCreateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.DocumentsEmailCreateResponse, error) {
	response, err := client.apiClient.DocumentsEmailCreateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsEmailCreateWithFormdataBody(ctx context.Context, id int, body api.DocumentsEmailCreateFormdataRequestBody) (*api.DocumentsEmailCreateResponse, error) {
	response, err := client.apiClient.DocumentsEmailCreateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsEmailCreate(ctx context.Context, id int, body api.DocumentsEmailCreateJSONRequestBody) (*api.DocumentsEmailCreateResponse, error) {
	response, err := client.apiClient.DocumentsEmailCreateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsHistoryList(ctx context.Context, id int, params *api.DocumentsHistoryListParams) (*api.DocumentsHistoryListResponse, error) {
	response, err := client.apiClient.DocumentsHistoryListWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsList(ctx context.Context, params *api.DocumentsListParams) (*api.DocumentsListResponse, error) {
	response, err := client.apiClient.DocumentsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsMetadataRetrieve(ctx context.Context, id int) (*api.DocumentsMetadataRetrieveResponse, error) {
	response, err := client.apiClient.DocumentsMetadataRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsNextAsnRetrieve(ctx context.Context) (*api.DocumentsNextAsnRetrieveResponse, error) {
	response, err := client.apiClient.DocumentsNextAsnRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsNotesCreateWithBody(ctx context.Context, id int, params *api.DocumentsNotesCreateParams, contentType string, body io.Reader) (*api.DocumentsNotesCreateResponse, error) {
	response, err := client.apiClient.DocumentsNotesCreateWithBodyWithResponse(ctx, id, params, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsNotesCreateWithFormdataBody(ctx context.Context, id int, params *api.DocumentsNotesCreateParams, body api.DocumentsNotesCreateFormdataRequestBody) (*api.DocumentsNotesCreateResponse, error) {
	response, err := client.apiClient.DocumentsNotesCreateWithFormdataBodyWithResponse(ctx, id, params, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsNotesCreate(ctx context.Context, id int, params *api.DocumentsNotesCreateParams, body api.DocumentsNotesCreateJSONRequestBody) (*api.DocumentsNotesCreateResponse, error) {
	response, err := client.apiClient.DocumentsNotesCreateWithResponse(ctx, id, params, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsNotesDestroy(ctx context.Context, id int, params *api.DocumentsNotesDestroyParams) (*api.DocumentsNotesDestroyResponse, error) {
	response, err := client.apiClient.DocumentsNotesDestroyWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsNotesList(ctx context.Context, id int, params *api.DocumentsNotesListParams) (*api.DocumentsNotesListResponse, error) {
	response, err := client.apiClient.DocumentsNotesListWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.DocumentsPartialUpdateResponse, error) {
	response, err := client.apiClient.DocumentsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.DocumentsPartialUpdateFormdataRequestBody) (*api.DocumentsPartialUpdateResponse, error) {
	response, err := client.apiClient.DocumentsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsPartialUpdate(ctx context.Context, id int, body api.DocumentsPartialUpdateJSONRequestBody) (*api.DocumentsPartialUpdateResponse, error) {
	response, err := client.apiClient.DocumentsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsPostDocumentCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.DocumentsPostDocumentCreateResponse, error) {
	response, err := client.apiClient.DocumentsPostDocumentCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsPreviewRetrieve(ctx context.Context, id int) (*api.DocumentsPreviewRetrieveResponse, error) {
	response, err := client.apiClient.DocumentsPreviewRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsRetrieve(ctx context.Context, id int, params *api.DocumentsRetrieveParams) (*api.DocumentsRetrieveResponse, error) {
	response, err := client.apiClient.DocumentsRetrieveWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsSelectionDataCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.DocumentsSelectionDataCreateResponse, error) {
	response, err := client.apiClient.DocumentsSelectionDataCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsSelectionDataCreate(ctx context.Context, body api.DocumentsSelectionDataCreateJSONRequestBody) (*api.DocumentsSelectionDataCreateResponse, error) {
	response, err := client.apiClient.DocumentsSelectionDataCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsSuggestionsRetrieve(ctx context.Context, id int) (*api.DocumentsSuggestionsRetrieveResponse, error) {
	response, err := client.apiClient.DocumentsSuggestionsRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsThumbRetrieve(ctx context.Context, id int) (*api.DocumentsThumbRetrieveResponse, error) {
	response, err := client.apiClient.DocumentsThumbRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.DocumentsUpdateResponse, error) {
	response, err := client.apiClient.DocumentsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsUpdateWithFormdataBody(ctx context.Context, id int, body api.DocumentsUpdateFormdataRequestBody) (*api.DocumentsUpdateResponse, error) {
	response, err := client.apiClient.DocumentsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) DocumentsUpdate(ctx context.Context, id int, body api.DocumentsUpdateJSONRequestBody) (*api.DocumentsUpdateResponse, error) {
	response, err := client.apiClient.DocumentsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) EmailDocumentsWithBody(ctx context.Context, contentType string, body io.Reader) (*api.EmailDocumentsResponse2, error) {
	response, err := client.apiClient.EmailDocumentsWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) EmailDocumentsWithFormdataBody(ctx context.Context, body api.EmailDocumentsFormdataRequestBody) (*api.EmailDocumentsResponse2, error) {
	response, err := client.apiClient.EmailDocumentsWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) EmailDocuments(ctx context.Context, body api.EmailDocumentsJSONRequestBody) (*api.EmailDocumentsResponse2, error) {
	response, err := client.apiClient.EmailDocumentsWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.GroupsCreateResponse, error) {
	response, err := client.apiClient.GroupsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsCreateWithFormdataBody(ctx context.Context, body api.GroupsCreateFormdataRequestBody) (*api.GroupsCreateResponse, error) {
	response, err := client.apiClient.GroupsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsCreate(ctx context.Context, body api.GroupsCreateJSONRequestBody) (*api.GroupsCreateResponse, error) {
	response, err := client.apiClient.GroupsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsDestroy(ctx context.Context, id int) (*api.GroupsDestroyResponse, error) {
	response, err := client.apiClient.GroupsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsList(ctx context.Context, params *api.GroupsListParams) (*api.GroupsListResponse, error) {
	response, err := client.apiClient.GroupsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.GroupsPartialUpdateResponse, error) {
	response, err := client.apiClient.GroupsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.GroupsPartialUpdateFormdataRequestBody) (*api.GroupsPartialUpdateResponse, error) {
	response, err := client.apiClient.GroupsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsPartialUpdate(ctx context.Context, id int, body api.GroupsPartialUpdateJSONRequestBody) (*api.GroupsPartialUpdateResponse, error) {
	response, err := client.apiClient.GroupsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsRetrieve(ctx context.Context, id int) (*api.GroupsRetrieveResponse, error) {
	response, err := client.apiClient.GroupsRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.GroupsUpdateResponse, error) {
	response, err := client.apiClient.GroupsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsUpdateWithFormdataBody(ctx context.Context, id int, body api.GroupsUpdateFormdataRequestBody) (*api.GroupsUpdateResponse, error) {
	response, err := client.apiClient.GroupsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) GroupsUpdate(ctx context.Context, id int, body api.GroupsUpdateJSONRequestBody) (*api.GroupsUpdateResponse, error) {
	response, err := client.apiClient.GroupsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) LogsList(ctx context.Context) (*api.LogsListResponse, error) {
	response, err := client.apiClient.LogsListWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountProcessWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.MailAccountProcessResponse2, error) {
	response, err := client.apiClient.MailAccountProcessWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountProcessWithFormdataBody(ctx context.Context, id int, body api.MailAccountProcessFormdataRequestBody) (*api.MailAccountProcessResponse2, error) {
	response, err := client.apiClient.MailAccountProcessWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountProcess(ctx context.Context, id int, body api.MailAccountProcessJSONRequestBody) (*api.MailAccountProcessResponse2, error) {
	response, err := client.apiClient.MailAccountProcessWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountTestWithBody(ctx context.Context, contentType string, body io.Reader) (*api.MailAccountTestResponse2, error) {
	response, err := client.apiClient.MailAccountTestWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountTestWithFormdataBody(ctx context.Context, body api.MailAccountTestFormdataRequestBody) (*api.MailAccountTestResponse2, error) {
	response, err := client.apiClient.MailAccountTestWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountTest(ctx context.Context, body api.MailAccountTestJSONRequestBody) (*api.MailAccountTestResponse2, error) {
	response, err := client.apiClient.MailAccountTestWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.MailAccountsCreateResponse, error) {
	response, err := client.apiClient.MailAccountsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsCreateWithFormdataBody(ctx context.Context, body api.MailAccountsCreateFormdataRequestBody) (*api.MailAccountsCreateResponse, error) {
	response, err := client.apiClient.MailAccountsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsCreate(ctx context.Context, body api.MailAccountsCreateJSONRequestBody) (*api.MailAccountsCreateResponse, error) {
	response, err := client.apiClient.MailAccountsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsDestroy(ctx context.Context, id int) (*api.MailAccountsDestroyResponse, error) {
	response, err := client.apiClient.MailAccountsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsList(ctx context.Context, params *api.MailAccountsListParams) (*api.MailAccountsListResponse, error) {
	response, err := client.apiClient.MailAccountsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.MailAccountsPartialUpdateResponse, error) {
	response, err := client.apiClient.MailAccountsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.MailAccountsPartialUpdateFormdataRequestBody) (*api.MailAccountsPartialUpdateResponse, error) {
	response, err := client.apiClient.MailAccountsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsPartialUpdate(ctx context.Context, id int, body api.MailAccountsPartialUpdateJSONRequestBody) (*api.MailAccountsPartialUpdateResponse, error) {
	response, err := client.apiClient.MailAccountsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsRetrieve(ctx context.Context, id int) (*api.MailAccountsRetrieveResponse, error) {
	response, err := client.apiClient.MailAccountsRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.MailAccountsUpdateResponse, error) {
	response, err := client.apiClient.MailAccountsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsUpdateWithFormdataBody(ctx context.Context, id int, body api.MailAccountsUpdateFormdataRequestBody) (*api.MailAccountsUpdateResponse, error) {
	response, err := client.apiClient.MailAccountsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailAccountsUpdate(ctx context.Context, id int, body api.MailAccountsUpdateJSONRequestBody) (*api.MailAccountsUpdateResponse, error) {
	response, err := client.apiClient.MailAccountsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.MailRulesCreateResponse, error) {
	response, err := client.apiClient.MailRulesCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesCreateWithFormdataBody(ctx context.Context, body api.MailRulesCreateFormdataRequestBody) (*api.MailRulesCreateResponse, error) {
	response, err := client.apiClient.MailRulesCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesCreate(ctx context.Context, body api.MailRulesCreateJSONRequestBody) (*api.MailRulesCreateResponse, error) {
	response, err := client.apiClient.MailRulesCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesDestroy(ctx context.Context, id int) (*api.MailRulesDestroyResponse, error) {
	response, err := client.apiClient.MailRulesDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesList(ctx context.Context, params *api.MailRulesListParams) (*api.MailRulesListResponse, error) {
	response, err := client.apiClient.MailRulesListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.MailRulesPartialUpdateResponse, error) {
	response, err := client.apiClient.MailRulesPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.MailRulesPartialUpdateFormdataRequestBody) (*api.MailRulesPartialUpdateResponse, error) {
	response, err := client.apiClient.MailRulesPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesPartialUpdate(ctx context.Context, id int, body api.MailRulesPartialUpdateJSONRequestBody) (*api.MailRulesPartialUpdateResponse, error) {
	response, err := client.apiClient.MailRulesPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesRetrieve(ctx context.Context, id int) (*api.MailRulesRetrieveResponse, error) {
	response, err := client.apiClient.MailRulesRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.MailRulesUpdateResponse, error) {
	response, err := client.apiClient.MailRulesUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesUpdateWithFormdataBody(ctx context.Context, id int, body api.MailRulesUpdateFormdataRequestBody) (*api.MailRulesUpdateResponse, error) {
	response, err := client.apiClient.MailRulesUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) MailRulesUpdate(ctx context.Context, id int, body api.MailRulesUpdateJSONRequestBody) (*api.MailRulesUpdateResponse, error) {
	response, err := client.apiClient.MailRulesUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) OauthCallbackRetrieve(ctx context.Context) (*api.OauthCallbackRetrieveResponse, error) {
	response, err := client.apiClient.OauthCallbackRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProcessedMailBulkDeleteCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.ProcessedMailBulkDeleteCreateResponse, error) {
	response, err := client.apiClient.ProcessedMailBulkDeleteCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProcessedMailBulkDeleteCreateWithFormdataBody(ctx context.Context, body api.ProcessedMailBulkDeleteCreateFormdataRequestBody) (*api.ProcessedMailBulkDeleteCreateResponse, error) {
	response, err := client.apiClient.ProcessedMailBulkDeleteCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProcessedMailBulkDeleteCreate(ctx context.Context, body api.ProcessedMailBulkDeleteCreateJSONRequestBody) (*api.ProcessedMailBulkDeleteCreateResponse, error) {
	response, err := client.apiClient.ProcessedMailBulkDeleteCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProcessedMailList(ctx context.Context, params *api.ProcessedMailListParams) (*api.ProcessedMailListResponse, error) {
	response, err := client.apiClient.ProcessedMailListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProcessedMailRetrieve(ctx context.Context, id int) (*api.ProcessedMailRetrieveResponse, error) {
	response, err := client.apiClient.ProcessedMailRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileDisconnectSocialAccountCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.ProfileDisconnectSocialAccountCreateResponse, error) {
	response, err := client.apiClient.ProfileDisconnectSocialAccountCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileDisconnectSocialAccountCreate(ctx context.Context, body api.ProfileDisconnectSocialAccountCreateJSONRequestBody) (*api.ProfileDisconnectSocialAccountCreateResponse, error) {
	response, err := client.apiClient.ProfileDisconnectSocialAccountCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileGenerateAuthTokenCreate(ctx context.Context) (*api.ProfileGenerateAuthTokenCreateResponse, error) {
	response, err := client.apiClient.ProfileGenerateAuthTokenCreateWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfilePartialUpdateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.ProfilePartialUpdateResponse, error) {
	response, err := client.apiClient.ProfilePartialUpdateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfilePartialUpdateWithFormdataBody(ctx context.Context, body api.ProfilePartialUpdateFormdataRequestBody) (*api.ProfilePartialUpdateResponse, error) {
	response, err := client.apiClient.ProfilePartialUpdateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfilePartialUpdate(ctx context.Context, body api.ProfilePartialUpdateJSONRequestBody) (*api.ProfilePartialUpdateResponse, error) {
	response, err := client.apiClient.ProfilePartialUpdateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileRetrieve(ctx context.Context) (*api.ProfileRetrieveResponse, error) {
	response, err := client.apiClient.ProfileRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileSocialAccountProvidersRetrieve(ctx context.Context) (*api.ProfileSocialAccountProvidersRetrieveResponse, error) {
	response, err := client.apiClient.ProfileSocialAccountProvidersRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileTotpCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.ProfileTotpCreateResponse, error) {
	response, err := client.apiClient.ProfileTotpCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileTotpCreate(ctx context.Context, body api.ProfileTotpCreateJSONRequestBody) (*api.ProfileTotpCreateResponse, error) {
	response, err := client.apiClient.ProfileTotpCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileTotpDestroy(ctx context.Context) (*api.ProfileTotpDestroyResponse, error) {
	response, err := client.apiClient.ProfileTotpDestroyWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ProfileTotpRetrieve(ctx context.Context) (*api.ProfileTotpRetrieveResponse, error) {
	response, err := client.apiClient.ProfileTotpRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) RemoteVersionRetrieve(ctx context.Context) (*api.RemoteVersionRetrieveResponse, error) {
	response, err := client.apiClient.RemoteVersionRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) RetrieveLog(ctx context.Context, id string, params *api.RetrieveLogParams) (*api.RetrieveLogResponse, error) {
	response, err := client.apiClient.RetrieveLogWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.SavedViewsCreateResponse, error) {
	response, err := client.apiClient.SavedViewsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsCreateWithFormdataBody(ctx context.Context, body api.SavedViewsCreateFormdataRequestBody) (*api.SavedViewsCreateResponse, error) {
	response, err := client.apiClient.SavedViewsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsCreate(ctx context.Context, body api.SavedViewsCreateJSONRequestBody) (*api.SavedViewsCreateResponse, error) {
	response, err := client.apiClient.SavedViewsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsDestroy(ctx context.Context, id int) (*api.SavedViewsDestroyResponse, error) {
	response, err := client.apiClient.SavedViewsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsList(ctx context.Context, params *api.SavedViewsListParams) (*api.SavedViewsListResponse, error) {
	response, err := client.apiClient.SavedViewsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.SavedViewsPartialUpdateResponse, error) {
	response, err := client.apiClient.SavedViewsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.SavedViewsPartialUpdateFormdataRequestBody) (*api.SavedViewsPartialUpdateResponse, error) {
	response, err := client.apiClient.SavedViewsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsPartialUpdate(ctx context.Context, id int, body api.SavedViewsPartialUpdateJSONRequestBody) (*api.SavedViewsPartialUpdateResponse, error) {
	response, err := client.apiClient.SavedViewsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsRetrieve(ctx context.Context, id int) (*api.SavedViewsRetrieveResponse, error) {
	response, err := client.apiClient.SavedViewsRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.SavedViewsUpdateResponse, error) {
	response, err := client.apiClient.SavedViewsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsUpdateWithFormdataBody(ctx context.Context, id int, body api.SavedViewsUpdateFormdataRequestBody) (*api.SavedViewsUpdateResponse, error) {
	response, err := client.apiClient.SavedViewsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SavedViewsUpdate(ctx context.Context, id int, body api.SavedViewsUpdateJSONRequestBody) (*api.SavedViewsUpdateResponse, error) {
	response, err := client.apiClient.SavedViewsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SearchAutocompleteList(ctx context.Context, params *api.SearchAutocompleteListParams) (*api.SearchAutocompleteListResponse, error) {
	response, err := client.apiClient.SearchAutocompleteListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) SearchRetrieve(ctx context.Context, params *api.SearchRetrieveParams) (*api.SearchRetrieveResponse, error) {
	response, err := client.apiClient.SearchRetrieveWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ShareLinksCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.ShareLinksCreateResponse, error) {
	response, err := client.apiClient.ShareLinksCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ShareLinksCreateWithFormdataBody(ctx context.Context, body api.ShareLinksCreateFormdataRequestBody) (*api.ShareLinksCreateResponse, error) {
	response, err := client.apiClient.ShareLinksCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ShareLinksCreate(ctx context.Context, body api.ShareLinksCreateJSONRequestBody) (*api.ShareLinksCreateResponse, error) {
	response, err := client.apiClient.ShareLinksCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ShareLinksDestroy(ctx context.Context, id int) (*api.ShareLinksDestroyResponse, error) {
	response, err := client.apiClient.ShareLinksDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ShareLinksList(ctx context.Context, params *api.ShareLinksListParams) (*api.ShareLinksListResponse, error) {
	response, err := client.apiClient.ShareLinksListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) ShareLinksRetrieve(ctx context.Context, id int) (*api.ShareLinksRetrieveResponse, error) {
	response, err := client.apiClient.ShareLinksRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StatisticsRetrieve(ctx context.Context) (*api.StatisticsRetrieveResponse, error) {
	response, err := client.apiClient.StatisticsRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StatusRetrieve(ctx context.Context) (*api.StatusRetrieveResponse, error) {
	response, err := client.apiClient.StatusRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.StoragePathsCreateResponse, error) {
	response, err := client.apiClient.StoragePathsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsCreateWithFormdataBody(ctx context.Context, body api.StoragePathsCreateFormdataRequestBody) (*api.StoragePathsCreateResponse, error) {
	response, err := client.apiClient.StoragePathsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsCreate(ctx context.Context, body api.StoragePathsCreateJSONRequestBody) (*api.StoragePathsCreateResponse, error) {
	response, err := client.apiClient.StoragePathsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsDestroy(ctx context.Context, id int) (*api.StoragePathsDestroyResponse, error) {
	response, err := client.apiClient.StoragePathsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsList(ctx context.Context, params *api.StoragePathsListParams) (*api.StoragePathsListResponse, error) {
	response, err := client.apiClient.StoragePathsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.StoragePathsPartialUpdateResponse, error) {
	response, err := client.apiClient.StoragePathsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.StoragePathsPartialUpdateFormdataRequestBody) (*api.StoragePathsPartialUpdateResponse, error) {
	response, err := client.apiClient.StoragePathsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsPartialUpdate(ctx context.Context, id int, body api.StoragePathsPartialUpdateJSONRequestBody) (*api.StoragePathsPartialUpdateResponse, error) {
	response, err := client.apiClient.StoragePathsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsRetrieve(ctx context.Context, id int, params *api.StoragePathsRetrieveParams) (*api.StoragePathsRetrieveResponse, error) {
	response, err := client.apiClient.StoragePathsRetrieveWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsTestCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.StoragePathsTestCreateResponse, error) {
	response, err := client.apiClient.StoragePathsTestCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsTestCreateWithFormdataBody(ctx context.Context, body api.StoragePathsTestCreateFormdataRequestBody) (*api.StoragePathsTestCreateResponse, error) {
	response, err := client.apiClient.StoragePathsTestCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsTestCreate(ctx context.Context, body api.StoragePathsTestCreateJSONRequestBody) (*api.StoragePathsTestCreateResponse, error) {
	response, err := client.apiClient.StoragePathsTestCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.StoragePathsUpdateResponse, error) {
	response, err := client.apiClient.StoragePathsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsUpdateWithFormdataBody(ctx context.Context, id int, body api.StoragePathsUpdateFormdataRequestBody) (*api.StoragePathsUpdateResponse, error) {
	response, err := client.apiClient.StoragePathsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) StoragePathsUpdate(ctx context.Context, id int, body api.StoragePathsUpdateJSONRequestBody) (*api.StoragePathsUpdateResponse, error) {
	response, err := client.apiClient.StoragePathsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.TagsCreateResponse, error) {
	response, err := client.apiClient.TagsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsCreateWithFormdataBody(ctx context.Context, body api.TagsCreateFormdataRequestBody) (*api.TagsCreateResponse, error) {
	response, err := client.apiClient.TagsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsCreate(ctx context.Context, body api.TagsCreateJSONRequestBody) (*api.TagsCreateResponse, error) {
	response, err := client.apiClient.TagsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsDestroy(ctx context.Context, id int) (*api.TagsDestroyResponse, error) {
	response, err := client.apiClient.TagsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsList(ctx context.Context, params *api.TagsListParams) (*api.TagsListResponse, error) {
	response, err := client.apiClient.TagsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.TagsPartialUpdateResponse, error) {
	response, err := client.apiClient.TagsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.TagsPartialUpdateFormdataRequestBody) (*api.TagsPartialUpdateResponse, error) {
	response, err := client.apiClient.TagsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsPartialUpdate(ctx context.Context, id int, body api.TagsPartialUpdateJSONRequestBody) (*api.TagsPartialUpdateResponse, error) {
	response, err := client.apiClient.TagsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsRetrieve(ctx context.Context, id int, params *api.TagsRetrieveParams) (*api.TagsRetrieveResponse, error) {
	response, err := client.apiClient.TagsRetrieveWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.TagsUpdateResponse, error) {
	response, err := client.apiClient.TagsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsUpdateWithFormdataBody(ctx context.Context, id int, body api.TagsUpdateFormdataRequestBody) (*api.TagsUpdateResponse, error) {
	response, err := client.apiClient.TagsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TagsUpdate(ctx context.Context, id int, body api.TagsUpdateJSONRequestBody) (*api.TagsUpdateResponse, error) {
	response, err := client.apiClient.TagsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TasksList(ctx context.Context, params *api.TasksListParams) (*api.TasksListResponse, error) {
	response, err := client.apiClient.TasksListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TasksRetrieve(ctx context.Context, id int, params *api.TasksRetrieveParams) (*api.TasksRetrieveResponse, error) {
	response, err := client.apiClient.TasksRetrieveWithResponse(ctx, id, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TasksRunCreateWithBody(ctx context.Context, params *api.TasksRunCreateParams, contentType string, body io.Reader) (*api.TasksRunCreateResponse, error) {
	response, err := client.apiClient.TasksRunCreateWithBodyWithResponse(ctx, params, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TasksRunCreateWithFormdataBody(ctx context.Context, params *api.TasksRunCreateParams, body api.TasksRunCreateFormdataRequestBody) (*api.TasksRunCreateResponse, error) {
	response, err := client.apiClient.TasksRunCreateWithFormdataBodyWithResponse(ctx, params, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TasksRunCreate(ctx context.Context, params *api.TasksRunCreateParams, body api.TasksRunCreateJSONRequestBody) (*api.TasksRunCreateResponse, error) {
	response, err := client.apiClient.TasksRunCreateWithResponse(ctx, params, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TokenCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.TokenCreateResponse, error) {
	response, err := client.apiClient.TokenCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TokenCreateWithFormdataBody(ctx context.Context, body api.TokenCreateFormdataRequestBody) (*api.TokenCreateResponse, error) {
	response, err := client.apiClient.TokenCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TokenCreate(ctx context.Context, body api.TokenCreateJSONRequestBody) (*api.TokenCreateResponse, error) {
	response, err := client.apiClient.TokenCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TrashCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.TrashCreateResponse, error) {
	response, err := client.apiClient.TrashCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TrashCreateWithFormdataBody(ctx context.Context, body api.TrashCreateFormdataRequestBody) (*api.TrashCreateResponse, error) {
	response, err := client.apiClient.TrashCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TrashCreate(ctx context.Context, body api.TrashCreateJSONRequestBody) (*api.TrashCreateResponse, error) {
	response, err := client.apiClient.TrashCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) TrashList(ctx context.Context, params *api.TrashListParams) (*api.TrashListResponse, error) {
	response, err := client.apiClient.TrashListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UiSettingsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.UiSettingsCreateResponse, error) {
	response, err := client.apiClient.UiSettingsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UiSettingsCreateWithFormdataBody(ctx context.Context, body api.UiSettingsCreateFormdataRequestBody) (*api.UiSettingsCreateResponse, error) {
	response, err := client.apiClient.UiSettingsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UiSettingsCreate(ctx context.Context, body api.UiSettingsCreateJSONRequestBody) (*api.UiSettingsCreateResponse, error) {
	response, err := client.apiClient.UiSettingsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UiSettingsRetrieve(ctx context.Context) (*api.UiSettingsRetrieveResponse, error) {
	response, err := client.apiClient.UiSettingsRetrieveWithResponse(ctx, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.UsersCreateResponse, error) {
	response, err := client.apiClient.UsersCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersCreateWithFormdataBody(ctx context.Context, body api.UsersCreateFormdataRequestBody) (*api.UsersCreateResponse, error) {
	response, err := client.apiClient.UsersCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersCreate(ctx context.Context, body api.UsersCreateJSONRequestBody) (*api.UsersCreateResponse, error) {
	response, err := client.apiClient.UsersCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersDeactivateTotpCreate(ctx context.Context, id int) (*api.UsersDeactivateTotpCreateResponse, error) {
	response, err := client.apiClient.UsersDeactivateTotpCreateWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersDestroy(ctx context.Context, id int) (*api.UsersDestroyResponse, error) {
	response, err := client.apiClient.UsersDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersList(ctx context.Context, params *api.UsersListParams) (*api.UsersListResponse, error) {
	response, err := client.apiClient.UsersListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.UsersPartialUpdateResponse, error) {
	response, err := client.apiClient.UsersPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.UsersPartialUpdateFormdataRequestBody) (*api.UsersPartialUpdateResponse, error) {
	response, err := client.apiClient.UsersPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersPartialUpdate(ctx context.Context, id int, body api.UsersPartialUpdateJSONRequestBody) (*api.UsersPartialUpdateResponse, error) {
	response, err := client.apiClient.UsersPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersRetrieve(ctx context.Context, id int) (*api.UsersRetrieveResponse, error) {
	response, err := client.apiClient.UsersRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.UsersUpdateResponse, error) {
	response, err := client.apiClient.UsersUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersUpdateWithFormdataBody(ctx context.Context, id int, body api.UsersUpdateFormdataRequestBody) (*api.UsersUpdateResponse, error) {
	response, err := client.apiClient.UsersUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) UsersUpdate(ctx context.Context, id int, body api.UsersUpdateJSONRequestBody) (*api.UsersUpdateResponse, error) {
	response, err := client.apiClient.UsersUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.WorkflowActionsCreateResponse, error) {
	response, err := client.apiClient.WorkflowActionsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsCreateWithFormdataBody(ctx context.Context, body api.WorkflowActionsCreateFormdataRequestBody) (*api.WorkflowActionsCreateResponse, error) {
	response, err := client.apiClient.WorkflowActionsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsCreate(ctx context.Context, body api.WorkflowActionsCreateJSONRequestBody) (*api.WorkflowActionsCreateResponse, error) {
	response, err := client.apiClient.WorkflowActionsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsDestroy(ctx context.Context, id int) (*api.WorkflowActionsDestroyResponse, error) {
	response, err := client.apiClient.WorkflowActionsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsList(ctx context.Context, params *api.WorkflowActionsListParams) (*api.WorkflowActionsListResponse, error) {
	response, err := client.apiClient.WorkflowActionsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.WorkflowActionsPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowActionsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.WorkflowActionsPartialUpdateFormdataRequestBody) (*api.WorkflowActionsPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowActionsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsPartialUpdate(ctx context.Context, id int, body api.WorkflowActionsPartialUpdateJSONRequestBody) (*api.WorkflowActionsPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowActionsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsRetrieve(ctx context.Context, id int) (*api.WorkflowActionsRetrieveResponse, error) {
	response, err := client.apiClient.WorkflowActionsRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.WorkflowActionsUpdateResponse, error) {
	response, err := client.apiClient.WorkflowActionsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsUpdateWithFormdataBody(ctx context.Context, id int, body api.WorkflowActionsUpdateFormdataRequestBody) (*api.WorkflowActionsUpdateResponse, error) {
	response, err := client.apiClient.WorkflowActionsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowActionsUpdate(ctx context.Context, id int, body api.WorkflowActionsUpdateJSONRequestBody) (*api.WorkflowActionsUpdateResponse, error) {
	response, err := client.apiClient.WorkflowActionsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.WorkflowTriggersCreateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersCreateWithFormdataBody(ctx context.Context, body api.WorkflowTriggersCreateFormdataRequestBody) (*api.WorkflowTriggersCreateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersCreate(ctx context.Context, body api.WorkflowTriggersCreateJSONRequestBody) (*api.WorkflowTriggersCreateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersDestroy(ctx context.Context, id int) (*api.WorkflowTriggersDestroyResponse, error) {
	response, err := client.apiClient.WorkflowTriggersDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersList(ctx context.Context, params *api.WorkflowTriggersListParams) (*api.WorkflowTriggersListResponse, error) {
	response, err := client.apiClient.WorkflowTriggersListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.WorkflowTriggersPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.WorkflowTriggersPartialUpdateFormdataRequestBody) (*api.WorkflowTriggersPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersPartialUpdate(ctx context.Context, id int, body api.WorkflowTriggersPartialUpdateJSONRequestBody) (*api.WorkflowTriggersPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersRetrieve(ctx context.Context, id int) (*api.WorkflowTriggersRetrieveResponse, error) {
	response, err := client.apiClient.WorkflowTriggersRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.WorkflowTriggersUpdateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersUpdateWithFormdataBody(ctx context.Context, id int, body api.WorkflowTriggersUpdateFormdataRequestBody) (*api.WorkflowTriggersUpdateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowTriggersUpdate(ctx context.Context, id int, body api.WorkflowTriggersUpdateJSONRequestBody) (*api.WorkflowTriggersUpdateResponse, error) {
	response, err := client.apiClient.WorkflowTriggersUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*api.WorkflowsCreateResponse, error) {
	response, err := client.apiClient.WorkflowsCreateWithBodyWithResponse(ctx, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsCreateWithFormdataBody(ctx context.Context, body api.WorkflowsCreateFormdataRequestBody) (*api.WorkflowsCreateResponse, error) {
	response, err := client.apiClient.WorkflowsCreateWithFormdataBodyWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsCreate(ctx context.Context, body api.WorkflowsCreateJSONRequestBody) (*api.WorkflowsCreateResponse, error) {
	response, err := client.apiClient.WorkflowsCreateWithResponse(ctx, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsDestroy(ctx context.Context, id int) (*api.WorkflowsDestroyResponse, error) {
	response, err := client.apiClient.WorkflowsDestroyWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsList(ctx context.Context, params *api.WorkflowsListParams) (*api.WorkflowsListResponse, error) {
	response, err := client.apiClient.WorkflowsListWithResponse(ctx, params, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsPartialUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.WorkflowsPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowsPartialUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsPartialUpdateWithFormdataBody(ctx context.Context, id int, body api.WorkflowsPartialUpdateFormdataRequestBody) (*api.WorkflowsPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowsPartialUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsPartialUpdate(ctx context.Context, id int, body api.WorkflowsPartialUpdateJSONRequestBody) (*api.WorkflowsPartialUpdateResponse, error) {
	response, err := client.apiClient.WorkflowsPartialUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsRetrieve(ctx context.Context, id int) (*api.WorkflowsRetrieveResponse, error) {
	response, err := client.apiClient.WorkflowsRetrieveWithResponse(ctx, id, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsUpdateWithBody(ctx context.Context, id int, contentType string, body io.Reader) (*api.WorkflowsUpdateResponse, error) {
	response, err := client.apiClient.WorkflowsUpdateWithBodyWithResponse(ctx, id, contentType, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsUpdateWithFormdataBody(ctx context.Context, id int, body api.WorkflowsUpdateFormdataRequestBody) (*api.WorkflowsUpdateResponse, error) {
	response, err := client.apiClient.WorkflowsUpdateWithFormdataBodyWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) WorkflowsUpdate(ctx context.Context, id int, body api.WorkflowsUpdateJSONRequestBody) (*api.WorkflowsUpdateResponse, error) {
	response, err := client.apiClient.WorkflowsUpdateWithResponse(ctx, id, body, client.authenticateRequest())
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}
