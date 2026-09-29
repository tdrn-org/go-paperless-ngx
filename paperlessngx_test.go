// generated from api.ClientWithResponsesInterface
package paperlessngx_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-paperless-ngx"
	"github.com/tdrn-org/go-paperless-ngx/api"
	"github.com/tdrn-org/go-paperless-ngx/mock"
)

func TestAPI(t *testing.T) {
	mockServer := mock.Start()
	defer mockServer.Stop(t.Context())

	client, err := paperlessngx.NewClient(mockServer.APIURL(), mock.APIKey)
	require.NoError(t, err)

	{
		ctx := t.Context()
		params := &api.AcknowledgeTasksParams{}
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.AcknowledgeTasksWithBody(ctx, params, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.AcknowledgeTasksParams{}
		body := api.AcknowledgeTasksJSONRequestBody{}
		_, err := client.AcknowledgeTasks(ctx, params, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.BulkEditObjectsWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.BulkEditObjectsJSONRequestBody{}
		_, err := client.BulkEditObjects(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.BulkEditWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.BulkEditJSONRequestBody{}
		_, err := client.BulkEdit(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.ConfigDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.ConfigList(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.ConfigPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.ConfigPartialUpdateFormdataRequestBody{}
		_, err := client.ConfigPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.ConfigPartialUpdateJSONRequestBody{}
		_, err := client.ConfigPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.ConfigRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.ConfigUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.ConfigUpdateFormdataRequestBody{BarcodeTagMapping: "{}", UserArgs: "[]"}
		_, err := client.ConfigUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.ConfigUpdateJSONRequestBody{}
		_, err := client.ConfigUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.CorrespondentsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.CorrespondentsCreateFormdataRequestBody{}
		_, err := client.CorrespondentsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.CorrespondentsCreateJSONRequestBody{}
		_, err := client.CorrespondentsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.CorrespondentsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.CorrespondentsListParams{}
		_, err := client.CorrespondentsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.CorrespondentsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CorrespondentsPartialUpdateFormdataRequestBody{}
		_, err := client.CorrespondentsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CorrespondentsPartialUpdateJSONRequestBody{}
		_, err := client.CorrespondentsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.CorrespondentsRetrieveParams{}
		_, err := client.CorrespondentsRetrieve(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.CorrespondentsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CorrespondentsUpdateFormdataRequestBody{}
		_, err := client.CorrespondentsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CorrespondentsUpdateJSONRequestBody{}
		_, err := client.CorrespondentsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.CustomFieldsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.CustomFieldsCreateFormdataRequestBody{}
		_, err := client.CustomFieldsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.CustomFieldsCreateJSONRequestBody{}
		_, err := client.CustomFieldsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.CustomFieldsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.CustomFieldsListParams{}
		_, err := client.CustomFieldsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.CustomFieldsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CustomFieldsPartialUpdateFormdataRequestBody{}
		_, err := client.CustomFieldsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CustomFieldsPartialUpdateJSONRequestBody{}
		_, err := client.CustomFieldsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.CustomFieldsRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.CustomFieldsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CustomFieldsUpdateFormdataRequestBody{}
		_, err := client.CustomFieldsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.CustomFieldsUpdateJSONRequestBody{}
		_, err := client.CustomFieldsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := "id"
		_, err := client.DocumentShareLinks(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentTypesCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.DocumentTypesCreateFormdataRequestBody{}
		_, err := client.DocumentTypesCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.DocumentTypesCreateJSONRequestBody{}
		_, err := client.DocumentTypesCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.DocumentTypesDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.DocumentTypesListParams{}
		_, err := client.DocumentTypesList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentTypesPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentTypesPartialUpdateFormdataRequestBody{}
		_, err := client.DocumentTypesPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentTypesPartialUpdateJSONRequestBody{}
		_, err := client.DocumentTypesPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentTypesRetrieveParams{}
		_, err := client.DocumentTypesRetrieve(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentTypesUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentTypesUpdateFormdataRequestBody{}
		_, err := client.DocumentTypesUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentTypesUpdateJSONRequestBody{}
		_, err := client.DocumentTypesUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentsBulkDownloadCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.DocumentsBulkDownloadCreateJSONRequestBody{}
		_, err := client.DocumentsBulkDownloadCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.DocumentsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsDownloadRetrieveParams{}
		_, err := client.DocumentsDownloadRetrieve(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentsEmailCreateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentsEmailCreateFormdataRequestBody{}
		_, err := client.DocumentsEmailCreateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentsEmailCreateJSONRequestBody{}
		_, err := client.DocumentsEmailCreate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsHistoryListParams{}
		_, err := client.DocumentsHistoryList(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.DocumentsListParams{}
		_, err := client.DocumentsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.DocumentsMetadataRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.DocumentsNextAsnRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsNotesCreateParams{}
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentsNotesCreateWithBody(ctx, id, params, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsNotesCreateParams{}
		body := api.DocumentsNotesCreateFormdataRequestBody{}
		_, err := client.DocumentsNotesCreateWithFormdataBody(ctx, id, params, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsNotesCreateParams{}
		body := api.DocumentsNotesCreateJSONRequestBody{}
		_, err := client.DocumentsNotesCreate(ctx, id, params, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsNotesDestroyParams{}
		_, err := client.DocumentsNotesDestroy(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsNotesListParams{}
		_, err := client.DocumentsNotesList(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentsPartialUpdateFormdataRequestBody{}
		_, err := client.DocumentsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentsPartialUpdateJSONRequestBody{}
		_, err := client.DocumentsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "multipart/form-data; boundary=bndry"
		body := strings.NewReader("--bndry\r\nContent-Disposition: form-data; name=\"document\"\r\n\r\ndocument\r\n--bndry--\r\n")
		_, err := client.DocumentsPostDocumentCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.DocumentsPreviewRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.DocumentsRetrieveParams{}
		_, err := client.DocumentsRetrieve(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentsSelectionDataCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.DocumentsSelectionDataCreateJSONRequestBody{}
		_, err := client.DocumentsSelectionDataCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.DocumentsSuggestionsRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.DocumentsThumbRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.DocumentsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentsUpdateFormdataRequestBody{}
		_, err := client.DocumentsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.DocumentsUpdateJSONRequestBody{}
		_, err := client.DocumentsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.EmailDocumentsWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.EmailDocumentsFormdataRequestBody{}
		_, err := client.EmailDocumentsWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.EmailDocumentsJSONRequestBody{}
		_, err := client.EmailDocuments(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.GroupsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.GroupsCreateFormdataRequestBody{}
		_, err := client.GroupsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.GroupsCreateJSONRequestBody{}
		_, err := client.GroupsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.GroupsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.GroupsListParams{}
		_, err := client.GroupsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.GroupsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.GroupsPartialUpdateFormdataRequestBody{}
		_, err := client.GroupsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.GroupsPartialUpdateJSONRequestBody{}
		_, err := client.GroupsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.GroupsRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.GroupsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.GroupsUpdateFormdataRequestBody{}
		_, err := client.GroupsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.GroupsUpdateJSONRequestBody{}
		_, err := client.GroupsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.LogsList(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailAccountProcessWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailAccountProcessFormdataRequestBody{}
		_, err := client.MailAccountProcessWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailAccountProcessJSONRequestBody{}
		_, err := client.MailAccountProcess(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailAccountTestWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.MailAccountTestFormdataRequestBody{}
		_, err := client.MailAccountTestWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.MailAccountTestJSONRequestBody{}
		_, err := client.MailAccountTest(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailAccountsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.MailAccountsCreateFormdataRequestBody{}
		_, err := client.MailAccountsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.MailAccountsCreateJSONRequestBody{}
		_, err := client.MailAccountsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.MailAccountsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.MailAccountsListParams{}
		_, err := client.MailAccountsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailAccountsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailAccountsPartialUpdateFormdataRequestBody{}
		_, err := client.MailAccountsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailAccountsPartialUpdateJSONRequestBody{}
		_, err := client.MailAccountsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.MailAccountsRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailAccountsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailAccountsUpdateFormdataRequestBody{}
		_, err := client.MailAccountsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailAccountsUpdateJSONRequestBody{}
		_, err := client.MailAccountsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailRulesCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.MailRulesCreateFormdataRequestBody{}
		_, err := client.MailRulesCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.MailRulesCreateJSONRequestBody{}
		_, err := client.MailRulesCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.MailRulesDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.MailRulesListParams{}
		_, err := client.MailRulesList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailRulesPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailRulesPartialUpdateFormdataRequestBody{}
		_, err := client.MailRulesPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailRulesPartialUpdateJSONRequestBody{}
		_, err := client.MailRulesPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.MailRulesRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.MailRulesUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailRulesUpdateFormdataRequestBody{}
		_, err := client.MailRulesUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.MailRulesUpdateJSONRequestBody{}
		_, err := client.MailRulesUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.OauthCallbackRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.ProcessedMailBulkDeleteCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ProcessedMailBulkDeleteCreateFormdataRequestBody{}
		_, err := client.ProcessedMailBulkDeleteCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ProcessedMailBulkDeleteCreateJSONRequestBody{}
		_, err := client.ProcessedMailBulkDeleteCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.ProcessedMailListParams{}
		_, err := client.ProcessedMailList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.ProcessedMailRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.ProfileDisconnectSocialAccountCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ProfileDisconnectSocialAccountCreateJSONRequestBody{}
		_, err := client.ProfileDisconnectSocialAccountCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.ProfileGenerateAuthTokenCreate(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.ProfilePartialUpdateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ProfilePartialUpdateFormdataRequestBody{}
		_, err := client.ProfilePartialUpdateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ProfilePartialUpdateJSONRequestBody{}
		_, err := client.ProfilePartialUpdate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.ProfileRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.ProfileSocialAccountProvidersRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.ProfileTotpCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ProfileTotpCreateJSONRequestBody{}
		_, err := client.ProfileTotpCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.ProfileTotpDestroy(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.ProfileTotpRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.RemoteVersionRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := "id"
		params := &api.RetrieveLogParams{}
		_, err := client.RetrieveLog(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.SavedViewsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.SavedViewsCreateFormdataRequestBody{}
		_, err := client.SavedViewsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.SavedViewsCreateJSONRequestBody{}
		_, err := client.SavedViewsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.SavedViewsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.SavedViewsListParams{}
		_, err := client.SavedViewsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.SavedViewsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.SavedViewsPartialUpdateFormdataRequestBody{}
		_, err := client.SavedViewsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.SavedViewsPartialUpdateJSONRequestBody{}
		_, err := client.SavedViewsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.SavedViewsRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.SavedViewsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.SavedViewsUpdateFormdataRequestBody{}
		_, err := client.SavedViewsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.SavedViewsUpdateJSONRequestBody{}
		_, err := client.SavedViewsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.SearchAutocompleteListParams{}
		_, err := client.SearchAutocompleteList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.SearchRetrieveParams{}
		_, err := client.SearchRetrieve(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.ShareLinksCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ShareLinksCreateFormdataRequestBody{}
		_, err := client.ShareLinksCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.ShareLinksCreateJSONRequestBody{}
		_, err := client.ShareLinksCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.ShareLinksDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.ShareLinksListParams{}
		_, err := client.ShareLinksList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.ShareLinksRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.StatisticsRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.StatusRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.StoragePathsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.StoragePathsCreateFormdataRequestBody{}
		_, err := client.StoragePathsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.StoragePathsCreateJSONRequestBody{}
		_, err := client.StoragePathsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.StoragePathsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.StoragePathsListParams{}
		_, err := client.StoragePathsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.StoragePathsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.StoragePathsPartialUpdateFormdataRequestBody{}
		_, err := client.StoragePathsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.StoragePathsPartialUpdateJSONRequestBody{}
		_, err := client.StoragePathsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.StoragePathsRetrieveParams{}
		_, err := client.StoragePathsRetrieve(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.StoragePathsTestCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.StoragePathsTestCreateFormdataRequestBody{}
		_, err := client.StoragePathsTestCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.StoragePathsTestCreateJSONRequestBody{}
		_, err := client.StoragePathsTestCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.StoragePathsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.StoragePathsUpdateFormdataRequestBody{}
		_, err := client.StoragePathsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.StoragePathsUpdateJSONRequestBody{}
		_, err := client.StoragePathsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.TagsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.TagsCreateFormdataRequestBody{}
		_, err := client.TagsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.TagsCreateJSONRequestBody{}
		_, err := client.TagsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.TagsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.TagsListParams{}
		_, err := client.TagsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.TagsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.TagsPartialUpdateFormdataRequestBody{}
		_, err := client.TagsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.TagsPartialUpdateJSONRequestBody{}
		_, err := client.TagsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.TagsRetrieveParams{}
		_, err := client.TagsRetrieve(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.TagsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.TagsUpdateFormdataRequestBody{}
		_, err := client.TagsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.TagsUpdateJSONRequestBody{}
		_, err := client.TagsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.TasksListParams{}
		_, err := client.TasksList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		params := &api.TasksRetrieveParams{}
		_, err := client.TasksRetrieve(ctx, id, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.TasksRunCreateParams{}
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.TasksRunCreateWithBody(ctx, params, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.TasksRunCreateParams{}
		body := api.TasksRunCreateFormdataRequestBody{}
		_, err := client.TasksRunCreateWithFormdataBody(ctx, params, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.TasksRunCreateParams{}
		body := api.TasksRunCreateJSONRequestBody{}
		_, err := client.TasksRunCreate(ctx, params, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.TokenCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.TokenCreateFormdataRequestBody{}
		_, err := client.TokenCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.TokenCreateJSONRequestBody{}
		_, err := client.TokenCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.TrashCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.TrashCreateFormdataRequestBody{}
		_, err := client.TrashCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.TrashCreateJSONRequestBody{}
		_, err := client.TrashCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.TrashListParams{}
		_, err := client.TrashList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.UiSettingsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.UiSettingsCreateFormdataRequestBody{}
		_, err := client.UiSettingsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.UiSettingsCreateJSONRequestBody{}
		_, err := client.UiSettingsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		_, err := client.UiSettingsRetrieve(ctx)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.UsersCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.UsersCreateFormdataRequestBody{}
		_, err := client.UsersCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.UsersCreateJSONRequestBody{}
		_, err := client.UsersCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.UsersDeactivateTotpCreate(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.UsersDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.UsersListParams{}
		_, err := client.UsersList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.UsersPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.UsersPartialUpdateFormdataRequestBody{}
		_, err := client.UsersPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.UsersPartialUpdateJSONRequestBody{}
		_, err := client.UsersPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.UsersRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.UsersUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.UsersUpdateFormdataRequestBody{}
		_, err := client.UsersUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.UsersUpdateJSONRequestBody{}
		_, err := client.UsersUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowActionsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.WorkflowActionsCreateFormdataRequestBody{}
		_, err := client.WorkflowActionsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.WorkflowActionsCreateJSONRequestBody{}
		_, err := client.WorkflowActionsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.WorkflowActionsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.WorkflowActionsListParams{}
		_, err := client.WorkflowActionsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowActionsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowActionsPartialUpdateFormdataRequestBody{}
		_, err := client.WorkflowActionsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowActionsPartialUpdateJSONRequestBody{}
		_, err := client.WorkflowActionsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.WorkflowActionsRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowActionsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowActionsUpdateFormdataRequestBody{}
		_, err := client.WorkflowActionsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowActionsUpdateJSONRequestBody{}
		_, err := client.WorkflowActionsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowTriggersCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.WorkflowTriggersCreateFormdataRequestBody{}
		_, err := client.WorkflowTriggersCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.WorkflowTriggersCreateJSONRequestBody{}
		_, err := client.WorkflowTriggersCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.WorkflowTriggersDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.WorkflowTriggersListParams{}
		_, err := client.WorkflowTriggersList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowTriggersPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowTriggersPartialUpdateFormdataRequestBody{}
		_, err := client.WorkflowTriggersPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowTriggersPartialUpdateJSONRequestBody{}
		_, err := client.WorkflowTriggersPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.WorkflowTriggersRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowTriggersUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowTriggersUpdateFormdataRequestBody{}
		_, err := client.WorkflowTriggersUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowTriggersUpdateJSONRequestBody{}
		_, err := client.WorkflowTriggersUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowsCreateWithBody(ctx, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.WorkflowsCreateFormdataRequestBody{}
		_, err := client.WorkflowsCreateWithFormdataBody(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		body := api.WorkflowsCreateJSONRequestBody{}
		_, err := client.WorkflowsCreate(ctx, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.WorkflowsDestroy(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		params := &api.WorkflowsListParams{}
		_, err := client.WorkflowsList(ctx, params)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowsPartialUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowsPartialUpdateFormdataRequestBody{}
		_, err := client.WorkflowsPartialUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowsPartialUpdateJSONRequestBody{}
		_, err := client.WorkflowsPartialUpdate(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		_, err := client.WorkflowsRetrieve(ctx, id)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		contentType := "application/json"
		body := strings.NewReader("{}")
		_, err := client.WorkflowsUpdateWithBody(ctx, id, contentType, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowsUpdateFormdataRequestBody{}
		_, err := client.WorkflowsUpdateWithFormdataBody(ctx, id, body)
		require.NoError(t, err)
	}
	{
		ctx := t.Context()
		id := 0
		body := api.WorkflowsUpdateJSONRequestBody{}
		_, err := client.WorkflowsUpdate(ctx, id, body)
		require.NoError(t, err)
	}
}
